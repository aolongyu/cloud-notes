import React, { FC, useEffect } from 'react';
import { NoteFolderModelState, ConnectProps, connect } from 'alita';
import Card from '@/components/noteFolder/index'
import NoMore from '@/components/noMore/index'
import styles from './index.less';

interface PageProps extends ConnectProps {
  noteFolder: NoteFolderModelState;
}

const NoteFolderPage: FC<PageProps> = ({ noteFolder, dispatch }) => {

  const userInfo = JSON.parse(localStorage.getItem('userInfo'))
  const Name = userInfo && userInfo.Name && 'cdw'
  console.log(userInfo)

  dispatch!({
    type: 'noteFolder/query',
    payload: {
      Name
    }
  });
  const { data } = noteFolder;
  const test = {Id: 'Id', Name: 'NameTest', Introduction: 'Introduction', ThumbsUp: 'ThumbsUp'}
  return (
    <div className={styles.container}>
      <Card cardName={test.Name} cardIntro={test.Introduction} cardModifyTime={test.ThumbsUp} />
      <Card cardName={test.Name} cardIntro={test.Introduction} cardModifyTime={test.ThumbsUp} />
      <Card cardName={test.Name} cardIntro={test.Introduction} cardModifyTime={test.ThumbsUp} />

      --------------------上面是假数据---------------
      {/* <Card cardName={data.Nasme} cardIntro={data.Introduction} cardModifyTime={data.ThumbsUp} /> */}
      <NoMore />
    </div>
  );
};

export default connect(({ noteFolder }: { noteFolder: NoteFolderModelState; }) => ({ noteFolder }))(NoteFolderPage);
