import request from '@/utils/request'
import { obj2Url } from '@/utils'

export function getFileListById(data){
  const para = obj2Url(data)
  console.log(para)
  return request({
    url: `file/kbid/${data.id}`,
    method: 'get'
  })
}

export function uploadFile(data, id) {
  const url = 'upload/'+ id
  return request({
    url: url,
    method: 'post',
    data
  })
}


// export function uploadFile(data) {
//   return request({
//     url: 'uploadFile',
//     method: 'post',
//     data
//   })
// }

