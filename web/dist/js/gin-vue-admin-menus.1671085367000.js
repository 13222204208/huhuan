/*! 
 Build based on gin-vue-admin 
 Time : 1671085367000 */
var e=(e,t,a)=>new Promise(((l,o)=>{var n=e=>{try{u(a.next(e))}catch(t){o(t)}},s=e=>{try{u(a.throw(e))}catch(t){o(t)}},u=e=>e.done?l(e.value):Promise.resolve(e.value).then(n,s);u((a=a.apply(e,t)).next())}));import{h as t,i as a,j as l}from"../gva/gin-vue-admin-index.1671085367000.js";import{u as o}from"./gin-vue-admin-authority.16710853670002.js";import{g as n,s}from"./gin-vue-admin-authorityBtn.1671085367000.js";import{r as u,i as d,o as r,j as i,q as c,k as m,t as h,C as y,O as f,A as p,v,d as g,R as I}from"./gin-vue-admin-vendor.1671085367000.js";const k={class:"clearflex"},w=p("确 定"),b={class:"custom-tree-node"},C={key:0},x=p(" 分配按钮 "),R={class:"dialog-footer"},_=p("取 消"),j=p("确 定"),D={name:"Menus"},z=Object.assign(D,{props:{row:{default:function(){return{}},type:Object}},emits:["changeRow"],setup(D,{expose:z,emit:N}){const A=D,B=u([]),E=u([]),O=u(!1),V=u({children:"children",label:function(e){return e.meta.title}});(()=>{e(this,null,(function*(){const e=yield t();B.value=e.data.menus;const l=(yield a({authorityId:A.row.authorityId})).data.menus,o=[];l.forEach((e=>{l.some((t=>t.parentId===e.menuId))||o.push(Number(e.menuId))})),E.value=o}))})();const P=t=>e(this,null,(function*(){const e=yield o({authorityId:A.row.authorityId,AuthorityName:A.row.authorityName,parentId:A.row.parentId,defaultRouter:t.name});0===e.code&&(g({type:"success",message:"设置成功"}),N("changeRow","defaultRouter",e.data.authority.defaultRouter))})),S=()=>{O.value=!0},T=u(null),q=()=>e(this,null,(function*(){const e=T.value.getCheckedNodes(!1,!0);0===(yield l({menus:e,authorityId:A.row.authorityId})).code&&g({type:"success",message:"菜单设置成功!"})}));z({enterAndNext:()=>{q()},needConfirm:O});const M=u(!1),U=u([]),F=u([]),G=u();let H="";const J=t=>e(this,null,(function*(){H=t.ID;const e=yield n({menuID:H,authorityId:A.row.authorityId});0===e.code&&(L(t),yield I(),e.data.selected&&e.data.selected.forEach((e=>{U.value.some((t=>{t.ID===e&&G.value.toggleRowSelection(t,!0)}))})))})),K=e=>{F.value=e},L=e=>{M.value=!0,U.value=e.menuBtn},Q=()=>{M.value=!1},W=()=>e(this,null,(function*(){const e=F.value.map((e=>e.ID));0===(yield s({menuID:H,selected:e,authorityId:A.row.authorityId})).code&&(g({type:"success",message:"设置成功"}),M.value=!1)}));return(e,t)=>{const a=d("el-button"),l=d("el-tree"),o=d("el-table-column"),n=d("el-table"),s=d("el-dialog");return r(),i("div",null,[c("div",k,[m(a,{class:"fl-right",size:"small",type:"primary",onClick:q},{default:h((()=>[w])),_:1})]),m(l,{ref_key:"menuTree",ref:T,data:B.value,"default-checked-keys":E.value,props:V.value,"default-expand-all":"","highlight-current":"","node-key":"ID","show-checkbox":"",onCheck:S},{default:h((({node:e,data:t})=>[c("span",b,[c("span",null,y(e.label),1),c("span",null,[m(a,{type:"text",size:"small",style:f({color:D.row.defaultRouter===t.name?"#E6A23C":"#85ce61"}),disabled:!e.checked,onClick:()=>P(t)},{default:h((()=>[p(y(D.row.defaultRouter===t.name?"首页":"设为首页"),1)])),_:2},1032,["style","disabled","onClick"])]),t.menuBtn.length?(r(),i("span",C,[m(a,{type:"text",size:"small",onClick:()=>J(t)},{default:h((()=>[x])),_:2},1032,["onClick"])])):v("",!0)])])),_:1},8,["data","default-checked-keys","props"]),m(s,{modelValue:M.value,"onUpdate:modelValue":t[0]||(t[0]=e=>M.value=e),title:"分配按钮","destroy-on-close":""},{footer:h((()=>[c("div",R,[m(a,{size:"small",onClick:Q},{default:h((()=>[_])),_:1}),m(a,{size:"small",type:"primary",onClick:W},{default:h((()=>[j])),_:1})])])),default:h((()=>[m(n,{ref_key:"btnTableRef",ref:G,data:U.value,"row-key":"ID",onSelectionChange:K},{default:h((()=>[m(o,{type:"selection",width:"55"}),m(o,{label:"按钮名称",prop:"name"}),m(o,{label:"按钮备注",prop:"desc"})])),_:1},8,["data"])])),_:1},8,["modelValue"])])}}});export{z as default};