import { Reducer } from 'redux';
import { query } from '@/services/api';
import { Effect } from '@/models/connect';

export interface LoginModelState {
  name: string;
}

export interface LoginModelType {
  namespace: 'login';
  state: LoginModelState;
  effects: {
    query: Effect;
  };
  reducers: {
    save: Reducer<LoginModelState>;
  };
}

const LoginModel: LoginModelType = {
  namespace: 'login',

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

export default LoginModel;
