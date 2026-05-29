class Solution:
    def maximalRectangle(self, matrix: list[list[str]]) -> int:
        cumlative = [0]*len(matrix[0])
        result = 0
        for i in range(len(matrix)):
            stk = [] # (value, index)
            for j in range(len(matrix[0])):
                cumlative[j] = cumlative[j] + 1 if matrix[i][j] == '1' else 0
                start = j
                while stk and stk[-1][0] > cumlative[j]:
                    value, start = stk.pop()
                    result = max(result, value*(j-start))
                stk.append((cumlative[j], start))

            while stk:
                value, start = stk.pop()
                result = max(result, value*(len(matrix[0])-start))
        return result

        