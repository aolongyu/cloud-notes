import React, { FC, useEffect } from 'react';
import { NoteFolderModelState, ConnectProps, connect, router } from 'alita';
import { SearchBar, Card, WingBlank, WhiteSpace } from 'antd-mobile';
import { FolderOpenTwoTone } from '@ant-design/icons'
import NoMore from '@/components/noMore/index'
import { createSocket } from '@/utils/websocket'
import styles from './index.less';

interface PageProps extends ConnectProps {
  noteFolder: NoteFolderModelState;
}

const NoteFolderPage: FC<PageProps> = ({ noteFolder, dispatch }) => {

  const userInfo = JSON.parse(localStorage.getItem('userInfo'))
  const Uid = userInfo && userInfo.Uid
  const Name = userInfo && userInfo.Name

  useEffect(() => {
    createSocket()
    dispatch!({
      type: 'noteFolder/queryNoteBookList',
      payload: {
        Uid
      }
    });
  }, []);

  const handleSearch = (msg: string) => {
    
  }

  const handle = (NoteBookId: any) => {
    router.push({
      pathname: '/noteList',
      query: {
        NoteBookId,
        Name,
        Uid
      },
    });
  }

  const { data } = noteFolder
  console.log(data)
  // const data = [{ Id: '111', Name: 'NameTest', Introduction: 'Introduction', ThumbsUp: 'ThumbsUp' }]

  return (
    <div className={styles.container}>
      {/* <SearchBar placeholder="查找笔记本" maxLength={15} onCancel={(val) => { handleSearch(val) }} cancelText="查找" /> */}
      {
        data && data.map(item => (
          <div key={item.Id} className={styles.card} onClick={() => { handle(item.Id) }}>
            <WingBlank size="lg">
              <WhiteSpace size="lg" />
              <Card>
                <Card.Header
                  title={item.Name}
                  thumb={<FolderOpenTwoTone />}
                  extra={<span>{item.ThumbsUp}</span>}
                />
                <Card.Body>
                  <div>{item.Introduction}</div>
                </Card.Body>
                {/* <Card.Footer content="footer content" extra={<div>extra footer content</div>} /> */}
              </Card>
            </WingBlank>
          </div>
        ))
      }
      <NoMore />
    </div>
  );
};

export default connect(({ noteFolder }: { noteFolder: NoteFolderModelState; }) => ({ noteFolder }))(NoteFolderPage);
