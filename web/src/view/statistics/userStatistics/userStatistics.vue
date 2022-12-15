<!--
 * @Author: your name
 * @Date: 2022-04-22 13:46:32
 * @LastEditTime: 2022-05-17 09:18:18
 * @LastEditors: yangpanda 573516293@qq.com
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /web/src/view/statistics/userStatistics/userStatistics.vue
-->
<template>
<div>
    <div class="gva-table-box">
    <el-tabs v-model="activeName" class="demo-tabs" @tab-click="handleClick">
        <div class="gva-search-box">
                      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
                    <!-- <el-form-item label="城市">
                    <el-input v-model="searchInfo.cityId" placeholder="搜索条件" />
                    </el-form-item> -->
                    <el-form-item label="日期范围">
          
                          <el-date-picker
                        v-model="rangeDate"
                        type="daterange"
                        unlink-panels
                        range-separator="到"
                        start-placeholder="开始日期"
                        end-placeholder="结束日期"
                        value-format="YYYY-MM-DD"
                        :shortcuts="shortcuts"
                      />
                    </el-form-item>
                    <el-form-item>
                    <el-button size="small" type="primary" icon="search" @click="onSubmit">查询</el-button>
                
                    </el-form-item>
                </el-form>
          </div>        
        <el-tab-pane label="新增用户" name="addUserInfo">
           

          <div class="gva-table-box">
              <el-row :gutter="12">
                <el-col :span="8">
                  <el-card shadow="always"> 
                    今日活跃用户：{{ tableData.active}}  
                    
                  </el-card>
                </el-col>
                <el-col :span="8">
                  <el-card shadow="hover">
                  用户总数：{{tableData.all}}  
                  </el-card>
                </el-col>
              </el-row>

                <el-table :data="tableData.list" style="width: 100%">
                  <el-table-column prop="date" label="日期" width="280" />
                  <el-table-column prop="total" label="用户数" width="180" />
                </el-table>
          </div>
          
        </el-tab-pane>
        <el-tab-pane label="用户排行" name="userTop">
          <div class="gva-table-box" style="display:flex;">
                 <el-table :data="tableData.buyList" style="width: 100%">
                  <el-table-column prop="uid" label="送送用户" width="100"  >
                              <template #default="scope">
                        <el-button type="text" @click="toUserDetail(scope.row.uid)">   查看用户</el-button> 
                        </template>
                  </el-table-column>
                  <el-table-column prop="total" label="订单数量" width="80" >
                    <template #default="scope">
             <el-button type="text" @click="toHelpOrder(scope.row.uid,'helpBuy')">  {{scope.row.total}}</el-button> 
                    </template>
                  </el-table-column>
                  <el-table-column prop="money" label="交易金额" width="80" />
                </el-table>

                <el-table :data="tableData.workList" style="width: 100%">
                  <el-table-column prop="uid" label="帮帮用户" width="100"  >
                              <template #default="scope">
                        <el-button type="text" @click="toUserDetail(scope.row.uid)">   查看用户</el-button> 
                        </template>
                  </el-table-column>

                  <el-table-column prop="total" label="订单数量" width="80" >
                    <template #default="scope">
             <el-button type="text" @click="toHelpOrder(scope.row.uid,'helpWork')">  {{scope.row.total}}</el-button> 
                    </template>
                  </el-table-column>
                  
                  <el-table-column prop="money" label="交易金额" width="80" />
                </el-table>

                <el-table :data="tableData.repairList" style="width: 100%">
                  <el-table-column prop="uid" label="急急急用户" width="100"  >
                              <template #default="scope">
                        <el-button type="text" @click="toUserDetail(scope.row.uid)">   查看用户</el-button> 
                        </template>
                  </el-table-column>
                
               <el-table-column prop="total" label="订单数量" width="80" >
                    <template #default="scope">
             <el-button type="text" @click="toHelpOrder(scope.row.uid,'helpRepair')">  {{scope.row.total}}</el-button> 
                    </template>
                  </el-table-column>

                  <el-table-column prop="money" label="交易金额" width="80" />
                </el-table>

                <el-table :data="tableData.fetchList" style="width: 100%">
                  <el-table-column prop="uid" label="取取用户" width="100"  >
                              <template #default="scope">
                        <el-button type="text" @click="toUserDetail(scope.row.uid)">   查看用户</el-button> 
                        </template>
                  </el-table-column>
                          <el-table-column prop="total" label="订单数量" width="80" >
                    <template #default="scope">
             <el-button type="text" @click="toHelpOrder(scope.row.uid,'helpFetch')">  {{scope.row.total}}</el-button> 
                    </template>
                  </el-table-column>
                  <el-table-column prop="money" label="交易金额" width="80" />
                </el-table>
          </div>
        </el-tab-pane>
    </el-tabs>
    </div>
</div>
</template>

<script>
export default {
  name: 'UserStatistics'
}
</script>

<script setup>
import{
    getNewUserCountList
}from '@/api/minappUser'
import { ref } from 'vue'
import { useRouter } from 'vue-router'
const router = useRouter()
const activeName = ref('addUserInfo')
const rangeDate = ref('')

var myDate = new Date()
    var time=(new Date).getTime()-7*24*60*60*1000;
        var targetTime=new Date(time);
        var month=targetTime.getMonth();
        var day=targetTime.getDate();
        targetTime=targetTime.getFullYear() + "-" + (targetTime.getMonth()> 9 ? (targetTime.getMonth() + 1) : "0" + (targetTime.getMonth() + 1)) + "-" +(targetTime.getDate()> 9 ? (targetTime.getDate()) : "0" + (targetTime.getDate()));

    console.log("开始时间 ",targetTime)

const searchInfo = ref({
  start:targetTime,
  end: dateFormat("YYYY-mm-dd",myDate),
  way:""
})

const toUserDetail = (uid) => {
  console.log("数据",uid)
  router.push({
    name: 'userDetail',
    params: {
      id: uid,
    },
  })
}

const toHelpOrder = (uid,name) =>{
  router.push({
    name:name,
    params:{
      id:uid
    }
  })
}

const onSubmit = () =>{
 
  searchInfo.value.start = rangeDate.value[0]
  searchInfo.value.end = rangeDate.value[1]
  getUserCountData()
}
const shortcuts = [
  {
    text: '最近一周',
    value: () => {
      const end = new Date()
      const start = new Date()
      start.setTime(start.getTime() - 3600 * 1000 * 24 * 7)
      console.log("最近一个周",start,end)
      return [start, end]
    },
  },
  {
    text: '最近一个月',
    value: () => {
      const end = new Date()
      const start = new Date()
      start.setTime(start.getTime() - 3600 * 1000 * 24 * 30)
      return [start, end]
    },
  },
  {
    text: '最近三个月',
    value: () => {
      const end = new Date()
      const start = new Date()
      start.setTime(start.getTime() - 3600 * 1000 * 24 * 90)
      return [start, end]
    },
  },
]


const handleClick = (tab, event) => {

  searchInfo.value.way = activeName.value
  getUserCountData()

}
const tableData = ref({})

const getUserCountData = async() =>{

    const res = await getNewUserCountList({...searchInfo.value})
    if (res.code === 0){
       tableData.value = res.data.data
    }
}
getUserCountData()
function dateFormat(fmt, date){
    let ret;
    const opt = {
        "Y+": date.getFullYear().toString(),        // 年
        "m+": (date.getMonth() + 1).toString(),     // 月
        "d+": date.getDate().toString(),            // 日
        "H+": date.getHours().toString(),           // 时
        "M+": date.getMinutes().toString(),         // 分
        "S+": date.getSeconds().toString()          // 秒
        // 有其他格式化字符需求可以继续添加，必须转化成字符串
    };
    for (let k in opt) {
        ret = new RegExp("(" + k + ")").exec(fmt);
        if (ret) {
            fmt = fmt.replace(ret[1], (ret[1].length == 1) ? (opt[k]) : (opt[k].padStart(ret[1].length, "0")))
        };
    };
    return fmt;
}
</script>
<style>
.demo-tabs > .el-tabs__content {
  color: #6b778c;
  font-size: 32px;
  font-weight: 600;
}
</style>
