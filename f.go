package main

import "fmt"

func main() {
	fmt.Println("hello")
	hello()
	fmt.Println("then we print even numbers between 1 and 50")
	for i:=1;i<=50;i++{
		if i%2==0 {
			fmt.Println(i)

		}

	}

}
func hello()  {
	fmt.Println("this is from hello")

}
