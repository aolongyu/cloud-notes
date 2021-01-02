import { Reducer } from 'redux';
import { queryUser, queryCloseuser, queryRePwd } from '@/services/api';
import { Effect } from '@/models/connect';

export interface AdminModelState {
  name: string;
}

export interface AdminModelType {
  namespace: 'admin';
  state: AdminModelState;
  effects: {
    queryUser: Effect;
    queryCloseuser: Effect;
    queryRePwd: Effect;
  };
  reducers: {
    save: Reducer<AdminModelState>;
  };
}

const AdminModel: AdminModelType = {
  namespace: 'admin',

  state: {
    name: '',
  },

  effects: {
    *queryUser({ payload }, { call, put }) {
      yield call(queryUser, payload);
      const data = JSON.parse(JSON.parse(window.cloud))
      console.log('从服务端获取对象：', data)
      yield put({
        type: 'save',
        payload: { data },
      });
    },
    *queryCloseuser({ payload }, { call, put }) {
      yield call(queryCloseuser, payload);
      const data = JSON.parse(JSON.parse(window.cloud))
      console.log('从服务端获取对象：', data)
      yield put({
        type: 'save',
        payload: { data },
      });
    },
    *queryRePwd({ payload }, { call, put }) {
      yield call(queryRePwd, payload);
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

export default AdminModel;
