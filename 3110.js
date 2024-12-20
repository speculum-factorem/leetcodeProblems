// let a = "hello"
// console.log(a.charCodeAt(0))

const scoreOfString = (s) => {

    sum = 0

    for (let i = 0; i < s.length - 1; i++) {
        sum += Math.abs(s.charCodeAt(i) - s.charCodeAt(i + 1))
    }

    return sum
}

console.log(scoreOfString("hello")) //13
