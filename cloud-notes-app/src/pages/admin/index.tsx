import React, { FC, useEffect } from 'react';
import { AdminModelState, ConnectProps, connect } from 'alita';
import { Card, Modal, WhiteSpace, Tabs, Button, SearchBar } from 'antd-mobile';
import { createSocket } from '@/utils/websocket'
import styles from './index.less';

const operation = Modal.operation;

interface PageProps extends ConnectProps {
  admin: AdminModelState;
}

const AdminPage: FC<PageProps> = ({ admin, dispatch }) => {
  // 这里发起了初始化请求
  useEffect(() => {
    createSocket()
    return () => {
    };
  }, []);
  // 注意，上面这里写空数组，表示初始化，如果需要监听某个字段变化再发起请求，可以在这里写明
  const { data } = admin;

  // const test = [
  //   { Customer_id: 'Customer1', LoginName: 'LoginName', Password: 'Password', UserStats: 'UserStats', ModifiedTime: 'ModifiedTime', CustomerLogincol: 'CustomerLogincol' },
  //   { Customer_id: 'Customer2', LoginName: 'LoginName', Password: 'Password', UserStats: 'UserStats', ModifiedTime: 'ModifiedTime', CustomerLogincol: 'CustomerLogincol' },
  //   { Customer_id: 'Customer3', LoginName: 'LoginName', Password: 'Password', UserStats: 'UserStats', ModifiedTime: 'ModifiedTime', CustomerLogincol: 'CustomerLogincol' },
  //   { Customer_id: 'Customer4', LoginName: 'LoginName', Password: 'Password', UserStats: 'UserStats', ModifiedTime: 'ModifiedTime', CustomerLogincol: 'CustomerLogincol' },
  //   { Customer_id: 'Customer5', LoginName: 'LoginName', Password: 'Password', UserStats: 'UserStats', ModifiedTime: 'ModifiedTime', CustomerLogincol: 'CustomerLogincol' }
  // ]

  const tabs = [
    { title: '文章管理' },
    { title: '用户管理' },
    { title: '违规封号' }
  ];

  const handleClick = (Sx: string) => {
    dispatch!({
      type: 'admin/queryUser',
      payload: {
        PageNo: 1,
        PageSize: 1,
        Sx
      }
    });
  }

  const handleCloseuser = (Uid: any) => {
    console.log(Uid)
    dispatch!({
      type: 'admin/queryCloseuser',
      payload: {
        Uid
      }
    });
  }

  return (<div className={styles.container}>
    <Tabs tabs={tabs} renderTabBar={props => <Tabs.DefaultTabBar {...props} page={tabs.length > 3 ? 3.5 : tabs.length} />}>
      <div className={styles.div1}>文章管理</div>
      <div className={styles.div2}>
        用户管理
      </div>

      <div className={styles.div3}>
          <SearchBar placeholder="查找用户" maxLength={20} onCancel={(val) => handleClick(val)} cancelText="查找" />
        {
          data && data.map((item: any) => (
            <div key={item.Customer_id} className={styles.box}>
              <WhiteSpace size="lg" />
              <Card className={styles.card}>
                <Card.Header
                  title={item.Customer_id}
                  extra={<Button className={styles.czBtn} size="small" onClick={() => operation([
                    { text: '发布违规内容', onPress: () => { handleCloseuser(item.Customer_id) } },
                    { text: '被大量用户投诉', onPress: () => { handleCloseuser(item.Customer_id) } },
                    { text: '其他违规', onPress: () => { handleCloseuser(item.Customer_id) } },
                  ])}
                  >操作</Button>}
                />
                <Card.Footer content={item.LoginName} extra={<div>{item.UserStats}</div>} />
                <Card.Footer content={item.CustomerLogincol} extra={<div>{item.ModifiedTime}</div>} />
              </Card>
            </div>
          ))
        }
      </div>
    </Tabs>
  </div>);
};

export default connect(({ admin }: { admin: AdminModelState; }) => ({ admin }))(AdminPage);
