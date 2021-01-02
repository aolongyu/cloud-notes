import { Reducer } from 'redux';
import { queryUpdateUserInfo } from '@/services/api';
import { Effect } from '@/models/connect';

export interface EditUserInfoModelState {
  name: string;
}

export interface EditUserInfoModelType {
  namespace: 'editUserInfo';
  state: EditUserInfoModelState;
  effects: {
    queryUpdateUserInfo: Effect;
  };
  reducers: {
    save: Reducer<EditUserInfoModelState>;
  };
}

const EditUserInfoModel: EditUserInfoModelType = {
  namespace: 'editUserInfo',

  state: {
    name: '',
  },

  effects: {
    *queryUpdateUserInfo({ payload }, { call, put }) {
      yield call(queryUpdateUserInfo, payload);
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

export default EditUserInfoModel;
