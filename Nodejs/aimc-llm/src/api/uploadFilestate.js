import request from '@/utils/request'
import { obj2Url } from '@/utils'


export function uploadFilestate(data,id) {
  console.log(data)
  const url = 'file/'+ id
  return request({
    url: url,
    method: 'put',
    data:data,
  })
}


