import React, { FC, useEffect } from 'react';
import { NoteListModelState, ConnectProps, connect } from 'alita';
import NoteBox from '@/components/noteBox/index'
import NoMore from '@/components/noMore/index'
import styles from './index.less';

interface PageProps extends ConnectProps {
  noteList: NoteListModelState;
}

const NoteListPage: FC<PageProps> = ({ noteList, dispatch }) => {
  // 这里发起了初始化请求
  // useEffect(() => {
  //   createSocket()

  //   return () => {
  //   };
  // }, []);
  dispatch!({
    type: 'noteList/query',
  });
  const { data } = noteList;

  const test = { Id: '1232id', Name: 'nama', Introduction: 'intorintorintorintorintorintorintorintorintorintorintor', Text: 'text', ThumbsUp: 'up' }
  return (
    <div className={styles.container}>
      <NoteBox Id={test.Id} Name={test.Name} Introduction={test.Introduction} Text={test.Text} ThumbsUp={test.ThumbsUp} />
      <NoteBox Id={test.Id} Name={test.Name} Introduction={test.Introduction} Text={test.Text} ThumbsUp={test.ThumbsUp} />
      <NoteBox Id={test.Id} Name={test.Name} Introduction={test.Introduction} Text={test.Text} ThumbsUp={test.ThumbsUp} />
      <NoteBox Id={test.Id} Name={test.Name} Introduction={test.Introduction} Text={test.Text} ThumbsUp={test.ThumbsUp} />

      --------------------上面是假数据---------------
      {
        // data.map((item: any) => <NoteBox Id={item.Id} Name={item.Name} Introduction={item.Introduction} Text={item.Text} ThumbsUp={item.ThumbsUp} />)
      }
      <NoMore text='没有更多了' />
    </div>
  );
};

export default connect(({ noteList }: { noteList: NoteListModelState; }) => ({ noteList }))(NoteListPage);
