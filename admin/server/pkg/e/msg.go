package e

var MsgFlags = map[int]string{
	SUCCESS:       "OK",
	ERROR:         "fail",
	InvalidParams: "请求参数错误",

	UserAddFailed:    "添加用户失败",
	UserGetFailed:    "获取用户失败",
	UserUpdateFailed: "更新用户失败",
	UserDeleteFailed: "删除用户失败",

	EmailAddFailed:    "添加邮箱失败",
	EmailGetFailed:    "获取邮箱失败",
	EmailUpdateFailed: "更新邮箱失败",

	LdapAddFailed:    "添加LDAP失败",
	LdapGetFailed:    "获得LDAP失败",
	LdapUpdateFailed: "更新LDAP失败",

	TxsmsAddFailed:    "添加Txsms失败",
	TxsmsGetFailed:    "获得Txsms失败",
	TxsmsUpdateFailed: "更新Txsms失败",

	MsgAddFailed:    "添加msg失败",
	MsgGetFailed:    "获得msg失败",
	MsgUpdateFailed: "更新msg失败",
	MsgDeleteFailed: "删除msg失败",

	DBAddFailed:    "添加db失败",
	DBGetFailed:    "获得db失败",
	DBUpdateFailed: "更新db失败",
	DBDeleteFailed: "删除db失败",

	DBUserAddFailed:    "添加dbuser失败",
	DBUserGetFailed:    "获得dbuser失败",
	DBUserUpdateFailed: "更新dbuser失败",
	DBUserDeleteFailed: "删除dbuser失败",

	IPAddFailed:    "添加ip失败",
	IPGetFailed:    "获得ip失败",
	IPUpdateFailed: "更新ip失败",
	IPDeleteFailed: "删除ip失败",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
