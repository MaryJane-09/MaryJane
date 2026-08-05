import sys

RED = "\033[1;31m"
GREEN = "\033[1;32m"
BLUE = "\033[1;34m"
CYAN = "\033[1;36m"
WHITE = "\033[1;37m"
RESET = "\033[0m"

valid_operators = {"+", "-", "*", "/", "%", "sqrt"}


def style(text:str, ansi:str) -> str:
    return ansi + text + RESET

def square_root(a: float) -> float:
    if a == 0:
        return 0 

    n = a
    for _ in range(10):
        n = (n + a / n) / 2
    return n

def goodbye():
    print(style("THANK YOU FOR USING BLAISE PASCHAL'S CALCULATOR", GREEN))

def help_user():
    text = """\
-------- AVAILABLE COMMAND -----------
<a> + <b>   → addition
<a> - <b>   → subtraction
<a> * <b>   → multiplication
<a> / <b>   → division
<a> %  <b>  → modulus
<a> sqrt    → square root
"""

    print(style(text, BLUE))

print()
print(style("********** WELCOME TO BLAISE PASCAL'S CALCULATOR🧮 ***********", CYAN))
print(style("TYPE; \n<help> FOR HELP\n<off> TO EXIT", CYAN))
print()


while True:
    while True:
        f_value = input(style("FIRST VALUE: ", WHITE)).lower()
        print()

        if f_value == "off":
            goodbye()
            sys.exit()

        if f_value == "help":
            help_user()
            continue

        try:
            num1 = float(f_value)
        except ValueError:
            print(style("INVALID NUMBER", RED))
            print(style("TRY AGAIN", RED))
            print()
        else:
            break

    while True:
        sign = input(style("OPERATOR: ", WHITE)).lower()
        print()

        if sign == "off":
            goodbye()
            sys.exit()

        if sign == "help":
            help_user()
            continue
        if sign == "sqrt":
            if num1 < 0:
                print(style("DOES NOT HANDLE NEGATIVE INPUT", RED))
                print(style("TRY AGAIN", RED))
                print()
                break

        if sign == "sqrt":
            print(style(f"√{num1:.2f} = {square_root(num1):.2f}", GREEN))
            print()
            break

        if sign not in valid_operators:
            print(style("THIS OPERATOR IS NOT SUPPORTED. TYPE <help> TO SEE AVAILABLE OPERATORS", RED))
            print(style("TRY AGAIN", RED))
            print()
            continue
        break

    if sign == "sqrt":
        continue

    while True:
        s_value = input(style("SECOND VALUE: ", WHITE)).lower()
        print()

        if s_value == "off":
            goodbye()
            sys.exit()

        if s_value == "help":
            help_user()
            continue

        try:
            num2 = float(s_value)

            if sign == "/" or sign == "%":
                if num2 == 0:
                    print(style("THE DIVISOR CANNOT BE ZERO", RED))
                    print(style("TRY AGAIN", RED))
                    print()
                    continue
        except ValueError:
            print(style("INVALID NUMBER", RED))
            print(style("TRY AGAIN", RED))
            print()
        else:
            break

    match sign:

        case "+":
            result = num1+num2
            print()
            print(style(f"{num1:.2f} {sign} {num2:.2f} = {result:.2f}", GREEN))

        case "-":
            result = num1 - num2
            print()
            print(style(f"{num1:.2f} {sign} {num2:.2f} = {result:.2f}", GREEN))

        case "*":
            result = num1 * num2
            print()
            print(style(f"{num1:.2f} {sign} {num2:.2f} = {result:.2f}", GREEN))

        case "/":
            result = num1 / num2
            print()
            print(style(f"{num1:.2f} {sign} {num2:.2f} = {result:.2f}", GREEN))

        case "%":
            result = num1 % num2
            print()
            print(style(f"{num1:.2f} {sign} {num2:.2f} = {result:.2f}", GREEN))

        case _:
            print(style("INVALID ARITHMETIC SYNTAX", RED))
            print(style("TRY AGAIN", RED))
            print()

    print()
    continue