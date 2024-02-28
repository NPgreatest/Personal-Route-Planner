import request from '@/utils/request'

export function openApp(data) {
  return request({
    url: `/issue/${data._id}`,
    method: 'post',
    data: data
  })
}