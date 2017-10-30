package main

import (
	"net/http"
	"github.com/abh/geoip"
	"github.com/gorilla/mux"
	"io"
	"log"
	"os"
	"strings"
)

var gi *geoip.GeoIP

func getPNG(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	country, _ := gi.GetCountry(vars["ip"])

	imgname := "flags/"+strings.ToLower(country)+".png"
	img, err := os.Open(imgname)
	if err != nil {
	}

	//w.Write([]byte(img))
	io.Copy(w, img)
}

func main() {
	var err error
	file := "/usr/share/GeoIP/GeoIP.dat"

	gi, err = geoip.Open(file)
	if err != nil {
		log.Fatal("Could not open GeoIP database\n")
	}

	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/ip/{ip}/png", getPNG)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))

}
