package main

import (
    "./person"
    "./structure"
    "./server"
)

func main() {
    ch := make(chan *structure.Req)
    quit := make(chan bool)
    go server.Forks(5, ch)
    go person.Philosopher(1, 0, 1, ch, quit)
    go person.Philosopher(2, 1, 2, ch, quit)
    go person.Philosopher(3, 2, 3, ch, quit)
    go person.Philosopher(4, 3, 4, ch, quit)
    go person.Philosopher(5, 4, 0, ch, quit)
    for n := 5; n > 0; n-- {
        <- quit
    }
}
