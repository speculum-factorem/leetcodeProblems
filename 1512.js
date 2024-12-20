const numIdenticalPairs = (nums) => {

    let resultCount = 0
    
    for (let i = 0; i < nums.length-1; i++) {
        for (let j = i+1; j < nums.length; j++) {
            if (nums[j] == nums[i]) {
                resultCount++
            }
        }
    }

    return resultCount
};

console.log(numIdenticalPairs([1,2,3,1,1,3])) //4