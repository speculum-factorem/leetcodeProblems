/**
 * @param {number} x
 * @return {number}
 */
var reverse = function(x) {

    let otric
    
    if (x < 0) {
        x *= -1
        otric = true
    } else {
        otric = false
    }
    
    let reverResult = 0

    while (x > 0) {
        const lastNum = x % 10
        reverResult = reverResult * 10 + lastNum
        x = Math.floor(x/10)
    }

    if (otric) {
        reverResult *= -1
    }
    if ((-2)**31 < reverResult && reverResult < (2**31-1)) {
        return reverResult
    } else {
        return 0
    }
};

console.log(reverse(1534236469))
console.log(9646324351 < 2**31-1)