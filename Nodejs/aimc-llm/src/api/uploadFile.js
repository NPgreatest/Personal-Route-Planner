import request from '@/utils/request'
import { obj2Url } from '@/utils'


export function uploadFile(data, id) {
  const url = 'upload/'+ id
  return request({
    url: url,
    method: 'post',
    data:data,
  })
}

export function getFileListByKBId(id) {
  return request({
    url: 'file/kbid/'+id,
    method: 'get'
  })
}

export function downloadFile(id) {
  return request({
    url: 'download/'+id,
    method: 'get'
  })
}

