import React, { FC, useEffect } from 'react';
import { NoteListModelState, ConnectProps, connect } from 'alita';
import NoteBox from '@/components/noteBox/index'
import NoMore from '@/components/noMore/index'
import { createSocket } from '@/utils/websocket'
import styles from './index.less';

interface PageProps extends ConnectProps {
  noteList: NoteListModelState;
}

const NoteListPage: FC<PageProps> = ({ noteList, dispatch }) => {
  // 这里发起了初始化请求
  useEffect(() => {
    createSocket()
    dispatch!({
      type: 'noteList/query',
    });
    return () => {
    };
  }, []);

  // const { data } = noteList;

  const data = [{ Id: '1232id', Name: 'nama', Introduction: 'intorintorintorintorintorintorintorintorintorintorintor', Text: 'text', ThumbsUp: 'up' }]
  return (
    <div className={styles.container}>
      {
        data.map((item: any) => <NoteBox Id={item.Id} Name={item.Name} Introduction={item.Introduction} Text={item.Text} ThumbsUp={item.ThumbsUp} />)
      }
      <NoMore text='没有更多了' />
    </div>
  );
};

export default connect(({ noteList }: { noteList: NoteListModelState; }) => ({ noteList }))(NoteListPage);
