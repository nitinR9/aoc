const fs = require('fs') ;

const [time, distances] = fs.readFileSync('big.txt', 'utf-8').split('\n').map(v => v.split(':')[1].trim().split(' ').filter(v => v).map(n => +n)) ;

function getCount(maxTime, maxDistance){
    let count = 0 ;
    for(let t=1; t<maxTime; t++){
        const d = (maxTime-t) * t ;
        if (maxDistance < d){
            count++ ;
        }
    }

    return count ;
}

function part1(){
    const maxLen = time.length, result = [] ;

    for(let i=0; i<maxLen; i++){
        const maxT = time[i] ;
        const prevRecord = distances[i] ;
        result.push(getCount(maxT, prevRecord)) ;
    }

    return result.reduce((acc, x) => acc*x, 1) ;
}

function part2(){
    const t = time.map(n => n.toString()).join('') ;

    const d = distances.map(n => n.toString()).join('') ;

    return getCount(t, d) ;

}

console.time('day6 part1') ;
console.log(part1()) ;
console.timeEnd('day6 part1') ;

console.time('day6 part2') ;
console.log(part2()) ;
console.timeEnd('day6 part2') ;