<template>
  <div class="app-wrapper">
    <div class="container">
      <div>
        <navbar />
      </div>
      <app-main />
    </div>
  </div>
</template>

<script>
import { AppMain, Navbar } from './components'
import { mapState } from 'vuex'

export default {
  name: 'Index',
  components: {
    AppMain,
    Navbar
  },
  computed: {
    ...mapState({
      device: state => state.app.device
    })
  },
  methods: {
  }
}
</script>

<style lang="scss" scoped>
  @import "~@/styles/mixin.scss";
  @import "~@/styles/variables.scss";

  .app-wrapper {
    @include clearfix;
    position: relative;
    height: 100%;
    width: 100%;
  }

  .drawer-bg {
    background: #000;
    opacity: 0.3;
    width: 100%;
    top: 0;
    height: 100%;
    position: absolute;
    z-index: 999;
  }

  .fixed-header {
    position: fixed;
    top: 0;
    right: 0;
    z-index: 9;
    width: calc(100% - #{$sideBarWidth});
    transition: width 0.28s;
  }

  .hideSidebar .fixed-header {
    width: calc(100% - 54px)
  }

  .mobile .fixed-header {
    width: 100%;
  }
</style>
