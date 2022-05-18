package main
import "fmt"

type Filter func(string) string // type declaration

func sayHello (name string, filter Filter){
	filterName := filterSpam(name)
	fmt.Println("Hello", filterName)

}

func filterSpam (name string) string{
	if(name == "Anjing"){
		return "..."
	}else{
		return name
	}
}

func main(){

	sayHello("Tama", filterSpam);
	sayHello("Anjing",filterSpam);
}