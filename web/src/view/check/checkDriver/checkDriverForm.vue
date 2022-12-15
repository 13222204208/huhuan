<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="姓名:">
          <el-input v-model="formData.name" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="联系方式:">
          <el-input v-model="formData.phone" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="驾照:">
          <el-input v-model="formData.licence" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="营业执照:">
          <el-input v-model="formData.charter" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="邮箱:">
          <el-input v-model="formData.mail" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="申请日期:">
          <el-input v-model="formData.time" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="用户ID:">
          <el-input v-model.number="formData.uid" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="状态:">
          <el-input v-model.number="formData.status" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" @click="save">保存</el-button>
          <el-button size="mini" type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'CheckDriver'
}
</script>

<script setup>
import {
  createCheckDriver,
  updateCheckDriver,
  findCheckDriver
} from '@/api/checkDriver'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref } from 'vue'
const route = useRoute()
const router = useRouter()
const type = ref('')
const formData = ref({
        name: '',
        phone: '',
        licence: '',
        charter: '',
        mail: '',
        time: '',
        uid: 0,
        status: 0,
        })

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findCheckDriver({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.recheckDriver
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      let res
      switch (type.value) {
        case 'create':
          res = await createCheckDriver(formData.value)
          break
        case 'update':
          res = await updateCheckDriver(formData.value)
          break
        default:
          res = await createCheckDriver(formData.value)
          break
      }
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '创建/更改成功'
        })
      }
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
