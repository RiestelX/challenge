package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	// var s = [][]int{
	// 	{59},
	// 	{73, 41},
	// 	{52, 40, 53},
	// 	{26, 53, 6, 34},
	// }

	data, err := os.ReadFile("./max-path/hard.json")
	if err != nil {
		fmt.Println("error ", err)
	}

	var jsonData [][]int
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		fmt.Println("error ", err)
	}

	fmt.Println(findMaxPath(jsonData))
}

func findMaxPath(data [][]int) int {
	for row := len(data) - 2; row >= 0; row-- {
		for col := 0; col < len(data[row]); col++ {
			data[row][col] += max(data[row+1][col], data[row+1][col+1])
		}
	}
	return data[0][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
