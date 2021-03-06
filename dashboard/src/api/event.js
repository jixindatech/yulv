import request from '@/utils/request'

export function getList(query, current = 1, size = 10) {
  return request({
    url: `/api/event/access`,
    method: 'get',
    params: { ...query, page: current, size }
  })
}

export function getInfo(query) {
  return request({
    url: `/api/event/info/access`,
    method: 'get',
    params: { ...query }
  })
}

export function getRuleList(query, current = 1, size = 10) {
  return request({
    url: `/api/event/rule`,
    method: 'get',
    params: { ...query, page: current, size }
  })
}

export function getRuleInfo(query) {
  return request({
    url: `/api/event/info/rule`,
    method: 'get',
    params: { ...query }
  })
}
