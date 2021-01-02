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
  return requestCloud('fiusn', msg)
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
  return requestCloud('FiNoByid', msg)
}

export async function queryUpdateNote(msg: object): Promise<any> {
  return requestCloud('editnote', msg)
}

export async function queryAddToBook(msg: object): Promise<any> {
  return requestCloud('AddToBook', msg)
}

export async function queryDeleteNote(msg: object): Promise<any> {
  return requestCloud('delnote', msg)
}

export async function queryUpdateNoteBook(msg: object): Promise<any> {
  return requestCloud('edNobook', msg)
}

export async function queryAllNote(msg: object): Promise<any> {
  return requestCloud('getallnot', msg)
}

export async function queryZan(msg: object): Promise<any> {
  return requestCloud('thumnote', msg)
}

export async function querySC(msg: object): Promise<any> {
  return requestCloud('conoadd', msg)
}

export async function queryJBNote(msg: object): Promise<any> {
  return requestCloud('renote', msg)
}

export async function queryUpdateUserInfo(msg: object): Promise<any> {
  return requestCloud('updpwd', msg)
}

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