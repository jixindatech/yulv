import request from '@/utils/request'

export function getList(query, current = 1, size = 10) {
  return request({
    url: `/api/db`,
    method: 'get',
    params: { ...query, page: current, size }
  })
}

export function add(data) {
  return request({
    url: `/api/db`,
    method: 'post',
    data
  })
}

export function update(id, data) {
  return request({
    url: `/api/db/${id}`,
    method: 'put',
    data
  })
}

export function getById(id) {
  return request({
    url: `/api/db/${id}`,
    method: 'get'
  })
}

export function deleteById(id) {
  return request({
    url: `/api/db/${id}`,
    method: 'delete'
  })
}
