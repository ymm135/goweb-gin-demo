<template>
    <div :class="mainDiv">
        <div>
            <a-button type="primary" style="float: right" @click="tripTo()"> 返回 </a-button>
        </div>

        <div class="header_write">{{infoData.header}}</div>
        <!--        模板内容-->
        <div v-for="(item, index) in infoData.contents" :key="index">
            <div class="title_weekly">{{item.title}}：</div>
            <quill-editor
                    :id="index"
                    style="margin-top: 20px"
                    v-model="item.content"
                    :ref="index"
                    @blur="onEditorBlur($event, index)" @focus="onEditorFocus($event, index)"
                    @change="onEditorChange($event)">
            </quill-editor>
        </div>

        <div class="title_weekly">图片上传</div>
        <div>
            <div class="clearfix">
                <a-upload
                        action="/week/fileUploadAndDownload/upload"
                        list-type="picture-card"
                        :file-list="fileList"
                        @preview="handlePreview"
                        @change="handleChange"
                        :showUploadList="showUploadList"
                        accept=".png,.jpeg,.jpg"
                        :headers="{ 'x-token': xToken, 'x-user-id':  xUserId}"
                        :remove="removeFile"
                >
                    <div v-if="canEdit">
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

        <div class="title_weekly">文件上传</div>
        <div>
            <div class="clearfix">
                <a-upload
                        action="/week/fileUploadAndDownload/upload"
                        :file-list="fileList2"
                        @change="handleChange2"
                        :showUploadList="showUploadList"
                        :headers="{ 'x-token': xToken, 'x-user-id':  xUserId}"
                        :remove="removeFile"
                >
                    <div v-if="canEdit">
                        <a-button> <a-icon type="upload" /> Upload </a-button>
                    </div>
                </a-upload>
            </div>
        </div>

        <div class="title_weekly" style="margin: 20px 0 0 0"><span style="color: red">*</span>发送给</div>
        <a-select
                mode="multiple"
                style="width: 100%"
                placeholder="请选择人员"
                @change="userChange"
                v-model="sends"
                :disabled="!canEdit"
                :class="sendToClass"
        >
            <a-select-option v-for="li in userList" :key="li.ID">
                {{li.nickName}}
            </a-select-option>
        </a-select>
        <span style="color: red" v-if="sendToClass === 'cannotSubmit'">请选择人员</span>

        <div style="text-align: center" v-if="canEdit">
            <a-button type="primary" class="commit_button" @click="editReports"> 提交周报 </a-button>
        </div>

        <div class="title_comments">评论</div>
        <div v-for="(li, index) in commentsList">
            <hr v-if="index !== 0"/>

            <div  class="title_weekly" style="font-size: 16px;">{{li.nickName}}：</div>
            <div>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;{{li.comment}}</div>
            <div style="width: 100%; text-align: right">{{li.CreatedAt}}</div>
        </div>
        <div>
            <a-textarea placeholder="请输入评论" v-model="comment" :rows="4" />
            <div style="width: 100%; text-align: center">
                <a-button type="primary" style="margin-top: 20px;" @click="addComment"> 提交评论 </a-button>
            </div>
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
        name: "InfoComponent",
        components: {quillEditor},
        created() {
            this.infoId = JSON.parse(this.$route.query.data);
            // 根据id查询周报
            this.findData();
        },
        data() {
            return {
                mainDiv: 'top0',
                infoId: '',
                infoData: {},
                canEdit: false,
                showUploadList: {},

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
                // 发送到(选中人员列表)
                sendTo: [],
                sendToClass: 'canSubmit',
                // 周报
                commentsList: [],
                comment: ''
            }
        },
        methods: {
            tripTo() {
                this.$router.push('viewWeeklyReport');
            },
            findData() {
                this.$axios.get(api.FIND_REPORT_BY_ID.url + `?id=${this.infoId}`).then(res => {
                    if (res.data.code === 0) {
                        this.infoData = res.data.data.rewtReports;

                        if (this.infoData.userId === Number(sessionStorage.getItem('userId'))) {
                            this.canEdit = true;
                        } else {
                            this.showUploadList = {showRemoveIcon: false};
                            setTimeout(() => {
                                for (const li of document.getElementsByClassName('ql-toolbar')) {
                                    li.style.display = 'none';
                                }
                            }, 80);
                            this.mainDiv = 'top1';
                        }

                        // 获取周报评论
                        this.findComments(this.infoData.ID);

                        this.findUserList().then(() => {
                            this.setData();
                        });
                    } else {
                        this.$message.error(res.data.msg);
                    }
                })
            },
            setData() {
                // 查询模板
                this.template = this.infoData.contents;
                const pic = [];
                for (let index = 0; index < this.infoData.pictures.length; index++) {
                    const temp = {};
                    temp.uid = index;
                    temp.name = this.infoData.pictures[index].name;
                    temp.key = this.infoData.pictures[index].key;
                    temp.url = this.BASEURL + api.DOWNLOAD_FILE.url + '?fileName=' + this.infoData.pictures[index].key;
                    temp.status = 'done';
                    pic.push(temp);
                }
                this.fileList = pic;

                const files = [];
                for (let index = 0; index < this.infoData.attachments.length; index++) {
                    const temp = {};
                    temp.uid = index;
                    temp.name = this.infoData.attachments[index].name;
                    temp.key = this.infoData.attachments[index].key;
                    temp.url = this.BASEURL + api.DOWNLOAD_FILE.url + '?fileName=' + this.infoData.attachments[index].key;
                    temp.status = 'done';
                    files.push(temp);
                }
                this.fileList2 = files;

                const user = [];
                for (const li of this.infoData.sendTo) {
                    let flag = false;
                    for (const tr of this.userList) {
                        if (tr.ID === li.id) {
                            flag = true;
                        }
                    }
                    if (!flag) {
                        const da = {};
                        da.ID = li.id;
                        da.userName = li.name;
                        this.userList.push(da);
                    }
                    user.push(li.id);
                }
                this.sends = user;
                this.sendTo = this.infoData.sendTo;
            },
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
                if (this.canEdit) {
                    const params = {};
                    if (file.response) {
                        params.key = file.response.data.file.key;
                    } else {
                        params.key = file.key;
                    }

                    const files = [];
                    for (const li of this.fileList) {
                        if (li.response) {
                            if (li.response.data.file.key !== params.key) {
                                files.push(li);
                            }
                        } else if (li.key !== params.key) {
                            files.push(li);
                        }
                    }
                    this.fileList = files;

                    const files2 = [];
                    for (const li of this.fileList2) {
                        if (li.response) {
                            if (li.response.data.file.key !== params.key) {
                                files2.push(li);
                            }
                        } else if (li.key !== params.key) {
                            files2.push(li);
                        }
                    }
                    this.fileList2 = files2;

                    this.$axios.post(api.DELETE_FILE.url, params).then(res => {
                        if (res.data.code === 0) {
                            this.$message.success(res.data.msg);
                        } else {
                            this.$message.error(res.data.msg);
                        }
                    });
                }
            },
            // 失去焦点事件
            onEditorBlur(event, id){
                document.getElementById(id).children[0].style.zIndex = '1';

            },
            // 获得焦点事件
            onEditorFocus(event, id){
                if (!this.canEdit) {
                    event.enable(false);
                } else {
                    event.enable(true);
                }
                document.getElementById(id).children[0].style.zIndex = '2';
            },
            // 内容改变事件
            onEditorChange(){
                // console.log(this.template.contents);
            },
            // 查询人员列表
            async findUserList() {
                const params = {};
                params.page = 1;
                params.pageSize = 99999;
                this.$axios.post(api.GET_USER_LIST.url, params).then(res => {
                    if (res.data.code === 0) {
                        this.userList = res.data.data.list;
                        return Promise.resolve(this.userList);
                    } else {
                        this.$message.error(res.data.msg);
                    }
                });
            },
            userChange(select) {
                const temp = [];
                if (select.length > 0) {
                    this.userList.forEach(li => {
                        for (const id of select) {
                            if (li.ID === id) {
                                const data = {};
                                data.id = id;
                                data.name = li.userName;
                                temp.push(data);
                            }
                        }
                    });
                }
                this.sendTo = temp;
                if (this.sendTo.length > 0) {
                    this.sendToClass = 'canSubmit';
                }
            },
            editReports() {
                const pic = [];
                if (this.fileList.length > 0) {
                    this.fileList.forEach(li => {
                        const da = {};
                        if (li.response) {
                            da.key = li.response.data.file.key;
                            da.name = li.response.data.file.name;
                        } else {
                            da.key = li.key;
                            da.name = li.name;
                        }
                        pic.push(da);
                    });
                }
                this.pictures = pic;

                const atta = [];
                if (this.fileList2.length > 0) {
                    this.fileList2.forEach(li => {
                        const da = {};
                        if (li.response) {
                            da.key = li.response.data.file.key;
                            da.name = li.response.data.file.name;
                        } else {
                            da.key = li.key;
                            da.name = li.name;
                        }
                        atta.push(da);
                    });
                }
                this.attachments = atta;

                if (this.sendTo.length === 0) {
                    this.sendToClass = 'cannotSubmit';
                } else {
                    const params = {};
                    params.id = this.infoData.ID;
                    params.userId = sessionStorage.getItem('userId');
                    params.userName = sessionStorage.getItem('userName');
                    params.sendTo = this.sendTo;
                    params.pictures = this.pictures;
                    params.attachments = this.attachments;
                    params.header = this.infoData.header;
                    params.contents = this.infoData.contents;
                    this.$axios.put(api.EDIT_REPORT.url, params).then(res => {
                        if (res.data.code === 0) {
                            this.$message.success(res.data.msg);
                            this.fileList = [];
                            this.previewVisible = false;
                            this.previewImage = '';
                            this.pictures = [];
                            this.fileList2 = [];
                            this.template = {};
                            this.userList = [];
                            this.attachments = [];
                            this.sendTo = [];
                            this.findData();
                        } else {
                            this.$message.error(res.data.msg);
                        }
                    })
                }
            },
            // 获取周报评论
            findComments(reportId) {
                this.$axios.get(api.FIND_COMMENT_LIST.url + `?reportId=${reportId}&page=1&pageSise=99999`).then(res => {
                    if (res.data.code === 0) {
                        this.commentsList = res.data.data.list;
                    } else {
                        this.$message.error(res.data.msg);
                    }
                });
            },
            addComment() {
                const params = {};
                params.reportId = this.infoData.ID;
                params.userName = sessionStorage.getItem('userName');

                params.comment = this.comment;
                this.$axios.post(api.ADD_COMMENT.url, params).then(res => {
                    if (res.data.code === 0) {
                        this.$message.success(res.data.msg);
                        this.findComments(this.infoData.ID);
                        this.comment = '';
                    } else {
                        this.$message.error(res.data.msg);
                    }
                });
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
    .readOnly {
        -webkit-user-modify: read-only;
    }
    .canSubmit > div {
        border-color: #d9d9d9;
    }
    .cannotSubmit > div {
        border-color: red;
    }
    .top0 {
        margin: 50px 0;
    }
    .top1 {
        margin: 0;
    }
    .title_weekly {
        font-weight: 600;
    }
    .commit_button {
        width: 400px;
        margin-top: 20px
    }
    .title_comments {
        font-weight: 600;
        margin: 20px 0;
        font-size: 18px;
    }
</style>
