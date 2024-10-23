package day3

import (
	"image"
	"math"
)

func moveRight(currentPoint image.Point) image.Point {
	return image.Pt(currentPoint.X+1, currentPoint.Y)
}
func moveLeft(currentPoint image.Point) image.Point {
	return image.Pt(currentPoint.X-1, currentPoint.Y)
}
func moveUp(currentPoint image.Point) image.Point {
	return image.Pt(currentPoint.X, currentPoint.Y+1)
}
func moveDown(currentPoint image.Point) image.Point {
	return image.Pt(currentPoint.X, currentPoint.Y-1)
}

func P1(input int) int {
	return generateSpiral(func(previousCellValue int, currentPosition image.Point) int {
		return previousCellValue + 1
	}, func(cellValue int, currentPosition image.Point) (int, bool) {
		if cellValue == input {
			return int(math.Abs(float64(currentPosition.X)) + math.Abs(float64(currentPosition.Y))), true
		}
		return 0, false
	})
}

func P2(input int) int {
	cells := map[image.Point]int{}
	cells[image.Pt(0, 0)] = 1
	return generateSpiral(func(previousCellValue int, currentPosition image.Point) int {
		newCellValue := 0
		neighbors := generateNeighbors(currentPosition)
		for _, neighbor := range neighbors {
			if value, present := cells[neighbor]; present {
				newCellValue += value
			}
		}
		cells[currentPosition] = newCellValue
		return newCellValue
	},
		func(cellValue int, currentPosition image.Point) (int, bool) {
			if cellValue > input {
				return cellValue, true
			}
			return 0, false
		})
}

func generateSpiral(valueCalculator func(int, image.Point) int, exitCondition func(int, image.Point) (int, bool)) int {
	currentPoint := image.Pt(0, 0)
	n := 1
	cellValue := 1
	rightCount := 0
	upCount := 0
	leftCount := 0
	downCount := 0
	for {
		if value, shouldExit := exitCondition(cellValue, currentPoint); shouldExit {
			return value
		}

		switch {
		case rightCount < n:
			rightCount += 1
			currentPoint = moveRight(currentPoint)
			cellValue = valueCalculator(cellValue, currentPoint)
		case upCount < n:
			upCount += 1
			currentPoint = moveUp(currentPoint)
			cellValue = valueCalculator(cellValue, currentPoint)
		case leftCount < n+1:
			leftCount += 1
			currentPoint = moveLeft(currentPoint)
			cellValue = valueCalculator(cellValue, currentPoint)
		case downCount < n+1:
			downCount += 1
			currentPoint = moveDown(currentPoint)
			cellValue = valueCalculator(cellValue, currentPoint)
		default:
			n += 2
			rightCount = 0
			upCount = 0
			leftCount = 0
			downCount = 0
		}
	}
}

func generateNeighbors(currentPoint image.Point) []image.Point {
	neighbors := make([]image.Point, 0)
	neighbors = append(neighbors, image.Pt(currentPoint.X+1, currentPoint.Y))
	neighbors = append(neighbors, image.Pt(currentPoint.X+1, currentPoint.Y+1))
	neighbors = append(neighbors, image.Pt(currentPoint.X, currentPoint.Y+1))
	neighbors = append(neighbors, image.Pt(currentPoint.X-1, currentPoint.Y+1))
	neighbors = append(neighbors, image.Pt(currentPoint.X-1, currentPoint.Y))
	neighbors = append(neighbors, image.Pt(currentPoint.X-1, currentPoint.Y-1))
	neighbors = append(neighbors, image.Pt(currentPoint.X, currentPoint.Y-1))
	neighbors = append(neighbors, image.Pt(currentPoint.X+1, currentPoint.Y-1))
	return neighbors
}
