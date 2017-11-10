package main

import(
	"generateDAG/dag"
)

func main() {

	var block uint64
	var dir string

	// enter block number
	// for default set to 0
	block = 0

	// enter path to directory to store DAG file
	dir = "/path/to/dir"


	dag.MakeDataset(block, dir)
}