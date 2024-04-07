s = 'salom' #expected (e(x(a(m)p)l)e), 'card' expected: (c(ar)d)
res = ''
mid = len(s) // 2 


for i in range(len(s)):
    if mid == i and not len(s) % 2:
        res += s[i]
    else:
        res += ')' + s[i] if mid < i else '(' + s[i]

res += ')'
print(res)
