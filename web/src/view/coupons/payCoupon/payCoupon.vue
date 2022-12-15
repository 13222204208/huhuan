<template>
  <div>

    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button size="small" type="primary" icon="plus" @click="openDialog">新增</el-button>
  
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="类型" width="200">
            <template #default="scope">
                {{scope.row.type == 1 ?"充值赠送优惠券":""}}
                {{scope.row.type == 2 ?"邀请用户赠送优惠券":""}}

            </template>
        </el-table-column>
        <el-table-column align="left" label="满" prop="money" width="120" />
        <el-table-column align="left" label="赠送优惠券" prop="title" width="120" />
        <el-table-column align="left" label="数量" prop="total" width="120" />
        <el-table-column align="left" label="按钮组">
            <template #default="scope">
            <el-button type="text" icon="edit" size="small" class="table-button" @click="updatePayCouponFunc(scope.row)">变更</el-button>
            <el-button type="text" icon="delete" size="small" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
      <el-form :model="formData" label-position="right" label-width="120px">

        <el-form-item label="类型">
          <el-select v-model="formData.type" style="width: 58rem;">
                <el-option
                  v-for="item in options"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value"
                />          
          </el-select>
        </el-form-item>
        <el-form-item label="满:">
          <el-input type="number" v-model.number="formData.money" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="选择优惠券:">
          <el-select v-model="formData.couponId"  placeholder="请选择优惠券" style="width: 58rem;">
            <el-option v-for='item in couponTitle' :key='item.ID' :value='item.ID' :label='item.title'>  			</el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="赠送数量:">
          <el-input type="number" v-model.number="formData.total" clearable placeholder="请输入" />
        </el-form-item>
 
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
  name: 'PayCoupon'
}
</script>

<script setup>
import {
  createPayCoupon,
  deletePayCoupon,
  deletePayCouponByIds,
  updatePayCoupon,
  findPayCoupon,
  getPayCouponList
} from '@/api/payCoupon'

import {
    getAllCouponTitle
} from '@/api/coupon'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'

const options = [
  {
    value: 1,
    label: '充值赠送优惠券',
  },
  {
    value: 2,
    label: '邀请好友赠送优惠券',
  },
]

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        money: 0,
        couponId: null,
        type:null,
        total: 1,
        })

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
const couponTitle = ref([])//获取所有的优惠券
// 重置
const onReset = () => {
  searchInfo.value = {}
}

// 搜索
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getPayCouponList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

//获取所有的优惠券
const getAllCoupons = async() => {
  const list = await getAllCouponTitle()
  if(list.code === 0){
    couponTitle.value = list.data.allCoupons
  }
}

getAllCoupons()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deletePayCouponFunc(row)
        })
    }


// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async() => {
      const ids = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.ID)
        })
      const res = await deletePayCouponByIds({ ids })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === ids.length && page.value > 1) {
          page.value--
        }
        deleteVisible.value = false
        getTableData()
      }
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updatePayCouponFunc = async(row) => {
    const res = await findPayCoupon({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.repayCoupon
        dialogFormVisible.value = true
    }
}


// 删除行
const deletePayCouponFunc = async (row) => {
    const res = await deletePayCoupon({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        money: 0,
        couponId: null,
        type:null,
        total: 1,
        }
}
// 弹窗确定
const enterDialog = async () => {
      let res
      switch (type.value) {
        case 'create':
          console.log(formData.value)
     
          res = await createPayCoupon(formData.value)
          break
        case 'update':
          res = await updatePayCoupon(formData.value)
          break
        default:
          res = await createPayCoupon(formData.value)
          break
      }
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '创建/更改成功'
        })
        closeDialog()
        getTableData()
      }
}
</script>

<style>
</style>
