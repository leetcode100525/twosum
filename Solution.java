import java.util.HashMap;

public class Solution {
    public int[] twoSum(int[] nums, int target) {
    HashMap<Integer,Integer> seen = new HashMap<>();
    for(int i=0;i<nums.length;i++){
        int num = nums[i];
        int dif = target-num;
        if(seen.containsKey(dif)){
            return new int[] {i,seen.get(dif)};
        }
        seen.put(nums[i],i);
    }
    return new int[0];
 }
}
