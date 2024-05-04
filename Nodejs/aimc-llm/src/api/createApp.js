import request from '@/utils/request'

export function createApp(params) {
  return  request({
     url: 'app',
     method: 'post',
     data:params
   })
 }