package main

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

// Тест для функции sensorReadings.
func TestSensorReadings(t *testing.T) {
	dataChannel := make(chan float64)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second) // Добавили тайм-аут.
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)
	go sensorReadings(ctx, dataChannel, &wg)

	// Проверяем, что мы получаем значения из канала.
	// Задаем тайм-аут, чтобы тест не зацикливался.
	receivedValues := []float64{}
	finish := time.After(3 * time.Second)

	for {
		select {
		case data, ok := <-dataChannel:
			if !ok {
				// Если канал закрыт, проверяем, что есть полученные значения.
				if len(receivedValues) == 0 {
					t.Error("Не считано ни одного значения")
				}
				wg.Wait() // Ожидаем завершения горутины.
				return    // Выход из цикла, если канал закрыт.
			}
			fmt.Printf("Считано значение: %.2f\n", data)
			receivedValues = append(receivedValues, data)

		case <-finish:
			// Прерываем, если время вышло.
			if len(receivedValues) == 0 {
				t.Error("Не считано ни одного значения")
			}
			wg.Wait() // Ожидаем завершения горутины.
			return    // Если получены значения, завершаем тест.
		}
	}
}

// Тест для функции processData.
func TestProcessData(t *testing.T) {
	dataChannel := make(chan float64)
	processedChannel := make(chan float64)
	var wg sync.WaitGroup
	wg.Add(1)
	go processData(dataChannel, processedChannel, &wg)

	// Отправляем 10 значений для обработки.
	for i := 0; i < 10; i++ {
		dataChannel <- float64(i)
	}
	close(dataChannel) // Закрываем канал входных данных.

	// Получаем все обработанные значения.
	go func() {
		wg.Wait()               // Ожидаем завершения горутины.
		close(processedChannel) // Закрываем канал обработанных данных.
	}()

	// Получаем среднее значение.
	avg := <-processedChannel
	expectedAvg := 4.5 // Среднее от 0 до 9
	if avg != expectedAvg {
		t.Errorf("Ожидалось среднее: %.2f, но получено: %.2f", expectedAvg, avg)
	}
}
