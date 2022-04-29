package main

import "fmt"

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	res := GetSet(&arr)
	fmt.Println(*res)
}

// GetSet O(n*2)
func GetSet(strs *[]string) *[]string {
	maps := make(map[string]struct{})
	sets := new([]string)
	for _, v := range *strs {
		maps[v] = struct{}{}
	}
	for k, _ := range maps {
		*sets = append(*sets, k)
	}
	return sets
}
