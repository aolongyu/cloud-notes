import { Reducer, history } from 'alita';
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
      const data = JSON.parse(JSON.parse(window.cloud))
      console.log('从服务端获取对象：', data)
      // 登录结果处理
      if(data.Status !== '0') { // data.Status为服务端返回来的登录状态
        Toast.success('登录成功', 1)
        // 缓存当前登录用户的登录信息
        localStorage.setItem('userInfo', JSON.stringify({...payload, Uid: data.Uid}))
      } else {
        Toast.fail('登录失败', 1)
      }
      yield put({
        type: 'save',
        payload: { name: data.Status },
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
