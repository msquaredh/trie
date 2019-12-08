package trie

type Trie struct {
    root *node
}

type node struct {
    children   map[rune]*node
    isTerminal bool
}

func New() *Trie {
    return &Trie{
        root: makeNode(),
    }
}

func makeNode() *node {
    return &node{
        children:   make(map[rune]*node),
        isTerminal: false,
    }
}

func (t *Trie) Insert(runes []rune) {
    current := t.root
    for _, r := range runes {
        if next, ok := current.children[r]; !ok {
            next = makeNode()
            current.children[r] = next
            current = next
        } else {
            current = next
        }
    }
    current.isTerminal = true
}

func (t *Trie) Search(runes []rune) bool {
    current := t.root
    for _, r := range runes {
        if next, ok := current.children[r]; ok {
            current = next
        } else {
            return false
        }
    }
    return current.isTerminal
}

func (t *Trie) Delete(runes []rune) bool {
    current := t.root
    var nodes []*node
    for _, r := range runes {
        if next, ok := current.children[r]; ok {
            nodes = append(nodes, next)
            current = next
        } else {
            return false
        }
    }
    if !current.isTerminal {
        return false
    }
    var (
        r      rune
        parent *node
    )
    for {
        if len(current.children) > 1 {
            current.isTerminal = false
            break
        }
        // Pop off the parent of current node
        parent, nodes = nodes[len(nodes)-2], nodes[:len(nodes)-2]
        // Pop off the last rune which is the current nodes key in the parent
        r, runes = runes[len(runes)-1], runes[:len(runes)-1]
        delete(parent.children, r)
        current = parent
    }
    return true
}
