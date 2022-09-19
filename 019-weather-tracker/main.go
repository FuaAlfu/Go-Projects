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

func query(city string)(weatherData, error){
	apiConfig, err := loginApiConfig(".apiConfig")
	if err != nil{
		return weatherData{}, err
	}

	http.GET("xxx")
	if err != nil{
		return weatherData{}, err
	}

	defer resp.Body.Close()
}

func hi(w http.ResponseWriter, r *http.Request){
	w.Header([]byte("welcom bb.. much love from golang\n"))
}

func server(){
	http.HandleFunc("/hi", hi)
	http.HandleFunc("/weather",
    func(w http.ResponseWriter, r *http.Request){
		city := strings.SplitN(e.URL.Path, "/", 3)[2]
		data, err := query(city)
		if err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=utf8")
		json.NewEncoder(w).Encode(data)
	})
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