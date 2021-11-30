<template>
    <div>
        <div>
            <span>人员</span>
            <a-select
                    class="select_div"
                    placeholder="请选择人员"
                    v-model="selectedUserId"
                    :allowClear="true"
                    :showSearch="true"
            >
                <a-select-option v-for="li in userList" :key="li.ID">
                    {{li.nickName}}
                </a-select-option>
            </a-select>

            <span class="span_title">起止时间</span>
            <a-range-picker show-time style="margin-left: 10px" v-model="rangeTime">
                <template slot="renderExtraFooter">
                    extra footer
                </template>
            </a-range-picker>

            <span class="span_title">周报内容</span>
            <a-input v-model="content" placeholder="请输入周报内容" class="content_input"></a-input>

            <a-button type="primary" @click="findDataList" class="span_title">
                查询
            </a-button>

            <div style="margin-top: 10px">
                <a-table :columns="columns" :data-source="tableList" rowKey="ID" :pagination="pagination" @change="findDataList">
                    <span slot="action" slot-scope="text, li">
                        <a @click="openInfo(li)">详情</a>
                    </span>
                    <span slot="contents" slot-scope="text" class="ecllipsis" :title="toContent(text)">
                        {{toContent(text)}}
                    </span>
                </a-table>
            </div>
        </div>
    </div>
</template>

<script>
    import api from '../../utils/api'

    export default {
        name: "MainComponent",
        data() {
            return {
                userList: [],
                selectedUserId: '',
                rangeTime: [],
                content: '',
                columns: [
                    {
                        title: '姓名',
                        dataIndex: 'nickName',
                        key: 'nickName'
                    },
                    {
                        title: '内容',
                        dataIndex: 'contents',
                        key: 'contents',
                        scopedSlots: { customRender: 'contents' },
                        width: '500px',
                    },
                    {
                        title: '评论数',
                        dataIndex: 'commentCount',
                        key: 'commentCount'
                    },
                    {
                        title: '时间',
                        dataIndex: 'CreatedAt',
                        key: 'CreatedAt'
                    },
                    {
                        title: '操作',
                        key: 'action',
                        scopedSlots: { customRender: 'action' },
                    },
                ],
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
            }
        },
        created() {
            this.findUserList();
            this.findDataList();
        },
        methods: {
            // 查询人员列表
            findUserList() {
                const params = {};
                params.page = 1;
                params.pageSize = 9999;

                this.$axios.post(api.GET_USER_LIST.url, params).then(res => {
                    if (res.data.code === 0) {
                        this.userList = res.data.data.list;
                    } else {
                        this.$message.error(res.data.msg);
                    }
                })
            },
            // 分页查询周报列表
            findDataList() {
                let startTime = '';
                let endTime = '';
                if (this.rangeTime.length !== 0) {
                    startTime = this.toDate(this.rangeTime[0]._d);
                    endTime = this.toDate(this.rangeTime[1]._d);
                }
                this.$axios.get(api.FIND_REPORT_LIST.url + `?page=${this.pagination.defaultCurrent}&pageSize=${this.pagination.defaultPageSize}&content=${this.content}&currUserId=${sessionStorage.getItem('userId')}&userId=${this.selectedUserId}&startTime=${startTime}&endTime=${endTime}`).then(res => {
                    if (res.data.code === 0) {
                        this.tableList = res.data.data.list;
                        this.pagination.total = res.data.data.total;
                    } else {
                        this.$message.error(res.data.msg);
                        this.pagination.total = 0;
                    }
                })
            },
            openInfo(data) {
                this.$router.push({
                    path:'/editWeeklyReport',
                    query:{data: JSON.stringify(data.ID)}
                })

            },
            // 日期转化
            toDate (date) {
                const y = date.getFullYear();
                let m = date.getMonth() + 1;
                m = m < 10 ? ('0' + m) : m;
                let d = date.getDate();
                d = d < 10 ? ('0' + d) : d;
                const h = date.getHours();
                let minute = date.getMinutes();
                minute = minute < 10 ? ('0' + minute) : minute;
                return y + '-' + m + '-' + d+' '+h+':'+minute;
            },
            // 截取中文
            toContent(data) {
                let str = '';
                data.forEach(li => {
                    str += li.title + this.toChinese(li.content) + ';';
                });
                return str;
            },
            toChinese(strValue) {
                if (strValue !== null && strValue !== '') {
                    const reg = /[\u4e00-\u9fa5]/g;
                    let content = '';
                    try {
                        if (strValue.match(reg) !== null) {
                            content = strValue.match(reg).join('');
                        }
                    } catch (e) {
                        console.log(e);
                    }
                    return content;
                }
                return '';
            }
        }
    }
</script>

<style xml:lang="less">
    .ecllipsis {
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
        display: inline-block;
        width: 500px;
    }
    .select_div {
        width: 200px;
        margin-left: 10px;
    }
    .span_title {
        margin-left: 20px;
    }
    .content_input {
        width: 200px;
        margin-left: 10px;
    }
</style>
