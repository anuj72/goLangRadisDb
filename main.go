package main

import (
	"fmt"
)


func main(){
	var myquery="Hyder"
	var LocalDb=getRedishData(myquery)
	if len(LocalDb) > 0 {
		fmt.Println("kkkk",LocalDb)
	}else {
		var clientData=getClientData(myquery)
		setRedishData(clientData,myquery)
		fmt.Println("yyyy",clientData)
	}

}

