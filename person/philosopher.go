package person

import(
    "../structure"
    "../server"
    "fmt"
    "time"
)

/*
 * 哲学者の振る舞い
 *
 * m 哲学者の番号
 * forkR 右側のフォーク
 * forkL 左側のフォーク
 * out 最初にリクエストを受け付けるチャネル
 * quit 終了通知の状態をもつチャネル
 */
func Philosopher(m, forkR, forkL int, out chan *structure.Req, quit chan bool) {
    in := make(chan bool)

    //goにwhile()はない
    for n := 2; n > 0;{
        fmt.Printf("Philosopher %d is thinking", m)
        fmt.Println()
        time.Sleep(1000 * time.Millisecond)
        server.GetFork(forkR, out, in)

        //左のフォークがとれるかどうか
        if server.GetFork1(forkL, out, in) {
            fmt.Printf("Philosopher %d is eating", m)
            fmt.Println()
            time.Sleep(500 * time.Millisecond)
            server.RetFork(forkR, out, in)
            server.RetFork(forkL, out, in)
            n--
        } else {
            //右のフォークを取って左が取れないときは右のフォークを元に戻す
            server.RetFork(forkR, out, in)
        }
    }

    fmt.Printf("Philosopher %d is sleeping", m)
    fmt.Println()
    quit <- true
}
