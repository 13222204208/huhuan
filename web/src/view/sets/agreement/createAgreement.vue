<template>
	<div>
	  <div class="gva-search-box">
		<el-form :model="form" label-width="80px">
		  <el-form-item label="标题">
			<el-input v-model="form.title" clearable placeholder="请输入"></el-input>
		  </el-form-item>
		  
		  <el-form-item label="内容">
			  <vue3-tinymce v-model="state.content" :setting="state.setting" />
		  </el-form-item>
		
		 <el-form-item>
			<el-button type="primary" @click="onSubmit">立即保存</el-button>
			<el-button>取消</el-button>
		  </el-form-item>
		</el-form>
	   </div>
	</div>
</template>

<script>
	
export default {
  name: 'Agreement'
}

</script>

<script setup>
	
	import {
		findAgreement,
	  updateAgreement,
	  createAgreement,
	} from '@/api/agreement'
	
import Vue3Tinymce from '@jsdawn/vue3-tinymce'
import { useUserStore } from '@/pinia/modules/user'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()

const form = ref({
			title:'',
			contents:''
		})

const userStore = useUserStore()
const type = ref('')

type.value = 'create'

//获取协议详情
const getAgreementData = async (agreementId) => {
	const res = await findAgreement({ ID: agreementId })
		console.log('详情',res)
		if (res.code === 0) {
			// form.value.title = res.data.reagreement.title
			// form.value.ID = res.data.reagreement.ID
			form.value = res.data.reagreement
			state.value.content = res.data.reagreement.contents	
		}	
}

const agreementId = router.currentRoute.value.query.id
console.log('id',agreementId)
if(agreementId){
	type.value = 'update'
	getAgreementData(agreementId)
}

//结束

console.log('表单信息',form.value)

const onSubmit = async () => {
	  form.value.contents = state.value.content
	  console.log(form.value)
      let res
      switch (type.value) {
        case 'create':
          res = await createAgreement(form.value)
          break
        case 'update':
          res = await updateAgreement(form.value)
          break
        default:
          res = await createAgreement(form.value)
          break
      }
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '创建/更改成功'
        })
        router.push({name:'agreement'})
      }
}

const state = ref({
  content: '',
  // editor 配置项
  setting: {
		 height: 560,
		  toolbar:
			'undo redo | fullscreen | formatselect alignleft aligncenter alignright alignjustify | link unlink | numlist bullist | image media table | fontsizeselect forecolor backcolor | bold italic underline strikethrough | indent outdent | superscript subscript | removeformat |',
		  toolbar_drawer: 'sliding',
		  quickbars_selection_toolbar:
			'removeformat | bold italic underline strikethrough | fontsizeselect forecolor backcolor',
		  plugins: 'link image media table lists fullscreen quickbars',
		  fontsize_formats: '12px 14px 16px 18px',
		  default_link_target: '_blank',
		  link_title: false,
		  nonbreaking_force_tab: true,
		  
		  
		    // 开启组件拓展的 上传图片功能，工具栏 图片按钮 弹框中出现 `upload` 选项
		    custom_images_upload: true,
		    // 复用 图片上传 api 地址
		    images_upload_url: import.meta.env.VITE_BASE_API + '/fileUploadAndDownload/upload',
		    // 上传成功回调函数，return 图片地址。默认 response.location
		   // custom_images_upload_callback: response => response.url,
		   custom_images_upload_callback: response =>import.meta.env.VITE_BASE_API +'/'+ response.data.file.url,
		
		    // 上传 api 请求头
		    custom_images_upload_header: { 'X-Token': userStore.token  },
		    // 上传 api 额外的参数。图片默认 file
		   // custom_images_upload_param: { id: 'xxxx01', age: 18 },
			
		  // 以中文简体为例
		  language: 'zh_CN',
		  language_url:
			'https://unpkg.com/@jsdawn/vue3-tinymce@1.1.6/dist/tinymce/langs/zh_CN.js'
			
  },
});




	
</script>