import React, { FC, useEffect, useState } from 'react';
import { IndexModelState, ConnectProps, connect, history } from 'alita';
import { Modal, List, Button, WhiteSpace, WingBlank, Icon, Toast } from 'antd-mobile';
import { createSocket } from '@/utils/websocket'

import styles from './index.less';

const prompt = Modal.prompt;

interface PageProps extends ConnectProps {
  index: IndexModelState;
}

const IndexPage: FC<PageProps> = ({ index, dispatch }) => {
  // 这里发起了初始化请求
  useEffect(() => {
    createSocket()
  }, []);

  const [visible, setVisible] = useState(false)
  const [visible2, setVisible2] = useState(false)

  const handle1 = () => {
    setVisible(true)
  }

  const handle2 = () => {
    setVisible2(true)
  }

  const handleSubmit1 = () => {
    const Uid = JSON.parse(localStorage.getItem('userInfo')).Uid
    const NoteName = document.getElementById('input0').value
    const NoteIntroduction = document.getElementById('input1').value
    // const Notebook_id = document.getElementById('input1').value
    const NoteText = '无'

    dispatch!({
      type: 'index/queryCrNote',
      payload: {
        Uid,
        NoteName,
        NoteIntroduction,
        NoteType: '0',
        Notebook_id: 10,
        NoteText
      }
    });

    setVisible(false)
  }

  const handleSubmit2 = () => {
    const Uid = JSON.parse(localStorage.getItem('userInfo')).Uid
    const NoteBookName = document.getElementById('input5').value
    const NoteBookIntroduction = document.getElementById('input7').value
    const NoteBookType = document.getElementById('input6').value

    dispatch!({
      type: 'index/queryCrNoBook',
      payload: {
        Uid,
        NoteBookName,
        NoteBookIntroduction,
        NoteBookType,
      }
    });

    setVisible2(false)
  }

  const { data } = index

  if(data && data.Status === '1') {
    data.Status = null
    Toast.success('创建成功')
    setTimeout(() => {
      history.push('/noteFolder')
    }, 1000)
  }

  return (
    <div className={styles.container}>
      <Modal
        popup
        visible={visible}
        onClose={() => { setVisible(false) }}
        animationType="slide-down"
        closable
      >
        <List renderHeader={() => <div>输入笔记基本信息</div>} className="popup-list">
          {['笔记名称', '笔记说明'].map((i, index) => (
            <List.Item key={index}>{i} <input id={`input${index}`} className={styles.input} type="text" /></List.Item>
          ))}
          <List.Item>
            <Button type="primary" onClick={handleSubmit1}>创建笔记</Button>
          </List.Item>
        </List>
      </Modal>
      <Modal
        popup
        visible={visible2}
        onClose={() => { setVisible2(false) }}
        animationType="slide-down"
        closable
      >
        <List renderHeader={() => <div>输入笔记本信息</div>} className="popup-list">
          {['笔记本名称', '笔记本类型', '笔记本说明'].map((i, index) => (
            <List.Item key={index}>{i} <input id={`input${index + 5}`} className={styles.input} type="text" /></List.Item>
          ))}
          <List.Item>
            <Button type="primary" onClick={handleSubmit2}>创建笔记本</Button>
          </List.Item>
        </List>
      </Modal>
      <input onClick={handle1} className={styles.input1} type="button" value="新建笔记" />
      <input onClick={handle2} className={styles.input2} type="button" value="新建笔记本" />
    </div>
  );
};

export default connect(({ index }: { index: IndexModelState }) => ({ index }))(IndexPage);
