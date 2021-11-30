export default({
  namespaced: true,
  state: {
    location: '',
    defaultSelectedKeys: [2],
  },
  getters: {},
  mutations: {
    setLocation (state, data) {
      state.location = data;
    },
    setDefaultSelectedKeys (state, data) {
      state.defaultSelectedKeys = data;
    },
  },
  actions: {}
})
