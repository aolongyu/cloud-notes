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

export async function queryNoteDetails(msg: object): Promise<any> {
  return requestCloud('ViewNote', msg)
}

export async function queryNoteBookList(msg: object): Promise<any> {
  return requestCloud('getNList', msg)
}

export async function queryCreateNote(msg: object): Promise<any> {
  return requestCloud('CrNoBook', msg)
}

export async function queryUser(msg: object): Promise<any> {
  return requestCloud('finduser', msg)
}

export async function queryCloseuser(msg: object): Promise<any> {
  return requestCloud('closeuser', msg)
}

export async function queryCrNote(msg: object): Promise<any> {
  return requestCloud('CrNote', msg)
}

export async function queryCrNoBook(msg: object): Promise<any> {
  return requestCloud('CrNoBook', msg)
}

export async function queryNoteList(msg: object): Promise<any> {
  return requestCloud('FindNote', msg)
}

// export async function queryNoteBookList(msg: object): Promise<any> {
//   return requestCloud('ViewNote', msg)
// }

// export async function queryNoteBookList(msg: object): Promise<any> {
//   return requestCloud('CrNoBook', msg)
// }

// export async function queryNoteBookList(msg: object): Promise<any> {
//   return requestCloud('CrNoBook', msg)
// }

// export async function queryNoteBookList(msg: object): Promise<any> {
//   return requestCloud('CrNoBook', msg)
// }

// export async function queryNoteBookList(msg: object): Promise<any> {
//   return requestCloud('CrNoBook', msg)
// }

// export async function queryNoteBookList(msg: object): Promise<any> {
//   return requestCloud('CrNoBook', msg)
// }