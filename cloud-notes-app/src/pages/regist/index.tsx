import React, { FC, useEffect } from 'react';
import { RegistModelState, ConnectProps, connect } from 'alita';
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
  return (
    <div className={styles.container}>
      <div className={styles.titleText}>regist</div>
      <div className={styles.registBox}>
        <input id='registName' className={styles.registName} placeholder='请输入用户名' type="text"/>
        <input id='registPassword1' className={styles.registPassword} placeholder='请输入密码' type="text"/>
        <input id='registPassword2' className={styles.registPassword} placeholder='请再次输入密码' type="text"/>
        <input id='registSubmit' className={styles.registSubmit} type="button" value="注册"/>
      </div>
    </div>
  );
};

export default connect(({ regist }: { regist: RegistModelState; }) => ({ regist }))(RegistPage);
