import {createSocket, sendWSPush} from './utils/websocket'

import logo from './logo.svg';
import './App.css';

function App() {
  createSocket('ws://192.168.1.4:8999')
  setTimeout(() => {
    sendWSPush('hello')
  }, 500);
  
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
