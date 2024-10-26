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
func (c *Counter) Increment() {
	c.mu.Lock()
	c.value++
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
	messageChannel := make(chan string, numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < incrementsPerGoroutine; j++ {
				c.Increment()
			}
			messageChannel <- fmt.Sprintf("Goroutine %d finished its work.", id)
		}(i)
	}

	go func() {
		wg.Wait()
		close(messageChannel)
	}()

	for message := range messageChannel {
		fmt.Println(message)
	}
}

func main() {
	counter := &Counter{}
	const numGoroutines = 5
	const incrementsPerGoroutine = 1000

	RunGoroutines(counter, numGoroutines, incrementsPerGoroutine)
	fmt.Printf("Final counter value: %d\n", counter.GetValue())
}
