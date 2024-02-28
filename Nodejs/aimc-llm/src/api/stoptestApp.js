import request from '@/utils/request'

export function stoptestApp(data) {
  return request({
    url: `/stoptestapp/${data}`,
    method: 'post',
  })
}