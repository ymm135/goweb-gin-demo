<template>
    <div class="content_wri">
        <div class="header_write">{{template.header}}</div>
<!--        模板内容-->
        <div v-for="(item, index) in template.contents" :key="index">
            <div class="title_weight">{{item.title}}：</div>
            <quill-editor
                    :id="index"
                    style="margin-top: 20px"
                    v-model="item.content"
                    :ref="index"
                    @blur="onEditorBlur($event, index)" @focus="onEditorFocus($event, index)">
            </quill-editor>
        </div>

        <div class="title_weight">图片上传</div>
        <div>
            <div class="clearfix">
                <a-upload
                        action="/week/fileUploadAndDownload/upload"
                        list-type="picture-card"
                        :file-list="fileList"
                        @preview="handlePreview"
                        @change="handleChange"
                        accept=".png,.jpeg,.jpg"
                        :headers="{ 'x-token': xToken, 'x-user-id':  xUserId}"
                        :remove="removeFile"
                >
                    <div v-if="fileList.length < 8">
                        <a-icon type="plus" />
                        <div class="ant-upload-text">
                            Upload
                        </div>
                    </div>
                </a-upload>
                <a-modal :visible="previewVisible" :footer="null" @cancel="handleCancel">
                    <img alt="example" style="width: 100%" :src="previewImage" />
                </a-modal>
            </div>
        </div>

        <div class="title_weight">文件上传</div>
        <div>
            <div class="clearfix">
                <a-upload
                        action="/week/fileUploadAndDownload/upload"
                        :file-list="fileList2"
                        @change="handleChange2"
                        :headers="{ 'x-token': xToken, 'x-user-id':  xUserId}"
                        :remove="removeFile"
                >
                    <div>
                        <a-button> <a-icon type="upload" /> Upload </a-button>
                    </div>
                </a-upload>
            </div>
        </div>

        <div style="font-weight: 600; margin: 20px 0 0 0"><span style="color: red">*</span>发送给</div>
        <a-select
                mode="multiple"
                style="width: 100%"
                placeholder="请选择人员"
                @change="userChange"
                v-model="sends"
                :class="sendToClass"
        >
            <a-select-option v-for="li in userList" :key="li.ID">
                {{li.nickName}}
            </a-select-option>
        </a-select>
        <span style="color: red" v-if="sendToClass === 'cannotSubmit'">请选择人员</span>

        <div style="text-align: center">
            <a-button type="primary" class="commit_button" @click="addReports"> 提交 </a-button>
        </div>
    </div>
</template>

<script>
    import api from '../../utils/api'
    import { quillEditor } from "vue-quill-editor";
    import 'quill/dist/quill.core.css';
    import 'quill/dist/quill.snow.css';
    import 'quill/dist/quill.bubble.css';

    function getBase64(file) {
        return new Promise((resolve, reject) => {
            const reader = new FileReader();
            reader.readAsDataURL(file);
            reader.onload = () => resolve(reader.result);
            reader.onerror = error => reject(error);
        });
    }
    export default {
        name: "WriteWeeklyReport",
        components: {quillEditor},
        data() {
            return {
                // 图片上传
                fileList: [],
                previewVisible: false,
                previewImage: '',
                xToken: sessionStorage.getItem('login'),
                xUserId: sessionStorage.getItem('userId'),
                pictures: [],
                // 文件上传
                fileList2: [],
                // 周报模板
                template: {},
                // 人员列表
                userList: [],
                sends: [],
                attachments: [],
                sendToClass: 'canSubmit',
                // 发送到(选中人员列表)
                sendTo: [],
            };
        },
        created() {
            // 查询模板
            this.findTempl();
            // 查询人员列表
            this.findUserList();
        },
        methods: {
            // 图片上传
            handleCancel() {
                this.previewVisible = false;
            },
            async handlePreview(file) {
                if (!file.url && !file.preview) {
                    file.preview = await getBase64(file.originFileObj);
                }
                this.previewImage = file.url || file.preview;
                this.previewVisible = true;
            },
            handleChange({ fileList }) {
                this.fileList = fileList;
            },
            // 文件上传
            handleChange2({ fileList }) {
                this.fileList2 = fileList;
            },
            // 删除文件
            removeFile(file) {
                const params = {};
                params.key = file.response.data.file.key;

                const files = [];
                this.fileList.forEach(li => {
                    if (li.response.data.file.key !== params.key) {
                        files.push(li);
                    }
                });
                this.fileList = files;

                const files2 = [];
                this.fileList2.forEach(li => {
                    if (li.response.data.file.key !== params.key) {
                        files2.push(li);
                    }
                });
                this.fileList2 = files2;

                this.$axios.post(api.DELETE_FILE.url, params).then(res => {
                    if (res.data.code === 0) {
                        this.$message.success(res.data.msg);
                    } else {
                        this.$message.error(res.data.msg);
                    }
                });
            },
            // 查询模板
            findTempl() {
                this.$axios.get(api.GET_TEMPLATE_LIST.url).then(res => {
                    if (res.data.code === 0) {
                        if (res.data.data.total > 0) {
                            this.template = res.data.data.list[0];
                        } else {
                            this.$message.error('请先创建周报模板');
                        }
                    } else {
                        this.$message.error(res.data.msg);
                    }
                });
            },
            // 失去焦点事件
            onEditorBlur(event, id){
                document.getElementById(id).children[0].style.zIndex = '1';

            },
            // 获得焦点事件
            onEditorFocus(event, id){
                document.getElementById(id).children[0].style.zIndex = '2';
            },
            // 查询人员列表
            findUserList() {
                const params = {};
                params.page = 1;
                params.pageSize = 99999;
                this.$axios.post(api.GET_USER_LIST.url, params).then(res => {
                    if (res.data.code === 0) {
                        this.userList = res.data.data.list;
                    } else {
                        this.$message.error(res.data.msg);
                    }
                })
            },
            userChange(select) {
                const temp = [];
                if (select.length > 0) {
                    this.userList.forEach(li => {
                        select.forEach(id => {
                            if (li.ID === id) {
                                const data = {};
                                data.id = id;
                                data.name = li.userName;
                                temp.push(data);
                            }
                        })
                    });
                }
                this.sendTo = temp;
                if (this.sendTo.length > 0) {
                    this.sendToClass = 'canSubmit';
                }
            },
            addReports() {
                const pic = [];
                if (this.fileList.length > 0) {
                    this.fileList.forEach(li => {
                        const da = {};
                        da.key = li.response.data.file.key;
                        da.name = li.response.data.file.name;
                        pic.push(da);
                    });
                }
                this.pictures = pic;

                const atta = [];
                if (this.fileList2.length > 0) {
                    this.fileList2.forEach(li => {
                        const da = {};
                        da.key = li.response.data.file.key;
                        da.name = li.response.data.file.name;
                        atta.push(da);
                    });
                }
                this.attachments = atta;

                if (this.sendTo.length === 0) {
                    this.sendToClass = 'cannotSubmit';
                } else {
                    const params = {};
                    params.userId = Number(sessionStorage.getItem('userId'));
                    params.userName = sessionStorage.getItem('userName');
                    params.sendTo = this.sendTo;
                    params.pictures = this.pictures;
                    params.attachments = this.attachments;
                    params.header = this.template.header;
                    params.contents = this.template.contents;
                    this.$axios.post(api.ADD_REPORT.url, params).then(res => {
                        if (res.data.code === 0) {
                            this.$message.success(res.data.msg);
                            this.findTempl();
                            this.findTempl();
                            this.fileList = [];
                            this.previewVisible = false;
                            this.previewImage = '';
                            this.pictures = [];
                            this.fileList2 = [];
                            this.template = {};
                            this.userList = [];
                            this.attachments = [];
                            this.sendTo = [];
                            this.sends = [];
                        } else {
                            this.$message.error(res.data.msg);
                        }
                    })
                }
            }
        },
        computed: {
            editor() {
                return this.$refs.myQuillEditor.quill;
            },
        },
    }
</script>

<style xml:lang="less">
    .ql-toolbar {
        position: fixed;
        top: 104px;
        background-color: white;
        width: calc(100% - 270px);
        left: 260px;
        border-radius: 4px;
        box-shadow: 1px 2px 1px 1px rgba(144,144,144,0.1);
    }
    .quill-editor {
        border: 1px solid #ccc;
        margin: 20px 0;
    }
    .ql-toolbar.ql-snow .ql-picker.ql-expanded .ql-picker-options {
        z-index: 3;
        top: 33px;
    }
    .ql-snow .ql-tooltip {
        left: 0!important;
    }
    .header_write {
        text-align: center;
        font-size: 18px;
        font-weight: 600;
    }
    .canSubmit > div {
        border-color: #d9d9d9;
    }
    .cannotSubmit > div {
        border-color: red;
    }
    .content_wri {
        margin: 50px 0;
    }
    .title_weight {
        font-weight: 600;
    }
    .commit_button {
        width: 400px;
        margin-top: 20px;
    }
</style>
