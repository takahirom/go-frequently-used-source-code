package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"encoding/xml"
)

type Route struct {
	Summary string `xml:"summary"`
}
type Response struct {
	Routes Route `xml:"route"`
}

func main() {
	body := get()
	fmt.Println(string(body))

	response := new(Response)
	if err := xml.Unmarshal(body, &response); err != nil {
		log.Println("JSON Unmarshal error:", err)
		return
	}
	fmt.Println(string(response.Routes.Summary))
}
func get() []byte {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://maps.googleapis.com/maps/api/directions/xml",
		nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	q := req.URL.Query()
	q.Add("origin", "tokyo")
	q.Add("destination", "osaka")
	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req)
	if err != nil {
		log.Print("Request error:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}