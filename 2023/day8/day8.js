const fs = require('fs') ;

function parseFile(){
    let [pattern, ...values] = fs.readFileSync('big.txt', 'utf-8').split('\n').filter(v => v) ;
    const instructions = values.map(i => i.split(' = ')).map(i => [i[0], i[1].match(/([A-Z]{3})/gm)]).reduce((obj, [key, value]) => {
        obj[key] = value ;
        return obj
    }, {}) ;
    return [pattern, instructions] ;
}

function part1(){
    const [pattern, instructions] = parseFile() ;
    let pIndex = 0, count = 0, inst = 'AAA'  ;

    while(inst !== 'ZZZ'){
        const [left, right] = instructions[inst] ;
        inst = pattern[pIndex] === 'L' ? left : right ;
        if (pIndex !== pattern.length-1){
            pIndex++
        }
        else{
            pIndex = 0 ;
        }
        count++ ;
    }

    return count ;
}

function GCD(a,b){
    let min = Math.min(a,b) ;
    while(min > 0){
        if (a % min === 0 && b % min === 0){
            break ;
        }
        min-- ;
    }

    return min ;
}

function part2(){
    const [pattern, instructions] = parseFile() ;
    const RegexZ = new RegExp(/..Z/) ;
    const startNodes = Object.keys(instructions).filter(node => node.match(/..A/)?.length) ;
    let counts = [] ;

    for(const start of startNodes){
        let pIndex = 0, current = start ;
        let count = 0 ;

        while(!RegexZ.test(current)){
            const [left, right] = instructions[current] ;
            current = pattern[pIndex] === 'L' ? left : right ;
            if (pIndex !== pattern.length - 1){
                pIndex++ ;
            }
            else{
                pIndex = 0 ;
            }

            count++ ;
        }

        counts.push(count) ;
    }
    let lcm = counts.shift() ;

    counts.forEach((count) => {
        const gcd = GCD(lcm, count) ;
        lcm = (lcm * count) / gcd ;
    }) ;

    return lcm ;    
}

console.time('Time for part1') ;
console.log(part1()) ;
console.timeEnd('Time for part1') ;

console.time('Time for part2') ;
console.log(part2()) ;
console.timeEnd('Time for part2') ;