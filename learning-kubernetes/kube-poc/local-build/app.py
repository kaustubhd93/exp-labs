import os
import time 

while True:
    print("*************************************************")
    print(os.environ['CLIENT_ID'])
    print(os.environ['SCOPE'])
    print("*************************************************")
    print(os.environ)
    print("\n")
    time.sleep(10)



