import { Reducer } from 'alita';
import { queryNoteList, queryAddToBook, queryDeleteNote } from '@/services/api';
import { Effect } from '@/models/connect';
import { Toast } from 'antd-mobile';

export interface NoteListModelState {
  name: string;
}

export interface NoteListModelType {
  namespace: 'noteList';
  state: NoteListModelState;
  effects: {
    queryNoteList: Effect;
    queryAddToBook: Effect;
    queryDeleteNote: Effect;
  };
  reducers: {
    save: Reducer<NoteListModelState>;
  };
}

const NoteListModel: NoteListModelType = {
  namespace: 'noteList',

  state: {
    name: '',
  },

  effects: {
    *queryNoteList({ payload }, { call, put }) {
      yield call(queryNoteList, payload);
      const data = JSON.parse(JSON.parse(window.cloud))
      console.log('从服务端获取对象：', data)
      yield put({
        type: 'save',
        payload: { data },
      });
    },
    *queryAddToBook({ payload }, { call, put }) {
      yield call(queryAddToBook, payload);
      const data = JSON.parse(JSON.parse(window.cloud))
      console.log('从服务端获取对象：', data)
      if(data.Status === '1') {
        Toast.success('移动成功')
      }
      // yield put({
      //   type: 'save',
      //   payload: { data },
      // });
    },
    *queryDeleteNote({ payload }, { call, put }) {
      yield call(queryDeleteNote, payload);
      const data = JSON.parse(JSON.parse(window.cloud))
      console.log('从服务端获取对象：', data)
      if(data.Status === '1') {
        Toast.success('移动成功')
      }
      // yield put({
      //   type: 'save',
      //   payload: { data },
      // });
    },
    
  },
  reducers: {
    save(state, action) {
      return {
        ...state,
        ...action.payload,
      };
    },
  },
};

export default NoteListModel;
