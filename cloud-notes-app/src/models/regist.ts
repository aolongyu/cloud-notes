import { Reducer, history } from 'alita';
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

export default RegistModel;
