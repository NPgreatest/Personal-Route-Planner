import request from '@/utils/request'
import { obj2Url } from '@/utils'

export function getHello() {
  return request({
    url: 'hello',
    method: 'get'
  })
}

export function upload(data) {
  return request({
    url: 'upload',
    method: 'post',
    data
  })
}

export function uploadKnowledgeBase(data) {
  console.log("shuchu",data)
  return request({
    url: 'uploadKnowledgeBase',
    method: 'post',
    data
  })
}
export function getKnowledgeList() {
  return request({
    url: 'getKnowledgeList',
    method: 'get'
  })
}

export function deleteKnowledge(data) {
  return request({
    url: 'deleteKnowledge',
    method: 'post',
    data
  })
}



