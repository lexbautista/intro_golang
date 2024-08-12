package main
import ( 
"fmt" 
"strconv")

func main() {
	var i int = 90
	var f float64 = float64(i)
	fmt.Printf("%.2f\n", f)
	fmt.Printf("%v\n", i)
	// 90.00
	// 90
	var s string = strconv.Itoa(i)
	fmt.Printf("%q \n", s) 
	// "90"
	var st string = "200"
	i, err := strconv.Atoi(st)
	fmt.Printf("%v, %T \n", i, i)
	fmt.Printf("%v, %T \n", err, err)
	var st2 string = "200abc"
	i, err2 := strconv.Atoi(st2)
	fmt.Printf("%v, %T \n", i, i)
	fmt.Printf("%v, %T \n", err2, err2)
}

//Itoa - comverts integer to string - returns 1 value
//Atoi - converts string to integer - returns 2 values

