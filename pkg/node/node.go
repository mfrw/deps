package node

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Name  string
	Id    int
	Type  string
	State string
}

func NewNode(name string, id int) *Node {
	// do some trickery here..
	// Probably an arena allocator
	// OR should we use the Cache ?
	// FOr now just the dumb New!
	n := new(Node)
	n.Name = name
	n.Id = id
	return n
}

// A very brittle function to parse a node out of the graph
func ProcessLineToNode(line string) (*Node, *Node, error) {
	if line[len(line)-1] == ';' {
		line = line[:len(line)-1]
	}

	parts := strings.Split(line, "->")
	from, err := parseNode(parts[0])
	if err != nil {
		return nil, nil, err
	}
	to, err := parseNode(parts[1])
	if err != nil {
		return nil, nil, err
	}

	return from, to, nil
}

func parseNode(line string) (*Node, error) {
	line = strings.TrimSpace(line)
	parts := strings.Split(line, " ")
	name := parts[0]
	id, typ, state, err := getIdTypeState(parts[1])
	if err != nil {
		return nil, err
	}
	n := NewNode(name, id)
	n.Type = typ
	n.State = state
	return n, nil
}

func getIdTypeState(line string) (int, string, string, error) {
	line = removeBothEnds(line)
	parts := strings.Split(line, ",")
	id, err := strconv.Atoi(strings.Split(parts[0], "=")[1])
	typ := strings.Split(parts[1], "=")[1]
	state := strings.Split(parts[2], "=")[1]
	return id, typ, state, err
}

func removeBothEnds(line string) string {
	return line[1 : len(line)-2]
}

func GetOnlyEdges(fname string) ([]string, error) {
	return slurpLines(fname, func(s string) bool { return strings.Contains(s, "->") })
}

// return only the lines that pass the predicate
// passing nil predicate return all lines
func slurpLines(fname string, pred func(ipt string) bool) ([]string, error) {
	f, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	if pred == nil {
		pred = func(s string) bool {
			return true
		}
	}

	res := []string{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// only consider files that actually have a '->' ignore rest
		l := scanner.Text()
		if pred(l) {
			res = append(res, l)
		}
	}

	return res, nil
}
