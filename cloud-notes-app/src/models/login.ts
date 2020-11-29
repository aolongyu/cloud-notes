import { Reducer } from 'alita';
import { queryLogin } from '@/services/api';
import { Effect } from '@/models/connect';
import { sendWSPush } from '@/utils/websocket';

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
      const data = yield call(queryLogin, payload);
      console.log(data)
      
      yield put({
        type: 'save',
        payload: { name: 'test' },
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
