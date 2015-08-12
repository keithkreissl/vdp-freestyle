package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"io/ioutil"
	"fmt"
)

/*func CityHandler(res http.ResponseWriter, req *http.Request) {
	data, _ := json.Marshal("{'cities':'San Francisco, Amsterdam, Berlin, New York, Tokyo, Chicago, Wheaton, blah, Glen Ellyn, Carol Stream'}")
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.Write(data)
}*/

type VehicleDetail struct{
	ListingDetail ListingDetailDto `json:"listingDetailDto"`
	LocalOffers LocalOffersDto `json:"localOffersDto"`
}

type ListingDetailDto struct{
	ListingId int `json:"listingId"`
	IsActiveListing bool `json:isActiveListing"`
	MakeName string `json:"makeName"`
	ModelName string `json:"modelName"`
	ModelYear int `json:"modelYear"`
	PriceFormatted string `json:"priceFormatted"`
}

type LocalOffersDto struct{
	Name string
}

func main() {
	port := os.Getenv("PORT") 
	if  port == "" {
		port = "5000"
	}
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/vehicledetail/detail", VehicleDetailHandler)
	err := http.ListenAndServe(":" + port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func IndexHandler(res http.ResponseWriter, req *http.Request){
	data, _ := json.Marshal("Vehicle Detail Page")
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.Write(data)
}

func VehicleDetailHandler(res http.ResponseWriter, req *http.Request) {
	vi := req.URL.Query().Get("vehicleId")
	resp, err := http.Get("http://www.cars.com/ajax/listingsapi/1.0/listing/detail/" + vi)
	
	if err != nil {
		//res writer handle err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	var vd VehicleDetail 
	json.Unmarshal(body, &vd)	

//	fmt.Printf("Keith the listing id is: %d\n", vd.ListingDetail.ListingId)

	res.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(res, "%d %s %s, %s", vd.ListingDetail.ModelYear, vd.ListingDetail.MakeName, vd.ListingDetail.ModelName, vd.ListingDetail.PriceFormatted)


}

