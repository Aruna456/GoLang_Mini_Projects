package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const baseURL = "http://ip-api.com/json/"

type GeoLoc struct {
	Country string  `json:"country"`
	Region  string  `json:"regionname"`
	City    string  `json:"city"`
	Zip     string  `json:"zip"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
	Status  string  `json:"status"`
	Message string  `json:"message"`
}

func main() {
	fmt.Println("===============Welcome to GeoIpFinder==================")
	fmt.Println("Enter the Ip address:")
	var ip string
	_, err := fmt.Scanln(&ip)
	if err != nil {
		panic(err)
	}
	fullurl := fmt.Sprintf("%s%s", baseURL, ip)
	response, err := http.Get(fullurl)
	if err != nil {
		panic(err)
	}
	var data GeoLoc
	json.NewDecoder(response.Body).Decode(&data)
	if data.Status != "success" {
		fmt.Println("Error:", data.Message)
		return
	}
	fmt.Printf("Fetched Location for IP:%s\n", ip)
	fmt.Println("Country:", data.Country)
	fmt.Println("Region:", data.Region)
	fmt.Println("City:", data.City)
	fmt.Println("Latitude:", data.Lat)
	fmt.Println("Longitude:", data.Lon)
	fmt.Println("ZipCode:", data.Zip)
	fmt.Println("==============Thank You=============")

}
