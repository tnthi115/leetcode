# @leet start
class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        diffs = {}
        for i, val in enumerate(nums):
            diff = target - val
            if diff in diffs:
                return [i, diffs.get(diff)]
            diffs[val] = i
        return []


# @leet end
