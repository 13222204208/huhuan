<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="头像:">
          <!-- <el-input v-model="formData.avatar" clearable placeholder="请输入" /> -->
              <CustomPic style="margin-top:8px" :pic-src="formData.avatar" />
        </el-form-item>
        <el-form-item label="昵称:">
          <!-- <el-input v-model="formData.nickname" disabled clearable placeholder="请输入" /> -->
          {{formData.nickname}}
        </el-form-item>
        <el-form-item label="手机号:">
          <!-- <el-input v-model="formData.phone" disabled clearable placeholder="请输入" /> -->
          {{formData.phone}}
        </el-form-item>
        <el-form-item label="地区:">
          <!-- <el-input v-model.number="formData.area_id" disabled clearable placeholder="请输入" /> -->
          {{formData.area}}
        </el-form-item>
        <el-form-item label="注册时间:">
         {{ formatDate(formData.CreatedAt) }}
        </el-form-item>
        <el-form-item label="用户编号:">
          {{formData.userNum}}
        </el-form-item>
<!--        <el-form-item label="余额:"> 
          {{formData.balance}}  <el-button style="margin-left:10px" type="primary"  @click="topUp">充值</el-button>
        </el-form-item> -->
		
		<el-form-item label="余额:">
		  {{formData.balance}}  <el-button style="margin-left:10px" type="primary"  @click="topUp">编辑</el-button>
		</el-form-item>
		
        <el-form-item label="明细:">
           <el-button type="primary" :icon="Search" @click="toCoupon">优惠券明细</el-button>
        </el-form-item>
 <!--       <el-form-item label="状态:">
          <el-select v-model="formData.status" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in intOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item> -->
        <el-form-item>
          <el-button  type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>

      <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
      <el-form :model="formData" label-position="right" label-width="80px">
          <el-form-item label="余额:">
          {{formData.balance}}
        </el-form-item>
        <el-form-item label="更改:">
          <el-radio-group v-model="change" class="ml-4">
            <el-radio label="1" size="large">增加</el-radio>
            <el-radio label="3" size="large">减少</el-radio>
          </el-radio-group>
        </el-form-item>

        <el-form-item label="金额:">
          <el-input-number v-model="record.money"  style="width:100%" :precision="2" clearable />
        </el-form-item>
<!--        <el-form-item label="备注">
            <el-input
            v-model="record.remark"
            :rows="2"
            type="textarea"
            placeholder="备注内容"
          />
        </el-form-item> -->
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>

  </div>
</template>

<script>
export default {
  name: 'MinappUser'
}
</script>

<script setup>
import {
  createMinappUser,
  updateMinappUser,
  findMinappUser
} from '@/api/minappUser'
import CustomPic from '@/components/customPic/index.vue'
import {  formatDate } from '@/utils/format'
import { Search } from '@element-plus/icons-vue'
// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref } from 'vue'
const route = useRoute()
const router = useRouter()
const type = ref('')
const intOptions = ref([])
const formData = ref({
        openid: '',
        avatar: '',
        nickname: '',
        phone: '',
        area: '',
        number: '',
        balance: 0,
        status: undefined,

        })
const record = ref({
  uid:0,
  money:0,
  remark:"",
  type:1,
})
const change = ref("1")

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例

    if (route.params.id) {
      const res = await findMinappUser({ ID: route.params.id })
      console.log('test',res)
      if (res.code === 0) {
        formData.value = res.data.reminappUser
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    intOptions.value = await getDictFunc('int')
}

init()

const dialogFormVisible = ref(false)

const topUp = async() =>{
  dialogFormVisible.value = true
}

const toCoupon = (row) => {
  router.push({
    name: 'releaseRecord',
    params: {
      id: route.params.id,
    },
  })
}

const enterDialog = async()=>{
  let res
  record.value.uid=formData.value.ID
  record.value.type = parseInt(change.value)
  console.log("结果",record.value)
  res = await updateMinappUser(record.value)
    if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '修改成功'
        })

        if( parseInt(change.value) == 1){
          formData.value.balance += record.value.money
        }else{
          formData.value.balance -= record.value.money
        }
         closeDialog()
      }
}
// 保存按钮

const closeDialog = ()=>{
  dialogFormVisible.value = false
}
// 返回按钮
const back = () => {
    router.push({
		name:'minappUser'
	})
}

</script>

<style>
</style>
