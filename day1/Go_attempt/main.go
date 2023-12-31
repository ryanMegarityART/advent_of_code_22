package main

import (
"fmt"
"os"
"strings"
"strconv"
"path/filepath"
"runtime"
"errors"
)

func check (e error) {
    if e != nil {
        panic(e)
    }
}

func readInputfile(filePath string) string {
    data, err := os.ReadFile(filePath)
    check(err)
    return string(data)
}

// Filename is the __filename equivalent
func Filename() (string, error) {
    _, filename, _, ok := runtime.Caller(1)
    if !ok {
        return "", errors.New("unable to get the current filename")
    }
    return filename, nil
}

// Dirname is the __dirname equivalent
func Dirname() (string, error) {
    filename, err := Filename()
    if err != nil {
        return "", err
    }
    return filepath.Dir(filename), nil
}

func main() {
    // read input file into an array
    path, err := Dirname()
    check(err)
    textInput := readInputfile(path + "/input.txt")
    textInput = strings.TrimSpace(textInput)
    textInputSliced := strings.Split(textInput, "\n\n")
    //fmt.Print(textInputSliced)

    // loop over array, split into sub array and then sum the elements
    maxSum := 0
    top3Slice := [3]int{0, 0, 0}
    for _, element := range textInputSliced {
        subSlice := strings.Split(element, "\n")
        sum := 0
        for _, subElement := range subSlice {
            calories, err := strconv.ParseInt(subElement, 10, 0)
            check(err)
            sum += int(calories)
        }
    
        if sum > maxSum {
            maxSum = sum
        } 

        // keep top 3 slice ordered to make comparison straightforward
        if sum > top3Slice[0] {
            if sum > top3Slice[1] {
                if sum > top3Slice[2] {
                    top3Slice[0] = top3Slice[1]
                    top3Slice[1] = top3Slice[2]
                    top3Slice[2] = sum
                } else {
                    top3Slice[0] = top3Slice[1]
                    top3Slice[1] = sum
                }
            } else {
                top3Slice[0] = sum
            }
        }
    }
    fmt.Println("maxSum: ", maxSum)
    fmt.Println("top3Slice: ", top3Slice)
    top3Sum := 0
        for _, e := range top3Slice {
            top3Sum += int(e)
        }
    fmt.Println("top3Sum: ", top3Sum)

}

