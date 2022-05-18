package main
import "fmt" 

func main(){

	 person :=map[string]string{
		 "name":"Tama",
		 "address": "Bogor",
		 
	 }
	 person["title"]= "programmer"

	fmt.Println(person) //map[address:Bogor name:Tama title:programmer]
	fmt.Println(len(person)) //3
	fmt.Println(person["name"])//Tama
	fmt.Println(person["address"])//Bogor


	var book map[string]string = make(map[string]string)
	book["title"]= "Belajar Golang"
	book["name"]= "Tamtama"
	book["ups"] = "salah"
	fmt.Println(book)//map[name:Tamtama title:Belajar Golang ups:salah]
	fmt.Println(len(book))//3

	delete(book,"ups")

	fmt.Println(book)//map[name:Tamtama title:Belajar Golang]
	fmt.Println(len(book))//2

}