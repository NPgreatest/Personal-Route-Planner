import request from '@/utils/request'

export function submitPrompt(data) {
  return request({
    url: `/app/prompt/${data._id}`,
    method: 'put',
    data:data
  })
}