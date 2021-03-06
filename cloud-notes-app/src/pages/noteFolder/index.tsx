import React, { FC, useEffect, useState } from 'react';
import { NoteFolderModelState, ConnectProps, connect, router, history } from 'alita';
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
        Uid: Number(Uid)
      }
    });
  }, []);

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

  const handleSubmit = () => {
    // 笔记本名称
    const Note_name = document.getElementById('input0').value
    // 笔记本说明
    const Note_introduction = document.getElementById('input1').value
    // 向服务端发起请求
    dispatch!({
      type: 'noteFolder/queryUpdateNoteBook',
      payload: {
        Id: Note_id,
        notebook_name: Note_name,
        introduction: Note_introduction,
        notebook_type: 0
      }
    })
    // 页面跳转
    history.push('/')
  }

  const handleSC = (index: number) => {
    router.push({
      pathname: '/noteList',
      query: {
        NoteBookId: index,
        Name,
        Uid
      },
    });
  }

  return (
    <div className={styles.container}>
      {/* <SearchBar placeholder="查找笔记本" maxLength={15} onCancel={(val) => { handleSearch(val) }} cancelText="查找" /> */}
      <div className={styles.card} onClick={() => { handleSC(2) }}>
        <WingBlank size="lg">
          <WhiteSpace size="lg" />
          <div className={styles.sc}>
            <Card.Header
              title="我的收藏"
              thumb={<FolderOpenTwoTone />}
              extra={<Icon type="right" />}
            />
          </div>
        </WingBlank>
      </div>
      <div className={styles.card} onClick={() => { handleSC(1) }}>
        <WingBlank size="lg">
          <WhiteSpace size="lg" />
          <div className={styles.sc}>
            <Card.Header
              title="默认笔记本"
              thumb={<FolderOpenTwoTone />}
              extra={<Icon type="right" />}
            />
          </div>
        </WingBlank>
      </div>
      {
        data && data.map((item: any) => (
          <div key={item.Id} className={(Number(item.Id) !== 1) && (Number(item.Id) !== 2) ? styles.card : styles.none}>
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
          {['笔记本名称', '笔记本说明'].map((i, index) => (
            <List.Item key={index}>{i} <input id={`input${index}`} className={styles.input} type="text" /></List.Item>
          ))}
          <List.Item>
            <Button type="primary" onClick={handleSubmit}>修改笔记本</Button>
          </List.Item>
        </List>
      </Modal>
      <NoMore />
    </div>
  );
};

export default connect(({ noteFolder }: { noteFolder: NoteFolderModelState; }) => ({ noteFolder }))(NoteFolderPage);
