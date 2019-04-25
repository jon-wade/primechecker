package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CalculatePrimacy(v int) bool {
	isPrime := true
	for i := 2; i < v; i++ {
		if v%i == 0 {
			fmt.Printf("%v is divisible by %v \n", v, i)
			isPrime = false
			break
		}
	}
	return isPrime
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for loop := 0; loop < 1; loop++ {
		fmt.Print("Enter an integer: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error: %s \n", err)
		}
		text = strings.TrimSuffix(text, "\n")
		result, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Please enter an integer value!")
			loop--
		} else {
			fmt.Printf("Result: %v \n", CalculatePrimacy(result))
		}
	}
}
