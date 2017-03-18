package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Route struct {
	Summary string `json:"summary"`
}
type Response struct {
	Routes []Route `json:"routes"`
}

func main() {
	body := get()
	fmt.Println(string(body))

	response := new(Response)
	if err := json.Unmarshal(body, &response); err != nil {
		log.Println("JSON Unmarshal error:", err)
		return
	}
	fmt.Println(string(response.Routes[0].Summary))
}
func get() []byte {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://maps.googleapis.com/maps/api/directions/json",
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
