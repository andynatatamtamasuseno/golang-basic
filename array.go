package main
import "fmt" 

func main(){

	var arr [3]string
	arr[0] = "nama"
	arr[1] = "saya"
	arr[2] = "tama"

	fmt.Println(arr[0])
	fmt.Println(arr[1])
	fmt.Println(arr[2])

	var value = [3]int{90,95,80}

	fmt.Println(value)

	var panjang [10] string


	fmt.Println("panjang array", len(arr))
	fmt.Println("panjang value", len(value))
	fmt.Println("panjang panjang", len(panjang))

}