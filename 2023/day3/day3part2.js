const fs = require('fs') ;

const input = fs.readFileSync('big.txt', 'utf-8').split('\n').map(l => l.split('\r')[0]) ;

const MAXROWS = input.length, MAXCOLS = input[0].length ;

function getNum(row, col){
    let leftPos = col, rightPos = col, line = input[row] ;

    while(leftPos >= 0 && rightPos < MAXCOLS){
        const [left, right] = [line[leftPos-1], line[rightPos+1]] ;

        if (isNaN(left) && isNaN(right)){
            break ;
        }
        if (!isNaN(left)){
            leftPos-- ;
        }
        if (!isNaN(right)){
            rightPos++
        }
    }

    return +line.slice(leftPos, rightPos+1) ;
}

function findGearRatio(row, col){
    const result = [], directions = [
        [-1,-1], [-1,0], [-1,1],
        [0,-1], [0, 1],
        [1,-1], [1,0], [1,1]
    ] ;

    for(let [y,x] of directions){
        const [rPos, cPos] = [row+y, col+x] ;
        if (!isNaN(input[rPos][cPos])){
            const num = getNum(rPos, cPos) ;
            if (num && !result.includes(num)){
                result.push(num) ;
            }
        }
    }
    if (result.length !== 2){
        return 0 ;
    }
    
    return result.reduce((acc,num) => acc*num, 1) ;
}

function day3part2(){
    let result = [] ;

    for(let row=0; row<MAXROWS; row++){
        for(let col=0; col<MAXCOLS; col++){
            const char = input[row][col] ;
            if (isNaN(char) && char !== '.'){
                const ratio = findGearRatio(row, col) ;
                if (ratio){
                    result.push(ratio) ;
                }
            }
        }
    }
    return result.reduce((acc, num) => acc+num, 0) ;
}

console.log(day3part2())