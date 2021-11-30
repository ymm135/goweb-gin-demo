<template>
    <div>
        <div class="static_ana">数据统计</div>
        <div class="title_ana">周报</div>
        <div class="content_ana">
            <div class="left_ana">
                <div>{{infoData.commitCount}}</div>
                <a-popover placement="topLeft">
                    <template slot="content">
                        {{commitPeoples}}
                    </template>
                    <span>已提交</span>
                </a-popover>
            </div>
            <div class="right_ana">
                <div>{{infoData.uncommitCount}}</div>
                <a-popover placement="topLeft">
                    <template slot="content">
                        {{uncommitPeoples}}
                    </template>
                    <span>未提交</span>
                </a-popover>
            </div>
        </div>

        <div class="export_ana">导出功能</div>
        <div class="content_ana">
            <div class="left_ana">
                <span class="title_ana">人员</span>
                <a-select
                        mode="multiple"
                        style="width: 400px; margin-left: 20px"
                        placeholder="请选择人员"
                        @change="userChange"
                        v-model="sends"
                >
                    <a-select-option v-for="li in userList" :key="li.ID">
                        {{li.userName}}
                    </a-select-option>
                </a-select>
            </div>
            <div class="right_ana">
                <span class="time_ana">起止时间</span>
                <a-range-picker @change="timeOnChange" />

                <a-button type="primary" @click="exportFile" style="margin-left: 20px"> 导出 </a-button>
            </div>
        </div>
    </div>
</template>

<script>
    import api from '../../utils/api'

    export default {
        name: "StatisticalExport",
        data() {
            return {
                infoData: {},
                uncommitPeoples: '',
                commitPeoples: '',
                // 人员列表
                userList: [],
                sends: [],
                reporters: [],
                startTime: '',
                endTime: '',
            }
        },
        created() {
            this.findResult();
            this.findUserList();
        },
        methods: {
            findResult() {
                this.$axios.get(api.FIND_RESULT.url + `?userId=${sessionStorage.getItem('userId')}`).then(res => {
                    if (res.data.code === 0) {
                        this.infoData = res.data.data.rewtOutput;
                        let un = '';
                        if (res.data.data.rewtOutput.uncommitPeoples) {
                            res.data.data.rewtOutput.uncommitPeoples.forEach(da => {
                                un += da.name + '，';
                            });
                        }
                        this.uncommitPeoples = un.substring(0, un.length - 1);

                        let is = '';
                        if (res.data.data.rewtOutput.commitPeoples) {
                            res.data.data.rewtOutput.commitPeoples.forEach(da => {
                                is += da.name + '，';
                            });
                        }
                        this.commitPeoples = is.substring(0, un.length - 1);
                    } else {
                        this.$message.error(res.data.msg);
                    }
                })
            },
            findUserList() {
                const params = {};
                params.page = 1;
                params.pageSize = 99999;
                this.$axios.post(api.GET_USER_LIST.url, params).then(res => {
                    if (res.data.code === 0) {
                        let data = [];
                        const li = {};
                        li.ID = 0;
                        li.userName = '全部';
                        data.push(li);
                        data = data.concat(res.data.data.list);
                        this.userList = data;

                        const all = [];
                        res.data.data.list.forEach(li => {
                            all.push(li.ID);
                        });
                        this.allId = all;
                    } else {
                        this.$message.error(res.data.msg);
                    }
                })
            },
            userChange(select) {
                let flag = false;
                select.forEach(id => {
                    if (id === 0) {
                        flag = true;
                    }
                });

                const temp = [];
                if (flag) {
                    this.sends = this.allId;
                    this.userList.forEach(li => {
                        if (li.ID !== 0) {
                            const data = {};
                            data.id = li.ID;
                            data.name = li.userName;
                            temp.push(data);
                        }
                    });
                } else {
                    if (select.length > 0) {
                        this.userList.forEach(li => {
                            select.forEach(id => {
                                if (li.ID === id) {
                                    const data = {};
                                    data.id = id;
                                    data.name = li.userName;
                                    temp.push(data);
                                }
                            });
                        });
                    }
                }
                this.reporters = temp;
            },
            timeOnChange(date, dateString) {
                this.startTime = dateString[0];
                this.endTime = dateString[1];
            },
            exportFile() {
                const name = this.startTime + '-' + this.endTime + '周报汇总.xlsx';
                this.$axios({
                    url: api.EXPORT_FILE.url + `?userIds=${this.sends}&startTime=${this.startTime}&endTime=${this.endTime}`,
                    method: 'get',
                    responseType: 'blob'
                }).then((res) => {
                    let blob = new Blob([res.data]);
                    let url = URL.createObjectURL(blob);
                    this.toFile(url, name)
                })
            },
            toFile(downUrl, fileName) {
                const aLinkUrl = document.createElement('a');
                aLinkUrl.href = downUrl;
                aLinkUrl.download = fileName;
                const clickAlink = (obj) => {
                    const ev = document.createEvent('MouseEvents');
                    ev.initMouseEvent('click', true, false, window, 0, 0, 0, 0, 0, false, false, false, false, 0, null);
                    obj.dispatchEvent(ev)
                };
                clickAlink(aLinkUrl)
            },
        }
    }
</script>

<style xml:lang="less">
    .static_ana {
        font-size: 18px;
        font-weight: 600;
        margin-bottom: 20px;
    }
    .title_ana {
        font-size: 16px;
        font-weight: 600;
        margin: 20px;
    }
    .content_ana {
        display: flex;
        width: 100%;
        font-size: 20px;
        font-weight: 600;
    }
    .left_ana {
        width: 50%;
        text-align: right;
        padding-right: 50px;
        border-right: 1px solid #ccc;
    }
    .right_ana {
        width: 50%;
        padding-left: 50px;
    }
    .export_ana {
        font-size: 18px;
        font-weight: 600;
        margin: 20px 0;
    }
    .time_ana {
        font-size: 16px;
        font-weight: 600;
        margin: 20px 20px 20px 0;
    }
</style>
