package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

/*
 * The algorithm works by maintaining a reservoir of size k, which initially
 * contains the first k items of the input. It then iterates over the remaining
 * items until the input is exhausted. Using one-based array indexing, let i > k
 * be the index of the item currently under consideration.
 *
 * The algorithm then generates a random number j between (and including) 1 and i.
 * If j is at most k, then the item is selected and replaces whichever item
 * currently occupies the j-th position in the reservoir. Otherwise, the item is
 * discarded.
 */

func main() {

	if len(os.Args) < 2 {
		usage(os.Args[0])
	}

	rand.Seed(time.Now().UnixNano() | int64(os.Getpid()))

	k, err := strconv.Atoi(os.Args[1])
	if err != nil {
		usage(os.Args[0], err)
	}

	var reservoir []string

	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < k && scanner.Scan(); i++ {
		line := scanner.Text()
		reservoir = append(reservoir, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	i := len(reservoir) // let i be the index of the item under consideration

	for scanner.Scan() {
		line := scanner.Text()
		i++

		// The algorithm then generates a random number j
		// between (and including) 1 and i.
		j := rand.Intn(i)

		// If j is at most k, then the item is selected and replaces whichever
		// item currently occupies the j-th position in the reservoir.
		// Otherwise, the item is discarded.

		if j < k {
			reservoir[j] = line
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for idx := range reservoir {
		fmt.Println(reservoir[idx])
	}
}

func usage(fileName string, errors ...error) {
	fmt.Fprintf(os.Stderr, "%s: reservoir sampling of lines of test on stdin\n", fileName)
	fmt.Fprintf(os.Stderr, "usage: %s <NUMBER>\n", fileName)
	fmt.Fprintf(os.Stderr, "puts at most NUMBER randomly-chosen and ordered lines of stdin on stdout\n")
	for _, err := range errors {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
	os.Exit(1)
}
