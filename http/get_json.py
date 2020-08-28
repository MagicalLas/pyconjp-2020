import json

json_string = """{
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
}"""

user = json.loads(json_string)
print(user["love"][0]["name"])
