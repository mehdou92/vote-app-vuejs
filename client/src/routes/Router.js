import VueRouter from "vue-router";

import Home from "../components/Home";
import Login from "../pages/User/Login";
import CreateUser from "../pages/User/CreateUser";

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
        }
    ]
});

export default router;
