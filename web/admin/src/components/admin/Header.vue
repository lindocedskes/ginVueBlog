<template>
  <a-button type="danger" @click="loginOut">退出</a-button>
</template>

<script>


export default {
  methods: {
    loginOut() {
      //清除token+返回登陆页面
      window.sessionStorage.clear('token')
      this.$router.push('/login')
    }
  }
}
</script>
