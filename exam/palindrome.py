def function(n):
    max_s = ''
    s = ''
    for i in range(len(n)):
        s += n[i]
        for j in range(i + 1, len(n)):
               
                s += n[j]
                if s != s[::-1]:
                    continue
                    # s+= n[j]
                elif s == s[::-1]:
                

                    if len(s) > len(max_s):
                        max_s = s
                        s = ''
                        # break
    return max_s
a = function(n=input("Enter the word: "))
print(a)
