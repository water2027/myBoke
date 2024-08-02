<template>
    <div class="tags">
        <button v-for="tag in tags" :key="tag" @click="currentTag = tag">{{ tag }}</button>
    </div>
    <div v-for="post in curPosts" :key="post.id" class="postLink">
        <router-link :to="`/post/${post.id}`">
            <div>
                <h2>{{ post.title }}</h2>
                <p v-html="post.jmjx"></p>
                <div class="time"><span>    {{ post.created_at }}</span></div>
            </div>
        </router-link>
    </div>
</template>
<script setup>
import { ref , inject, watch, onMounted } from 'vue';
import { RouterLink } from 'vue-router';
const tags = ref(['闲谈','技术','更新'])
const currentTag = ref('');
const allPosts = inject('allPosts')
const curPosts = ref([])
watch(() => currentTag.value, (val) => {
    curPosts.value = allPosts.value.filter(post => post.tag == val)
})
</script>

<style scoped>
.tags{
    display: flex;
    justify-content: space-between;
}
p{
    background: linear-gradient(to right, #3ac5d5, #00ccfa);
    background-clip: text;
    color: transparent;
}
</style>