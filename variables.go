package main
import "fmt"
//global variable
var name string = "Hailey"
func main() {
	var city string = "Manila"
	var user string = "Lexander"
	var grades int = 42
	fmt.Println(name)
	fmt.Print("Welcome to ", city, ", \n", user)
	fmt.Println(city)
	fmt.Println(user)
	fmt.Printf("Nice to see you here, at %v \n",city)
	fmt.Printf("Marks: %d\n", grades)
    fmt.Printf("Hey %v!  You have a score of %d/100 in Mathh", user, grades)
    var s,t string = "foo", "bar"
	fmt.Printf("\n %v and %v \n", s, t)
	var (
		l string = "lex"
		i int = 5
	)
	fmt.Println(l)
	fmt.Println(i)
    //short variable Declaration
	name:= "Lisa"
	name ="Peter"
	//result will be peter
	fmt.Println(name)

	/*Zero Values
	bool = false
	int =0
	float64 = 0.0
	string ="'
    */
}