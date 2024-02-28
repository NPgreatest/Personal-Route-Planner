import request from '@/utils/request'

export function submitApp(id) {
  return  request({
     url:  `/issue/${id}`,
     method: 'get',
     
   })
 }