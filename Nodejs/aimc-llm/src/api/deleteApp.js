import request from '@/utils/request'

export function deleteApp(params) {
  console.log(params)
  return request({
    url:`app/${params._id}`,
    method: 'delete',
    data: params
  })
}
