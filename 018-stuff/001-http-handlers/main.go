package main

import(
	"fmt"
	"io"
	"compress/gzip"
	"net/http"
)

//example..
func dataHandler(w http.ResponseWriter, r *http.Request){
	var wtr io.Writer = w
	if r.Header.Get("Accept-Encoding") == "gzip"{
		w.Header().Set("Accept-Encoding", "gzip")
		gzw := gzip.NewWriter(w)
		defer gzw.Close()
		wtr = gzw
	}
	w.Header().Set("Content-Type", "text/csv")
	cw := csv.NewWriter(wtr)
	defer cw.Flush()
	for _, row := range query(){
		cw.Write(row)
	}
}

func main() {
	//dataHandler()
}