# here's the same example, but using Python - much more familiar to me :((
import requests

print("Hello, GO Standard Library!")

response = requests.get("http://jsonplaceholder.typicode.com/posts/1")

if response.status_code == 200:
    print(f"HTTP Response Status: {response.status_code}")

else:
    print(f"Error: {response.status_code} {response.reason}")
