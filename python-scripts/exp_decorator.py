# By definition, a decorator is a function that takes another function and extends 
# the behavior of the latter function without explicitly modifying it.
# For example: the @app.route() function in Flask

# This function can be kept in another file and imported as a module.
def decorate_this(func_obj):
    def func_to_be_decorated():
        print("Somthing before the function is called.")
        func_obj()
        print("Something after the function is called.")
    return func_to_be_decorated

# Applying a decorator to a function the pythonic way.
@decorate_this
def say_hello():
    print("Hello, there")

# Decorating function in a simple manner.
# print("Before decorating")
# say_hello()
# say_hello = decorate_this(say_hello)
# print("After decorating")
# say_hello()

say_hello()

