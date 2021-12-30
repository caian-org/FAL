package main

import "C"

// export SumAndMultiplies
func SumAndMultiplies(value int) int {
    return (value + 1) * value
}

func main() {}
