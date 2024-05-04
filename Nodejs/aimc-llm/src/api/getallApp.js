import request from '@/utils/request'

export function getallApp() {
    return request({
      url: 'app',
      method: 'get'
    })
  }