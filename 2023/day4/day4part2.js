const fs = require('fs') ;

const input = fs.readFileSync('big.txt', 'utf-8').split('\r\n') ;

const cards = input.map(card => card.split(':')[1].trim().split('|').map(nums => nums.trim().replace(/  /g, ' ')).map(nums => nums.split(' '))) ;

const mapMatch = new Array(cards.length).fill(0), result = new Array(cards.length).fill(1) ;

function calcCopies(index){
    if (!mapMatch[index]){
        return ;
    }

    const max = mapMatch[index]+index+1 > cards.length ? cards.length : mapMatch[index]+index ;

    for(let i=index+1; i<= max; i++){
        result[i] += 1 ;
    }
}

function day4part2(){
    for(let i=0; i<cards.length; i++){
        const [wins, nums] = cards[i] ;
        mapMatch[i] = wins.filter(win => nums.includes(win)).length ;
    }

    for(let i=0; i< cards.length; i++){
        for(let j=0; j< result[i]; j++){
            calcCopies(i) ;
        }
    }

    return result.reduce((acc, num) => acc+num, 0) ;
}

console.log(day4part2()) ;