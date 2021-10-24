package shortestpath

type Graph interface {
	Vertices() []Vertex
	Neighbors(v Vertex) []Vertex
	Weight(u, v Vertex) int
}

// Nonnegative integer ID of vertex
type Vertex int

type GraphS struct {
	Vert  []Vertex
	Edges map[Vertex]map[Vertex]int
}

func (g GraphS) Edge(u, v Vertex, w int) {
	if _, ok := g.Edges[u]; !ok {
		g.Edges[u] = make(map[Vertex]int)
	}
	g.Edges[u][v] = w
}

func (g GraphS) Vertices() []Vertex { return g.Vert }

func (g GraphS) Neighbors(v Vertex) (vs []Vertex) {
	for k := range g.Edges[v] {
		vs = append(vs, k)
	}
	return vs
}

func (g GraphS) Weight(u, v Vertex) int { return g.Edges[u][v] }

func (g GraphS) Path(vv []Vertex) (int, int) {
	var lk = 0

	if len(vv) == 0 {
		return 0, 0
	}
	//s = strconv.Itoa(int(vv[0]))
	for _, v := range vv[1:] {
		//s += " -> " + strconv.Itoa(int(v))
		lk = int(v)
	}

	return int(vv[0]), lk
}

const Infinity = int(^uint(0) >> 1)

func FloydWarshall(g Graph) (dist map[Vertex]map[Vertex]int, next map[Vertex]map[Vertex]*Vertex) {
	vert := g.Vertices()
	dist = make(map[Vertex]map[Vertex]int)
	next = make(map[Vertex]map[Vertex]*Vertex)
	for _, u := range vert {
		dist[u] = make(map[Vertex]int)
		next[u] = make(map[Vertex]*Vertex)
		for _, v := range vert {
			dist[u][v] = Infinity
		}
		dist[u][u] = 0
		for _, v := range g.Neighbors(u) {
			v := v
			dist[u][v] = g.Weight(u, v)
			next[u][v] = &v
		}
	}
	for _, k := range vert {
		for _, i := range vert {
			for _, j := range vert {
				if dist[i][k] < Infinity && dist[k][j] < Infinity {
					if dist[i][j] > dist[i][k]+dist[k][j] {
						dist[i][j] = dist[i][k] + dist[k][j]
						next[i][j] = next[i][k]
					}
				}
			}
		}
	}
	return dist, next
}

func Path(u, v Vertex, next map[Vertex]map[Vertex]*Vertex) (path []Vertex) {
	if next[u][v] == nil {
		return
	}
	path = []Vertex{u}
	for u != v {
		u = *next[u][v]
		path = append(path, u)
	}
	return path
}
