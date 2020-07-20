# Example of Attribute error

class Parent:
    def __init__(self, param):
        self.v1 = param

class Child(Parent):
    def __init__(self, param):
        self.v2 = param
        #self.v1 = param

obj = Child(11)
print(obj.v2)
print(obj.v1)
