<template>
  <el-dropdown :show-timeout="100" trigger="click">
    <el-button plain>
      {{ is_original?'原创':'转载' }}
      <i class="el-icon-caret-bottom el-icon--right" />
    </el-button>
    <el-dropdown-menu slot="dropdown" class="no-padding">
      <el-dropdown-item>
        <el-radio-group v-model="is_original" style="padding: 10px;">
          <el-radio :label="true">
            原创
          </el-radio>
          <el-radio :label="false">
            转载
          </el-radio>
        </el-radio-group>
      </el-dropdown-item>
    </el-dropdown-menu>
  </el-dropdown>
</template>

<script>
export default {
  props: {
    value: {
      type: Boolean,
      default: false
    }
  },
  computed: {
    is_original: {
      get() {
        return this.value
      },
      set(val) {
        this.$emit('input', val)
      }
    }
  }
}
</script>
