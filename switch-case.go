package main
import "fmt" 

func main(){

	var name = "Tama"

	switch name {
	case "Andy": 
		fmt.Println("Hello Andy")
	case "Tama":
		fmt.Println("Hello Tama")
	default:
		fmt.Println("Hello, Boleh Kenalan")	
	}
	

	switch length := len(name); length>4{
		case true: 
			fmt.Println("nama terlalu panjang")
		case false:
			fmt.Println("nama sudah benar")
	}

	panjang :=len(name)

	switch{
		case panjang >4 : 
			fmt.Println("nama lebih dari 4 karakter")
		case panjang <4:
			fmt.Println("nama kurang dari 4 karakter")
		case panjang ==4 :
			fmt.Println("nama sama dengan 4 karakter")
	}
}