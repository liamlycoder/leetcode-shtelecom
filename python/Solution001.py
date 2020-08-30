class Solution:
    def twoSum(self, nums, target):
        map = {}
        for i in range(len(nums)):
            if nums[i] in map.keys():
                return [i, map[nums[i]]]
            map[target-nums[i]] = i   # i的键就是i的另一半
        return []


if __name__ == '__main__':
    solution = Solution().twoSum([2, 7, 11, 15], 9)
    print(solution)
