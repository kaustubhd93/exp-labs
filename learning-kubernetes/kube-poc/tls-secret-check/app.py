import os
import time 

while True:
    print("*************************************************")
    with open("/opt/api-keys/public_key.pem") as file_data:
        for line in file_data:
            print(line)
    print("*************************************************")
    with open("/opt/api-keys/private_key.pem") as file_data:
        for line in file_data:
            print(line)
    print("*************************************************")
    print("\n")
    time.sleep(10)



