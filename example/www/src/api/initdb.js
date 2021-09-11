import request from '@/utils/request'

// 初始化用户数据库
export function initDB(data) {
  return request({
    url: '/init/initdb',
    method: 'post',
    data
  })
}

//  初始化用户数据库
export function checkDB() {
  return request({
    url: '/init/checkdb',
    method: 'get'
  })
}
