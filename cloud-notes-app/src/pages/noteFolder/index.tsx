import React, { FC, useEffect } from 'react';
import { NoteFolderModelState, ConnectProps, connect } from 'alita';
import { SearchBar, Card, WingBlank, WhiteSpace } from 'antd-mobile';
import { FolderOpenTwoTone } from '@ant-design/icons'
import NoMore from '@/components/noMore/index'
import { createSocket } from '@/utils/websocket'
import styles from './index.less';

interface PageProps extends ConnectProps {
  noteFolder: NoteFolderModelState;
}

const NoteFolderPage: FC<PageProps> = ({ noteFolder, dispatch }) => {

  // const userInfo = JSON.parse(localStorage.getItem('userInfo'))
  // const Name = userInfo && userInfo.Name && 'cdw'
  // console.log(userInfo)

  useEffect(() => {
    createSocket()
    return () => {
    };
  }, []);
  dispatch!({
    type: 'noteFolder/query',
    payload: {
      Name: 'cdw'
    }
  });
  const { data } = noteFolder;

  const handleSearch = (msg: string) => {
    console.log(msg)
  }

  const test = [{ Id: 'Id', Name: 'NameTest', Introduction: 'Introduction', ThumbsUp: 'ThumbsUp' }]

  return (
    <div className={styles.container}>
      <SearchBar placeholder="查找笔记本" maxLength={15} onCancel={(val) => { handleSearch(val) }} cancelText="查找" />
      {
        test.map(item => (
          <div className={styles.card}>
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
              <WhiteSpace size="lg" />
            </WingBlank>
          </div>
        ))
      }
      {/* <Card cardName={test.Name} cardIntro={test.Introduction} cardModifyTime={test.ThumbsUp} />
      <Card cardName={test.Name} cardIntro={test.Introduction} cardModifyTime={test.ThumbsUp} />
      <Card cardName={test.Name} cardIntro={test.Introduction} cardModifyTime={test.ThumbsUp} /> */}

      --------------------上面是假数据---------------
      {/* <Card cardName={data.Name} cardIntro={data.Introduction} cardModifyTime={data.ThumbsUp} /> */}
      <NoMore />
    </div>
  );
};

export default connect(({ noteFolder }: { noteFolder: NoteFolderModelState; }) => ({ noteFolder }))(NoteFolderPage);
