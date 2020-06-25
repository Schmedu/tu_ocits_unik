import time

import requests

r = requests.get('http://127.0.0.1:5000/stop/python_App')

print("being busy...")
time.sleep(5)
print("..bye")
