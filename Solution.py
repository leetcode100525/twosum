from typing import List
class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        nmap = {}
        for i in range(len(nums)):
            c = target - nums[i]
            if c in nmap:
                idx = nmap[c]
                return [idx,i]
            nmap[nums[i]]=i
        return []