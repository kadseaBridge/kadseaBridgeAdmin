import { login, logout, getInfo,verifyGoogleCode } from '@/api/login'
import { getToken, setToken, removeToken } from '@/utils/auth'
import {getUpFileUrl} from "@/utils/ruoyi";

const user = {
  state: {
    token: getToken(),
    name: '',
    avatar: '',
    roles: [],
    permissions: []
  },

  mutations: {
    SET_TOKEN: (state, token) => {
      state.token = token
    },
    SET_NAME: (state, name) => {
      state.name = name
    },
    SET_AVATAR: (state, avatar) => {
      state.avatar = avatar
    },
    SET_ROLES: (state, roles) => {
      state.roles = roles
    },
    SET_PERMISSIONS: (state, permissions) => {
      state.permissions = permissions
    }
  },

  actions: {
    // 登录
    Login({ commit }, userInfo) {
      const username = userInfo.username.trim()
      const password = userInfo.password
      const google = userInfo.google
      const code = userInfo.code
      const uuid = userInfo.uuid
      return new Promise((resolve, reject) => {
        login(username, password, code, uuid, google)
          .then(response => {
            const { data } = response
            console.log("收到服务器返回的数据:", data)

            try {
              if (data.bindGoogleAuth || data.googleAuthRequired) {
                console.log("返回谷歌验证数据：", data)
                resolve(data) // 返回数据以便前端处理谷歌验证
              } else {
                commit('SET_TOKEN', data.token)
                setToken(data.token)
                resolve(data)
              }
            } catch (error) {
              console.error("处理服务器返回的数据时发生错误:", error)
              reject(error)
            }
          })
          .catch(error => {
            console.error("Login action 捕获到错误:", error)
            reject(error)
          })
      })
    },

    VerifyGoogleCode({ commit }, { username, code }) {
      return new Promise((resolve, reject) => {
        verifyGoogleCode(username, code)
          .then(response => {
            const { data } = response
            commit('SET_TOKEN', data.token)
            setToken(data.token)
            resolve(response)
          })
          .catch(error => {
            reject(error)
          })
      })
    },

    // 获取用户信息
    GetInfo({ commit, state }) {
      return new Promise((resolve, reject) => {
        getInfo(state.token).then(res => {
          const user = res.data.user
          const avatar = user.avatar == "" ? require("@/assets/image/profile.jpg") : getUpFileUrl(process.env.VUE_APP_BASE_API , user.avatar);

          if (res.data.roles && res.data.roles.length > 0) { // 验证返回的roles是否是一个非空数组
            commit('SET_ROLES', res.data.roles)
            commit('SET_PERMISSIONS', res.data.permissions)
          } else {
            commit('SET_ROLES', ['ROLE_DEFAULT'])
          }
          commit('SET_NAME', user.user_name)
          commit('SET_AVATAR', avatar)
          resolve(res)
        }).catch(error => {
          reject(error)
        })
      })
    },

    // 退出系统
    LogOut({ commit, state }) {
      return new Promise((resolve, reject) => {
        logout(state.token).then(() => {
          commit('SET_TOKEN', '')
          commit('SET_ROLES', [])
          commit('SET_PERMISSIONS', [])
          removeToken()
          resolve()
        }).catch(error => {
          reject(error)
        })
      })
    },

    // 前端 登出
    FedLogOut({ commit }) {
      return new Promise(resolve => {
        commit('SET_TOKEN', '')
        removeToken()
        resolve()
      })
    },

    // VerifyGoogleCode({ commit }, { username, code }) {
    //   return new Promise((resolve, reject) => {
    //     verifyGoogleCode(username, code).then(response => {
    //       resolve(response)
    //     }).catch(error => {
    //       reject(error)
    //     })
    //   })
    // },
  }
}

export default user
