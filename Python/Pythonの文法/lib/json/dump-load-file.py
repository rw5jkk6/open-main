import json

src = {
    "Name":"Hitoshi",
    "Pl":3,
    "BeastPl":5
}

print(src)
# {'Name': 'Hitoshi', 'Pl': 3, 'BeastPl': 5}

# dict -> json(file)
with open("test.json", "w") as f:
    json.dump(src, f)


# json(file) -> dict
with open("test.json", "r") as f:
    a = json.load(f)
print(a)
# {'Name': 'Hitoshi', 'Pl': 3, 'BeastPl': 5}
