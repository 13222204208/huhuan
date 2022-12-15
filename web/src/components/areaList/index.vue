<!--
 * @Author: your name
 * @Date: 2022-04-12 18:48:27
 * @LastEditTime: 2022-04-13 10:21:57
 * @LastEditors: your name
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /web/src/components/areaList/index.vue
-->
<template>
            <div class="example-block">
              <el-cascader
                placeholder="地区名称"
                :options="options"
                 @change="handleChange"
                change-on-select
                filterable
              />
            </div>
</template>

<script>
export default {
  name: 'AreaList'
}
</script>

<script setup>
import {
  getAreaList,
} from '@/api/area'
import { ref } from 'vue'

const options = ref([])
const searchInfo = ref({})



const handleChange = async(e) =>{
      var end = e.pop()
      console.log("数据",end)
      emit('selectArea', end)
}
const emit = defineEmits(['selectArea'])
// 查询
const getTableAreaData = async() => {
  const table = await getAreaList({ page: 1, pageSize: 10000, ...searchInfo.value })
  
  let listData = table.data.list
  for (let index = 0; index < listData.length; index++) {
    listData[index].value = listData[index].title
    listData[index].label = listData[index].title
  }

  if (table.code === 0) {
    options.value = arraytotree(listData)

  }
}

getTableAreaData()

const arraytotree = (arr)=> {
                var top = [], sub = [], tempObj = {};
                arr.forEach(function (item) {
                    if (item.pid == 0) { // 顶级分类
                        top.push(item)
                    } else {
                        sub.push(item) // 其他分类
                    }
                    item.children = []; // 默然添加children属性
                    tempObj[item.ID] = item // 用当前分类的id做key，存储在tempObj中
                })
                sub.forEach(function (item) {
                    // 取父级
                    var parent = tempObj[item.pid] || {'children': []}
                    // 把当前分类加入到父级的children中
                    parent.children.push(item)
                })
                return top
            }
</script>