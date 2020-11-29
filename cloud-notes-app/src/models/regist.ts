import { Reducer } from 'redux';
import { query } from '@/services/api';
import { Effect } from '@/models/connect';

export interface RegistModelState {
  name: string;
}

export interface RegistModelType {
  namespace: 'regist';
  state: RegistModelState;
  effects: {
    query: Effect;
  };
  reducers: {
    save: Reducer<RegistModelState>;
  };
}

const RegistModel: RegistModelType = {
  namespace: 'regist',

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

export default RegistModel;
