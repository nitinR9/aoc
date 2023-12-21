const fs = require('fs') ;

const input = fs.readFileSync('big2.txt', 'utf-8') ;

const games = input.split('\n').map(game => {
    return game.split(':')[1].trim().split(';').map(sub => {
        return sub.split(',').map(cubes => cubes.trim()).map(cube => {
            const [num, type] = cube.split(' ') ;
            return num + type[0].toUpperCase() ;
        }) ;
    }) ;
}) ;

let result = games.map(game => {
    const finalSet = new Array(3).fill(0) ;
    game.flat().forEach(set => {
        const [R, G, B] = finalSet ;
        const type = set.at(-1) ;
        const num = +set.slice(0, set.length-1) ;
        if (type === 'R'){
            finalSet[0] = R < num ? num : R ;
        }
        else if (type === 'G'){
            finalSet[1] = G < num ? num : G ;
        }
        else if (type === 'B'){
            finalSet[2] = B < num ? num : B ;
        }
    }) ;

    return finalSet ;
}).reduce((acc, [R,G,B]) => acc + (R*G*B), 0);

console.log(result)