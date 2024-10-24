package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("err!:%v", err)
	}
	input = strings.TrimSpace(input)
	inputInt, err2 := strconv.Atoi(input)
	if err2 != nil {
		fmt.Printf("convert err!:%v", err2)
	} else {
		fmt.Println(inputInt)
	}
}
