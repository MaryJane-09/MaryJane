package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func style(text, ansi string) string {
    return ansi + text + "\033[0m"
}

func squareRoot(a float64) float64 {

    if a == 0 {

        return 0
    }
    if a < 0 {
        fmt.Println(style("DOES NOT HANDLE NEGATIVE INPUTS", Red))
        return 0
    }
    n := a
    for i := 0; i < 10; i++ {
        n = (n + a / n) / 2
    }
    return n
}

// func power(a float64) float64 {

// }

func helpUser(){
        text := []string{
            "",
            "-------- AVAILABLE COMMAND -----------",
            "<a> + <b>   → addition",
            "<a> - <b>   → subtraction",
            "<a> * <b>   → multiplication",
            "<a> / <b>   → division",
            "<a> %  <b>  → modulus (FOR INTEGERS ONLY!)",
            "<a> sqrt    → square root",
            "",
    }
    fmt.Println(style(strings.Join(text, "\n"), Blue))
}

const (
    Red    = "\033[1;31m"
    Green  = "\033[1;32m"
    Yellow = "\033[1;33m"
    Blue   = "\033[1;34m"
)

func main() {
    fmt.Println()
    fmt.Println(style("********** WELCOME TO BLAISE PASCAL'S CALCULATOR🧮 ***********", Yellow))
    fmt.Println(style("TYPE; \n<help> FOR HELP\n<off> TO EXIT", Yellow))
    fmt.Println()

    reader := bufio.NewReader(os.Stdin)

start :
    fmt.Print("FIRST VALUE: ")
    first_value, _ := reader.ReadString('\n')
    first_value = strings.TrimSpace(strings.ToLower(first_value))

    if first_value == "off" {
        fmt.Println(style("Thank you for using Blaise Pascal's Calculator!", Green))       
        return
    }
    if first_value == "help"{
        helpUser()
        goto start
    }

    num1, err1 := strconv.ParseFloat(first_value, 64)

    if err1 != nil {
        fmt.Println(style("INVALID NUMBER", Red))
        fmt.Println(style("TRY AGAIN", Red))
        fmt.Println()
        goto start
    }


    fmt.Print("OPERATOR: ")
    sign, _ := reader.ReadString('\n')
    sign = strings.TrimSpace(strings.ToLower(sign))

    if sign == "off" {
        fmt.Println(style("Thank you for using Blaise Pascal's Calculator!", Green))
        return
    }
    if sign == "help"{
        helpUser()
        goto start
    }

    if sign == "sqrt" {
        fmt.Println()
        fmt.Printf(style("√%.2f = %.2f\n", Green), num1, squareRoot(num1))
        fmt.Println()
        goto start
    }

start2 :
    fmt.Print("SECOND VALUE: ")
    second_value, _ := reader.ReadString('\n')
    second_value = strings.TrimSpace(strings.ToLower(second_value)) 

    if second_value == "off" {
        fmt.Println(style("Thank you for using Blaise Pascal's Calculator!", Green))
        return
    }
    if second_value == "help" {
        helpUser()
        goto start2
    }

    num2, err2 := strconv.ParseFloat(second_value, 64)

    if err2 != nil {
        fmt.Println(style("INVALID NUMBER", Red))
        fmt.Println(style("TRY AGAIN", Red))
        fmt.Println()
        goto start2
    }

    switch sign {
    case "+" :
        fmt.Println()
        fmt.Printf(style("%.2f %s %.2f = %.2f\n", Green), num1, sign, num2, num1 + num2)

    case "-" :
        fmt.Println()
        fmt.Printf(style("%.2f %s %.2f = %.2f\n", Green), num1, sign, num2, num1 - num2)

    case "*" :
        fmt.Println()
        fmt.Printf(style("%.2f %s %.2f = %.2f\n", Green), num1, sign, num2, num1 * num2)

    case "/" :
        if num2 == 0 {
			fmt.Println(style("THE DIVISOR CANNOT BE ZERO", Red))
			fmt.Println(style("TRY AGAIN", Red))
            fmt.Println()
			goto start2
        }
        fmt.Println()
        fmt.Printf(style("%.2f %s %.2f = %.2f\n", Green), num1, sign, num2, num1 / num2)

    case "%" :
        if num2 == 0 {
			fmt.Println(style("THE DIVISOR CANNOT BE ZERO", Red))
			fmt.Println(style("TRY AGAIN", Red))
            fmt.Println()
			goto start2
        }
        if num1 == float64(num1) || num2 == float64(num2){
            fmt.Println(style("FOR INTEGERS ONLY!", Red))
            fmt.Println(style("TRY AGAIN", Red))
            fmt.Println()
            goto start
        }
        fmt.Println()
        fmt.Printf(style("%v %s %v = %v\n", Green), num1, sign, num2, int(num1) % int(num2))

    default :
        fmt.Println(style("INVALID ARITHMETIC SYNTAX", Red))
        fmt.Println(style("TRY AGAIN", Red))
        fmt.Println()
        goto start
    }
    fmt.Println()
    goto start
}