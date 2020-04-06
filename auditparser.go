package main

import (
	"os"
	"strings"
)

type Product struct{
	Number uint8 `json:"number"`
	Price uint8 `json:"price"`
	Vends uint8 `json:"vends"`
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