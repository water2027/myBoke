<template>
    <div class="tags" ref="tagsContainer">
        <div v-for="tag in tags" :key="tag">
            <button @click="currentTag = tag">{{ tag }}</button>
        </div>
    </div>
    <div v-for="post in curPosts" :key="post.id" class="postLink">
        <router-link :to="`/post/${post.id}`">
            <div>
                <h2>{{ post.title }}</h2>
                <p v-html="post.jmjx"></p>
                <div class="time"><span> {{ post.created_at }}</span></div>
            </div>
        </router-link>
    </div>
</template>
<script setup>
import { ref, inject, watch } from 'vue';
import { RouterLink } from 'vue-router';
const tags = ref(['闲谈', '技术', '更新'])
const currentTag = ref('');
const allPosts = inject('allPosts')
const curPosts = ref([])
const tagsContainer = ref(null)

watch(() => currentTag.value, (val) => {
    curPosts.value = allPosts.value.filter(post => post.tag === val)
})



</script>

<style scoped>
.tags {
    display: flex;
    justify-content: space-between;
}

.tags div {
    --color: rgba(255, 255, 255, 0.5);
    width: 100px;
    height: 100px;
    border-radius: 50%;
    display: flex;
    justify-content: center;
    align-items: center;
    background: var(--color);
    backdrop-filter: blur(5px);
    scale: 1;
    transition: all 0.5s ease;
}

.tags div:hover {
    scale: 1.1;
}

p {
    background: linear-gradient(to right, #3ac5d5, #00ccfa);
    background-clip: text;
    color: transparent;
}
</style>