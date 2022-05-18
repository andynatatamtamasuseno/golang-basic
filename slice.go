package main
import "fmt" 

func main(){

	var months = [...]string{
		"January",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"Sepmtember",
		"Oktober",
		"November",
		"Desember",
	}
	
	slice1:= months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1)) //3
	fmt.Println(cap(slice1)) //8

	months[5] ="DiUbah"
	fmt.Println(slice1) //[Mei DiUbah Juli]


	slice1[0]= "Mei Lagi"

	fmt.Println(months) //[January Februari Maret April Mei Lagi DiUbah Juli Agustus Sepmtember Oktober November Desember]



	slice2:= months[10:]
	fmt.Println(slice2)
	



	slice3:= append(slice2,"Tama")
	fmt.Println(slice3)
	slice3[1]="Bukan Desember"
	fmt.Println(slice3)//[November Bukan Desember Tama]

	fmt.Println(slice2) //[November Desember]
	fmt.Println(months) //[January Februari Maret April Mei Lagi DiUbah Juli Agustus Sepmtember Oktober November Desember]

	slice21:= months[2:4]
	fmt.Println(slice21)
	slice31:= append(slice21,"Tama")
	fmt.Println(slice31)
	slice31[1]="Bukan Desember"
	fmt.Println(slice31)//[November Bukan Desember Tama]

	fmt.Println(slice21) //[Maret Bukan Desember]
	fmt.Println(months) //[January Februari Maret Bukan Desember Tama DiUbah Juli Agustus Sepmtember Oktober November Desember]
	

	newSlice := make([]string,2, 5)

	newSlice[0] = "Tamtama"
	newSlice[1] = "Suseno"
	
	fmt.Println(newSlice) //[Tamtama Suseno]
	fmt.Println(len(newSlice))//2
	fmt.Println(cap(newSlice))//5


	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)//[Tamtama Suseno]


	iniArray := [5]int{1,2,3,4,5}
	iniSlice := []int{1,2,3,4,5}

	fmt.Println(iniArray)//[1 2 3 4 5]
	fmt.Println(iniSlice)//[1 2 3 4 5]



}