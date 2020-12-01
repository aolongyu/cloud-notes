import React, { FC, useEffect } from 'react';
import { NoteDetailsModelState, ConnectProps, connect } from 'alita';
import styles from './index.less';

interface PageProps extends ConnectProps {
  noteDetails: NoteDetailsModelState;
}

const NoteDetailsPage: FC<PageProps> = ({ noteDetails, dispatch }) => {
  dispatch!({
    type: 'noteDetails/query',
  });
  const { data } = noteDetails;
  return (
    <div className={styles.container}>

    </div>
  );
};

export default connect(({ noteDetails }: { noteDetails: NoteDetailsModelState; }) => ({ noteDetails }))(NoteDetailsPage);
