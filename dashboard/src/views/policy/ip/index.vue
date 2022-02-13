<template>
  <div
    class="app-container"
  >
    <el-tabs v-model="activeName" @tab-click="handleClick">
      <el-tab-pane label="允许IP" name="allow">
        <IP :typ="1" />
      </el-tab-pane>
      <el-tab-pane label="拒绝IP" name="deny">
        <IP :typ="2" />
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script>
import IP from './components/index'
export default {
  components: { IP },
  data() {
    return {
      activeName: 'allow'
    }
  },
  methods: {
    handleClick(tab, event) {
      console.log(tab, event)
    }
  }
}
</script>
