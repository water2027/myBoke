<template>
    <div>
        <div v-for="friend in myFriends" class="postLink" :key="friend.id">
            <a :href="friend.link" target="_blank">
                <h3>{{ friend.name }}</h3>
                <p>{{ friend.desc }}</p>
            </a>
        </div>
    </div>
</template>
<script setup>
import { ref, onMounted } from 'vue';

const myFriends = ref([]);

const getMyFriends = () => {
    myFriends.value = [
        {
            id: 1,
            name: 'water2027',
            desc: '本项目链接',
            link: 'https://github.com/water2027/myBoke'
        },
        {
            id: 2,
            name: 'steamedfresh',
            desc: 'steamedfresh，我的一个同学',
            link: 'https://steamedfresh.online/'
        },
    ]
    //fetch操作，暂时不写
}

onMounted(() => {
    getMyFriends()
})

</script>
<style scoped>
p {
    background: linear-gradient(to right, #3ac5d5, #00ccfa);
    background-clip: text;
    color: transparent;
}
</style>