package main

func main() {
	composite := NewComposite("root")

	leaf := NewLeaf("leaf.txt")
	composite.Add(leaf)
	composite.Remove(leaf)


	composite2 := NewComposite("2Level")

	leaf2 := NewLeaf("test.txt")
	composite2.Add(leaf2)

	composite3 := NewComposite("3Level")
	leaf3 := NewLeaf("3LevelLeaf.txt")
	composite3.Add(leaf3)
	composite2.Add(composite3)
	composite.Add(composite2)

	composite.Operation()

}
