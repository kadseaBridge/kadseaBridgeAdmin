<template>
  <div class="login">
    <el-form ref="loginForm" :model="loginForm" :rules="loginRules" class="login-form">
      <h3 class="title">gfast后台管理系统</h3>
      <el-form-item prop="username">
        <el-input v-model="loginForm.username" type="text" auto-complete="off" placeholder="账号">
          <svg-icon slot="prefix" icon-class="user" class="el-input__icon input-icon"/>
        </el-input>
      </el-form-item>
      <el-form-item prop="password">
        <el-input
          v-model="loginForm.password"
          type="password"
          auto-complete="off"
          placeholder="密码"
          @keyup.enter.native="handleLogin"
        >
          <svg-icon slot="prefix" icon-class="password" class="el-input__icon input-icon"/>
        </el-input>
      </el-form-item>
      <el-form-item prop="code">
        <el-input
          v-model="loginForm.code"
          auto-complete="off"
          placeholder="验证码"
          style="width: 63%"
          @keyup.enter.native="handleLogin"
        >
          <svg-icon slot="prefix" icon-class="validCode" class="el-input__icon input-icon"/>
        </el-input>
        <div class="login-code">
          <img :src="codeUrl" @click="getCode"/>
        </div>
      </el-form-item>
      <el-checkbox v-model="loginForm.rememberMe" style="margin:0px 0px 25px 0px;">记住密码</el-checkbox>
      <el-form-item style="width:100%;">
        <el-button
          :loading="loading"
          size="medium"
          type="primary"
          style="width:100%;"
          @click.native.prevent="handleLogin"
        >
          <span v-if="!loading">登 录</span>
          <span v-else>登 录 中...</span>
        </el-button>
      </el-form-item>
    </el-form>

    <!--  底部  -->
    <div class="el-login-footer">
      <span>Copyright © 2021-2023 g-fast.cn All Rights Reserved.</span>
    </div>

    <!-- 谷歌验证绑定对话框 -->
    <el-dialog
      title="绑定谷歌验证"
      :visible.sync="googleAuthDialogVisible"
      width="30%"
      @close="handleGoogleAuthDialogClose"
    >
      <div v-if="googleAuthQRCode">
        <img :src="'data:image/png;base64,' + googleAuthQRCode" alt="QR Code" style="width: 100%; height: auto;"/>
      </div>
      <span slot="footer" class="dialog-footer">
<!--        <el-button @click="googleAuthDialogVisible = false">取消</el-button>-->
        <el-button @click="unBind">取消</el-button>
        <el-button type="primary" @click="confirmGoogleAuth">确认</el-button>
      </span>
    </el-dialog>

    <!-- 谷歌验证码输入对话框 -->
    <el-dialog
      title="谷歌验证"
      :visible.sync="googleCodeDialogVisible"
      width="30%"
      @close="handleGoogleCodeDialogClose"
    >
      <el-input
        v-model="googleCode"
        placeholder="请输入谷歌验证码"
        @keyup.enter.native="verifyGoogleCode"
      />
      <span slot="footer" class="dialog-footer">
        <el-button @click="googleCodeDialogVisible = false;googleCode = '' "  >取消</el-button>
        <el-button type="primary" @click="verifyGoogleCode">确认</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
import {getCodeImg} from "@/api/login";
import Cookies from "js-cookie";
import {encrypt, decrypt} from '@/utils/jsencrypt'

import {bindGoogleAuth} from "@/api/system/user";

export default {
  name: "Login",
  data() {
    return {
      codeUrl: "",
      cookiePassword: "",
      googleAuthQRCode: "", // 添加谷歌二维码的URL
      googleCode: "", // 添加谷歌验证码输入框的内容
      googleAuthDialogVisible: false, // 控制谷歌验证绑定对话框的显示
      googleCodeDialogVisible: false, // 控制谷歌验证码输入对话框的显示
      googleAuthRequired: false,
      bindGoogleAuth: false,
      qrcode: "",
      loginForm: {
        // username: "admin",
        // password: "123456",
        username: "",
        password: "",
        rememberMe: false,
        code: "1222",
        uuid: "",
        google: "",
      },
      loginRules: {
        username: [
          {required: true, trigger: "blur", message: "用户名不能为空"}
        ],
        password: [
          {required: true, trigger: "blur", message: "密码不能为空"}
        ],
        code: [{required: true, trigger: "change", message: "验证码不能为空"}]
      },
      loading: false,
      redirect: undefined
    };
  },
  watch: {
    $route: {
      handler: function (route) {
        this.redirect = route.query && route.query.redirect;
      },
      immediate: true
    }
  },
  created() {
    this.getCode();
    this.getCookie();
  },
  methods: {
    getCode() {
      getCodeImg().then(res => {
        this.codeUrl = res.data.base64stringC;
        this.loginForm.uuid = res.data.idKeyC;
      });
    },
    getCookie() {
      const username = Cookies.get("username");
      const google = Cookies.get("google");
      const password = Cookies.get("password");
      const rememberMe = Cookies.get('rememberMe')
      this.loginForm = {
        username: username === undefined ? this.loginForm.username : username,
        password: password === undefined ? this.loginForm.password : decrypt(password),
        google: google === undefined ? this.loginForm.google : google,
        rememberMe: rememberMe === undefined ? false : Boolean(rememberMe)
      };
    },
    handleLogin() {
      this.$refs.loginForm.validate(valid => {
        if (valid) {
          this.loading = true;
          if (this.loginForm.rememberMe) {
            Cookies.set("username", this.loginForm.username, {expires: 30});
            Cookies.set("password", encrypt(this.loginForm.password), {expires: 30});
            Cookies.set('rememberMe', this.loginForm.rememberMe, {expires: 30});
          } else {
            Cookies.remove("username");
            Cookies.remove("password");
            Cookies.remove('rememberMe');
          }
          this.$store
            .dispatch("Login1", this.loginForm)
            .then(response => {
              try {
                if (response.bindGoogleAuth) {
                  this.googleAuthQRCode = response.qrcode;
                  this.googleAuthDialogVisible = true;
                } else if (response.googleAuthRequired) {
                  this.googleCodeDialogVisible = true;
                } else {
                  this.$router.push({path: this.redirect || "/"});
                }
              } catch (e) {
                this.loading = false;
              }
            })
            .catch(() => {
              this.loading = false;
              this.getCode();
            });
        }
      });
    },

    verifyGoogleCode() {
      this.$store
        .dispatch('VerifyGoogleCode', {username: this.loginForm.username, googleCode: this.googleCode})

        .then(() => {
          this.googleCodeDialogVisible = false;
          this.$router.push({path: this.redirect || "/"});
        })
        .catch(() => {
          this.loading = false;
          this.getCode();
        });

      this.googleCode = ""
    },

    unBind(){
      this.$store.dispatch('UnBind',{username: this.loginForm.username})
        .then(() => {
          // this.googleCodeDialogVisible = false;
          this.googleAuthDialogVisible = false
          this.$router.push({path: this.redirect || "/"});
        })
        .catch(() => {
          this.loading = false;
          this.getCode();
        });
    },


    confirmGoogleAuth() {
      // 验证谷歌二维码绑定
      this.googleAuthDialogVisible = false;
      this.googleCodeDialogVisible = true;
    },
    handleGoogleAuthDialogClose() {
      this.googleAuthDialogVisible = false;
      this.loading = false;
      this.getCode();
    },
    handleGoogleCodeDialogClose() {
      this.googleCodeDialogVisible = false;
      this.loading = false;
      this.getCode();
    }

  }
};
</script>

<style rel="stylesheet/scss" lang="scss">
.login {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%;
  background-image: url("../assets/image/login-background.jpg");
  background-size: cover;
}

.title {
  margin: 0px auto 30px auto;
  text-align: center;
  color: #707070;
}

.login-form {
  border-radius: 6px;
  background: #ffffff;
  width: 400px;
  padding: 25px 25px 5px 25px;

  .el-input {
    height: 38px;

    input {
      height: 38px;
    }
  }

  .input-icon {
    height: 39px;
    width: 14px;
    margin-left: 2px;
  }
}

.login-tip {
  font-size: 13px;
  text-align: center;
  color: #bfbfbf;
}

.login-code {
  width: 33%;
  height: 38px;
  float: right;

  img {
    cursor: pointer;
    vertical-align: middle;
    width: 120px;
  }
}

.el-login-footer {
  height: 40px;
  line-height: 40px;
  position: fixed;
  bottom: 0;
  width: 100%;
  text-align: center;
  color: #fff;
  font-family: Arial;
  font-size: 12px;
  letter-spacing: 1px;
}
</style>
