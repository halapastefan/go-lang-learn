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
func main() {

	csvFile, err := os.Open("callingCodes.csv")

	t := trie.New()

	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(csvFile)
	start := time.Now()
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
	elapsed := start.Sub(time.Now())
	fmt.Printf("Trie inicialization %s\n", elapsed)

	start = time.Now()

	i := 0
	for {
		r, _ := findPrefix("9378985320", t)
		t, _ := time.Parse(timeFormat, "2019-12-04T00:00:00.00Z")
		matchDate(*r, t)

		i++
		if i == 1 {
			break
		}
	}

	elapsed = start.Sub(time.Now())
	fmt.Printf("Prefix search took %s", elapsed)
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