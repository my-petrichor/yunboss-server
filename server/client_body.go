package server

type ClientLoginBody struct {
	Type string `json:"type"`
	Uid  string `json:"uid"`
	Body string `json:"body"`
}

type ClientHeartBeatBody struct {
	Type  string `json:"type"`
	UID   string `json:"uid"`
	Token string `json:"token"`
	Body  string `json:"body"`
}

type ClientPushBody struct {
	Type  string `json:"type"`
	UID   string `json:"uid"`
	Token string `json:"token"`
	Body  string `json:"body"`
}

// server return client message after client sendmsg
type ClientReturnBody struct {
	Type   string `json:"type"`
	Status string `json:"status"`
	Msg    string `json:"msg"`
}
