import random 

with open('file.txt', 'w') as file:
    for i in range(5):
        r = input("Enter the number: ")
        file.write(r + '\n')

with open('file.txt', 'r') as file, open('file2.txt', 'w') as file2:
    content = file.read().split()
    for i in content:
        if int(i[0]) == 5:
            file2.write(str(i) + '\n')
with open('file.txt', 'r') as file, open('file3.txt', 'w') as file3:
    content = file.read().split()
    for i in content:
        if int(i[0]) != 5:
            file3.write(str(i) + '\n')

    
    