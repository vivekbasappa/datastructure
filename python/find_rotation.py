# Implement your function below.
def is_rotation(list1, list2):
	if len(list1) != len(list2):
		return False

	j = -1

	while j < len(list2) :
		if list1[0] == list2[j]:
			break
		j += 1	
	else:
		return False	

	for index, value in enumerate(list1):
		i = (j + index) % len(list1)
		if list1[index] != list2[j]:
			return False
  
		return True


# NOTE: The following input values will be used for testing your solution.
list1 = [1, 2, 3, 4, 5, 6, 7]
list2a = [4, 5, 6, 7, 8, 1, 2, 3]
print(is_rotation(list1, list2a)) #should return False.
list2b = [4, 5, 6, 7, 1, 2, 3]
print(is_rotation(list1, list2b)) #should return True.
list2c = [4, 5, 6, 9, 1, 2, 3]
print(is_rotation(list1, list2c)) #should return False.
list2d = [4, 6, 5, 7, 1, 2, 3]
print(is_rotation(list1, list2d)) #should return False.
list2e = [4, 5, 6, 7, 0, 2, 3]
print(is_rotation(list1, list2e)) #should return False.
list2f = [1, 2, 3, 4, 5, 6, 7]
print(is_rotation(list1, list2f)) #should return True.
list2g = [7, 1, 2, 3, 4, 5, 6]
print(is_rotation(list1, list2g)) #should return True.
