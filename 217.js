const containsDuplicateArray = (nums) => {

    const collection = []

    for ( let i = 0; i < nums.length; i++) {
        if (collection.includes(nums[i])) {
            return true
        } else {
            collection.push(nums[i])
        }
    }

    return false
}


//217
const containsDuplicate = (nums) => {

    const collection = new Set() //Самое главвное то, что вместо массива тут используется
                                //множесто Set, что намного ускоряет скорость алгоритма
    for ( let i = 0; i < nums.length; i++) {
        if (collection.has(nums[i])) {
            return true
        } else {
            collection.add(nums[i])
        }
    }

    return false
}


console.log(containsDuplicate([1,2,3,4]))