class Solution(object):
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        hash_table = {}
        for i, val in enumerate(nums):
            diff = target - val
            if diff not in hash_table:
                hash_table[val] = i
            else:
                return [hash_table[diff], i]
