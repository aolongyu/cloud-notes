import { Reducer } from 'redux';
import { query } from '@/services/api';
import { Effect } from '@/models/connect';

export interface NoteFolderModelState {
  name: string;
}

export interface NoteFolderModelType {
  namespace: 'noteFolder';
  state: NoteFolderModelState;
  effects: {
    query: Effect;
  };
  reducers: {
    save: Reducer<NoteFolderModelState>;
  };
}

const NoteFolderModel: NoteFolderModelType = {
  namespace: 'noteFolder',

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

export default NoteFolderModel;
