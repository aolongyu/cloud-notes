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
  // 这里发起了初始化请求
  useEffect(() => {
    dispatch!({
      type: 'noteSquare/queryAllNote',
    });
    return () => {
      // 这里写一些需要消除副作用的代码
      // 如: 声明周期中写在 componentWillUnmount
    };
  }, []);
  // 注意，上面这里写空数组，表示初始化，如果需要监听某个字段变化再发起请求，可以在这里写明
  // const { data } = noteSquare;

  const data = [1,2,2,6,7,8]

  return (
    <div className={styles.container}>
      <List renderHeader={() => '笔记广场'} className="my-list">
        {
          data && data.map(item =>
            <Item
              arrow="horizontal"
              multipleLine
              onClick={() => { router.push({
                pathname: 'noteDetails',
                query: {
                  noteId: ''
                }
              }) }}
              platform="android"
            >
              ListItem （Android）<Brief>There may have water ripple effect of <br /> material if you set the click event.</Brief>
            </Item>
          )
        }
      </List>
    </div>
  )
};

export default connect(({ noteSquare }: { noteSquare: NoteSquareModelState; }) => ({ noteSquare }))(NoteSquarePage);
