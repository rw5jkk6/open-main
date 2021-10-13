def f(n):
    s = 0
    for i in range(n):
        s += i
    return s

s = f(2<<25)
print(s)

# 2251799780130816

# real    0m3.337s
# user    0m3.325s
# sys     0m0.008s
