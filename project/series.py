class Solution:
    def maxSubArray(self, nums):
        # Initialize our variables using the first element.
        current_subarray = max_subarray = nums[0]
        
        # Start with the 2nd element since we already used the first one.
        for num in nums[1:]:
            # If current_subarray is negative, throw it away. Otherwise, keep adding to it.
            current_subarray = max(num, current_subarray + num)
            max_subarray = max(max_subarray, current_subarray)
        
        return max_subarray

nums = [ 2, -3, 4, -1, -6, 10, 2]
solution = Solution()
print(solution.maxSubArray(nums))