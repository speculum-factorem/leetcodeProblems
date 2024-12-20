// let a = [1, 2, 3, 4]
// console.log(a.slice(0, 2), a.slice(2, 4))

const shuffle = (nums, n) => {

    let left = nums.slice(0, n)
    let right = nums.slice(n, nums.length)

    let result = []

    for (let i = 0; i < left.length; i++) {
        result.push(left[i], right[i])
    }

    return result
}

console.log(shuffle([1, 2, 3, 4], 2))