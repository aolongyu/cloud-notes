import { Reducer } from 'alita';
import { queryNoteList } from '@/services/api';
import { Effect } from '@/models/connect';

export interface NoteListModelState {
  name: string;
}

export interface NoteListModelType {
  namespace: 'noteList';
  state: NoteListModelState;
  effects: {
    query: Effect;
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
    *query({ payload }, { call, put }) {
      yield call(queryNoteList, payload);
      const data = JSON.parse(JSON.parse(window.cloud))
      console.log('从服务端获取对象：', data)
      yield put({
        type: 'save',
        payload: { data },
      });
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
