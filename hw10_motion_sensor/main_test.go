package main

import (
	"testing"
	"time"
)

// проверяем корректность работы функции Sensor.
func TestSensor1(t *testing.T) {
	max := 1000
	emulationSensor := Sensor(int64(max), time.Minute)

	for i := 0; i < max; i++ {
		output, ok := <-emulationSensor
		if !ok || output > max {
			t.Error(i)
			break
		}
	}
}

// проверяем завершение работы сенсора через установленное время.
func TestSensorTimeout(t *testing.T) {
	max := 10000000000000
	delay := time.Second
	emulationSensor := Sensor(int64(max), delay)
	start := time.Now()
	stop := start.Add(2 * delay)

	for i := 0; i < max; i++ {
		_, ok := <-emulationSensor
		if time.Now().After(stop) {
			if ok {
				t.Error(ok, i)
				break
			}
			break
		}
	}
}

// проверяем корректность работы функции Reader.
func TestReader(t *testing.T) {
	delay := time.Second
	emulationSensor := Sensor(100000, 2*delay)
	readData := Reader(100, emulationSensor)
	start := time.Now()
	stop := start.Add(delay)

	for {
		if time.Now().After(stop) {
			break
		}
		_, ok := <-emulationSensor
		if !ok {
			t.Error("emulationSensor error")
			break
		}
		_, ok = <-readData
		if !ok {
			t.Error("readData error")
			break
		}
	}
}
