#
# @lc app=leetcode id=48 lang=python3
#
# [48] Rotate Image
#

# @lc code=start
class Solution:
  
    def print_matrix(self, matrix: List[List[int]]):
        for row in matrix:
            print(row)
        
    def rotate(self, matrix: List[List[int]]) -> None:
        """ 
        Rotate the picture 90 degree
        Do not return anything, modify matrix in-place instead.
        """
        
        # 1.inverse each row
        # 2.swap numbers using the diagonal up right corner to down left
        len_matrix = len(matrix)  
        for row in matrix:
            for j in range((int(len_matrix/2))):
                temp = row[j]
                row[j] = row[-j-1]
                row[-j-1] = temp
        
        for i in range(len_matrix):
            for j in range(len_matrix):
                if i+j >= len_matrix-1:
                    break          
                temp = matrix[i][j]
                matrix[i][j] = matrix[-j-1][-i-1]
                matrix[-j-1][-i-1] = temp
        return
# @lc code=end

