<template>
  <a-config-provider :locale="zh_CN">
    <div id="app">
      <router-view></router-view>
    </div>
  </a-config-provider>
</template>

<script>
import zh_CN from 'ant-design-vue/lib/locale-provider/zh_CN'

export default {
  name: 'App',
  data() {
    return{
      zh_CN
    }
  }
}
</script>

<style>
html, body {
  height: 100%;
  width: 100%;
}
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  height: 100%;
  width: 100%;
}
main {
  height: 100%;
  overflow-y: auto;
  overflow-x: hidden;
  width: 100%;
}
main > div {
  width: calc(100vw - 310px);
  width: 100%;
}
</style>
