import React, { FC } from 'react';
import { CreateNoteModelState, ConnectProps, connect } from 'alita';
import styles from './index.less';

interface PageProps extends ConnectProps {
  createNote: CreateNoteModelState;
}

const CreateNotePage: FC<PageProps> = ({ createNote, dispatch }) => {
  // 这里发起了初始化请求
  // useEffect(() => {
  //   dispatch!({
  //     type: 'createNote/query',
  //   });
  //   return () => {
  //     // 这里写一些需要消除副作用的代码
  //     // 如: 声明周期中写在 componentWillUnmount
  //   };
  // }, []);
  // 注意，上面这里写空数组，表示初始化，如果需要监听某个字段变化再发起请求，可以在这里写明
  // const { name } = createNote;

  const handleClick = () => {
    const NoteBookName = document.getElementById('NoteBookName').value
    const NoteBookIntroduction = document.getElementById('NoteBookIntroduction').value
    const NoteBookType = document.getElementById('NoteBookType').value
    dispatch!({
      type: 'createNote/query',
      payload: {
        NoteBookName,
        NoteBookIntroduction,
        NoteBookType
      },
    });
  }

  return (
    <div className={styles.container}>
      <input type="text" className={styles.NoteBookName} id="NoteBookName"/>
      <input type="text" className={styles.NoteBookIntroduction} id="NoteBookIntroduction"/>
      <input type="text" className={styles.NoteBookType} id="NoteBookType"/>
      <input type="button" className={styles.submitBtn} onClick={handleClick} value="创建"/>
    </div>
  );
};

export default connect(({ createNote }:{ createNote: CreateNoteModelState; }) => ({ createNote }))(CreateNotePage);
