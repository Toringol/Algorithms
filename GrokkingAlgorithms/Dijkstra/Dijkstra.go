package main

import "fmt"

const Infinity = int(^uint(0) >> 1)

func NewCostTable(graph map[string]map[string]int, startNode string) map[string]int {
	costs := make(map[string]int)

	for node, cost := range graph {
		for root, value := range cost {
			if node == startNode {
				costs[root] = value
			}
			if _, ok := costs[root]; !ok {
				costs[root] = Infinity
			}

		}
	}

	return costs
}

func NewParentTable(graph map[string]map[string]int, startNode string) map[string]string {
	parents := make(map[string]string)

	for node, cost := range graph {
		for root := range cost {
			if node == startNode {
				parents[root] = startNode
			}
			if _, ok := parents[root]; !ok {
				parents[root] = "-"
			}
		}
	}

	return parents
}

func FindLowestNode(costs map[string]int, processedNodes map[string]struct{}) string {
	lowestCost := Infinity
	lowestNode := ""

	for node, cost := range costs {
		if _, ok := processedNodes[node]; !ok && cost < lowestCost {
			lowestCost = cost
			lowestNode = node
		}
	}

	return lowestNode
}

func Dijkstra(graph map[string]map[string]int, startNode string, endNode string) int {

	costs := NewCostTable(graph, startNode)
	parents := NewParentTable(graph, startNode)
	processedNodes := make(map[string]struct{})

	node := FindLowestNode(costs, processedNodes)

	fmt.Println(costs)
	fmt.Println(parents)
	fmt.Println(node)
	fmt.Println(graph[node])

	for i := 0; i < len(graph)-2; i++ {
		cost := costs[node]
		neighbors := graph[node]
		for n := range neighbors {
			newCost := cost + neighbors[n]
			if costs[n] > newCost {
				costs[n] = newCost
				parents[n] = node
			}
		}
		processedNodes[node] = struct{}{}
		node = FindLowestNode(costs, processedNodes)
	}

	fmt.Println(costs)
	fmt.Println(parents)

	return 0
}

func GenerateNewGraph() map[string]map[string]int {
	graph := make(map[string]map[string]int)

	graph["start"] = map[string]int{}
	graph["start"]["A"] = 6
	graph["start"]["B"] = 2

	graph["A"] = map[string]int{}
	graph["A"]["fin"] = 1

	graph["B"] = map[string]int{}
	graph["B"]["A"] = 3
	graph["B"]["fin"] = 5
	graph["fin"] = map[string]int{}

	return graph
}

func main() {
	graph := GenerateNewGraph()
	fmt.Println(graph)
	Dijkstra(graph, "start", "fin")
}
