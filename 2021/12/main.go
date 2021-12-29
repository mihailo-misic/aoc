package main

import (
	"fmt"
	"strings"

	. "../utils"
)

type Vertex struct {
	Key      string
	Vertices map[string]*Vertex
}

func NewVertex(key string) *Vertex {
	return &Vertex{
		Key:      key,
		Vertices: map[string]*Vertex{},
	}
}

type Graph struct {
	Vertices map[string]*Vertex
}

func NewGraph() *Graph {
	return &Graph{
		Vertices: map[string]*Vertex{},
	}
}

func (g *Graph) AddVertex(key string) {
	if _, ok := g.Vertices[key]; ok {
		return
	}
	g.Vertices[key] = NewVertex(key)
}

func (g *Graph) AddEdge(k1, k2 string) {
	v1 := g.Vertices[k1]
	v2 := g.Vertices[k2]

	if _, ok := v1.Vertices[v2.Key]; ok {
		return
	}

	v1.Vertices[v2.Key] = v2
	v2.Vertices[v1.Key] = v1
}

var counter = 0
var paths = [][]string{}

type Seeker struct {
	Jumped  bool
	Key     int
	Visited []string
	At      *Vertex
}

var SP = map[int]*Seeker{}

func NewSeeker(visited []string, at *Vertex, jumped bool) *Seeker {
	counter++

	s := &Seeker{
		Jumped:  jumped,
		Key:     counter,
		Visited: visited,
		At:      at,
	}

	SP[s.Key] = s

	return s
}

func (s *Seeker) GoTo(v *Vertex) {
	s.At = v
	isSmall := v.Key == strings.ToLower(v.Key)
	if isSmall && Includes(s.Visited, v.Key) {
		s.Jumped = true
	}
	s.Visited = append(s.Visited, v.Key)

	if v.Key == "end" {
		paths = append(paths, s.Visited)
		s.Die()
	}
}
func IncludesTwice(sl []string, str string) bool {
	found := 0
	for _, c := range sl {
		if c == str {
			found++
			if found == 2 {
				return true
			}

		}
	}
	return false
}
func (s *Seeker) FindOptions() (options []*Vertex) {
	for _, v := range s.At.Vertices {
		isBig := v.Key == strings.ToUpper(v.Key)
		isStart := v.Key == "start"
		wasVisited := Includes(s.Visited, v.Key)
		if !isBig && !isStart && wasVisited && s.Jumped == false {
			wasVisited = IncludesTwice(s.Visited, v.Key)
		}

		if isBig || (!isStart && !wasVisited) {
			options = append(options, v)
		}
	}
	return
}
func (s *Seeker) SendCloneTo(v *Vertex) {
	newVisited := []string{}
	for _, v := range s.Visited {
		newVisited = append(newVisited, v)
	}
	clone := NewSeeker(newVisited, s.At, s.Jumped == true)
	clone.GoTo(v)
}
func (s *Seeker) Die() {
	delete(SP, s.Key)
}

func main() {
	input := ReadFile("./input.txt")

	g := NewGraph()

	for _, line := range input {
		rooms := strings.Split(line, "-")
		g.AddVertex(rooms[0])
		g.AddVertex(rooms[1])
		g.AddEdge(rooms[0], rooms[1])
	}

	NewSeeker([]string{"start"}, g.Vertices["start"], false)

	for len(SP) > 0 {
		spKeys := []int{}
		for k := range SP {
			spKeys = append(spKeys, k)
		}

		for _, key := range spKeys {
			s := SP[key]
			options := s.FindOptions()
			for _, o := range options {
				s.SendCloneTo(o)
			}
			s.Die()
		}
	}

	fmt.Println("Answer", len(paths))
}
