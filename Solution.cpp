#include <vector>
#include <map>
using namespace std;
class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        map<int, int> nmap;
        for(int i=0; i< nums.size(); i++){
            int c = target - nums[i];
            if(nmap.count(c)){
                int idx = nmap[c];
                return {idx,i};
            }
            nmap[nums[i]]=i;
        }
        return {};
    }
};