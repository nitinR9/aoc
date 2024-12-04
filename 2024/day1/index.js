const fs = require('fs')

const file = fs.readFileSync('./input.txt', 'utf-8')

function parseFile1(value){
    const input = file.split('\n').map(s => s.split('   ').map(s => +s)).reduce((obj, val, index) => {
        obj.left.push(val[0])
        obj.right.push(val[1])
        return obj
    }, { left: [], right: [] })

    input.left.sort((a,b) => a-b)
    input.right.sort((a,b) => a-b)

    return input
}

function parseFile2(file){
    const map = new Map()
    const input = file.split('\n').map(s => s.split('   ').map(s => +s)).reduce((obj, val) => {
        obj.left.push(val[0])
        let num = map.get(val[1])
        if (num){
            map.set(val[1], num+1)
        }
        else{
            map.set(val[1], 1)
        }
        return obj
    }, { left: [] })

    return {
        ...input,
        right: map
    }
}

function part1(){
    const { left, right } = parseFile1(file)
    let sum = 0
    for(let i=0; i<left.length; i++){
        sum += Math.abs( left[i] - right[i] )
    }

    return sum
}

function part2(){
    const { left, right } = parseFile2(file)
    let sum = 0
    for(let i=0; i<left.length; i++){
        const val = left[i]
        const multiplier = right.get(val)
        sum += val * (multiplier === undefined ? 0 : multiplier)
    }

    return sum
}

(() => {
    console.time('Time for part 1')
    console.log('Part 1:', part1())
    console.timeEnd('Time for part 1')
    console.time('Time for part 2')
    console.log('Part 2:', part2())
    console.timeEnd('Time for part 2')
})()
