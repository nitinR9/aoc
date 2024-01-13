const fs = require('fs') ;

const input = fs.readFileSync('input.txt', 'utf-8').split('\n').filter(lines => lines !== '') ;

const seeds = input[0].split('seeds:')[1].trim().replace('  ', ' ').split(' ').map(n => +n) ;
const almanac = []

function parseFile(lines){
    let index = -1 ;
    for(let i=1; i< lines.length; i++){
        const char = lines[i][0] ;
        if (isNaN(char)){
            ++index ;
        }
        else if (!isNaN(char)){
            if (!almanac[index]){
                almanac[index] = [lines[i].split(' ').map(n => +n)] ;
            }
            else{
                almanac[index].push(lines[i].split(' ').map(n => +n)) ;
            }
        }
    }
}

function evaluateRange(value, ranges){
    for(let i=0; i<ranges.length; i++){
        const [dest, source, length] = ranges[i] ;
        if (source <= value && value < source+length){
            if (source > dest){
                value -= (source - dest) ;
                break ;
            }
            else{
                value += (dest - source) ;
                break ;
            }
        }
    }

    return value ;
}

function getLocation(seed){
    let value = seed ;
    for(let i=0; i<almanac.length; i++){
        value = evaluateRange(value, almanac[i]) ;
    }

    return value ;
}

function day5part1(){
    let locations = [] ;
    for(let i=0; i<seeds.length; i++){
        locations.push(getLocation(seeds[i]));
    }
    
    return Math.min(...locations) ;
}

parseFile(input) ;
console.time('time') ;
console.log(day5part1())
console.timeEnd('time') ;