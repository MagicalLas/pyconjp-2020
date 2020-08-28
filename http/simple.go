package http

import (
	"net/http"

	"io/ioutil"
)

func main() {
	response, err := http.Get("https://pycon.jp")
	if err != nil {
		// error catch
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		// error catch
	}
	println(string(body))
}
