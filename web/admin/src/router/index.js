import Vue from 'vue'
import VueRouter from 'vue-router'
import Admin from '../views/Admin.vue'
import Login from "@/views/Login.vue";
Vue.use(VueRouter)

//页面路由组件
const Index = () => import(/* webpackChunkName: "AddArt" */ '../components/admin/Index.vue')
const AddArt = () => import(/* webpackChunkName: "AddArt" */ '../components/article/AddArt.vue')
const ArtList = () => import(/* webpackChunkName: "ArtList" */ '../components/article/ArtList.vue')
const CateList = () => import(/* webpackChunkName: "CateList" */ '../components/category/CateList.vue')
const UserList = () => import(/* webpackChunkName: "UserList" */ '../components/user/UserList.vue')

const Profile = () => import(/* webpackChunkName: "Index" */ '../components/user/Profile.vue')

const routes = [
  {
    path: '/login',
    name: 'login',
    meta: {
      title: '请登录'
    },
    component: Login
  },
  {
    path: '/admin',
    name: 'admin',
    component:Admin,
    children:[
      {
      path: 'index',
      component: Index,
      meta: {
        title: 'GinBlog 后台管理页面'
      }
    },
      {
        path: 'addart',
        component: AddArt,
        meta: {
          title: '新增文章'
        }
      },
      {
        path: 'addart/:id',
        component: AddArt,
        meta: {
          title: '编辑文章'
        },
        props: true
      },
      {
        path: 'artlist',
        component: ArtList,
        meta: {
          title: '文章列表'
        }
      },
      {
        path: 'catelist',
        component: CateList,
        meta: {
          title: '分类列表'
        }
      },
      {
        path: 'userlist',
        component: UserList,
        meta: {
          title: '用户列表'
        }

      },
      {
        path: 'profile',
        component: Profile,
        meta: {
          title: '个人设置'
        }
      }
    ]

  },
]

const router = new VueRouter({
  routes
})

//路由守卫，控制路由访问权限
router.beforeEach((to, from,next) => {
  const token = window.sessionStorage.getItem('token')
  if (to.path==='/login') return next() //请求登陆页面——放行
  //TODO 服务端验证token有效性
  if(!token){//其他页面，没有token
    next('/login') //跳转到登陆页面
  }else{
    next() //有token放行
  }
  return false
})
export default router
