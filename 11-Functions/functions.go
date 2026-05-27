package main

import "fmt"

//Functions

func plus(a int, b int) int{
	return  a+ b
}

func plusPlus(a, b, c int) int{
	return  a + b + c 
}
func main() {
	res := plus(1,2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1,2,3)
	fmt.Println("1+2+3 =", res)

	//Multiple Return values
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)
	_, c:= vals()
	fmt.Println(c)

	//Variadic Functions
	sum(1,2)
	sum(1,2,3)
	nums := []int{1,2,3,4}
	sum(nums...)
}

//Multiple Return Values

func vals() (int, int) {
	return 3, 7
}


//variadic Function

func sum(nums ...int){
	fmt.Println(nums, " ")
	total := 0

	for _, num := range nums {
		total +=num
	}
	fmt.Println(total)
}




