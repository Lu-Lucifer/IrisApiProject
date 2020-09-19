import request from '@/utils/request'

export function searchUser(name) {
  return request({
    url: '/admin/users',
    method: 'get',
    params: { name }
  })
}

export function transactionList(query) {
  return request({
    url: '/admin/transaction/list',
    method: 'get',
    params: query
  })
}
