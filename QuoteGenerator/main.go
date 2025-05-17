package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const baseurl = "https://zenquotes.io/api/random"

type Quote struct {
	Quote  string `json:"q"`
	Author string `json:"a"`
}

func main() {
	fmt.Println("=================Quote of the Day==================")
	response, err := http.Get(baseurl)
	if err != nil {
		panic(err)
	}
	//Debugging Tool
	// data, err := io.ReadAll(response.Body)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(data))

	var quotedata []Quote
	if err := json.NewDecoder(response.Body).Decode(&quotedata); err != nil {
		panic(err)
	}
	fmt.Println(quotedata[0].Quote)
	fmt.Println("By", quotedata[0].Author)

}
