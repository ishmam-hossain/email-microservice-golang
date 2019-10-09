import requests

def send_1000_post_requests():
    for i in range(1000):
        req_data = {
            "id": i,
            "name": "ishmam",
            "age": 26,
            "email": "ishmam@finder-lbs.com"
        }

        response = requests.post("http://127.0.0.1:9999/post", json=req_data)
        print(response.text)

send_1000_post_requests()
