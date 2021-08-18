package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://api.sleeper.app/v1/user/445264202878152704")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(body))
}
