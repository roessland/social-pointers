package main

func generateIds(ids chan<- int) {
    id := 1
    for {
        ids <- id
        id++
    }
}

func newIdGenerator() <-chan int {
    idGen := make(chan int)
    go generateIds(idGen)
    return idGen
}

var idGen = newIdGenerator()
