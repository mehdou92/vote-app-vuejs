import VueRouter from "vue-router";
import HeaderBoard from "../components/Header/Board";
import HeaderList from "../components/Header/List";

import Board from "../components/Board.vue";
import List from "../components/List.vue";
import ListHome from "../components/ListHome.vue";

const router = new VueRouter({
    mode: "history",
    routes: [
        {
            path: "/",
            name:"HomeBoard",
            components: {
                header: HeaderBoard,
                default: Board
            }
        },
        {
            path: "/list",
            name:"HomeList",
            components: {
                header: HeaderList,
                default: ListHome
            },
            children: [
                {
                    path: ":listid", 
                    component: List,
                    props: (route) =>  ({
                        name: router.params.listid,
                        cards: []
                    })
                }
            ]

        },
    ]
});

export default router;
