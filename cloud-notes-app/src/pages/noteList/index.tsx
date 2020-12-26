import React, { FC, useEffect } from 'react';
import { NoteListModelState, ConnectProps, connect, router } from 'alita';
import NoteBox from '@/components/noteBox/index'
import NoMore from '@/components/noMore/index'
import { createSocket } from '@/utils/websocket'
import styles from './index.less';

interface PageProps extends ConnectProps {
  noteList: NoteListModelState;
}

const NoteListPage: FC<PageProps> = ({ noteList, dispatch, location }) => {

  const { NoteBookId, Name } = location.query

  // 这里发起了初始化请求
  useEffect(() => {
    createSocket()
    dispatch!({
      type: 'noteList/queryNoteList',
      payload: {
        Uid: NoteBookId
      }
    });
    return () => {
    };
  }, []);

  const { data } = noteList;

  const click = (NoteId: any) => {
    router.push({
      pathname: '/noteDetails',
      query: {
        NoteId,
        Name
      },
    });
  }

  // const data = [{ Id: '1232id', Name: 'nama', Introduction: 'intorintorintorintorintorintorintorintorintorintorintor', Text: 'text', ThumbsUp: 'up' }]
  return (
    <div className={styles.container}>
      {
        data && data.map((item: any) => <NoteBox key={item.Id} {...item} click={click} />)
      }
      <NoMore text='没有更多了' />
    </div>
  );
};

export default connect(({ noteList }: { noteList: NoteListModelState; }) => ({ noteList }))(NoteListPage);
