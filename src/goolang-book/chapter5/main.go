// define package main
package main
// import package fmt
import "fmt"

// this is a comment
// define function main
func main() {
	for i := 1; i <= 10; i++ {
		if i % 2 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}
	}
	fmt.Println("Tadam!\n")

	for i:= 1; i <=7; i++ {
		switch i {
		case 0: fmt.Println(i, "Zero")
		case 1: fmt.Println(i, "One")
		case 2: fmt.Println(i, "Two")
		case 3: fmt.Println(i, "Three")
		case 4: fmt.Println(i, "Four")
		case 5: fmt.Println(i, "Five")
		default: fmt.Println(i, "Unknown number")
		}
	}
}
