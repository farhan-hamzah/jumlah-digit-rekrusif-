package main
import "fmt"

func jumlahDigit(n int)int{
	if n == 0{
		return 0
	}else{
		return n%10 + jumlahDigit(n/10)
	}
}
func main(){
	var num int
	fmt.Scan(&num)
	fmt.Print(jumlahDigit(num))
}


