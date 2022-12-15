<!--
 * @Author: your name
 * @Date: 2022-04-15 14:55:40
 * @LastEditTime: 2022-04-15 15:42:42
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /web/src/components/orderstate/check.vue
-->
<template>
           <el-select  filterable placeholder="请选择" v-model="selectState"   @change="handleChange">
            <el-option
              v-for="item in options"
              :key="item.value"
              :label="item.label"
              :value="item.value"
           
            />
          </el-select>
</template>

<script>
    export default{
        name:"CheckState"
    }
</script>

<script setup>
import { ref } from 'vue'
import{
    findSysDictionary
} from '@/api/sysDictionary'

const selectState = ref('')
const emit = defineEmits(['selectState'])
const handleChange = async(e) =>{
      emit('selectState', e)
}
const options = ref([])
const checkTableData = async() =>{
    const stateData = await findSysDictionary({ID:11})
    if (stateData.code == 0) {
      options.value = stateData.data.resysDictionary.sysDictionaryDetails
    }
} 
checkTableData()
</script>