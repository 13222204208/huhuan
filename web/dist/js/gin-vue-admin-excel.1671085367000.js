/*! 
 Build based on gin-vue-admin 
 Time : 1671085367000 */
import{s as e,_ as a,u as l,g as t}from"../gva/gin-vue-admin-index.1671085367000.js";import{d as o,r as s,i as n,o as i,j as d,q as r,k as p,t as c,x as u,C as m,A as v}from"./gin-vue-admin-vendor.1671085367000.js";const x=(e,a)=>{if(void 0!==e.data){if("application/json"===e.data.type){const a=new FileReader;a.onload=function(){const e=JSON.parse(a.result).msg;o({showClose:!0,message:e,type:"error"})},a.readAsText(new Blob([e.data]))}}else{var l=window.URL.createObjectURL(new Blob([e])),t=document.createElement("a");t.style.display="none",t.href=l,t.download=a;var s=new MouseEvent("click");t.dispatchEvent(s)}},f=()=>e({url:"/excel/loadExcel",method:"get"});const h={class:"upload"},w={class:"gva-table-box"},b={class:"gva-btn-list"},g=v("导入"),y=v("导出"),E=v("下载模板"),_={name:"Excel"};var k=a(Object.assign(_,{setup(a){const o=s("/api"),v=s(1),_=s(0),k=s(999),j=s([]),z=(e=(()=>{}))=>{return a=this,l=null,t=function*(){const a=yield e({page:v.value,pageSize:k.value});0===a.code&&(j.value=a.data.list,_.value=a.data.total,v.value=a.data.page,k.value=a.data.pageSize)},new Promise(((e,o)=>{var s=e=>{try{i(t.next(e))}catch(a){o(a)}},n=e=>{try{i(t.throw(e))}catch(a){o(a)}},i=a=>a.done?e(a.value):Promise.resolve(a.value).then(s,n);i((t=t.apply(a,l)).next())}));var a,l,t};z(t);const I=l(),T=a=>{a&&"string"==typeof a||(a="ExcelExport.xlsx"),((a,l)=>{e({url:"/excel/exportExcel",method:"post",data:{fileName:l,infoList:a},responseType:"blob"}).then((e=>{x(e,l)}))})(j.value,a)},C=()=>{z(f)},N=()=>{var a;e({url:"/excel/downloadTemplate",method:"get",params:{fileName:a="ExcelTemplate.xlsx"},responseType:"blob"}).then((e=>{x(e,a)}))};return(e,a)=>{const l=n("el-button"),t=n("el-upload"),s=n("el-table-column"),v=n("el-table");return i(),d("div",h,[r("div",w,[r("div",b,[p(t,{class:"excel-btn",action:`${o.value}/excel/importExcel`,headers:{"x-token":u(I).token},"on-success":C,"show-file-list":!1},{default:c((()=>[p(l,{size:"small",type:"primary",icon:"upload"},{default:c((()=>[g])),_:1})])),_:1},8,["action","headers"]),p(l,{class:"excel-btn",size:"small",type:"primary",icon:"download",onClick:a[0]||(a[0]=e=>T("ExcelExport.xlsx"))},{default:c((()=>[y])),_:1}),p(l,{class:"excel-btn",size:"small",type:"success",icon:"download",onClick:a[1]||(a[1]=e=>N())},{default:c((()=>[E])),_:1})]),p(v,{data:j.value,"row-key":"ID"},{default:c((()=>[p(s,{align:"left",label:"ID","min-width":"100",prop:"ID"}),p(s,{align:"left","show-overflow-tooltip":"",label:"路由Name","min-width":"160",prop:"name"}),p(s,{align:"left","show-overflow-tooltip":"",label:"路由Path","min-width":"160",prop:"path"}),p(s,{align:"left",label:"是否隐藏","min-width":"100",prop:"hidden"},{default:c((e=>[r("span",null,m(e.row.hidden?"隐藏":"显示"),1)])),_:1}),p(s,{align:"left",label:"父节点","min-width":"90",prop:"parentId"}),p(s,{align:"left",label:"排序","min-width":"70",prop:"sort"}),p(s,{align:"left",label:"文件路径","min-width":"360",prop:"component"})])),_:1},8,["data"])])])}}}),[["__scopeId","data-v-4ff7d823"]]);export{k as default};
