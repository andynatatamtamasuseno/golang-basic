package main
import "fmt" 

func main(){

	counter :=1

	for counter <= 10{
		fmt.Println("perulangan ke",counter)
		counter+=1
	}

	for number:=1;number<=10; number++{
		fmt.Println("number ke",number)
	}

	sliceData := []string{"andynata","tamtama", "suseno"}

	for number:=0;number<len(sliceData); number++ {
		fmt.Println(sliceData[number])
	}

	for i, value:= range sliceData {
		fmt.Println("index",i,"value",value)
	}

	person := make(map[string]string)
	person["name"]= "Tama"
	person["address"] ="Visar"


	for key, value := range person{
		fmt.Println(key, "=",value)
	}

	for i:=0; i<100;i++{
		if i ==10{
			break
		}else if i%2 ==0{
			continue
		}

		fmt.Println("no-",i)
		
	}
}