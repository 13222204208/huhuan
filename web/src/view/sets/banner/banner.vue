<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="城市">
          <el-input v-model="searchInfo.cityId" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="分类">
          <el-input v-model="searchInfo.classId" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="small" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button size="small" type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-popover v-model:visible="deleteVisible" placement="top" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin-top: 8px;">
                <el-button size="small" type="text" @click="deleteVisible = false">取消</el-button>
                <el-button size="small" type="primary" @click="onDelete">确定</el-button>
            </div>
            <template #reference>
                <el-button icon="delete" size="small" style="margin-left: 10px;" :disabled="!multipleSelection.length">删除</el-button>
            </template>
            </el-popover>
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <!-- <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column> -->
        <el-table-column align="center" label="标题" prop="title" width="120" />
        <el-table-column align="center" label="城市" prop="city" width="120" />
        <el-table-column align="center" label="分类" prop="class" width="120" >
          <template #default="scope">
               {{scope.row.class == 1 ? "首页":""}}
               {{scope.row.class == 2 ? "国内捎带":""}}
               {{scope.row.class == 3 ? "组团拼车":""}}
               {{scope.row.class == 4 ? "团购接龙":""}}
               {{scope.row.class == 5 ? "二手闲置":""}}
          </template>
        </el-table-column>
        <el-table-column align="center" label="图片" width="220">
          <template #default="scope">
            <CustomPic pic-type="file" :pic-src="scope.row.img" />
          </template>
        </el-table-column>
        <el-table-column align="center" label="链接" prop="url" width="120" />
        <el-table-column align="center" label="按钮组">
            <template #default="scope">
            <el-button type="text" icon="edit" size="small" class="table-button" @click="updateBannerFunc(scope.row)">变更</el-button>
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
      <ChooseImg ref="chooseImgRef" @enter-img="enterImg" />
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="标题:">
          <el-input v-model="formData.title" clearable placeholder="请输入" />
        </el-form-item>
        
        <el-form-item label="分类:">
          <el-select v-model="formData.class"  placeholder="请选择分类" style="width: 58rem;">
            <el-option v-for='item in helpMeTitleData' :key='item.ID' :value='item.ID' :label='item.title'>  			</el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="城市:" v-if="formData.class == 1">
          <!-- <el-input v-model.number="formData.cityId" clearable placeholder="请输入" /> -->
          <AreaList @select-area="selectArea" />
        </el-form-item>

        <el-form-item label="图片:">
<!--            <div
              class="user-headpic-update"
              :style="{
                'background-image': `url(${
                  path
                })`,
                'background-repeat': 'no-repeat',
                'background-size': 'cover',
              }"
            >
              <span class="update" @click="openChooseImg">
                <el-icon>
                  <edit />
                </el-icon>
                重新上传</span>
            </div> -->
			<div
			  class="user-headpic-update"
			  :style="{
						'background-image': `url(${
						  path
						})`,
						'background-repeat': 'no-repeat',
						'background-size': 'cover',
						'width':'200px',
					  }"
			>
			</div>
		
			<el-upload
			  :action="`${urlPath}/fileUploadAndDownload/upload`"
			  :before-upload="checkFile"
			  :headers="{ 'x-token': userStore.token }"
			  :on-error="uploadError"
			  :on-success="uploadSuccess"
			  :show-file-list="false"
			  class="upload-btn"
			>
			  <el-button size="small" type="primary">重新上传</el-button>
			</el-upload>
			
        </el-form-item>
        <el-form-item label="链接:">
          <el-input v-model="formData.url" clearable placeholder="请输入" />
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
  name: 'Banner'
}
</script>

<script setup>
import {
  createBanner,
  deleteBanner,
  deleteBannerByIds,
  updateBanner,
  findBanner,
  getBannerList
} from '@/api/banner'


import ChooseImg from '@/components/chooseImg/index.vue'
import CustomPic from '@/components/customPic/index.vue'
import AreaList from '@/components/areaList/index.vue'
// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'
import { useUserStore } from '@/pinia/modules/user'
const urlPath = ref(import.meta.env.VITE_BASE_API)
const userStore = useUserStore()

const path = ref("https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif?imageView2/1/w/80/h/80")
// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        title: '',
        city: '',
        class: '',
        img: '',
        url: '',
        })

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
const helpMeTitleData = ref([
  {'ID': 1,'title':'首页'},
  // {'ID': 2,'title':'国内捎带'},
  // {'ID': 3,'title':'组团拼车'},
  // {'ID': 4,'title':'团购接龙'},
  // {'ID': 5,'title':'二手闲置'},
])
const selectArea = async(e) => {
  console.log("子组件传的值",e)
  formData.value.city = e
}
//图片上传

const enterImg = async(url) => {
  console.log('图片地址',url)
  path.value = import.meta.env.VITE_BASE_API+'/'+url
  formData.value.img = url
}

const fullscreenLoading = ref(false)
const checkFile = (file) => {
  fullscreenLoading.value = true
  const isJPG = file.type === 'image/jpeg'
  const isPng = file.type === 'image/png'
  const isLt2M = file.size / 1024 / 1024 < 0.5
  if (!isJPG && !isPng) {
    ElMessage.error('上传图片只能是 jpg或png 格式!')
    fullscreenLoading.value = false
  }
  if (!isLt2M) {
    ElMessage.error('未压缩未见上传图片大小不能超过 500KB，请使用压缩上传')
    fullscreenLoading.value = false
  }
  return (isPng || isJPG) && isLt2M
}

const uploadSuccess = (res) => {
  fullscreenLoading.value = false
  console.log("图片",res)
  if (res.code === 0) {
	  var url = res.data.file.url
	  path.value = import.meta.env.VITE_BASE_API+'/'+url
	  console.log("图片地址",path.value)
	  formData.value.img = url
	  
    ElMessage({
      type: 'success',
      message: '上传成功'
    })
    if (res.code === 0) {
      getTableData()
    }
  } else {
    ElMessage({
      type: 'warning',
      message: res.msg
    })
  }
}
const uploadError = () => {
  ElMessage({
    type: 'error',
    message: '上传失败'
  })
  fullscreenLoading.value = false
}

const chooseImgRef = ref(null)

const openChooseImg = () => {
  chooseImgRef.value.open()
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

// 查询
const getTableData = async() => {
  const table = await getBannerList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

//获取用户分类名称
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
            deleteBannerFunc(row)
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
      const res = await deleteBannerByIds({ ids })
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
const updateBannerFunc = async(row) => {
    const res = await findBanner({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.rebanner
        path.value = import.meta.env.VITE_BASE_API+'/'+res.data.rebanner.img

        dialogFormVisible.value = true
    }
}


// 删除行
const deleteBannerFunc = async (row) => {
    const res = await deleteBanner({ ID: row.ID })
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
        city: '',
        class: '',
        img: '',
        url: '',
        }
}
// 弹窗确定
const enterDialog = async () => {
      let res
      switch (type.value) {
        case 'create':
          res = await createBanner(formData.value)
          break
        case 'update':
          res = await updateBanner(formData.value)
          break
        default:
          res = await createBanner(formData.value)
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

<style lang="scss">
    .user-headpic-update {
      //width: 120px;
      height: 120px;
      line-height: 120px;
     // margin: 0 auto;
      display: flex;
      justify-content: center;
      border-radius: 20px;
      &:hover {
        color: #fff;
        background: linear-gradient(
            to bottom,
            rgba(255, 255, 255, 0.15) 0%,
            rgba(0, 0, 0, 0.15) 100%
          ),
          radial-gradient(
              at top center,
              rgba(255, 255, 255, 0.4) 0%,
              rgba(0, 0, 0, 0.4) 120%
            )
            #989898;
        background-blend-mode: multiply, multiply;
        .update {
          color: #fff;
        }
      }
      .update {
        height: 120px;
        width: 220px;
        text-align: center;
        color: transparent;
      }
    }
</style>
