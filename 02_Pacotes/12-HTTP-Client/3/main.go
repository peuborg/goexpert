package main

import (
	"io"
	"net/http"
)

func main() {
	//clientHTTP
	c := http.Client{}
	//Objeto de Request
	req, err := http.NewRequest("GET", "http://www.google.com", nil)
	if err != nil {
		panic(err)
	}
	//header da request
	req.Header.Set("Accept", "application/json")
	//response
	resp, err := c.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	//body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))
}
