import { Reducer, router } from 'alita';
import { Toast } from 'antd-mobile'
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
      yield call(queryLogin, payload);
      const data = window.cloud
      if(data !== '0') {
        Toast.success('登录成功', 1)
        setTimeout(() => {
          router.replace('/')
        }, 1000);
      } else {
        Toast.fail('登录失败', 1)
      }
      yield put({
        type: 'save',
        payload: { name: data },
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
