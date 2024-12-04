const fs = require('fs')

function parser(input, checkOnlyMul = true){
    const ins = []
    let disable = false, index = 0

    while(index < input.length){
        let start = 0
        const slice = input.slice(index, index+3)

        if (!checkOnlyMul && slice === 'don'){
            index += 8;
            disable = true
            continue
        }
        if (!checkOnlyMul && slice === 'do('){
            index += 5;
            disable = false
            continue
        }
        if (checkOnlyMul){
            index
        }
        
    }

}

function run(){
    let ans1 = 0, ans2 = 0
    let file = fs.readFileSync('./input.txt', 'utf-8')
    const insRegexp = new RegExp(/mul\(([\d]{1,3}),([\d]{1,3})\)/g)
    const instructions1 = file.matchAll(insRegexp)

    for(let ins of instructions1){
        const [x, y] = [+ins[1], +ins[2]]
        ans1 += x * y
    }

    const instructions2 = file
                            .split("do()")
                            .map(line => line.split("don't()")[0])
                            .map(input => [...input.matchAll(insRegexp)])
                            .flat()
                            .map(m => m[1] * m[2])
    
    for(let ins of instructions2){
        ans2 += ins
    }
    return [ans1, ans2]
}

(() => {
    console.time('Time taken for both')
    const [ans1, ans2] = run()
    console.log('Part 1:', ans1)
    console.log('Part 2:', ans2)
    console.timeEnd('Time taken for both')
})()