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
      <el-form-item label="规则内容" prop="sql">
        <el-input v-model="formData.sql" maxlength="254" />
      </el-form-item>
      <el-form-item label="规则水印" prop="res">
        <el-input v-model="formData.res" maxlength="254" />
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
import { test } from '@/api/rule'
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
    remoteClose: {
      type: Function,
      default: function() {}
    }
  },

  data() {
    return {
      formData: {
        sql: '',
        res: ''
      },
      rules: {
        sql: [{ required: true, message: '请输入名称', trigger: 'blur' }]
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
      response = await test(this.formData)
      this.formData.res = response.data.item
      if ((response.code === 0)) {
        this.$message({ message: '测试成功', type: 'success' })
      } else {
        this.$message({ message: '测试失败', type: 'error' })
      }
    },
    handleClose() {
      this.$refs['formData'].resetFields()
      this.remoteClose()
    }
  }
}
</script>
