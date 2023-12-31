package pkg

type NodeHeap struct {
	Nodes  []NodeLoss
	Target Location
}

func (h NodeHeap) Len() int {
	return len(h.Nodes)
}

func (h NodeHeap) Less(i, j int) bool {
	iCost := h.Nodes[i].MinLossToHere + h.Nodes[i].ManhattanDist(h.Target)
	jCost := h.Nodes[j].MinLossToHere + h.Nodes[j].ManhattanDist(h.Target)
	return iCost < jCost || (iCost == jCost && h.Nodes[i].ArrivalDir < h.Nodes[j].ArrivalDir)
}

func (h NodeHeap) Swap(i, j int) {
	h.Nodes[j], h.Nodes[i] = h.Nodes[i], h.Nodes[j]
}

func (h *NodeHeap) Push(x any) {
	h.Nodes = append(h.Nodes, x.(NodeLoss))
}

func (h *NodeHeap) Pop() any {
	old := h.Nodes
	n := len(old)
	x := old[n-1]
	h.Nodes = old[0 : n-1]
	return x
}
