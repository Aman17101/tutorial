package main

import "fmt"

func main(){
	var myMap map[string]int
	myMap=map[string]int{}
	myMap["a"]=12
	myMap["b"]=13
	myMap["c"]=34
	myMap["d"]=54
	myMap["e"]=36

	delete(myMap,"e")

	fmt.Println(myMap)
}
