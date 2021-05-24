// define package main
package main
// import package fmt
import "fmt"

// this is a comment
// define function main
func main() {
	/*x := [5]float64 { 98, 93, 77, 82, 83 }
	var total float64 = 0

	for _, value := range x {
		total += value
	}

	fmt.Println(total / float64(len(x)))*/
	/*arr := [5]float64 {1, 2, 3, 4, 5}
	x := arr[:]
	fmt.Println()*/

	slice1 := []int {1,2,3}
	//slice2 := append(slice1, 4, 5)
	slice2 := make([]int, 2)
	fmt.Println(slice1, slice2)

	copy(slice2, slice1)
	fmt.Println(slice1, slice2)

	x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x, x["key"])

	elements := map[string]map[string]string {
		"H": {
			"name": "Hydrogen",
			"state": "gas",
		},
		"He": {
			"name": "Helium",
			"state": "gas",
		},
		"Li": {
			"name": "Lithium",
			"state": "solid",
		},
		"Be": {
			"name": "Beryllium",
			"state": "solid",
		},
		"B": {
			"name": "Boron",
			"state": "solid",
		},
		"C": {
			"name": "Carbon",
			"state": "solid",
		},
		"N": {
			"name": "Nitrogen",
			"state": "gas",
		},
		"O": {
			"name": "Oxygen",
			"state": "gas",
		},
		"F": {
			"name": "Flourine",
			"state": "gas",
		},
		"Ne": {
			"name": "Neon",
			"state": "gas",
		},
	}

	/*fmt.Println(elements["Li"])

	if name, isOk := elements["Uh"]; isOk {
		fmt.Println(name, isOk)
	} else {
		fmt.Println(isOk)
	}*/

	if el, ok := elements["Li"]; ok {
		fmt.Println(el["name"], el["state"], ok)
	}

	array := []int {78, 45, 13, 13, 1, 4, 6, 9, 0, 213, 5}
	min := array[0]

	for i := 1; i < len(array); i++ {
		if min > array[i] {
			min = array[i]
		}
	}
	fmt.Println(array)

	fmt.Println("min of array =", min)
}
