class A:
    add = None
    def __del__(self):
        print("Bye")

class B:
    add = None
    def __del__(self):
        print("Bye")

a = A()
b = B()
a.add = b
b.add = a
print(a.add)

# <__main__.B object at 0x103fd6790>
# Bye
# Bye

# 循環参照しないようだ
