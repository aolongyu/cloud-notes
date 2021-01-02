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
    dispatch!({
      type: 'admin/queryAllNote',
    })
    return () => {
    };
  }, []);
  // 注意，上面这里写空数组，表示初始化，如果需要监听某个字段变化再发起请求，可以在这里写明
  const { data } = admin;

  const tabs = [
    { title: '文章管理', value: 0 },
    { title: '用户管理', value: 1 },
    { title: '违规封号', value: 2 }
  ];

  const handleClick = (Sx: string) => {
    dispatch!({
      type: 'admin/queryUser',
      payload: {
        PageNo: 1,
        PageSize: 99,
        Sx
      }
    });
  }

  const handleCloseuser = (Name: any) => {
    console.log(Name)
    dispatch!({
      type: 'admin/queryCloseuser',
      payload: {
        tname: Name
      }
    });
  }
  console.log(data)

  const changeTab = (e) => {
    console.log(e)
    if (e.value === 0) {
      dispatch!({
        type: 'admin/queryAllNote',
      })
    } else {
      dispatch!({
        type: 'admin/queryUser',
        payload: {
          PageNo: 1,
          PageSize: 99,
          Sx: ''
        }
      })
    }
  }

  const handleRePwd = (Id: number) => {
    dispatch!({
      type: 'admin/queryRePwd',
      payload: {
        id: Number(Id)
      }
    });
  }

  return (<div className={styles.container}>
    <Tabs tabs={tabs} onChange={changeTab} renderTabBar={props => <Tabs.DefaultTabBar {...props} page={tabs.length > 3 ? 3.5 : tabs.length} />}>
      <div className={styles.div1}>
        <SearchBar placeholder="查找笔记" maxLength={20} onCancel={(val) => handleClick(val)} cancelText="查找" />
        {
          data && data.map && data.map((item: any) => (
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
      <div className={styles.div2}>
        <SearchBar placeholder="查找用户" maxLength={20} onCancel={(val) => handleClick(val)} cancelText="查找" />
        {
          data && data.map && data.map((item: any) => (
            <div key={item.Customer_id} className={styles.box}>
              <WhiteSpace size="lg" />
              <Card className={styles.card}>
                <Card.Header
                  title={item.Name}
                  extra={<Button className={styles.czBtn} size="small" onClick={() => operation([
                    { text: '重置密码', onPress: () => { handleRePwd(item.Id) } },
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
      <div className={styles.div3}>
        <SearchBar placeholder="用户ID" maxLength={20} onCancel={(val) => handleClick(val)} cancelText="查找" />
        {
          data && data.map && data.map((item: any) => (
            <div key={item.Name} className={styles.box}>
              <WhiteSpace size="lg" />
              <Card className={styles.card}>
                <Card.Header
                  title={<div><span>({item.Id}) </span>{item.Name}</div>}
                  extra={<Button className={styles.czBtn} size="small" onClick={() => operation([
                    { text: '发布违规内容', onPress: () => { handleCloseuser(item.Name) } },
                    { text: '被大量用户投诉', onPress: () => { handleCloseuser(item.Name) } },
                    { text: '其他违规', onPress: () => { handleCloseuser(item.Name) } },
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
