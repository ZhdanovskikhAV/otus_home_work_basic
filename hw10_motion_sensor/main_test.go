package main

import (
	"fmt"
	"sync"
	"testing"
)

// Тест для функции sensorReadings.
func TestSensorReadings(t *testing.T) {
	dataChannel := make(chan float64)
	processedChannel := make(chan float64)
	var wg sync.WaitGroup
	wg.Add(1)

	go processData(dataChannel, processedChannel, &wg)

	// Отправляем значения в канал.
	go func() {
		for _, value := range []float64{16.0, 48.0, 70.0} {
			dataChannel <- value
			fmt.Printf("Считано значение: %.2f\n", value)
		}
		close(dataChannel) // Закрываем канал после отправки всех значений.
	}()

	// Создаем отдельную горутину для закрытия processedChannel после завершения обработки.
	go func() {
		wg.Wait()               // Ожидаем завершения горутины.
		close(processedChannel) // Закрываем канал обработанных данных.
	}()
	// Предварительно выделяем память для processedValues.
	processedValues := make([]float64, 0, 1) // Предполагаем, что будет 1 среднее значение.

	// Получаем все обработанные значения.
	for avg := range processedChannel {
		processedValues = append(processedValues, avg)
	}

	// Проверяем, что мы получили правильное среднее значение.
	expectedAvg := (16.0 + 48.0 + 70.0) / 3 // Среднее от трех значений
	if len(processedValues) != 1 || processedValues[0] != expectedAvg {
		t.Errorf("Ожидалось среднее: %.2f, но получено: %.2f", expectedAvg, processedValues)
	}
}

// Тест для функции processData.
func TestProcessData(t *testing.T) {
	dataChannel := make(chan float64)
	processedChannel := make(chan float64)
	var wg sync.WaitGroup
	wg.Add(1)

	// Запускаем горутину для обработки данных.
	go processData(dataChannel, processedChannel, &wg)

	// Отправляем 10 значений для обработки.
	for i := 0; i < 10; i++ {
		dataChannel <- float64(i)
	}
	close(dataChannel) // Закрываем канал входных данных.

	// Создаем отдельную горутину для закрытия processedChannel после завершения обработки.
	go func() {
		wg.Wait()               // Ожидаем завершения горутины.
		close(processedChannel) // Закрываем канал обработанных данных.
	}()
	// Предварительно выделяем память для processedValues.
	processedValues := make([]float64, 0, 1) // Предполагаем, что будет 1 среднее значение.

	// Получаем все обработанные значения.
	for avg := range processedChannel {
		processedValues = append(processedValues, avg)
	}

	// Проверяем, что мы получили правильное среднее значение.
	expectedAvg := 4.5 // Среднее от 0 до 9
	if len(processedValues) != 1 || processedValues[0] != expectedAvg {
		t.Errorf("Ожидалось среднее: %.2f, но получено: %.2f", expectedAvg, processedValues)
	}
}
