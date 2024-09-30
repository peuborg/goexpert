package main

import (
	"context"
	"io"
	"net/http"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	//Objeto de Request
	req, err := http.NewRequestWithContext(ctx, "GET", "http://www.google.com", nil)
	if err != nil {
		panic(err)
	}
	//response
	resp, err := http.DefaultClient.Do(req)
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
