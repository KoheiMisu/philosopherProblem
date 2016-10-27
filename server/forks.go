package server

import (
    "../structure"
    "time"
)


/*
 * フォークを管理する
 * forループ中でリクエストを受け付ける
 * チャネルchからリクエストを取り出して変数rにセットする
 *
 * n フォークの本数
 * ch リクエストを受け付けるチャネル
 */
func Forks(n int, ch chan *structure.Req) {
    forkTable := make([]bool, n)

    for i := 0; i < n; i++ {
        forkTable[i] = true
    }

    for {
        r := <- ch

        switch r.Req {

            //フォークの取得
            case structure.GET:
                if forkTable[r.Fork] {
                    forkTable[r.Fork] = false
                    r.Reply <- true
                } else {
                    r.Reply <- false
                }

            //フォークの返却
            case structure.RET:
                forkTable[r.Fork] = true
                r.Reply <- true
        }
    }
}

func GetFork(fork int, out chan *structure.Req, in chan bool) int {
    r := structure.NewReq(structure.GET, fork, in)

    for {
        out <- r
        if <- in {
            time.Sleep(100 * time.Millisecond)
            return fork
        } else {
            time.Sleep(500 * time.Millisecond)
        }
    }
}

func RetFork(fork int, out chan *structure.Req, in chan bool) bool {
    time.Sleep(100 * time.Millisecond)
    out <- structure.NewReq(structure.RET, fork, in)
    return <- in
}