package service

const cacheUser = "user"
const cacheIP = "ip"
const cacheReqRule = "reqrule"
const cacheRespRule = "resprule"

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

func SetupCacheConfig() error {
	userSrv := DBUser{}
	err := userSrv.Distribute()
	if err != nil {
		return err
	}

	ipSrv := IP{}
	err = ipSrv.Distribute()
	if err != nil {
		return err
	}

	ruleSrv := Rule{}
	err = ruleSrv.Distribute()
	if err != nil {
		return err
	}

	return nil
}
