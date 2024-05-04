import request from '@/utils/request'
import axios from 'axios';  

export function chat(data) {
  return axios({
    url: '/app/chat',
    method: 'post',
    headers: {"Content-Type":"application/json"},
    data:JSON.stringify(data),
  })
}

export function getChatStream(data) {
  return fetch('/app/chat',
    {
      method:'post',
      headers: {"Content-Type":"application/json"},
      body:JSON.stringify(data)
    });
}

export function chatmessages() {
  return axios({
    url: 'app/model',
    method: 'get',
  })
}