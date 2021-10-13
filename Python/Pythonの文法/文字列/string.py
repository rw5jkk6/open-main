str = ["Hitoshi", "serina", "akari"]

# 1.  "Hitoshi, serina, akari"
print(','.join(str))
s = ','.join(str)

# 2. "serina"は含まれているか
print("serina" in str)  # True

# 3. "serina好き"を100回繰り返す
print("serina好き" * 100)

# 4. strの"serina"を"haruka"に書き換える
print(s.replace("serina", "haruka"))
# 5. strの文字列からスペースを取り除く
print(s.replace(" ", ""))
# 6. strを大文字にする
print(s.upper())
# 7. "Hitoshi, serina, akari" -> ["Hitoshi", "serina", "akari"]に変換する
print(s.split(','))
