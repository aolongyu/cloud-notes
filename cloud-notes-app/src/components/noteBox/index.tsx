import React, { useState, useEffect } from 'react';

import { Modal, List, Button, WhiteSpace, WingBlank, Icon, Toast, Popover, Card } from 'antd-mobile';

import styles from './index.less';

const Item1 = Popover.Item;


interface CardProps {
  Id?: string;
  Name?: string;
  Introduction?: string;
  Text: string;
  ThumbsUp: string;
  index: number
  click: (index: number) => {};
  onSelect: () => {}
}

const NoteBox = ({ Id = '', Name = '', Introduction = '', Text = '', ThumbsUp = '', click, onSelect, index }: CardProps) => {

  const Item = List.Item;
  const Brief = Item.Brief;

  const [visible, setVisible] = useState(false)
  // const handleVisibleChange = (visible: boolean) => {
  //   setVisible(!visible)
  // };

  return (
    <div className={styles.container}>
      <Item
        arrow="empty"
        multipleLine
        onClick={() => { click(index) }}
        platform="android"
      >
        ({ThumbsUp}) {Name}：<Brief>{Introduction}</Brief>
      </Item>
      <Popover mask
        overlayClassName="fortest"
        overlayStyle={{ color: 'currentColor' }}
        visible={visible}
        overlay={[
          (<Item1 key="4" dataSeed="1">加入到笔记本</Item1>),
          (<Item1 key="5" dataSeed="2">删除</Item1>)
        ]}
        align={{
          overflow: { adjustY: 0, adjustX: 0 },
          offset: [-10, 0],
        }}
        onSelect={(e) => { onSelect(e, Id) }}
      >
        <div style={{
          position: 'absolute',
          top: '50px',
          right: '30px'
        }}
        >
          <Icon type="ellipsis" />
        </div>
      </Popover>
    </div>
  );
};

export default React.memo(NoteBox);
