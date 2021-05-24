// define package main
package main
// import package fmt
import "fmt"

// this is a comment
// define function main
func main() {
	// calling function Print from fmt package
	fmt.Println("1 + 1 =", 1 + 1)
	fmt.Println("1 + 1 =", 1.0 + 1.0)
	fmt.Println("1 - 1 =", 1 - 1)
	fmt.Println("1 * 1 =", 1 * 1)
	fmt.Println("1 % 1 =", 1 % 1)
	fmt.Println("1 % 1 =", 1 % 1)
	fmt.Println("32132 × 42452 =", 32132 * 42452)
	fmt.Println("Length of 'Length' =", len("Length"))
	fmt.Println()
	fmt.Println("true && true -", true && true)
    fmt.Println("true && false -", true && false)
    fmt.Println("true || true -", true || true)
    fmt.Println("true || false -", true || false)
    fmt.Println("!true -", !true)
}