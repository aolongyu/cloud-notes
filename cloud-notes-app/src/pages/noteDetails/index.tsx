import React, { FC, useEffect } from 'react';
import { NoteDetailsModelState, ConnectProps, connect } from 'alita';
import NoMore from '@/components/noMore/index'
import styles from './index.less';

interface PageProps extends ConnectProps {
  noteDetails: NoteDetailsModelState;
}

const NoteDetailsPage: FC<PageProps> = ({ noteDetails, dispatch, location }) => {
  const { NoteBookId } = location.query
  // console.log(NoteBookId)
  dispatch!({
    type: 'noteDetails/queryNoteDetails',
    payload: {
      
    }
  });

  const { data } = noteDetails;

  const test = {
    noteName: '软件工程',
    text: '前言：发现以前写的就像是笔记，哪像博客啊，这里再次修改。问题描述： 在固定宽度的p元素里（任何块级元素同理），长单词不自动换行，中文字符会自动换行，效果如：http://codepen.io/aliceluojuan/pen/rrxbpO产生原因：1.英文会将不包含空格、换行的连续文本认为是一个词，所以在默认情况下不换行;2.中文的话标点文字都是独立的，所以会自动换行;解决方案：在英文字不改变内容的情况下，通过设置p元素的前言：发现以前写的就像是笔记，哪像博客啊，这里再次修改。问题描述： 在固定宽度的p元素里（任何块级元素同理），长单词不自动换行，中文字符会自动换行，效果如：http://codepen.io/aliceluojuan/pen/rrxbpO产生原因：1.英文会将不包含空格、换行的连续文本认为是一个词，所以在默认情况下不换行;2.中文的话标点文字都是独立的，所以会自动换行;解决方案：在英文字不改变内容的情况下，通过设置p元素的前言：发现以前写的就像是笔记，哪像博客啊，这里再次修改。问题描述： 在固定宽度的p元素里（任何块级元素同理），长单词不自动换行，中文字符会自动换行，效果如：http://codepen.io/aliceluojuan/pen/rrxbpO产生原因：1.英文会将不包含空格、换行的连续文本认为是一个词，所以在默认情况下不换行;2.中文的话标点文字都是独立的，所以会自动换行;解决方案：在英文字不改变内容的情况下，通过设置p元素的前言：发现以前写的就像是笔记，哪像博客啊，这里再次修改。问题描述： 在固定宽度的p元素里（任何块级元素同理），长单词不自动换行，中文字符会自动换行，效果如：http://codepen.io/aliceluojuan/pen/rrxbpO产生原因：1.英文会将不包含空格、换行的连续文本认为是一个词，所以在默认情况下不换行;2.中文的话标点文字都是独立的，所以会自动换行;解决方案：在英文字不改变内容的情况下，通过设置p元素的',
    hot: '23',
    modifyDate: '2020.11.23'
  }

  return (
    <div className={styles.container}>
      <div className={styles.noteTitle}>
        <span className={styles.noteName}>{test.noteName}</span>
        <span className={styles.hot}>{test.hot}</span>
      </div>
      <div className={styles.noteInfo}>
        <span className={styles.author}>敖敖</span>
        <span className={styles.modityDate}>{test.modifyDate}</span>
      </div>
      <hr/>
      <div className={styles.mainText}>
        <span className={styles.text}>{test.text}</span>
      </div>
      <div className={styles.rightFloat}>
        <div className={styles.edit}></div>
        <div className={styles.share}></div>
      </div>
      <NoMore text='没有更多了' />
    </div>
  );
};

export default connect(({ noteDetails }: { noteDetails: NoteDetailsModelState; }) => ({ noteDetails }))(NoteDetailsPage);
