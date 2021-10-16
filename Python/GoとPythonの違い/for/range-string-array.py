# open-main/Go言語/Goの文法/for/range-string-array.go 

name = ["satou", "suzuki", "yamada"]

for i, v in enumerate(name):
    print(i, v)

print("-" * 10)

for i, _ in enumerate(name):
    print(i)

print("-" * 10)

for v in name:
    print(v)
