package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"io/ioutil"
	"fmt"
	"github.com/julienschmidt/httprouter"
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
	router := httprouter.New()
	router.GET("/", IndexHandler)
	router.GET("/vehicledetail/detail/:vehicleId/overview", VehicleDetailHandler)
	err := http.ListenAndServe(":" + port, router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	fmt.Fprint(w, "Welcome! to the Vdp in Go making a change!\n")
}

func VehicleDetailHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	resp, _ := http.Get("http://www.cars.com/ajax/listingsapi/1.0/listing/detail/" + ps.ByName("vehicleId"))
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var vd VehicleDetail 
	json.Unmarshal(body, &vd)	

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "%d %s %s, %s", vd.ListingDetail.ModelYear, vd.ListingDetail.MakeName, vd.ListingDetail.ModelName, vd.ListingDetail.PriceFormatted)
}

