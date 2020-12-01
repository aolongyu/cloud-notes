import { Reducer } from 'redux';
import { query } from '@/services/api';
import { Effect } from '@/models/connect';

export interface NoteDetailsModelState {
  name: string;
}

export interface NoteDetailsModelType {
  namespace: 'noteDetails';
  state: NoteDetailsModelState;
  effects: {
    query: Effect;
  };
  reducers: {
    save: Reducer<NoteDetailsModelState>;
  };
}

const NoteDetailsModel: NoteDetailsModelType = {
  namespace: 'noteDetails',

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

export default NoteDetailsModel;
