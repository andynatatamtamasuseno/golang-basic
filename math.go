package main
import "fmt" 

func main(){
	a:=10;
	b:=5;
	jumlah := a+b;
	kurang := a-b;
	kali := a*b;
	bagi := a/b;
	modulus := a%b;
	fmt.Println("ketika a=10 & b= 5",jumlah );
	fmt.Println("jumlah",jumlah );
	fmt.Println("kurang",kurang );
	fmt.Println("kali",kali );
	fmt.Println("bagi",bagi );
	fmt.Println("modulus",modulus );
}