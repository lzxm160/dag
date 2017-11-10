package main

import(
	"github.com/karachain/dag/dagUtils"
	"flag"
)

func main() {

	flagPath := flag.String("path", "/tmp", "path to store dag file")
	flagBlock := flag.Uint64("block", 0, "block number")
	flag.Parse()


	dagUtils.MakeDataset(*flagBlock, *flagPath)
}
