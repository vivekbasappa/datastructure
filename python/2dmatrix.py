class Solution:
    def binary_search(self, matrix, target, start, vertical):
        lo = start
        hi = len(matrix[0]) - 1 if vertical else len(matrix) - 1

        while hi >= lo:
            mid = (lo + hi) // 2
            if vertical: # searching a column
                if matrix[start][mid] < target:
                    lo = mid + 1
                elif matrix[start][mid] > target:
                    hi = mid - 1
                else:
                    return True
            else: # searching a row
                if matrix[mid][start] < target:
                    lo = mid + 1
                elif matrix[mid][start] > target:
                    hi = mid - 1
                else:
                    return True
        
        return False

    def searchMatrix(self, matrix, target) -> bool:
        # an empty matrix obviously does not contain `target`
        if not matrix:
            return False

        # iterate over matrix diagonals starting in bottom left.
        for i in range(min(len(matrix), len(matrix[0]))):
            vertical_found = self.binary_search(matrix, target, i, True)
            horizontal_found = self.binary_search(matrix, target, i, False)
            if vertical_found or horizontal_found:
                return True
        
        return False

solution = Solution()
matrix = [
	[ 1, 7, 13, 19,25 , 31],
	[ 2, 8, 14, 20, 26, 32],
	[ 3, 9, 15, 21, 27, 33],
	[ 4, 10, 16, 22, 28, 34],
	[ 5, 11, 17, 23, 29, 35],
	[ 6, 12, 18, 24, 30, 36],
]
print("target found in matrix: ", solution.searchMatrix(matrix, 28))
