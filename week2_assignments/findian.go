package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan(){
		str := scanner.Text()
		str = strings.ToLower(str)
		if str[0] == 'i' && str[len(str)-1] == 'n' && strings.Index(str, "a") != -1{
			fmt.Println("Found!")
		} else {
			fmt.Println("Not Found!")
		}

	}
}
