package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	var inventory = make(map[int]int)
	f, err := os.Open("../../inputs/day01/day01.txt")
	if err != nil {
		fmt.Printf("opening file %s", err)
		return
	}
	defer f.Close()
	buf := bufio.NewScanner(f)
	buf.Split(bufio.ScanLines)

	elfNum := 0
	for buf.Scan() {
		s := buf.Text()
		if s != "" {
			val, err := strconv.Atoi(s)
			if err != nil {
				fmt.Println("converting amount", err)
			}
			inventory[elfNum] += val
		} else {
			elfNum += 1
		}

	}
	temp := inventory
	first, firstScore := getWinner(temp)
	temp[first] = 0
	second, secondScore := getWinner(temp)
	temp[second] = 0
	third, thirdScore := getWinner(temp)
	temp[third] = 0

	fmt.Println(inventory[first])
	fmt.Printf("%d is carrying the most calories with %d \n", first, firstScore)
	fmt.Printf("Top three elves are carrying %d calories\n", firstScore+secondScore+thirdScore)

}

func getWinner(m map[int]int) (k, v int) {
	max := 0
	winner := 0
	for k, v := range m {
		if v > max {
			max = v
			winner = k
		}
	}
	return winner, max
}
