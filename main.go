package main

import (
	"fmt"
	Route2 "github.com/andrethiel/goland"
)

func main(){
	route := Route2.Route{
		ID: "1"
		ClientID: "1"
	}

	route.LoadPositions()

	stringjson, _ := route.ExportJsonPositions()
	fmt.Println(stringjson[0])

}