def function(n):

    max_s = ''
    s = ''
    for i in range(len(n)-1):
        s += n[i]
        for j in range(i + 1, len(n)):
            if n[j] not in s:
                s += n[j]
                if len(s) > len(max_s):
                    max_s = s
            else:
                s = ''
                break
    return 1 if s == '' else len(max_s)
a = function(n=input("Enter the word: "))
print(a)
