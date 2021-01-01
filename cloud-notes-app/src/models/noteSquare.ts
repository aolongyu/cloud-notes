import { Reducer } from 'redux';
import { queryAllNote } from '@/services/api';
import { Effect } from '@/models/connect';

export interface NoteSquareModelState {
  name: string;
}

export interface NoteSquareModelType {
  namespace: 'noteSquare';
  state: NoteSquareModelState;
  effects: {
    queryAllNote: Effect;
  };
  reducers: {
    save: Reducer<NoteSquareModelState>;
  };
}

const NoteSquareModel: NoteSquareModelType = {
  namespace: 'noteSquare',

  state: {
    name: '',
  },

  effects: {
    *queryAllNote({ payload }, { call, put }) {
      yield call(queryAllNote, payload);
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

export default NoteSquareModel;
