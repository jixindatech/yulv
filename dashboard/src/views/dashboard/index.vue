<template>
  <div
    v-permission="['GET:/api/db',
                   'GET:/api/dbuser',
                   'GET:/api/event/info/access',
                   'GET:/api/event/info/rule'
    ]"
    class="dashboard-container"
  >
    <panel-group
      :db-total="dbTotal"
      :db-user-total="dbUserTotal"
      :event-total="eventTotal"
      :rule-total="ruleTotal"
    />
    <el-form :inline="true" :model="query" size="mini">
      <el-form-item>
        <el-date-picker
          v-model="queryTime"
          type="datetimerange"
          :picker-options="pickerOptions"
          range-separator="-"
          start-placeholder=""
          end-placeholder=""
          value-format="timestamp"
          align="right"
        />
      </el-form-item>
      <el-form-item label="数据库" prop="database">
        <el-select v-model="query.database" placeholder="请选择数据库">
          <el-option v-for="(item,index) in databases" :key="index" :label="item" :value="item" />
        </el-select>
      </el-form-item>
      <el-form-item>
        <el-button
          icon="el-icon-search"
          type="primary"
          @click="queryData"
        >查询</el-button>
        <el-button
          icon="el-icon-refresh"
          @click="reload"
        >重置</el-button>
      </el-form-item>
    </el-form>
    <el-row style="margin-top:30px">
      <el-card>
        <FingerPrintLineChart :data="fingerprintInfo" name="事件指纹分布" />
      </el-card>
    </el-row>
    <el-row style="margin-top:30px">
      <el-card>
        <el-table
          :data="fingerprintList"
          stripe
          border
          style="width: 100%"
        >
          <el-table-column align="center" type="index" label="序号" width="60" />
          <el-table-column align="center" prop="key" label="指纹" />
          <el-table-column align="center" prop="num" label="数量" width="80" />
        </el-table>
      </el-card>
    </el-row>
    <el-row style="margin-top:30px">
      <el-card>
        <FingerPrintLineChart :data="ruleFingerprintInfo" name="规则事件指纹分布" />
      </el-card>
    </el-row>
    <el-row style="margin-top:30px">
      <el-card>
        <el-table
          :data="ruleFingerprintList"
          stripe
          border
          style="width: 100%"
        >
          <el-table-column align="center" type="index" label="序号" width="60" />
          <el-table-column align="center" prop="key" label="指纹" />
          <el-table-column align="center" prop="num" label="数量" width="80" />
        </el-table>
      </el-card>
    </el-row>
  </div>
</template>

<script>

import PanelGroup from './components/PanelGroup'
import FingerPrintLineChart from './components/FingerPrintLineChart.vue'
import * as db from '@/api/db'
import * as dbuser from '@/api/dbuser'
import * as event from '@/api/event'

export default {
  name: 'Dashboard',
  components: { PanelGroup, FingerPrintLineChart },
  data() {
    return {
      dbTotal: 0,
      dbUserTotal: 0,
      eventTotal: 0,
      ruleTotal: 0,

      flag: false, // 判断是否显示图表组件
      categoryTotal: {}, // 每个分类下的文章数
      topInfo: {}, // 查询近6个月发布文章数

      query: {},
      databases: [],
      queryTime: [],
      pickerOptions: {
        shortcuts: [{
          text: '最近30分钟',
          onClick(picker) {
            const end = new Date()
            const start = new Date()
            start.setTime(start.getTime() - 1800 * 1000)
            picker.$emit('pick', [start, end])
          }
        }, {
          text: '最近一小时',
          onClick(picker) {
            const end = new Date()
            const start = new Date()
            start.setTime(start.getTime() - 3600 * 1000)
            picker.$emit('pick', [start, end])
          }
        }, {
          text: '最近24小时',
          onClick(picker) {
            const end = new Date()
            const start = new Date()
            start.setTime(start.getTime() - 3600 * 1000 * 24)
            picker.$emit('pick', [start, end])
          }
        },
        {
          text: '最近一周',
          onClick(picker) {
            const end = new Date()
            const start = new Date()
            start.setTime(start.getTime() - 3600 * 1000 * 24 * 7)
            picker.$emit('pick', [start, end])
          }
        },
        {
          text: '最近一个月',
          onClick(picker) {
            const end = new Date()
            const start = new Date()
            start.setTime(start.getTime() - 3600 * 1000 * 24 * 30)
            picker.$emit('pick', [start, end])
          }
        }]
      },
      fingerprintInfo: {},
      fingerprintList: [],
      ruleFingerprintInfo: {},
      ruleFingerprintList: []

    }
  },
  async created() {
    this.fetchData()
  },
  methods: {
    async fetchData() {
      this.queryTime[0] = new Date().getTime() - 3600 * 1000 * 24 * 7
      this.queryTime[1] = new Date().getTime()
      this.getDatabase()
      this.getDbUser()
      if (this.queryTime.length > 0) {
        this.query['start'] = this.queryTime[0]
        this.query['end'] = this.queryTime[1]
      }
      var res = await event.getInfo(this.query)
      this.fingerprintInfo = res['data']['data']
      var tmpData = this.transferInfoToList(this.fingerprintInfo)
      this.fingerprintList = tmpData.data
      this.eventTotal = tmpData.num

      res = await event.getRuleInfo(this.query)
      this.ruleFingerprintInfo = res['data']['data']
      tmpData = this.transferInfoToList(this.ruleFingerprintInfo)
      this.ruleFingerprintList = tmpData.data
      this.ruleTotal = tmpData.num
    },

    async getDatabase() {
      const res = await db.getList({}, 0, 0)
      const dbs = res['data']['list']

      this.databases = []
      var item
      for (item of dbs) {
        this.databases.push(item.name)
      }
      this.dbTotal = res['data']['total']
    },

    async getDbUser() {
      const res = await dbuser.getList({}, 0, 0)
      this.dbUserTotal = res['data']['total']
    },

    queryData() {
      if (this.queryTime.length > 0) {
        this.query['start'] = this.queryTime[0]
        this.query['end'] = this.queryTime[1]
      }
      this.fetchData()
    },

    reload() {
      this.query = {}
      this.queryTime[0] = new Date().getTime() - 3600 * 1000 * 24 * 7
      this.queryTime[1] = new Date().getTime()
      if (this.queryTime.length > 0) {
        this.query['start'] = this.queryTime[0]
        this.query['end'] = this.queryTime[1]
      }

      this.fetchData()
    },

    transferInfoToList(info) {
      const keys = info['item']
      const nums = info['num']
      var tmpData = { data: [], num: 0 }
      for (var i = 0; i < keys.length; i++) {
        const tmp = {
          'key': keys[i],
          'num': nums[i]
        }
        tmpData.num += nums[i]
        tmpData.data.push(tmp)
      }
      return tmpData
    }
  }
}
</script>

<style lang="scss" scoped>
.dashboard {
  &-container {
    margin: 30px;
  }
  &-text {
    font-size: 30px;
    line-height: 46px;
  }
}
</style>
