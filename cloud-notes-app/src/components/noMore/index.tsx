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
  text?: string
}

const NoMore = ({ text = '我是有底线的' }: CardProps) => {

  return (
    <div className={styles.container}>
      <span>------------------------   {text}   ------------------------</span>
    </div>
  );
};

export default React.memo(NoMore);
