package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Print("InPut score 입력 : ")
	reader := bufio.NewReader(os.Stdin)
	inputScore, err := reader.ReadString('\n') //옵션 2
	log.Fatal(err)
	fmt.Println((inputScore))
}
