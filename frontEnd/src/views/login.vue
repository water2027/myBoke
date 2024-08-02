<template>
    <div class="pageWithLoginButton">
        <span class="title">{{ isLogin ? 'Login' : 'Register' }}</span>
        <div v-if="isLogin" class="loginAndRegPage">
            <div class="inputData">
                <input ref="email" type="email" id="loginEmail" required>
                <div class="underline"></div>
                <label for="loginEmail">邮箱</label>
            </div>
            <div class="inputData">
                <input ref="password" type="password" id="loginpwd" required>
                <div class="underline"></div>
                <label for="loginpwd">密码</label>
            </div>
            <input type="checkbox" ref="remembered" class="checkbox">记住我
        </div>
        <div v-else class="loginAndRegPage">
            <div class="inputData">
                <input ref="email" type="email" id="regEmail" required>
                <div class="underline"></div>
                <label for="regEmail">邮箱</label>
            </div>
            <div class="inputData">
                <input ref="username" type="text" id="username" required>
                <div class="underline"></div>
                <label for="username">用户名</label>
            </div>
            <div class="inputData">
                <input ref="password" type="password" id="regpwd" required>
                <div class="underline"></div>
                <label for="regpwd">密码</label>
            </div>
            <div class="inputData">
                <input ref="code" type="text" required>
                <div class="underline"></div>
                <label>验证码</label>
            </div>
            <button @click="sendCode" class="codeButton">发送验证码</button>
        </div>
        <div class="buttonDiv">
            <div>
                <button @click="login">{{ isLogin ? '登录' : '已有账号？' }} </button>
            </div>
            <div>
                <button @click="register">{{ isLogin ? '还没账号？' : '注册' }}</button>
            </div>
        </div>

    </div>
</template>
<script setup>
import { ref, onMounted, computed } from 'vue'
import { useRouter } from 'vue-router'
const email = ref(null)
const username = ref(null)
const password = ref(null)
const router = useRouter()
const isLogin = ref(true)
const code = ref(null)
const remembered = ref(false)
onMounted(() => {
    if (document.cookie.indexOf('userInfo') != -1) {
        router.push('/')
    }
})

/**
 * @description 判断登录是否符合要求
 */
const loginIsOk = computed(() => {
    return email.value.value && password.value.value
})

/**
 * @description 判断注册是否符合要求
 */
const regIsOk = computed(() => {
    return email.value.value && username.value.value && password.value.value
})

/**
 * @description 登录
 */
const login = () => {
    if (isLogin.value) {
        if (!loginIsOk.value) {
            if (!email.value.value) {
                alert('email is required')
            }
            else if (!password.value.value) {
                alert('password is required')
            }
            return
        }

        //fetch function
        let formData = `email=${encodeURIComponent(email.value.value)}&password=${encodeURIComponent(password.value.value)}`;
        fetch('/api/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/x-www-form-urlencoded'
            },
            credentials: 'include',
            body: formData
        }).then(res => res.json()).then(data => {
            if (data.code == 1) {
                if (remembered.value) {
                    localStorage.setItem('loginData', formData)
                }
                router.push('/')
            }
            else {
                alert(data.message)
            }
        })

    }
    else {
        isLogin.value = true
    }
}

/**
 * @description 注册
 */
const register = () => {
    if (isLogin.value) {
        isLogin.value = false
    }
    else if (regIsOk.value) {
        //register success
        //fetch function
        let formData = `code=${encodeURIComponent(code.value.value)}`;
        fetch("/api/register", {
            method: 'POST',
            headers: {
                'Content-Type': 'application/x-www-form-urlencoded',
            },
            body: formData,
            credentials: 'include'
        })
            .then(res => res.json())
            .then(data => {
                if (data.code == 1) {
                    router.push('/')
                } else {
                    alert(data.message)
                }
            })
            .catch(error => {
                console.error('There was a problem with your fetch operation:', error);
            });
    }
    else {
        if (!email.value) {
            alert('email is required')
        }
        else if (!username.value) {
            alert('username is required')
        }
        else if (!password.value) {
            alert('password is required')
        }
    }
}

/**
 * @description 发送验证码
 */
const sendCode = () => {
    if ((!email.value.value) || (!username.value.value) || (!password.value.value)) {
        alert('email, username, password is required')
        return
    }
    let formData = `email=${encodeURIComponent(email.value.value)}&username=${encodeURIComponent(username.value.value)}&password=${encodeURIComponent(password.value.value)}`;
    fetch("/api/sendCode", {
        method: 'POST',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
        },
        body: formData,
        credentials: 'include'
    })
        .then(res => res.json())
        .then(data => {
            alert(data.message)
        })
}

</script>

<style scoped>
.title {
    font-size: 30px;
    margin-bottom: 20px;
}

.pageWithLoginButton {
    width: 50%;
    margin: auto;
    margin-top: 15%;
    padding-top: 5%;
    padding-bottom: 5%;
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
    background-color: rgba(255, 255, 255, 0.472);
    backdrop-filter: blur(4px);
    border-radius: 10px;
}

span::before {
    content: '';
    display: inline-block;
    width: 0 !important; 
}

@media screen and (max-width: 768px) {
    .pageWithLoginButton {
        width: 100%;
        margin-bottom: 50px;
    }
}

.loginAndRegPage {
    width: 60%;
    padding: 40px;

}

.loginAndRegPage .inputData {
    position: relative;
    width: 100%;
    height: 40px;
    margin-bottom: 15px;
}

.loginAndRegPage .inputData input {
    width: 100%;
    height: 100%;
    border: none;
    font-size: 17px;
    border-bottom: 2px solid #c0c0c0;
    background-color: rgba(255, 255, 255, 0.472);
}


.inputData input:valid~label,
.inputData input:focus~label {
    transform: translateX(-48px) !important;
    font-size: 16px;
    color: #eb6b26;
    font-weight: bold;
}

.inputData label {
    position: absolute;
    bottom: 10px;
    left: 0px;
    color: #808080;
    pointer-events: none;
    transition: all 0.3s ease;
}

.underline{
    position: absolute;
    bottom: 0;
    height: 2px;
    width: 100%;
    background: linear-gradient(90deg, #eb6b26, #eb6b26);
    transform: scaleX(0);
    transition: all 0.3s ease;
}

.inputData input:focus~.underline,
.inputData input:valid~.underline{
    transform: scaleX(1);
}

.codeButton{
    margin-top: 20px;
}

.buttonDiv {
    display: flex;
    justify-content: space-between;
    width: 100%;
    margin-top: 10px;
}

.buttonDiv div {
    margin-left: 12%;
    margin-right: 12%
}

.checkbox {
    margin-top: 10px;
    
}
</style>