package main
import "fmt" 

func sayHello(){
	fmt.Println("Hello")
}
func penjumlahan(a int,b int){
	fmt.Println(a+b)
}

func pengebalian(name string)string{
	return "Hello "+ name
}


func multipleValue(name string)string{
	return "Hello "+ name
}

func getFullname ()(string,string,string){
	return "andynata", "tamtama", "suseno"
}

func getFullname2 ()(firstName string, middleName string, lastName string){ //return value tidak perlu di panggil ketika sudah di definisikan
	firstName= "Andynata"
	middleName= "Tamtama"
	lastName= "Suseno"

	return 
}

func main(){

	sayHello()
	penjumlahan(10,20)
	firstname,_,lastname := getFullname()

	fmt.Println(firstname,lastname) // andynata suseno

	fmt.Println(pengebalian("Tama"))// Hello nama

	fmt.Println(getFullname())//andynata tamtama suseno

	fmt.Println(getFullname2())//Andynata Tamtama Suseno
}