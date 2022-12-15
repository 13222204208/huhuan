<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="订单号:">
          <el-input v-model="formData.orderNum" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="车主姓名:">
          <el-input v-model="formData.ownerName" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="乘车人姓名:">
          <el-input v-model="formData.riderName" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="姓别:">
          <el-input v-model="formData.gender" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="联系电话:">
          <el-input v-model="formData.contactPhone" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="接单时间:">
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
        <el-form-item label="费用/几人:">
          <el-input v-model="formData.price" clearable placeholder="请输入" />
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
        <el-form-item label="乘车人数:">
          <el-input v-model.number="formData.rederNum" clearable placeholder="请输入" />
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
        <el-form-item label="出发地址:">
          <el-input v-model="formData.startAddress" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="送达地址:">
          <el-input v-model="formData.arriveAddress" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="车牌号:">
          <el-input v-model="formData.carNumber" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="用车人数:">
          <el-input v-model.number="formData.useNum" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="车辆已预约人数:">
          <el-input v-model.number="formData.carUsedNum" clearable placeholder="请输入" />
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
  name: 'HelpCarpool'
}
</script>

<script setup>
import {
  createHelpCarpool,
  updateHelpCarpool,
  findHelpCarpool
} from '@/api/helpCarpool'

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
        ownerName: '',
        riderName: '',
        gender: '',
        contactPhone: '',
        time: '',
        remark: '',
        couponId: 0,
        couponAmount: 0,
        price: '',
        takeName: '',
        takePhone: '',
        takeTime: '',
        userNum: '',
        rederNum: 0,
        status: undefined,
        doneTime: '',
        city: '',
        area: '',
        type: 0,
        startAddress: '',
        arriveAddress: '',
        carNumber: '',
        useNum: 0,
        carUsedNum: 0,
        })

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findHelpCarpool({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rehelpCarpool
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
          res = await createHelpCarpool(formData.value)
          break
        case 'update':
          res = await updateHelpCarpool(formData.value)
          break
        default:
          res = await createHelpCarpool(formData.value)
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
