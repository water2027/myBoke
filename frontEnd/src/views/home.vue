<template>
    <div v-for="post in posts" :key="post.id" class="postLink">
        <router-link :to="`/post/${post.id}`">
            <div>
                <h2>{{ post.title }}</h2>
                <p v-html="post.jmjx"></p>
                <div class="time"><span>{{ post.created_at }}</span></div>
            </div>
        </router-link>
    </div>
</template>
<script setup>
import { RouterLink } from 'vue-router'
import { inject } from 'vue'
const posts = inject('posts')

</script>
<style scoped>
p{
    background: linear-gradient(to right, #3ac5d5, #00ccfa);
    background-clip: text;
    color: transparent;
}
</style>