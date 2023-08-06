import { promises as fs } from 'fs';

const INPUT_FILE_PATH = `${__dirname}/input.txt`

async function readFromFile(filePath: string) {
    const fileText = await fs.readFile(filePath)
    return fileText.toString()
}

function convertInputToArray(inputText: string) {
    const trimmedInput = inputText.trim()
    const splitOnBlankNewLine = trimmedInput.split(/\n\n/)
    const splitOnNewLine = splitOnBlankNewLine.map((a) => a.split(/\n/))
    return splitOnNewLine
}

async function main() {
    // read input file into a nested array
    const inputText = await readFromFile(INPUT_FILE_PATH)
    const inputArray = convertInputToArray(inputText)

    // calculate the sums of the nested arrays
    const summedArray: number[] = inputArray.map(a => a.reduce((prev: number, curr: string) => { return parseInt(curr) + prev }, 0))
    summedArray.sort((a, b) => b - a)
    const maxSingle = summedArray[0]
    const top3Sum = summedArray.slice(0, 3).reduce((a, b) => a + b, 0)
    return [maxSingle, top3Sum]
}

(async () => {
    const [maximumCalories, top3Sum] = await main()
    console.log("maximumCalories: ", maximumCalories)
    console.log("top3Sum: ", top3Sum)
})()

