package main

import "fmt"

func loop () {
	fmt.Println("home work - 2")
	var i int
	for i = 100 ; i > 0 ; i-- {
		if i % 4 == 0{
			fmt.Println(i)
		}
	}
}