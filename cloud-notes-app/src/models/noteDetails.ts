import { Reducer } from 'alita';
import { queryNoteDetails, queryUpdateNote } from '@/services/api';
import { Effect } from '@/models/connect';
import { Toast } from 'antd-mobile';

export interface NoteDetailsModelState {
  name: string;
}

export interface NoteDetailsModelType {
  namespace: 'noteDetails';
  state: NoteDetailsModelState;
  effects: {
    queryNoteDetails: Effect;
    queryUpdateNote: Effect;
  };
  reducers: {
    save: Reducer<NoteDetailsModelState>;
  };
}

const NoteDetailsModel: NoteDetailsModelType = {
  namespace: 'noteDetails',

  state: {
    name: '',
  },

  effects: {
    *queryNoteDetails({ payload }, { call, put }) {
      yield call(queryNoteDetails, payload);
      const data = JSON.parse(JSON.parse(window.cloud))
      console.log('从服务端获取对象：', data)
      yield put({
        type: 'save',
        payload: { data },
      });
    },
    *queryUpdateNote({ payload }, { call, put }) {
      yield call(queryUpdateNote, payload);
      const data = JSON.parse(JSON.parse(window.cloud))
      if(data.Status === '1') {
        Toast.success('修改成功')
      }
      console.log('从服务端获取对象：', data)
      // yield put({
      //   type: 'save',
      //   payload: { data1 },
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

export default NoteDetailsModel;
