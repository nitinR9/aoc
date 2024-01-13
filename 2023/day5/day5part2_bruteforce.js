const fs = require('fs') ;

const input = fs.readFileSync('input.txt', 'utf-8').split('\n').filter(lines => lines !== '') ;

const initSeeds = input[0].split('seeds:')[1].trim().replace('  ', ' ').split(' ').map(n => +n) ;

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
    // console.log('evaluate range', value, ranges)
    for(let i=0; i<ranges.length; i++){
        const [dest, source, length] = ranges[i] ;
        if (source <= value && value < source+length){  
            const diff = dest - source ;
            value += diff ;
            break ;
        }
    }

    return value ;
}

function getLocation(start, length){
    let location = Infinity ;

    for(let i=start; i<start+length; i++){
        let value = i ;
        // console.log('seed start', i)
        for(let j=0; j<almanac.length; j++){
            value = evaluateRange(value, almanac[j]) ;
            // console.log('each step:', value)
        }
        // console.log('got location for seed:', i, value) ;
        location = Math.min(location, value) ; 
    }

    console.log('got min location for range', start, length, location)
    return location ;
}

function day5part2(){
    let location = Infinity ;
    
    for(let i=0; i<initSeeds.length; i+=2){
        console.log('for', initSeeds[i], initSeeds[i+1])
        location = Math.min(location, getLocation(initSeeds[i], initSeeds[i+1]));
    }
    
    return location ;
}

console.time('parse file') ;
parseFile(input) ;
console.timeEnd('parse file') ;
console.time('day5part2') ;
console.log(day5part2())
console.timeEnd('day5part2') ;

