package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"ssh_server/core"
	"ssh_server/global"
	"ssh_server/proxy/sshd/term"
	"ssh_server/proxy/sshd/termin"
	"strconv"
)

var UP = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// 解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handler(res http.ResponseWriter, req *http.Request) {
	// 服务升级
	ws, err := UP.Upgrade(res, req, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println()
	query := req.URL.Query()
	cols, _ := strconv.Atoi(query.Get("cols"))
	rows, _ := strconv.Atoi(query.Get("rows"))

	var xterm = "xterm-256color"
	ssh := global.Config.Ssh
	nextTerminal, err := term.NewNextTerminal(ssh.DestIP, ssh.DestPort, ssh.User, ssh.Pwd, "", "", rows, cols, "", xterm, true)

	if err != nil {
		return
	}

	if err := nextTerminal.RequestPty(xterm, rows, cols); err != nil {
		return
	}

	if err := nextTerminal.Shell(); err != nil {
		return
	}
	termHandler := termin.NewTermHandler("", "", "123", false, ws, nextTerminal)
	termHandler.Start()
	defer termHandler.Stop()

	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			// web socket会话关闭后主动关闭ssh会话
			break
		}
		fmt.Println(string(message))

		msg, err := ParseMessage(string(message))
		if err != nil {
			fmt.Println(err)
			continue
		}

		switch msg.Type {
		case 0:
			type Win struct {
				Cols int `json:"cols"`
				Rows int `json:"rows"`
			}
			byteData, _ := json.Marshal(msg.Content)

			var win Win
			json.Unmarshal(byteData, &win)

			if err := termHandler.WindowChange(win.Rows, win.Cols); err != nil {
			}
			fmt.Println(1, msg.Content)

		case 1: // 发送数据
			input := []byte(msg.Content.(string))
			fmt.Println(msg.Content)
			err := termHandler.Write(input)
			if err != nil {
				break
			}

		}
	}
	fmt.Println("服务关闭")
}

func ParseMessage(content string) (msg Message, err error) {
	err = json.Unmarshal([]byte(content), &msg)
	return
}

type Message struct {
	Type    int `json:"type"`
	Content any `json:"content"`
}

func (s Message) String() string {
	byteData, _ := json.Marshal(s)
	return string(byteData)
}

func main() {
	core.InitConf()
	addr := global.Config.System.Addr()
	fmt.Printf("ssh_server 运行在：%s\n", addr)
	http.HandleFunc("/", handler)
	http.ListenAndServe(addr, nil)
}
