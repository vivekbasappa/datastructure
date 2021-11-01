def is_one_away(s1, s2):
	len_s1 = len(s1)
	len_s2 = len(s2)
	if len_s1 - len_s2 > 1 or len_s2 - len_s1 > 1:
		return False
	if len_s1 == len_s2:
		return both_string_of_same_lenght(s1, s2)
	elif len_s2 > len_s2:
		return both_string_of_diff_length(s1, s2)
	else:
		return both_string_of_diff_length(s2, s1)


def both_string_of_same_lenght(s1, s2):
	count_diff = 0 
	for index, val in enumerate(s1):
		if s1[index] != s2[index]:
			count_diff += 1
			if count_diff > 1: 
				return False
	return True


def both_string_of_diff_length(s1, s2):
		count_diff = 0
		i = 0 
		while i + count_diff < len(s1):
			if s1[i + count_diff] == s2[i]:
				i += 1
			else:
				count_diff += 1
				if count_diff > 1:
					return False
		return True

# NOTE: The following input values will be used for testing your solution.
print(is_one_away("abcdee", "abcd"))  # should return True
print(is_one_away("abde", "abcde"))  # should return True
print(is_one_away("a", "a"))  # should return True
print(is_one_away("abcdef", "abqdef")) # should return True
print(is_one_away("abcdef", "abccef"))  # should return True
print(is_one_away("abcdef", "abcde"))  # should return True
print(is_one_away("aaa", "abc"))  # should return False
print(is_one_away("abcde", "abc"))  # should return False
print(is_one_away("abc", "abcde"))  # should return False
print(is_one_away("abc", "bcc"))  # should return False

