import React, { FC, useEffect } from 'react';
import { LoginModelState, ConnectProps, connect } from 'alita';
import { createSocket, sendWSPush } from '@/utils/websocket'
import Logo from '@/assets/login/logo.png'
import styles from './index.less';

interface PageProps extends ConnectProps {
  login: LoginModelState;
}

const LoginPage: FC<PageProps> = ({ login, dispatch }) => {
  // 这里发起了初始化请求
  useEffect(() => {
    createSocket()
    
    
    return () => {
    };
  }, []);
  // 注意，上面这里写空数组，表示初始化，如果需要监听某个字段变化再发起请求，可以在这里写明
  const { name } = login;

  console.log(name)

  // return <div className={styles.center}>Hello {name}</div>;

  const handleClick = () => {
    const Name = document.getElementById('name').value
    const Password = document.getElementById('password').value
    // console.log(Name, Password)
    // const sendMsg = {Name, Password}
    // sendWSPush('login', JSON.stringify(sendMsg))
    dispatch!({
      type: 'login/query',
      payload: {
        Name,
        Password
      },
    });
  }

  return (
    <div className={styles.container}>
      <div className={styles.titleBox}>
        <img className={styles.titleLogo} src={Logo} alt="在线云笔记" />
        <span className={styles.titleText}>在线云笔记</span>
      </div>
      <div className={styles.inputBox}>
        <input id='name' type="text" placeholder='请输入账号' /><br />
        <input id='password' type="password" placeholder='请输入密码' /><br />
        <input onClick={handleClick} type="button" value="登录" />
      </div>
      <div className={styles.footer}>
        <span>Copyright ©2020 fjut 福建省福州市闽侯县学院路33号</span>
      </div>
    </div>
  );
};

export default connect(({ login }: { login: LoginModelState; }) => ({ login }))(LoginPage);
