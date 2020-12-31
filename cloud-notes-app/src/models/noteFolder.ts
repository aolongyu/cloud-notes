import { Reducer } from 'alita';
import { queryFolder } from '@/services/api';
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
      yield call(queryFolder, payload);
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

export default NoteFolderModel;
