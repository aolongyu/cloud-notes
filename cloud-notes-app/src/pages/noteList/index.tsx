import React, { FC, useEffect, useState } from 'react';
import { NoteListModelState, ConnectProps, connect, router, setPageNavBar, history } from 'alita';
import { Modal, List, Button, WhiteSpace, WingBlank, Icon, Toast, Popover, Radio } from 'antd-mobile';
import NoteBox from '@/components/noteBox/index'
import NoMore from '@/components/noMore/index'
import { createSocket } from '@/utils/websocket'
import styles from './index.less';

const RadioItem = Radio.RadioItem;

interface PageProps extends ConnectProps {
  noteList: NoteListModelState;
}

const NoteListPage: FC<PageProps> = ({ noteList, dispatch, location }) => {

  const { NoteBookId, Name, Uid } = location.query

  const [visiblePop, setVisiblePop] = useState(false)
  const [nid, setNid] = useState(null)
  const [checked, setChecked] = useState(13)

  // 这里发起了初始化请求
  useEffect(() => {
    createSocket()
    dispatch!({
      type: 'noteList/queryNoteList',
      payload: {
        Id: Number(NoteBookId)
      }
    });
  }, []);

  const { data } = noteList;

  let noteBookData = []
  const noteFolder = JSON.parse(localStorage.getItem('noteFolder'))
  console.log(noteFolder)
  noteFolder.forEach((item: any) => {
    noteBookData.push({value: item.Id, label: item.Name})
  })

  const click = (index: number) => {
    console.log(data[index])
    router.push({
      pathname: '/noteDetails',
      query: { ...data[index], author: Name }
    })
  }

  const handleSubmit = () => {
    // const bid = document.getElementById('inputbid').value
    dispatch!({
      type: 'noteList/queryAddToBook',
      payload: {
        obid: Number(NoteBookId),
        nid: Number(nid),
        bid: Number(checked)
      }
    });
    setTimeout(() => {
      history.replace('/noteFolder')
    }, 1000);
  }

  const onChange = (val: number) => {
    setChecked(val)
    console.log(val)
  }

  const onSelect = (e: any, mnid: any) => {
    console.log(e.props.dataSeed)
    if (e.props.dataSeed === '1') {
      setVisiblePop(true)
      setNid(mnid)
    } else {
      dispatch!({
        type: 'noteList/queryDeleteNote',
        payload: {
          Sid: Uid,
          Note_id: mnid
        }
      });
      setTimeout(() => {
        history.replace('/noteFolder')
      }, 1000);
    }
  }

  const prompt = Modal.prompt;
  // const data = [{ Id: '1232id', Name: 'nama', Introduction: 'intorintorintorintorintorintorintorintorintorintorintor', Text: 'text', ThumbsUp: 'up' }]
  return (
    <div className={styles.container}>
      {
        data && data.map((item: any, index: number) => <NoteBox key={item.Id} {...item} index={index} click={click} onSelect={onSelect} />)
      }
      <NoMore text='没有更多了' />
      <Modal
        popup
        visible={visiblePop}
        onClose={() => { setVisiblePop(false) }}
        animationType="slide-down"
        closable
      >
        <List renderHeader={() => <div>选择要加入的笔记本</div>} className="popup-list">
          {/* <List.Item>笔记本名称： <input id='inputbid' className={styles.input} type="text" /></List.Item> */}
          {noteBookData.map(i => (
            <RadioItem key={i.value} checked={checked === i.value} onChange={() => onChange(i.value)}>
              {i.label}
            </RadioItem>
          ))}
          <List.Item>
            <Button type="primary" onClick={handleSubmit}>移动笔记</Button>
          </List.Item>
        </List>
      </Modal>
    </div>
  );
};

export default connect(({ noteList }: { noteList: NoteListModelState; }) => ({ noteList }))(NoteListPage);
