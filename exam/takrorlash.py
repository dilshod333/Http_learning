# with open('s.txt', 'w') as file:
#     n = input('Enter a word: ')
#     file.write(n)

# with open('s.txt', 'r') as file:
#     content = file.read().split()
#     print(content)
#     lst = []
#     for i in content:
#         if len(list(i.lower())) == len(list(set(i.lower()))):
#             lst.append(i)
#     print(lst)
#     s = ' '.join(lst)

#     l = 0
#     r = len(lst) - 1
 
#     while l < r:
#         lst[l], lst[r] = lst[r], lst[l]
#         l += 1
#         r -= 1

#     print(lst)

    
#     reversed_lst = [word[::-1] for word in lst]
#     print(reversed_lst)
# def findit(s):
#     if not s:
#         return 0

#     max_len = ""  # Change to 0 for storing length
#     sub = ""

#     for char in s:
#         if char in sub:
#             idx = sub.index(char) + 1
#             sub = sub[idx:]
#         sub += char
#         max_len = max(len(max_len), len(sub))  # Update max_len with length of sub

#     return max_len

# n = input("Enter the word: ")
# a = findit(n)
# print(a)
# def findit(s):
#     if not s:
#         return ""

#     max_sub = ""
#     sub = ""

#     for char in s:
#         if char in sub:
#             idx = sub.index(char) + 1
#             sub = sub[idx:]
#         sub += char
#         if len(sub) > len(max_sub):
#             max_sub = sub

#     return max_sub

# n = input("Enter the word: ")
# a = findit(n)
# print(a)
# def findit(s):
#     if not s:
#         return 0

#     max_len = ""  
#     sub = ""

#     for char in s:
#         if char in sub:
#             idx = sub.index(char) + 1
#             sub = sub[idx:]
#         sub += char
#         if len(sub) > len(max_len):
#             max_len = sub 

#     return max_len

# n = input("Enter the word: ")
# a = findit(n)
# print(a)
s = "fgabn"
n = 'n'
s = max(s, n)
print(s)