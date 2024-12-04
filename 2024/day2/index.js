const fs = require('fs')

const file = fs.readFileSync('./input.txt', 'utf-8')

function parseFile1(value){
    return file.split('\n').map(s => s.split(' ').map(s => +s))
}

function checkReport(report){
    const orderedReport = [...report]
    orderedReport.sort((a,b) => a-b)
    const reversedReport = [...orderedReport]
    reversedReport.reverse()

    if (report.toString() !== orderedReport.toString() && report.toString() !== reversedReport.toString())
        return true

    const set = new Set(report)

    if (set.size !== report.length)
        return true

    for(let i=0; i< report.length-1; i++)
        if (Math.abs(report[i] - report[i+1]) > 3)
            return true

    return
}

function run(){
    const reports = parseFile1(file)
    let ansP1 = 0, ansP2 = 0
    for(let report of reports){
        let hasFailed = checkReport(report)

        if (!hasFailed){
            ansP1++
            ansP2++
            continue
        }

        hasFailed = false
        for(let i=0; i< report.length; i++){
            const reportCopy = [...report]
            reportCopy.splice(i, 1)

            hasFailed = checkReport(reportCopy)
            
            if (!hasFailed)
                break
        }

        if (!hasFailed)
            ansP2++
    }
    return [ansP1, ansP2]
}

(() => {
    console.time('Time taken for both')
    const [ans1, ans2] = run()
    console.log('Part 1:', ans1)
    console.log('Part 2:', ans2)
    console.timeEnd('Time taken for both')
})()