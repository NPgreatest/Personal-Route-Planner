import request from '@/utils/request'
import { obj2Url } from '@/utils'

export function deletekb(id) {
  return request({
    url: 'knowledgebase/'+id,
    method: 'delete'
  })
}
