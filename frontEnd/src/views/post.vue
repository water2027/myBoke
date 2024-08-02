<template>
    <div class="article">
        <h2>{{ post.title }}</h2>
        <article v-html="post.body"></article>
        <div class="articleTime">{{ post.created_at }}</div>
    </div>
    <div>
        <textarea ref="replyInput" placeholder="想说点什么吗"></textarea>
        <button @click="submitReply">提交</button>
    </div>
    <div>
        <div v-for="reply in replys" :key="reply.id" class="article">
            <p>{{ reply.name }}</p>
            <div>{{ reply.reply }}</div>
            <div>{{ reply.time }}</div>
            <button @click="del(reply.id)">删除</button>
        </div>
    </div>
</template>
<script setup>
import { ref, inject, onMounted } from 'vue'
const replyInput = ref(null)
const posts = inject('posts')
const props = defineProps({
    id:{
        type: String,
        required: true
    }
})

onMounted(() => {
    getReplys();
})

const post = ref(posts.value.filter(post => post.id == props.id)[0])
const replys = ref([])

/**
 * @description 获取回复
 */
const getReplys = () => {
    fetch('/api/getReply?title=' + encodeURIComponent(post.value.title), {
        method: 'GET'
    })
        .then(response => response.json())
        .then(data => {
            replys.value = data.data;
            replys.value.reverse();
        });
}

/**
 * @description 提交回复
 */
const submitReply = () => {
    const reply = replyInput.value.value;
    if(reply === ''){
        alert('请输入内容');
        return;
    }
    if(document.cookie.indexOf('userInfo') === -1){
        alert('请先登录');
        return;
    }
    const formData = `title=${post.value.title}&reply=${reply}`;
    fetch('/api/addReply', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded'
        },
        body: formData,
        credentials: 'include'
    })
        .then(response => response.json())
        .then(data => {
            getReplys();
        })
        .catch((error) => {
            console.log('失败了?' + error);
        });
}

/**
 * @description 删除回复
 * @param {String} id 回复的id，不是文章的id
 */
const del = (id) => {
    fetch('/api/delReply?id='+id, {
        method: 'GET',
        credentials: 'include'
    })
        .then(response => response.json())
        .then(data => {
            getReplys();
        })
        .catch((error) => {
            console.log('失败了?' + error);
        });
}
</script>

<style scoped>
.article{
    background: rgba(255, 255, 255, 0.374);
    padding: 10px;
    backdrop-filter: blur(2px);
}

.articleTime{
    text-align: right;
}

</style>