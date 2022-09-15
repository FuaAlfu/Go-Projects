package main

import(
	"fmt"
	"log"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type(
	apiConfigData struct{
		OpenWeatherMapApiKey string `json: "OpenWeatherMapApiKey"`
	}

	weatherData struct{
		Name string `json: "name"`
		Main struct{
			ke float64 `json: "temp"`
		} `json: "main"`
	}
)

func loadApiConfig(filename string)(apiConfigData, error){
	bytes, err := ioutil.ReadFile(filename)
	if err != nil{
		return apiConfigData{}, nil
	}

	var c apiConfigData

	err = json.Unmarshal(bytes, &c)
	if err != nil{
		return apiConfigData{}, err
	}
	return c, nil
}

func server(){
	http.HandleFunc("/hi", hi)
	//http.ListenAndServe(":9090",nil)
	//or
	fmt.Printf("servering at port 9090\n")
	if err := http.ListenAndServe(":9090",nil); err != nil{
		log.Fatal(err)
	}
}

func main(){
	server()
}