// const a = {
//     a: 10,
//     b:11
// }

// for (let i in a) {
//     if (a[i] == 11) {
//         console.log(i)
//     }
// }

/**
 * @param {number[]} nums
 * @return {number}
 */
var singleNumber = function(nums) {

    const collection = {}
    let setNum = new Set(nums)

    
    for (let item of setNum) {
        collection[item] = 0
    }

    for (let item of nums) {
        collection[item]++
    }

    for (let item in collection) {
        console.log(item)
        if (collection[item] == 1) {
            return item
        }
    }

};

singleNumber([2, 2, 1])