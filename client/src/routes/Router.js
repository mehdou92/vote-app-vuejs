import VueRouter from "vue-router";

import Home from "../components/Home";
import Login from "../pages/User/Login";
import CreateUser from "../pages/User/CreateUser";
import ListUsers from "../pages/User/ListUser";
import User from "../pages/User/GetUser";
import CreateLaw from "../pages/Vote/AddLaw";
import ListLaws from "../pages/Vote/ListLaws";
import Law from "../pages/Vote/GetLaw";
import Error from "../pages/error/404";
import store from "@/store/user";

const router = new VueRouter({
    mode: "history",
    routes: [
        {
            path: '/',
            name: 'homepage',
            component: Home
        },
        {
            path: '/login',
            name: 'login',
            component: Login
        },
        {
            path: '/create-user',
            name: 'Create user',
            component: CreateUser,
            meta: { requiresAuth: true }
        },
        {
            path:'/users',
            name: 'List of users',
            component: ListUsers,
            meta: { requiresAuth: true }
        },
        {
            path: '/user/:id',
            name: 'Get user',
            component: User,
            meta: { requiresAuth: true }
        },
        {
            path: '/create-law',
            name: 'Create new law',
            component: CreateLaw
        },
        {
            path: '/laws',
            name: 'List of laws',
            component: ListLaws
        },
        {
            path: '/law/:id',
            name: 'Get law',
            component: Law
        },
        {
            path: '*',
            name: '404 error',
            component: Error
        }
    ]
});

router.beforeEach((to, from, next) => {
    if (to.matched.some(record => record.meta.requiresAuth)) {
      if (store.getters.adminPermission.access_level != 2) {
        next({
          path: '/*',
          query: { redirect: to.fullPath }
        })
      }
    } else {
      next()
    }
  })

export default router;
