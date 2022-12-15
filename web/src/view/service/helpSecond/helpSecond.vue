<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="下单日期">
          <el-date-picker v-model="searchInfo.creationDate" type="date" 
          value-format="YYYY-MM-DD"
          placeholder="选择日期" />
        </el-form-item>
        <el-form-item label="订单号">
          <el-input v-model="searchInfo.orderNum" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="商品名称">
          <el-input v-model="searchInfo.goodsName" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="我的地址">
          <el-input v-model="searchInfo.address" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="接单人">
          <el-input v-model="searchInfo.takeName" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="用户编号">
          <el-input v-model="searchInfo.userNum" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="订单状态">
            <el-select v-model="searchInfo.status" class="m-2" placeholder="选择状态">
            <el-option
              v-for="item in options"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="地区">
             <AreaList @select-area="selectArea" />
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="small" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      
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
        <el-table-column align="left" label="订单号" prop="orderNum" width="220" />
        <el-table-column align="left" label="商品详情" prop="goodsDetail" width="120" />
        <el-table-column align="left" label="商品图片" prop="goodsImg" width="120" />
        <el-table-column align="left" label="商品名称" prop="goodsName" width="120" />
        <el-table-column align="left" label="商品原价" prop="originaPrice" width="120" />
        <el-table-column align="left" label="商品售价" prop="price" width="120" />
        <el-table-column align="left" label="我的地址" prop="address" width="220" />
        <el-table-column align="left" label="接单人" prop="takeName" width="120" />
        <el-table-column align="left" label="接单人电话" prop="takePhone" width="120" />
        <el-table-column align="left" label="接单时间" prop="takeTime" width="120" />
        <el-table-column align="left" label="用户编号" prop="userNum" width="220" />
        <el-table-column align="left" label="订单状态" prop="status" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.status,orderStateOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="城市" prop="city" width="120" />
        <el-table-column align="left" label="地区" prop="area" width="120" />
        <el-table-column align="left" label="类别" prop="category" width="120" />
        <el-table-column align="left" label="按钮组">
            <template #default="scope">
            <el-button type="text" icon="edit" size="small" class="table-button" @click="updateHelpSecondFunc(scope.row)">变更</el-button>
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
          <el-input-number v-model="formData.originaPrice"  style="width:100%" :precision="2" clearable />
        </el-form-item>
        <el-form-item label="商品售价:">
          <el-input-number v-model="formData.price"  style="width:100%" :precision="2" clearable />
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
          <el-select v-model="formData.status" placeholder="请选择" style="width:100%" clearable>
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
  name: 'HelpSecond'
}
</script>

<script setup>
import {
  createHelpSecond,
  deleteHelpSecond,
  deleteHelpSecondByIds,
  updateHelpSecond,
  findHelpSecond,
  getHelpSecondList
} from '@/api/helpSecond'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'
import AreaList from '@/components/areaList/index.vue'
const selectArea = async(e) => {
  console.log("子组件传的值",e)
  searchInfo.value.area = e
}
const options = ref([
  {
    value:1,
    label:"上架"
  },
  {
    value:2,
    label:"下架"
  }
])
// 自动化生成的字典（可能为空）以及字段
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

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

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
  const table = await getHelpSecondList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    orderStateOptions.value = await getDictFunc('orderState')
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
            deleteHelpSecondFunc(row)
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
      const res = await deleteHelpSecondByIds({ ids })
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
const updateHelpSecondFunc = async(row) => {
    const res = await findHelpSecond({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.rehelpSecond
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteHelpSecondFunc = async (row) => {
    const res = await deleteHelpSecond({ ID: row.ID })
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
        }
}
// 弹窗确定
const enterDialog = async () => {
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
        closeDialog()
        getTableData()
      }
}
</script>

<style>
</style>
