import { createRouter, createWebHistory } from "vue-router";
import Home from "../views/home.vue";
import post from "../views/post.vue";
const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: "/",
            name: "home",
            component: Home,
        },
        {
            path: "/login",
            name: "login",
            component: () =>
                import(/* webpackPrefetch: true */ "../views/login.vue"),
        },
        {
            path: "/tags",
            name: "tags",
            component: () =>
                import(/* webpackPrefetch: true */ "../views/tags.vue"),
        },
        {
            path: "/friends",
            name: "friends",
            component: () =>
                import(/* webpackPrefetch: true */ "../views/friends.vue"),
        },
        {
            path: "/chatroom",
            name: "chatroom",
            component: () =>
                import(/* webpackPrefetch: true */ "../views/chatroom.vue"),
        },
        {
            path: "/post/:id",
            name: "post",
            props: true,
            component: post,
        },
        {
            path: "/game",
            name: "game",
            component: () =>
                import(/* webpackPrefetch: true */ "../views/game.vue"),
        },
        {
            path: "/edit",
            name: "edit",
            component: () =>
                import(/* webpackPrefetch: true */ "../views/edit.vue"),
        }
    ],
});

export default router;
