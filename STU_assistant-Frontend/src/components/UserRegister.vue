<template>
    <div class="container-register">
        <el-form ref="ruleFormRef" style="max-width: 500px" size="large" :model="ruleForm" :rules="rules" label-position="right" label-width="auto" class="demo-ruleForm">
            <el-form-item label="昵称:" prop="name">
                <el-input style="width: 300px" v-model="ruleForm.name" placeholder="请输入昵称" clearable />
            </el-form-item>
            <el-form-item label="密码:" prop="password">
                <el-input style="width: 300px" v-model="ruleForm.password" type="password" placeholder="请输入密码" clearable
                    show-password />
            </el-form-item>
            <el-form-item label="确认密码:" prop="password_confirm">
                <el-input style="width: 300px" v-model="ruleForm.password_confirm" type="password" placeholder="请确认密码" clearable
                    show-password />
            </el-form-item>
        </el-form>
        <el-form class="form2">
            <div class="but1">
                <el-button type="primary" style="width: 100px" size="large" @click="switch2Login">返回登录</el-button>
            </div>
            <div class="but2">
                <el-button style="width: 100px" size="large" @click="submitForm(ruleForm)">注册</el-button>
            </div>
        </el-form>
    </div>
</template>

<script>
import { userRegister } from '@/utils/api/api';
import { ref, reactive } from 'vue'
import { ElNotification } from 'element-plus'

export default {
    name: 'UserRegister',
    data() {
        return {
            ruleForm: reactive({
                name: '',
                password: '',
                password_confirm: '',
            }),
            ruleFormRef: ref(null), // 初始化为 ref
            rules: {
                name: [{ validator: this.validatename, trigger: 'blur' }],
                password: [{ validator: this.validatePassword, trigger: '' }],
                password_confirm: [{ validator: this.validatePassword, trigger: '' }],
            },
            switch2login: false,
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
                callback(new Error('密码不能为空'))
            } else {
                this.showErrmsg = false;
                callback();
            }
        },
        submitForm() {
            this.$refs.ruleFormRef.validate((valid) => {
                if (valid) {
                    if(this.ruleForm.password !== this.ruleForm.password_confirm){
                        ElNotification({
                                title: 'Error',
                                message: '两次输入密码需一致',
                                type: 'error',
                        })
                        return ;
                    }
                    userRegister(this.ruleForm.name, this.ruleForm.password, this.ruleForm.password_confirm).then(response => {
                        if (response.status === 200) {
                            ElNotification({
                                title: 'Success',
                                message: '注册成功，点击返回登录',
                                type: 'success',
                            })
                        } else {
                            ElNotification({
                                title: 'Error',
                                message: '注册失败',
                                type: 'error',
                            })
                        }
                    }).catch(error => {
                        console.log(error);
                    })
                } else {
                    console.log('请检查输入')
                }
            })
        },
        switch2Login(){
            // transmit为父组件定义的监听器
            // this.allowLogin为发送的数据
            this.$emit('transmit', this.switch2login)
        }
    },
}
</script>

<style scoped>
.container-register {
    margin-left: 72px;
    height: 100%;
    width: 100%;
    display: flex;
    flex-direction: column;
    /* align-items: center; */
    /* justify-content: center; */
}
.form2 {
    margin-top: 5px;
    display: flex;
}

.but1 {
    margin-left: 70px;
}

.but2 {
    padding-left: 100px;
}

</style>