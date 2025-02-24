package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func eprint(a ...interface{}) {
	fmt.Fprintln(os.Stderr, a...)
}

func print(a ...interface{}) {
	fmt.Println(a...)
}

func Perm(a []string, f func([]string)) {
	perm(a, f, 0)
}

// Permute the values at index i to len(a)-1.
func perm(a []string, f func([]string), i int) {
	if i > len(a) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}

func fight(a, b string) (bool, bool) {
	iWin := make(map[string]bool)
	iWin["feueau"] = false
	iWin["eaufeu"] = true
	iWin["feuplante"] = true
	iWin["plantefeu"] = false
	iWin["feuglace"] = true
	iWin["glacefeu"] = false
	iWin["eauplante"] = false
	iWin["planteeau"] = true
	iWin["eausol"] = false
	iWin["soleau"] = true
	iWin["plantepoison"] = true
	iWin["poisonplante"] = false
	iWin["plantesol"] = false
	iWin["solplante"] = true
	iWin["plantevol"] = true
	iWin["volplante"] = false

	battle, found := iWin[fmt.Sprintf("%s%s", a, b)]
	return battle, found
}

func battleRound(sasha, my *list.List) {
	sashaCard := sasha.Front()
	myCard := my.Front()
	battle, found := fight(fmt.Sprintf("%s", myCard.Value), fmt.Sprintf("%s", sashaCard.Value))

	if found == false || myCard.Value == sashaCard.Value {
		sasha.Remove(sashaCard)
		my.Remove(myCard)
	} else if battle == true {
		sasha.Remove(sashaCard)
	} else {
		my.Remove(myCard)
	}
}

const workers = 100

func wDelete(jobs <-chan []string, results chan bool, winDeck chan []string, sasha []string) {
	for deck := range jobs {
		myDeck := list.New()
		// for each permutation, create a list for my deck
		for d := range deck {
			myDeck.PushBack(deck[d])
		}
		sashaDeck := list.New()
		for i := range sasha {
			sashaDeck.PushBack(sasha[i])
		}

		// and battleRound decks
		for sashaDeck.Len() > 0 && myDeck.Len() > 0 {
			battleRound(sashaDeck, myDeck)
		}
		// then verify if my deck is winning
		if myDeck.Front() != nil {
			results <- true
			winDeck <- deck
			break
		}
		results <- false
		winDeck <- []string{}
	}
}

func contestResponse() {
	eprint("=============== BEGIN INPUT ===============")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // first line is maybe not needed
	n, _ := strconv.Atoi(scanner.Text())
	eprint(n)

	// read inputs
	scanner.Scan()
	sasha := strings.Split(scanner.Text(), " ")
	scanner.Scan()
	my := strings.Split(scanner.Text(), " ")

	eprint("=============== END INPUT ===============")
	eprint("")

	// initialize vars
	var myDecks [][]string

	// generate all permutations for my deck
	Perm(my, func(a []string) {
		myDecks = append(myDecks, a)
	})

	numJobs := len(myDecks)
	eprint("number of jobs :", numJobs)
	jobs := make(chan []string, numJobs)
	winDeck := make(chan []string, numJobs)
	results := make(chan bool, numJobs)

	for w := 0; w < numJobs; w++ {
		go wDelete(jobs, results, winDeck, sasha)
	}

	for x := range myDecks {
		jobs <- myDecks[x]
	}
	close(jobs)
	for a := 0; a < numJobs; a++ {
		foo := <- results
		deck := <- winDeck
		if foo {
			print(strings.Join(deck, " "))
			os.Exit(0)
		}
	}
	print("-1")
}

func main() {
	contestResponse()
}
