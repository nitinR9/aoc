const fs = require('fs') ;

const input = fs.readFileSync('big.txt', 'utf-8').split('\r\n') ;

const cards = input.map(l => l.split(':')[1].trim().split('|').map(nums => nums.trim().replace(/\ \ /g, ' ')).map(nums => nums.split(' '))) ;

function getPoints(wins, nums){
    let points = 0 ;
    for(let i=0; i<wins.length; i++){
        if (nums.includes(wins[i])){
            if (points==0){
                points = 1 ;
            }
            else{
                points *= 2 ;
            }
        }
    }
    if (!points){
        return null ;
    }
    return points ;
}

function day4part1(){
    let result = [] ;

    for(let [wins, nums] of cards){
            const points = getPoints(wins, nums) ;
            if (points){
                result.push(points) ;
            }
    }

    return result.reduce((acc, num) => acc+num, 0) ;
}

console.log(day4part1()) ;

