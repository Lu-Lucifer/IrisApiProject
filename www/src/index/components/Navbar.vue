<template>
  <div class="navbar">
    <div class="text-center">  <pan-thumb width="150px" height="150px" image="https://thirdwx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTIPbZRufW9zPiaGpfdXgU7icRL1licKEicYyOiace8QQsYVKvAgCrsJx1vggLAD2zJMeSXYcvMSkw9f4pw/132">
      snowlyg
    </pan-thumb></div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import PanThumb from '@/components/PanThumb'

export default {
  components: { PanThumb },
  computed: {
    ...mapGetters(['device'])
  },
  methods: {}
}
</script>

<style lang="scss" scoped>
.navbar {
  height: 250px;

  //overflow: hidden;
  //position: relative;
  //background: #fff;
  //-webkit-box-shadow: 0 1px 3px 0 rgba(0, 0, 0, .1), 0 1px 2px 0 rgba(0, 0, 0, .06);
  //box-shadow: 0 1px 3px 0 rgba(0, 0, 0, .1), 0 1px 2px 0 rgba(0, 0, 0, .06);
  //border-top: 4px solid #368cda;
  //color: #777;
}

</style>
