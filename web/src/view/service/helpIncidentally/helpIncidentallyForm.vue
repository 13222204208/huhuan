<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="订单号:">
          <el-input v-model="formData.orderNum" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="物品名称:">
          <el-input v-model="formData.goodsName" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="物品数量:">
          <el-input v-model="formData.goodsCount" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="联系人:">
          <el-input v-model="formData.linkman" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="联系电话:">
          <el-input v-model="formData.contactPhone" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="取货地址:">
          <el-input v-model="formData.fetchAddress" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="收货地址:">
          <el-input v-model="formData.receiveAddress" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="出发时间:">
          <el-input v-model="formData.time" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="留言:">
          <el-input v-model="formData.remark" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="优惠券ID:">
          <el-input v-model.number="formData.couponId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="优惠金额:">
          <el-input-number v-model="formData.couponAmount" :precision="2" clearable></el-input-number>
        </el-form-item>
        <el-form-item label="赏金:">
          <el-input-number v-model="formData.price" :precision="2" clearable></el-input-number>
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
        <el-form-item label="商品价值:">
          <el-input-number v-model="formData.goodsPrice" :precision="2" clearable></el-input-number>
        </el-form-item>
        <el-form-item label="订单状态:">
          <el-select v-model="formData.status" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in orderStateOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="完成时间:">
          <el-input v-model="formData.doneTime" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="城市:">
          <el-input v-model="formData.city" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="地区:">
          <el-input v-model="formData.area" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="类型:">
          <el-input v-model.number="formData.type" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="出发城市:">
          <el-input v-model="formData.startCity" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="抵达城市:">
          <el-input v-model="formData.arriveCity" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="航班号:">
          <el-input v-model="formData.flightNumber" clearable placeholder="请输入" />
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
  name: 'HelpIncidentally'
}
</script>

<script setup>
import {
  createHelpIncidentally,
  updateHelpIncidentally,
  findHelpIncidentally
} from '@/api/helpIncidentally'

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
        goodsName: '',
        goodsCount: '',
        linkman: '',
        contactPhone: '',
        fetchAddress: '',
        receiveAddress: '',
        time: '',
        remark: '',
        couponId: 0,
        couponAmount: 0,
        price: 0,
        takeName: '',
        takePhone: '',
        takeTime: '',
        userNum: '',
        goodsPrice: 0,
        status: undefined,
        doneTime: '',
        city: '',
        area: '',
        type: 0,
        startCity: '',
        arriveCity: '',
        flightNumber: '',
        })

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findHelpIncidentally({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rehelpIncidentally
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
          res = await createHelpIncidentally(formData.value)
          break
        case 'update':
          res = await updateHelpIncidentally(formData.value)
          break
        default:
          res = await createHelpIncidentally(formData.value)
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
