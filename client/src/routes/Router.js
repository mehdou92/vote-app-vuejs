import VueRouter from "vue-router";

import Home from "../components/Home";
import Login from "../pages/User/Login";
import CreateUser from "../pages/User/CreateUser";
import ListUser from "../pages/User/ListUser";

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
            path:'/list-users',
            name: 'List of users',
            component: ListUser
        }
    ]
});

export default router;
