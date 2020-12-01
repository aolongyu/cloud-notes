import React, { useState, useEffect } from 'react';

import styles from './index.less';

/**
 * 卡片组件，接受类型详见 interface。
 * @param isHeader 可选，是否显示 header（默认 true）
 * @param isBody 可选，是否需要 body （默认 true）
 * @param isCollapse 可选，是否点击折叠 body （默认 false）
 * @param defaultExpand 可选，折叠控件是否展开 body（默认 false）
 * @param header 头部元素
 * @param body 卡身元素
 * ****
 * **默认样式修改**
 * ```css
 * * :global {
 * *   .selector {
 * *     ..........
 * *   }
 * * }
 * ```
 */

interface CardProps {
  cardName?: string;
  cardIntro?: string;
  cardModifyTime?: string;
}

const Card = ({ cardName = '无名文件', cardIntro = '作者很懒，没有说明', cardModifyTime = '2020.12.01' }: CardProps) => {

  return (
    <div className={styles.container}>
      <div className={styles.line1}>
        <span className={styles.left}>{cardName}</span>
        <span className={styles.right}>{cardModifyTime}</span>
      </div>
      <hr/>
      <div className={styles.line2}>
        <span>{cardIntro}</span>
      </div>
    </div>
  );
};

export default React.memo(Card);
