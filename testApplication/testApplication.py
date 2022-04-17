import requests
from datetime import datetime
from threading import Thread
from random import randint

headers = {
    "User-Agent": """Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.99 Safari/537.36"""
}


def createRandomInput():
    return {
        "ip": "192.92.168.2",
        "ua": "Mozilla/5.0",
        "utm": "utm_medium=facebook",
        "url": "https://google.com",
        "sessionId": "12345",
        "clientId": randint(1, 10000)
    }


def makeRequest():
    time1 = datetime.now()
    r = requests.post('http://18.228.8.134:80/access',
                      json=createRandomInput(), headers=headers)
    time2 = datetime.now()
    print(f"{time1}")
    print(r.text)
    print(f"{time2}")


def makeManyRequests(nRequests):
    for i in range(0, nRequests):
        Thread(target=makeRequest).start()


if __name__ == '__main__':
    makeManyRequests(1000)
