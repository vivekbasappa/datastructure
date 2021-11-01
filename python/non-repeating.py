from collections import OrderedDict

# Implement your function below.
def non_repeating(given_string):
	str_map = OrderedDict()
	for character in given_string:
		if character not in str_map:
			str_map[character] = 1
		else:
			str_map[character] += 1
	for key, value in str_map.items():
		if value == 1:
			return key
	else:
		return None


# NOTE: The following input values will be used for testing your solution.
print(non_repeating("abcab"))  # should return 'c'
print(non_repeating("abab"))  # should return None
print(non_repeating("aabbbc"))  # should return 'c'
print(non_repeating("aabbdbc"))  # should return 'd
