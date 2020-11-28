import React, { FC, useEffect } from 'react';
import { LoginModelState, ConnectProps, connect } from 'alita';
import { createSocket, sendWSPush } from '@/utils/websocket'
import styles from './index.less';

interface PageProps extends ConnectProps {
  login: LoginModelState;
}

const LoginPage: FC<PageProps> = ({ login, dispatch }) => {
  // 这里发起了初始化请求
  useEffect(() => {
    dispatch!({
      type: 'login/query',
    });
    return () => {
      // 这里写一些需要消除副作用的代码
      // 如: 声明周期中写在 componentWillUnmount
    };
  }, []);
  // 注意，上面这里写空数组，表示初始化，如果需要监听某个字段变化再发起请求，可以在这里写明
  const { name } = login;
  createSocket('ws://localhost:8999')
  
  // return <div className={styles.center}>Hello {name}</div>;

  const handleClick = () => {
    const Name = document.getElementById('name').value
    const Password = document.getElementById('password').value
    console.log(Name, Password)
    const sendMsg = {Name, Password}
    sendWSPush('login', JSON.stringify(sendMsg))
  }

  return (
    <div className="container">
      <input id='name' type="text" placeholder='请输入账号' value='aolyu@qq.com' /><br/>
      <input id='password' type="password" placeholder='请输入密码' value='123' /><br/>
      <input onClick={handleClick}  type="button" value="登录"/>
    </div>
  );
};

export default connect(({ login }: { login: LoginModelState; }) => ({ login }))(LoginPage);
