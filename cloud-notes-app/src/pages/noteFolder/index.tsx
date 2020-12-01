import React, { FC, useEffect } from 'react';
import { NoteFolderModelState, ConnectProps, connect } from 'alita';
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
  // 注意，上面这里写空数组，表示初始化，如果需要监听某个字段变化再发起请求，可以在这里写明
  const { name } = noteFolder;
  return <div className={styles.center}>Hello {name}</div>;
};

export default connect(({ noteFolder }:{ noteFolder: NoteFolderModelState; }) => ({ noteFolder }))(NoteFolderPage);
