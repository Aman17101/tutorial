package main

import "fmt"


	type Books struct{
		title string
		id int



	}
	func main(){
		var book1 Books
		var book2 Books
		
		book1.title="developer"
		book1.id=1

		book2.title="programmer"
		book2.id=2
   
		printBook(book1)
		printBook(book2)
		
}
func printBook(Book Books){
	fmt.Printf("Book title %s\n",Book.title)
		fmt.Printf("Book id %d\n",Book.id)

		

		

}
