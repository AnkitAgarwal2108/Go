package main
import ("fmt")

//value of b can't be changed once assigned
func main() {
	const b =10
	fmt.Println(b)
	b = 30
	fmt.Println(b)
}