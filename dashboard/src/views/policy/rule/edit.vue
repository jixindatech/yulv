<template>
  <el-dialog
    :visible.sync="visible"
    width="500px"
    :before-close="handleClose"
  >
    <template slot="title">
      <span style="position:relative">
        <span>{{ title }}</span>
        <el-tooltip effect="dark" placement="top">
          <div slot="content">
            <p>匹配项中,需要至少匹配一个条件</p>
          </div>
          <i class="el-icon-question table-msg" />
        </el-tooltip>
      </span>
    </template>

    <el-form
      ref="formData"
      :rules="rules"
      :model="formData"
      label-width="100px"
      label-position="right"
      style="width: 400px"
      status-icon
    >
      <el-form-item label="规则名称" prop="name">
        <el-input v-model="formData.name" :disabled="typeof(formData.id) !== 'undefined' && formData.id !== 0" maxlength="30" />
      </el-form-item>
      <el-form-item prop="ip">
        <template slot="label">
          <span style="position:relative">
            <span>客户IP</span>
            <el-tooltip effect="dark" placement="top">
              <div slot="content">
                <p>支持CIDR形式的IP</p>
              </div>
              <i class="el-icon-question table-msg" />
            </el-tooltip>
          </span>
        </template>
        <el-input v-model="formData.ip" maxlength="254" />
      </el-form-item>
      <el-form-item label="数据库用户" prop="user">
        <el-input v-model="formData.user" maxlength="254" />
      </el-form-item>
      <el-form-item label="数据库" prop="database">
        <el-input v-model="formData.database" maxlength="254" />
      </el-form-item>
      <el-form-item label="请求方式" prop="type">
        <el-select v-model="formData.type" placeholder="请选择请求方式">
          <el-option v-for="(item,index) in SQL_TYPE_OPTIONS" :key="index" :label="item.label" :value="item.value" />
        </el-select>
      </el-form-item>
      <el-form-item label="规则水印" prop="sql">
        <el-input v-model="formData.sql" maxlength="254" />
      </el-form-item>
      <el-form-item prop="match">
        <template slot="label">
          <span style="position:relative">
            <span>匹配方式</span>
            <el-tooltip effect="dark" placement="top">
              <div slot="content">
                <p>选择了匹配方式之后,必须填写匹配内容</p>
              </div>
              <i class="el-icon-question table-msg" />
            </el-tooltip>
          </span>
        </template>
        <el-select v-model="formData.match" placeholder="请选择匹配方式">
          <el-option v-for="(item,index) in RULE_MATCH_OPTIONS" :key="index" :label="item.label" :value="item.value" />
        </el-select>
      </el-form-item>
      <el-form-item v-if="formData.match !== undefined" label="匹配内容" prop="pattern">
        <el-input v-model="formData.pattern" maxlength="254" />
      </el-form-item>
      <el-form-item prop="rows">
        <template slot="label">
          <span style="position:relative">
            <span>影响行数</span>
            <el-tooltip effect="dark" placement="top">
              <div slot="content">
                <p>默认影响行数为0,不进行匹配,匹配条件为大于或等于该值</p>
              </div>
              <i class="el-icon-question table-msg" />
            </el-tooltip>
          </span>
        </template>

        <el-input v-model.number="formData.rows" />
      </el-form-item>
      <el-form-item label="匹配动作" prop="action">
        <el-select v-model="formData.action" placeholder="请选择规则类型">
          <el-option v-for="(item,index) in RULE_TYPE_OPTIONS" :key="index" :label="item.label" :value="item.value" />
        </el-select>
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
import { add, update } from '@/api/rule'
import { RULE_TYPE_OPTIONS, RULE_MATCH_OPTIONS, SQL_TYPE_OPTIONS } from '@/utils/const'
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
      RULE_TYPE_OPTIONS,
      RULE_MATCH_OPTIONS,
      SQL_TYPE_OPTIONS,
      rules: {
        name: [{ required: true, message: '请输入名称', trigger: 'blur' }],
        action: [{ required: true, message: '请选择动作类型', trigger: 'change' }]
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
    }
  }
}
</script>
