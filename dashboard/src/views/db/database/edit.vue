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
      <el-form-item label="数据库名称" prop="name">
        <el-input v-model="formData.name" :disabled="typeof(formData.id) !== 'undefined' && formData.id !== 0" maxlength="30" />
      </el-form-item>
      <el-form-item label="数据库用户" prop="user">
        <el-input v-model="formData.user" maxlength="254" />
      </el-form-item>
      <el-form-item label="数据库密码" prop="password">
        <el-input v-model="formData.password" maxlength="254" />
      </el-form-item>
      <el-form-item label="数据库主机" prop="host">
        <el-input v-model="formData.host" maxlength="254" />
      </el-form-item>
      <el-form-item label="数据库端口" prop="port">
        <el-input v-model.number="formData.port" />
      </el-form-item>
      <el-form-item label="字符集" prop="character">
        <el-select v-model="formData.charset" placeholder="请选择字符集" @change="charsetChanged">
          <el-option v-for="(item,index) in SQL_CHARSET" :key="index" :label="item.label" :value="item.value" />
        </el-select>
      </el-form-item>
      <el-form-item label="字符序" prop="collation">
        <el-select v-model="formData.collation" placeholder="请选择字符序">
          <el-option v-for="(item,index) in collation" :key="index" :label="item.label" :value="item.value" />
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
import { add, update } from '@/api/db'
import { SQL_CHARSET } from '@/utils/const'
import {
  SQL_COLLATION_BIG5,
  SQL_COLLATION_DEC8,
  SQL_COLLATION_CP850,
  SQL_COLLATION_HP8,
  SQL_COLLATION_KOI8R,
  SQL_COLLATION_LATIN1,
  SQL_COLLATION_LATIN2,
  SQL_COLLATION_SWE7,
  SQL_COLLATION_ASCII,
  SQL_COLLATION_UJIS,
  SQL_COLLATION_SJIS,
  SQL_COLLATION_HEBREW,
  SQL_COLLATION_TIS620,
  SQL_COLLATION_EUCKR,
  SQL_COLLATION_KOI8U,
  SQL_COLLATION_GB2312,
  SQL_COLLATION_GREEK,
  SQL_COLLATION_CP1250,
  SQL_COLLATION_GBK,
  SQL_COLLATION_LATIN5,
  SQL_COLLATION_ARMSCII8,
  SQL_COLLATION_UTF8,
  SQL_COLLATION_UCS2,
  SQL_COLLATION_CP886,
  SQL_COLLATION_KEYBCS2,
  SQL_COLLATION_MACCE,
  SQL_COLLATION_MACROMAN,
  SQL_COLLATION_CP852,
  SQL_COLLATION_LATIN7,
  SQL_COLLATION_UTF8MB4,
  SQL_COLLATION_CP1251,
  SQL_COLLATION_UTF16,
  SQL_COLLATION_UTF16LE,
  SQL_COLLATION_CP1256,
  SQL_COLLATION_CP1257,
  SQL_COLLATION_UTF32,
  SQL_COLLATION_BINARAY,
  SQL_COLLATION_GEOSTD8,
  SQL_COLLATION_CP932,
  SQL_COLLATION_EUCJPMS
} from '@/utils/const'
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
      SQL_CHARSET,
      collation: [],
      SQL_COLLATION_BIG5,
      SQL_COLLATION_DEC8,
      SQL_COLLATION_CP850,
      SQL_COLLATION_HP8,
      SQL_COLLATION_KOI8R,
      SQL_COLLATION_LATIN1,
      SQL_COLLATION_LATIN2,
      SQL_COLLATION_SWE7,
      SQL_COLLATION_ASCII,
      SQL_COLLATION_UJIS,
      SQL_COLLATION_SJIS,
      SQL_COLLATION_HEBREW,
      SQL_COLLATION_TIS620,
      SQL_COLLATION_EUCKR,
      SQL_COLLATION_KOI8U,
      SQL_COLLATION_GB2312,
      SQL_COLLATION_GREEK,
      SQL_COLLATION_CP1250,
      SQL_COLLATION_GBK,
      SQL_COLLATION_LATIN5,
      SQL_COLLATION_ARMSCII8,
      SQL_COLLATION_UTF8,
      SQL_COLLATION_UCS2,
      SQL_COLLATION_CP886,
      SQL_COLLATION_KEYBCS2,
      SQL_COLLATION_MACCE,
      SQL_COLLATION_MACROMAN,
      SQL_COLLATION_CP852,
      SQL_COLLATION_LATIN7,
      SQL_COLLATION_UTF8MB4,
      SQL_COLLATION_CP1251,
      SQL_COLLATION_UTF16,
      SQL_COLLATION_UTF16LE,
      SQL_COLLATION_CP1256,
      SQL_COLLATION_CP1257,
      SQL_COLLATION_UTF32,
      SQL_COLLATION_BINARAY,
      SQL_COLLATION_GEOSTD8,
      SQL_COLLATION_CP932,
      SQL_COLLATION_EUCJPMS,
      rules: {
        name: [{ required: true, message: '请输入名称', trigger: 'blur' }],
        user: [{ required: true, message: '请输入用户', trigger: 'blur' }],
        password: [{ required: true, message: '请输入密码', trigger: 'blur' }],
        host: [{ required: true, message: '请输入主机地址', trigger: 'blur' }],
        port: [{ required: true, message: '请输入主机端口', trigger: 'change' }],
        status: [{ required: true, message: '请选择', trigger: 'change' }]
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
    },
    charsetChanged(value) {
      if (value === 'big5') {
        this.collation = SQL_COLLATION_BIG5
      } else if (value === 'dec8') {
        this.collation = SQL_COLLATION_DEC8
      } else if (value === 'cp850') {
        this.collation = SQL_COLLATION_CP850
      } else if (value === 'hp8') {
        this.collation = SQL_COLLATION_HP8
      } else if (value === 'koi8r') {
        this.collation = SQL_COLLATION_KOI8R
      } else if (value === 'latin1') {
        this.collation = SQL_COLLATION_LATIN1
      } else if (value === 'latin2') {
        this.collation = SQL_COLLATION_LATIN2
      } else if (value === 'swe7') {
        this.collation = SQL_COLLATION_SWE7
      } else if (value === 'ascii') {
        this.collation = SQL_COLLATION_ASCII
      } else if (value === 'ujis') {
        this.collation = SQL_COLLATION_UJIS
      } else if (value === 'sjis') {
        this.collation = SQL_COLLATION_SJIS
      } else if (value === 'hebrew') {
        this.collation = SQL_COLLATION_HEBREW
      } else if (value === 'tis620') {
        this.collation = SQL_COLLATION_TIS620
      } else if (value === 'euckr') {
        this.collation = SQL_COLLATION_EUCKR
      } else if (value === 'koi8u') {
        this.collation = SQL_COLLATION_KOI8U
      } else if (value === 'gb2312') {
        this.collation = SQL_COLLATION_GB2312
      } else if (value === 'greek') {
        this.collation = SQL_COLLATION_GREEK
      } else if (value === 'cp1250') {
        this.collation = SQL_COLLATION_CP1250
      } else if (value === 'gbk') {
        this.collation = SQL_COLLATION_GBK
      } else if (value === 'latin5') {
        this.collation = SQL_COLLATION_LATIN5
      } else if (value === 'armscii8') {
        this.collation = SQL_COLLATION_ARMSCII8
      } else if (value === 'utf8') {
        this.collation = SQL_COLLATION_UTF8
      } else if (value === 'ucs2') {
        this.collation = SQL_COLLATION_UCS2
      } else if (value === 'cp866') {
        this.collation = SQL_COLLATION_CP886
      } else if (value === 'keybcs2') {
        this.collation = SQL_COLLATION_KEYBCS2
      } else if (value === 'macce') {
        this.collation = SQL_COLLATION_MACCE
      } else if (value === 'cp852') {
        this.collation = SQL_COLLATION_CP852
      } else if (value === 'latin7') {
        this.collation = SQL_COLLATION_LATIN7
      } else if (value === 'utf8mb4') {
        this.collation = SQL_COLLATION_UTF8MB4
      } else if (value === 'cp1251') {
        this.collation = SQL_COLLATION_CP1251
      } else if (value === 'utf16') {
        this.collation = SQL_COLLATION_UTF16
      } else if (value === 'utf16le') {
        this.collation = SQL_COLLATION_UTF16LE
      } else if (value === 'cp1256') {
        this.collation = SQL_COLLATION_CP1251
      } else if (value === 'cp1257') {
        this.collation = SQL_COLLATION_CP1257
      } else if (value === 'utf32') {
        this.collation = SQL_COLLATION_UTF32
      } else if (value === 'binary') {
        this.collation = SQL_COLLATION_BINARAY
      } else if (value === 'geostd8') {
        this.collation = SQL_COLLATION_GEOSTD8
      } else if (value === 'cp932') {
        this.collation = SQL_COLLATION_CP932
      } else if (value === 'eucjpms') {
        this.collation = SQL_COLLATION_EUCJPMS
      } else {
        this.collation = []
      }
    }
  }
}
</script>
