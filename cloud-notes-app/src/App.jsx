import {createSocket, sendWSPush} from './utils/websocket'

import logo from './logo.svg';
import './App.css';

function App() {
  createSocket('ws://10.12.83.81:8999')
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
