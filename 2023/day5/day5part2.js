const fs = require('fs') ;

const input = fs.readFileSync('input.txt', 'utf-8').split('\n').filter(lines => lines !== '') ;

const seedValues = input.shift().split('seeds:')[1].trim().replace('  ', ' ').split(' ').map(n => +n) ;

let blocks = [], seeds = [] ;

function parseFile(){
    let index = -1 ;
    for(let i=0; i< input.length; i++){
        const char = input[i][0] ;
        if (isNaN(char)){
            ++index ;
        }
        else if (!isNaN(char)){
            if (!blocks[index]){
                blocks[index] = [input[i].split(' ').map(n => +n)] ;
            }
            else{
                blocks[index].push(input[i].split(' ').map(n => +n)) ;
            }
        }
    }

    for(let i=0; i<seedValues.length; i+=2){
        seeds.push([seedValues[i], seedValues[i]+seedValues[i+1]]) ;
    }
}

function day5part2(){
    let loc = Infinity ;

    // console.log(seeds,blocks) ;

    for(let block of blocks){
        // console.log(ranges)
        const newRange = [] ;
        while(seeds.length > 0){
            const [st,end] = seeds.pop() ;
            console.log(st,end)
            for(const [d, s, l] of block){
                let os = Math.max(st, s) ;
                let oe = Math.min(end, d+l) ;

                if (os < oe){
                    newRange.push([os-s+d, oe-s+d]) ;

                    if (os >  s){
                        seeds.push([s, os]) ;
                    }
                    if (end > oe){
                        seeds.push([oe, end]) ;
                    }
                    break ;
                }
                else{
                }
            }
            newRange.push([st,end]) ;
        }
        seeds = newRange;
    }

    console.log(seeds, seeds.length)

    return loc ;
}

parseFile() ;
console.time('day5part2') ;
console.log(day5part2()) ;
console.timeEnd('day5part2') ;