package graph

import (
	"errors"
	"math"
	"sort"
)

func SelectBestChild(n *Node) (*Node, error) {
	if n.IsTerminal() {
		return nil, errors.New("valueError: this is a terminal node, has no child")
	} else if len(n.Children) == 1 {
		return n.Children[0], nil
	}
	maxChild, maxUCT := n.Children[0], uctReward(n, n.Children[0], 1/math.Sqrt(2))
	for _, child := range n.Children[1:] {
		childReward := uctReward(n, child, 1/math.Sqrt(2))
		if childReward > maxUCT {
			maxUCT = childReward
			maxChild = child
		}
	}
	return maxChild, nil
}

func SelectSecureChild(n *Node) (*Node, error) {
	if n.IsTerminal() {
		return nil, errors.New("valueError: this is a terminal node, has no child")
	} else if len(n.Children) == 1 {
		return n.Children[0], nil
	}

	maxChild, maxLCT := n.Children[0], lctReward(n, n.Children[0], 1/math.Sqrt(2))
	for _, child := range n.Children[1:] {
		childReward := lctReward(n, child, 1/math.Sqrt(2))
		if childReward > maxLCT {
			maxLCT = childReward
			maxChild = child
		}
	}
	return maxChild, nil
}

func SelectMaxChild(n *Node) (*Node, error) {
	if n.IsTerminal() {
		return nil, errors.New("valueError: this is a terminal node, has no child")
	}
	if len(n.Children) == 0 {
		return nil, errors.New("valueError: Selecting max child from an unexplored node")
	} else if len(n.Children) == 1 {
		return n.Children[0], nil
	}
	sort.Slice(n.Children, func(i, j int) bool {
		return n.Children[i].Reward/n.Children[i].Visits < n.Children[j].Reward/n.Children[j].Visits
	})
	return n.Children[len(n.Children)-1], nil
}

func SelectRobustChild(n *Node) (*Node, error) {
	if n.IsTerminal() {
		return nil, errors.New("valueError: this is a terminal node, has no child")
	} else if len(n.Children) == 1 {
		return n.Children[0], nil
	}
	sort.Slice(n.Children, func(i, j int) bool {
		return n.Children[i].Visits < n.Children[j].Visits
	})
	return n.Children[len(n.Children)-1], nil
}

func uctReward(root *Node, child *Node, explorationConst float64) float64 {
	return (child.Reward / child.Visits) + (explorationConst * math.Sqrt(2*math.Log(root.Visits)/child.Visits))
}

func lctReward(root *Node, child *Node, explorationConst float64) float64 {
	return (child.Reward / child.Visits) - (explorationConst * math.Sqrt(2*math.Log(root.Visits)/child.Visits))
}
