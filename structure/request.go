package structure

const (
    GET = 0 //フォークの取得
    RET = 1 //フォークの返却
)

/*
 * フォークのリクエスト
 *
 * Req リクエストの種類
 * Fork フォークの番号
 * Reply 応答用のチャネル
 */
type Req struct {
    Req, Fork int
    Reply chan<-bool
}

//リクエストの生成
func NewReq(req, fork int, reply chan bool) *Req {
    p := new(Req)
    p.Req = req
    p.Fork = fork
    p.Reply = reply
    return p
}