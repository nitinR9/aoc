const fs = require('fs')
const file = fs.readFileSync(__dirname + '/input.txt', 'utf-8')
const [keys, values] = file.split('\n').reduce((prev, value, i) => {
    let [key, string] = value.split(': ')
    string = string.split(' ').map(Number)
    prev[0].push(+key)
    prev[1].push(string)
    return prev
}, [[],[]])

function cartesianProduct(str, n) {
    if (n <= 0) return [];
    if (n === 1) return str.split("");
    
    const smallerProduct = cartesianProduct(str, n - 1);
    const result = [];
    
    for (const char of str) {
        for (const subset of smallerProduct) {
            result.push([char,...subset]);
        }
    }
    return result;
}

function totalCalibration(string){
    let sum = 0
    
    let testOpComb = (combo, arr) => {
        let result = arr[0]
        for(i=1; i<arr.length; i++){
            if (combo[i-1] === '*')
                result *= arr[i]
            else if (combo[i-1] === '+')
                result += arr[i]
            else if (combo[i-1] === '|')
                result = +`${result}${arr[i]}`
        }
        return result
    }

    for(let j=0; j<values.length; j++){
        let key = keys[j]
        let value = values[j]

        for(let combo of cartesianProduct(string, value.length-1)){
            const result = testOpComb(combo, value)
            if (result === key){
                sum += result
                break
            }
        }
    }

    return sum
}

(() => {
    console.time('Time part1')
    console.log('Part 1:', totalCalibration('*+'))
    console.timeEnd('Time part1')

    console.time('Time part2')
    console.log('Part2:', totalCalibration('*+|'))
    console.timeEnd('Time part2')
})()