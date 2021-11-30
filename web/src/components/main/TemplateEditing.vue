<template>
    <div class="content_div">
        <div>
            <span class="title_tem">模板标题</span>
            <a-input v-model="data.header" placeholder="请输入模板标题" class="modal_title"></a-input>

            <span class="create_people">创建人</span>
            <span v-if="isCreate">{{userName}}</span>
            <span v-else>{{data.userName}}</span>

            <a-button type="primary" style="float: right" @click="saveData"> <a-icon type="save" /> 保存 </a-button>
        </div>
        <div>
            <div v-if="!isCreate">
                <div v-for="(li, index) in data.contents" :key="index">
                    <div class="content1">
                        <div class="content2">
                            <a-input v-model="li.title" placeholder="请输入字段标题" class="input_temp"></a-input>
                            <a-icon type="delete" style="cursor: pointer" @click="deleteContent(index)"/>
                        </div>
                        <div class="content_tem">
                            <span>
                                待填写者输入
                            </span>
                        </div>
                    </div>
                </div>
            </div>
            <div v-else>
                <div v-for="(li, index) in addContents" :key="index">
                    <div class="content1">
                        <div class="content2">
                            <a-input v-model="li.title" placeholder="请输入字段标题" class="input_temp"></a-input>
                            <a-icon type="delete" style="cursor: pointer" @click="deleteContent(index)"/>
                        </div>
                        <div class="content_tem">
                            <span>
                                待填写者输入
                            </span>
                        </div>
                    </div>
                </div>
            </div>

            <a-button type="primary" class="add_button_tem" @click="addContent"> <a-icon type="plus" /> </a-button>
        </div>
    </div>
</template>

<script>
    import api from '../../utils/api'

    export default {
        name: "TemplateEditing",
        created() {
            this.findData();
        },
        data() {
            return {
                data: {},
                // 创建还是编辑
                isCreate: true,
                userName: '',
                addContents: []
            }
        },
        methods: {
            findData() {
                this.$axios.get(api.GET_TEMPLATE_LIST.url).then(res => {
                    if (res.data.code === 0) {
                        if (res.data.data.total > 0) {
                            this.isCreate = false;
                            this.data = res.data.data.list[0];
                        } else {
                            this.isCreate = true;
                            this.data.userName = sessionStorage.getItem('userName');
                            this.userName = sessionStorage.getItem('userName');
                        }
                    } else {
                        this.$message.error(res.data.msg);
                    }
                });
            },
            addContent() {
                const li = {};
                li.title = '';
                li.content = '';
                if (this.data.contents) {
                    this.data.contents.push(li);
                    this.addContents.push(li);
                } else {
                    const con = [];
                    con.push(li);
                    this.data.contents = con;
                    this.addContents.push(li);
                }
            },
            deleteContent(index) {
                this.data.contents.splice(index, 1);
            },
            saveData() {
                if (this.isCreate) {
                    // 创建接口
                    this.$axios.post(api.ADD_TEMPLATE.url, this.data).then(res => {
                        if (res.data.code === 0) {
                            this.$message.success(res.data.msg);
                        } else {
                            this.$message.error(res.data.msg);
                        }
                    });
                } else {
                    // 编辑接口
                    this.$axios.put(api.EDIT_TEMPLATE.url, this.data).then(res => {
                        if (res.data.code === 0) {
                            this.$message.success(res.data.msg);
                        } else {
                            this.$message.error(res.data.msg);
                        }
                    });
                }
            }
        }
    }
</script>

<style xml:lang="less">
    .content_div {
        height: 100%;
        padding: 0 200px;
    }
    .content1 {
        margin-top: 20px;
        box-shadow: 0 2px 12px 0 #ccc;
        border-radius: 8px;
        padding: 20px;
    }
    .content2 {
        border-bottom: 1px solid #ccc;
    }
    .input_temp {
        width: calc(100% - 50px);
        border: none;
        font-size: 16px;
        font-weight: 600;
    }
    .input_temp:focus {
        border: none!important;
        box-shadow: none!important;
    }
    .title_tem {
        font-size: 20px;
        font-weight: 600;
    }
    .modal_title {
        width: 350px;
        margin-left: 10px;
    }
    .create_people {
        margin-left: 100px;
        margin-right: 10px;
    }
    .content_tem {
        margin: 10px 0 0 10px;
    }
    .add_button_tem {
        width: 100%;
        margin-top: 20px;
    }
</style>
