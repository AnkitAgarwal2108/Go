package main
import "fmt"

func main() {
    //declaring a integer variable x
    var x int
    x=3 //assigning x the value 3
    fmt.Println("x:", x) //prints 3

    //declaring a integer variable y with value 20 in a single statement and prints it
    var y int=20
    fmt.Println("y:", y)

    //declaring a variable z with value 50 and prints it
    //Here type int is not explicitly mentioned
    var z=50
    fmt.Println("z:", z)

    //Multiple variables are assigned in single line- i with an integer and j with a string
    var i, j = 100,"hello"
    fmt.Println("i and j:", i,j)

    //declaring a float32 variable
    var a float32
    a = 3.50
    fmt.Println("a:", a)

    //declaring a boolean variable
    var b bool
    b = true
    fmt.Println("b:", b)

    //declaring a string variable
    var c string
    c = "John Doe"
    fmt.Println("c:", c)