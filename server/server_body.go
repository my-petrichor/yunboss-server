package server

type ServerLoginBody struct {
	Type   string `json:"type"`
	Ip     string `json:"ip"`
	UID    string `json:"uid"`
	Body   string `json:"body"`
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Token  string `json:"token"`
}

type ServerHeartBeatBody struct {
	Type   string `json:"type"`
	Status string `json:"status"`
	Msg    string `json:"msg"`
}

type ServerPushBody struct {
	Type string `json:"type"`
	Body string `json:"body"`
	URL  string `json:"url"`
}

// client return server message after server sendmsg
type ServerReturnBody struct {
	Type   string `json:"type"`
	Msg    string `json:"msg"`
	Status int    `json:"status"`
	Body   string `json:"body"`
}
