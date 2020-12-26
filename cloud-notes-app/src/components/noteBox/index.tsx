import React, { useState, useEffect } from 'react';

import styles from './index.less';

interface CardProps {
  Id?: string;
  Name?: string;
  Introduction?: string;
  Text: string;
  ThumbsUp: string;
  click: () => {};
}

const NoteBox = ({ Id = '', Name = '', Introduction = '', Text = '', ThumbsUp = '', click }: CardProps) => {

  return (
    <div className={styles.container} onClick={click}>
      <div className={styles.name}>{Name}</div>
      <span className={styles.intro}>{Introduction}</span>
      <div className={styles.up}>{ThumbsUp}</div>
    </div>
  );
};

export default React.memo(NoteBox);
