package main

import "fmt"

func main(){
	var a int=100

	var b int =200

fmt.Printf("value of a  before swap is :%d\n",a)
fmt.Printf("value of b  before swap is %d\n",b)

swap(a,b)                                                     //address is swap in call by value ,there is no change in value
fmt.Printf("value of a after swap%d\n",a)
fmt.Printf("value of b after swap%d\n",b)
}

func swap(x int,y int){
	var temp int
	temp=x
	x=y
	y=temp

}
