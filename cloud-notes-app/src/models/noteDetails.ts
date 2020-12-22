import { Reducer } from 'alita';
import { queryNoteDetails } from '@/services/api';
import { Effect } from '@/models/connect';

export interface NoteDetailsModelState {
  name: string;
}

export interface NoteDetailsModelType {
  namespace: 'noteDetails';
  state: NoteDetailsModelState;
  effects: {
    queryNoteDetails: Effect;
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
