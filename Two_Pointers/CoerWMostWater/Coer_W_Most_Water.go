package main

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxW := 0

	for left < right {

		// минимальная высота
		minH := height[left]
		if height[right] < minH {
			minH = height[right]
		}

		// площадь равна минимальной высоте * расстояние между двумя стенами
		width := right - left
		area := width * minH
		if maxW < area {
			maxW = area
		}

		// если левая стена меньше правой, то двигаем левую стену вправо
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxW
}
