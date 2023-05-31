<template>
  <a-layout id="components-layout-demo-basic">
      <Nav></Nav>
    <a-layout>
      <a-layout-header class="headerBtn">
        <Header></Header>
      </a-layout-header>
      <a-layout-content>
        <router-view :key="$route.fullPath"></router-view>
      </a-layout-content>
      <a-layout-footer class="footer">
        <Footer></Footer>
      </a-layout-footer>
    </a-layout>
  </a-layout>
</template>

<script>
import Nav from '../components/admin/Nav.vue'
import Footer from '../components/admin/Footer.vue'
import Header from '../components/admin/Header.vue'

export default {
  name: "Admin",
  components:{Nav,Footer,Header}
}
</script>

<style scoped>
#components-layout-demo-basic {
  height: 100%;
}
.headerBtn {
  display: flex;
  justify-content: flex-end;
  align-items: center;
}
</style>
