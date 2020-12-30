import React, { FC, useEffect, useState } from 'react';
import { NoteFolderModelState, ConnectProps, connect, router } from 'alita';
import { Card, WingBlank, WhiteSpace, Icon, Modal, List, Button } from 'antd-mobile';
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

  const [visible, setVisible] = useState(false)
  const [Note_id, setNoteId] = useState()

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
  // console.log(data)
  // const data = [{ Id: '111', Name: 'NameTest', Introduction: 'Introduction', ThumbsUp: 'ThumbsUp' }]

  const handleSubmit = () => {
    const Note_name = document.getElementById('input0').value
    const Note_introduction = document.getElementById('input1').value
    dispatch!({
      type: 'noteFolder/queryUpdateNoteBook',
      payload: {
        Id: Note_id,
        notebook_name: Note_name,
        introduction: Note_introduction,
        notebook_type: 0
      }
    });
  }

  return (
    <div className={styles.container}>
      {/* <SearchBar placeholder="查找笔记本" maxLength={15} onCancel={(val) => { handleSearch(val) }} cancelText="查找" /> */}
      {
        data && data.map(item => (
          <div key={item.Id} className={styles.card}>
            <WingBlank size="lg">
              <WhiteSpace size="lg" />
              <Card>
                <Card.Header
                  title={item.Name}
                  thumb={<FolderOpenTwoTone />}
                  extra={<div><span>{item.ThumbsUp}</span>&nbsp;&nbsp;&nbsp;&nbsp;<Icon onClick={() => { setVisible(true); setNoteId(item.Id) }} type="ellipsis" /></div>}
                />
                <Card.Body onClick={() => { handle(item.Id) }}>
                  <div>{item.Introduction}</div>
                </Card.Body>
                {/* <Card.Footer content="footer content" extra={<div>extra footer content</div>} /> */}
              </Card>
            </WingBlank>
          </div>
        ))
      }
      <Modal
        popup
        visible={visible}
        onClose={() => { setVisible(false) }}
        animationType="slide-down"
        closable
      >
        <List renderHeader={() => <div>修改笔记本信息</div>} className="popup-list">
          {['笔记名称', '笔记笔记说明'].map((i, index) => (
            <List.Item key={index}>{i} <input id={`input${index}`} className={styles.input} type="text" /></List.Item>
          ))}
          <List.Item>
            <Button type="primary" onClick={handleSubmit}>创建笔记</Button>
          </List.Item>
        </List>
      </Modal>
      <NoMore />
    </div>
  );
};

export default connect(({ noteFolder }: { noteFolder: NoteFolderModelState; }) => ({ noteFolder }))(NoteFolderPage);
