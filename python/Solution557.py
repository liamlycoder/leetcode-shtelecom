class Solution:
    def reverseWords(self, s):
        word_list = s.split(" ")
        for i in range(len(word_list)):
            word_list[i] = word_list[i][::-1]
        result = ""
        for i in range(len(word_list)):
            result += word_list[i]
            if i != len(word_list) - 1:
                result += " "
        return result


if __name__ == '__main__':
    result = Solution().reverseWords("take LeetCode contest")
    print(result)
