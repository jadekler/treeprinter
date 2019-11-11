# treeprinter
Print trees.

```
package treeprinter_test

import (
	"fmt"

	"github.com/jadekler/treeprinter"
)

func Example() {
	trie := &trieNode{letter: ' ', succs: make(map[rune]*trieNode)}
	for _, w := range []string{"FOOD", "FOOL", "FAIL", "FOOLISH", "SAD", "SWAM"} {
		trie.addSuccessive(w)
	}
	fmt.Println(trie)
	// Prints:
	// root
	// └──F
	// │  └──O
	// │  │  └──O
	// │  │     └──D
	// │  │     │  └──
	// │  │     └──L
	// │  │        └──
	// │  │        └──I
	// │  │           └──S
	// │  │              └──H
	// │  │                 └──
	// │  └──A
	// │     └──I
	// │        └──L
	// │           └──
	// └──S
	//    └──A
	// │     └──D
	// │        └──
	// └──W
	//    └──A
	//       └──M
	//          └──
}

// trieNode is a basic trie.
type trieNode struct {
	letter rune
	succs  map[rune]*trieNode
}

// addSuccess adds words to a trie.
func (tn *trieNode) addSuccessive(s string) {
	if len(s) == 0 {
		tn.succs[' '] = &trieNode{}
		return
	}
	r := rune(s[0])
	if _, ok := tn.succs[r]; !ok {
		tn.succs[r] = &trieNode{letter: r, succs: make(map[rune]*trieNode)}
	}
	tn.succs[r].addSuccessive(s[1:])
}

// String prints the trie with the treeprinter prettifier, so that it shows up
// nicely in the terminal.
func (t *trieNode) String() string {
	root := treeprinter.New("root")
	var dfs func(level *trieNode) *treeprinter.Node
	dfs = func(level *trieNode) *treeprinter.Node {
		newRoot := treeprinter.New(string(level.letter))
		for _, n := range level.succs {
			newRoot.Add(dfs(n))
		}
		return newRoot
	}
	for _, n := range t.succs {
		root.Add(dfs(n))
	}
	return root.Print()
}
```