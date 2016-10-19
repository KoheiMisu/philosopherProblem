package structure

const (
    GET = 0 //フォークの取得
    RET = 1 //フォークの返却
)

//フォークのリクエスト
type Req struct {
    Req, Fork int
    Reply chan<-bool
}

func NewReq(req, fork int, reply chan bool) *Req {
    p := new(Req)
    p.Req = req
    p.Fork = fork
    p.Reply = reply
    return p
}