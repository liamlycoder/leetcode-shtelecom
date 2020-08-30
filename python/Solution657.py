class Solution:
    def judgeCircle(self, moves: str):
        move_table = {
            "L": 0,
            "R": 0,
            "U": 0,
            "D": 0
        }

        for i in moves:
            move_table[i] += 1
        return move_table['L'] == move_table['R'] and move_table['U'] == move_table['D']

if __name__ == '__main__':
    result = Solution().judgeCircle("UD")
    print(result)

