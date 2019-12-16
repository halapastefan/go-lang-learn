package main

import (
	"encoding/csv"
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/derekparker/trie"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
)
var timeFormat = "2006-01-02T15:04:05.00Z"

type nodeValue struct {
	value []entry
}
type entry struct {
	price string
	initial int
	increment int
	startDate time.Time
}

func main() {

	csvFile, err := os.Open("callingCodes.csv")

	if err != nil {
		log.Fatal(err)
	}
	start := time.Now()
	t := createTrie(csvFile)
	elapsed := start.Sub(time.Now())
	fmt.Printf("Trie inicialization %s\n", elapsed)

	start = time.Now()

	i := 0
	for {
		res := findPrice(t, "939898123985", "2019-12-04T00:00:00.00Z")
		fmt.Println(res)
		i++
		if i == 1 {
			break
		}
	}

	elapsed = start.Sub(time.Now())
	fmt.Printf("Prefix search took %s", elapsed)
}

func findPrice(t *trie.Trie, calling string, date string) entry {
	r, _ := findPrefix(calling, t)
	d, _ := time.Parse(timeFormat, date)
	return matchDate(*r, d)
}

func findPrefix(nm string, trie *trie.Trie) (*nodeValue, bool) {

	i := 0
	for range nm {
		search := nm[0:len(nm)-i]
		node, found := trie.Find(search)
		if found {
			retVal := node.Meta().(nodeValue)
			return &retVal, true
		}
		i++
	}

	return nil, false
}
func matchDate(value nodeValue, date time.Time) entry {
	for _, v := range value.value{
		start := v.startDate
		end := start.AddDate(0, 1, 0)
		if date.After(start) && date.Before(end) {
			return v
		}
	}
	sort.Slice(value.value, func(i, j int) bool {
		return value.value[j].startDate.Before(value.value[i].startDate)
	})

	return value.value[0]
}

//From csv line create entry that will go to the trie
func createEntry(line []string) *entry  {
	price := line[3]
	initial, _ := strconv.Atoi(line[4])
	increment, _ := strconv.Atoi(line[5])
	start, dErr := time.Parse(timeFormat, line[6])
	if dErr != nil {
		fmt.Println(dErr)
	}

	return &entry{
		price:     price,
		initial:   initial,
		increment: increment,
		startDate: start,
	}
}

//Meta data of node will be appended if node already exists
func appendNode(t *trie.Trie, line []string){
	node, found := t.Find(line[0])
	if found {
		v := node.Meta().(nodeValue)
		v.value = append(v.value, *createEntry(line))
		t.Add(line[0], v)
	} else {
		en := make([]entry, 0)
		en = append(en, *createEntry(line))
		nv := nodeValue{
			value:en,
		}
		t.Add(line[0], nv)
	}
}

func createTrie(file *os.File) *trie.Trie {
	t := trie.New()

	reader := csv.NewReader(file)
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		if line[0]!="prefix" {
			appendNode(t, line)
		}

	}

	return t
}