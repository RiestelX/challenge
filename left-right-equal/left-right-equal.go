package main

import (
	"bufio"
	"fmt"
	"os"
)

func decoderV2(text string) string {
	const left = 'L'
	const right = 'R'
	const equal = '='

	temp := make([]int, len(text)+1)
	countL, countR := 0, 0

	for i := 0; i < len(text); i++ {
		if text[i] == left {
			countL++
			countR = 0
			for j := i; j >= i-(countL-1); j-- {
				if temp[j] <= temp[j+1] {
					temp[j] = temp[j+1] + 1
				}
			}
		} else if text[i] == right {
			countL = 0
			countR++
			temp[i+1] = temp[i] + 1
		} else if text[i] == equal {
			countL = 0
			countR = 0
			temp[i+1] = temp[i]
		}
	}

	for i := len(temp) - 1; i > 0; i-- {
		if text[i-1] == equal {
			temp[i-1] = temp[i]
		}
	}

	result := ""
	for i := 0; i < len(temp)-1; i++ {
		result += fmt.Sprintf("%d", temp[i])
	}

	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("encoded : ")
	text, _ := reader.ReadString('\n')
	text = text[:len(text)-1]
	fmt.Printf("o : %s\n", decoderV2(text))
}
