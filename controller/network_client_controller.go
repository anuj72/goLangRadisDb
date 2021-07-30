package controller

import (
	"RedisProject/config"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var key string = config.NewRepository().GooglePlacesKey()

func GetClientData(myquery string) string {
	req, err := http.NewRequest("GET",
		"https://maps.googleapis.com/maps/api/place/autocomplete/json",
		nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	q := req.URL.Query()
	q.Add("input", myquery)
	q.Add("key", key)
	req.URL.RawQuery = q.Encode()
	//fmt.Println(req.URL.String())
	resp, err := http.Get(req.URL.String())
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return string(body)
}
