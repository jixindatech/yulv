import request from '@/utils/request'

export function getList(query, current = 1, size = 10) {
  return request({
    url: `/api/dbuser`,
    method: 'get',
    params: { ...query, page: current, size }
  })
}

export function add(data) {
  return request({
    url: `/api/dbuser`,
    method: 'post',
    data
  })
}

export function update(id, data) {
  return request({
    url: `/api/dbuser/${id}`,
    method: 'put',
    data
  })
}

export function getById(id) {
  return request({
    url: `/api/dbuser/${id}`,
    method: 'get'
  })
}

export function deleteById(id) {
  return request({
    url: `/api/dbuser/${id}`,
    method: 'delete'
  })
}

export function updateUserDB(id, data) {
  return request({
    url: `/api/dbuser/${id}/db`,
    method: 'put',
    data
  })
}

export function Distribute() {
  return request({
    url: `/api/dbuser/distribute`,
    method: 'post'
  })
}
