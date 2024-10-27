package main

import (
	"fmt"
	"sync"
)

// Counter - структура, представляющая общий счетчик.
type Counter struct {
	value int
	mu    sync.Mutex
}

// Increment - метод для инкрементации счетчика.
func (c *Counter) Increment(increment int) {
	c.mu.Lock()
	c.value += increment
	c.mu.Unlock()
}

// GetValue - метод для получения текущего значения счетчика.
func (c *Counter) GetValue() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

// RunGoroutines - функция, запускающая горутины для инкрементации счетчика.
func RunGoroutines(c *Counter, numGoroutines, incrementsPerGoroutine int) {
	var wg sync.WaitGroup

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// Накапливаем временный результат
			tempSum := 0
			for j := 0; j < incrementsPerGoroutine; j++ {
				tempSum++
			}
			// Добавляем временный результат к общему счетчику за одну блокировку
			c.Increment(tempSum)
		}()
	}

	wg.Wait()
}

func main() {
	counter := &Counter{}
	const numGoroutines = 5
	const incrementsPerGoroutine = 1000

	RunGoroutines(counter, numGoroutines, incrementsPerGoroutine)
	fmt.Printf("Final counter value: %d\n", counter.GetValue())
}
