<template>
    <div class="container">
        <div class="login-container">
            <div class="login-bg"></div>
            <div class="login">
                <div class="text">Welcome Back :)</div>
                <div class="form">
                    <component :is="componentnext" @transmit="getMessage"></component>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import DefaultLogin from '@/components/DefaultLogin.vue';
import UserRegister from '@/components/userRegister.vue';
export default {
    name: 'UserLogin',
    components:{
        DefaultLogin,
        UserRegister,
    },
    data() {
        return {
            register: false,
            componentnext: 'DefaultLogin',
        }
    },
    methods: {
        getMessage(data){
            this.register = data;
            if(this.register){
                this.componentnext = 'UserRegister'
            }else{
                this.componentnext = 'DefaultLogin'
            }
        }
    },
}
</script>

<style scoped>
.container {
    width: 100%;
    height: 100%;
    top: 0;
    left: 0;
    display: flex;
    justify-content: center;
    /* 水平居中 */
    align-items: center;
    /* 垂直居中 */
    position: fixed;
    /* 固定位置 */
    background-size: cover;
    /* 背景图片覆盖容器 */
    background-position: center;
    /* 背景图片居中对齐 */
    background-repeat: no-repeat;
    /* 背景图片不重复 */
    overflow: hidden;
    /* 隐藏背景区域超出部分 */
    background-color: black;
    background-image: url('/background.png');
}

.login-container {
    display: flex;
    /* 启用 Flexbox 布局 */
    width: 1200px;
    height: 500px;
    position: fixed;
    background: rgba(255, 255, 255, 0.95);
    border: 1px solid #404248;
    border-radius: 8px;
    box-shadow: 0px 2px 8px 0px rgba(0, 0, 0, 0.30);
    /* 设置高斯模糊效果 */
    backdrop-filter: blur(10px);
    /* 隐藏超出视口的部分 */
    overflow: hidden;
}

.login-bg {
    width: 50%;
    height: 100%;
    background-position: center;
    /* 背景图片居中对齐 */
    background-repeat: no-repeat;
    /* 背景图片不重复 */
    overflow: hidden;
    background-image: url('/login-bg.png');
}

.login {
    width: 50%;
    /* 右侧占据 40% */
    height: 100%;
    margin: 0, auto;
    /* 水平居中 */
    display: flex;
    flex-direction: column;
    /* 子元素垂直排列 */
}

.text {
    text-align: center;
    /* 设置文字水平居中 */
    color: #000;
    font-size: 40px;
    user-select: none;
    margin-top: 90px;
    width: 100%;
    height: 50px;
}
.form {
    height: 90%;
    display: flex;
    margin-top: 40px;
    
}
</style>