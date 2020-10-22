package main

import "fmt"

func mayor(args ...int) int {
	may := 0
	for _, v := range args {
		if may < v {
			may = v
		}
	}
	return may
}

func main() {
	a := []int{7, 5, 4, 1, 2, 4, 78, 6}

	fmt.Println(mayor(a...))

}
