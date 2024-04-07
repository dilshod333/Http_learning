def change(n):
    mid = len(n) // 2
    first = n[:mid]
    second = n[mid:]

    res = second + first
    return res

a =change(n=input('enter the word: '))
print(a)