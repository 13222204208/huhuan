<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="昵称">
          <el-input v-model="searchInfo.nickname" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="手机号">
          <el-input v-model="searchInfo.phone" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="地区">
          <el-input v-model="searchInfo.area_id" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="用户编号">
          <el-input v-model="searchInfo.user_num" placeholder="搜索条件" />
        </el-form-item>
        <!-- <el-form-item label="证件号">
          <el-input v-model="searchInfo.number" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="状态">
          <el-input v-model="searchInfo.status" placeholder="搜索条件" />
        </el-form-item> -->
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
        <!-- <el-table-column align="left" label="Openid" prop="openid" width="120" /> -->
        <el-table-column align="left" label="头像" >
            <template #default="scope">
                <img :src="scope.row.avatar" width="60" />
            </template>
          
        </el-table-column>

        <el-table-column align="left" label="昵称" prop="nickname" width="120" />
        <el-table-column align="left" label="手机号" prop="phone" width="120" />
        <el-table-column align="left" label="地区" prop="area_id" width="120" />
        <el-table-column align="left" label="证件号" prop="number" width="120" />
        <el-table-column align="left" label="余额" prop="balance" width="120" />
        <el-table-column align="left" label="状态" prop="status" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.status,intOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="按钮组">
            <template #default="scope">
            <!-- <el-button type="text" icon="edit" size="small" class="table-button" @click="updateMinappUserFunc(scope.row)">查看用户详情</el-button> -->
            <el-button type="text" icon="edit" size="small" class="table-button" @click="toUserDetail(scope.row)">查看用户详情</el-button>
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
        <el-form-item label="Openid:">
          <el-input v-model="formData.openid" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="头像:">
          <el-input v-model="formData.avatar" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="昵称:">
          <el-input v-model="formData.nickname" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="手机号:">
          <el-input v-model="formData.phone" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="地区:">
          <el-input v-model.number="formData.area_id" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="证件号:">
          <el-input v-model="formData.number" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="余额:">
          <el-input-number v-model="formData.balance"  style="width:100%" :precision="2" clearable />
        </el-form-item>
<!--        <el-form-item label="状态:">
          <el-select v-model="formData.status" placeholder="请选择" style="width:100%" clearable>
            <el-option v-for="(item,key) in intOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
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


  deleteMinappUserByIds,
  updateMinappUser,
  findMinappUser,
  getMinappUserList
} from '@/api/minappUser'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useRouter } from 'vue-router'
import { ref } from 'vue'

// 自动化生成的字典（可能为空）以及字段
const intOptions = ref([])
const formData = ref({
        openid: '',
        avatar: '',
        nickname: '',
        phone: '',
        area_id: 0,
        number: '',
        balance: 0,
        status: undefined,
        })

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
const router = useRouter()
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
  const table = await getMinappUserList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    intOptions.value = await getDictFunc('int')
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
            deleteMinappUserFunc(row)
        })
    }

const toUserDetail = (row) => {
  router.push({
    name: 'userDetail',
    params: {
      id: row.ID,
    },
  })
}


// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateMinappUserFunc = async(row) => {
    const res = await findMinappUser({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.reminappUser
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteMinappUserFunc = async (row) => {
    const res = await deleteMinappUser({ ID: row.ID })
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
        openid: '',
        avatar: '',
        nickname: '',
        phone: '',
        area_id: 0,
        number: '',
        balance: 0,
        status: undefined,
        }
}
// 弹窗确定
const enterDialog = async () => {
      let res
      switch (type.value) {
        case 'create':
          res = await createMinappUser(formData.value)
          break
        case 'update':
          res = await updateMinappUser(formData.value)
          break
        default:
          res = await createMinappUser(formData.value)
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
