import VueRouter from "vue-router";

import Home from "../components/Home";
import Login from "../pages/User/Login";

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
        }
    ]
});

export default router;
