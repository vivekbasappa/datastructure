def generateParenthesis(n):
		"""
		:type n: int
		:rtype: List[str]
		"""
		output_arr = []
		backtrack(output_arr, "", 0, 0 , n)
		return output_arr

def backtrack(output_arr, current_string, opend, closed, max):
		if len(current_string) == max * 2:
				output_arr.append(current_string)
				return;
		if opend < max :
				print ("curent string:", current_string, ":")
				backtrack(output_arr, current_string + "(", opend+1, closed, max)
		if closed < opend: 
				print ("curent string:", current_string, ":")
				backtrack(output_arr, current_string + ")" , opend, closed+1, max)

generateParenthesis(10)
