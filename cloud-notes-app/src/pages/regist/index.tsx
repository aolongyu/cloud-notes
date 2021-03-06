import React, { FC, useEffect } from 'react';
import { RegistModelState, ConnectProps, connect, history, setPageNavBar, router } from 'alita';
import { Toast } from 'antd-mobile'
import { createSocket } from '@/utils/websocket'
import styles from './index.less';

interface PageProps extends ConnectProps {
  regist: RegistModelState;
}

const RegistPage: FC<PageProps> = ({ regist, dispatch, location }) => {


  useEffect(() => {
    createSocket()
    dispatch!({
      type: 'noteList/query',
    });
    setPageNavBar({
      pagePath: location.pathname,
      navBar: {
        onLeftClick: () => {
          router.push('/login');
        }
      },
    });
    return () => {
    };
  }, []);

  const { name } = regist

  console.log(name)

  if (String(name) === '1') {
    Toast.success('注册成功，请登录', 1)
    setTimeout(() => {
      history.replace('login')
    }, 1000);
  } else if(String(name) === '0') {
    Toast.fail('注册失败', 1)
  }

  const handleClick = () => {
    const Name = document.getElementById('registName').value
    const Password = document.getElementById('registPassword1').value
    const Password2 = document.getElementById('registPassword2').value
    if (Password === Password2 && judgeIsTrue(Password)) {  // 判断输入的数据是否满足校验，如空格等特殊字符
      dispatch!({
        type: 'regist/query',
        payload: {
          Name,
          Password
        }
      })
    } else {
      Toast.fail('两次密码输入不相同', 1)
    }
  }

  return (
    <div className={styles.container}>
      <div className={styles.titleText}>注册</div>
      <div className={styles.registBox}>
        <input id='registName' className={styles.registName} placeholder='请输入用户名' type="text" />
        <input id='registPassword1' type="password" className={styles.registPassword} placeholder='请输入密码' />
        <input id='registPassword2' type="password" className={styles.registPassword} placeholder='请再次输入密码' />
        <input id='registSubmit' className={styles.registSubmit} onClick={handleClick} type="button" value="注册" />
      </div>
    </div>
  );
};

export default connect(({ regist }: { regist: RegistModelState; }) => ({ regist }))(RegistPage);
