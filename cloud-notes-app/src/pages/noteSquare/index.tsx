import React, { FC, useEffect } from 'react';
import { NoteSquareModelState, ConnectProps, connect, router } from 'alita';
import { List } from 'antd-mobile';

import styles from './index.less';

const Item = List.Item;
const Brief = Item.Brief;

interface PageProps extends ConnectProps {
  noteSquare: NoteSquareModelState;
}

const NoteSquarePage: FC<PageProps> = ({ noteSquare, dispatch }) => {

  const { Name } = JSON.parse(localStorage.getItem('userInfo'))
  // 这里发起了初始化请求
  useEffect(() => {
    dispatch!({
      type: 'noteSquare/queryAllNote',
    });
    return () => {
    };
  }, []);
  const { data } = noteSquare;

  return (
    <div className={styles.container}>
      <List renderHeader={() => '笔记广场'} className="my-list">
        {
          data && data.map(item =>
            <Item
              arrow="horizontal"
              multipleLine
              onClick={() => {
                router.push({
                  pathname: 'lookNote',
                  query: { ...item, author: Name }
                })
              }}
              platform="android"
            >
              ({item.Thumbs}) {item.Name} ({item.Uname})<Brief>{item.Introduction}</Brief>
            </Item>
          )
        }
      </List>
    </div>
  )
};

export default connect(({ noteSquare }: { noteSquare: NoteSquareModelState; }) => ({ noteSquare }))(NoteSquarePage);
