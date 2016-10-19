package person

import(
    "../structure"
    "../server"
    "fmt"
    "time"
)

func Philosopher(m, forkR, forkL int, out chan *structure.Req, quit chan bool) {
    in := make(chan bool)
    for n := 2; n > 0; n-- {
        fmt.Printf("Philosopher %d is thinking", m)
        fmt.Println()
        time.Sleep(1000 * time.Millisecond)
        server.GetFork(forkR, out, in)
        server.GetFork(forkL, out, in)
        fmt.Printf("Philosopher %d is eating", m)
        fmt.Println()
        server.RetFork(forkR, out, in)
        server.RetFork(forkL, out, in)
    }

    fmt.Printf("Philosopher %d is sleeping", m)
    fmt.Println()
    quit <- true
}
