import { Reducer } from 'redux';
import { queryZan, querySC, queryJBNote } from '@/services/api';
import { Toast } from 'antd-mobile'
import { Effect } from '@/models/connect';

export interface LookNoteModelState {
  name: string;
}

export interface LookNoteModelType {
  namespace: 'lookNote';
  state: LookNoteModelState;
  effects: {
    queryZan: Effect;
    querySC: Effect;
    queryJBNote: Effect;
  };
  reducers: {
    save: Reducer<LookNoteModelState>;
  };
}

const LookNoteModel: LookNoteModelType = {
  namespace: 'lookNote',

  state: {
    name: '',
  },

  effects: {
    *queryZan({ payload }, { call, put }) {
      yield call(queryZan, payload);
      const data = JSON.parse(JSON.parse(window.cloud))
      console.log('从服务端获取对象：', data)
      if(data.Status === '1') {
        Toast.success('操作成功')
      }
      yield put({
        type: 'save',
        payload: { ...data },
      });
    },
    *querySC({ payload }, { call, put }) {
      yield call(querySC, payload);
      const data = JSON.parse(JSON.parse(window.cloud))
      console.log('从服务端获取对象：', data)
      if(data.Status === '1') {
        Toast.success('收藏成功')
      }
      yield put({
        type: 'save',
        payload: { ...data },
      });
    },
    *queryJBNote({ payload }, { call, put }) {
      yield call(queryJBNote, payload);
      const data = JSON.parse(JSON.parse(window.cloud))
      console.log('从服务端获取对象：', data)
      if(data.Status === '1') {
        Toast.success('举报成功')
      }
      yield put({
        type: 'save',
        payload: { ...data },
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

export default LookNoteModel;
