package gee

import (
	"log"
	"sort"
	"strings"
)

type node struct {
	pattern  string  // 待匹配路由，例如 /p/:lang
	part     string  // 路由中的一部分，例如 :lang
	children []*node // 子节点，例如 [doc, tutorial, intro]
	isWild   bool    // 是否模糊匹配，part 含有 : 或 * 时为true
	isLeaf   bool    //是否为叶子节点
}

// 第一个匹配成功的节点，用于插入
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		//完全相等 或者都是:xxx 或者都是* 的情况
		if child.part == part || (len(child.part) > 0 && len(part) > 0 && (child.part[0] == part[0] && part[0] == ':')) {
			return child
		}
	}
	return nil
}

// 所有匹配成功的节点，用于查找
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

func (n *node) insert(pattern string, parts []string, height int) {

	if len(parts) == height {
		if n.isLeaf {
			log.Fatalf("冲突: pattern: %s 与 pattern: %s 同义!", n.pattern, pattern)
		}
		n.isLeaf = true
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{part: part, isWild: len(part) > 0 && (part[0] == ':' || part[0] == '*')}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	//if height > len(parts) {
	//	return nil
	//}

	if height == len(parts) {
		if n.isLeaf {
			//精准匹配
			if n.part == parts[height-1] {
				return n
			}
			//:xxx /*
			if n.isWild {
				return n
			}
		} else {
			return nil
		}
	} else if n.part == "**" {
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)
	indexMap := map[string]int{"": 0, ":": 1, "*": 2, "**": 3}
	//children
	sort.Slice(children, func(i, j int) bool {
		a := children[i]
		b := children[j]
		af := ""
		bf := ""
		if a.isWild {
			if strings.HasPrefix(a.part, ":") {
				af = ":"
			} else if strings.HasPrefix(a.part, "*") {
				af = "*"
			} else {
				af = "**"
			}
		}
		if b.isWild {
			if strings.HasPrefix(b.part, ":") {
				bf = ":"
			} else if strings.HasPrefix(b.part, "*") {
				bf = "*"
			} else {
				bf = "**"
			}
		}
		return indexMap[af] < indexMap[bf]
		//children[j]

		//// 优先精准
		//if !a.isWild && b.isWild {
		//	return true
		//}
		//// :xx -> /* -> /**
		//if a.isWild && b.isWild {
		//	if a.part == "**" {
		//		return false
		//	}
		//	if a.part == "*" {
		//		return false
		//	}
		//}
		//return true
	})
	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}
