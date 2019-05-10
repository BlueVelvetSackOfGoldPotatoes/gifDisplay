package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func main() {
	fmt.Println("Press q to quit")
	reader := bufio.NewReader(os.Stdin)

	for {
		gif := rand.Intn(38)
		fmt.Println(gif)
		fmt.Println("Press a key")

		key, _ := reader.ReadString('\n')

		if strings.Compare("q", key) == 1 {

		}
	}
}
