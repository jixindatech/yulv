package service

const cacheUser = "users"
const cacheIP = "ip"
const cacheReqRule = "reqrule"

type cacheHead struct {
	ID        uint        `json:"id"`
	Timestamp int64       `json:"timestamp"`
	Config    interface{} `json:"config"`
}

type database struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
}

type dbUser struct {
	User     string     `json:"user"`
	Password string     `json:"password"`
	Database []database `json:"database"`
}

type ip struct {
	Type string   `json:"type"`
	Data []string `json:"data"`
}

type rule struct {
}
