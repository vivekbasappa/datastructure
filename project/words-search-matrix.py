class Solution(object):
    def findWords(self, board, words):
        """
        :type board: List[List[str]]
        :type words: List[str]
        :rtype: List[str]
        """
        self.ROWS = len(board)
        self.COLS = len(board[0])
        self.board = board

        words_found = []
        for word in words:
            for row in range(self.ROWS):
                for col in range(self.COLS):
                    if self.backtrack(row, col, word):
                        words_found.append(word)
        return words_found

    def backtrack(self, row, col, suffix):
        if len(suffix) == 0:
            return True

        if row < 0 or row == self.ROWS or col < 0 or col == self.COLS or \
                self.board[row][col] != suffix[0]:
            return False

        ret = False
        self.board[row][col] = '#'

        for row_offset, col_offset in [(0, 1), (1, 0), (0, -1), (-1, 0)]:
            ret = self.backtrack(row+row_offset, col + col_offset, suffix[1:])
            if ret:
                return True

        self.board[row][col] = suffix[0]

        return ret


board = [	["o", "a", "a", "n"], 
					["e", "t", "a", "e"], 
					["i", "h", "k", "r"], 
					["i", "n", "i", "a"]
				]
words = ["oath", "pea", "eat", "rain"]

solution = Solution()

print("is there a solution for", board, "with word ",
      words, " result:", solution.findWords(board, words))
