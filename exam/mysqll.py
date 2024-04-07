import mysql.connector

mydatabase = mysql.connector.connect(
    host = 'localhost',
    username = 'root',
    password = 'Dilshod@2005',
    database = 'learn',
)
mycursor = mydatabase.cursor()

query = 'create table if not exists me(name varchar(36), username varchar(36), age int)'

mycursor.execute(query)
mydatabase.commit()

sql = 'insert into me(name, username, age) values(%s, %s, %s)'
data = [
    ('dilshod', 'dilmurodov', 18),
    ('alibek', 'alibekov', 15),
    ('valibek', 'valibekov', 17),
    ('salim', 'salimov', 23),
    ('olim', 'olimov', 22),

]

mycursor.executemany(sql, data)
mydatabase.commit()
mycursor.execute('select * from me')
for x in mycursor.fetchall():
    print(x)
