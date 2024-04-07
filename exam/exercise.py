n = input('Enter the words: ')

with open('t.txt', 'w') as file:
    file.write(n)
with open('t.txt', 'r') as file:
    content = file.read()
    print(content)
    content = content.split()
    print(content)
    m = content[0]
 
    for i in content:
        if len(i) > len(m):
            m = i
    print(f"maximum size is {m}")