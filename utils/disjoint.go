package utils

type DisjointSet[T comparable] struct {
	parent map[T]T
	rank   map[T]int
	count  int
}

func NewDisjointSet[T comparable]() *DisjointSet[T] {
	return &DisjointSet[T]{
		parent: make(map[T]T),
		rank:   make(map[T]int),
		count:  0,
	}
}

func (ds *DisjointSet[T]) Add(x T) {
	if _, ok := ds.parent[x]; ok {
		return
	}
	ds.parent[x] = x
	ds.rank[x] = 0
	ds.count++
}

func (ds *DisjointSet[T]) Find(x T) T {
	if _, ok := ds.parent[x]; !ok {
		ds.parent[x] = x
		ds.rank[x] = 0
		ds.count++
		return x
	}

	if ds.parent[x] == x {
		return x
	}
	root := ds.Find(ds.parent[x])
	ds.parent[x] = root
	return root
}

func (ds *DisjointSet[T]) Union(a, b T) bool {
	ra := ds.Find(a)
	rb := ds.Find(b)
	if ra == rb {
		return false
	}

	raRank := ds.rank[ra]
	rbRank := ds.rank[rb]

	if raRank < rbRank {
		ds.parent[ra] = rb
	} else if raRank > rbRank {
		ds.parent[rb] = ra
	} else {
		ds.parent[rb] = ra
		ds.rank[ra]++
	}
	ds.count--
	return true
}

func (ds *DisjointSet[T]) Connected(a, b T) bool {
	return ds.Find(a) == ds.Find(b)
}

func (ds *DisjointSet[T]) Count() int {
	return ds.count
}

func (ds *DisjointSet[T]) Sets() map[T][]T {
	groups := make(map[T][]T)
	for x := range ds.parent {
		root := ds.Find(x)
		groups[root] = append(groups[root], x)
	}
	return groups
}
