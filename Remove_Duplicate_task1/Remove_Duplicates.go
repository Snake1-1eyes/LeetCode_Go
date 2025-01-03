// Учитывая целочисленный массив чисел, отсортированный в неубывающем порядке,
// удалите дубликаты на месте так, чтобы каждый уникальный элемент появлялся только один раз.
// Относительный порядок элементов должен оставаться неизменным. Затем верните количество уникальных элементов в числах.
// Считайте, что количество уникальных элементов чисел равно k. Чтобы вас приняли, вам нужно сделать следующее:
// Измените массив nums так, чтобы первые k элементов nums содержали уникальные элементы в том порядке,
//
//	в котором они присутствовали в nums изначально. Остальные элементы nums не важны, как и размер nums. Вернуть К.
package Rem_Dup

func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	k := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[k] = nums[i]
			k++
		}
	}

	return k
}
