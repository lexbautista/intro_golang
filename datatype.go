package main
import "fmt"
import "reflect"

func main () {
	var grades int = 42
	var message string = "hello world"
	var isCheck bool = true
	var amount float32 = 5466.54

	fmt.Printf("variable grades = %v is of type %T \n", grades, grades)
	fmt.Printf("variable message = '%v' is of type %T \n", message, message)
	fmt.Printf("variable isCheck = '%v' if of type %T \n", isCheck, isCheck)
	fmt.Printf("variable amount = %v is of type %T \n", amount, amount)

	fmt.Printf("Type %v \n", reflect.TypeOf(1000))
	fmt.Printf("variables message='%v' is of type %v \n", message, reflect.TypeOf(message))

}