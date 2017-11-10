package main

import(
	"generateDAG/dag"
	"flag"
)

func main() {

	flagPath := flag.String("path", "/tmp", "path to store dag file")
	flagBlock := flag.Uint64("block", 0, "block number")
	flag.Parse()


	dag.MakeDataset(*flagBlock, *flagPath)
}