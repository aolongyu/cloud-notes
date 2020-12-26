import { Reducer } from 'alita';
import { queryNoteBookList } from '@/services/api';
import { Effect } from '@/models/connect';

export interface NoteFolderModelState {
  name: string;
}

export interface NoteFolderModelType {
  namespace: 'noteFolder';
  state: NoteFolderModelState;
  effects: {
    queryNoteBookList: Effect;
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
    *queryNoteBookList({ payload }, { call, put }) {
      yield call(queryNoteBookList, payload);
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
