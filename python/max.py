# Implement your function below.
def most_frequent(given_list):
    max_item = None
    item_map = {}
    for number in given_list:
        if number not in item_map:
            item_map[number] = 0
        else:
            item_map[number] += 1

    max_number = (None, None)
    for key, value in item_map.items():
        if max_number[1] == None or max_number[1] < value:
            max_number = (key, value)

    return max_number[0]


# NOTE: The following input values will be used for testing your solution.
list1 = [1, 3, 1, 3, 2, 1]
print(most_frequent(list1)) #should return 1
list2 = [3, 3, 1, 3, 2, 1]
print(most_frequent(list2)) # should return 3
list3 = []
print(most_frequent(list3)) # should return None
list4 = [0]
print(most_frequent(list4)) # should return 0
list5 = [0, -1, 10, 10, -1, 10, -1, -1, -1, 1]
print(most_frequent(list5)) # should return -1
