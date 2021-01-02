import React, { FC, useEffect, useState } from 'react';
import { NoteListModelState, ConnectProps, connect, router, setPageNavBar } from 'alita';
import { Modal, List, Button, WhiteSpace, WingBlank, Icon, Toast, Popover } from 'antd-mobile';
import NoteBox from '@/components/noteBox/index'
import NoMore from '@/components/noMore/index'
import { createSocket } from '@/utils/websocket'
import styles from './index.less';

interface PageProps extends ConnectProps {
  noteList: NoteListModelState;
}

const NoteListPage: FC<PageProps> = ({ noteList, dispatch, location }) => {

  const { NoteBookId, Name, Uid } = location.query

  const [visiblePop, setVisiblePop] = useState(false)
  const [nid, setNid] = useState(null)

  // 这里发起了初始化请求
  useEffect(() => {
    createSocket()
    dispatch!({
      type: 'noteList/queryNoteList',
      payload: {
        Id: Number(NoteBookId)
      }
    });
    setPageNavBar({
      pagePath: location.pathname,
      navBar: {
        onLeftClick: () => {
          router.goBack()
        }
      },
    });
    return () => {
    };
  }, []);

  const { data } = noteList;

  const click = (index: number) => {
    console.log(data[index])
    router.push({
      pathname: '/noteDetails',
      query: { ...data[index], author: Name }
    })
  }

  const handleSubmit = () => {
    const bid = document.getElementById('inputbid').value
    dispatch!({
      type: 'noteList/queryAddToBook',
      payload: {
        nid: Number(nid),
        bid: Number(bid)
      }
    });
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
    }
  };
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
        <List renderHeader={() => <div>输入要加入的笔记本</div>} className="popup-list">
          {['笔记名称'].map((i, index) => (
            <List.Item key={index}>{i} <input id='inputbid' className={styles.input} type="text" /></List.Item>
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
