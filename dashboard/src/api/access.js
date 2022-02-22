import request from '@/utils/request'

export function getList(query, current = 1, size = 10) {
  return request({
    url: `/api/event/access`,
    method: 'get',
    params: { ...query, page: current, size }
  })
}