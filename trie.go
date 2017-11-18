package main

type Node struct {
	end  bool
	nmap map[byte]*Node
}

func (n *Node) Add(entry []byte) {
	if len(entry) == 0 {
		n.end = true
		return
	}

	b := entry[0] // current byte
	if n.nmap[b] == nil {
		n.nmap[b] = &Node{end: false, nmap: make(map[byte]*Node)}
	}
	n.nmap[b].Add(entry[1:])
}

func main() {
	n := &Node{end: false, nmap: make(map[byte]*Node)}
	n.Add([]byte("a"))
	n.Add([]byte("b"))
	n.Add([]byte("aa"))
	n.Add([]byte("abc"))
}
