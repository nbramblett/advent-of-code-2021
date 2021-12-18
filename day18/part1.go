package day18

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/nbramblett/advent-of-code-2021/util"
)

func Solve1() {
	ns := ReadLines()

	n := ns[0]
	for i := 1; i < len(ns); i++ {
		n = Add(n, ns[i])
		Resolve(n, true)
	}
	log.Println(n)
	log.Println(n.Magnitude())
}

// Recurse is false-able for step-wise debugging in case there's a loop.
func Resolve(n *Node, recurse bool) {
	ns := n.FindNested(0)
	for len(ns) > 0 {
		ns[0].Explode()
		ns = n.FindNested(0)
	}
	ns = n.FindSplits()
	if len(ns) > 0 {
		ns[0].Split()
		if recurse {
			Resolve(n, recurse)
		}
	}
}

type Node struct {
	parent      *Node
	left, right *Node
	val         int
}

func (n *Node) Magnitude() int {
	if n == nil {
		return 0
	}
	if n.IsRegular() {
		return n.val
	}
	return 3*n.left.Magnitude() + 2*n.right.Magnitude()
}

func (n *Node) String() string {
	if n == nil {
		return ""
	}
	if n.left == nil && n.right == nil {
		return fmt.Sprintf("%d", n.val)
	}
	return fmt.Sprintf("[%s,%s]", n.left.String(), n.right.String())
}

func (n *Node) Split() {
	n.left = &Node{parent: n, val: n.val / 2}
	n.right = &Node{parent: n, val: (n.val + 1) / 2}
	n.val = 0
}

func (n *Node) Explode() {
	explode(n, n.left.val, true)
	explode(n, n.right.val, false)
	n.left, n.right = nil, nil
	n.val = 0
}

func explode(n *Node, x int, isLeft bool) {
	p := n.parent
	target := p.left
	if !isLeft {
		target = p.right
	}
	if !target.IsRegular() {
		if target == n {
			if p.parent != nil {
				explode(p, x, isLeft)
			}
			return
		}
		descend(target, x, isLeft)
		return
	}
	target.val += x
}

func descend(n *Node, x int, isLeft bool) {
	if n.IsRegular() {
		n.val += x
		return
	}
	// Reverse it, I know it's weird to think about
	target := n.left
	if isLeft {
		target = n.right
	}
	descend(target, x, isLeft)
}

func (n *Node) IsRegular() bool {
	if n == nil {
		return false
	}
	return n.left == nil && n.right == nil
}

func (n *Node) FindNested(depth int) []*Node {
	if n == nil {
		return nil
	}
	ns := []*Node{}
	if depth == 4 && !n.IsRegular() {
		return []*Node{n}
	}
	ns = append(ns, n.left.FindNested(depth+1)...)
	ns = append(ns, n.right.FindNested(depth+1)...)
	return ns
}

func (n *Node) FindSplits() []*Node {
	if n == nil {
		return nil
	}
	ns := []*Node{}
	if n.val >= 10 {
		return []*Node{n}
	}
	ns = append(ns, n.left.FindSplits()...)
	ns = append(ns, n.right.FindSplits()...)
	return ns
}

func Add(n1, n2 *Node) *Node {
	p := &Node{left: n1, right: n2}
	n1.parent, n2.parent = p, p
	return p
}

func ReadLines() []*Node {
	file, err := os.Open("day18/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// first line of input is the numbers to read
	nodes := []*Node{}
	for scanner.Scan() {
		nodes = append(nodes, ParseNode(scanner.Text()))
	}
	return nodes
}

func ParseNode(s string) *Node {
	var n *Node
	for _, r := range s {
		if r == '[' {
			if n == nil {
				n = &Node{}
			} else {
				p := n
				n = &Node{parent: p}
				if p.left == nil {
					p.left = n
				} else {
					p.right = n
				}
			}
		} else if strings.Contains("1234567890", string(r)) {
			k := &Node{parent: n, val: MustParse(string(r))}
			if n.left == nil {
				n.left = k
			} else {
				n.right = k
			}
		} else if r == ']' {
			if n.parent != nil {
				n = n.parent
			}
		}
	}
	return n
}

func MustParse(s string) int {
	i, err := strconv.Atoi(s)
	util.PanicIf(err)
	return i
}
