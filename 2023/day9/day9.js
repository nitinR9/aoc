const fs = require('fs') ;

function parseFile(){
    const input = fs.readFileSync('big.txt', 'utf-8').split('\n') ;
    return input.map(line => line.split(' ').map(n => +n)) ;
}

function part1(){
    const histories = parseFile() ;
    let sum = 0 ;

    const calcPredictions = (history) => {
        const sequences = [history] ;
        while(sequences.at(-1).some(v => v !== 0)){
            const next = [], sq = sequences.at(-1) ;
            for(let i=1, k=0; i<sq.length; i++, k++){
                next[k] = sq[i] - sq[i-1] ;
            }
            sequences.push(next) ;
        }

        return sequences.map(s => s.at(-1)).reduce((acc, c) => acc+c, 0) ;
    } ;

    for(const history of histories){
        sum += calcPredictions(history) ;
    }

    return sum ;
}

function part2(){
    const histories = parseFile() ;
    let sum = 0 ;

    const calcPredictions = (history) => {
        const sequences = [history] ;
        while(sequences.at(-1).some(v => v !== 0)){
            const next = [], sq = sequences.at(-1) ;
            for(let i=1, k=0; i<sq.length; i++, k++){
                next[k] = sq[i] - sq[i-1] ;
            }
            sequences.push(next) ;
        }

        return sequences.reverse().map(s => s.at()).reduce((acc,c) => c-acc, 0) ;
    }

    for(const history of histories){
        sum += calcPredictions(history) ;
    }

    return sum ;
}

console.time('Time for part 1') ;
console.log(part1()) ;
console.timeEnd('Time for part 1') ;

console.time('Time for part 1') ;
console.log(part2()) ;
console.timeEnd('Time for part 1') ;