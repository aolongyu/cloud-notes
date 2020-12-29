import React, { useState, useEffect } from 'react';

import { Modal, List, Button, WhiteSpace, WingBlank, Icon, Toast, Popover } from 'antd-mobile';

import styles from './index.less';

const Item = Popover.Item;


interface CardProps {
  Id?: string;
  Name?: string;
  Introduction?: string;
  Text: string;
  ThumbsUp: string;
  click: () => {};
  onSelect: () => {}
}

const NoteBox = ({ Id = '', Name = '', Introduction = '', Text = '', ThumbsUp = '', click,onSelect }: CardProps) => {

  const [visible, setVisible] = useState(false)
  // const handleVisibleChange = (visible: boolean) => {
  //   setVisible(!visible)
  // };
  
  return (
    <div className={styles.container}>
      <div className={styles.name} onClick={() => { click(Id) }}>{Name}</div>
      <span className={styles.intro}>{Introduction}</span>
      <div className={styles.up}>{ThumbsUp}</div>
      <Popover mask
        overlayClassName="fortest"
        overlayStyle={{ color: 'currentColor' }}
        visible={visible}
        overlay={[
          (<Item key="4" data-seed="logId">加入到笔记本</Item>)
        ]}
        align={{
          overflow: { adjustY: 0, adjustX: 0 },
          offset: [-10, 0],
        }}
        // onVisibleChange={handleVisibleChange}
        onSelect={() => {onSelect(Id)}}
      >
        <div style={{
          position: 'absolute',
          top: '15px',
          right: '15px'
        }}
        >
          <Icon type="ellipsis" />
        </div>
      </Popover>
      
    </div>
  );
};

export default React.memo(NoteBox);
