<template>
    <div class="edit">
        <nav>
            <div>
                <span @click="toggleCheck(0)">技术</span>
                <div v-if="isChecked == 0" v-for="post in jiuu" :key="post.id">
                    <span @click="get(post.id)">{{ post.title }}</span>
                </div>
            </div>
            <div>
                <span @click="toggleCheck(1)">闲谈</span>
                <div v-if="isChecked == 1" v-for="post in xmtj" :key="post.id">
                    <span @click="get(post.id)">{{ post.title }}</span>
                </div>
            </div>
            <div>
                <span @click="toggleCheck(2)">更新</span>
                <div v-if="isChecked == 2" v-for="post in ggxn" :key="post.id">
                    <span @click="get(post.id)">{{ post.title }}</span>
                </div>
            </div>
        </nav>
        <div class="edit-container">
            <div class="edit-title">
                <input type="text" placeholder="请输入标题" ref="titleInput">
            </div>
            <div class="edit-tag">
                <input type="text" placeholder="请输入标签" ref="tagInput">
            </div>
            <div class="edit-jmjx">
                <textarea placeholder="请输入简介" ref="jmjxInput"></textarea>
            </div>
            <div class="edit-content">
                <textarea v-model="body" placeholder="请输入内容" ref="bodyInput"></textarea>
            </div>
            <div class="buttons">
                <div class="edit-submit">
                    <button @click="submit">提交</button>
                </div>
                <div class="edit-submit">
                    <button @click="change">更改</button>
                </div>
            </div>
        </div>
    </div>
</template>
<script setup>
import { useRouter } from 'vue-router';
import { inject, onMounted, ref } from 'vue';
const titleInput = ref(null)
const tagInput = ref(null)
const jmjxInput = ref(null)
const bodyInput = ref(null)
const jiuu = ref([])
const xmtj = ref([])
const ggxn = ref([])
const router = useRouter()
const allPosts = inject('allPosts')
const token = ref('')
onMounted(() => {
    // jiuu.value = allPosts.filter(post => post.tag === '技术')
    // xmtj.value = allPosts.filter(post => post.tag === '闲谈')
    // ggxn.value = allPosts.filter(post => post.tag === '更新')
    if(localStorage.getItem('token')) {
        token.value = localStorage.getItem('token')
    }
    if(!token.value){
        token.value = prompt('请先登录')
        localStorage.setItem('token', token.value)
    }
    if(!token.value){
        router.push('/login')
    }
})

/**
 * @description 提交文章
 */
const submit = () => {
    //fetch function
    // 发起fetch请求
    // 提交数据
    const title = titleInput.value.value
    const tag = tagInput.value.value
    const jmjx = jmjxInput.value.value
    const body = bodyInput.value.value
    console.log(title, tag, jmjx, body)
    let formData = `title=${encodeURIComponent(title)}&jmjx=${encodeURIComponent(jmjx.value)}&body=${encodeURIComponent(body.value)}&tag=${encodeURIComponent(tag.value)}`;
    fetch("/api/writePost", {
        method: 'POST',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
            'token': token.value
        },
        body: formData
    }).then(res => {
        if (res.ok) {
            alert('提交成功')
        } else {
            alert('提交失败')
        }
    })
}

/**
 * @description 修改文章
 */
const change = () => {
    let formData = `title=${encodeURIComponent(title.value)}&jmjx=${encodeURIComponent(jmjx.value)}&body=${encodeURIComponent(body.value)}&tag=${encodeURIComponent(tag.value)}`;
    fetch("/api/editPost", {
        method: 'POST',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded'
        },
        body: formData
    }).then(res => {
        if (res.ok) {
            alert('修改成功')
        } else {
            alert('修改失败')
        }
    })    
}

/**
 * @description 获取文章
 * @param {String} id 文章id
 */
const get = (id) => {
    var aimPost = allPosts.find(post => post.id === id)
    title.value = aimPost.title
    tag.value = aimPost.tag
    jmjx.value = aimPost.jmjx
    body.value = aimPost.body
}
const isChecked = ref(-1);

/**
 * @description 切换选中
 * @param {Number} index 选中的索引
 */
const toggleCheck = (index) => {
    if (isChecked.value === index) {
        isChecked.value = -1;
    } else {
        isChecked.value = index;
    }
}
</script>
<style scoped>
.edit {
    display: flex;
    flex-direction: column;
    align-items: center;
    margin-top: 20px;
}

nav {
    display: flex;
    justify-content: space-between;
    width: 70%;
    margin-bottom: 20px;
}

nav div {
    display: flex;
    flex-direction: column;
    align-items: center;
}

nav div span {
    cursor: pointer;
    margin-right: 10px;
}

nav div div {
    flex-direction: column;
    align-items: center;
}

nav div input:checked+div {
    display: flex;
}

.edit-container {
    width: 70%;
    display: flex;
    flex-direction: column;
    align-items: center;
}

.edit-title {
    width: 30%;
    margin-bottom: 20px;
}

.edit-title input {
    width: 100%;
    height: 30px;
    padding: 0 10px;
    border: 1px solid #ccc;
    border-radius: 5px;
}

.edit-tag input {
    width: 100%;
    height: 30px;
    padding: 0 10px;
    border: 1px solid #ccc;
    border-radius: 5px;
}

.edit-jmjx input {
    width: 100%;
    height: 30px;
    padding: 0 10px;
    border: 1px solid #ccc;
    border-radius: 5px;
}

.edit-content textarea {
    width: 100%;
    height: 300px;
    padding: 10px;
    border: 1px solid #ccc;
    border-radius: 5px;
}

.edit-submit button {
    width: 100px;
    height: 30px;
    background-color: #409eff;
    color: #fff;
    border: none;
    border-radius: 5px;
    cursor: pointer;
}

.edit-submit button:hover {
    background-color: #66b1ff;
}

.edit-submit button:active {
    background-color: #3a8ee6;
}

.edit-submit button:focus {
    outline: none;
}

.edit-submit button:disabled {
    background-color: #dcdfe6;
    color: #bbb;
    cursor: not-allowed;
}

.edit-submit button:disabled:hover {
    background-color: #dcdfe6;
}

.edit-submit button:disabled:active {
    background-color: #dcdfe6;
}

.edit-submit button:disabled:focus {
    outline: none;
}

.edit-submit button:disabled:disabled {
    background-color: #dcdfe6;
    color: #bbb;
    cursor: not-allowed;
}

.edit-submit button:disabled:disabled:hover {
    background-color: #dcdfe6;
}

.edit-submit button:disabled:disabled:active {
    background-color: #dcdfe6;
}

.edit-submit button:disabled:disabled:focus {
    outline: none;
}

.edit-submit button:disabled:disabled:disabled {
    background-color: #dcdfe6;
    color: #bbb;
    cursor: not-allowed;
}

.edit-submit button:disabled:disabled:disabled:hover {
    background-color: #dcdfe6;
}

.edit-submit button:disabled:disabled:disabled:active {
    background-color: #dcdfe6;
}

.edit-submit button:disabled:disabled:disabled:focus {
    outline: none;
}

.buttons {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
    width: 30%;
    margin-top: 20px;
}

.buttons div {
    margin-right: 10px;
    margin-left: 10px;
}
</style>
