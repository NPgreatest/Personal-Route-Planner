import request from '@/utils/request'

export function getKnowledgeList() {
    return request({
      url: '/knowledgebase',
      method: 'get'
    })
  }