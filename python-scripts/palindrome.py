import sys

istring = sys.argv[1]
w = ""

for ele in istring:
    w = ele + w
    
if w == istring:
    print("Input string is a palindrome.")
else:
    print("Input string is not a palindrome.")

