package main
import (
	"fmt"
	"os"
)
func checkError(err error,msg string){
	if err != nil {
		fmt.Println(msg)
	}
}
 func main(){
 	var products []Product;
 	var headers = []string{"DXS","ST","G85","SE","DXE"}
 	fileID,err := os.Open("./source/ddcmp.txt")
 	checkError(err,"(Main) Opening File")
 	fmt.Println(headers)
 	var index = 0
 	for count:=0;count<147;count++{
 		line,byteCount:=StreamLine(*fileID, index)
 		dataElement := GetFirstDataElement(line)
		if IsHeader(dataElement,headers)!=true{
			if dataElement == "LA1"{
				products = append(products,ParseProduct(line))
				fmt.Println(products)
			}
			if dataElement == "EA1" {
				fmt.Println(line)
			}
		}
 		index += byteCount
	}
 }