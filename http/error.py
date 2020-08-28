import requests

try:
    response = requests.get("https://pyconjp.jp")
except Exception as e:
    print(e)

print(response.text)
