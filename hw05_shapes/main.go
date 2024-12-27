package main

import (
	"errors"
	"fmt"
	"math"
)

// Интерфейс Shape с методом для вычисления площади.
type Shape interface {
	Area() float64
}

// Структура для круга.
type Circle struct {
	Radius float64
}

// Метод поиска площади для круга.
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Структура для прямоугольника.
type Rectangle struct {
	Width  float64
	Height float64
}

// Метод поиска площади для прямоугольника.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Структура для треугольника.
type Triangle struct {
	Base   float64
	Height float64
}

// Метод поиска площади  для треугольника.
func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

// Функция для вычисления площади.
func calculateArea(s any) (float64, error) {
	shape, ok := s.(Shape)
	if !ok {
		return 0, errors.New("переданный объект не является фигурой")
	}
	return shape.Area(), nil
}

func main() {
	// Создание объектов разных типов.
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 10, Height: 5}
	triangle := Triangle{Base: 8, Height: 6}

	// Вычисление и вывод площадей.
	shapes := []any{circle, rectangle, triangle}

	for _, shape := range shapes {
		area, err := calculateArea(shape)
		if err != nil {
			fmt.Println("Ошибка:", err)
		} else {
			switch s := shape.(type) {
			case Circle:
				fmt.Printf("Круг: радиус %.1f\nПлощадь: %.10f\n", s.Radius, area)
			case Rectangle:
				fmt.Printf("Прямоугольник: ширина %.1f, высота %.1f\nПлощадь: %.1f\n", s.Width, s.Height, area)
			case Triangle:
				fmt.Printf("Треугольник: основание %.1f, высота %.1f\nПлощадь: %.1f\n", s.Base, s.Height, area)
			}
		}
	}
}
