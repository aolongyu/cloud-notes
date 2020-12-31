import { Reducer, router } from 'alita';
import { Toast } from 'antd-mobile'
import { queryCreateNote } from '@/services/api';
import { Effect } from '@/models/connect';

export interface CreateNoteModelState {
  name: string;
}

export interface CreateNoteModelType {
  namespace: 'createNote';
  state: CreateNoteModelState;
  effects: {
    query: Effect;
  };
  reducers: {
    save: Reducer<CreateNoteModelState>;
  };
}

const CreateNoteModel: CreateNoteModelType = {
  namespace: 'createNote',

  state: {
    name: '',
  },

  effects: {
    *query({ payload }, { call, put }) {
      yield call(queryCreateNote, payload);
      const data = JSON.parse(JSON.parse(window.cloud))
      console.log('从服务端获取对象：', data)

      if(data.Status) {
        Toast.success(`创建笔记本: ${payload.NoteBookName}成功`, 1)
        router.push('/noteList')
      } else {
        Toast.fail(`创建笔记本: ${payload.NoteBookName}失败`, 1)
      }

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

export default CreateNoteModel;
