package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"unicode"
)

// Counter struct
type Counter struct {
	values [][]byte
	keys   []int
}

func main() {
	//reading text file
	var m1, m2 runtime.MemStats
	runtime.ReadMemStats(&m1)
	data, err := ioutil.ReadFile("mobydick.txt")
	if err != nil {
		log.Fatal(err)
	}
	//Trim if its not letter and divide there
	//to create matrix of bytes
	f := func(c rune) bool {
		return !unicode.IsLetter(c)
	}

	words := bytes.FieldsFunc(data, f)
	//declaring Struct
	var counter Counter

	counter.uniq(words)
	counter.sort()
	// counter.keys, counter.values = quicksort(counter.keys, counter.values)
	counter.print()
	runtime.ReadMemStats(&m2)
	fmt.Println("memory used", m2.Alloc-m1.Alloc)

}

//creats uniq words and creats counts of them
// func (c *Counter) uniq(words [][]byte) {
// 	for _, word := range words {
// 		n := exist(word, c.values)
// 		if n == -1 {
// 			c.values = append(c.values, word)
// 			c.keys = append(c.keys, 1)
// 		} else {
// 			c.keys[n] = c.keys[n] + 1
// 		}
// 	}
// }

func (c *Counter) uniq(words [][]byte) {
	for i := range words {
		n := exist(words[i], c.values)
		if n == -1 {
			c.values = append(c.values, words[i])
			c.keys = append(c.keys, 1)
		} else {
			c.keys[n] = c.keys[n] + 1
		}
	}
}

//checking word for existens, if not exist returnig flag,
//if yes returning index of the word
// func exist(source []byte, target [][]byte) int {
// 	for i, word := range target {
// 		if bytes.EqualFold(word, source) {
// 			return i
// 		}
// 	}
// 	return -1
// }

func exist(source []byte, target [][]byte) int {
	for i := range target {
		if bytes.EqualFold(target[i], source) {
			return i
		}
	}
	return -1
}

//sorting keys and by them sorting values
// func (c *Counter) sort() {
// 	for {
// 		isChanged := false

// 		for k := 0; k < len(c.keys); k++ {
// 			if k+1 == len(c.keys) {
// 				break
// 			}
// 			if c.keys[k] < c.keys[k+1] {
// 				c.keys[k], c.keys[k+1] = c.keys[k+1], c.keys[k]
// 				c.values[k], c.values[k+1] = c.values[k+1], c.values[k]
// 				isChanged = true
// 			}
// 		}
// 		if !isChanged {
// 			break
// 		}
// 	}

// }

func (c *Counter) sort() {
	var n = len(c.keys)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if c.keys[j-1] < c.keys[j] {
				c.keys[j-1], c.keys[j] = c.keys[j], c.keys[j-1]
				c.values[j-1], c.values[j] = c.values[j], c.values[j-1]
			}
			j = j - 1
		}
	}
}

// func quicksort(a []int, b [][]byte) ([]int, [][]byte) {
// 	if len(a) < 2 {
// 		return a, b
// 	}

// 	left, right := 0, len(a)-1

// 	pivot := rand.Int() % len(a)
// 	if len(a) < 2 {
// 		return a, b
// 	}

// 	left, right := 0, len(a)-1

// 	pivot := rand.Int() % len(a)

// 	a[pivot], a[right] = a[right], a[pivot]

// 	for i := range a {
// 		if a[i] > a[right] {
// 			a[left], a[i] = a[i], a[left]
// 			b[left], b[i] = b[i], b[left]
// 			left++
// 		}
// 	}

// 	a[left], a[right] = a[right], a[left]

// 	quicksort(a[:left], b[:left])
// 	quicksort(a[left+1:], b[left+1:])

// 	return a, ba[left]

// 	quicksort(a[:left], b[:left])
// 	quicksort(a[left+1:], b[left+1:])

// 	return a, b
// }

//printing by Stdout Writer
//and before print converting letter to lower case
// func (c *Counter) print() {
// 	for i, word := range c.values[:20] {
// 		word = bytes.ToLower(word)
// 		fmt.Printf("%d ", c.keys[i])
// 		os.Stdout.Write(word)
// 		fmt.Println()
// 	}
// }

func (c *Counter) print() {
	words20 := c.values[:20]
	for i := range words20 {
		words20[i] = bytes.ToLower(words20[i])
		fmt.Printf("%d ", c.keys[i])
		os.Stdout.Write(words20[i])
		fmt.Println()
	}
}
