package main

import "fmt" 

func hasAntonBFS(graph map[string][]string, start string) bool {
	var queue []string
	queue = append(queue, graph[start]...)
	searched := make(map[string]bool)

	for len(queue) > 0 {
		person := queue[0]
		queue = queue[1:]

		if _, notSearched := searched[person]; !notSearched {
			if isAnton(person) { return true }
			queue = append(queue, graph[person]...)
			searched[person] = true
		}
	}
	return false
}


func isAnton(name string) bool {
	return name == "anton"
}

func main() {
	graph := make(map[string][]string)
	graph["me"] = []string{"alice", "bob", "claire"}

	graph["alice"] = []string{"peggy", "mike"}
	graph["bob"] = []string{"tom", "jonny"}
	graph["claire"] = []string{"kate"}

	graph["peggy"] = []string{"lucas"}
	graph["mike"] = []string{"lily"}
	graph["tom"] = []string{}
	graph["jonny"] = []string{}
	graph["kate"] = []string{"sophia"}

	graph["lucas"] = []string{"emily"}
	graph["lily"] = []string{"noah"}
	graph["sophia"] = []string{"anton"}

	graph["emily"] = []string{}
	graph["noah"] = []string{}
	graph["anton"] = []string{}
	
	if hasAntonBFS(graph, "me") { 
		fmt.Println("You have a friend Anton") 
	} else {
		fmt.Println("You don't have a friend Anton")
	}
}