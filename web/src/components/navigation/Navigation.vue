<template>
    <a-layout class="content">
        <a-layout-header class="header">
            <div>
                <div class="title_nav">
                    周报系统
                </div>
                <div @click="toggleCollapsed" class="control_button">
                    <a-icon :type="collapsed ? 'menu-unfold' : 'menu-fold'" />
                </div>
            </div>

            <div class="tip_div">
                欢迎您：
                <span class="tip_div_nickName">{{nickName}}</span>|
                <span @click="logout" class="logout_span">注销</span>
            </div>
        </a-layout-header>
        <a-layout>
            <div class="left_div_nav">
                <div :class="left_div">
                    <a-menu
                            class="menu_nav"
                            mode="inline"
                            theme="dark"
                            :inline-collapsed="collapsed"
                            style="text-align: left"
                            @click="menuClick"
                            :defaultSelectedKeys="defaultSelectedKeys"
                            :defaultOpenKeys="[1,6]"
                    >
                        <template v-for="item in menuList">
                            <a-menu-item v-if="!item.children" :key="item.ID" :title="item.name">
                                <a-icon type="pie-chart" />
                                <span>{{ item.name }}</span>
                            </a-menu-item>
                            <sub-menu v-else :key="item.ID" :menu-info="item" />
                        </template>
                    </a-menu>
                </div>
                <div :class="right_div" class="right_div_nav">
                    <div class="current_div">
                        <span class="current_span">当前位置：</span>
                        <span>{{location}}</span>
                    </div>
                    <a-layout class="layout_nav">
                        <a-layout-content
                                :style="{ background: '#fff', padding: '20px', margin: 0, minHeight: '280px' }"
                        >
                            <router-view></router-view>
                        </a-layout-content>
                    </a-layout>
                </div>
            </div>
        </a-layout>
    </a-layout>
</template>

<script>
    import { Menu } from 'ant-design-vue';
    import api from '../../utils/api';

    const SubMenu = {
        template: `
      <a-sub-menu :key="menuInfo.ID" v-bind="$props" v-on="$listeners">
        <span slot="title" :title="menuInfo.name">
          <a-icon :type="menuInfo.icon" /><span>{{ menuInfo.name }}</span>
        </span>
        <template v-for="item in menuInfo.children">
          <a-menu-item v-if="!item.children" :key="item.ID" :title="item.name">
            <a-icon :type="item.icon" />
            <span>{{ item.name }}</span>
          </a-menu-item>
          <sub-menu v-else :key="item.ID" :menu-info="item" />
        </template>
      </a-sub-menu>
    `,
        name: 'SubMenu',
        // must add isSubMenu: true
        isSubMenu: true,
        props: {
            ...Menu.SubMenu.props,
            // Cannot overlap with properties within Menu.SubMenu.props
            menuInfo: {
                type: Object,
                default: () => ({}),
            },
        },
    };

    export default {
        components: {
            'sub-menu': SubMenu,
        },
        watch: {
            // 监听vuex
            '$store.state.pageRouterData.location' (newval) {
                this.location = newval;
            },
            '$store.state.pageRouterData.defaultSelectedKeys' (newval) {
                this.defaultSelectedKeys = newval;
            }
        },
        name: "Navigation",
        data() {
            return {
                collapsed: false,
                left_div: 'left_div_open',
                right_div: 'right_div_open',
                location: this.$store.state.pageRouterData.location,
                defaultSelectedKeys: this.$store.state.pageRouterData.defaultSelectedKeys,
                menuList: [],
                // 当前登录用户
                nickName: sessionStorage.getItem('nickName'),
            }
        },
        created() {
            setTimeout(this.findMenuList(), 1000);
        },
        methods: {
            // 路由跳转
            trip(name, path) {
                this.location = name;
                this.$router.push(path);
            },
            toggleCollapsed() {
                this.collapsed = !this.collapsed;
                if (this.left_div === 'left_div_open') {
                    this.left_div = 'left_div_close';
                    this.right_div = 'right_div_close';

                    if (document.getElementsByClassName('ql-toolbar').length > 0) {
                        document.getElementsByClassName('ql-toolbar').forEach(li => {
                            li.style.left = '90px';
                            li.style.width = 'calc(100% - 100px)';
                        });
                    }
                } else{
                    this.left_div = 'left_div_open';
                    this.right_div = 'right_div_open';
                    document.getElementsByClassName('ql-toolbar').forEach(li => {
                        li.style.left = '260px';
                        li.style.width = 'calc(100% - 270px)';
                    });
                }
            },
            findMenuList() {
                this.$axios.defaults.headers.common['x-token'] = sessionStorage.getItem('login');
                this.$axios.defaults.headers.common['x-user-id'] = sessionStorage.getItem('userId');
                this.$axios.post(api.GET_MENU.url).then(res => {
                    if (res.data.code === 0) {
                        this.menuList = res.data.data.menus;
                    }
                });
            },
            logout() {
                sessionStorage.clear();
                this.$store.commit('pageRouterData/setLocation', '查看周报');
                this.$store.commit('pageRouterData/setDefaultSelectedKeys', [2]);
                this.$router.push('/login');
            },
            // eslint-disable-next-line no-unused-vars
            menuClick({ item, key }) {
                this.getCheckedNodes(this.menuList, key);
                const keyArray = [key];
                this.$store.commit('pageRouterData/setDefaultSelectedKeys', keyArray);
            },
            getCheckedNodes(data, value) {
                data.forEach(item => {
                    if (item.ID === value) {
                        this.$store.commit('pageRouterData/setLocation', item.name);
                        this.trip(item.name, item.path);
                    } else {
                        if (item.children && item.children.length > 0) {
                            this.getCheckedNodes(item.children, value);
                        }
                    }
                });
            },
        }
    }
</script>

<style scoped>
    .content {
        height: 100%;
    }
    section {
        overflow: hidden;
        background-color: rgba(255, 255, 255, 0.4);
    }
    .left_div_open {
        width: 250px;
        background-color: rgb(0, 21, 41);
        height: 100%;
        overflow-y: auto;
    }
    .left_div_close {
        width: 80px;
        background-color: rgb(0, 21, 41);
        height: 100%;
        overflow-y: auto;
    }
    .right_div_open {
        width: calc(100% - 250px)
    }
    .right_div_close {
        width: calc(100% - 80px)
    }
    .title_nav {
        color: white;
        font-size: 25px;
        font-weight: 600;
        text-align: left;
        margin-left: -25px;
        width: 225px;
        float: left;
    }
    .control_button {
        float: left;
        color: white;
        cursor: pointer;
    }
</style>
<style xml:lang="less">
    .ant-layout-header {
        height: 64px;
        padding: 0 50px;
        line-height: 64px;
        background: #0a42ab!important;
    }
    .__tabs {
        display: -webkit-box;
        margin-top: -1px;
        width: 100%;
        overflow-x: hidden;
    }
    .__tab-item {
        white-space: nowrap;
        padding: 8px;
        font-size: 14px;
        border: 1px solid #cccccc4d;
        border-left: none;
        border-bottom: 0;
        line-height: 14px;
        cursor: pointer;
        transition: color 0.3s cubic-bezier(0.645, 0.045, 0.355, 1),
        padding 0.3s cubic-bezier(0.645, 0.045, 0.355, 1);
        overflow: hidden;
    }
    .el-icon-close {
        width: 12px;
        margin-right: 0;
    }
    .__is-active {
        border-bottom: 2px solid #409eff;
        color: #409eff;
        background-color: #f6f6f6;
    }
    .__is-active .el-icon-close {
        width: 12px;
        margin-right: 0;
        margin-left: 2px;
    }
    .__contextmenu {
        margin: 0;
        border: 1px solid #e4e7ed;
        background: #fff;
        z-index: 3000;
        position: absolute;
        list-style-type: none;
        padding: 5px 0;
        border-radius: 4px;
        font-size: 14px;
        color: #333;
        box-shadow: 1px 1px 3px 0 rgba(0, 0, 0, 0.1);
    }
    .__contextmenu li {
        margin: 0;
    }
    .__contextmenu li :hover {
        background: #f2f2f2;
        cursor: pointer;
    }
    .__contextmenu li button {
        color: #2c3e50;
        border: none;
    }
    .ant-menu-item {
        margin: 0!important;
    }
    .tip_div {
        float: right;
        color: white;
        width: 180px;
        display: flex;
    }
    .tip_div_nickName {
        margin-right: 10px
    }
    .logout_span {
        margin-left: 10px;
        cursor: pointer;
    }
    .left_div_nav {
        display: flex;
        height: 100%;
    }
    .right_div_nav {
        background-color: rgba(231, 234, 237, 0.3);
    }
    .current_div {
        background-color: white;
        height: 30px;
        text-align: left;
        line-height: 30px;
        display: flex;
    }
    .current_span {
        font-weight: 600;
        font-size: 14px;
        padding-left: 10px;
    }
    .layout_nav {
        margin: 10px;
        height: calc(100% - 50px);
    }
</style>
