<template>
  <el-dialog
    :title="title"
    :visible.sync="visible"
    width="500px"
    :before-close="handleClose"
  >
    <el-form
      ref="formData"
      :rules="rules"
      :model="formData"
      label-width="100px"
      label-position="right"
      style="width: 400px"
      status-icon
    >
      <el-form-item label="IP名称" prop="name">
        <el-input v-model="formData.name" :disabled="typeof(formData.id) !== 'undefined' && formData.id !== 0" maxlength="30" />
      </el-form-item>
      <el-form-item label="规则类型" prop="type">
        <el-select v-model="formData.type" placeholder="请选择规则类型">
          <el-option v-for="(item,index) in IP_TYPE_OPTIONS" :key="index" :label="item.label" :value="item.value" />
        </el-select>
      </el-form-item>
      <el-form-item>
        <span slot="label">IP
          <el-tooltip placement="right">
            <div slot="content">
              支持CIDR方式
            </div>
            <i class="el-icon-question" />
          </el-tooltip>
        </span>

        <el-table
          :row-style="{height:'10px'}"
          :cell-style="{padding:'1px'}"
          style="font-size: 8px; margin-top: 0px;"
          size="mini"
          :show-header="false"
          :data="ips"
        >
          <el-table-column align="center" label="IP" width="270px">
            <template slot-scope="scope">
              <el-form-item :prop="'ips.' + scope.$index + '.ip'" :rules="rules.ip">
                <el-input v-model="scope.row.ip" size="mini" placeholder="请输入IP" />
              </el-form-item>
            </template>
          </el-table-column>
          <el-table-column align="center" width="30px">
            <template slot-scope="scope">
              <el-button type="text" icon="el-icon-delete" size="medium" @click="deleteIP(scope.row, scope.$index)" />
            </template>
          </el-table-column>
        </el-table>
        <el-button type="text" icon="el-icon-plus" size="mini" style="margin-bottom: 20px;" @click="addIP()">新增IP</el-button>
      </el-form-item>
      <el-form-item label="备注：" prop="remark">
        <el-input v-model="formData.remark" type="textarea" />
      </el-form-item>
    </el-form>
    <div slot="footer" class="dialog-footer">
      <el-button
        type="primary"
        size=""
        @click="submitForm('formData')"
      >确定</el-button>
      <el-button size="" @click="handleClose">取消</el-button>
    </div>
  </el-dialog>
</template>

<script>
import { add, update } from '@/api/ip'
import { validIP } from '@/utils/validate'
import { IP_TYPE_OPTIONS } from '@/utils/const'

const checkIP = (rule, value, callback) => {
  const items = value.split('\n')
  for (const i in items) {
    if (validIP(items[i]) === false) {
      return callback(new Error('请输入正确的IP地址'))
    }
  }
  return callback()
}

export default {
  props: {
    title: {
      type: String,
      default: ''
    },
    visible: {
      type: Boolean,
      default: false
    },
    formData: {
      type: Object,
      default: function() { return {} }
    },
    remoteClose: {
      type: Function,
      default: function() {}
    }
  },

  data() {
    return {
      IP_TYPE_OPTIONS,
      ips: [],
      rules: {
        name: [{ required: true, message: '请输入名称', trigger: 'blur' }],
        type: [{ required: true, message: '请选择类型', trigger: 'blur' }],
        ip: [
          { required: true, message: '请输入正确IP', trigger: 'blur' },
          { validator: checkIP }
        ]
      }
    }
  },
  watch: {
    visible(newVal, oldVal) {
      if (newVal) {
        this.ips = []
        if (this.formData.ip === undefined) {
          this.formData.ip = []
          const item = { ip: '' }
          this.ips.push(item)
        } else {
          this.formData.ip.forEach(element => {
            const item = { ip: element }
            this.ips.push(item)
          })
        }
        this.formData.ips = this.ips
      }
    }
  },
  methods: {
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.submitData()
        } else {
          return false
        }
      })
    },

    async submitData() {
      let response = null

      this.formData.ip = []
      this.ips.forEach(element => {
        this.formData.ip.push(element.ip)
      })

      if (this.formData.id) {
        response = await update(this.formData.id, this.formData)
      } else {
        response = await add(this.formData)
      }

      if ((response.code === 0)) {
        this.$message({ message: '保存成功', type: 'success' })
        this.handleClose()
      } else {
        this.$message({ message: '保存失败', type: 'error' })
      }
    },

    handleClose() {
      this.$refs['formData'].resetFields()
      this.remoteClose()
    },
    addIP() {
      const item = { ip: '' }
      this.ips.push(item)
    },
    deleteIP(row, index) {
      this.ips.splice(index, 1)
    }

  }
}
</script>
