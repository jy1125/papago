package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	answer := rand.Intn(100) + 1 //1~100

	fmt.Println("Guess Game ( 1 ~ 100 ) :")
	fmt.Println(answer)

	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < 10; i++ {
		fmt.Println("You chance :", 10-i)
		fmt.Print("Guess number : ")
		inputNumberString, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		inputNumberString = strings.TrimSpace(inputNumberString)

		inputNumber, err := strconv.Atoi(inputNumberString)

		if inputNumber == answer {
			fmt.Println(("정답~"))
			break
		} else if inputNumber < answer {
			fmt.Println("Ur guess number is lower ghan answer.") //정답이 더 크다
		} else if inputNumber > answer {
			fmt.Println("Ur guess number is higher ghan answer.") //정답이 더 작다
		}
	}
}
