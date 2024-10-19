package main

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// Тест для функции sensorReadings.
func TestSensorReadings(t *testing.T) {
	dataChannel := make(chan float64)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second) // Добавили тайм-аут.
	defer cancel()

	go sensorReadings(ctx, dataChannel)

	// Проверяем, что мы получаем значения из канала.
	// Задаем тайм-аут, чтобы тест не зацикливался.
	receivedValues := []float64{}
	finish := time.After(3 * time.Second)

	for {
		select {
		case data, ok := <-dataChannel:
			if !ok {
				return // Выход из цикла, если канал закрыт.
			}
			fmt.Printf("Считано значение: %.2f\n", data)
			receivedValues = append(receivedValues, data)

		case <-finish:
			// Прерываем, если время вышло.
			if len(receivedValues) == 0 {
				t.Error("Не считано ни одного значения")
			}
			return // Если получены значения, завершаем тест.
		}
	}
}

// Тест для функции processData.
func TestProcessData(t *testing.T) {
	dataChannel := make(chan float64)
	processedChannel := make(chan float64)
	go processData(dataChannel, processedChannel)

	// Отправляем 10 значений для обработки.
	for i := 0; i < 10; i++ {
		dataChannel <- float64(i)
	}
	close(dataChannel) // Закрываем канал входных данных.

	avg := <-processedChannel // Получаем среднее значение
	expectedAvg := 4.5        // Среднее от 0 до 9
	if avg != expectedAvg {
		t.Errorf("Ожидалось среднее: %.2f, но получено: %.2f", expectedAvg, avg)
	}
}
