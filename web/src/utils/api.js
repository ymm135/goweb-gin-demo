const GET_CAPTCHA = {
    url: '/base/captcha',
    deacriptions: '获取验证码'
};
const LOGIN = {
    url: '/base/login',
    deacriptions: '登录'
};
const GET_MENU = {
    url: '/menu/getMenu',
    deacriptions: '获取动态路由'
};
const CHANGE_PASSWORD = {
    url: '/user/changePassword',
    deacriptions: '重置密码'
};
const GET_USER_LIST = {
    url: '/user/getUserList',
    deacriptions: '分页获取用户列表'
};
const GET_AUTHORITY_LIST = {
    url: '/authority/getAuthorityList',
    deacriptions: '获取角色列表'
};
const DELETE_USER = {
    url: '/user/deleteUser',
    deacriptions: '删除用户'
};
const ADD_USER = {
    url: '/user/register',
    deacriptions: '新增用户'
};
const EDIT_USER = {
    url: '/user/setUserInfo',
    deacriptions: '编辑用户'
};
const UPLOAD_FILE = {
    url: '/fileUploadAndDownload/upload',
    deacriptions: '文件上传'
};
const DELETE_FILE = {
    url: '/fileUploadAndDownload/deleteFile',
    deacriptions: '文件删除'
};
const DOWNLOAD_FILE = {
    url: '/fileUploadAndDownload/download',
    deacriptions: '文件下载'
};
const GET_TEMPLATE_LIST = {
    url: '/wtTemplates/getWtTemplateList',
    deacriptions: '获取模板'
};
const ADD_TEMPLATE = {
    url: '/wtTemplates/createWtTemplate',
    deacriptions: '新建模板'
};
const EDIT_TEMPLATE = {
    url: '/wtTemplates/updateWtTemplate',
    deacriptions: '编辑模板'
};
const ADD_REPORT = {
    url: '/wtReports/createWtReports',
    deacriptions: '创建周报'
};
const FIND_REPORT_LIST = {
    url: '/wtReports/getWtReportsList',
    deacriptions: '分页查询周报'
};
const EDIT_REPORT = {
    url: '/wtReports/updateWtReports',
    deacriptions: '更新周报'
};
const FIND_REPORT_BY_ID = {
    url: '/wtReports/findWtReports',
    deacriptions: '根据id查询周报'
};
const FIND_COMMENT_LIST = {
    url: '/wtComment/getWtCommentList',
    deacriptions: '获取周报评论'
};
const ADD_COMMENT = {
    url: '/wtComment/createWtComment',
    deacriptions: '创建周报评论'
};
const ADD_RULE = {
    url: '/wtRule/createWtRule',
    deacriptions: '创建规则'
};
const UPDATE_RULE = {
    url: '/wtRule/updateWtRule',
    deacriptions: '编辑规则'
};
const FIND_RULE_LIST = {
    url: '/wtRule/getWtRuleList',
    deacriptions: ' 查询规则'
};
const FIND_RESULT = {
    url: '/wtOutput/GetStatResult',
    deacriptions: ' 查询统计结果'
};
const EXPORT_FILE = {
    url: '/wtOutput/ExportReportToExcel',
    deacriptions: ' 导出报表'
};

//一定要注册才可以使用
export default {
    GET_CAPTCHA: GET_CAPTCHA,
    LOGIN: LOGIN,
    GET_MENU: GET_MENU,
    CHANGE_PASSWORD: CHANGE_PASSWORD,
    GET_USER_LIST: GET_USER_LIST,
    GET_AUTHORITY_LIST: GET_AUTHORITY_LIST,
    DELETE_USER: DELETE_USER,
    ADD_USER: ADD_USER,
    EDIT_USER: EDIT_USER,
    UPLOAD_FILE: UPLOAD_FILE,
    DELETE_FILE: DELETE_FILE,
    DOWNLOAD_FILE: DOWNLOAD_FILE,
    GET_TEMPLATE_LIST: GET_TEMPLATE_LIST,
    ADD_TEMPLATE: ADD_TEMPLATE,
    EDIT_TEMPLATE: EDIT_TEMPLATE,
    ADD_REPORT: ADD_REPORT,
    FIND_REPORT_LIST: FIND_REPORT_LIST,
    EDIT_REPORT: EDIT_REPORT,
    FIND_REPORT_BY_ID: FIND_REPORT_BY_ID,
    FIND_COMMENT_LIST: FIND_COMMENT_LIST,
    ADD_COMMENT: ADD_COMMENT,
    ADD_RULE: ADD_RULE,
    UPDATE_RULE: UPDATE_RULE,
    FIND_RULE_LIST: FIND_RULE_LIST,
    FIND_RESULT: FIND_RESULT,
    EXPORT_FILE: EXPORT_FILE,
}
