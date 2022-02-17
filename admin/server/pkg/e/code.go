package e

const (
	SUCCESS       = 0
	ERROR         = 500
	InvalidParams = 400

	UserAddFailed     = 10000
	UserGetFailed     = 10001
	UserUpdateFailed  = 10002
	UserDeleteFailed  = 10003
	EmailAddFailed    = 20000
	EmailGetFailed    = 20001
	EmailUpdateFailed = 20002
	LdapAddFailed     = 30000
	LdapGetFailed     = 30001
	LdapUpdateFailed  = 30002
	TxsmsAddFailed    = 40000
	TxsmsGetFailed    = 40001
	TxsmsUpdateFailed = 40002
	MsgAddFailed      = 50000
	MsgGetFailed      = 50001
	MsgUpdateFailed   = 50002
	MsgDeleteFailed   = 50003

	DBAddFailed    = 60000
	DBGetFailed    = 60001
	DBUpdateFailed = 60002
	DBDeleteFailed = 60003

	DBUserAddFailed        = 70000
	DBUserGetFailed        = 70001
	DBUserUpdateFailed     = 70002
	DBUserDeleteFailed     = 70003
	DBUserDistributeFailed = 70004

	IPAddFailed        = 80000
	IPGetFailed        = 80001
	IPUpdateFailed     = 80002
	IPDeleteFailed     = 80003
	IPDistributeFailed = 80004

	RuleAddFailed    = 90000
	RuleGetFailed    = 90001
	RuleUpdateFailed = 90002
	RuleDeleteFailed = 90003

	AccessEventGetFailed = 11001
)
