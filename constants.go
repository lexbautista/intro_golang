// const <const name>
// datatype = <value>
// typed and Untyped

// sample untyped constants

package main
import "fmt"

const PI float64 = 3.14
func main(){
	const name = " Harry Potter"
	const is_muggle = false
	const age = 12

	fmt.Printf("%v: %T \n", name, name)
	fmt.Printf("%v: %T \n", is_muggle, is_muggle)
	fmt.Printf("%v: %T \n", age, age)

	var radius float64 = 5.0
	var area float64 
	area = PI * radius * radius  
	fmt.Println("Area of Circle is :", area)
}