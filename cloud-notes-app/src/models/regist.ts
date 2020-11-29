import { Reducer, router } from 'alita';
import {Toast} from 'antd-mobile'
import { queryRegist } from '@/services/api';
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
      yield call(queryRegist, payload);
      const data = JSON.parse(JSON.parse(window.cloud))
      console.log('从服务端获取对象：', data)
      if(String(data.Status) !== '0') {
        Toast.success('注册成功，请登录', 1)
        setTimeout(() => {
          router.replace('login')
        }, 1000);
      } else {
        Toast.fail('注册失败', 1)
      }
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
