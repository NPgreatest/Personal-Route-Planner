import request from '@/utils/request'
import { obj2Url } from '@/utils'

export function deleteFile(id) {
  return request({
    url: 'remove/'+id,
    method: 'delete'
  })
}
