import VueRouter from "vue-router";

import Home from "../components/Home";
import Login from "../pages/User/Login";
import CreateUser from "../pages/User/CreateUser";
import ListUsers from "../pages/User/ListUser";
import User from "../pages/User/GetUser";
import CreateLaw from "../pages/Vote/AddLaw";
import ListLaws from "../pages/Vote/ListLaws";
import Law from "../pages/Vote/GetLaw";

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
            component: CreateUser
        },
        {
            path:'/users',
            name: 'List of users',
            component: ListUsers
        },
        {
            path: '/user/:id',
            name: 'Get user',
            component: User
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
        }
    ]
});

export default router;
