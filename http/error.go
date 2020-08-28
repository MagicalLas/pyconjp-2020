package http

import "net/http"

func Do() (*http.Response, error){
	res, err := http.Get("worng url")
	if err != nil {
		return nil, err
	}
	return res, nil
}
