package http

import "encoding/json"

const jsonString = `{
    "name": "Las",
    "age": 19,
    "love": [
        {
            "name": "programming",
            "score": 2
        },
        {
            "name": "python",
            "score": 2
        }
    ]
}`
func g() {
	user := map[string]interface{}{}
	err := json.Unmarshal([]byte(jsonString), user)
	if err != nil {
		// error catch
	}

	loves, ok := user["love"].([]map[string]interface{})
	if !ok {
		// error catch
	}

	name, ok := loves[0]["name"].(string)
	if !ok {
		// error catch
	}
	print(name)
}
