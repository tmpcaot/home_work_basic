package main

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
	"time"
)

// имитируем работу сенсора, генерируя случайные числа и отправляя их в канал.
func Sensor(max int64, delay time.Duration) chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		timeout := time.After(delay) //   таймаут для остановки работы сенсора.
		for {
			select {
			case <-timeout:
				return
			default:
				r, err := rand.Int(rand.Reader, big.NewInt(max))
				if err != nil {
					return
				}
				select {
				case c <- int(r.Int64()):
				case <-timeout:
					return
				}
			}
		}
	}()
	return c
}

// обрабатываем поток данных из канала, вычисляем среднее значение каждые 10 элементов и передаем результат в другой канал.
func Reader(depth int, inputCh chan int) chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		counter := 0
		storage := 0
		for a := range inputCh {
			storage += a
			counter++
			if counter == depth {
				mid := math.Ceil(float64(storage) / float64(depth))
				c <- int(mid)
				counter = 0
				storage = 0
			}
		}
	}()
	return c
}

func main() {
	emulationSensor := Sensor(1000, time.Minute)
	readData := Reader(10, emulationSensor)

	for output := range readData {
		fmt.Println("Среднее значение:", output)
	}
}
