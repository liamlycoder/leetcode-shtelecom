class Solution:
    def isNumber(self, s):
        try:
            answer = float(s)
            return True
        except:
            return False


if __name__ == '__main__':
    result = "5 4 32".replace(" ", "")
    print(result)
