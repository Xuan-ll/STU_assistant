<template>
    <div class="home-container">
        <el-container>
            <el-aside class="aside">
                <div class="user">
                    <el-icon color="#E5EAF3" :size="40" @click="turn2Login">
                        <User />
                    </el-icon>
                    <div class="user-name">{{ userData.user_name }}</div>
                </div>
                <el-divider />
                <div class="list">
                    <div class="list-login">
                        <el-button text class="but" @click="turn2Login">
                            <el-icon color="#E5EAF3" :size="35">
                                <Position />
                            </el-icon>
                            <div class="but-text">返回登录</div>
                        </el-button>
                    </div>
                    <div class="list-note">
                        <el-button text class="but" @click="switch2Note">
                            <el-icon color="#E5EAF3" :size="35">
                                <Memo />
                            </el-icon>
                            <div class="but-text">备忘录</div>
                        </el-button>
                    </div>
                    <div class="list-schedule">
                        <el-button text class="but" @click="switch2Schedule">
                            <el-icon color="#E5EAF3" :size="35">
                                <Calendar />
                            </el-icon>
                            <div class="but-text">课表</div>
                        </el-button>
                    </div>
                </div>
            </el-aside>
            <el-container>
                <el-main class="main">
                    <component :is="componentnext" @transmit="getMessage"></component>
                </el-main>
            </el-container>
        </el-container>
    </div>
</template>
<script>
import Note from '@/components/Note.vue';
import Schedule from '@/components/Schedule.vue';
import { getToken } from '@//utils/api/api';
export default {
    name: 'HomeView',
    components: {
        Note,
        Schedule,
    },
    data() {
        return {
            componentnext: 'Note',
            userData: JSON.parse(localStorage.getItem('user')),
            componentnext: 'Note',
            intervalId: null,
        }
    },
    created() {
        this.auroRefreshToken();
    },
    beforeUnmount() {
        this.stopAutoRefresh(); // 组件销毁时停止自动刷新
    },
    methods: {
        switch2Note() {
            this.componentnext = 'Note';
        },
        switch2Schedule() {
            this.componentnext = 'Schedule';
        },
        turn2Login() {
            this.$router.push({ path: '/login' });
        },
        auroRefreshToken() {
            // 设置每 10 秒刷新一次
            this.intervalId = setInterval(() => {
                this.refreshToken();
            }, 29 * 60 * 1000); // 1000 毫秒 = 1 秒 
        },
        stopAutoRefresh() {
            // 清除定时器
            if (this.intervalId) {
                clearInterval(this.intervalId);
                this.intervalId = null;
            }
        },
        refreshToken() {
            getToken().then(response => {
                // console.log(response);
                if(!response.data.token){
                    console.error('Get refresh token failed.');
                    this.$router.push({ path: '/login' });// 重定向到登录页
                }else{
                    const token = response.data.token;
                    localStorage.setItem('token', token);
                }
            }).catch (refreshError => {
                // 如果刷新失败，处理错误（如跳转到登录页面）
                console.error('Refresh token failed:', refreshError);
                localStorage.clear(); // 清除 Token
                this.$router.push({ path: '/login' });// 重定向到登录页
            })
        }
    },
}
</script>

<style scoped>
.home-container {
    display: flex;
    height: 100%;
    width: 100%;
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

.aside {
    width: 400px;
    height: 100%;
    /* background-color: #0A0A0A; */
    background: rgba(30, 31, 34, 0.6);
    border: 1px solid #404248;
    box-shadow: 0px 2px 8px 0px rgba(0, 0, 0, 0.30);
    /* 设置高斯模糊效果 */
    backdrop-filter: blur(10px);
    /* 隐藏超出视口的部分 */
    overflow: hidden;
}

.user {
    height: 50px;
    margin-top: 30px;
    margin-left: 20px;
    display: flex;
    align-items: center;
}

.user-name {
    color: #E5EAF3;
    margin-left: 110px;
    font-size: 25px;
    user-select: none;
}

.list {
    height: 600px;
    margin-top: 100px;
    display: flex;
    flex-direction: column;
}

.list-note,
.list-schedule,
.list-login {
    margin-top: 50px;
    display: flex;
    align-items: center;
}

.but {
    /* margin-left: 15px; */
    width: 100%;
    color: #E5EAF3;
    font-size: 25px;
    display: flex;
    justify-content: flex-start;
    /* 图标和文字靠左对齐 */
    gap: 10px;
    /* 图标和文字之间的间距 */
}

::v-deep(.but:hover) {
    background-color: #1b1b1b !important;
    /* 强制覆盖 */
    color: rgb(228, 228, 228);
    /* 如果需要，也可以修改文字颜色 */
    transition: background-color 0.3s ease;
    /* 增加过渡效果 */
}

.but-text {
    padding-left: 120px;
}

.main {
    height: 100%;
    background: rgba(30, 31, 34, 0.9);
    box-shadow: 0px 2px 8px 0px rgba(0, 0, 0, 0.90);
    /* 设置高斯模糊效果 */
    backdrop-filter: blur(15px);
    /* 隐藏超出视口的部分 */
    overflow: hidden;
}
</style>