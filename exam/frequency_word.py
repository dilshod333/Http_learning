with open('freq.txt', 'w') as file:
    for i in range(3):
        n = input('enter the word: ')
        file.write(n + '\n')
with open('freq.txt', 'r') as file:
    content = file.read().split()
    print(content)
    dct = {}
    for i in content:
        if i in dct:
            dct[i] += 1
        else:
            dct[i] = 1
    print(dct)