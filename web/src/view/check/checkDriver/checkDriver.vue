<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="状态">
            <CheckState @select-state="selectState" />
        </el-form-item>

        <el-form-item label="姓名">
          <el-input v-model="searchInfo.name" placeholder="搜索姓名" />
        </el-form-item>

        <el-form-item label="申请日期">
          <el-date-picker
            v-model="searchInfo.time"
            type="date"
            value-format="YYYY-MM-DD"
            placeholder="选择日期"
          />
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
        >

        <el-table-column align="left" label="申请日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="姓名" prop="name" width="120" />
        <el-table-column align="left" label="联系方式" prop="phone" width="120" />
        <el-table-column align="left" label="驾照" prop="licence" width="120" >

                       <template #default="scope">
                  <el-image
                style="width: 100px;"
                :src="path+scope.row.licence"
                @click="ShowPic(scope.row.licence)"
              />
             </template>
        </el-table-column>
        <el-table-column align="left" label="营业执照" prop="charter" width="120" >
                       <template #default="scope">
                  <el-image
                style="width: 100px;"
                :src="path+scope.row.charter"
                @click="ShowPic(scope.row.charter)"
              />
             </template>

        </el-table-column>
        <el-table-column align="left" label="邮箱" prop="mail" width="120" />
        <el-table-column align="left" label="用户" prop="uid" width="120" >
          <template #default="scope">
                <el-button type="text" icon="search" size="small" class="table-button" @click="toUserDetail(scope.row)">{{scope.row.minAppUser.nickname}}</el-button>
          </template>

        </el-table-column>
        <el-table-column align="left" label="状态" prop="status" width="120" >
          <template #default="scope">
            <span style="color:blue">  {{scope.row.status == 1 ?"未审核":""}}</span>
            <span style="color:green">  {{scope.row.status == 2 ?"审核通过":""}}</span>
            <span style="color:red">   {{scope.row.status == 3 ?"已拒绝":""}}</span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="按钮组">
            <template #default="scope">
            <el-button type="text" icon="edit" size="small" class="table-button" @click="updateCheckDriverFunc(scope.row)">审核</el-button>
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
    <el-dialog v-model="dialogShowPic" :before-close="closeImg" title="认证图片">
      <el-image :src="showPicUrl"  />
    </el-dialog>

    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
      <el-form :model="formData" label-position="right" label-width="80px">

        <el-form-item label="状态:">
            <el-radio-group v-model="formData.status">
              <el-radio :label=1>未审核 </el-radio>
              <el-radio :label=2>审核通过 </el-radio>
              <el-radio :label=3>审核拒绝 </el-radio>
            </el-radio-group>
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
  name: 'CheckDriver'
}
</script>

<script setup>
import {
  createCheckDriver,
  deleteCheckDriver,
  deleteCheckDriverByIds,
  updateCheckDriver,
  findCheckDriver,
  getCheckDriverList
} from '@/api/checkDriver'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'
import CheckState from'@/components/checkState/index.vue'
const selectState = async(e) => {
  console.log("子组件传的值",e)
  searchInfo.value.status = e
}
import { useRouter } from 'vue-router'
const router = useRouter()
//获取用户详情
const toUserDetail = (row) => {
  router.push({
    name: 'userDetail',
    params: {
      id: row.uid,
    },
  })
}
const path = ref(import.meta.env.VITE_BASE_API + '/')
const ShowPic = (val) => {
 
  showPicUrl.value = path.value+val
  dialogShowPic.value = true
}
const closeImg = () => {
  dialogShowPic.value = false
}
const showPicUrl = ref("")
const dialogShowPic = ref(false)

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        name: '',
        phone: '',
        licence: '',
        charter: '',
        mail: '',
        time: '',
        uid: 0,
        status: 1,
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
  const table = await getCheckDriverList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
            deleteCheckDriverFunc(row)
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
      const res = await deleteCheckDriverByIds({ ids })
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
const updateCheckDriverFunc = async(row) => {
    const res = await findCheckDriver({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.recheckDriver
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteCheckDriverFunc = async (row) => {
    const res = await deleteCheckDriver({ ID: row.ID })
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
        name: '',
        phone: '',
        licence: '',
        charter: '',
        mail: '',
        time: '',
        uid: 0,
        status: 0,
        }
}
// 弹窗确定
const enterDialog = async () => {
      let res
      switch (type.value) {
        case 'create':
          res = await createCheckDriver(formData.value)
          break
        case 'update':
          res = await updateCheckDriver(formData.value)
          break
        default:
          res = await createCheckDriver(formData.value)
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
