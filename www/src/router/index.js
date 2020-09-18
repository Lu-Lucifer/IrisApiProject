import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

/* Layout */
import Layout from '@/layout'

/* Router Modules */
/* import componentsRouter from './modules/components'
 import chartsRouter from './modules/charts'*/
// import tableRouter from './modules/table'
/* import nestedRouter from './modules/nested'*/

/**
 * Note: sub-menu only appear when route children.length >= 1
 * Detail see: https://panjiachen.github.io/vue-element-admin-site/guide/essentials/router-and-nav.html
 *
 * hidden: true                   if set true, item will not show in the sidebar(default is false)
 * alwaysShow: true               if set true, will always show the root menu
 *                                if not set alwaysShow, when item has more than one children route,
 *                                it will becomes nested mode, otherwise not show the root menu
 * redirect: noRedirect           if set noRedirect will no redirect in the breadcrumb
 * name:'router-name'             the name is used by <keep-alive> (must set!!!)
 * meta : {
    roles: ['admin','editor']    control the page roles (you can set multiple roles)
    title: 'title'               the name show in sidebar and breadcrumb (recommend set)
    icon: 'svg-name'/'el-icon-x' the icon show in the sidebar
    noCache: true                if set true, the page will no be cached(default is false)
    affix: true                  if set true, the tag will affix in the tags-view
    breadcrumb: false            if set false, the item will hidden in breadcrumb(default is true)
    activeMenu: '/example/list'  if set path, the sidebar will highlight the path you set
  }
 */

/**
 * constantRoutes
 * a base page that does not have permission requirements
 * all roles can be accessed
 */
export const constantRoutes = [
  {
    path: '/redirect',
    component: Layout,
    hidden: true,
    children: [
      {
        path: '/redirect/:path(.*)',
        component: () => import('@/views/redirect/index')
      }
    ]
  },
  {
    path: '/admin/login',
    component: () => import('@/views/login/index'),
    hidden: true
  },
  {
    path: '/auth-redirect',
    component: () => import('@/views/login/auth-redirect'),
    hidden: true
  },
  {
    path: '/404',
    component: () => import('@/views/error-page/404'),
    hidden: true
  },
  {
    path: '/401',
    component: () => import('@/views/error-page/401'),
    hidden: true
  },
  {
    path: '/admin',
    component: Layout,
    redirect: '/admin/dashboard',
    children: [
      {
        path: '/admin/dashboard',
        component: () => import('@/views/dashboard/admin/index'),
        name: '控制面板',
        meta: { title: '控制面板', icon: 'dashboard', affix: true }
      }
    ]
  },
  {
    path: '/admin/profile',
    component: Layout,
    redirect: '/admin/profile/index',
    hidden: true,
    children: [
      {
        path: '/admin/profile/index',
        component: () => import('@/views/profile/index'),
        name: 'Profile',
        meta: { title: 'Profile', icon: 'user', noCache: true }
      }
    ]
  }
]

/**
 * asyncRoutes
 * the routes that need to be dynamically loaded based on user roles
 */
export const asyncRoutes = [
  {
    path: '/admin/permission',
    component: Layout,
    redirect: '/admin/permission/page',
    alwaysShow: true, // will always show the root menu
    name: '系统权限',
    meta: {
      title: '系统权限',
      icon: 'lock',
      roles: ['admin', 'editor'] // you can set roles in root nav
    },
    children: [
      {
        path: '/admin/permission/page',
        component: () => import('@/views/permission/perm'),
        name: '权限页面',
        meta: {
          title: '权限页面',
          roles: ['admin'] // or you can only set roles in sub nav
        }
      },
      {
        path: '/admin/permission/directive',
        component: () => import('@/views/permission/user'),
        name: '用户页面',
        meta: {
          title: '用户页面',
          roles: ['admin']
        }
      },
      {
        path: '/admin/permission/role',
        component: () => import('@/views/permission/role'),
        name: '角色页面',
        meta: {
          title: '角色页面',
          roles: ['admin']
        }
      }
    ]
  },
  {
    path: '/admin/article',
    component: Layout,
    redirect: '/admin/article/index',
    alwaysShow: true, // will always show the root menu
    name: '文章管理',
    meta: {
      title: '文章管理',
      icon: 'lock',
      roles: ['admin', 'editor'] // you can set roles in root nav
    },
    children: [
      {
        path: '/admin/article/index',
        component: () => import('@/views/article/list'),
        name: '文章列表',
        meta: {
          title: '文章列表',
          roles: ['admin'] // or you can only set roles in sub nav
        }
      },
      {
        path: '/admin/article/add',
        component: () => import('@/views/article/create'),
        name: '添加文章',
        meta: {
          title: '添加文章',
          roles: ['admin'] // or you can only set roles in sub nav
        }
      },
      // {
      //   path: '/admin/article/tag',
      //   component: () => import('@/views/article/tag'),
      //   name: '文章标签',
      //   meta: {
      //     title: '文章标签',
      //     roles: ['admin']
      //   }
      // },
      // {
      //   path: '/admin/article/type',
      //   component: () => import('@/views/article/type'),
      //   name: '文章分类',
      //   meta: {
      //     title: '文章分类',
      //     roles: ['admin']
      //   }
      // }
    ]
  },

  /** when your routing map is too long, you can split it into small modules **/
  // componentsRouter,
  // chartsRouter,
  // nestedRouter,
  // tableRouter,

  // 404 page must be placed at the end !!!
  { path: '*', redirect: '/404', hidden: true }
]

const createRouter = () => new Router({
  // mode: 'history', // require service support
  scrollBehavior: () => ({ y: 0 }),
  routes: constantRoutes
})

const router = createRouter()

// Detail see: https://github.com/vuejs/vue-router/issues/1234#issuecomment-357941465
export function resetRouter() {
  const newRouter = createRouter()
  router.matcher = newRouter.matcher // reset router
}

export default router
