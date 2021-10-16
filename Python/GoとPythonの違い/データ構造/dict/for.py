d = {"apple":2, "banana": 3, "orange": 1}

# keyを出力する  d.keys()でも同じ
for a in d:
    print(a)

print("*" * 20)

# key, valueを出力する
for a, b in d.items():
    print(a, b)

print("*" * 20)

# valueを出力する
for a in d.values():
    print(a)
