<template>
    <div>
        <div style="height: 50px">
            <a-button type="primary" @click="openAdd" style="float: right">
                <a-icon type="plus" />
                新增
            </a-button>
        </div>
        <div>
            <a-table :columns="columns" :data-source="tableList" rowKey="ID" :pagination="pagination" @change="findUserList">
                <template slot="authorityId" slot-scope="text, li">
                    <span v-for="al in authorityList" :key="al.ID">
                        <span v-if="al.authorityId === li.authorityId">{{al.authorityName}}</span>
                    </span>
                </template>
                <span slot="action" slot-scope="text, li">
                  <a @click="openEdit(li)">编辑</a>
                  <a-divider type="vertical" />
                  <a @click="openDelete(li)">删除</a>
                </span>
            </a-table>
        </div>

        <a-modal v-model="addVisible" title="新增用户" @ok="submitAdd">
            <a-form :form="addForm" :label-col="{ span: 5 }" :wrapper-col="{ span: 12 }">
                <a-form-item label="用户名">
                    <a-input v-decorator="[ 'username', { rules: [{ required: true, message: '用户名不能为空!' }] }, ]"
                             type="text"
                             placeholder="请输入用户名" class="input_width_use">
                    </a-input>
                </a-form-item>
                <a-form-item label="昵称">
                    <a-input v-decorator="[ 'nickname', { rules: [{ required: true, message: '昵称不能为空!' }] }, ]"
                             type="text"
                             placeholder="请输入昵称" class="input_width_use">
                    </a-input>
                </a-form-item>
                <a-form-item label="密码">
                    <a-input v-decorator="[ 'password', { rules: [{ required: true, message: '密码不能为空!' }] }, ]"
                             type="password"
                             placeholder="请输入密码" class="input_width_use">
                    </a-input>
                </a-form-item>
                <a-form-item label="角色">
                    <a-select  style="width: 320px"
                               placeholder="请输入角色"
                               v-decorator="[ 'authorityId', { rules: [{ required: true, message: '角色不能为空!' }] }, ]"
                               @change="addAuthorityChange">
                        <a-select-option v-for="li in authorityList" :key="li.authorityId">
                            {{ li.authorityName }}
                        </a-select-option>
                    </a-select>
                </a-form-item>
            </a-form>
        </a-modal>

        <a-modal v-model="editVisible" title="编辑用户" @ok="submitEdit">
            <a-form :form="editForm" :label-col="{ span: 5 }" :wrapper-col="{ span: 12 }">
                <a-input v-decorator="['ID']"
                         type="text" class="input_width_use" style="display: none">
                </a-input>
                <a-form-item label="用户名">
                    <a-input v-decorator="[ 'username', { rules: [{ required: true, message: '用户名不能为空!' }] }, ]"
                             type="text"
                             :disabled="true"
                             placeholder="请输入用户名" class="input_width_use">
                    </a-input>
                </a-form-item>
                <a-form-item label="昵称">
                    <a-input v-decorator="[ 'nickname', { rules: [{ required: true, message: '昵称不能为空!' }] }, ]"
                             type="text"
                             placeholder="请输入昵称" class="input_width_use">
                    </a-input>
                </a-form-item>
                <a-form-item label="密码">
                    <a-input v-decorator="[ 'password', { rules: [{ required: true, message: '密码不能为空!' }] }, ]"
                             type="password"
                             placeholder="请输入密码" class="input_width_use">
                    </a-input>
                </a-form-item>
                <a-form-item label="角色">
                    <a-select  style="width: 320px"
                               placeholder="请输入角色"
                               v-decorator="[ 'authorityId', { rules: [{ required: true, message: '角色不能为空!' }] }, ]"
                               @change="editAuthorityChange">
                        <a-select-option v-for="li in authorityList" :key="li.authorityId">
                            {{ li.authorityName }}
                        </a-select-option>
                    </a-select>
                </a-form-item>
            </a-form>
        </a-modal>
    </div>
</template>

<script>
    import api from '../../utils/api'

    export default {
        name: "UserManagement",
        data() {
            return {
                columns: [
                    {
                        title: '用户名',
                        dataIndex: 'userName',
                        key: 'userName',
                    },
                    {
                        title: '昵称',
                        dataIndex: 'nickName',
                        key: 'nickName',
                    },
                    {
                        title: '角色',
                        dataIndex: 'authorityId',
                        key: 'authorityId',
                        scopedSlots: { customRender: 'authorityId' }
                    },
                    {
                        title: '操作',
                        key: 'action',
                        scopedSlots: { customRender: 'action' },
                    },
                ],
                authorityList: [],
                tableList: [],
                pagination: {
                    total: 0,
                    defaultCurrent: 1,
                    defaultPageSize: 10,
                    showSizeChanger: true,
                    pageSizeOptions: ['10', '20', '50', '100'],
                    onShowSizeChange: (current, pageSize) => {
                        this.pagination.defaultCurrent = current;
                        this.pagination.defaultPageSize = pageSize;
                    },
                    onChange: (current, pageSize) => {
                        this.pagination.defaultCurrent = current;
                        this.pagination.defaultPageSize = pageSize;
                    },
                },
                page: 1,
                pageSize: 10,
                addVisible: false,
                addForm: this.$form.createForm(this, { name: 'addUser' }),
                addAuthorityIds: [],

                editVisible: false,
                editData: '',
                editForm: this.$form.createForm(this, { name: 'editUser' }),
                editAuthorityIds: [],

                oldPassword: '',
            }
        },
        created() {
            this.findUserList();
            this.findAuthorityList();
        },
        methods: {
            findUserList() {
                const params = {};
                params.page = this.pagination.defaultCurrent;
                params.pageSize = this.pagination.defaultPageSize;
                this.$axios.post(api.GET_USER_LIST.url, params).then(res => {
                    if (res.data.code === 0) {
                        this.tableList = res.data.data.list;
                        this.pagination.total = res.data.data.total;
                    } else {
                        this.$message.error(res.data.msg);
                        this.pagination.total = 0;
                    }
                })
            },
            // 获取角色列表
            findAuthorityList() {
                const params = {};
                params.page = 1;
                params.pageSize = 9999;
                this.$axios.post(api.GET_AUTHORITY_LIST.url, params).then(res => {
                    if (res.data.code === 0) {
                        this.authorityList = res.data.data.list;
                    } else {
                        this.$message.error(res.data.msg);
                    }
                })
            },
            // 新增用户
            openAdd() {
                this.addForm.resetFields();
                this.addVisible = true;
            },
            addAuthorityChange(value) {
                this.addAuthorityIds = [value];
            },
            submitAdd(e) {
                e.preventDefault();
                this.addForm.validateFields((err, values) => {
                    values.authorityIds = this.addAuthorityIds;
                    if (!err) {
                        this.$axios.post(api.ADD_USER.url, values).then(res => {
                            if (res.data.code === 0) {
                                this.$message.success('新增成功！');
                                this.addForm.resetFields();
                                this.addVisible = false;
                                this.findUserList();
                            } else {
                                this.$message.error(res.data.msg);
                            }
                        })
                    }
                });
            },
            // 编辑用户
            openEdit(data) {
                this.editForm.resetFields();
                this.editVisible = true;
                this.editData = data;

                this.oldPassword = data.password;

                this.$nextTick(() => {
                    this.editForm.setFieldsValue({
                        ID: data.ID,
                        username: data.userName,
                        nickname: data.nickName,
                        password: data.password,
                        authorityId: data.authorityId
                    });
                });
                this.editAuthorityIds = [data.authorityId];
            },
            editAuthorityChange(value) {
                this.editAuthorityIds = [value];
            },
            submitEdit(e) {
                e.preventDefault();
                this.editForm.validateFields((err, values) => {
                    values.ID = this.editData.ID;
                    values.authorityIds = this.editAuthorityIds;

                    if (this.oldPassword === values.password) {
                        values.password = null;
                    }
                    if (!err) {
                        this.$axios.put(api.EDIT_USER.url, values).then(res => {
                            if (res.data.code === 0) {
                                this.$message.success('编辑成功！');
                                this.editForm.resetFields();
                                this.editVisible = false;
                                this.findUserList();
                            } else {
                                this.$message.error(res.data.msg);
                            }
                        })
                    }
                });
            },
            // 删除用户
            openDelete(data) {
                this.$confirm({
                    title: '提示信息：',
                    content: '确定删除数据？',
                    onOk: () => this.submitDelete(data)
                });
            },
            submitDelete(data) {
                this.$axios.delete(api.DELETE_USER.url, {data: {ID:  data.ID}}).then(res => {
                    if (res.data.code === 0) {
                        this.$message.success('删除成功！');
                    } else {
                        this.$message.error(res.data.msg);
                    }
                    this.findUserList();
                })
            },
        }
    }
</script>

<style xml:lang="less">
    .input_width_use {
        width: 320px;
    }
</style>
