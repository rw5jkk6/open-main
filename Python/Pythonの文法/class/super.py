class A:
    def __init__(self, a):
        self.a = a

    def walk(self):
        print("walk")

class B(A):
    def __init__(self, a, b):
        super().__init__(a)
        self.b = b
    def walk_run(self):
        super().walk()
        print(f'{self.a}で{self.b}する。')
    
b = B("test","run")
b.walk_run()

# walk
# testでrunする。
