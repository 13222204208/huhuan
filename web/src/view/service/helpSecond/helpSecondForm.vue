<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="订单号:">
          <el-input v-model="formData.orderNum" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="商品详情:">
          <el-input v-model="formData.goodsDetail" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="商品图片:">
          <el-input v-model="formData.goodsImg" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="商品名称:">
          <el-input v-model="formData.goodsName" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="商品原价:">
          <el-input-number v-model="formData.originaPrice" :precision="2" clearable></el-input-number>
        </el-form-item>
        <el-form-item label="商品售价:">
          <el-input-number v-model="formData.price" :precision="2" clearable></el-input-number>
        </el-form-item>
        <el-form-item label="我的地址:">
          <el-input v-model="formData.address" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="接单人:">
          <el-input v-model="formData.takeName" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="接单人电话:">
          <el-input v-model="formData.takePhone" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="接单时间:">
          <el-input v-model="formData.takeTime" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="用户编号:">
          <el-input v-model="formData.userNum" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="订单状态:">
          <el-select v-model="formData.status" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in orderStateOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="城市:">
          <el-input v-model="formData.city" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="地区:">
          <el-input v-model="formData.area" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="类别:">
          <el-input v-model="formData.category" clearable placeholder="请输入" />
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
  name: 'HelpSecond'
}
</script>

<script setup>
import {
  createHelpSecond,
  updateHelpSecond,
  findHelpSecond
} from '@/api/helpSecond'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref } from 'vue'
const route = useRoute()
const router = useRouter()
const type = ref('')
const orderStateOptions = ref([])
const formData = ref({
        orderNum: '',
        goodsDetail: '',
        goodsImg: '',
        goodsName: '',
        originaPrice: 0,
        price: 0,
        address: '',
        takeName: '',
        takePhone: '',
        takeTime: '',
        userNum: '',
        status: undefined,
        city: '',
        area: '',
        category: '',
        })

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findHelpSecond({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rehelpSecond
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    orderStateOptions.value = await getDictFunc('orderState')
}

init()
// 保存按钮
const save = async() => {
      let res
      switch (type.value) {
        case 'create':
          res = await createHelpSecond(formData.value)
          break
        case 'update':
          res = await updateHelpSecond(formData.value)
          break
        default:
          res = await createHelpSecond(formData.value)
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
