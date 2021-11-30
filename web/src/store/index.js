import Vue from 'vue'
import Vuex from 'vuex'
import pageRouterData from './modules/pageRouterData'
import createPersistedState from 'vuex-persistedstate'

Vue.use(Vuex);

export default new Vuex.Store({
    modules: {
        pageRouterData,
    },
    state: {},
    mutations: {},
    // vuex持久化插件
    plugins: [createPersistedState(
        {storage: window.sessionStorage}
    )]
})
