package http

import "encoding/json"

type Love struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}
type JSON struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Love []Love `json:"love"`
}

func j() {
	user := JSON{
		Name: "Las",
		Age: 19,
		Love: []Love{
			{
				Name: "programming",
				Score: 2,
			},
			{
				Name: "python",
				Score: 2,
			},
		},
	}
	bytes, err := json.Marshal(user)
	if err != nil {
		// error catch
	}
	print(string(bytes))
}
