import React, { FC, useEffect, useState } from 'react';
import { IndexModelState, ConnectProps, connect, history, router } from 'alita';
import { Modal, List, Button, WhiteSpace, WingBlank, Icon, Toast, PickerView } from 'antd-mobile';
import { createSocket } from '@/utils/websocket'

import styles from './index.less';

const prompt = Modal.prompt;

interface PageProps extends ConnectProps {
  index: IndexModelState;
}

const IndexPage: FC<PageProps> = ({ index, dispatch, location }) => {
  // 这里发起了初始化请求
  useEffect(() => {
    createSocket()
  }, []);

  const [visible, setVisible] = useState(false)
  const [visible2, setVisible2] = useState(false)
  const [selectNumber, setSelectNumber] = useState('1')

  const { Uid, Name } = JSON.parse(localStorage.getItem('userInfo'))

  const handle1 = () => {
    setVisible(true)
  }

  const handle2 = () => {
    setVisible2(true)
  }

  const handle3 = () => {
    router.push({
      pathname: '/noteSquare',
      query: {
        Uid: Number(Uid),
        Name
      }
    })
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
        Uid: Number(Uid),
        NoteName,
        NoteIntroduction,
        NoteType: 0,
        Notebook_id: 1,
        NoteText
      }
    });
    setVisible(false)
  }

  const handleSubmit2 = () => {
    const Uid = JSON.parse(localStorage.getItem('userInfo')).Uid
    const NoteBookName = document.getElementById('input5').value
    const NoteBookIntroduction = document.getElementById('input7').value
    // const NoteBookType = document.getElementById('input6').value

    dispatch!({
      type: 'index/queryCrNoBook',
      payload: {
        Uid: Number(Uid),
        NoteBookName,
        NoteBookIntroduction,
        NoteBookType: selectNumber,
      }
    });

    setVisible2(false)
  }

  const { data } = index

  if (data && data.Status === '1') {
    data.Status = null
    Toast.success('创建成功')
    setTimeout(() => {
      history.push('/noteFolder')
    }, 1000)
  }

  const types = [
    {
      label: '日记',
      value: '1',
    },
    {
      label: '课程笔记',
      value: '2',
    },
    {
      label: '随手笔记',
      value: '3',
    },
    {
      label: '开心笔记',
      value: '4',
    },
    {
      label: '快乐笔记',
      value: '5',
    },
  ];

  const handleChange = (val: any) => {
    setSelectNumber(val[0])
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
          <List.Item>笔记本名称<input id={`input5`} className={styles.input} type="text" /></List.Item>
          <List.Item>笔记本说明<input id={`input7`} className={styles.input} type="text" /></List.Item>
          <List.Item>笔记本类型
            <PickerView
              onChange={handleChange}
              data={types}
              cascade={false}
            /></List.Item>
          <List.Item>
            <Button type="primary" onClick={handleSubmit2}>创建笔记本</Button>
          </List.Item>
        </List>
      </Modal>
      <input onClick={handle1} className={styles.input1} type="button" value="新建笔记" />
      <input onClick={handle2} className={styles.input2} type="button" value="新建笔记本" />
      <input onClick={handle3} className={styles.input3} type="button" value="笔记广场" />
    </div>
  );
};

export default connect(({ index }: { index: IndexModelState }) => ({ index }))(IndexPage);
