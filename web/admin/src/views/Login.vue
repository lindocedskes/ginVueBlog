<template>
<div class="container">
  <div class="loginBox">
    <a-form-model ref="loginFormRef" :rules="rules" :model="formdata" class="loginForm">
      <a-form-model-item prop="username">
        <a-input v-model="formdata.username" placeholder="请输入用户名">
          <a-icon slot="prefix" type="user" style="color:rgba(0,0,0,.25)" />
        </a-input>
      </a-form-model-item>

      <a-form-model-item pr  op="password">
        <a-input
          v-model="formdata.password"
          placeholder="请输入密码"
          type="password"
          v-on:keyup.enter="login"
        >
          <a-icon slot="prefix" type="lock" style="color:rgba(0,0,0,.25)" />
        </a-input>
      </a-form-model-item>

      <a-form-model-item class="loginBtn">
        <a-button type="primary" style="margin:10px" @click="login">登录</a-button>
        <a-button type="info" style="margin:10px" @click="resetForm">取消</a-button>
      </a-form-model-item>
    </a-form-model>
  </div>
</div>
</template>

<script>
export default {
  name: "Login",
  data() {
    return {
      formdata: {
        username: '',
        password: '',
      },
      rules: {
        username: [
          { required: true, message: '请输入用户名', trigger: 'blur' },
          {
            min: 4,
            max: 12,
            message: '用户名必须在4到12个字符之间',
            trigger: 'blur',
          },
        ],
        password: [
          { required: true, message: '请输入密码', trigger: 'blur' },
          {
            min: 6,
            max: 20,
            message: '密码必须在6到20个字符之间',
            trigger: 'blur',
          },
        ],
      },
    }
  },
  methods: {
    resetForm() {
      this.$refs.loginFormRef.resetFields()
    },
    login() {
      //参数 valid 是通过回调函数的方式传递进来的，valid，它的值是一个布尔值，表示表单是否验证通过
      this.$refs.loginFormRef.validate(async (valid) => {
        if (!valid) return this.$message.error('输入非法数据，请重新输入')
        const { data: res } = await this.$http.post('login', this.formdata) //提交给后台
        console.log(res)
        if (res.status != 200) return this.$message.error(res.message)//非200返回错误
        window.sessionStorage.setItem('token', res.token) //浏览器session保存token
        this.$router.push('/admin/index') //路由跳转
      })
    },
  },
}
</script>

<style scoped>
.container {
  height: 100%;
  background-color: #282c34;
}

.loginBox {
  width: 450px;
  height: 300px;
  background-color: #fff;
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  border-radius: 9px;
}
.loginForm {
  width: 100%;
  position: absolute;
  bottom: 10%;
  padding: 0 20px;
  box-sizing: border-box;
}

.loginBtn {
  display: flex;
  justify-content: flex-end;
}
</style>
