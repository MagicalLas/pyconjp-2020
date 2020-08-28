package http

import "encoding/json"

func j() {
	user := map[string]interface{}{
		"name": "Las",
		"age":  19,
		"love": []map[string]interface{}{
			{
			"name": "programming",
			"score": 2,
			},
			{
				"name": "python",
				"score": 2,
			},
		},
	}
	bytes, err := json.Marshal(user)
	if err != nil {
		// error catch
	}
	print(string(bytes))
}
