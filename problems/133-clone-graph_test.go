package problems

import "testing"

func TestCloneGraph(t *testing.T) {
	graph := &GraphNode{
		Val:       1,
		Neighbors: []*GraphNode{},
	}
	graph.Neighbors = append(graph.Neighbors, &GraphNode{
		Val:       2,
		Neighbors: []*GraphNode{},
	})
	graph.Neighbors[0].Neighbors = append(graph.Neighbors[0].Neighbors, graph)
	cloned := cloneGraph(graph)
	if cloned == nil {
		t.Fatal("cloned graph is nil")
	}
	if cloned.Val != 1 {
		t.Fatalf("cloned graph has wrong value: got %d, want 1", cloned.Val)
	}
	if len(cloned.Neighbors) != 1 {
		t.Fatalf("cloned graph has wrong number of neighbors: got %d, want 1", len(cloned.Neighbors))
	}
	if cloned.Neighbors[0].Val != 2 {
		t.Fatalf("cloned graph has wrong neighbor value: got %d, want 2", cloned.Neighbors[0].Val)
	}
	if len(cloned.Neighbors[0].Neighbors) != 1 {
		t.Fatalf("cloned graph's neighbor has wrong number of neighbors: got %d, want 1", len(cloned.Neighbors[0].Neighbors))
	}
	if cloned.Neighbors[0].Neighbors[0].Val != 1 {
		t.Fatalf("cloned graph's neighbor has wrong neighbor value: got %d, want 1", cloned.Neighbors[0].Neighbors[0].Val)
	}
}
