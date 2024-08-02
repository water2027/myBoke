<template>
    <div>
        <h2>chatRoom</h2>
        <div class="chatRoomContent">
            <div class="chatContent" v-for="item in chatContent" :key="item.id" :title="item.created_at">
                <div class="chatContentHeader">
                    <span>{{ item.username }}</span>
                </div>
                <p class="chatContentBody">{{ item.message }}</p>
                <span>{{ item.time }}</span>
            </div>
        </div>
        <div class="sendMsgDiv">
            <textarea v-model="msg" placeholder="saySomething"></textarea>
            <button @click="sendMsg">send</button>
        </div>
    </div>
</template>
<script setup>
import { ref, onMounted } from 'vue'
const chatContent = ref([])
const ws = ref(null)
const msg = ref('')
onMounted(() => {
    ws.value = new WebSocket('wss://watering.top/api/chatroom')
    ws.value.onopen = () => {
        console.log('open')
    }
    ws.value.onmessage = (e) => {
        getMsg(JSON.parse(e.data))
    }
})

const sendMsg = () => {
    if (msg.value) {
        ws.value.send(JSON.stringify({
            'message': msg.value,
        }))
        msg.value = ''
    }
}

let getMsg = (e) => {
    if (e.username === "system") {
        chatContent.value = e.formerMsg
        getMsg = (e) => {
            chatContent.value.push(e)
        }
    }
}
</script>

<style scoped>
.chatRoomContent {
    height: 500px;
    overflow: auto;
    background-color: rgba(240, 248, 255, 0.622);
    margin-bottom: 10px;
    border-radius: 10px;
}

.chatContent {
    margin: 10px;
    padding: 10px;
}

.chatContentHeader {
    display: flex;
    justify-content: space-between;
}

.chatContentBody {
    margin: 10px;
}

.sendMsgDiv {
    display: flex;
    justify-content: space-between;
}

.sendMsgDiv textarea {
    width: 80%;
    height: 100%;
    border: none;
    font-size: 17px;
    border-bottom: 2px solid #c0c0c0;
    background-color: rgba(255, 255, 255, 0.472);
}

.sendMsgDiv button {
    width: 20%;
    font-size: 32px;
    border: none;
    cursor: pointer;

}
</style>