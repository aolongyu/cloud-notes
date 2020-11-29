/**
 * websocket封装
 * 
 * @author aolyu
 * @created 2020/11/09 18:32:51
 */

/**
 * 在main.js或需要使用的地方引入并建立连接
 * import { createSocket } from 'websocket.js'
 * createSocket('wss://api.baidu.com')
 */

/**
 * 发送消息
 * import { sendWSPush } from 'websocket.js'
 * sendWSPush(data)
 */

/**
 * 接收消息
 * const getsocketData = e => {  // 创建接收消息函数
 *   const data = e && e.detail.data
 *   console.log(data)
 * }
 */

/**
 * 注册监听事件
 * window.addEventListener('onmessageWS', getsocketData)
 */

/**
 * 在需要的时候卸载监听事件，比如离开页面
 * window.removeEventListener('onmessageWS', getsocketData)
 */

let Socket = ''
const setIntervalWesocketPush = null

/**
 * 建立websocket连接
 * @param {string} url ws地址
 */

export const createSocket = (url = 'ws://localhost:8999') => {
  // Socket && Socket.close()
  if (!Socket) {
    if (url === undefined) {
      console.log(
        '%c%s',
        'color: white; background: red;',
        '服务器down了'
      )
    } else {
      console.log(
        '%c%s',
        'color: white; background: #5DAC81;',
        `建立websocket连接:${url}`
      )
    }
    Socket = new WebSocket(url)
    Socket.onopen = onopenWS
    Socket.onmessage = onmessageWS
    Socket.onerror = onerrorWS
    Socket.onclose = oncloseWS
  } else {
    console.log(
      '%c%s',
      'color: white; background: #5DAC81;',
      'websocket已连接'
    )
  }
}

/**打开WS之后发送心跳 */
const onopenWS = () => {
  sendPing()
}

/**连接失败重连 */
const onerrorWS = () => {
  Socket.close()
  clearInterval(setIntervalWesocketPush)
  console.log(
    '%c%s',
    'color: white; background: orange;',
    '连接失败重连中'
  )
  // console.log('连接失败重连中')
  if (Socket.readyState !== 3) {
    Socket = null
    createSocket()
  }
}

/**WS数据接收统一处理 */
const onmessageWS = e => {
  console.log(
    '%c%s',
    'color: white; background: #5DAC81;',
    `${'从服务端接收到: ' + e.data}`
  )
  window.dispatchEvent(new CustomEvent('onmessageWS', {
    detail: {
      data: e.data
    }
  }))
  return e.data
}

/**
 * 发送数据但连接未建立时进行处理等待重发
 * @param {any} message 需要发送的数据
 */
const connecting = message => {
  setTimeout(() => {
    if (Socket.readyState === 0) {
      connecting(message)
    } else {
      Socket.send(message)
    }
  }, 1000)
}

const strTo10Length = (str) => {
  const len = str.length
  if (len > 0 && len <= 10) {
    let addStr = ''
    for (let i = len; i < 10; i++) {
      addStr += ' '
    }
    console.log('strTo10Length: ', str + addStr, (str + addStr).length)
    return str + addStr
  }
  return false
}


/**
 * 发送数据
 * @param {string} type 需要发送的数据
 * @param {object} message 需要发送的数据
 */
export const sendWSPush = (type, message) => {
  // console.log(JSON.stringify(type + message))
  // console.log('Socket.readyState: ', Socket.readyState)
  if (Socket !== null && Socket.readyState === 3) {
    Socket.close()
    createSocket()
  } else if (Socket.readyState === 1) {
    console.log(
      '%c%s',
      'color: white; background: #5DAC81;',
      `向服务端发送message: ${strTo10Length(type) && JSON.stringify(strTo10Length(type) + message)}`
    )
    console.log(strTo10Length(type) && JSON.stringify(strTo10Length(type) + JSON.stringify(message)))
    Socket.send(strTo10Length(type) && JSON.stringify(strTo10Length(type) + JSON.stringify(message)))
  } else if (Socket.readyState === 0) {
    connecting(message)
  }
}

/**
 * 断开重连
 */
const oncloseWS = () => {
  clearInterval(setIntervalWesocketPush)
  console.log(
    '%c%s',
    'color: white; background: red;',
    'websocket已断开....正在尝试重连'
  )
  if (Socket.readyState !== 2) {
    Socket = null
    createSocket()
  }
}

/**
 * 发送心跳
 * @param {number} time 心跳间隔毫秒 默认5000
 * @param {string} ping 心跳名称 默认字符串ping
 */
export const sendPing = (time = 5000, ping = 'pingaolyu') => {
  clearInterval(setIntervalWesocketPush)
  // Socket.send(ping)
  // setIntervalWesocketPush = setInterval(() => {
  //   Socket.send(JSON.stringify(ping))
  // }, time)
} 