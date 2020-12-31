import React, { useState, useEffect } from 'react';

import Folder from '@/assets/noteFolder/folder.png'

import styles from './index.less';

interface CardProps {
  cardName?: string;
  cardIntro?: string;
  cardModifyTime?: string;
}

const Card = ({ cardName = '无名文件', cardIntro = '作者很懒，没有说明', cardModifyTime = '2020.12.01' }: CardProps) => {

  return (
    <div className={`${styles.container} clearfix`}>
      <div className={styles.left}>
        <img className={styles.folderImg} src={Folder} alt=""/>
      </div>
      <div className={styles.right}>
        <div className={styles.line1}>
          <span className={styles.line1Left}>{cardName}</span>
          <span className={styles.line1Right}>赞: {cardModifyTime}</span>
        </div>
        <hr />
        <div className={styles.line2}>
          <span className={styles.intro}>{cardIntro}</span>
        </div>
      </div>
    </div>
  );
};

export default React.memo(Card);
