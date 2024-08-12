package main
import "fmt"

func main() {
	var name string
	fmt.Print("Enter your name:")
	fmt.Scanf("%s", &name)
	fmt.Println("Hey there,", name)

	var a string
	var b int
	fmt.Print("Enter a string and a number: ")
	count, err :=fmt.Scanf("%s %d", &a, &b)
	fmt.Println("count :", count)
	fmt.Println("error: ", err)
	fmt.Println("a: ", a)
	fmt.Println("b: ",b)

}
