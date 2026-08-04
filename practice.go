package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	maxretries = 4
)

func styles(text, ansi string) string {
	return ansi + text + "\033[0m"
}

func containsAtSymbol(email string) bool {
	return strings.Contains(email, "@")
}

func capitalize(s string) string {
	if s == "" {
		return s
	}

	return strings.ToUpper(string(s[0])) + strings.ToLower(string(s[1:]))
}

func main() {

	//FIRSNAME VERIFICATION

	fmt.Println(" ")
	fmt.Println(style("********** IDENTITY VERIFICATION ***********", "\033[1;33m"))
	fmt.Println(" ")
	fmt.Println(style("PLEASE ENTER YOUR DETAILS TO PROCEED ", "\033[1;33m"))
	fmt.Println(" ")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("INPUT YOUR FIRSTNAME: ")
start:
	Firstname, _ := reader.ReadString('\n')
	Firstname = strings.TrimSpace(capitalize(Firstname))
	if Firstname == "" {
		fmt.Println(style("ERROR❗❗ FIRSTNAME CANNOT BE EMPTY", "\033[1;31m"))
		fmt.Println(style("TRY AGAIN", "\033[1;31m"))
		fmt.Println("")
		fmt.Println("RE-ENTER YOUR FIRSTNAME")
		goto start
	}

	if strings.Contains(Firstname, " ") {
		fmt.Println(style("ERROR❗❗ FIRSTNAME CANNOT CONTAIN SPACES IN-BETWEEN", "\033[1;31m"))
		fmt.Println(style("TRY AGAIN", "\033[1;31m"))
		fmt.Println("")
		fmt.Println("RE-ENTER YOUR FIRSTNAME")
		goto start
	}
	for _, char := range Firstname {
		if (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') {
		} else {
			fmt.Println(" ")
			fmt.Println(style("ERROR❗❗ ONLY ALPHABETS ARE ALLOWED", "\033[1;31m"))
			fmt.Println(style("TRY AGAIN", "\033[1;31m"))
			fmt.Println(" ")
			fmt.Print("PLEASE RE-ENTER YOUR FIRSTNAME: ")
			goto start
		}
	}

	// LASTNAME VERIFICATION

	fmt.Println(" ")

	reader = bufio.NewReader(os.Stdin)
	fmt.Print("INPUT YOUR LASTNAME: ")
start2:
	Lastname, _ := reader.ReadString('\n')
	Lastname = strings.TrimSpace(capitalize(Lastname))
	if Lastname == "" {
		fmt.Println(style("ERROR❗❗ LASTNAME CANNOT BE EMPTY", "\033[1;31m"))
		fmt.Println(style("TRY AGAIN", "\033[1;31m"))
		fmt.Println("")
		fmt.Print("PLEASE RE-ENTER YOUR LASTNAME: ")
		goto start2
	}

	if strings.Contains(Lastname, " ") {
		fmt.Println(style("ERROR❗❗ LASTNAME CANNOT CONTAIN SPACES IN-BETWEEN", "\033[1;31m"))
		fmt.Println(style("TRY AGAIN", "\033[1;31m"))
		fmt.Println("")
		fmt.Print("PLEASE RE-ENTER YOUR LASTNAME: ")
		goto start2

	}

	for _, char := range Lastname {
		if (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') {
		} else {
			fmt.Println(" ")
			fmt.Println(style("ERROR❗❗ ONLY ALPHABETS ARE ALLOWED", "\033[1;31m"))
			fmt.Println(style("TRY AGAIN", "\033[1;31m"))
			fmt.Println(" ")
			fmt.Print("PLEASE RE-ENTER YOUR LASTNAME: ")
			goto start2
		}
	}

	//PRINTING THE VERIFIED FIRSTNAME AND LASTNAME
	fmt.Println(" ")
	fmt.Printf(style("HELLO %s %s! YOU ARE VERIFIED\n", "\033[1;32m"), Firstname, Lastname)
	fmt.Println(" ")
	fmt.Println(style("=========================================", "\033[1;32m"))
	fmt.Println(" ")
	fmt.Println(style("TO ENTER VIP SECTION, YOU NEED TO INPUT AN ADMIN PASSWORD", "\033[1;33m"))

	// PASSWORD VERIFICATION

	reader = bufio.NewReader(os.Stdin)
	attempt := 1

	for attempt <= maxretries {
		fmt.Println(style("PLEASE ENTER YOUR ADMIN PASSWORD: ", "\033[1;33m"))
		fmt.Println("YOU HAVE", maxretries-attempt+1, "ATTEMPT(S) LEFT")

		password, _ := reader.ReadString('\n')
		password = strings.TrimSpace(password)
		if password == "" || strings.Contains(password, " ") {
			fmt.Println(style("ERROR❗❗ PASSWORD CANNOT BE EMPTY OR CONTAIN SPACES", "\033[1;31m"))
			fmt.Println(style("TRY AGAIN", "\033[1;31m"))
			fmt.Println(" ")
			attempt++
			continue
		}
		if len(password) < 8 {
			fmt.Println(style("ERROR❗❗ PASSWORD MUST BE AT LEAST 8 CHARACTERS LONG", "\033[1;31m"))
			fmt.Println(style("TRY AGAIN", "\033[1;31m"))
			fmt.Println(" ")
			attempt++
			continue
		}

		var hasUpper bool
		var hasLower bool
		var hasNumber bool
		var hasSymbol bool

		for _, char := range password {
			if char >= 'A' && char <= 'Z' {
				hasUpper = true
			}
			if char >= 'a' && char <= 'z' {
				hasLower = true
			}
			if char >= '0' && char <= '9' {
				hasNumber = true
			}
			if char == '!' || char == '@' || char == '#' {
				hasSymbol = true
			}
		}

		if !hasUpper || !hasLower || !hasNumber || !hasSymbol {
			fmt.Println(style("PASSWORD IS TOO WEAK", "\033[1;31m"))
			fmt.Println(style("TRY AGAIN", "\033[1;31m"))
			fmt.Println(" ")
			attempt++
			continue
		}

		reader = bufio.NewReader(os.Stdin)
		fmt.Print(style("CONFIRM PASSWORD: ", "\033[1;33m"))
		confirm, _ := reader.ReadString('\n')
		confirm = strings.TrimSpace(confirm)
		if password != confirm {
			fmt.Println(style("PASSWORDS DO NOT MATCH", "\033[1;31m"))
			fmt.Println(style("TRY AGAIN", "\033[1;31m"))
			fmt.Println(" ")
			attempt++
			continue
		}

		fmt.Println(style("PASSWORD VERIFIED", "\033[1;32m"))
		return
	}

	fmt.Println(style("YOU HAVE EXCEEDED THE MAXIMUM NUMBER OF ATTEMPTS. ACCESS DENIED.", "\033[1;31m"))
	fmt.Println(style("PLEASE CONTACT THE ADMINISTRATOR FOR ASSISTANCE", "\033[1;31m"))
	fmt.Println(" ")
}
