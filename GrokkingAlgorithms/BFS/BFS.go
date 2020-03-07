package main

import (
	"fmt"

	"github.com/tebeka/deque"
)

func BFS(graph map[string][]string, startNode string, endNode string) (bool, error) {
	searchQueue := deque.New()
	for _, item := range graph[startNode] {
		searchQueue.Append(item)
	}

	seachedElems := []string{}

	for searchQueue.Len() > 0 {
		fmt.Println(searchQueue)
		checkElem, err := searchQueue.PopLeft()
		if err != nil {
			return false, err
		}

		var alreadySearched bool
		for _, item := range seachedElems {
			if item == checkElem {
				alreadySearched = true
				break
			}
		}

		if !alreadySearched {
			if checkElem == endNode {
				return true, nil
			}
			seachedElems = append(seachedElems, checkElem.(string))
			for _, item := range graph[checkElem.(string)] {
				searchQueue.Append(item)
			}
		}

	}
	return false, nil
}

func main() {
	hashTable := make(map[string][]string)
	hashTable["Mary"] = []string{"Vasya", "Oleg", "Sergey"}
	hashTable["Vasya"] = []string{"Oleg", "Tanya"}
	hashTable["Oleg"] = []string{"Sergey", "Tanya"}
	hashTable["Sergey"] = []string{"Katya"}
	hashTable["Tanya"] = []string{"Katya"}
	hashTable["Katya"] = []string{}
	if found, err := BFS(hashTable, "Mary", "Katya"); found && (err == nil) {
		fmt.Println("Found")
	} else if !found && (err == nil) {
		fmt.Println("Not found")
	} else {
		fmt.Println(err)
	}
}
