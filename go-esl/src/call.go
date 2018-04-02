package main

import (
        "fmt"
        "log"
	"time"
        "github.com/vma/esl"
	"github.com/go-redis/redis"
)

type Handler struct {
        CallId  string
        BgJobId string
}

const (
        Caller = "018882000"
        Callee = "2348148515007"
	Codec  = "PCMU,PCMA"

)

var Client = redis.NewClient(&redis.Options{
                Addr:     "localhost:6379",
                Password: "", // no password set
                DB:       0,  // use default DB
})

func main() {
	fmt.Println("main")
        handler := &Handler{}
        con, err := esl.NewConnection("127.0.0.1:8021", handler)
        if err != nil {
                log.Fatal("ERR connecting to freeswitch:", err)
        }
        con.HandleEvents()
	for{
	}
}

func (h *Handler) OnConnect(con *esl.Connection) {
        con.SendRecv("event", "plain", "ALL")
	for {
                lPop := Client.LPop("calllist")
                if lPop.Val() == "" {
                        fmt.Println("hello")
			time.Sleep(2000 * time.Millisecond)
                }else{
                        BgJobId, _ := con.BgApi("originate", "{origination_caller_id_number="+Caller+",absolute_codec_string="+Codec+",execute_on_answer='transfer ANSWEREDCALL XML default'}sofia/gateway/178.62.27.239/"+lPop.Val(), "&park()")
                        log.Println("originate bg job id:", BgJobId)
                        time.Sleep(2000 * time.Millisecond)
                }
        }

}

func (h *Handler) OnDisconnect(con *esl.Connection, ev *esl.Event) {
        log.Println("esl disconnected:", ev)
}

func (h *Handler) OnClose(con *esl.Connection) {
        log.Println("esl connection closed")
}

func (h *Handler) OnEvent(con *esl.Connection, ev *esl.Event) {
        log.Printf("%s - event %s %s %s\n", ev.UId, ev.Name, ev.App, ev.AppData)
       // fmt.Println(ev) // send to stderr as it is very verbose
	fmt.Println("callid")
	fmt.Println(ev.UId)
        switch ev.Name {
        case esl.BACKGROUND_JOB:
                log.Printf("bg job result:%s\n", ev.GetTextBody())
        case esl.CHANNEL_ANSWER:
                log.Println("call answered, starting moh")
           //     con.Execute("playback", ev.UId, "/tmp/UDOM.wav")
        case esl.CHANNEL_HANGUP:
                hupcause := ev.Get("Hangup-Cause")
                log.Printf("call terminated with cause %s", hupcause)
               // con.Close()
        }
}
