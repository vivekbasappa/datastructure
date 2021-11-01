class Solution(object):
    def lengthOfLongestSubstringKDistinct(self, s, k):
        """
        :type s: str
        :type k: int
        :rtype: int
        """
        heighest = 0
        substring = ''
        if len(s) == 0:
            return 0
        char_map = {}
        
        for character in s:
            print(char_map)
            if character not in substring:
                char_map[character] = 1
                substring += character 
            else:
                length_substring = len(substring)
                if length_substring > heighest:
                    heighest = length_substring
                for index, c in enumerate(substring):
                    if c == character and char_map[character] < k :
                        char_map[character] += 1
                        substring += character
                        print("split1:", substring)
                    elif char_map[character] == k :
                        char_map = { character : 1 }
                        print("split2:index", substring, ":", index)
                        substring = substring[index+2:] + character
                        

        print("split3:", substring)                
        length_substring = len(substring)
        if length_substring > heighest:
            heighest = length_substring
        
        return heighest 

sol = Solution()
print(sol.lengthOfLongestSubstringKDistinct("eceba", 2))