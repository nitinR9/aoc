const fs = require('fs') ;

const map = fs.readFileSync('input.txt', 'utf-8').split('\n') ;

const mapWeight = {
    '|': [-1,0],
    '-': [0,-1],
    'L': [-1,1],
    'J': [-1,-1],
    '7': [1,-1],
    'F': [1, 1]
} ;

// top, bottom, left, right
const validMoves = [['|', 'F', '7'], ['|', 'J', 'L'], ['F', 'L', '-'], ['7', '-', 'J']] ;

/*

    .   T   .
    L  [c]  R
    .   B   .
*/

function isValidTop(x,y){
    return y > 0 &&  validMoves[0].includes(map[y-1][x]) ;
}

function isValidBottom(x,y){
    return y < map.length-1 &&  validMoves[1].includes(map[y+1][x]) ;
}

function isValidLeft(x,y){
    return x > 0 && validMoves[2].includes(map[y][x-1]) ;
}

function isValidRight(x,y){
    return x < map[0].length-1 && validMoves[3].includes(map[y][x+1]) ;
}

function isEqual(p1, p2){
    return p1[0] === p2[0] && p1[1] === p2[1] ;
}

function getPoints(currentPos){
    const [y,x] = currentPos ;
    const moves = [] ;
    if (isValidTop(x,y)){
        moves.push([y-1, x]) ;
    }
    if (isValidBottom(x,y)){
        moves.push([y+1, x]) ;
    }
    if (isValidLeft(x,y)){
        moves.push([y, x-1]) ;
    }
    if (isValidRight(x,y)){
        moves.push([y, x+1]) ;
    }
    return moves ;
}

function getNextMove(current, prev){
    const moves = getPoints(current) ;
    return moves.filter(move => !isEqual(move, prev))[0] ;
}



function part1(){
    let pos = null, [prevp1, prevp2] = [null, null] , [p1,p2] = [null, null] ;

    for(let i=0; i<map.length; i++){
        for(let j=0; j<map[0].length; j++){
            if (map[i][j] === 'S'){
                pos = [i,j] ;
                prevp1 = pos ;
                prevp2 = pos ;
                [p1,p2] = getPoints(pos) ;
                console.log('initial next moves', p1, p2)
            }
        }
    }

    let count1 = 1, count2 = 1 ;
    while(!isEqual(p1, p2)){
        console.log('looping') ;
        const temp1 = getNextMove(p1, prevp1) ;
        const temp2 = getNextMove(p2, prevp2) ;
        prevp1 = p1 ;
        prevp2 = p2 ;
        [p1, p2] = [temp1, temp2] ;
        console.log('moves', prevp1, '==>', temp1, '&', prevp2, '==>',  prevp2)
        count1++ ;
        count2++ ;
    }
}


console.time('Time of part 1') ;
part1()
console.timeEnd('Time of part 1') ;