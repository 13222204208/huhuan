<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="名称:">
          <el-input v-model="formData.title" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="使用条件:">
          <el-input v-model.number="formData.condition" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="开始时间:">
          <el-input v-model="formData.start" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="结束时间:">
          <el-input v-model="formData.over" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="发放总数:">
          <el-input v-model.number="formData.send" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="已发放:">
          <el-input v-model.number="formData.sent" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="积分兑换:">
          <el-input v-model.number="formData.integral" clearable placeholder="请输入" />
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
  name: 'Coupon'
}
</script>

<script setup>
import {
  createCoupon,
  updateCoupon,
  findCoupon
} from '@/api/coupon'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref } from 'vue'
const route = useRoute()
const router = useRouter()
const type = ref('')
const formData = ref({
        title: '',
        condition: 0,
        start: '',
        over: '',
        send: 0,
        sent: 0,
        integral: 0,
        })

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findCoupon({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.recoupon
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
          res = await createCoupon(formData.value)
          break
        case 'update':
          res = await updateCoupon(formData.value)
          break
        default:
          res = await createCoupon(formData.value)
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
