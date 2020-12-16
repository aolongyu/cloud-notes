import React, { FC, useEffect } from 'react';
import { AdminModelState, ConnectProps, connect } from 'alita';
import { Tabs } from 'antd-mobile'
import FhBox from './components/fhBox/index'
import styles from './index.less';

interface PageProps extends ConnectProps {
  admin: AdminModelState;
}

const AdminPage: FC<PageProps> = ({ admin, dispatch }) => {
  // 这里发起了初始化请求
  // useEffect(() => {
  //   dispatch!({
  //     type: 'admin/query',
  //   });
  //   return () => {
  //     // 这里写一些需要消除副作用的代码
  //     // 如: 声明周期中写在 componentWillUnmount
  //   };
  // }, []);
  // 注意，上面这里写空数组，表示初始化，如果需要监听某个字段变化再发起请求，可以在这里写明
  // const { data } = admin;

  const test = [
    { Customer: 'Customer', LoginName: 'LoginName', Password: 'Password', UserStats: 'UserStats', ModifiedTime: 'ModifiedTime', CustomerLogincol: 'CustomerLogincol' },
    { Customer: 'Customer', LoginName: 'LoginName', Password: 'Password', UserStats: 'UserStats', ModifiedTime: 'ModifiedTime', CustomerLogincol: 'CustomerLogincol' },
    { Customer: 'Customer', LoginName: 'LoginName', Password: 'Password', UserStats: 'UserStats', ModifiedTime: 'ModifiedTime', CustomerLogincol: 'CustomerLogincol' },
    { Customer: 'Customer', LoginName: 'LoginName', Password: 'Password', UserStats: 'UserStats', ModifiedTime: 'ModifiedTime', CustomerLogincol: 'CustomerLogincol' },
    { Customer: 'Customer', LoginName: 'LoginName', Password: 'Password', UserStats: 'UserStats', ModifiedTime: 'ModifiedTime', CustomerLogincol: 'CustomerLogincol' }
  ]


  const tabs = [
    { title: '文章管理' },
    { title: '用户管理' },
    { title: '违规封号' }
  ];

  const handleClick = () => {
    const Sx = document.getElementById('getSx').value
    dispatch!({
      type: 'admin/query',
      payload: {
        PageNo: 1,
        PageSize: 1,
        I: 1,
        Sx
      }
    });
  }

  const handle = () => {
    console.log('123123123132')
  }

  return (<div className={styles.container}>
    <Tabs tabs={tabs} renderTabBar={props => <Tabs.DefaultTabBar {...props} page={tabs.length > 3 ? 3.5 : tabs.length} />}>
      <div className={styles.div1}>文章管理</div>
      <div className={styles.div2}></div>
      <div className={styles.div3}>
        <div className={styles.searchBox}>
          <input className={styles.div3Input} id="getSx" type="text" placeholder="查找用户" />
          <input className={styles.div3Btn} type="button" value="查找" onClick={handleClick} />
        </div>
        {
          test.map(item => <FhBox {...item} clickFunc={handle} />)
        }
      </div>
    </Tabs>
  </div>);
};

export default connect(({ admin }: { admin: AdminModelState; }) => ({ admin }))(AdminPage);
