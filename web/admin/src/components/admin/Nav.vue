<template >
  <a-layout-sider breakpoint="lg" v-model="collapsed">
    <div class="log">
      <span>{{ collapsed ? 'Blog' : 'Lyt\' Blog' }}</span>
    </div>

    <a-menu theme="dark" mode="inline" @click="goToPage">
      <a-menu-item key="index">
        <a-icon type="dashboard"></a-icon>
        <span>仪表盘</span>
      </a-menu-item>

      <a-sub-menu>
        <span slot="title">
          <a-icon type="file" />
          <span>文章管理</span>
        </span>
        <a-menu-item key="addart">
          <a-icon type="form" />
          <span>写文章</span>
        </a-menu-item>
        <a-menu-item key="artlist">
          <a-icon type="snippets" />
          <span>文章列表</span>
        </a-menu-item>
      </a-sub-menu>
      <a-menu-item key="catelist">
        <a-icon type="book" />
        <span>分类列表</span>
      </a-menu-item>
      <a-menu-item key="userlist">
        <a-icon type="user" />
        <span>用户列表</span>
      </a-menu-item>

      <a-menu-item key="profile">
        <a-icon type="setting" />
        <span>个人设置</span>
      </a-menu-item>
    </a-menu>
  </a-layout-sider>
</template>

<script>
export default {
  data() {
    return {
      collapsed: false,
    }
  },
  methods: {
    goToPage(item) {
      this.$router.push('/admin/' + item.key).catch((err) => err)
    },
  },
}
</script>

<style scoped>
.log {
  height: 50px;
  margin: 16px;
  background-color: #00c0e2;
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 22px;
  font-weight: bold;
  border-radius: 5px;
  box-shadow: 0px 2px 5px rgba(0, 0, 0, 0.2);
  text-transform: uppercase;
  letter-spacing: 1px;
  cursor: pointer;
  transition: all 0.3s ease;
}
.log:hover {
  transform: translateY(-3px);
  box-shadow: 0px 4px 8px rgba(0, 0, 0, 0.3);
}
</style>
