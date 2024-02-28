import request from '@/utils/request'

export function stopApp(id) {
  return  request({
     url:  `/stopapp/${id}`,
     method: 'post',

   })
 }