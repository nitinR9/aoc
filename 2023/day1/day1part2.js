const fs = require('fs') ;

const input = fs.readFileSync('biginput.txt', 'utf-8') ;

const lines = input.split('\n').map(l => l.split('\r')[0]) ;

const map = {
    one: 1,
    two: 2,
    three: 3,
    four: 4,
    five: 5,
    six: 6,
    seven: 7,
    eight: 8,
    nine: 9
} ;

const keys = [...new Set(Object.keys(map).map(k => k[0]))] ;

const res = lines.map(line => {
    let index = 0, result = [null, null] ;
    
    while(index < line.length ){
        const char = line[index] ;
        if (!isNaN(+char)){
            result[result[0] ? 1: 0] = +char ;
            index++ ;
        }
        else if (keys.indexOf(char) !== -1){
            const m3 = map[line.substring(index, index+3)] ;
            const m4 = map[line.substring(index, index+4)] ;
            const m5 = map[line.substring(index, index+5)] ;

            if (m3){
                result[result[0] ? 1: 0] = m3 ;
                index+=2;
            }
            else if (m4){
                result[result[0] ? 1: 0] = m4 ;
                index+=3 ;
            }
            else if (m5){
                result[result[0] ? 1: 0] = m5 ;
                index+=4 ;
            }
            else{
                index++ ;
            }
        }
        else{
            index++ ;
        }
    }

    if (!result[1]){
        result[1] = result[0] ;
    }

    return +result.join('')
})

console.log(res.reduce((acc, value) => acc + value, 0))