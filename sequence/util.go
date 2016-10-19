package sequence

/* Streamはint型のチャネル */
type Stream chan int

func MakeInt(n, m int) Stream {

    s := make(Stream)
    go func() {
        for i :=n; i <= m; i++ {
            /* sにiを送信 */
            s <- i
        }
        close(s)
    }()

    return s
}

func MakeNum(n int) Stream {
    s := make(Stream)

    go func() {
        for { s <- n }
    }()

    return s;
}

func MakeFibo() Stream {
    s := make(Stream)

    go func() {
        a, b := 1, 1
        for {
            s <- a
            a, b = b, a+b
            if a < 0 { break }
        }
        close(s)
    }()

    return s
}

func StreamMap(f func(int) int, in Stream) Stream {
    s := make(Stream)

    go func() {
        for {
            x, ok := <- in
            if !ok { break }
            s <- f(x)
        }
        close(s)
    }()

    return s
}

func StreamFilter(f func(int) bool, in Stream) Stream {
    s := make(Stream)
    go func () {
        for {
            x, ok := <-in
            if !ok { break }
            if f(x) {
                s <- x
            }
        }
        close(s)
    }()
    return s
}

func filter(n int, in Stream) Stream {
    return StreamFilter(func(x int) bool { return x% n != 0 }, in)
}

func Sieve(n int) Stream {
    s := make(Stream)

    go func() {
        in := MakeInt(2, n)
        for {
            x, ok := <-in
            if !ok { break }
            s<-x
            if x * x <= n {
                in = filter(x, in)
            }
        }
        close(s)
    }()
    return s
}