package main

import "fmt"

func main() {
	var start int
	start = 2
	var end int = 11
	for ; start <= end; start++ {
    var k = 0
    var j = 1
		for ; j <= start; j++ {
			if (start % j) == 0 {
				k = k + 1
			}
		}
		if (k == 2) {
			fmt.Println(start)
		}
	}
}