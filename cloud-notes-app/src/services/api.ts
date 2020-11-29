import { request } from 'alita';
import { sendWSPush } from '@/utils/websocket'

export async function query(): Promise<any> {
  return request('/api/hello');
}

export async function queryList(data: any): Promise<any> {
  return request('/api/list', { data });
}

export async function queryLogin(msg: any): Promise<any> {
  sendWSPush('login', msg)
  return new Promise((resolve, reject) => {
    addEventListener('onmessageWS', (e) => {
      console.log(e.detail.data)
      resolve(e.detail.data)
    })
  }).then((data) => {
    window.cloud = data
  })
}
