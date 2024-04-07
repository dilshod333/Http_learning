# # # # class Bookstore:
# # # #     def __init__(self, name, author, year_publish,price):
# # # #         self.name = name
# # # #         self.author = author
# # # #         self.year_publish = year_publish
# # # #         self.price = price
# # # # class Library(Bookstore):
# # # #     def __init__(self, name, author, year_publish,price,sale):
# # # #         super().__init__(name, author,year_publish,price)
# # # #         self.sale = sale
# # # #     def giveSale(self):
# # # #         self.price -= (self.price * (self.sale/ 100))
# # # #         return f"{self.name}, {self.author},after sale price is {self.price} "
# # # # book = []
# # # # for i in range(5):
# # # #     name = input("Enter the name: ")
# # # #     author = input("Enter the author: ")
# # # #     year_publish = input('Enter the year_publish: ')
# # # #     price = float(input('Enter the price: '))
# # # #     sale = int(input("Enter the sale percent: "))
# # # #     book.append(Library(name, author, year_publish,price, sale))
# # # # for i in book:
# # # #     print(i.giveSale())

# # # import mysql.connector

# # # mydatabse = mysql.connector.connect(
# # #     host = 'localhost',
# # #     username = 'root',
# # #     password = 'Dilshod@2005',
# # #     database = 'Learn',
# # # )

# # # mycursor = mydatabse.cursor()
# # # query = 'create table if not exists bookshop(id int auto_increment primary key,name varchar(50),author varchar(50),year_publish int,price real, sale int,genre varchar(50),len_page int )'
# # # mycursor.execute(query)

# # # query = "insert into bookshop(name, author, year_publish, price,sale,genre, len_page  values(%s,%s,%s,%s,%s,%s,%s))"
# # # data = [
# # #     ("Vincent Theo","Raymund Slayny",2000,500,10,"westren",500),
# # #     ("charlie chan's", "nessy gibby", 2001, 600,15, "animation",600),
# # #     ("pin", "sheillah halliwell", 2002, 700, 5,'drama', 300),
# # #     ("same old song on connal", "cully book", 2003,800,25,"thriller",800),
# # #     ('santa claus',  'leslie hasling',2004,900,10,"Drama",900),
# # #     ('north atlantic', 'nattasa turfold', 2005,1000,15,'drama',1000),
# # #     ('heist', 'rick andrews', 2006, 1100, 5,'documatry', 1100),
# # #     ('game of chance', 'aldin McConnal',2008,1300,10,'threiller',1300),
# # #     ('anne frank', 'agnese sproji', 2009, 1400,15,'westren,1400',1000),
    

# # # ]
# # # mycursor.executemany(query,data)
# # # mycursor.execute("select * from bookshop")

# # # query = 'select * from bookshop where castyear_publish  > 10'
# # with open('t.txt', 'w') as file:
    
# #     a = input("Enter the word: ")
# #     file.write(a)
# # with open('t.txt', 'r') as file:
# #     res = ''
# #     content = file.read()
# #     print(content)
# #     empty_lst = []
# #     a = 0
# #     for i in content:
# #         for j in range(i+ 1, len(content)):
# #             if content[i] != ' ' and content[j] != ' ':
# #                 if content[i] == content[j]:
# #                     break
# #                 else:
# #                     res += content[i]

        
              
        
# #     print(res)

# n = int(input("Enter the number: "))
# count = 0
# i  =1
# if str(n) == str(n)[::-1]:
#     print("palindrome")
# else:
#     while i < 1000:
#         a = int(str(n)[::-1])
#         n = n + a 
#         if str(n) != str(n)[::-1]:
#             count+=1
#             i+=1
#         if str(n) == str(n)[::-1]:
#             print(n, count)
#             break
           


with open('t.txt', 'w') as file:
    
    a = input("Enter the word: ")
    file.write(a)
with open('t.txt', 'r') as file:
    res = ''
    content = file.read().split()
    lst = []
    for i in content:
        count =+ 1
        for j in range(i, len(content)):
            if content[i] != content[count: i]:
                res+= i
    s = ""
    s2= ""
    for i in range(len(res)):

        if i % 2 == 0:
            s += res[i]
        else:
            s2+= res[i]
    print(s2 + s)
            
    


   


        


