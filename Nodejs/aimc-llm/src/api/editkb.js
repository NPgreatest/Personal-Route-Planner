import request from '@/utils/request'
import { obj2Url } from '@/utils'

export function editkb(data) {
  return request({
    url: 'knowledgebase/'+ data._id,
    method: 'put',
    data: data
  })
}


