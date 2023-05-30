package main

import "fmt"

const LoginToken string = "asdasdasdasdasdasdasdasd" //Login token is starterd with capital letter so it is accessible outside the package, public variable globally

func main() {

	// string
	var username string = "Roy"
	fmt.Println("Hello", username)
	fmt.Printf("Variable is of a type: %T \n", username)

	// boolean
	var isLoggedIn bool = true
	fmt.Println("Logged in status:", isLoggedIn)
	fmt.Printf("Variable is of a type: %T \n", isLoggedIn)

	// -----------------------------------------------------------------------------------------------------------------
	// Numeric types
	var smallVal uint8 = 255 // max value 255 for uint8
	fmt.Println("Small value:", smallVal)
	fmt.Printf("Variable is of a type: %T \n", smallVal)

	// uint8       the set of all unsigned  8-bit integers (0 to 255)
	// uint16      the set of all unsigned 16-bit integers (0 to 65535)
	// uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
	// uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

	// int8        the set of all signed  8-bit integers (-128 to 127)
	// int16       the set of all signed 16-bit integers (-32768 to 32767)
	// int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
	// int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

	// float32     the set of all IEEE-754 32-bit floating-point numbers
	// float64     the set of all IEEE-754 64-bit floating-point numbers

	// complex64   the set of all complex numbers with float32 real and imaginary parts
	// complex128  the set of all complex numbers with float64 real and imaginary parts

	// byte        alias for uint8
	// rune        alias for int32

	// -----------------------------------------------------------------------------------------------------------------

	var smallFloat float32 = 255.4555555
	fmt.Println("Small float:", smallFloat)
	fmt.Printf("Variable is of a type: %T \n", smallFloat)

	// default values and aliases
	var anotherVariable int
	fmt.Println("Another variable:", anotherVariable)
	fmt.Printf("Variable is of a type: %T \n", anotherVariable)

	// implicit type
	var website = "learngolang.org"
	fmt.Println("Website:", website)
	fmt.Printf("Variable is of a type: %T \n", website)

	// no var style
	numberOfUsers := 300000.0
	fmt.Println("Number of users:", numberOfUsers)
	fmt.Printf("Variable is of a type: %T \n", numberOfUsers)

	//LoginToken
	fmt.Println("Login token:", LoginToken)
	fmt.Printf("Variable is of a type: %T \n", LoginToken)
}
