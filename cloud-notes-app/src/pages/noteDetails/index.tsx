import React, { FC, useEffect, useState } from 'react';
import { NoteDetailsModelState, ConnectProps, connect, setPageNavBar, history } from 'alita';
import NoMore from '@/components/noMore/index'
import styles from './index.less';

interface PageProps extends ConnectProps {
  noteDetails: NoteDetailsModelState;
}

const NoteDetailsPage: FC<PageProps> = ({ noteDetails, dispatch, location }) => {
  const { Id, Introduction, Name, Text, ThumbsUp, author } = location.query

  // const { data } = noteDetails;

  // console.log(`${location.pathname}${location.search}`)

  useEffect(() => {
    // dispatch!({
    //   type: 'noteDetails/queryNoteDetails',
    //   payload: {
    //     id: Number(NoteId)
    //   }
    // });
  }, []);

  const [readOnly, setReadOnly] = useState(true)

  // const msg = data && data[0]

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
        Note_id: Number(Id),
        Note_name: Name,
        Note_introduction: Introduction,
        Note_type: 0,
        Note_text: document.getElementById('text').value
      }
    });
    history.replace(`${location.pathname}${location.search}`)
  }
  const handleShare = () => { }

  return (
    <div className={styles.container}>
      <div className={styles.noteTitle}>
        <span className={styles.noteName}>{Name}</span>
        <span className={styles.hot}>{ThumbsUp}</span>
      </div>
      <div className={styles.noteInfo}>
        <span className={styles.author}>{author}</span>
        <div>{Introduction}</div>
      </div>
      <hr />
      <div className={styles.mainText}>
        <textarea readOnly={readOnly} name="text" id="text" className={styles.text}>{Text}</textarea>
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
