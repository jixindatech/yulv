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
  { label: '拒绝类型', value: 2 }
]

const RULE_TYPE_MAP = {
  1: '允许类型',
  2: '拒绝类型'
}

export {
  ROLE_OPTIONS,
  LOGIN_OPTIONS,
  LDAP_TYPE_OPTIONS,
  USER_STATUS_OPTIONS,
  RULE_TYPE_OPTIONS,
  RULE_TYPE_MAP
}
