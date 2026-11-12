class Solution:
    def hasDuplicate(self, nums: List[int]) -> bool:
        is_present = set()
        for num in nums:
            if num in is_present:
                return True
            else:
                is_present.add(num)
        return False
            