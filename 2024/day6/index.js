const fs = require('fs')
const file = fs.readFileSync(__dirname + '/input.txt', 'utf-8')
const input = file.split('\n').map(s => s.split(''))

const MAX_ROWS = input.length
const MAX_COLS = input[0].length

const dirMap = {
    '^': [-1,0],
    'v': [1, 0],
    '>': [0,1],
    '<': [0,-1]
}

const moveMap = {
    '^': '>',
    'v': '<',
    '>': 'v',
    '<': '^'
}

function getStartPos(){
    let dir = null
    for(let row=0; row < MAX_ROWS; row++){
        for(let col=0; col < MAX_COLS; col++ ){
            const val = input[row][col]
            switch(val){
                case '^': dir = '^'
                break;
                case 'v': dir = 'v'
                break;
                case '<': dir = '<'
                break;
                case '>': dir = '>'
                break;
            }
            if (dir !== null){
                return {
                    pos: [row, col],
                    dir
                }
            }
        }
    }
}

function nextPos(pos, dir){
    let next = [...pos]
    const [R, C] = dirMap[dir]
    next[0] += R
    next[1] += C
    return next
}

function isOutOfBounds(row, col){
    return !((0 <= row && row < MAX_ROWS) && ( 0 <= col && col < MAX_COLS))
}

function travel(pos, direction){
    let [row, col] = pos, dir = direction , uniquePos = new Set(), map = input

    while(true){
        uniquePos.add(`${row},${col}`)

        const [r1, c1] = nextPos([row, col], dir)

        if (isOutOfBounds(r1, c1))
            break

        if (['.', 'X'].includes(map[r1][c1])){
            [row, col] = [r1, c1]
            map[row][col] = 'X'
        }
        else if(map[r1][c1] === '#'){
            dir = moveMap[dir]
        }
    }

    return uniquePos.size
}

function isLooping(arr, pos, direction){
    let [row, col] = pos, dir = direction, map = arr, turns = 0
    
    while(true){
        const [r1, c1] = nextPos([row, col], dir)
        turns++

        if (turns === MAX_COLS * MAX_ROWS + 1){
            return true
        }

        if (isOutOfBounds(r1, c1)){
            return false
        }

        if (map[r1][c1] !== '#'){
            [row, col] = [r1, c1]
        }
        else{
            dir = moveMap[dir]
        }
    }
}

function hasVisitedOnPath(row, col){
  return Object.values(dirMap).map(([x,y]) => [row+x, col+y]).filter(([r, c]) => !isOutOfBounds(r, c)).some(([x,y]) => input[x][y] === 'X')
}

function getAllObstructions(pos, dir){
    let numPos = 0, map = input.map(arr => arr.slice())

    for(let i=0; i< MAX_ROWS; i++){
        for(let j=0; j< MAX_COLS; j++){
            if (hasVisitedOnPath(i, j)){
                const temp = map[i][j]
                map[i][j] = '#'
                if (isLooping(map, pos, dir)){
                    numPos++
                }
                map[i][j] = temp
            }
        }
    }
    
    return numPos
}

(() => {
    let { pos, dir } = getStartPos()
    let [x,y] = pos
    input[x][y] = '.'
    console.time('Time part1')
    console.log('Part 1:', travel(pos, dir))
    console.timeEnd('Time part1')
    console.time('Time part2')
    console.log('Part2:', getAllObstructions(pos, dir))
    console.timeEnd('Time part2')    
})()