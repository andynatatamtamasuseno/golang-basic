package main
import "fmt" 

func main(){
	var nint32 int32=10001
	var nint64 int64= int64(nint32)
	var nint8 int8= int8(nint32)
	fmt.Println("number int 32" ,nint32);
	fmt.Println("number int 64" ,nint64);
	fmt.Println("number int 8" ,nint8);

name:= "Eko";
var chart byte = name[0];
var eString string= string(chart);

fmt.Println("name",name);
fmt.Println("chart/byte", chart)
fmt.Println("string",eString)
}