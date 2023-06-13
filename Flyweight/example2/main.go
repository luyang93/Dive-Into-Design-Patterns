package main

import (
	"fmt"
	"runtime"
	"strconv"
)

// TreeType is a flyweight that contains a portion of the state of a tree.
type TreeType struct {
	Name, Color, Texture string
}

// Tree is a contextual object that contains the extrinsic part of the tree state.
type Tree struct {
	X, Y int
	Type *TreeType
}

type Tree2 struct {
	X, Y                 int
	Name, Color, Texture string
}

// Forest is a collection of trees.
type Forest struct {
	Trees  []*Tree
	Trees2 []*Tree2
}

// TreeFactory is a flyweight factory that decides whether to reuse existing flyweight or to create a new object.
type TreeFactory struct {
	TreeTypes map[string]*TreeType
}

func NewTreeFactory() *TreeFactory {
	return &TreeFactory{TreeTypes: make(map[string]*TreeType)}
}

// GetTreeType returns a tree type with the given name, color, and texture.
// If such a tree type does not exist, a new one is created and stored.
func (f *TreeFactory) GetTreeType(name, color, texture string) *TreeType {
	if treeType, ok := f.TreeTypes[name]; ok {
		return treeType
	} else {
		treeType := &TreeType{name, color, texture}
		f.TreeTypes[name] = treeType
		return treeType
	}
}

// PlantTree plants a tree of the given type at the given coordinates in the forest.
func (f *Forest) PlantTree(x, y int, t *TreeType) {
	tree := &Tree{x, y, t}
	f.Trees = append(f.Trees, tree)
}

func (f *Forest) PlantTree2(x, y int) {
	tree := &Tree2{x, y, strconv.Itoa(x), strconv.Itoa(x), strconv.Itoa(x)}
	f.Trees2 = append(f.Trees2, tree)
}

// Draw prints all trees in the forest to the console.
func (f *Forest) Draw() {
	for _, t := range f.Trees {
		fmt.Printf("Tree at %d,%d of type %s\n", t.X, t.Y, t.Type.Name)
	}
}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Printf("Memory Usage: Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func main() {
	PrintMemUsage()
	factory := NewTreeFactory()
	forest := &Forest{}

	greenTree := factory.GetTreeType("green", "green", "tree")
	for i := 0; i < 10000000; i++ {
		forest.PlantTree(1, 1, greenTree)
	}
	//forest.Draw()
	PrintMemUsage()
	for i := 0; i < 10000000; i++ {
		forest.PlantTree2(1, 1)
	}
	PrintMemUsage()
}
