const fs = require('fs') ;

const input = fs.readFileSync('big.txt', 'utf-8') ;

const map = {
    G: 13,
    R: 12,
    B: 14
}

const games = input.split('\r\n').map(game => {
    return game.split(':')[1].trim().split(';').map(sub => {
        return sub.split(',').map(cubes => cubes.trim()).map(cube => {
            const [num, type] = cube.split(' ') ;
            return num + type[0].toUpperCase() ;
        }) ;
    }) ;
}) ;

let result = []

games.forEach((game, index) => {
    let bool = false ;
    outer: for(let sets of game){
        bool = sets.some(set => {
            const num = set.slice(0, set.length-1) ;
            const type = set[set.length-1] ;
            return map[type] < +num
        }) ;

        if (bool){
            break outer;
        }
    }
    if (!bool){
        result.push(index+1) ;
    }
}) ;

console.log(result.reduce((acc, num) => acc+num, 0)) ;