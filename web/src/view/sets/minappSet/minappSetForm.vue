<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="appid:">
          <el-input v-model="formData.appid" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="秘钥:">
          <el-input v-model="formData.app_secret" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="商户号:">
          <el-input v-model="formData.mch_id" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="支付秘钥:">
          <el-input v-model="formData.pay_secret" clearable placeholder="请输入" />
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
  name: 'MinappSet'
}
</script>

<script setup>
import {
  createMinappSet,
  updateMinappSet,
  findMinappSet
} from '@/api/minappSet'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref } from 'vue'
const route = useRoute()
const router = useRouter()
const type = ref('')
const formData = ref({
        appid: '',
        app_secret: '',
        mch_id: '',
        pay_secret: '',
        })

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findMinappSet({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reminappSet
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
          res = await createMinappSet(formData.value)
          break
        case 'update':
          res = await updateMinappSet(formData.value)
          break
        default:
          res = await createMinappSet(formData.value)
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
