package main

import(
	"fmt"
	"os"
	"bufio"
)

// Using console comand
func main() {
	in := os.Stdin
	scanner := bufio.NewScanner(in)
	fmt.Print("Print our name: ")
	scanner.Scan()
	name := scanner.Text()
	fmt.Printf("Hello %s!", name)
}

// Using arguments
// func main() {
// 	name := os.Args[1]
// 	fmt.Printf("Hello %s!", name)
// }