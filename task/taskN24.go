package task

import (
	"fmt"
	"math"
)

// структура Point представляет точку в 2D пространстве
type Point struct {
	x, y float64
}

// конструктор NewPoint создает и возвращает новый экземпляр Point
func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

// функция Distance вычисляет расстояние между двумя точками
func Distance(p1, p2 *Point) float64 {
	result := math.Sqrt(math.Pow(p2.x-p1.x, 2) + math.Pow(p2.y-p1.y, 2))
	return result
}

// функция taskN24 демонстрирует использование структуры Point и функции Distance
func TaskN24() {
	p1 := NewPoint(4, 5)
	p2 := NewPoint(0, 0)

	dis := Distance(p1, p2)

	fmt.Printf("Distance between p1 and p2 is %.2f\n", dis)
}
