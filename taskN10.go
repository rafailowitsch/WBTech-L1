package main

import (
	"fmt"
	"math"
	"sort"
)

// функция taskN10 демонстрирует группировку температурных значений с шагом 10 градусов
func taskN10() {
	sl := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	m := GroupTemperatures(sl)

	fmt.Println(m)
}

// функция GroupTemperatures группирует температурные значения по диапазонам с шагом 10 градусов
func GroupTemperatures(temp []float64) map[int][]float64 {
	m := make(map[int][]float64)
	index := 0

	sort.Float64s(temp) // сортируем температуры по возрастанию

	for index < len(temp) {
		series := sign(temp[index]) * int(math.Floor(math.Abs(temp[index])/10)*10)

		if _, exists := m[series]; !exists {
			m[series] = make([]float64, 0, 10)
		}

		for ; index < len(temp) &&
			(series-10) <= int(math.Ceil(temp[index])) &&
			int(math.Ceil(temp[index])) <= (series+10); index++ {
			m[series] = append(m[series], temp[index])
		}
	}

	return m
}

// функция sign возвращает знак числа
func sign(num float64) int {
	if num > 0 {
		return 1
	} else if num < 0 {
		return -1
	}
	return 0
}
