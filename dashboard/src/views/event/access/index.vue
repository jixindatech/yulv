<template>
  <div
    v-permission="['GET:/api/event/access']"
    class="app-container"
  >
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

    <el-table
      ref="dataTable"
      v-loading="listLoading"
      :data="list"
      element-loading-text="Loading"
      :header-cell-style="{'text-align':'center'}"
      :cell-style="{'text-align':'center'}"
      border
      fit
      highlight-current-row
    >
      <el-table-column align="center" type="index" label="序号" width="60" />
      <el-table-column align="center" prop="_source.ip" label="IP" />
      <el-table-column align="center" prop="_source.user" label="数据库用户" />
      <el-table-column align="center" prop="_source.database" label="数据库" />
      <el-table-column align="center" prop="_source.event" label="事件类型" />
      <el-table-column align="center" prop="_source.cmd" label="请求类型">
        <template slot-scope="scope">
          <el-input :value="SQL_CMD_OPTIONS[scope.row._source.cmd]" />
        </template>
      </el-table-column>
      <el-table-column align="center" prop="_source.sql" label="请求内容" />
      <el-table-column align="center" prop="_source.fingerprint" label="请求水印" />
      <el-table-column align="center" prop="_source.rows" label="影响行数" />
      <el-table-column align="center" prop="_source.timestamp" label="时间">
        <template slot-scope="scope">
          {{ new Date(scope.row._source.timestamp * 1000).toLocaleString() }}
        </template>
      </el-table-column>

    </el-table>
    <el-pagination
      :current-page="page.current"
      :page-sizes="[10, 20, 50]"
      :page-size="page.size"
      layout="total, sizes, prev, pager, next, jumper"
      :total="page.total"
      @size-change="handleSizeChange"
      @current-change="handleCurrentChange"
    />
  </div>
</template>

<script>
import { getList } from '@/api/event'
import { SQL_CMD_OPTIONS } from '@/utils/const'
export default {
  data() {
    return {
      SQL_CMD_OPTIONS,
      queryTime: [],
      query: {},
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
      edit: {
        title: '',
        visible: false,
        formData: {}
      },
      page: {
        current: 1,
        size: 10,
        total: 0
      },
      list: [],
      listLoading: true
    }
  },
  created() {
    this.queryTime[0] = new Date().getTime() - 3600 * 1000 * 24 * 7
    this.queryTime[1] = new Date().getTime()
    this.fetchData()
  },
  methods: {
    async fetchData() {
      this.listLoading = true
      if (this.queryTime.length > 0) {
        this.query['start'] = this.queryTime[0]
        this.query['end'] = this.queryTime[1]
      }
      await getList(
        this.query,
        this.page.current,
        this.page.size
      ).then(response => {
        const { data } = response
        this.list = data.list
        this.page.total = data.total
        this.listLoading = false
      })
    },
    queryData() {
      this.page.current = 1
      this.fetchData()
    },
    reload() {
      this.query = {}
      this.queryTime[0] = new Date().getTime() - 3600 * 1000 * 24 * 7
      this.queryTime[1] = new Date().getTime()
      this.fetchData()
    },
    handleSizeChange(val) {
      this.page.size = val
      this.fetchData()
    },
    handleCurrentChange(val) {
      this.page.current = val
      this.fetchData()
    }
  }
}
</script>
