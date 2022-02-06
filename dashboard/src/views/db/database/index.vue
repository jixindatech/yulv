<template>
  <div
    class="app-container"
  >
    <el-form :inline="true" :model="query" size="mini">
      <el-form-item label="数据库名称:">
        <el-input v-model.trim="query.name" />
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
          v-if="!userDbs"
          icon="el-icon-circle-plus-outline"
          type="primary"
          @click="openAdd"
        >新增</el-button>
        <el-button
          v-if="userDbs"
          icon="el-icon-circle-plus-outline"
          type="success"
          @click="handleUserDbs"
        >关联数据库</el-button>
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
      @selection-change="handleSelectionChange"
    >
      <el-table-column
        v-if="userDbs"
        align="center"
        reserve-selection
        type="selection"
        width="55"
      />
      <el-table-column prop="name" label="数据库名称" />
      <el-table-column prop="user" label="数据库用户" />
      <el-table-column prop="host" label="数据库主机" />
      <el-table-column prop="port" label="数据库端口" />
      <el-table-column v-if="!userDbs" prop="updateAt" label="创建时间" width="220">
        <template slot-scope="scope">
          <i class="el-icon-time" />
          <span>{{ scope.row.createdAt }}</span>
        </template>
      </el-table-column>
      <el-table-column v-if="!userDbs" prop="updateAt" label="更新时间" width="220">
        <template slot-scope="scope">
          <i class="el-icon-time" />
          <span>{{ scope.row.updateAt }}</span>
        </template>
      </el-table-column>
      <el-table-column v-if="!userDbs" align="center" label="操作" width="250">
        <template slot-scope="scope">
          <el-button
            type="success"
            size="mini"
            @click="handleEdit(scope.row.id)"
          >编辑</el-button>
          <el-button
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

  </div>
</template>

<script>
import { getList, deleteById, getById } from '@/api/db'
import Edit from './edit'
export default {
  components: { Edit },
  props: {
    userDbs: {
      type: Array,
      default: function() { return null }
    }
  },
  data() {
    return {
      query: {},
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
      listLoading: true,
      checkedUserDbList: []
    }
  },
  watch: {
    userDbs(newVal, oldVal) {
      console.log('---user: ', this.userDbs)
      console.log('watch :', newVal)
      if (newVal !== null) {
        this.query = {}
        this.queryData()
      }
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
      this.checkedUserDb()
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
      this.edit.title = '新增数据库'
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
    handleSelectionChange(val) {
      this.checkedUserDbList = val
    },
    checkedUserDb() {
      this.$refs.dataTable.clearSelection()
      if (this.userDbs) {
        this.list.forEach((item) => {
          if (this.userDbs.indexOf(item.id) !== -1) {
            this.$refs.dataTable.toggleRowSelection(item, true)
          }
        })
      }
    },
    handleUserDbs() {
      const checkedDbs = []
      this.checkedUserDbList.forEach((item) => {
        checkedDbs.push(item.id)
      })
      this.checkedUserDbList = []
      this.$emit('saveUserDb', checkedDbs)
    }
  }
}
</script>
