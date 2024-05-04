import request from '@/utils/request'
import { obj2Url } from '@/utils'


export function uploadKnowledgeBase(data) {
  return request({
    url: 'knowledgebase',
    method: 'post',
    data
  })
}

export function getKnowledgeBase() {
  return request({
    url: 'knowledgebase',
    method: 'get',
  })
}



