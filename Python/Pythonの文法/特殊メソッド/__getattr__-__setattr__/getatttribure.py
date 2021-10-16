"""
属性にないものが取得されたら、self.fooを呼び出す
"""

class Foo:
    def __init__(self):
        self.foo = "foo"
        self.bar = "bar"
    def __getattr__(self, name):
        return self.foo

f = Foo()
print(f.foo)
print(f.z)

class Bar:
    def __init__(self):
        self.foo = "foo"
        self.bar = "bar"
    def __getattribute__(self, name):
        # super()だとエラーになる
        return object.__getattribute__(self,"foo")

b = Bar()
print(b.foo)
print(b.z)

# foo
# foo
# foo
# foo
