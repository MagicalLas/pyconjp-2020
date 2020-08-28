import json

user = {
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
}

print(json.dumps(user))