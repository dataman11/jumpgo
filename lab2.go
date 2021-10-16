package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Println("The no is ", i)
	switch i {
	case 1:
		fmt.Println("The no is 1")
	case 2:
		fmt.Println("The no is 2")
	case 3:
		fmt.Println("The no is 3")

	}
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("weekend is here")
	case time.Wednesday:
		fmt.Println("middle day ")
	default:
		fmt.Println("all other day")

	}

	c := func(j interface{}) {
		switch m := j.(type) {
		case int:
			fmt.Println("The no is integer")
		case bool:
			fmt.Println("the no is boolean")
		default:
			fmt.Printf("the no is string %T\n", m)
		}

	}
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
	c("12")

	var a [5]int
	a[2] = 2
	fmt.Println(a)
	fmt.Println(len(a))

	b := [3]int{1, 2, 4}
	fmt.Printf("the be is %T ", b)
	fmt.Println(len(b))

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	p := make([]int, 2)
	fmt.Println(p)
	p[0] = 1
	p[1] = 2
	fmt.Println(p)

	s := make([]string, 4)

	s[0] = "alan"
	s[1] = "one"
	s[3] = "one"
	s[2] = "wwww"
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println((s[:3]))
	// fmt.Println()
	// s = append(s, "1", "4", "1", "2", "3", "4", "w")
	// s[0] = "alan"
	// fmt.Println(s)

	// sl1 := make([][]int, 3)
	// fmt.Println(sl1)
	// for i := 0; i < 3; i++ {
	// 	inner := i + 1
	// 	sl1[i] = make([]int, inner)
	// 	for j := 0; j < inner; j++ {
	// 		sl1[i][j] = i + j
	// 	}
	// }
	// fmt.Println(sl1)

	arr := make([][]int, 3)
	fmt.Println(arr)
	for i := 0; i < 3; i++ {
		arr[i] = make([]int, i+3)
		n := i + 1
		for j := 0; j < n; j++ {
			arr[i][j] = j + 1

		}

	}
	fmt.Println(arr)

}
