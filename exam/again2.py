with open('t.txt', 'w') as file:
    n = input("Enter the word: ")
    file.write(n)
with open('t.txt', 'r') as file:
    content = file.read().split()
    lst = []
    for i in content:
        a = i
        if len(a) == len(set(i)):
            lst.append(i)
    print(lst)
    res = ''
    res1 = ''
    for i in lst:
        i = i.strip()
        res += i[len(i) // 2:]
        res += i[:len(i) // 2]
        print(res)
        res = ''

       

   


        

