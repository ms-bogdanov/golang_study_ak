package main

import "fmt"

func main() {
	map1 := map[string]int{"apple": 3, "banana": 2}
	map2 := map[string]int{"orange": 5, "grape": 4}

	mergeMap := mergeMaps(map1, map2)
	for key, value := range mergeMap {
		fmt.Printf("%s: %d\n", key, value)
	}
}

func mergeMaps(map1, map2 map[string]int) map[string]int {
	res := make(map[string]int)

	for k, v := range map1 {
		res[k] = v
	}
	for k, v := range map2 {
		res[k] = v
	}
	return res
}
