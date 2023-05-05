package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strings"
)

type Bigramms struct {
	firstLetters map[string]int
	lastLetters  map[string]int
	midLetters   map[string]int
}

type Probability struct {
	firstLetter map[string]float64
	letters     map[string]float64
	lastLetter  map[string]float64
}

const allNames = 32033

func main() {
	names, err := readTxt()
	if err != nil {
		log.Println(err)
		return
	}

	bigramms := makeThreeMaps(names)

	prblty := probability(bigramms)

	name := newName(prblty)
	fmt.Println("New name:", name)
}

func probability(bgrams Bigramms) Probability {
	prbs := Probability{}
	prbs.firstLetter = make(map[string]float64)
	prbs.lastLetter = make(map[string]float64)
	prbs.letters = make(map[string]float64)

	for k, v := range bgrams.firstLetters {
		val := float64(v) / allNames
		prbs.firstLetter[k] = math.Round(val*10000) / 10000
	}

	for k, v := range bgrams.lastLetters {
		val := float64(v) / allNames
		prbs.lastLetter[k] = math.Round(val*10000) / 10000
	}

	for k, v := range bgrams.midLetters {
		val := float64(v) / allNames
		prbs.letters[k] = math.Round(val*10000) / 10000
	}

	return prbs
}

func makeThreeMaps(names []string) Bigramms {
	bgrms := Bigramms{}
	bgrms.firstLetters = make(map[string]int)
	bgrms.lastLetters = make(map[string]int)
	bgrms.midLetters = make(map[string]int)
	var le2ers string

	for _, name := range names {
		bgrms.firstLetters[string(name[0])]++
		bgrms.lastLetters[string(name[len(name)-1])]++
		for j := 0; j < len(name)-1; j++ {
			le2ers = string(name[j]) + string(name[j+1])
			bgrms.midLetters[le2ers]++
		}
	}
	return bgrms
}

func newName(prblty Probability) string {
	first := randomFirstLetter()
	res := first
	max := 0.0
	var maxComb string

	fmt.Println("first letter:", first)

	for len(res) < 10 {
		fmt.Println("letter looking for the most likely pair:", first)
		for k, v := range prblty.letters {
			if first == string(k[0]) {
				if v > max {
					max = v
					maxComb = k
				}
			}
		}

		fmt.Printf("the most suitable couple: '%s', with probability: %f\n", maxComb, max)

		if len(res) > 3 {
			if len(res) == 4 && res[len(res)-4:len(res)-2] == res[len(res)-2:] {
				fmt.Printf("here we hit a couple: '%s' that will be repeated at the end of the name, so we leave what happened\n", maxComb)
				res = res[:len(res)-1]
				break
			}

			if res[len(res)-3:len(res)-1] == maxComb {
				fmt.Printf("here we hit a couple: '%s' that will be repeated at the end of the name, so we leave what happened\n", maxComb)
				break
			}
		}

		first = string(maxComb[len(maxComb)-1])
		res += string(maxComb[len(maxComb)-1])
		max = 0

		fmt.Println("generation name:", res)

	}

	return res
}

func readTxt() ([]string, error) {
	names, err := os.ReadFile("names.txt")
	if err != nil {
		return nil, err
	}
	arrnames := strings.Split(string(names), "\n")
	return arrnames, nil
}

func randomFirstLetter() string {
	letters := []rune("abcdefghijklmnopqrstuvwxyz")
	b := make([]rune, 1)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
