import React, { FC, useEffect, useState } from 'react';
import { NoteDetailsModelState, ConnectProps, connect } from 'alita';
import NoMore from '@/components/noMore/index'
import styles from './index.less';

interface PageProps extends ConnectProps {
  noteDetails: NoteDetailsModelState;
}

const NoteDetailsPage: FC<PageProps> = ({ noteDetails, dispatch, location }) => {
  const { NoteId, Name } = location.query

  useEffect(() => {
    dispatch!({
      type: 'noteDetails/queryNoteDetails',
      payload: {
        NoteId
      }
    });
  }, []);

  const [readOnly, setReadOnly] = useState(true)

  const { data } = noteDetails;
  const msg = data && data[0]
  console.log(msg)

  // const data = {
  //   Id: '',
  //   Name: '软件工程',
  //   Introduction: 'dsefdsffds',
  //   Text: '前言：发现以前写的就像是笔记，哪像博客啊，这里再次修改。问题描述： 在固定宽度的p元素里（任何块级元素同理），长单词不自动换行，中文字符会自动换行，效果如：http://codepen.io/aliceluojuan/pen/rrxbpO产生原因：1.英文会将不包含空格、换行的连续文本认为是一个词，所以在默认情况下不换行;2.中文的话标点文字都是独立的，所以会自动换行;解决方案：在英文字不改变内容的情况下，通过设置p元素的前言：发现以前写的就像是笔记，哪像博客啊，这里再次修改。问题描述： 在固定宽度的p元素里（任何块级元素同理），长单词不自动换行，中文字符会自动换行，效果如：http://codepen.io/aliceluojuan/pen/rrxbpO产生原因：1.英文会将不包含空格、换行的连续文本认为是一个词，所以在默认情况下不换行;2.中文的话标点文字都是独立的，所以会自动换行;解决方案：在英文字不改变内容的情况下，通过设置p元素的前言：发现以前写的就像是笔记，哪像博客啊，这里再次修改。问题描述： 在固定宽度的p元素里（任何块级元素同理），长单词不自动换行，中文字符会自动换行，效果如：http://codepen.io/aliceluojuan/pen/rrxbpO产生原因：1.英文会将不包含空格、换行的连续文本认为是一个词，所以在默认情况下不换行;2.中文的话标点文字都是独立的，所以会自动换行;解决方案：在英文字不改变内容的情况下，通过设置p元素的前言：发现以前写的就像是笔记，哪像博客啊，这里再次修改。问题描述： 在固定宽度的p元素里（任何块级元素同理），长单词不自动换行，中文字符会自动换行，效果如：http://codepen.io/aliceluojuan/pen/rrxbpO产生原因：1.英文会将不包含空格、换行的连续文本认为是一个词，所以在默认情况下不换行;2.中文的话标点文字都是独立的，所以会自动换行;解决方案：在英文字不改变内容的情况下，通过设置p元素的',
  //   ThumbsUp: '23',
  // }

  const handleEdit = () => {
    setReadOnly(false)
    const element = document.getElementById('text')
    const edit = document.getElementById('edit')
    const save = document.getElementById('save')
    element.style.backgroundColor = '#fff'
    save.style.display = 'block'
    edit.style.display = 'none'
  }

  const handleSave = () => {
    setReadOnly(true)
    const element = document.getElementById('text')
    const edit = document.getElementById('edit')
    const save = document.getElementById('save')
    element.style.backgroundColor = 'transparent'
    edit.style.display = 'block'
    save.style.display = 'none'
    dispatch!({
      type: 'noteDetails/queryUpdateNote',
      payload: {
        NoteId,
      }
    });
   }

  const handleShare = () => { }

  return (
    <div className={styles.container}>
      <div className={styles.noteTitle}>
        <span className={styles.noteName}>{msg && msg.Name}</span>
        <span className={styles.hot}>{msg && msg.ThumbsUp}</span>
      </div>
      <div className={styles.noteInfo}>
        <span className={styles.author}>{Name}</span>
      </div>
      <hr />
      <div className={styles.mainText}>
        <textarea readOnly={readOnly} name="text" id="text" className={styles.text}>{msg && msg.Text}</textarea>
      </div>
      <div className={styles.rightFloat}>
        <div id="edit" className={styles.edit} onClick={handleEdit}></div>
        <div id="save" className={styles.save} onClick={handleSave}></div>
        <div id="share" className={styles.share} onClick={handleShare}></div>
      </div>
      <NoMore text='没有更多了' />
    </div>
  );
};

export default connect(({ noteDetails }: { noteDetails: NoteDetailsModelState; }) => ({ noteDetails }))(NoteDetailsPage);
