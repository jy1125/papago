package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("InPut score 입력 : ")
	reader := bufio.NewReader(os.Stdin)
	inputScore, _ := reader.ReadString('\n') //옵션 1
	fmt.Println((inputScore))
}
