package main

import (
	"fmt"
	"os"

	"github.com/CalvoM/DexParser/fileparser"
)

func checkError(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
	}
}
func main() {
	var products []fileparser.Product
	var headers = []string{"DXS", "ST", "G85", "SE", "DXE"}
	fileID, err := os.Open("./source/ddcmp.txt")
	checkError(err, "(Main) Opening File")
	fmt.Println(headers)
	var index = 0
	for count := 0; count < 147; count++ {
		line, byteCount := fileparser.StreamLine(*fileID, index)
		dataElement := fileparser.GetFirstDataElement(line, "*")
		if fileparser.IsHeaderPresent(dataElement, headers) != true {
			if dataElement == "LA1" {
				products = append(products, fileparser.ParseProduct(line))
				fmt.Println(products)
			}
			if dataElement == "EA1" {
				fmt.Println(line)
			}
		}
		index += byteCount
	}
}
