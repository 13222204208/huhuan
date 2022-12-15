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
            <!-- <el-popover v-model:visible="deleteVisible" placement="top" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin-top: 8px;">
                <el-button size="small" type="text" @click="deleteVisible = false">取消</el-button>
                <el-button size="small" type="primary" @click="onDelete">确定</el-button>
            </div>
            <template #reference>
                <el-button icon="delete" size="small" style="margin-left: 10px;" :disabled="!multipleSelection.length">删除</el-button>
            </template>
            </el-popover> -->
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column align="left" label="有效期" width="230">
            <template #default="scope">{{ formatDate(scope.row.start,"day")}} 到 {{formatDate(scope.row.over,"day")}}</template>
        </el-table-column>
        <el-table-column align="left" label="用户" prop="uid" width="120" >
            <template #default="scope">
              {{scope.row.UserInfo.nickname}}
            </template>
        </el-table-column>
        <el-table-column align="left" label="优惠券" prop="title" width="120" />
        <el-table-column align="left" label="使用条件" width="180">
          <template #default="scope">
            满{{scope.row.condition}}可用
          </template>
        </el-table-column>
        <el-table-column align="left" label="数量" prop="total" width="120" />
        <el-table-column align="left" label="状态" prop="status" width="120" >
          <template #default="scope">
          <span style="color:green">  {{scope.row.status == 0 ? "未使用":""}}</span>
          <span style="color:red" > {{scope.row.status == 1 ? "已使用":""}}</span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="按钮组">
            <template #default="scope">
            <!-- <el-button type="text" icon="edit" size="small" class="table-button" @click="updateReleaseRecordFunc(scope.row)">变更</el-button> -->
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
      <el-form :model="formData" label-position="right" label-width="110px">

        <el-form-item label="选择优惠券:">
          <el-select v-model="formData.couponId"  placeholder="请选择优惠券" style="width: 58rem;">
            <el-option v-for='item in couponTitle' :key='item.ID' :value='item.ID' :label='item.title'>  			</el-option>
          </el-select>
        </el-form-item>

        <el-form-item label="每人发送张数:">
          <el-input v-model.number="formData.total" clearable placeholder="请输入" />
          此处受总数限制，如果剩余张数不足以发放给选定会员数量，则无法发放
        </el-form-item>
        <el-form-item label="关联会员:">
          <el-button type="primary" :icon="Search" @click="selectUser">选择会员</el-button>
           <span style="color:blue; margin-left:5px"> 已选中 {{ids.length}} 人</span> 
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>

       <el-dialog v-model="dialogSelectUser" :before-close="closeSelectUser" title="选择用户">
             <div class="gva-search-box">
      <el-form :inline="true" :model="searchUserInfo" class="demo-form-inline">
        <el-form-item label="昵称">
          <el-input v-model="searchUserInfo.nickname" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="手机号">
          <el-input v-model="searchUserInfo.phone" placeholder="搜索条件" />
        </el-form-item>

        <el-form-item label="编号">
          <el-input v-model="searchUserInfo.uid" placeholder="搜索条件" />
        </el-form-item>

        <el-form-item label="地区">
          <el-input v-model="searchUserInfo.area" placeholder="搜索条件" />
        </el-form-item>

        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmitUser">查询</el-button>
          <el-button size="small" icon="refresh" @click="onResetUser">重置</el-button>
        </el-form-item>
      </el-form>
    </div>

        <div class="gva-btn-list">
            <!-- <el-button size="small" type="primary" icon="plus" @click="openDialog">新增</el-button> -->
            <el-popover v-model:visible="userVisible" placement="top" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin-top: 8px;">
                <el-button size="small" type="text" @click="userVisible = false">取消</el-button>
                <!-- <el-button size="small" type="primary" @click="onDelete">确定</el-button> -->
            </div>
            <template #reference>
                <el-button  size="small" style="margin-left: 10px;" :disabled="!multipleSelectionUser.length" @click="onSure">确定选择</el-button>
            </template>
            </el-popover>
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableDataUser"
        row-key="ID"
        @selection-change="handleSelectionChangeUser"
        >
        <el-table-column type="selection" width="55" />
        <!-- <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column> -->

        <el-table-column align="left" label="头像" >
            <template #default="scope">
                <img :src="scope.row.avatar" width="60" />
            </template>
          
        </el-table-column>

        <el-table-column align="left" label="昵称" prop="nickname" width="120" />
        <el-table-column align="left" label="手机号" prop="phone" width="120" />
        <el-table-column align="left" label="地区" prop="area_id" width="120" />
        <el-table-column align="left" label="证件号" prop="number" width="120" />

        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="pageUser"
            :page-size="pageSizeUser"
            :page-sizes="[10, 30, 50, 100]"
            :total="totalUser"
            @current-change="handleCurrentChangeUser"
            @size-change="handleSizeChangeUser"
            />
        </div>

    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'ReleaseRecord'
}
</script>

<script setup>
import {
  createReleaseRecord,
  deleteReleaseRecord,
 // deleteReleaseRecordByIds,
  updateReleaseRecord,
  findReleaseRecord,
  getReleaseRecordList
} from '@/api/releaseRecord'

import {
  getMinappUserList
} from '@/api/minappUser'

import {
    getAllCouponTitle
} from '@/api/coupon'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'
import {  Search } from '@element-plus/icons-vue'
import { useRoute, useRouter } from "vue-router"
// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        uid: 0,
        couponId: null,
        total: 1,
        status: 0,
        })

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
const couponTitle = ref([])//获取所有的优惠券

const pageUser = ref(1)
const totalUser = ref(0)
const pageSizeUser = ref(10)
const tableDataUser = ref([])
const searchUserInfo = ref({})

const dialogSelectUser = ref(false)
const route = useRoute()

const ids = ref([])

const init = async() => {
  if (route.params.id){
    searchInfo.value.uid = route.params.id
  }
}
init()

const selectUser = () => {
  dialogSelectUser.value = true
}

const onSubmitUser = () => {
  pageUser.value = 1
  pageSizeUser.value = 10
  getTableDataUser()
}

// 分页
const handleSizeChangeUser = (val) => {
  pageSizeUser.value = val
  getTableDataUser()
}

// 修改页面容量
const handleCurrentChangeUser = (val) => {
  pageUser.value = val
  getTableDataUser()
}

// 查询
const getTableDataUser = async() => {
  const table = await getMinappUserList({ page: page.value, pageSize: pageSize.value, ...searchUserInfo.value })
  if (table.code === 0) {
    tableDataUser.value = table.data.list
    totalUser.value = table.data.total
    pageUser.value = table.data.page
    pageSizeUser.value = table.data.pageSize
  }
}

getTableDataUser()

const onResetUser = () => {
  searchUserInfo.value = {}
}

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


//获取所有的优惠券
const getAllCoupons = async() => {
  const list = await getAllCouponTitle()
  if(list.code === 0){
    couponTitle.value = list.data.allCoupons
  }
}

getAllCoupons()

// 查询
const getTableData = async() => {
  const table = await getReleaseRecordList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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

// 多选数据
const multipleSelectionUser = ref([])


// 多选
const handleSelectionChangeUser = (val) => {
    multipleSelectionUser.value = val
    console.log("测试",multipleSelectionUser.value)
      const selectId = []
    
      if (multipleSelectionUser.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelectionUser.value &&
        multipleSelectionUser.value.map(item => {
          selectId.push(item.ID)
        })
        console.log('选中',selectId)
        ids.value = selectId
    console.log("选中的会员",ids)

}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteReleaseRecordFunc(row)
        })
    }


// 批量删除控制标记
// const deleteVisible = ref(false)
 const userVisible = ref(false)

// // 多选删除
// const onDelete = async() => {
//       const ids = []
//       if (multipleSelection.value.length === 0) {
//         ElMessage({
//           type: 'warning',
//           message: '请选择要删除的数据'
//         })
//         return
//       }
//       multipleSelection.value &&
//         multipleSelection.value.map(item => {
//           ids.push(item.ID)
//         })
//       const res = await deleteReleaseRecordByIds({ ids })
//       if (res.code === 0) {
//         ElMessage({
//           type: 'success',
//           message: '删除成功'
//         })
//         if (tableData.value.length === ids.length && page.value > 1) {
//           page.value--
//         }
//         deleteVisible.value = false
//         getTableData()
//       }
//     }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateReleaseRecordFunc = async(row) => {
    const res = await findReleaseRecord({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.rereleaseRecord
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteReleaseRecordFunc = async (row) => {
    const res = await deleteReleaseRecord({ ID: row.ID })
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
const closeSelectUser = () => {
    dialogSelectUser.value = false
   
}

const onSure = () =>{
   dialogSelectUser.value = false
  console.log("测试",ids)
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        uid: 0,
        couponId: null,
        total: 1,
        status: 0,
        }
}
// 弹窗确定
const enterDialog = async () => {


      if(formData.value.couponId < 1){
          ElMessage.error('请选择优惠券')
          return
      }

      if(ids.value.length < 1){
          ElMessage.error('请选择会员')
          return
      }

      formData.value.uid = ids.value
      let res
      switch (type.value) {
        case 'create':
          res = await createReleaseRecord(formData.value)
          break
        case 'update':
          res = await updateReleaseRecord(formData.value)
          break
        default:
          res = await createReleaseRecord(formData.value)
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
