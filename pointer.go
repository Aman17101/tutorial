package main

import "fmt"

func main(){
	var a int =20
	var ip*int
	 ip=&a
	 fmt.Printf("adress of  variable %x\n",&a)
	 fmt.Printf("adress stored in  ip variable %x\n",ip)
	 fmt.Printf("value of ip variable %d\n",*ip)

}
