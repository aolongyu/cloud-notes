import { Effect, Reducer, history } from 'alita';
import { queryCrNote, queryCrNoBook } from '@/services/api';
export interface IndexModelState {
  name: string;
}

export interface IndexModelType {
  namespace: 'index';
  state: IndexModelState;
  effects: {
    queryCrNote: Effect;
    queryCrNoBook: Effect;
  };
  reducers: {
    save: Reducer<IndexModelState>;
  };
}

const IndexModel: IndexModelType = {
  namespace: 'index',

  state: {
    name: '',
  },

  effects: {
    *queryCrNote({ payload }, { call, put }) {
      yield call(queryCrNote, payload);
      const data = JSON.parse(JSON.parse(window.cloud))
      console.log('从服务端获取对象：', data)
      yield put({
        type: 'save',
        payload: { data },
      });
    },
    *queryCrNoBook({ payload }, { call, put }) {
      yield call(queryCrNoBook, payload);
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

export default IndexModel;
