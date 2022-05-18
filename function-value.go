package main
import "fmt" 

func getGoodBye(name string)string{
	return "Good bye "+ name
}

func main(){
	result:= getGoodBye
	fmt.Println(result("tama"))

}