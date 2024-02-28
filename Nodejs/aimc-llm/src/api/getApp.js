import request from '@/utils/request'

export function getApp(params) {
    return request({
      url: `/app/${params._id}`,
      method: 'get',
      data:params
    })
  }