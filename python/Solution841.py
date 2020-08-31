class Solution:
    rooms_status = []

    def canVisitAllRooms(self, rooms):
        global rooms_status
        nums = len(rooms)
        rooms_status = [False for _ in range(nums)]
        rooms_status[0] = True
        self.dfs(rooms, 0)
        return False not in rooms_status

    def dfs(self, rooms, i):
        global rooms_status
        key_list = rooms[i]
        for key in key_list:
            if (not rooms_status[key]):
                rooms_status[key] = True
                self.dfs(rooms, key)


if __name__ == '__main__':
    result = Solution().canVisitAllRooms([[6, 7, 8], [5, 4, 9], [], [8], [4], [], [1, 9, 2, 3], [7], [6, 5], [2, 3, 1]])
    print(result)
