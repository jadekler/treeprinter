// Treeprinter prints trees in a visually appealing manner.
package treeprinter

// NOTE: Original inspiration from https://github.com/DiSiqueira/GoTree.

const (
	newLine      = "\n"
	emptySpace   = "   "
	continueItem = "│  "
	lastItem     = "└──"
)

// A node holds some text and references to other nodes.
type Node struct {
	text  string
	succs []*Node
}

// New creates a new node.
func New(text string) *Node {
	return &Node{
		text:  text,
		succs: []*Node{},
	}
}

// Add adds a node.
func (t *Node) Add(n *Node) {
	t.succs = append(t.succs, n)
}

// Val returns the value of the node.
func (t *Node) Val() string {
	return t.text
}

// Print prints the node and its successors with formatting.
func (t *Node) Print() string {
	return t.Val() + newLine + printSuccs(t.succs, []bool{})
}

func printSuccs(succs []*Node, spaces []bool) string {
	var result string
	for i, s := range succs {
		last := i == len(succs)-1
		result += printText(s.Val(), spaces)
		if len(s.succs) > 0 {
			spacesChild := append(spaces, last)
			result += printSuccs(s.succs, spacesChild)
		}
	}
	return result
}

func printText(text string, spaces []bool) string {
	var result string
	for _, space := range spaces {
		if space {
			result += emptySpace
		} else {
			result += continueItem
		}
	}
	return result + lastItem + text + newLine
}
