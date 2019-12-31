package src

import "fmt"

func checkThis() {
	res, _ := GetAllProductWith("shake")
	fmt.Println(res)
}
