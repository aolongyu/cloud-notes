import React, { useState, useEffect } from 'react';
import { Modal, Button, WhiteSpace, WingBlank } from 'antd-mobile';

import styles from './index.less';

const operation = Modal.operation;

interface CardProps {
  Customer: string;
  LoginName: string;
  Password: string;
  UserStats: string;
  ModifiedTime: string;
  CustomerLogincol: string;
  clickFunc: () => {}
}



const FhBox = (data: CardProps) => (
  <div className={styles.container}>
    <p>{data.Customer}</p>
    <p>{data.LoginName}</p>
    <p>{data.Password}</p>
    <p>{data.UserStats}</p>
    <p>{data.ModifiedTime}</p>
    <p>{data.CustomerLogincol}</p>
    <WingBlank size="lg">
      <Button onClick={() => operation([
        { text: '发布违规内容', onPress: data.clickFunc },
        { text: '被大量用户投诉', onPress: data.clickFunc },
        { text: '其他违规', onPress: data.clickFunc },
      ])}
      >操作</Button>
      <WhiteSpace size="lg" />
    </WingBlank>
  </div>
);

export default React.memo(FhBox);
