import request from '@/utils/request'

export function testApp(data) {
  return request({
    url: `/testapp/${data}`,
    method: 'post',
  })
}