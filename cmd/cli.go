package main

import "os"
import "fmt"
import "bufio"
import "github.com/fuele/totp-client/pkg/totp"

// required base32 secret
// algorithm
// interval

func main() {
	var secret string
	if isInputFromPipe(){
		secret = getInputFromPipe()
	}else{
		secret = os.Args[1]
	}

	run(secret)
}

func isInputFromPipe() bool {
    fileInfo, _ := os.Stdin.Stat()
    return fileInfo.Mode() & os.ModeCharDevice == 0
}

func getInputFromPipe() string{
	var input string
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	for scanner.Scan() {
		input = scanner.Text()
	}

	return input
}

func run(secret string){
	t := totp.Totp {
		Digits: 6,
		Algorithm: "sha1", 
		TimeZero: 0,
		TimeStep: 30,
	}
	code, err := t.ComputeTotp(secret)

	if err != nil {
		os.Stderr.WriteString("Failed to generate code. Error: "+err.Error())
		os.Exit(-1)
	}
	
	for _, digit := range code {
		fmt.Print(digit)
	}
	fmt.Println()
}
