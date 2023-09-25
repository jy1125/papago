package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("InPut score 입력 : ")
	reader := bufio.NewReader(os.Stdin)
	inputScore := reader.ReadString('\n')
	fmt.Println((inputScore))
}
