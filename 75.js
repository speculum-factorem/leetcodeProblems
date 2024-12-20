var sortColors = function(nums) {
    
    let count0 = 0
    let count1 = 0
    let count2 = 0
    
    result = []
 
    for (i of nums) {
        switch (i) {
            case 0:
                count0++
                break
            case 1:
                count1++
                break
            case 2:
                count2++
                break
        }
    }

    // let index = 0

    while (count0 > 0) {
        result.push(0)
        count0--
    }
    while (count1 != 0) {
        result.push(1)
        count1--
    }
    while (count2 != 0) {
        result.push(2)
        count2--
    }

    return result
};

console.log(sortColors([2,0,2,1,1,0]))