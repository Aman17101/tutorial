package main

import "fmt"

func main() {
	const number2=45638

	var number int
	number = 19
	var addressofnumber *int
	addressofnumber=&number

	number3:=6758
	fmt.Println(number3)

	var decNum float32
	decNum = 39.904040
	var addressoffloat *float32
	addressoffloat=&decNum

	var flags bool
	flags = true
	var adressofflag *bool
	adressofflag=&flags

	var name string
	name = "Aman"
	var adressofname*string
	adressofname=&name

	fmt.Printf("Number value=%d,decimal value=%f,flag value=%v,name value=%s,\n", number, decNum, flags, name)
	fmt.Printf("adressofNumber=%v,adressoffloat=%v,adressofflag =%v,adressofname =%v,\n",addressofnumber, addressoffloat, adressofflag, adressofname)

}
