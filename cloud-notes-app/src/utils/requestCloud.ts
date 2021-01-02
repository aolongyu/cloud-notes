/**
 * 封装请求参数
 * 
 * @author aolyu
 * @created 2020/11/29 18:00:10
 */

import { sendWSPush } from '@/utils/websocket'

// eslint-disable-next-line func-names
const requestCloud = function (type: string, msg: object | string | number | undefined | null) {
  return new Promise(resolve => {
    sendWSPush(type, msg)
    // eslint-disable-next-line no-restricted-globals
    addEventListener('onmessageWS', (e) => {
      resolve(e.detail.data)
    })
  }).then((data: any) => {
    window.cloud = JSON.stringify(data.replace(/(^\s*)|(\s*$)/g, "")); 
  })
}

export default requestCloud