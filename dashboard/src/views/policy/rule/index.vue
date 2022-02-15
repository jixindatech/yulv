<template>
  <div
    v-permission="['GET:/api/rule']"
    class="app-container"
  >
    <el-form :inline="true" :model="query" size="mini">
      <el-form-item label="规则名称:">
        <el-input v-model.trim="query.name" />
      </el-form-item>
      <el-form-item label="规则类型" prop="action">
        <el-select v-model="query.action" placeholder="请选择规则类型">
          <el-option v-for="(item,index) in RULE_TYPE_OPTIONS" :key="index" :label="item.label" :value="item.value" />
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
        <el-button
          v-permission="['POST:/api/rule']"
          icon="el-icon-circle-plus-outline"
          type="primary"
          @click="openAdd"
        >新增</el-button>
        <el-button
          v-permission="['POST:/api/rule/distribute']"
          icon="el-icon-circle-plus-outline"
          type="success"
          @click="distribute"
        >下发配置</el-button>
        <el-button
          v-permission="['GET:/api/rule/sql/test']"
          icon="el-icon-circle-plus-outline"
          type="info"
          @click="testSql"
        >测试水印</el-button>
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
      row-key="id"
    >
      <el-table-column prop="name" label="规则名称" />
      <el-table-column prop="sql" label="请求水印" />
      <el-table-column prop="action" label="匹配动作">
        <template slot-scope="scope">
          {{ RULE_TYPE_MAP[scope.row.action] }}
        </template>
      </el-table-column>
      <el-table-column prop="ip" label="客户端IP" />
      <el-table-column prop="user" label="数据库用户" />
      <el-table-column prop="database" label="数据库" />
      <el-table-column prop="updateAt" label="创建时间" width="220">
        <template slot-scope="scope">
          <i class="el-icon-time" />
          <span>{{ scope.row.createdAt }}</span>
        </template>
      </el-table-column>
      <el-table-column prop="updateAt" label="更新时间" width="220">
        <template slot-scope="scope">
          <i class="el-icon-time" />
          <span>{{ scope.row.updateAt }}</span>
        </template>
      </el-table-column>
      <el-table-column align="center" label="操作" width="250">
        <template slot-scope="scope">
          <el-button
            v-permission="['PUT:/api/rule/:id']"
            type="success"
            size="mini"
            @click="handleEdit(scope.row.id)"
          >编辑</el-button>
          <el-button
            v-permission="['DELETE:/api/rule/:id']"
            type="danger"
            size="mini"
            @click="handleDelete(scope.row.id)"
          >删除</el-button>
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

    <edit
      :title="edit.title"
      :form-data="edit.formData"
      :visible="edit.visible"
      :remote-close="remoteClose"
    />

    <TestRule
      :title="test.title"
      :visible="test.visible"
      :remote-close="testClose"
    />

  </div>
</template>

<script>
import { getList, deleteById, getById, Distribute } from '@/api/rule'
import Edit from './edit'
import TestRule from './test'
import { RULE_TYPE_OPTIONS, RULE_TYPE_MAP } from '@/utils/const'
export default {
  components: { Edit, TestRule },
  data() {
    return {
      RULE_TYPE_OPTIONS,
      RULE_TYPE_MAP,
      query: {},
      edit: {
        title: '',
        visible: false,
        formData: {}
      },
      test: {
        title: '',
        visible: false
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
    this.fetchData()
  },
  methods: {
    async fetchData() {
      this.listLoading = true
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
      this.fetchData()
    },
    openAdd() {
      this.edit.title = '新增规则'
      this.edit.visible = true
    },
    remoteClose() {
      this.edit.formData = {}
      this.edit.visible = false
      this.fetchData()
    },
    handleSizeChange(val) {
      this.page.size = val
      this.fetchData()
    },
    handleCurrentChange(val) {
      this.page.current = val
      this.fetchData()
    },
    handleEdit(id) {
      getById(id).then((response) => {
        const { data } = response
        this.edit.formData = data.item
        this.edit.title = '编辑'
        this.edit.visible = true
      })
    },
    handleDelete(id) {
      this.$confirm('确认删除这条记录吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(() => {
          deleteById(id).then((response) => {
            this.$message({
              type: response.code === 0 ? 'success' : 'error',
              message: '删除成功!'
            })
            this.fetchData()
          })
        })
        .catch(() => {
        })
    },
    testSql() {
      this.test.title = '测试规则'
      this.test.visible = true
    },
    testClose() {
      this.test.visible = false
    },
    distribute() {
      Distribute().then((response) => {
        if (response.code === 0) {
          this.$message({ message: '下发成功', type: 'success' })
        } else {
          this.$message({ message: '下发失败', type: 'error' })
        }
      })
    }
  }
}
</script>
