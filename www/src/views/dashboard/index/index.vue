<template>
  <div class="components-container">
    <div class="dashboard-editor-container">
      <el-row>
        <el-col v-for="(item) in aritcles" :key="item.id" :span="24">
          <el-card :body-style="{ padding: '0px' }">
            <div style="padding: 14px;">
              <img src="item.image_uri" alt="">
              <span> {{ item.title }}</span>
              <div> {{ item.content_short }}</div>
              <div class="bottom clearfix">
                <span> [{{ item.author }}] </span>
                <time class="time">{{ item.display_time | parseTime('{y}-{m}-{d} {h}:{i}:{s}') }}</time>
              </div>
            </div>
          </el-card>
        </el-col>
      </el-row>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import { indexList } from '@/api/article'

export default {
  name: 'DashboardEditor',
  components: {},
  data() {
    return {
      aritcles: []
    }
  },
  created() { this.getData() },
  methods: {
    getData() {
      this.listLoading = true
      indexList().then(response => {
        this.aritcles = response.data.items
      })
    }
  },
  // eslint-disable-next-line vue/order-in-components
  computed: {
    ...mapGetters([])
  }
}
</script>

<style lang="scss" scoped>
.bottom {
  margin-top: 13px;
  line-height: 12px;
}

.button {
  padding: 0;
  float: right;
}

.image {
  width: 100%;
  display: block;
}

.clearfix:before,
.clearfix:after {
  display: table;
  content: "";
}

.clearfix:after {
  clear: both
}

.components-container {
  margin: 0;
  position: relative;
  height: 100vh;

}

.dashboard-editor-container {
  background-color: #e3e3e3;
  min-height: 100vh;
  padding: 60px 350px;
}

.pan-info-roles {
  font-size: 12px;
  font-weight: 700;
  color: #333;
  display: block;
}

.info-container {
  position: relative;
  margin-left: 190px;
  height: 150px;
  line-height: 200px;

  .display_name {
    font-size: 48px;
    line-height: 48px;
    color: #212121;
    position: absolute;
    top: 25px;
  }
}
</style>
