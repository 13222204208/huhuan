/*! 
 Build based on gin-vue-admin 
 Time : 1671085367000 */
import{C as a}from"./gin-vue-admin-github.1671085367000.js";import{f as s}from"./gin-vue-admin-date.1671085367000.js";import{_ as e}from"../gva/gin-vue-admin-index.1671085367000.js";import{r as m,o as t,j as o,q as i,J as l,L as d,F as n,G as r,s as c,C as v}from"./gin-vue-admin-vendor.1671085367000.js";const f={class:"commit-table"},u=(a=>(n("data-v-02421634"),a=a(),r(),a))((()=>i("div",{class:"commit-table-title"}," 更新日志 ",-1))),g={class:"log"},p={class:"flex-1 flex key-box"},h={class:"flex-5 flex message"},x={class:"flex-3 flex form"},y={name:"DashboardTable"};var b=e(Object.assign(y,{setup(e){const n=m(!0),r=m([]);return a(0).then((({data:a})=>{n.value=!1,a.forEach(((a,e)=>{a.commit.message&&e<10&&r.value.push({from:s(a.commit.author.date,"yyyy-MM-dd"),title:a.commit.author.name,showDayAndMonth:!0,message:a.commit.message})}))})),(a,s)=>(t(),o("div",f,[u,i("div",g,[(t(!0),o(l,null,d(r.value,((a,s)=>(t(),o("div",{key:s,class:"log-item"},[i("div",p,[i("span",{class:c(["key",s<3&&"top"])},v(s+1),3)]),i("div",h,v(a.message),1),i("div",x,v(a.from),1)])))),128))])]))}}),[["__scopeId","data-v-02421634"]]);export{b as default};