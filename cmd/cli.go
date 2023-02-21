package main

import "fmt"
import "github.com/fuele/totp-client/pkg"

func main() {
	fmt.Println("Starting program")
	fmt.Println(totp_generator.OutputCode())
	fmt.Println("Ending program")
}
