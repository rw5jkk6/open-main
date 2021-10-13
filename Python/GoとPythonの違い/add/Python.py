def add(n, m):
    return n + m

num = add(1, 2)
print(num) // 3


def calc_multi_return(n, m):
    a = n // m
    b = n % m
    return a, b

num = calc_multi_return(3, 2)
print(f'割り算は{num[0]} 剰余は{num[1]}')
// 割り算は1 剰余は1
