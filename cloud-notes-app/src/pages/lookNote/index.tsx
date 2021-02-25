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

  const { Uid } = (JSON.parse(localStorage.getItem('userInfo')))

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
      type: 'lookNote/querySC',
      payload: {
        noteid: Number(Id), // 提供被收藏笔记本id
        bid: 2  // 放到迷人笔记本当中
      }
    });
  }

  const handleZan = () => {
    dispatch!({
      type: 'lookNote/queryZan',
      payload: {
        uid: Number(Uid), // 点赞用户本人id
        nid: Number(Id) // 提供被点赞的笔记id
      }
    });
  }

  const handleJB = () => {
    dispatch!({
      type: 'lookNote/queryJBNote',
      payload: {
        uid: Number(Uid), // 举报人id
        nid: Number(Id) // 被举报笔记id
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
