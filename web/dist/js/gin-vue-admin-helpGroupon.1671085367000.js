/*! 
 Build based on gin-vue-admin 
 Time : 1671085367000 */
var e=Object.defineProperty,l=Object.getOwnPropertySymbols,a=Object.prototype.hasOwnProperty,o=Object.prototype.propertyIsEnumerable,t=(l,a,o)=>a in l?e(l,a,{enumerable:!0,configurable:!0,writable:!0,value:o}):l[a]=o,u=(e,l,a)=>new Promise(((o,t)=>{var u=e=>{try{i(a.next(e))}catch(l){t(l)}},d=e=>{try{i(a.throw(e))}catch(l){t(l)}},i=e=>e.done?o(e.value):Promise.resolve(e.value).then(u,d);i((a=a.apply(e,l)).next())}));import{g as d,d as i,c as r,u as n}from"./gin-vue-admin-helpGroupon.16710853670002.js";import{f as m,a as s,g as p}from"./gin-vue-admin-format.1671085367000.js";import{_ as v}from"./gin-vue-admin-index.167108536700024.js";import{r as c,i as f,o as g,j as b,q as h,k as V,t as y,J as _,L as w,A as k,C as N,x as U,e as j,d as T,H as x}from"./gin-vue-admin-vendor.1671085367000.js";import"../gva/gin-vue-admin-index.1671085367000.js";import"./gin-vue-admin-date.1671085367000.js";import"./gin-vue-admin-dictionary.1671085367000.js";import"./gin-vue-admin-dictionary.16710853670002.js";import"./gin-vue-admin-sysDictionary.1671085367000.js";import"./gin-vue-admin-area.16710853670002.js";const D={class:"gva-search-box"},z=k("查询"),C=k("重置"),P={class:"gva-table-box"},I=k("删除"),O={class:"gva-pagination"},S={class:"dialog-footer"},Y=k("取 消"),A=k("确 定"),B={name:"HelpGroupon"},G=Object.assign(B,{setup(e){const B=e=>u(this,null,(function*(){console.log("子组件传的值",e),L.value.area=e})),G=c([]),H=c({orderNum:"",title:"",goodsName:"",goodsTotal:"",price:0,time:"",goodsImg:"",goodsDetail:"",takeName:"",takePhone:"",takeTime:"",userNum:"",status:void 0,city:"",area:""}),M=c(1),q=c(0),E=c(10),J=c([]),L=c({}),F=()=>{L.value={}},K=()=>{M.value=1,E.value=10,W()},Q=e=>{E.value=e,W()},R=e=>{M.value=e,W()},W=()=>u(this,null,(function*(){const e=yield d(((e,u)=>{for(var d in u||(u={}))a.call(u,d)&&t(e,d,u[d]);if(l)for(var d of l(u))o.call(u,d)&&t(e,d,u[d]);return e})({page:M.value,pageSize:E.value},L.value));0===e.code&&(J.value=e.data.list,q.value=e.data.total,M.value=e.data.page,E.value=e.data.pageSize)}));W();(()=>{u(this,null,(function*(){G.value=yield p("orderState")}))})();const X=c([]),Z=e=>{X.value=e};c(!1);const $=c(""),ee=e=>u(this,null,(function*(){0===(yield i({ID:e.ID})).code&&(T({type:"success",message:"删除成功"}),1===J.value.length&&M.value>1&&M.value--,W())})),le=c(!1),ae=()=>{le.value=!1,H.value={orderNum:"",title:"",goodsName:"",goodsTotal:"",price:0,time:"",goodsImg:"",goodsDetail:"",takeName:"",takePhone:"",takeTime:"",userNum:"",status:void 0,city:"",area:""}},oe=()=>u(this,null,(function*(){let e;switch($.value){case"create":e=yield r(H.value);break;case"update":e=yield n(H.value);break;default:e=yield r(H.value)}0===e.code&&(T({type:"success",message:"创建/更改成功"}),ae(),W())}));return(e,l)=>{const a=f("el-date-picker"),o=f("el-form-item"),t=f("el-input"),u=f("el-option"),d=f("el-select"),i=f("el-button"),r=f("el-form"),n=f("el-table-column"),p=f("el-table"),c=f("el-pagination"),T=f("el-input-number"),W=f("el-dialog");return g(),b("div",null,[h("div",D,[V(r,{inline:!0,model:L.value,class:"demo-form-inline"},{default:y((()=>[V(o,{label:"下单日期"},{default:y((()=>[V(a,{modelValue:L.value.creationDate,"onUpdate:modelValue":l[0]||(l[0]=e=>L.value.creationDate=e),type:"date","value-format":"YYYY-MM-DD",placeholder:"选择日期"},null,8,["modelValue"])])),_:1}),V(o,{label:"订单号"},{default:y((()=>[V(t,{modelValue:L.value.orderNum,"onUpdate:modelValue":l[1]||(l[1]=e=>L.value.orderNum=e),placeholder:"搜索条件"},null,8,["modelValue"])])),_:1}),V(o,{label:"商品名称"},{default:y((()=>[V(t,{modelValue:L.value.goodsName,"onUpdate:modelValue":l[2]||(l[2]=e=>L.value.goodsName=e),placeholder:"搜索条件"},null,8,["modelValue"])])),_:1}),V(o,{label:"接单人"},{default:y((()=>[V(t,{modelValue:L.value.takeName,"onUpdate:modelValue":l[3]||(l[3]=e=>L.value.takeName=e),placeholder:"搜索条件"},null,8,["modelValue"])])),_:1}),V(o,{label:"用户编号"},{default:y((()=>[V(t,{modelValue:L.value.userNum,"onUpdate:modelValue":l[4]||(l[4]=e=>L.value.userNum=e),placeholder:"搜索条件"},null,8,["modelValue"])])),_:1}),V(o,{label:"订单状态"},{default:y((()=>[V(d,{modelValue:L.value.status,"onUpdate:modelValue":l[5]||(l[5]=e=>L.value.status=e),class:"m-2",placeholder:"选择状态"},{default:y((()=>[(g(!0),b(_,null,w(G.value,(e=>(g(),x(u,{key:e.value,label:e.label,value:e.value},null,8,["label","value"])))),128))])),_:1},8,["modelValue"])])),_:1}),V(o,{label:"地区"},{default:y((()=>[V(v,{onSelectArea:B})])),_:1}),V(o,null,{default:y((()=>[V(i,{size:"small",type:"primary",icon:"search",onClick:K},{default:y((()=>[z])),_:1}),V(i,{size:"small",icon:"refresh",onClick:F},{default:y((()=>[C])),_:1})])),_:1})])),_:1},8,["model"])]),h("div",P,[V(p,{ref:"multipleTable",style:{width:"100%"},"tooltip-effect":"dark",data:J.value,"row-key":"ID",onSelectionChange:Z},{default:y((()=>[V(n,{align:"left",label:"日期",width:"180"},{default:y((e=>[k(N(U(m)(e.row.CreatedAt)),1)])),_:1}),V(n,{align:"left",label:"订单号",prop:"orderNum",width:"220"}),V(n,{align:"left",label:"团购标题",prop:"title",width:"120"}),V(n,{align:"left",label:"商品名称",prop:"goodsName",width:"120"}),V(n,{align:"left",label:"剩于可团人数",prop:"goodsTotal",width:"120"},{default:y((e=>[k(N(e.row.goodsTotal-e.row.follow<1?"抢光了":e.row.goodsTotal-e.row.follow),1)])),_:1}),V(n,{align:"left",label:"商品价格",prop:"price",width:"120"}),V(n,{align:"left",label:"团购时长",prop:"time",width:"120"}),V(n,{align:"left",label:"商品图片",prop:"goodsImg",width:"120"}),V(n,{align:"left",label:"商品详情",prop:"goodsDetail",width:"120"}),V(n,{align:"left",label:"接单人",prop:"takeName",width:"120"}),V(n,{align:"left",label:"接单人电话",prop:"takePhone",width:"120"}),V(n,{align:"left",label:"接单时间",prop:"takeTime",width:"120"}),V(n,{align:"left",label:"用户编号",prop:"userNum",width:"220"}),V(n,{align:"left",label:"订单状态",prop:"status",width:"120"},{default:y((e=>[k(N(U(s)(e.row.status,G.value)),1)])),_:1}),V(n,{align:"left",label:"城市",prop:"city",width:"120"}),V(n,{align:"left",label:"地区",prop:"area",width:"120"}),V(n,{align:"left",label:"按钮组"},{default:y((e=>[V(i,{type:"text",icon:"delete",size:"small",onClick:l=>{return a=e.row,void j.confirm("确定要删除吗?","提示",{confirmButtonText:"确定",cancelButtonText:"取消",type:"warning"}).then((()=>{ee(a)}));var a}},{default:y((()=>[I])),_:2},1032,["onClick"])])),_:1})])),_:1},8,["data"]),h("div",O,[V(c,{layout:"total, sizes, prev, pager, next, jumper","current-page":M.value,"page-size":E.value,"page-sizes":[10,30,50,100],total:q.value,onCurrentChange:R,onSizeChange:Q},null,8,["current-page","page-size","total"])])]),V(W,{modelValue:le.value,"onUpdate:modelValue":l[21]||(l[21]=e=>le.value=e),"before-close":ae,title:"弹窗操作"},{footer:y((()=>[h("div",S,[V(i,{size:"small",onClick:ae},{default:y((()=>[Y])),_:1}),V(i,{size:"small",type:"primary",onClick:oe},{default:y((()=>[A])),_:1})])])),default:y((()=>[V(r,{model:H.value,"label-position":"right","label-width":"80px"},{default:y((()=>[V(o,{label:"订单号:"},{default:y((()=>[V(t,{modelValue:H.value.orderNum,"onUpdate:modelValue":l[6]||(l[6]=e=>H.value.orderNum=e),clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),V(o,{label:"团购标题:"},{default:y((()=>[V(t,{modelValue:H.value.title,"onUpdate:modelValue":l[7]||(l[7]=e=>H.value.title=e),clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),V(o,{label:"商品名称:"},{default:y((()=>[V(t,{modelValue:H.value.goodsName,"onUpdate:modelValue":l[8]||(l[8]=e=>H.value.goodsName=e),clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),V(o,{label:"商品数量:"},{default:y((()=>[V(t,{modelValue:H.value.goodsTotal,"onUpdate:modelValue":l[9]||(l[9]=e=>H.value.goodsTotal=e),clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),V(o,{label:"商品价格:"},{default:y((()=>[V(T,{modelValue:H.value.price,"onUpdate:modelValue":l[10]||(l[10]=e=>H.value.price=e),style:{width:"100%"},precision:2,clearable:""},null,8,["modelValue"])])),_:1}),V(o,{label:"团购时长:"},{default:y((()=>[V(t,{modelValue:H.value.time,"onUpdate:modelValue":l[11]||(l[11]=e=>H.value.time=e),clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),V(o,{label:"商品图片:"},{default:y((()=>[V(t,{modelValue:H.value.goodsImg,"onUpdate:modelValue":l[12]||(l[12]=e=>H.value.goodsImg=e),clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),V(o,{label:"商品详情:"},{default:y((()=>[V(t,{modelValue:H.value.goodsDetail,"onUpdate:modelValue":l[13]||(l[13]=e=>H.value.goodsDetail=e),clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),V(o,{label:"接单人:"},{default:y((()=>[V(t,{modelValue:H.value.takeName,"onUpdate:modelValue":l[14]||(l[14]=e=>H.value.takeName=e),clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),V(o,{label:"接单人电话:"},{default:y((()=>[V(t,{modelValue:H.value.takePhone,"onUpdate:modelValue":l[15]||(l[15]=e=>H.value.takePhone=e),clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),V(o,{label:"接单时间:"},{default:y((()=>[V(t,{modelValue:H.value.takeTime,"onUpdate:modelValue":l[16]||(l[16]=e=>H.value.takeTime=e),clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),V(o,{label:"用户编号:"},{default:y((()=>[V(t,{modelValue:H.value.userNum,"onUpdate:modelValue":l[17]||(l[17]=e=>H.value.userNum=e),clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),V(o,{label:"订单状态:"},{default:y((()=>[V(d,{modelValue:H.value.status,"onUpdate:modelValue":l[18]||(l[18]=e=>H.value.status=e),placeholder:"请选择",style:{width:"100%"},clearable:""},{default:y((()=>[(g(!0),b(_,null,w(G.value,((e,l)=>(g(),x(u,{key:l,label:e.label,value:e.value},null,8,["label","value"])))),128))])),_:1},8,["modelValue"])])),_:1}),V(o,{label:"城市:"},{default:y((()=>[V(t,{modelValue:H.value.city,"onUpdate:modelValue":l[19]||(l[19]=e=>H.value.city=e),clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),V(o,{label:"地区:"},{default:y((()=>[V(t,{modelValue:H.value.area,"onUpdate:modelValue":l[20]||(l[20]=e=>H.value.area=e),clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1})])),_:1},8,["model"])])),_:1},8,["modelValue"])])}}});export{G as default};