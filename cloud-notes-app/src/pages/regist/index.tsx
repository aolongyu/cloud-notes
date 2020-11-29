import React, { FC, useEffect } from 'react';
import { RegistModelState, ConnectProps, connect } from 'alita';
import {Toast} from 'antd-mobile'
import styles from './index.less';

interface PageProps extends ConnectProps {
  regist: RegistModelState;
}

const RegistPage: FC<PageProps> = ({ regist, dispatch }) => {
  // 这里发起了初始化请求
  // useEffect(() => {
  //   dispatch!({
  //     type: 'regist/query',
  //   });
  //   return () => {
  //     // 这里写一些需要消除副作用的代码
  //     // 如: 声明周期中写在 componentWillUnmount
  //   };
  // }, []);
  // 注意，上面这里写空数组，表示初始化，如果需要监听某个字段变化再发起请求，可以在这里写明
  // const { name } = regist;

  const handleClick = () => {
    const Name = document.getElementById('registName').value
    const Password = document.getElementById('registPassword1').value
    const Password2 = document.getElementById('registPassword2').value
    if(Password === Password2) {
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
        <input id='registName' className={styles.registName} placeholder='请输入用户名' type="text"/>
        <input id='registPassword1' className={styles.registPassword} placeholder='请输入密码' type="text"/>
        <input id='registPassword2' className={styles.registPassword} placeholder='请再次输入密码' type="text"/>
        <input id='registSubmit' className={styles.registSubmit} onClick={handleClick} type="button" value="注册"/>
      </div>
    </div>
  );
};

export default connect(({ regist }: { regist: RegistModelState; }) => ({ regist }))(RegistPage);
