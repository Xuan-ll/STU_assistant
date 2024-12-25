<template>
    <div class="container-defaultLogin">
        <el-form ref="ruleFormRef" style="max-width: 500px" size="large" :model="ruleForm" :rules="rules"
            label-width="auto" class="demo-ruleForm">
            <el-form-item label="昵称:" prop="name">
                <el-input style="width: 300px" v-model="ruleForm.name" placeholder="请输入昵称" clearable />
            </el-form-item>
            <el-form-item label="密码:" prop="password">
                <el-input style="width: 300px" v-model="ruleForm.password" type="password" placeholder="请输入密码" clearable
                    show-password />
            </el-form-item>
            <transition name="slide-down">
                    <div class="errMsg" v-if="showErrmsg">{{ errMessage }}</div>
            </transition>
        </el-form>
        <el-form class="form2">
            <div class="but1">
                <el-button type="primary" style="width: 100px" size="large" @click="submitForm(ruleForm)">登录</el-button>
            </div>
            <div class="but2">
                <el-button style="width: 100px" size="large" @click="switch2Register">注册</el-button>
            </div>
        </el-form>
    </div>
</template>

<script>
import { userLogin } from '@/utils/api/api';
import { ref, reactive } from 'vue'

export default {
    name: 'DefaultLogin',
    data() {
        return {
            ruleForm: reactive({
                name: '',
                password: '',
            }),
            ruleFormRef: ref(null), // 初始化为 ref
            rules: {
                name: [{ validator: this.validatename, trigger: 'blur' }],
                password: [{ validator: this.validatePassword, trigger: '' }],
            },
            errMessage: '登录失败，请检查用户名/密码，或进行注册',
            showErrmsg: false,
            switch2register: true,
        }
    },
    methods: {
        validatename(rule, value, callback) {
            if (value === '') {
                callback(new Error('昵称不能为空'));
            } else if (!/^[\u4e00-\u9fa5a-zA-Z0-9_]+$/.test(value)) {
                callback(new Error('用户名只能包含中文、英文、数字和下划线'));
            } else {
                callback();
            }
        },
        validatePassword(rule, value, callback) {
            if (value === '') {
                this.showErrmsg = false;
                callback(new Error('密码不能为空'))
            } else {
                this.showErrmsg = false;
                callback();
            }
        },
        submitForm() {
            this.$refs.ruleFormRef.validate((valid) => {
                if (valid) {
                    this.showErrmsg = false;
                    userLogin(this.ruleForm.name, this.ruleForm.password).then(response => {
                        if (response.status === 200) {
                            const token = response.data.token;
                            const user = response.data.user.user_detail;
                            localStorage.setItem('token', token);
                            localStorage.setItem('user', JSON.stringify(user));
                            this.$router.push({ path: '/home' });
                        } else {
                            this.showErrmsg = true;
                        }
                    }).catch(error => {
                        this.showErrmsg = true;
                        console.log(error);
                    })
                } else {
                    console.log('请检查输入')
                }
            })
        },
        switch2Register(){
            // transmit为父组件定义的监听器
            // this.allowLogin为发送的数据
            this.$emit('transmit', this.switch2Register)
        }
    },
}
</script>

<style scoped>
.container-defaultLogin {
    margin-left: 100px;
    height: 100%;
    width: 100%;
    display: flex;
    flex-direction: column;
    /* align-items: center; */
    /* justify-content: center; */
}
.form2 {
    margin-top: 15px;
    display: flex;
}

.but1 {
    margin-left: 40px;
}

.but2 {
    padding-left: 102px;
}

/* 定义滑入效果的动画 */
.slide-down-enter-active {
    animation: slideDown 0.3s ease-out;
    /* 进入动画持续 0.3s */
}

.slide-down-leave-active {
    animation: slideUp 0.3s ease-out;
    /* 离开动画持续 0.3s */
}

/* 自定义从上到下的滑入动画 */
@keyframes slideDown {
    from {
        transform: translateY(-10px);
        opacity: 0;
    }

    to {
        transform: translateY(0);
        opacity: 1;
    }
}

/* 自定义向上消失的动画 */
@keyframes slideUp {
    from {
        transform: translateY(0);
        opacity: 1;
    }

    to {
        transform: translateY(-10px);
        opacity: 0;
    }
}

.errMsg {
    color: #ff4d4f;
    /* 错误提示红色 */
    font-size: 12px;
    height: 10px;
    margin-top: -10px;
    margin-left: 45px;
}
</style>