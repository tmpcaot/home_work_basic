package main

import (
	"fmt"
	"sync"
)

// Counter хранит общее количество выполненых задач.
type Counter struct {
	count int
	mux   sync.Mutex
}

// Increment увеличивает значение счетчика на единицу.
func (c *Counter) Increment() {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.count++
}

// Get возвращает текущее значение счетчика.
func (c *Counter) Get() int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.count
}

// Worker выполняет задание и сообщает о завершении.
func worker(id int, counter *Counter, wg *sync.WaitGroup) {
	defer wg.Done()
	counter.Increment()
	fmt.Printf("Worker %d finished its task.\n", id)
}

func main() {
	var wg sync.WaitGroup
	counter := &Counter{}

	numWorkers := 5
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go worker(i, counter, &wg)
	}

	wg.Wait()
	fmt.Printf("Total tasks completed: %d\n", counter.Get())
}
