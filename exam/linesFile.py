n = int(input('enter how many line u want to write the txt: '))
with open('line.txt', 'w') as file:
    for i in range(n):
        a = input('enter the word: ')
        file.write(a + '\n')
with open('line.txt', 'r') as file:
    content = file.read()
    count = 0
    for i in content:
        if i == '\n':
            count+=1
    print(f"lines {count}")
    