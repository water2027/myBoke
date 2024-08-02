<template>
    <audio controls autoplay loop>
        <source src="./assets/hello.m4a" type="audio/mpeg">
    </audio>
    <div @keyup.ctrl.shift.enter="enterEditView" @click="createHeart" id="body" tabindex="0">
        <header>
            <h1>Water的开发日志</h1>
            <button id="audioPlayer" @click="audioPlay"></button>
            <button v-if="isMobileView" @click="toggleNav" class="hamburgerMenu">
                <span></span>
                <span></span>
                <span></span>
            </button>
            <nav v-else>
                <router-link to="/login" v-if="!isLogin">登录</router-link>
                <router-link to="/">时间轴</router-link>
                <router-link to="/tags">分类</router-link>
                <router-link to="/friends">友链</router-link>
                <router-link to="/chatroom">聊天室</router-link>
                <router-link to="/game">小游戏</router-link>
            </nav>
        </header>
        <div v-if="isNavOpen" class="sidebar">
            <nav @click="toggleNav">
                <router-link to="/login" v-if="!isLogin">登录</router-link>
                <router-link to="/">时间轴</router-link>
                <router-link to="/tags">分类</router-link>
                <router-link to="/friends">友链</router-link>
                <router-link to="/chatroom">聊天室</router-link>
                <router-link to="/game">小游戏</router-link>
            </nav>
        </div>
        <main>
            <section class="leftContent">
                <router-view />
                <nav v-if="isNavShow">
                    <button @click="getMorePosts(0)">
                        < </button>
                            <button @click="getMorePosts(1)">
                                >
                            </button>
                </nav>
            </section>
            <aside>
                <div class="aside">
                    <div id="head">
                        <div class="headImage">
                            <img src="./assets/water.png" alt="Profile picture description">
                        </div>
                        <div id="headInfo">
                            <h3 class="riloText">Water</h3>
                            <span>摆烂，摆大烂</span><br>
                            <button @click="follow">{{ followed ? '取消？' : '订阅？' }}</button>
                        </div>
                    </div>
                    <div>
                        <a href="https://github.com/water2027"><span class="gradientText">我的GitHub</span></a>
                    </div>

                </div>
                <div class="aside">
                    <h3>{{ username }}你好！</h3>
                    <div>
                        <p>
                            {{ userInfo ? userInfo.tip : '你好，欢迎访问我的博客！' }}
                        </p>
                    </div>
                </div>
            </aside>
        </main>
        <footer>
            <a href="https://beian.miit.gov.cn/" target="_blank">
                <span class="github">粤ICP备2024236068号-1</span>
            </a>
        </footer>
    </div>
</template>

<script setup>
import { ref, onMounted, provide, computed } from 'vue';
import { RouterLink, RouterView, useRoute, useRouter } from 'vue-router';

const posts = ref([]);
const allPosts = ref([]);
const len = ref(0);
// watch(allPosts, (val) => {
//     len.value = val.length;
//     firstGetPosts();
//     firstGetPosts = () => { };
// });
const userInfo = ref(null);
const start = ref(0);
const route = useRoute();
const router = useRouter()
const followed = ref(false);
const username = ref('访客')
const isLogin = ref(false);
const isMobileView = window.innerWidth < 768;
const isNavOpen = ref(false);

const toggleNav = () => {
    isNavOpen.value = !isNavOpen.value
};

const audioPlay = () => {
    let audio = document.querySelector('audio');
    if (audio.paused) {
        audio.play();
    } else {
        audio.pause();
    }
};

provide('posts', posts)
provide('allPosts', allPosts)

const enterEditView = () => {
    // if (document.cookie.indexOf('userInfo') == -1) {
    //     alert('请先登录')
    //     return
    // }
    router.push('/edit')
};

const fetchAllPosts = async () => {
    //fetch function
    //流式传输，传过来的是五个一组，全都加进allPosts里
    // 发起fetch请求
    const response = await fetch("/api/getPosts");
    const data = await response.json();
    allPosts.value = Object.freeze(data.data.reverse());
    len.value = data.data.length;
    getPosts();
    // const reader = response.body.getReader();
    // const decoder = new TextDecoder();
    // alert("fetch")
    // console.log(response)

    // // 读取数据
    // reader.read().then(function processStream({ done, value }) {
    //     if (done) {
    //         console.log("Stream complete");
    //         return;
    //     }

    //     // 将Uint8Array转换为字符串
    //     const str = decoder.decode(value, { stream: true });
    //     console.log(str)
    //     try {
    //         // 假设服务器发送的是JSON数组
    //         const array = JSON.parse(str);

    //         allPosts.value = allPosts.value.concat(array);
    //     } catch (e) {
    //         console.error("Failed to parse JSON", e);
    //     }

    //     // 继续读取下一段数据
    //     reader.read().then(processStream);
    // });

}

// let firstGetPosts = () => {
//     posts.value = allPosts.value.slice(start.value, (start.value + 5) < len.value ? (start.value + 5) : len.value);
// }

const getPosts = () => {
    //一次只加载部分文章，默认为5
    posts.value = allPosts.value.slice(start.value, (start.value + 5) < len.value ? (start.value + 5) : len.value);
};

const getMorePosts = (direction) => {
    //API接口应该传入两个个数字参数，第一个是文章的起始位置，第二个是文章的数量
    //使用GET方法请求文章数据
    //然后更新posts.value
    // start.value += 5;
    if (direction === 0) {
        start.value -= 5;
        start.value = start.value < 0 ? 0 : start.value;
    } else {
        start.value += 5;
    }
    getPosts();
};

const getUserInfo = () => {
    let sendUserInfo = (userInfo) => {
        const data = {
            time: userInfo.time,
            week: userInfo.week,
            ip: userInfo.ip,
            location: userInfo.location,
            browser: userInfo.browser,
            deviceSystem: userInfo.system
        }
        fetch('/api/getUserInfo', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)
        })
    };
    fetch('https://api.vvhan.com/api/visitor.info')
        .then(response => response.json())
        .then(data => {
            userInfo.value = data;
            sendUserInfo(data);
        })
        .catch(err => {
            console.log(err)
        })
};

const isNavShow = computed(() => {
    return route.path == '/'
});

onMounted(() => {
    fetchAllPosts();
    getUserInfo();
    let loginData = localStorage.getItem('loginData');
    if (loginData && document.cookie.indexOf('userInfo') == -1) {

        fetch('/api/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/x-www-form-urlencoded'
            },
            credentials: 'include',
            body: loginData
        }).then(res => res.json()).then(data => {
            if (data.code == 1) {
                isLogin.value = true;
            }
            else {
                alert("Oh no 发生了一些错误，自动登录失败了")
            }
        })
    }

});

const createHeart = (e) => {
    let x = e.clientX;
    let y = e.clientY;
    let heart = document.createElement('b');
    heart.style.left = `${x}px`;
    heart.style.top = `${y}px`;
    heart.innerHTML = '❤';
    document.body.appendChild(heart);
    heart.animate([
        { top: `${y}px`, opacity: 1 },
        { top: `${y - 100}px`, opacity: 0 }
    ], {
        duration: 1500,
        easing: 'ease'
    });
    setTimeout(() => {
        heart.remove();
    }, 1000);
};

const follow = () => {
    fetch('/api/follow', {
        method: 'GET',
        credentials: 'include'
    })
        .then(response => {
            if (response.status == 401) {
                alert('请先登录')
                return
            }
            return response.json()
        })
        .then(data => {
            if (data.code == 1) {
                followed.value = true
            }
            else {
                followed.value = false
            }
        })
};

</script>

<style scoped>
.hamburgerMenu {
    display: flex;
    flex-direction: column;
    justify-content: space-around;
    width: 30px;
    height: 30px;
    background: transparent;
    border: none;
    cursor: pointer;
    padding: 0;
    z-index: 9998;
}

.hamburgerMenu span {
    width: 100%;
    height: 3px;
    background-color: #333;
    border-radius: 2px;
    transition: all 0.3s;
}

#audioPlayer {
    /* 三角形播放按钮 */
    margin-left: auto;
    width: 0;
    height: 0;
    border-top: 10px solid transparent;
    border-bottom: 10px solid transparent;
    border-left: 15px solid #333;
    cursor: pointer;
    top: 0;
    right: 0;
    position: relative;
    z-index: 8000;
    margin-right: 20px;
}

audio {
    display: none;
    position: absolute;
    pointer-events: none;
}

.sidebar {
    position: fixed;
    top: 0;
    right: 0;
    width: 100%;
    height: 100%;
    background-color: rgba(255, 255, 255, 0.207);
    box-shadow: 2px 0 5px rgba(0, 0, 0, 0.5);
    z-index: 9000;
    padding: 20px;
    box-sizing: border-box;
    backdrop-filter: blur(2px);
}

.sidebar nav {
    width: 100%;
    margin-top: 10%;
    display: flex;
    flex-direction: column;
    align-items: center;
    font-size: larger;
}

.sidebar nav a {
    margin-bottom: 15px;
    margin-top: 15px;
}

@media (max-width: 767px) {
    nav {
        display: none;
    }

    .hamburger-menu {
        display: flex;
    }
}

@media (min-width: 768px) {
    .hamburger-menu {
        display: none;
    }
}

#body {
    background-color: #f0f0f0;
    width: 100%;
    min-height: 100vh;
    height: 100%;
    margin: 0;
    padding: 0;
    display: flex;
    flex-direction: column;
    align-items: center;
    background-image: url('./assets/wow.jpg');
    background-size: cover;
    background-attachment: fixed;
    background-repeat: no-repeat;
}

header {
    margin-left: 15%;
    margin-right: 15%;
    width: 70%;
    height: 30px;
    margin-top: 10px;
    background-color: rgba(255, 255, 255, 0.3);
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 0 20px;
}

section {
    width: 70%;
}

h1 {
    font-size: 20px;
    text-shadow: 2px 2px 15px #494949;
    color: rgb(248, 248, 248);
}

nav {
    display: flex;
    justify-content: space-between;
    width: 30%;
}

main {
    display: flex;
    justify-content: space-between;
    width: 70%;
    height: 100%;
    padding: 20px;
}

aside {
    width: 25%;
    display: flex;
    flex-direction: column;
}

@media (max-width: 768px) {
    main {
        flex-direction: column;
    }

    aside {
        width: 100%;
    }

    .leftContent {
        width: 100%;
    }
}

.aside {
    width: 100%;
    padding-left: 5%;
    padding-bottom: 1%;
    margin-bottom: 5%;
    display: flex;
    flex-direction: column;
    background-color: rgba(255, 255, 255, 0.7);
    border-radius: 10px;
}

#head {
    display: flex;
    align-items: center;
    margin-bottom: 20px;
}

footer {
    width: 100%;
    height: 50px;
    margin: auto;
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
    bottom: 0;
}

p {
    background: linear-gradient(to right, #3ac5d5, #00ccfa);
    background-clip: text;
    color: transparent;
}

span::before {
    content: '';
    display: inline-block;
    width: 16px !important;
}

</style>