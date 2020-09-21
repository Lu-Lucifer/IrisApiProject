<template>
  <div class="navbar text-center">
    <div class="text-center"> <span>© 2020-2012 SNOWLYG.COM </span> | <a href="https://beian.miit.gov.cn/#/Integrated/index">湘ICP备20005575号</a></div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'

export default {
  components: {},
  computed: {
    ...mapGetters(['device'])
  },
  methods: {}
}
</script>

<style lang="scss" scoped>
.navbar {
  height: 50px;
  overflow: hidden;
  position: relative;
  background:  #1b1c1d;
  box-shadow: 0 1px 4px rgba(0,21,41,.08);
  color: rgb(137, 137, 140);
  line-height: 50px;
  }
.navbar a{
  size: 12px;
}
</style>
