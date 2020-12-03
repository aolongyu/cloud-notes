import React, { FC, useEffect } from 'react';
import { NoteFolderModelState, ConnectProps, connect } from 'alita';
import Card from '@/components/noteFolder/index'
import styles from './index.less';

interface PageProps extends ConnectProps {
  noteFolder: NoteFolderModelState;
}

const NoteFolderPage: FC<PageProps> = ({ noteFolder, dispatch }) => {
  // 这里发起了初始化请求
  useEffect(() => {
    dispatch!({
      type: 'noteFolder/query',
    });
    return () => {
      // 这里写一些需要消除副作用的代码
      // 如: 声明周期中写在 componentWillUnmount
    };
  }, []);

  const userInfo = JSON.parse(localStorage.getItem('userInfo'))
  const Name = userInfo.Name
  console.log(userInfo)

  dispatch!({
    type: 'noteFolder/query',
    payload: {
      Name
    }
  });
  const { data } = noteFolder;
  const test = {NoteName: 'noteName', NoteIntro: 'noteIntro', NoteModifyTime: 'noteModifyTime'}
  return (
    <div className={styles.container}>
      <Card cardName={test.NoteName} cardIntro={test.NoteIntro} cardModifyTime={test.NoteModifyTime} />
    </div>
  );
};

export default connect(({ noteFolder }: { noteFolder: NoteFolderModelState; }) => ({ noteFolder }))(NoteFolderPage);
