import { request } from 'alita';
import { sendWSPush } from '@/utils/websocket'

export async function query(): Promise<any> {
  return request('/api/hello');
}

export async function queryList(data: any): Promise<any> {
  return request('/api/list', { data });
}

export async function queryLogin(msg: any): Promise<any> {
  console.log(msg)
  await sendWSPush('login', msg)
  return 
}