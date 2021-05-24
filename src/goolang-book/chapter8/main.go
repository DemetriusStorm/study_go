// define package main
package main
// import package fmt
import (
	"fmt"
	//"os"
)

// this is a comment
//define average function
func average(xs []float64) float64 {
	//panic("Not Implemented")
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

// define add function
func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

// wrapper
func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

// recursion
func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x - 1)
}

func first() {
	fmt.Println("1st")
}

func second() {
	fmt.Println("2nd")
}

func main() {
	// отложенный вызов функции
	// выполнится перед завершением функции, в самом конце
	defer second() // выполнится после panic
	first() // выполнится первой
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	
	panic("PANIC!")
	// f, _ := os.Open(filename)
	// defer f.Close()
}

// define function main
/*func main() {
	xs := []float64 { 98, 93, 77, 82, 83 }
	fmt.Println(average(xs))

	fmt.Println(add(1, 2, 3, 4, 5))

	v := []int { 98, 93, 77, 82, 83 }
	fmt.Println(add(v...))
}
*/
/*func main() {
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(1, 2))
}*/
/*func main() {
	x := 0
	increment := func() int {
		 x++
		 return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
}*/
