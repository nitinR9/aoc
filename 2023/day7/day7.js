const fs = require('fs') ;

const cardMap1 = {
    A: 13,
    K: 12,
    Q: 11,
    J: 10,
    T: 9,
    9: 8,
    8: 7,
    7: 6,
    6: 5,
    5: 4,
    4: 3,
    3: 2,
    2: 1
}, cardMap2 = {
    A: 13,
    K: 12,
    Q: 11,
    J: 1,
    T: 10,
    9: 9,
    8: 8,
    7: 7,
    6: 6,
    5: 5,
    4: 4,
    3: 3,
    2: 2
} ;

function getP1input(){
    return fs.readFileSync('big.txt', 'utf-8').split('\n').map(l => l.split(' ')).map(([h, bid]) => [h, +bid])
}

function getHandType(hand, enablePart2=false){
    let map = {}, carry = 0 ;
    
    for(let char of hand){
        if (enablePart2 && char === 'J'){
            carry++ ;
        }
        else{
            map[char] = 1 + (map[char] || 0) ;
        }
    }
        
    if (carry === 5){
        map['J'] = 5 ;
    }
    else{
        const key = Object.keys(map).reduce((a,b) => map[a] > map[b] ? a: b) ;
        map[key] += carry ;
    }
    
    const keys = Object.keys(map) ;
    const values = Object.values(map) ;
    
    
    // five of a kind
    if (keys.length === 1){
        return 6 ;
    }
    // four of a kind
    else if (keys.length === 2 && values.includes(4)){
        return 5 ;
    }
    // full house
    else if (keys.length === 2){
        return 4 ;
    }
    // three of a kind
    else if (keys.length === 3 && values.includes(3)){
        return 3 ;
    }
    // two pair
    else if (keys.length === 3){
        return 2 ;
    }
    // one pair
    else if (keys.length === 4){
        return 1 ;
    }
    // high card
    else{
        return 0 ;
    }
}

function part1(part2 = false){
    const values = Object.fromEntries(getP1input()) ;
    let handPairs = new Array(7) ;
    
    for(let hand of Object.keys(values)){
        const type = getHandType(hand, part2) ;
        
        if (handPairs[type]){
            handPairs[type].push(hand) ;
        }
        else{
            handPairs[type] = [hand] ;
        }
    }
    
    const finalOrder = handPairs.map(pair => {
        if (pair && pair.length > 1){
            return pair.sort((a,b) => {
                for(let i=0; i<a.length; i++){
                    if (part2){
                        if (cardMap2[b[i]] > cardMap2[a[i]]){
                            return -1 ;
                        }
                        else if (cardMap2[b[i]] < cardMap2[a[i]]){
                            return 1 ;
                        }
                    }
                    else{
                        if (cardMap1[b[i]] > cardMap1[a[i]]){
                            return -1 ;
                        }
                        else if (cardMap1[b[i]] < cardMap1[a[i]]){
                            return 1 ;
                        }
                    }
                }
            }) ;
        }
        
        return pair ;
    }) ;
    
    return finalOrder.flat().reduce((acc, x, i) => acc+(values[x] * (i+1)), 0) ;
}

console.time('day7 part1') ;
console.log(part1()) ;
console.timeEnd('day7 part1') ;

console.time('day7 part2') ;
console.log(part1(true)) ;
console.timeEnd('day7 part2') ;