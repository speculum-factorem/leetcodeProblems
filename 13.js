//13 самый сок в объявлении свойств оъекта
const romanToInt = (s) => {

    const collection = {
        I: 1,
        V: 5,
        X: 10,
        L: 50,
        C: 100,
        D: 500,
        M: 1000
    }
    
    let sum = 0
    let i = 0

    while (i < s.length) {

        const curentSymbol = collection[s[i]]
        const nextSymbol = collection[s[i+1]]

        if (curentSymbol < nextSymbol) {
            sum += nextSymbol - curentSymbol
            i += 2
        } else {
            sum += curentSymbol
            i++
        }
    }

    return sum

}

console.log(romanToInt("III"))







// a = "12"
// console.log(a.length)

// const collection = {
//     "I": 1,
//     V: 5,
//     X: 10,
//     L: 50,
//     C: 100,
//     D: 500,
//     M: 1000
// }


// console.log(collection["I"])



