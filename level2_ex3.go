package main
//Create TYPED and UNTYPED constants. Print the values of the constants.
import "fmt"

const(
	a = 5
	b int = 5
	c string = "123dasd"
)
func main() {
	fmt.Println(a, b, c)
}
