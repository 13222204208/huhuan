<template>
    <div>
        <div class="gva-table-box">
            <el-form :model="form" label-width="120px">        
              <el-form-item label="消息内容">
                <el-input v-model="form.contents" type="textarea" />
              </el-form-item>
            <el-form-item label="选择用户">
                <el-radio-group v-model="form.resource">
                  <el-radio label="1" >指定用户发送</el-radio>
                  <el-radio label="2" >全部发送</el-radio>
            <!--      <el-radio label="3" >按地区发送</el-radio> -->
                </el-radio-group>
            </el-form-item>

            <div v-if="form.resource == 1">
              <el-form-item label="关联会员:">
                <el-button type="primary" :icon="Search" @click="selectUser">选择会员</el-button>
                <span style="color:blue; margin-left:5px"> 已选中 {{ids.length}} 人</span> 
              </el-form-item>
            </div>
              <el-form-item>
                <el-button type="primary" @click="onSubmit">确认发送</el-button>
            
              </el-form-item>
            </el-form>
        </div>

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
                <el-button  size="small" style="margin-left: 10px;"  :disabled="!multipleSelectionUser.length" @click="onSure">确定选择</el-button>
            </template>
            </el-popover>
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableDataUser"
        @selection-change="handleSelectionChangeUser"
        row-key="ID"
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

<script  setup>
import { reactive } from 'vue'
import { ref } from 'vue'
import {  Search } from '@element-plus/icons-vue'

import {
  getMinappUserList
} from '@/api/minappUser'

import {
  createSystemMessage
} from '@/api/systemMessage'
import { ElMessage, ElMessageBox } from 'element-plus'
const form = reactive({
  uid: [],
  contents: '',
  resource: '1'
})
const type = ref('')
const pageUser = ref(1)
const totalUser = ref(0)
const pageSizeUser = ref(10)
const tableDataUser = ref([])
const searchUserInfo = ref({})
const ids = ref([])
const selectUser = () => {
  dialogSelectUser.value = true
}
const userVisible = ref(false)
const dialogSelectUser = ref(false)
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
// 分页
const handleSizeChangeUser = (val) => {
  pageSizeUser.value = val
  getTableDataUser()
}
const onResetUser = () => {
  searchUserInfo.value = {}
}

const onSubmitUser = () => {
  pageUser.value = 1
  pageSizeUser.value = 10
  getTableDataUser()
}
const onSure = () =>{
   dialogSelectUser.value = false
  console.log("测试",ids)
}
// 修改页面容量
const handleCurrentChangeUser = (val) => {
  pageUser.value = val
  getTableDataUser()
}
const multipleSelectionUser = ref([])


// 多选
const handleSelectionChangeUser = (val) => {
    multipleSelectionUser.value = val
    console.log("测试",multipleSelectionUser.value)
      const selectId = []
    
      if (multipleSelectionUser.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选数据'
        })
        return
      }
      multipleSelectionUser.value &&
        multipleSelectionUser.value.map(item => {
          let uid = item.ID.toString()
          selectId.push(uid)
        })
        console.log('选中',selectId)
        ids.value = selectId
    console.log("选中的会员",ids)

}
// 查询
const getTableDataUser = async() => {
  const table = await getMinappUserList({ page: pageUser.value, pageSize: pageSizeUser.value, ...searchUserInfo.value })
  if (table.code === 0) {
    tableDataUser.value = table.data.list
    totalUser.value = table.data.total
    pageUser.value = table.data.page
    pageSizeUser.value = table.data.pageSize
  }
}

const onSubmit = async() => {

  if (form.contents == ''){
          ElMessage.error('请填写消息内容')
          return
  }

  if (form.resource == 1){
          form.uid = ids.value
          if (form.uid.length < 1){
                 ElMessage.error('请选择会员')
          return
          }
  }
  let res
  console.log(form)
  res = await createSystemMessage(form)
        if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '发送成功'
        })
      }
}
</script>
