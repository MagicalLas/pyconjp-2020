package main

import (
	"net/http"

	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	response, err := http.Get("https://pycon.jp")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(response)

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(body))
}
