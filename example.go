package main

import (
	"fmt"
	"os"
	// import MPI binding for Go as mpi.
	mpi "github.com/JohannWeging/go-mpi"
)

func main() {
	mpi.Init(&os.Args)
	worldSize, _ := mpi.Comm_size(mpi.COMM_WORLD)
	rank, _ := mpi.Comm_rank(mpi.COMM_WORLD)
	name, _ := mpi.Get_processor_name()
	fmt.Printf("Hello world from processor %s, rank %d out of %d processors\n", name, rank, worldSize)
	mpi.Finalize()
}
