class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        num_map = dict()
        for i, num in enumerate(nums):
            t = target-num
            if t in num_map:
                return [num_map[t], i]
            else:
                num_map[num] = i