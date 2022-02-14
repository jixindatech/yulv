const ROLE_OPTIONS = [
  { value: 'admin', lable: 'admin' }
]

const LOGIN_OPTIONS = [
  { value: 'standard', label: 'standard' },
  { value: 'ldap', label: 'ldap' }
]

const LDAP_TYPE_OPTIONS = [
  { value: 'tcp', label: 'tcp' },
  { value: 'udp', label: 'udp' }
]

const USER_STATUS_OPTIONS = [
  { value: '未锁定', label: 1 },
  { value: '已锁定', label: 2 }
]

const RULE_TYPE_OPTIONS = [
  { label: '允许类型', value: 1 },
  { label: '拒绝类型', value: 2 },
  { label: '日志', value: 3 }
]

const RULE_TYPE_MAP = {
  1: '允许类型',
  2: '拒绝类型',
  3: '日志'
}

const IP_TYPE_OPTIONS = [
  { label: '允许类型', value: 1 },
  { label: '拒绝类型', value: 2 }
]

const IP_TYPE_MAP = {
  1: '允许类型',
  2: '拒绝类型'
}

const RULE_MATCH_OPTIONS = [
  { label: '字符串查找', value: 'str_find' },
  { label: '正则匹配', value: 're' }
]

const SQL_TYPE_OPTIONS = [
  { label: 'INSERT', value: 'insert' },
  { label: 'UPDATE', value: 'update' },
  { label: 'DELETE', value: 'delete' },
  { label: 'REPLACE', value: 'replace' },
  { label: 'SET', value: 'set' },
  { label: 'BEGIN', value: 'begin' },
  { label: 'COMMIT', value: 'commit' },
  { label: 'ROLLBACK', value: 'rollback' },
  { label: 'ADMIN', value: 'admin' },
  { label: 'SELECT', value: 'select' },
  { label: 'USE', value: 'use' },
  { label: 'START', value: 'start' },
  { label: 'TRANSACTION', value: 'transaction' },
  { label: 'SHOW', value: 'show' },
  { label: 'TRUNCATE', value: 'truncate' }
]
export {
  ROLE_OPTIONS,
  LOGIN_OPTIONS,
  LDAP_TYPE_OPTIONS,
  USER_STATUS_OPTIONS,
  RULE_TYPE_OPTIONS,
  RULE_TYPE_MAP,
  IP_TYPE_OPTIONS,
  IP_TYPE_MAP,
  RULE_MATCH_OPTIONS,
  SQL_TYPE_OPTIONS
}
