class Solution:
    def minArray(self, numbers):
        l, r = 0, len(numbers) - 1
        while l < r:
            mid = ((r - l) >> 1) + l
            if numbers[r] > numbers[mid]:
                r = mid
            elif numbers[r] < numbers[mid]:
                l = mid + 1
            else:
                r -= 1
        return numbers[l]

if __name__ == '__main__':
    print(Solution().minArray([3,4,5,1,2]))
