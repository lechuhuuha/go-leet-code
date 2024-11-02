1. What data structure is used to represent a set of linked objects?

- A graph is a representation of a set of objects that's connected by links
- The links connect vertices, which are points.
- The basic operations on a graph are the addition and removal of links and vertices

2. What is a two-dimensional matrix with Boolean values called?

- Its called zero matrix
- A null or a zero matrix is a matrix entirely consisting of zero values.

```go
var matrix = [3][3] int{
 {0,0,0},
 {0,0,0},
 {0,0,0}
}
```

3. Give a code example for a network representation using a graph.

```go
package graph

import "fmt"

type (
 SocialGraph struct {
  Size  int
  Links [][]Link
 }

 Link struct {
  Vertex1     int
  Vertex2     int
  LightWeight int
 }
)

func NewSocialGraph(num int) *SocialGraph {
 return &SocialGraph{
  Size:  num,
  Links: make([][]Link, num),
 }
}

func (s *SocialGraph) AddLink(vertex1, vertex2, weight int) {
 s.Links[vertex1] = append(s.Links[vertex1], Link{
  Vertex1:     vertex1,
  Vertex2:     vertex2,
  LightWeight: weight,
 })
}

func (s *SocialGraph) PrintLinks() {
 vertex := 0
 for _, link := range s.Links[vertex] {
  fmt.Printf("Link : %d -> %d (%d)\n", link.Vertex1, link.Vertex2, link.LightWeight)
 }

 for _, adjacent := range s.Links {
  for _, link := range adjacent {
   fmt.Printf("Link : %d -> %d (%d)\n", link.Vertex1, link.Vertex2, link.LightWeight)
  }
 }
}

```

4. Which metrics can be calculated from a social graph?

- Sparse matrix can use social graph. It can be used for solving large-scale
problems that do not require dense matrices

5. What is a cartographic design?

- Cartographic design is referred to as map creation using geographic
information.

```go
type (
 Place struct {
  Name      string
  Latitude  float64
  Longitude float64
 }

 MapLayout struct {
  GraphNodes map[Place]struct{}
  Links      map[Place]map[Place]struct{}
 }
)

```

6. Give an example of a knowledge graph and define the class, slots, and facets.

```go
type Class struct {
 Name string
}

type KnowledgeGraph struct {
 GraphNodes map[Class]struct{}
 Links      map[Class]map[Class]struct{}
}

func NewKnowledgeGraph() *KnowledgeGraph {
 return &KnowledgeGraph{
  GraphNodes: map[Class]struct{}{},
  Links:      map[Class]map[Class]struct{}{},
 }
}
```

7. What are the applications of sparse matrices?

- Image Processing: Sparse matrices are utilized in image processing applications, especially in methods such as compressed sensing and image compression,
    where images can be represented using fewer non-zero elements.

8. Define a list of lists and write a code sample.

```go
type Cell struct {
 Row    int
 Column int
 Value  float64
}

type SparseMatrix struct {
 Cells []Cell
 shape [2]int
}

func NewSparseMatrix(m int, n int) *SparseMatrix {
 return &SparseMatrix{
  Cells: []Cell{},
  shape: [2]int{m, n},
 }
}


```

9. What is a map layout?

- A map layout is a geographical visualization of locations that are linked together.
- The nodes in the graph of a map consist of geo-based information.
- The node will have information such as the name of the location, latitude, and longitude.

```go
type (
 Place struct {
  Name      string
  Latitude  float64
  Longitude float64
 }

 MapLayout struct {
  GraphNodes map[Place]struct{}
  Links      map[Place]map[Place]struct{}
 }
)

func NewMapLayout() *MapLayout {
 return &MapLayout{
  GraphNodes: map[Place]struct{}{},
  Links:      map[Place]map[Place]struct{}{},
 }
}

```

10. What different operations can be performed using graphs?

- The basic operations on a graph are the addition and removal of links and vertices.
