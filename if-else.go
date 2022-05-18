package main
import "fmt" 

func main(){

	var name = "tamasssss"

	if name=="tama" {
		fmt.Println("ini Tama")
	} else if name== "tamas"{
		fmt.Println("ini Tamas")
	} else {
		fmt.Println("ini siapa?")
	}

	if length := len(name); length>5 {
		fmt.Println("nama terlalu panjang")
	}
	
}