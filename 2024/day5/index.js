const fs = require('fs')
const file = fs.readFileSync('input.txt', 'utf-8')
const input = file.split('\n\n').map(v => v.split('\n'))

const rules = input[0].map(v => v.split('|')).reduce((prev, value) => {
    const [x,y] = value
    if (prev[x]){
        prev[x].push(+y)
    }
    else{
        prev[x] = [+y]
    }
    return prev
}, {})

const updates = input[1].map(u => u.split(',').map(Number))

function isValidPage(num, leftArr){
    const checkValues = rules[num]
    if (!checkValues)
        return true
    return !leftArr.some(v => checkValues.includes(v))
}

function getSwitchIndex(num, leftArr){
    const checkValues = rules[num]

    if (!checkValues)
        return [false, null]
    
    const index = leftArr.findIndex((v) => checkValues.includes(v))

    return index === -1 ? [false, null] : [true, index]
}

function switchAndUpdate(update){

    // initial check if recursive update param is valid, if yes then return it otherwise switch the invalid positions
    let valid = true
    for(let i=update.length-1; i>0; i--){
        if (!isValidPage(update[i], update.slice(0, i))){
            valid = false
            break
        }
    }
    if (valid)
        return update

    // get the switch index and switch it with current position and recursively return with new update
    for(let i=update.length-1; i>0 ; i--){
        const [switchRequired, pos] = getSwitchIndex(update[i], update.slice(0, i))
        if (switchRequired){
            let temp = update[i]
            update[i] = update[pos]
            update[pos] = temp
        }
    }
    return switchAndUpdate(update)
}

function run(){
    let value1 = 0, value2 = 0, index = 0, invalidUpdateIndexes = []

    // part1
    for(let update of updates){
        let valid = true
        for(let i=update.length-1; i>0 ; i--){
            if (!isValidPage(update[i], update.slice(0, i))){
                valid = false
                invalidUpdateIndexes.push(index)
                break
            }
        }
        if (valid){
            value1 += update[Math.floor(update.length/2)]
        }
        index++
    }

    // part2
    for(let index of invalidUpdateIndexes){
        let update = updates[index]
        update = switchAndUpdate(update)
        value2 += update[Math.floor(update.length/2)]
        iter = 0
    }

    return [value1, value2]
}

(() => {
    console.time('Time taken for both')
    const [r1, r2] = run()
    console.log('Part 1:', r1)
    console.log('Part 2:', r2)
    console.timeEnd('Time taken for both')
})()