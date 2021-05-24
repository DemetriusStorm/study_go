// define package main
package main
// import package fmt
import "fmt"

// this is a comment
// define function main
func main() {
	// define variables string type
	var x string = "Hello, World!"
	var y string
	// x := "Hello, World!"
	// y := "Hello, World!"
	// z := 5
	// var f = 5
	const constant = "Hello!"
	fmt.Println(constant)
	// Error!!!
	// constant = "1"
	// fmt.Println(constant)

	/*var (
		a = 1
		b = 2
		c = 3
	)*/

	y = "Well, I'am Okay. And you?"
	// calling function Print from fmt package
    fmt.Println(x)   
    fmt.Println(y)

    x = "First word. "
    y = "Second word.. "
    fmt.Println(x)   
    fmt.Println(y)    

    x += "First "
    y = x + y
    fmt.Println(x)
    fmt.Println(y)

    x = "hello"
	y = "world"
	fmt.Println(x == y)

	x = "hello"
	y = "hello"
	fmt.Println(x == y)

	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2
	fmt.Println(output)
}
