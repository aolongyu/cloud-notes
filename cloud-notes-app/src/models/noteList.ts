import { Reducer } from 'redux';
import { query } from '@/services/api';
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
      const data = yield call(query, payload);
      console.log(data)
      yield put({
        type: 'save',
        payload: { name: data.text },
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
