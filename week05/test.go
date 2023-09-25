package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("InPut score 입력 : ")
	reader := bufio.NewReader(os.Stdin)
	inputScore, err := reader.ReadString('\n') //옵션 2
	if err != nil {                        
		log.Fatal(err)
	}
	inputScore = strings.TrimSpace(inputScore)
	score, err := strconv.ParseInt(inputScore, 32)



	if inputScore >= 90 {
		grade := "A grade"
	}
	else {
		grade := "under A grade"
	}
	fmt.Println(inpuinputScore)
	fmt.Println("you will get", grade)
}
