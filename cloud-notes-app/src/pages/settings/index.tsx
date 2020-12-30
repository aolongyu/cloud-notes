import React, { FC, useEffect } from 'react';
import { setPageNavBar, connect, SettingsModelState, ConnectProps, history } from 'alita';
import { Modal } from 'antd-mobile';

import styles from './index.less';

const alert = Modal.alert;
interface PageProps extends ConnectProps {
  settings: SettingsModelState;
}

const SettingsPage: FC<PageProps> = ({ settings, dispatch, location }) => {

  useEffect(() => {
    dispatch!({
      type: 'settings/query',
    });
  }, []);
  // const { name } = settings;

  const userInfo = JSON.parse(localStorage.getItem('userInfo'))
  const Name = userInfo.Name
  const Uid = userInfo.Uid

  const handleExit = () => {
    alert('提示', '确认退出吗？', [
      { text: '取消', onPress: () => { } },
      { text: '确定', onPress: () => { history.replace('/login') } },
    ])
  }

  return (
    <div className={styles.container}>
      <div className={styles.line}><span className={styles.left}>用户ID: </span><span className={styles.right}>{Uid}</span></div>
      <hr/>
      <div className={styles.line}><span className={styles.left}>用户名: </span><span className={styles.right}>{Name}</span></div>
      <input className={styles.exitBtn} type="button" value="退出" onClick={handleExit} />
    </div>
  );
};

export default connect(({ settings }: { settings: SettingsModelState }) => ({ settings }))(
  SettingsPage,
);
