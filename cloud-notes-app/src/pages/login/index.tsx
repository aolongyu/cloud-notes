import React, { FC, useEffect } from 'react';
import { LoginModelState, ConnectProps, connect, router } from 'alita';
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
  const { name } = login;

  console.log(name)

  // return <div className={styles.center}>Hello {name}</div>;

  const handleClick = () => {
    const Name = document.getElementById('name').value
    const Password = document.getElementById('password').value
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
        <input id='name' className={styles.name} type="text" placeholder='请输入账号' /><br />
        <input id='password' className={styles.password} type="password" placeholder='请输入密码' /><br />
        <input className={styles.submitBtn} onClick={handleClick} type="button" value="登录" />
        <a onClick={() => {router.push('regist')}} className={styles.registText}>没账号？去注册</a>
      </div>
      <div className={styles.footer}>
        <span>Copyright ©2020 fjut 福建省福州市闽侯县学院路33号</span>
      </div>
    </div>
  );
};

export default connect(({ login }: { login: LoginModelState; }) => ({ login }))(LoginPage);
