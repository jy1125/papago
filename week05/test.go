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
	if err != nil {                            // 조건문 들어갔늨데 이해 안감 뭐라는거여
		log.Fatal(err)
	}
	fmt.Println((inputScore))
}
