import React, { FC, useEffect } from 'react';
import { LookNoteModelState, ConnectProps, connect } from 'alita';
import { Card, WingBlank, WhiteSpace } from 'antd-mobile';
import SCImg from '@/assets/lookNote/shoucang.png'
import ZanImg from '@/assets/lookNote/zan.png'
import JBImg from '@/assets/lookNote/jb.png'
import styles from './index.less';

interface PageProps extends ConnectProps {
  lookNote: LookNoteModelState;
}

const LookNotePage: FC<PageProps> = ({ lookNote, dispatch, location }) => {

  const { Id, Introduction, Name, Text, ThumbsUp, author } = location.query

  const Uid = (JSON.parse(localStorage.getItem('userInfo'))).value
  // 这里发起了初始化请求
  // useEffect(() => {
  //   dispatch!({
  //     type: 'lookNote/queryNoteDetails',
  //   });
  //   return () => {
  //   };
  // }, []);

  // const { data } = lookNote;

  const handleSC = (bid: any) => {
    dispatch!({
      type: 'lookNote/queryZan',
      payload: {
        noteid: 9,
        bid: Number(bid)
      }
    });
  }

  const handleZan = () => {
    dispatch!({
      type: 'lookNote/querySC',
      payload: {
        nid: '',
        bid: ''
      }
    });
  }

  const handleJB = () => {
    dispatch!({
      type: 'lookNote/queryJBNote',
      payload: {
        uid: '',
        nid: ''
      }
    });
  }

  return (
    <>
      <WingBlank size="lg">
        <WhiteSpace size="lg" />
        <Card>
          <Card.Header
            title={Name}
            thumb="https://gw.alipayobjects.com/zos/rmsportal/MRhHctKOineMbKAZslML.jpg"
            extra={
              <div>
                <span>hot:{ThumbsUp}</span>
                <img onClick={handleSC} className={styles.SCImg} src={SCImg} />
                <img onClick={handleZan} className={styles.ZanImg} src={ZanImg} />
              </div>}
          />
          <Card.Body>
            <span className={styles.text}>{Text}</span>
          </Card.Body>
          <Card.Footer content={author} extra={<div><img onClick={handleJB} className={styles.JBImg} src={JBImg}/></div>} />
        </Card>
        <WhiteSpace size="lg" />
      </WingBlank>
    </>
  )
};

export default connect(({ lookNote }: { lookNote: LookNoteModelState; }) => ({ lookNote }))(LookNotePage);
