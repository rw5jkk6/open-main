import json

src = {
    "Name":"Hitoshi",
    "Pl":3,
    "BeastPl":5
}

print(src)
# {'Name': 'Hitoshi', 'Pl': 3, 'BeastPl': 5}

# liblary -> json
print(json.dumps(src))
# {"Name": "Hitoshi", "Pl": 3, "BeastPl": 5}

# json -> liblary
a = json.dumps(src)
print(json.loads(a))
# {'Name': 'Hitoshi', 'Pl': 3, 'BeastPl': 5}
