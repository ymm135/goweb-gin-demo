<template>
    <div>
        <a-form :form="form" class="edit_password">
            <a-form-item label="当前密码">
                <a-input v-decorator="['oldPassword',{rules: [{required: true,message: '请输入当前密码',}],},]"
                         type="password"
                         required="true"
                         placeholder="请输入当前密码" class="input_width_pass">
                </a-input>
            </a-form-item>
            <a-form-item label="新密码">
                <a-input v-decorator="['password1',{rules: [{required: true,message: '请输入密码',},{validator: validateToNextPassword}]}]"
                         type="password"
                         placeholder="请输入新密码" class="input_width_pass">
                </a-input>
            </a-form-item>
            <a-form-item label="确认密码">
                <a-input v-decorator="['password2',{rules: [{required: true,message: '请输入密码',},{validator: compareToFirstPassword}]}]"
                         type="password"
                         placeholder="请输入新密码" class="input_width_pass">
                </a-input>
            </a-form-item>
            <a-form-item label="">
                <a-button type="primary" @click="submitPassword" style="margin-left: 210px">
                    确定
                </a-button>
            </a-form-item>
        </a-form>
    </div>
</template>

<script>
    import api from '../../utils/api'

    export default {
        name: "UserPassword",
        data() {
            return {
                // 编辑
                editVisible: false,
                form: this.$form.createForm(this, { name: 'editPassword' }),
                confirmDirty: false,
            }
        },
        methods: {
            compareToFirstPassword(rule, value, callback) {
                const form = this.form;
                if (value && value !== form.getFieldValue('password1')) {
                    callback('两次密码不一致');
                } else {
                    callback();
                }
            },
            validateToNextPassword(rule, value, callback) {
                const form = this.form;
                if (value && this.confirmDirty) {
                    form.validateFields(['password2'], { force: true });
                }
                callback();
            },
            // 修改账户密码提交事件
            submitPassword(e) {
                // 阻止与事件关联的默认动作
                e.preventDefault();
                this.form.validateFields((err, values) => {
                    const loginData = {};
                    loginData.username = sessionStorage.getItem('userName');
                    loginData.password = values.oldPassword;
                    loginData.newPassword = values.password1;
                    if (!err) {
                        this.$axios.post(api.CHANGE_PASSWORD.url, loginData).then(res => {
                            if (res.data.code === 0) {
                                this.form.resetFields();
                                this.$message.success(res.data.msg);
                            } else {
                                this.$message.error(res.data.msg);
                            }
                        })
                    }
                });
            },
        }
    }
</script>

<style xml:lang="less">
    .edit_password > .ant-form-item {
        display: flex;
    }
    .edit_password > .ant-form-item > .ant-form-item-label {
        width: 100px;
    }
    .input_width_pass {
        width: 320px;
    }
</style>
