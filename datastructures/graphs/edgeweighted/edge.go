package graphs

import "log"

// Edge struct used in edge-weighted (di)graphs
type Edge struct {
	v, w, weight int
}

// NewEdge constructs an edge with a given weight
func NewEdge(v int, w int, weight int) *Edge {
	return &Edge{v: v, w: w, weight: weight}
}

// Weight returns the weight of the edge
func (e Edge) Weight() int {
	return e.weight
}

// Either returns the vertex v
func (e Edge) Either() int {
	return e.v
}

// Other returns the opposite of the edge in the parameter
func (e Edge) Other(vertex int) int {
	if e.v == vertex {
		return e.w
	}
	if e.w == vertex {
		return e.v
	}
	log.Fatal("Inconsistent edge")
	return -1
}

// Compare two edges
func (e Edge) Compare(that Edge) int {
	if e.Weight() > that.Weight() {
		return +1
	}
	if e.Weight() < that.Weight() {
		return -1
	}
	return 0
}
