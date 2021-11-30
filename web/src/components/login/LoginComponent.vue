<template>
    <div class="login_div">
        <div class="login_content">
            <div class="login_title">周报系统</div>
            <a-form :form="form" :label-col="{ span: 5 }" :wrapper-col="{ span: 12 }" @submit="handleSubmit">
                <a-form-item class="login_explain">
                    <a-input v-decorator="['username', { rules: [{ required: true, message: '用户名不能为空!' }] },]"
                             placeholder="请输入用户名" class="input_login" autocomplete="off">
                        <a-icon slot="prefix" type="user" class="icon_color" />
                    </a-input>
                </a-form-item>
                <a-form-item class="login_explain">
                    <a-input v-decorator="[ 'password', { rules: [{ required: true, message: '密码不能为空!' }] }, ]"
                             type="password"
                             placeholder="请输入密码" class="input_login">
                        <a-icon slot="prefix" type="lock" class="icon_color" />
                    </a-input>
                </a-form-item>
                <a-form-item class="login_explain captcha">
                    <a-input v-decorator="[ 'captcha', { rules: [{ required: true, message: '验证码不能为空!' }] }, ]"
                             @keyup.enter="handleSubmit"
                             placeholder="请输入验证码" style="width: 220px; margin-left: 30px" autocomplete="off">
                    </a-input>
                    <img :src="captcha" class="captcha_img" title="点击刷新" @click="getCaptcha" alt="">
                </a-form-item>
                <a-form-item>
                    <a-button type="primary" class="login_button" @click="handleSubmit">
                        登录
                    </a-button>
                </a-form-item>
            </a-form>
        </div>
    </div>
</template>

<script>
    import api from '../../utils/api'

    export default {
        name: "LoginComponent",
        data() {
            return {
                formLayout: 'horizontal',
                captcha: '',
                captchaId: '',
                form: this.$form.createForm(this, { name: 'login' }),
                message: '',
            };
        },
        mounted() {
            // 获取验证码
            this.getCaptcha();
        },
        methods: {
            getCaptcha() {
                this.$axios.post(api.GET_CAPTCHA.url).then(res => {
                    if (res.data.code === 0) {
                        this.captcha = res.data.data.picPath;
                        this.captchaId = res.data.data.captchaId;
                    } else {
                        this.$message.error(res.data.msg);
                    }
                })
            },
            handleSubmit(e) {
                e.preventDefault();
                this.form.validateFields((err, values) => {
                    const loginData = {};
                    // loginData.username = this.cryptos.encrypto(values.username);
                    // loginData.password = this.cryptos.encrypto(values.password);
                    loginData.username = values.username;
                    loginData.password = values.password;
                    loginData.captchaId = this.captchaId;
                    loginData.captcha = values.captcha;
                    if (!err) {
                        this.$axios.defaults.withCredentials = true;
                        this.$axios.post(api.LOGIN.url, loginData).then(res => {
                            if (res.data.code === 0) {
                                // 存登录返回的session
                                sessionStorage.setItem('login', res.data.data.token);
                                sessionStorage.setItem('userId', res.data.data.user.ID);
                                sessionStorage.setItem('userName', res.data.data.user.userName);
                                sessionStorage.setItem('nickName', res.data.data.user.nickName);
                                this.$router.push('/');
                            } else {
                                this.getCaptcha();
                                this.$message.error(res.data.msg);
                            }
                        })
                    }
                });
            },
        },
    }
</script>

<style>
    .login_div {
        background-color: #6495ED;
        height: 100%;
        width: 100%;
        min-height: 100%;
        display: flex;
        align-items: center;
        background-repeat: no-repeat;
        background-position: center;
        background-size: cover;
    }
    .login_content {
        width: 386px;
        margin-left: calc(50% - 193px);
        border-radius: 8px;
        background-color: rgba(255, 255, 255, 0.4);
        text-align: center;
    }
    .login_title {
        font-size: 35px;
        text-align: center;
        margin-bottom: 25px;
        font-weight: 600;
        margin-top: 20px;
    }
    .login_button {
        margin-left: 160px;
    }
    .login_explain > div > div > .ant-form-explain {
        text-align: left;
        margin-left: 30px;
    }
    .captcha > div > div > .ant-form-item-children {
        display: flex;
    }
    .icon_color {
        color: rgba(0,0,0,.25);
    }
    .captcha_img {
        height: 32px;
        margin-left: 30px;
        cursor: pointer;
    }
    .input_login {
        width: 320px;
        margin-left: 30px
    }
</style>
