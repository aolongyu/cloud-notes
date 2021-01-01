import { Reducer } from 'alita';
import { queryNoteBookList, queryUpdateNoteBook } from '@/services/api';
import { Effect } from '@/models/connect';
import { Toast } from 'antd-mobile';

export interface NoteFolderModelState {
  name: string;
}

export interface NoteFolderModelType {
  namespace: 'noteFolder';
  state: NoteFolderModelState;
  effects: {
    queryNoteBookList: Effect;
    queryUpdateNoteBook: Effect;
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
    *queryUpdateNoteBook({ payload }, { call, put }) {
      yield call(queryUpdateNoteBook, payload);
      const data = JSON.parse(JSON.parse(window.cloud))
      console.log('从服务端获取对象：', data)
      if (data.Status === '1') {
        Toast.success('修改笔记信息成功')
      } else {
        Toast.fail('修改笔记信息失败')
      }

      // yield put({
      //   type: 'save',
      //   payload: { data },
      // });
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
