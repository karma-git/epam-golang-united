package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReverseString(input string) (output string) {
	rs := []rune(input)

	reversed_bs := make([]rune, len(rs))
	for i := 0; i < len(reversed_bs); i++ {
		reversed_bs[i] = rs[len(reversed_bs)-i-1]
	}
	output = string(reversed_bs)
	return
}

func main() {
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	s = ReverseString(s)
	fmt.Println(s)
}
