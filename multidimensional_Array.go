package main
import "fmt"

func main(){
	var a=[5][2]int {{0,10},{1,2},{2,5},{3,48},{4,76}}
	var i,j int
	for i=0;i<5;i++{
		for j=0;j<2;j++{
			fmt.Printf("a[%d][%d]=%d\n",i,j,a[i][j])
		}
	}
	}
