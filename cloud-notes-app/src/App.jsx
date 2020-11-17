import {createSocket, sendWSPush} from './utils/websocket'

import logo from './logo.svg';
import './App.css';

function App() {
  // 192.168.1.4
  createSocket('ws://localhost:8999')
  setTimeout(() => {
    sendWSPush('陈丹伟')
  }, 100);
  
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.js</code> and save to reload.
        </p>
      </header>
    </div>
  );
}

export default App;
