import request from '@/utils/request'

export function getRoles() {
  return request({
    url: '/admin/roles',
    method: 'get'
  })
}

export function getUsers() {
  return request({
    url: '/admin/users',
    method: 'get'
  })
}

export function addUser(data) {
  return request({
    url: '/admin/users',
    method: 'post',
    data
  })
}

export function updateUser(id, data) {
  return request({
    url: `/admin/users/${id}`,
    method: 'put',
    data
  })
}

export function deleteUser(id) {
  return request({
    url: `/admin/users/${id}`,
    method: 'delete'
  })
}
