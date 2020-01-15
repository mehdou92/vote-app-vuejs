import VueRouter from "vue-router";

import Home from "../components/Home";

const router = new VueRouter({
    mode: "history",
    routes: [
        {
            path: '/',
            name: 'homepage',
            component: Home
        }
    ]
});

export default router;
