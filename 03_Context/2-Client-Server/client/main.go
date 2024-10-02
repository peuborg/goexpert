package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	//Objeto de Request
	req, err := http.NewRequestWithContext(ctx, "GET", "http://192.168.16.175:8080", nil)
	if err != nil {
		panic(err)
	}
	//response
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}
