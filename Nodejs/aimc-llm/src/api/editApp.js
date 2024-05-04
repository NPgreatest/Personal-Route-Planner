import request from '@/utils/request'

export function editApp(params) {
  return  request({
     url: `/app/${params._id}`,
     method: 'put',
     data:params
   })
 }