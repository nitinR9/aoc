const fs = require('fs') ;

const lines = fs.readFileSync('big1.txt', 'utf-8') ;

const sum = lines.split('\n').map( l => {
    let nums = [];
    for(let c of l){
        if (!isNaN(parseInt(c))){
            nums.push(c) ;
        }
    }
    if (nums.length < 2){
        return nums[0]+nums[0] ;
    }
    else{
        return nums.at(0) + nums.at(-1) ;
    }
}).reduce((acc, current) => acc + (+current), 0)

console.log(sum)