/**
 * @param {number} n
 * @return {number}
 */

var fib = function(n) {
    
    let canvas = new Array(n+1).fill([])

    for (let i = 0; i < canvas.length; i++) {
        if (i == 0) {
            canvas[i] = 0
        }
        if (i == 1 || i == 2) {
            canvas[i] = 1
        }
        if (i > 2) {
            canvas[i] = canvas[i-1] + canvas[i-2]
        }
    }

    return canvas[n]
};

console.log(fib(4))