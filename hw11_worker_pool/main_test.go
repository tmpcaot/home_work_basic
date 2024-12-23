package main

import (
	"os"
	"sync"
	"testing"
)

// проверяем работу счетчика.
func TestCounter(t *testing.T) {
	counter := &Counter{}

	// Увеличить счетчик 5 раз.
	for i := 0; i < 5; i++ {
		counter.Increment()
	}

	// Проверить, что счетчик равен 5.
	if counter.Get() != 5 {
		t.Errorf("Expected count to be 5, but got %d", counter.Get())
	}
}

// проверяем инкрементацию счетчика несколькими горутинами.
func TestConcurrentIncrement(t *testing.T) {
	counter := &Counter{}
	var wg sync.WaitGroup

	numWorkers := 10
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()

	// проверяем, что счетчик равен количеству запущенных горутин.
	if counter.Get() != numWorkers {
		t.Errorf("Expected count to be %d, but got %d", numWorkers, counter.Get())
	}
}

// проверяем выполнение задания каждым воркером.
func TestWorker(t *testing.T) {
	counter := &Counter{}
	var wg sync.WaitGroup

	numWorkers := 5
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go worker(i, counter, &wg)
	}

	wg.Wait()

	// проверяем, что счетчик равен количеству запущенных горутин.
	if counter.Get() != numWorkers {
		t.Errorf("Expected count to be %d, but got %d", numWorkers, counter.Get())
	}
}

func TestMain(m *testing.M) {
	retCode := m.Run()

	os.Exit(retCode)
}
