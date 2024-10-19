package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

// Функция для генерации безопасного случайного числа в заданном диапазоне.
func CryptoRand(limit int) uint64 {
	safeNum, err := rand.Int(rand.Reader, big.NewInt(int64(limit)))
	if err != nil {
		return 0
	}
	return uint64(safeNum.Int64())
}

// Функция для генерации случайных данных.
func sensorReadings(ctx context.Context, ch chan<- float64) {
	defer close(ch) // Закрываем канал при завершении работы горутины.

	for {
		select {
		case <-ctx.Done():
			// Контекст завершен, выходим из функции.
			fmt.Println("Горутина считывания завершена по просьбе!")
			return
		default:
			// Генерация случайного значения в диапазоне [0, 100).
			data := float64(CryptoRand(100))
			ch <- data
			time.Sleep(1 * time.Second) // Пауза между считываниями.
		}
	}
}

// Функция для обработки данных.
func processData(ch <-chan float64, processedCh chan<- float64) {
	defer close(processedCh) // Закрываем канал при завершении обработки.

	var sum float64
	var count int

	for data := range ch {
		sum += data
		count++

		// Каждые 10 считанных значений.
		if count == 10 {
			average := sum / float64(count)
			processedCh <- average
			sum = 0   // Сбрасываем сумму.
			count = 0 // Сбрасываем счетчик.
		}
	}

	// Если есть оставшиеся данные, вычисляем среднее и отправляем.
	if count > 0 {
		average := sum / float64(count)
		processedCh <- average
	}
}

// Основная функция.
func main() {
	dataChannel := make(chan float64)
	processedChannel := make(chan float64)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Вызываем cancel при завершении main.

	go sensorReadings(ctx, dataChannel)           // Запускаем горутину для считывания данных.
	go processData(dataChannel, processedChannel) // Запускаем горутину для обработки данных.

	// Главная горутина читает обработанные данные и выводит их на экран.
	for avg := range processedChannel {
		fmt.Printf("Среднее арифметическое: %.2f\n", avg)
	}

	// Завершаем счетчик.
	cancel() // Завершаем контекст, что вызывает завершение другой горутины.
}
