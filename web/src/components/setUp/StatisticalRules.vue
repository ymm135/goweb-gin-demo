<template>
    <div>
        <div class="rule_div">设定规则</div>
        <div>
            <span>需要提交人</span>
            <a-select
                    mode="multiple"
                    class="select_sta"
                    placeholder="请选择人员"
                    @change="userChange"
                    v-model="sends"
            >
                <a-select-option v-for="li in userList" :key="li.ID">
                    {{li.nickName}}
                </a-select-option>
            </a-select>
        </div>

        <div class="main_sta_div">
            <span class="time_sta_span">起止时间</span>
            <a-select v-model="startWeek" class="select_sta_time">
                <a-select-option :value="1">
                    周一
                </a-select-option>
                <a-select-option :value="2">
                    周二
                </a-select-option>
                <a-select-option :value="3">
                    周三
                </a-select-option>
                <a-select-option :value="4">
                    周四
                </a-select-option>
                <a-select-option :value="5">
                    周五
                </a-select-option>
                <a-select-option :value="6">
                    周六
                </a-select-option>
                <a-select-option :value="7">
                    周日
                </a-select-option>
            </a-select>

            <a-select v-model="startHour" class="select_str">
                <a-select-option :key="0">
                    0:00
                </a-select-option>
                <a-select-option v-for="li in 23" :key="li">
                    {{li}}:00
                </a-select-option>
            </a-select>

            <span style="margin-left: 20px">~</span>

            <a-select v-model="endWeek" class="select_str">
                <a-select-option :value="1">
                    周一
                </a-select-option>
                <a-select-option :value="2">
                    周二
                </a-select-option>
                <a-select-option :value="3">
                    周三
                </a-select-option>
                <a-select-option :value="4">
                    周四
                </a-select-option>
                <a-select-option :value="5">
                    周五
                </a-select-option>
                <a-select-option :value="6">
                    周六
                </a-select-option>
                <a-select-option :value="7">
                    周日
                </a-select-option>
            </a-select>

            <a-select v-model="endHour" style="width: 80px; margin-left: 10px">
                <a-select-option :key="0">
                    0:00
                </a-select-option>
                <a-select-option v-for="li in 23" :key="li">
                    {{li}}:00
                </a-select-option>
            </a-select>

            <a-button type="primary" @click="submitData" style="margin-left: 20px"> 提交 </a-button>
        </div>
    </div>
</template>

<script>
    import api from '../../utils/api'

    export default {
        name: "StatisticalRules",
        data() {
            return {
                isCreate: true,
                // 人员列表
                userList: [],
                sends: [],
                reporters: [],
                // 全选id
                allId: [],
                startWeek: '',
                startHour: '',
                endWeek: '',
                endHour: '',
                editId: '',
            }
        },
        created() {
          this.findDataList();
        },
        methods: {
            findDataList() {
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

                        this.$axios.get(api.FIND_RULE_LIST.url + `?page=1&pageSize=1&userId=${sessionStorage.getItem('userId')}`).then(resp => {
                            if (resp.data.code === 0) {
                                if (resp.data.data.list) {
                                    this.isCreate = false;
                                    this.editId = resp.data.data.list[0].ID;
                                    this.startWeek = resp.data.data.list[0].startWeek;
                                    this.startHour = resp.data.data.list[0].startHour;
                                    this.endWeek = resp.data.data.list[0].endWeek;
                                    this.endHour = resp.data.data.list[0].endHour;

                                    const user = [];
                                    const temp = [];

                                    for (const li of resp.data.data.list[0].reporters) {
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

                                        const data = {};
                                        data.id = li.id;
                                        data.name = li.name;
                                        temp.push(data);
                                    }
                                    this.sends = user;
                                    this.reporters = temp;
                                }
                            } else {
                                this.$message.error(resp.data.msg);
                            }
                        })
                    } else {
                        this.$message.error(res.data.msg);
                    }
                })
            },
            userChange(select) {
                let flag = false;
                select.forEach(li => {
                    if (li === 0) {
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
                            })
                        });
                    }
                }
                this.reporters = temp;
            },
            submitData() {
                if (this.isCreate) {
                    // 新增
                    const params = {};
                    params.userId = Number(sessionStorage.getItem('userId'));
                    params.startWeek = this.startWeek;
                    params.startHour = this.startHour;
                    params.endWeek = this.endWeek;
                    params.endHour = this.endHour;
                    params.reporters = this.reporters;
                    this.$axios.post(api.ADD_RULE.url, params).then(res => {
                        if (res.data.code === 0) {
                            this.$message.success(res.data.msg);
                        } else {
                            this.$message.error(res.data.msg);
                        }
                    })
                } else {
                    // 编辑
                    const params = {};
                    params.ID = this.editId;
                    params.userId = Number(sessionStorage.getItem('userId'));
                    params.startWeek = this.startWeek;
                    params.startHour = this.startHour;
                    params.endWeek = this.endWeek;
                    params.endHour = this.endHour;
                    params.reporters = this.reporters;
                    this.$axios.put(api.UPDATE_RULE.url, params).then(res => {
                        if (res.data.code === 0) {
                            this.$message.success(res.data.msg);
                        } else {
                            this.$message.error(res.data.msg);
                        }
                    })
                }
            }
        }
    }
</script>

<style xml:lang="less">
    .rule_div {
        font-size: 18px;
        font-weight: 600;
        margin-bottom: 20px;
    }
    .select_sta {
        width: 500px;
        margin-left: 20px;
    }
    .main_sta_div {
        margin-top: 10px;
    }
    .time_sta_span {
        margin-left: 14px;
    }
    .select_sta_time {
        width: 80px;
        margin-left: 20px;
    }
    .select_str {
        width: 80px;
        margin-left: 10px;
    }
</style>
