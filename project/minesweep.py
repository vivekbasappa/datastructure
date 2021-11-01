import random

n = 3
m = 2
matrix = [ [0] * n ] * n 
print(matrix)
random_list = random.sample(range(0, n*n), m)
for value in random_list:
	row = int(value / 2)
	column = value % n
	matrix[row][column] = 'M'
print(matrix)