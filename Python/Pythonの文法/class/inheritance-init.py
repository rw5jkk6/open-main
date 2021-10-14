class A:
    def __init__(self, a):
        self.a = a

class B(A):
    pass

b = B("test")
print(b.a) # test
