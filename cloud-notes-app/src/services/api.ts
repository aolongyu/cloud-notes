import { request } from 'alita';
import requestCloud from '@/utils/requestCloud'

export async function query(): Promise<any> {
  return request('/api/hello');
}

export async function queryList(data: any): Promise<any> {
  return request('/api/list', { data });
}

export async function queryLogin(msg: object): Promise<any> {
  return requestCloud('login', msg)
}

export async function queryRegist(msg: object): Promise<any> {
  return requestCloud('regist', msg)
}
