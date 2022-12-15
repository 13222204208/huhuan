<template>
  <div>
    <!-- <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="small" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div> -->
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
        <el-table-column align="left" label="名称" prop="title" width="120" />
        <el-table-column align="left" label="使用条件" prop="condition" width="120" >
          <template #default="scope">
            {{scope.row.condition < 1 ?"不限制": "满"+scope.row.condition+"元"}}
          </template>
        </el-table-column>

        <el-table-column align="left" label="优惠方式" prop="way" width="120" >
          <template #default="scope"> 
             立减 {{scope.row.way}} 元
          </template>
        </el-table-column>

        <el-table-column align="left" label="开始时间" prop="start" width="120" />
        <el-table-column align="left" label="结束时间" prop="over" width="120" />
       
        <el-table-column align="left" label="发放总数" prop="send" width="120" />
     
        <el-table-column align="left" label="积分兑换" prop="integral" width="120" />
        <el-table-column align="left" label="按钮组">
            <template #default="scope">
            <el-button type="text" icon="edit" size="small" class="table-button" @click="updateCouponFunc(scope.row)">变更</el-button>
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
      <el-form :model="formData" label-position="right" :rules="rules" label-width="80px">
        <el-form-item label="名称:" prop="title">
          <el-input v-model="formData.title" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="使用条件:">
          <el-input v-model.number="formData.condition" clearable placeholder="请输入" />
          <span>消费满多少可用, 空或0 不限制</span>
        </el-form-item>
  
        <el-form-item label="有效日期:" prop="dateRange">
              <el-date-picker
                v-model="dateRange"
                type="daterange"
                range-separator="至"
                format="YYYY/MM/DD"
                value-format="YYYY-MM-DD"
                start-placeholder="开始日期"
                end-placeholder="结束日期">
              </el-date-picker>
        </el-form-item>
        <el-form-item label="优惠方式:">
          
          <el-input v-model.number="formData.way" clearable placeholder="请输入" >
            <template #prepend>立减</template>
            <template #append>元</template>
            </el-input>
        </el-form-item>
        <el-form-item label="发放总数:">
          <el-input v-model.number="formData.send" clearable placeholder="请输入" />
          <span>优惠券总数量，没有不能领取或发放,-1 为不限制张数</span>
        </el-form-item>
        
        <el-form-item label="积分兑换:">
          <el-input v-model.number="formData.integral" clearable placeholder="请输入" />
          <span>使用积分兑换优惠券, 0, 为不能兑换</span>
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
  name: 'Coupon'
}
</script>

<script setup>
import {
  createCoupon,
  deleteCoupon,
  deleteCouponByIds,
  updateCoupon,
  findCoupon,
  getCouponList
} from '@/api/coupon'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import {reactive, ref } from 'vue'

const rules = reactive({
  title: [{ required: true, message: '请输入优惠券名称', trigger: 'blur' }],

})

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        title: '',
        condition: 0,
        start: '',
        over: '',
        send: -1,
        way: 1,
        integral: 0,
        })

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
const dateRange = ref([])

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
  const table = await getCouponList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

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
            deleteCouponFunc(row)
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
      const res = await deleteCouponByIds({ ids })
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
const updateCouponFunc = async(row) => {
    const res = await findCoupon({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.recoupon
        dateRange.value[0] = res.data.recoupon.start
        dateRange.value[1] = res.data.recoupon.over
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteCouponFunc = async (row) => {
    const res = await deleteCoupon({ ID: row.ID })
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
        title: '',
        condition: 0,
        start: '',
        over: '',
        send: -1,
        way: 1,
        integral: 0,
        }
}
// 弹窗确定
const enterDialog = async () => {

      if(formData.value.title == ''){
        ElMessage.error('请填写优惠券名称')
        return
      }

      if(dateRange.value.length == 0){
        ElMessage.error('请选择有效日期范围')
        return
      }
      formData.value.over = dateRange.value[1]
      formData.value.start = dateRange.value[0]
  
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
        closeDialog()
        getTableData()
      }
}
</script>

<style>
</style>
