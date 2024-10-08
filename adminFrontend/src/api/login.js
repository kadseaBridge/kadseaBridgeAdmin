import request from '@/utils/request'
// var Web3 = require("web3")

// var web3
// if (window.ethereum) {
//     var web3 = new Web3(window.ethereum)
//     ethereum.enable()
// }


// 登录方法
export function login1(username, password, code, uuid) {
  console.log(username, password, code, uuid)
  const data = {
    username,
    password,
    verifyCode:code,
    verifyKey:uuid
  }
  return request({
    url: '/system/auth/login1',
    method: 'post',
    data: data
  })
}

// 获取用户详细信息
export function getInfo() {
  return request({
    url: '/system/user/getInfo',
    method: 'get'
  })
}

// 退出方法
export function logout() {
  return request({
    url: '/system/logout',
    method: 'post'
  })
}

// 获取验证码
export function getCodeImg() {
  return request({
    url: '/captcha/get',
    method: 'get'
  })
}

export function verifyGoogleCode(username, googleCode) {
  return request({
    url: '/system/verifyGoogleCode',
    method: 'post',
    data: { username, googleCode }
  })
}

// 取消绑定谷歌验证
export function unBind(userName) {
  return request({
    url: '/system/auth/unBind',
    method: 'post',
    data: { userName }
  })
}
