import React, { FC, useEffect } from 'react';
import { EditUserInfoModelState, ConnectProps, connect, history } from 'alita';
import { Toast } from 'antd-mobile'
import styles from './index.less';

interface PageProps extends ConnectProps {
  editUserInfo: EditUserInfoModelState;
}

const EditUserInfoPage: FC<PageProps> = ({ editUserInfo, dispatch, location }) => {

  const { Uid, Name, Password } = location.query
  // 这里发起了初始化请求
  useEffect(() => {
    dispatch!({
      type: 'editUserInfo/query',
    });
    return () => {
      // 这里写一些需要消除副作用的代码
      // 如: 声明周期中写在 componentWillUnmount
    };
  }, []);

  const { data } = editUserInfo

  if (String(data && data.Status) === '1') {
    Toast.success('修改成功，请从新登录', 1)
    setTimeout(() => {
      history.replace('login')
    }, 1000);
  } else if (String(data && data.Status) === '0') {
    Toast.fail('修改失败', 1)
  }

  const handleClick = () => {
    const Password = document.getElementById('registPassword1').value
    const Password2 = document.getElementById('registPassword2').value
    if (Password === Password2) {
      dispatch!({
        type: 'editUserInfo/queryUpdateUserInfo',
        payload: {
          uid: Number(Uid),
          pwd: Password
        }
      })
    } else {
      Toast.fail('两次密码输入不相同', 1)
    }
  }

  const { name } = editUserInfo;
  return (
    <div className={styles.container}>
      <div className={styles.titleText}>修改您的信息</div>
      <div className={styles.registBox}>
        <input id='registPassword1' type="password" className={styles.registPassword} placeholder='请输入密码' />
        <input id='registPassword2' type="password" className={styles.registPassword} placeholder='请再次输入密码' />
        <input id='registSubmit' className={styles.registSubmit} onClick={handleClick} type="button" value="提交" />
      </div>
    </div>
  )
};

export default connect(({ editUserInfo }: { editUserInfo: EditUserInfoModelState; }) => ({ editUserInfo }))(EditUserInfoPage);
