<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="用户:">
          <el-input v-model.number="formData.uid" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="类型:">
          <el-input v-model.number="formData.type" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="订单号:">
          <el-input v-model="formData.orderNum" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="金额:">
          <el-input-number v-model="formData.money" :precision="2" clearable></el-input-number>
        </el-form-item>
        <el-form-item label="备注:">
          <el-input v-model="formData.remark" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="状态:">
          <el-input v-model.number="formData.status" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="日期:">
          <el-input v-model="formData.creationDate" clearable placeholder="请输入" />
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
  name: 'MoneyRecord'
}
</script>

<script setup>
import {
  createMoneyRecord,
  updateMoneyRecord,
  findMoneyRecord
} from '@/api/moneyRecord'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref } from 'vue'
const route = useRoute()
const router = useRouter()
const type = ref('')
const formData = ref({
        uid: 0,
        type: 0,
        orderNum: '',
        money: 0,
        remark: '',
        status: 0,
        creationDate: '',
        })

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findMoneyRecord({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.remoneyRecord
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
          res = await createMoneyRecord(formData.value)
          break
        case 'update':
          res = await updateMoneyRecord(formData.value)
          break
        default:
          res = await createMoneyRecord(formData.value)
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
