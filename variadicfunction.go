package main

import "fmt"

func variadic(param  ...int){
	fmt.Printf("%T\n%v\n",param,param)
}
func main(){
	variadic(1,2,3,4,5,6,7,8,9)
}
