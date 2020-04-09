package main

import (
	"os"
	"strings"
)

type Product struct{
	Number string `json:"number"`
	Price string `json:"price"`
	Vends string `json:"vends"`
}
type Event struct {
	Type string `json:"type"`
	Duration uint64 `json:"duration"`
}
func StreamLine(fileID os.File,index int)(string, int){
	var buffer = make([]byte,0)
	var char = make([]byte,1)
	var count =0
	_,err :=fileID.ReadAt(char,int64(index))
	checkError(err,"(StreamLine) Error Reading")
	buffer = append(buffer,char[0])
	for char[0] !=10 {
		index++
		count++
		_, err = fileID.ReadAt(char,int64(index))
		checkError(err,"(StreamLine) Loop Error Reading")
		buffer = append(buffer,char[0])
	}
	count++
	return string(buffer),count
}
func GetFirstDataElement(line string)string{
 	dataElement := strings.Split(line,"*")[0]
 	return dataElement
}
func IsHeader(header string,headers []string)(isPresent bool){
	isPresent = false
	for _, h := range headers {
		if h==header{
			isPresent = true
			break
		}
	}
	return
}
func ParseProduct(line string) (product Product){
	var productDetails = strings.Split(line, "*")
	product.Number = productDetails[2]
	product.Price = productDetails[3]
	product.Vends = productDetails[5]
	return
}
