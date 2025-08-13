package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		cleanedUserInput := cleanInput(scanner.Text())
		if len(cleanedUserInput) == 0 {
			break
		}
		fmt.Println("Your command was:", cleanedUserInput[0])
	}

}
