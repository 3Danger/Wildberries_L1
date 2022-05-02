package main

import "fmt"

/*
	Имеется последовательность строк - (cat, cat, dog, cat, tree)
	создать для нее собственное множество.
*/

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	res := GetSet(&arr)
	fmt.Println(*res)
}

// GetSet O(n*2)
func GetSet(strs *[]string) *[]string {
	// Бинарное дерево будет быстрее
	// Для нашей задачи подходит лучше всего
	// тем самым избежим лишние проходы по слайсу
	maps := make(map[string]struct{})
	for _, v := range *strs {
		maps[v] = struct{}{}
	}

	result := new([]string)
	for k, _ := range maps {
		*result = append(*result, k)
	}
	return result
}
