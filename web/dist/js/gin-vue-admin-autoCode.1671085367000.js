/*! 
 Build based on gin-vue-admin 
 Time : 1671085367000 */
import{s as t}from"../gva/gin-vue-admin-index.1671085367000.js";const o=o=>t({url:"/autoCode/preview",method:"post",data:o}),a=o=>t({url:"/autoCode/createTemp",method:"post",data:o,responseType:"blob"}),e=()=>t({url:"/autoCode/getDB",method:"get"}),d=o=>t({url:"/autoCode/getTables",method:"get",params:o}),s=o=>t({url:"/autoCode/getColumn",method:"get",params:o}),r=o=>t({url:"/autoCode/getSysHistory",method:"post",data:o}),u=o=>t({url:"/autoCode/rollback",method:"post",data:o}),m=o=>t({url:"/autoCode/getMeta",method:"post",data:o}),l=o=>t({url:"/autoCode/delSysHistory",method:"post",data:o});export{s as a,e as b,a as c,m as d,r as e,l as f,d as g,o as p,u as r};
