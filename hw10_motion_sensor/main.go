package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"sync"
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
func sensorReadings(ctx context.Context, ch chan<- float64, wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшаем счетчик WaitGroup при завершении работы горутины.
	defer close(ch) // Закрываем канал при завершении работы горутины.

	// Создаем таймер, который будет срабатывать через 1 минуту.
	timer := time.NewTimer(1 * time.Minute)
	defer timer.Stop() // Останавливаем таймер при завершении функции.

	for {
		select {
		case <-ctx.Done():
			// Контекст завершен, выходим из функции.
			fmt.Println("Горутина считывания завершена по просьбе!")
			return
		case <-timer.C:
			// Таймер сработал, выходим из функции.
			fmt.Println("Горутина считывания завершена по времени!")
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
func processData(ch <-chan float64, processedCh chan<- float64, wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшаем счетчик WaitGroup при завершении работы горутины.

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

	var wg sync.WaitGroup

	wg.Add(1)
	go sensorReadings(ctx, dataChannel, &wg) // Запускаем горутину для считывания данных.

	wg.Add(1)
	go processData(dataChannel, processedChannel, &wg) // Запускаем горутину для обработки данных.

	// Главная горутина читает обработанные данные и выводит их на экран.
	go func() {
		for avg := range processedChannel {
			fmt.Printf("Среднее арифметическое: %.2f\n", avg)
		}
	}()

	// Ждем завершения горутины считывания.
	wg.Wait()
	// После завершения считывания, отменяем контекст, чтобы завершить обработку.
	cancel()

	// Ждем завершения горутины обработки данных.
	wg.Wait()

	fmt.Println("Все горутины завершены. Программа завершена.")
}
