const fs = require('fs')

const file = fs.readFileSync('./input.txt', 'utf-8')
const input = file.split('\n')
const MAX_ROWS = input.length, MAX_COLS = input[0].length

function hasXMAS(row, col, dirX, dirY){
    let result = ''
    while((row >= 0 && row < MAX_ROWS) && (col >= 0 && col < MAX_COLS) && result.length !== 4){
        result += input[row][col]
        row += dirX
        col += dirY
    }

    return result === 'XMAS'
}

function hasMAS(row, col){
    const positions = [input[row-1][col-1], input[row-1][col+1], input[row+1][col-1], input[row+1][col+1]]

    if (positions.some(v => !['M', 'S'].includes(v))){
        return false
    }

    const format = positions.toString()
    const validPositions = ['S,M,S,M', 'M,S,M,S', 'S,S,M,M', 'M,M,S,S']
    
    if (validPositions.includes(format)){
        return true
    }
    
    return false
}

function getTotal(i,j){
    const directions = [[0,1], [1,0], [0,-1], [-1,0], [1,1], [-1,-1], [1,-1], [-1,1]]

    return directions.reduce((prev, val) => {
        if (hasXMAS(i, j, val[0], val[1])){
            return prev+1
        }
        return prev
    }, 0)
}

function run(){
    let ans1 = 0, ans2 = 0

    for(let i=0; i<MAX_ROWS; i++){
        for(let j=0; j<MAX_COLS; j++){
            if (input[i][j] === 'X'){
                ans1 += getTotal(i,j)
            }
        }
    }

    for(let i=1; i<MAX_ROWS-1; i++){
        for(let j=1; j<MAX_COLS-1; j++){
            if (input[i][j] === 'A' && hasMAS(i, j)){
                ans2++
            }
        }
    }
    return [ans1, ans2]
}

(() => {
    console.time('Time taken for both')
    const [ans1, ans2] = run()
    console.log('Part 1:', ans1)
    console.log('Part 2:', ans2)
    console.timeEnd('Time taken for both')
})()