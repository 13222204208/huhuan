/*! 
 Build based on gin-vue-admin 
 Time : 1671085367000 */
var e=Object.defineProperty,l=Object.getOwnPropertySymbols,a=Object.prototype.hasOwnProperty,o=Object.prototype.propertyIsEnumerable,u=(l,a,o)=>a in l?e(l,a,{enumerable:!0,configurable:!0,writable:!0,value:o}):l[a]=o,t=(e,l,a)=>new Promise(((o,u)=>{var t=e=>{try{n(a.next(e))}catch(l){u(l)}},d=e=>{try{n(a.throw(e))}catch(l){u(l)}},n=e=>e.done?o(e.value):Promise.resolve(e.value).then(t,d);n((a=a.apply(e,l)).next())}));import{g as d,d as n,c as r,u as i}from"./gin-vue-admin-helpFetch.16710853670002.js";import{f as m,a as p,g as s}from"./gin-vue-admin-format.1671085367000.js";import{_ as c}from"./gin-vue-admin-index.167108536700024.js";import{U as v,r as f,i as b,o as h,j as V,q as g,k as _,t as k,J as y,L as w,A as U,C as N,x as A,e as j,d as C,H as P}from"./gin-vue-admin-vendor.1671085367000.js";import"../gva/gin-vue-admin-index.1671085367000.js";import"./gin-vue-admin-date.1671085367000.js";import"./gin-vue-admin-dictionary.1671085367000.js";import"./gin-vue-admin-dictionary.16710853670002.js";import"./gin-vue-admin-sysDictionary.1671085367000.js";import"./gin-vue-admin-area.16710853670002.js";const x={class:"gva-search-box"},T=U("查询"),z=U("重置"),I={class:"gva-table-box"},D=U("删除"),O={class:"gva-pagination"},S={class:"dialog-footer"},Y=U("取 消"),M=U("确 定"),B={name:"HelpFetch"},F=Object.assign(B,{setup(e){const B=v(),F=e=>t(this,null,(function*(){console.log("子组件传的值",e),K.value.area=e})),H=f([]),q=f({orderNum:"",goodsName:"",goodsCount:"",linkman:"",contactPhone:"",fetchAddress:"",receiveAddress:"",time:"",remark:"",couponId:0,couponAmount:0,price:0,takeName:"",takePhone:"",takeTime:"",userNum:"",status:void 0,doneTime:"",city:"",area:""}),E=f(1),J=f(0),L=f(10),G=f([]),K=f({}),Q=()=>{K.value={}},R=()=>{E.value=1,L.value=10,Z()},W=e=>{L.value=e,Z()},X=e=>{E.value=e,Z()},Z=()=>t(this,null,(function*(){const e=yield d(((e,t)=>{for(var d in t||(t={}))a.call(t,d)&&u(e,d,t[d]);if(l)for(var d of l(t))o.call(t,d)&&u(e,d,t[d]);return e})({page:E.value,pageSize:L.value},K.value));0===e.code&&(G.value=e.data.list,J.value=e.data.total,E.value=e.data.page,L.value=e.data.pageSize)}));Z();(()=>{t(this,null,(function*(){B.params.id&&(K.value.uid=B.params.id,Z())}))})();(()=>{t(this,null,(function*(){H.value=yield s("orderState"),console.log("字典",H.value)}))})();const $=f([]),ee=e=>{$.value=e};f(!1);const le=f(""),ae=e=>t(this,null,(function*(){0===(yield n({ID:e.ID})).code&&(C({type:"success",message:"删除成功"}),1===G.value.length&&E.value>1&&E.value--,Z())})),oe=f(!1),ue=()=>{oe.value=!1,q.value={orderNum:"",goodsName:"",goodsCount:"",linkman:"",contactPhone:"",fetchAddress:"",receiveAddress:"",time:"",remark:"",couponId:0,couponAmount:0,price:0,takeName:"",takePhone:"",takeTime:"",userNum:"",status:void 0,doneTime:"",city:"",area:""}},te=()=>t(this,null,(function*(){let e;switch(le.value){case"create":e=yield r(q.value);break;case"update":e=yield i(q.value);break;default:e=yield r(q.value)}0===e.code&&(C({type:"success",message:"创建/更改成功"}),ue(),Z())}));return(e,l)=>{const a=b("el-date-picker"),o=b("el-form-item"),u=b("el-input"),t=b("el-option"),d=b("el-select"),n=b("el-button"),r=b("el-form"),i=b("el-table-column"),s=b("el-table"),v=b("el-pagination"),f=b("el-input-number"),C=b("el-dialog");return h(),V("div",null,[g("div",x,[_(r,{inline:!0,model:K.value,class:"demo-form-inline"},{default:k((()=>[_(o,{label:"下单日期"},{default:k((()=>[_(a,{modelValue:K.value.creationDate,"onUpdate:modelValue":l[0]||(l[0]=e=>K.value.creationDate=e),type:"date","value-format":"YYYY-MM-DD",placeholder:"选择日期"},null,8,["modelValue"])])),_:1}),_(o,{label:"订单号"},{default:k((()=>[_(u,{modelValue:K.value.orderNum,"onUpdate:modelValue":l[1]||(l[1]=e=>K.value.orderNum=e),placeholder:"搜索条件"},null,8,["modelValue"])])),_:1}),_(o,{label:"联系人"},{default:k((()=>[_(u,{modelValue:K.value.linkman,"onUpdate:modelValue":l[2]||(l[2]=e=>K.value.linkman=e),placeholder:"搜索条件"},null,8,["modelValue"])])),_:1}),_(o,{label:"联系电话"},{default:k((()=>[_(u,{modelValue:K.value.contactPhone,"onUpdate:modelValue":l[3]||(l[3]=e=>K.value.contactPhone=e),placeholder:"搜索条件"},null,8,["modelValue"])])),_:1}),_(o,{label:"接单人"},{default:k((()=>[_(u,{modelValue:K.value.takeName,"onUpdate:modelValue":l[4]||(l[4]=e=>K.value.takeName=e),placeholder:"搜索条件"},null,8,["modelValue"])])),_:1}),_(o,{label:"用户编号"},{default:k((()=>[_(u,{modelValue:K.value.userNum,"onUpdate:modelValue":l[5]||(l[5]=e=>K.value.userNum=e),placeholder:"搜索条件"},null,8,["modelValue"])])),_:1}),_(o,{label:"订单状态"},{default:k((()=>[_(d,{modelValue:K.value.status,"onUpdate:modelValue":l[6]||(l[6]=e=>K.value.status=e),class:"m-2",placeholder:"选择状态"},{default:k((()=>[(h(!0),V(y,null,w(H.value,(e=>(h(),P(t,{key:e.value,label:e.label,value:e.value},null,8,["label","value"])))),128))])),_:1},8,["modelValue"])])),_:1}),_(o,{label:"地区"},{default:k((()=>[_(c,{onSelectArea:F})])),_:1}),_(o,null,{default:k((()=>[_(n,{size:"small",type:"primary",icon:"search",onClick:R},{default:k((()=>[T])),_:1}),_(n,{size:"small",icon:"refresh",onClick:Q},{default:k((()=>[z])),_:1})])),_:1})])),_:1},8,["model"])]),g("div",I,[_(s,{ref:"multipleTable",style:{width:"100%"},"tooltip-effect":"dark",data:G.value,"row-key":"ID",onSelectionChange:ee},{default:k((()=>[_(i,{align:"left",label:"日期",width:"180"},{default:k((e=>[U(N(A(m)(e.row.CreatedAt)),1)])),_:1}),_(i,{align:"left",label:"订单号",prop:"orderNum",width:"220"}),_(i,{align:"left",label:"物品名称",prop:"goodsName",width:"220"}),_(i,{align:"left",label:"物品数量",prop:"goodsCount",width:"120"}),_(i,{align:"left",label:"联系人",prop:"linkman",width:"120"}),_(i,{align:"left",label:"联系电话",prop:"contactPhone",width:"120"}),_(i,{align:"left",label:"取件地址",prop:"fetchAddress",width:"220"}),_(i,{align:"left",label:"收货地址",prop:"receiveAddress",width:"220"}),_(i,{align:"left",label:"送达时间",prop:"time",width:"120"}),_(i,{align:"left",label:"留言",prop:"remark",width:"120"}),_(i,{align:"left",label:"优惠券",prop:"couponInfo",width:"120"},{default:k((e=>[U(N(""==e.row.couponInfo.title?"无":e.row.couponInfo.title),1)])),_:1}),_(i,{align:"left",label:"优惠金额",prop:"couponAmount",width:"120"}),_(i,{align:"left",label:"赏金",prop:"price",width:"120"}),_(i,{align:"left",label:"接单人",prop:"takeName",width:"120"}),_(i,{align:"left",label:"接单人电话",prop:"takePhone",width:"120"}),_(i,{align:"left",label:"接单时间",prop:"takeTime",width:"120"}),_(i,{align:"left",label:"用户编号",prop:"userNum",width:"120"}),_(i,{align:"left",label:"订单状态",prop:"status",width:"120"},{default:k((e=>[U(N(A(p)(e.row.status,H.value)),1)])),_:1}),_(i,{align:"left",label:"完成时间",prop:"doneTime",width:"120"}),_(i,{align:"left",label:"城市",prop:"city",width:"120"}),_(i,{align:"left",label:"地区",prop:"area",width:"120"}),_(i,{align:"left",label:"按钮组"},{default:k((e=>[_(n,{type:"text",icon:"delete",size:"small",onClick:l=>{return a=e.row,void j.confirm("确定要删除吗?","提示",{confirmButtonText:"确定",cancelButtonText:"取消",type:"warning"}).then((()=>{ae(a)}));var a}},{default:k((()=>[D])),_:2},1032,["onClick"])])),_:1})])),_:1},8,["data"]),g("div",O,[_(v,{layout:"total, sizes, prev, pager, next, jumper","current-page":E.value,"page-size":L.value,"page-sizes":[10,30,50,100],total:J.value,onCurrentChange:X,onSizeChange:W},null,8,["current-page","page-size","total"])])]),_(C,{modelValue:oe.value,"onUpdate:modelValue":l[27]||(l[27]=e=>oe.value=e),"before-close":ue,title:"弹窗操作"},{footer:k((()=>[g("div",S,[_(n,{size:"small",onClick:ue},{default:k((()=>[Y])),_:1}),_(n,{size:"small",type:"primary",onClick:te},{default:k((()=>[M])),_:1})])])),default:k((()=>[_(r,{model:q.value,"label-position":"right","label-width":"80px"},{default:k((()=>[_(o,{label:"订单号:"},{default:k((()=>[_(u,{modelValue:q.value.orderNum,"onUpdate:modelValue":l[7]||(l[7]=e=>q.value.orderNum=e),clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),_(o,{label:"物品名称:"},{default:k((()=>[_(u,{modelValue:q.value.goodsName,"onUpdate:modelValue":l[8]||(l[8]=e=>q.value.goodsName=e),clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),_(o,{label:"物品数量:"},{default:k((()=>[_(u,{modelValue:q.value.goodsCount,"onUpdate:modelValue":l[9]||(l[9]=e=>q.value.goodsCount=e),clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),_(o,{label:"联系人:"},{default:k((()=>[_(u,{modelValue:q.value.linkman,"onUpdate:modelValue":l[10]||(l[10]=e=>q.value.linkman=e),clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),_(o,{label:"联系电话:"},{default:k((()=>[_(u,{modelValue:q.value.contactPhone,"onUpdate:modelValue":l[11]||(l[11]=e=>q.value.contactPhone=e),clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),_(o,{label:"取件地址:"},{default:k((()=>[_(u,{modelValue:q.value.fetchAddress,"onUpdate:modelValue":l[12]||(l[12]=e=>q.value.fetchAddress=e),clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),_(o,{label:"收货地址:"},{default:k((()=>[_(u,{modelValue:q.value.receiveAddress,"onUpdate:modelValue":l[13]||(l[13]=e=>q.value.receiveAddress=e),clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),_(o,{label:"送达时间:"},{default:k((()=>[_(u,{modelValue:q.value.time,"onUpdate:modelValue":l[14]||(l[14]=e=>q.value.time=e),clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),_(o,{label:"留言:"},{default:k((()=>[_(u,{modelValue:q.value.remark,"onUpdate:modelValue":l[15]||(l[15]=e=>q.value.remark=e),clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),_(o,{label:"优惠券ID:"},{default:k((()=>[_(u,{modelValue:q.value.couponId,"onUpdate:modelValue":l[16]||(l[16]=e=>q.value.couponId=e),modelModifiers:{number:!0},clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),_(o,{label:"优惠金额:"},{default:k((()=>[_(f,{modelValue:q.value.couponAmount,"onUpdate:modelValue":l[17]||(l[17]=e=>q.value.couponAmount=e),style:{width:"100%"},precision:2,clearable:""},null,8,["modelValue"])])),_:1}),_(o,{label:"赏金:"},{default:k((()=>[_(f,{modelValue:q.value.price,"onUpdate:modelValue":l[18]||(l[18]=e=>q.value.price=e),style:{width:"100%"},precision:2,clearable:""},null,8,["modelValue"])])),_:1}),_(o,{label:"接单人:"},{default:k((()=>[_(u,{modelValue:q.value.takeName,"onUpdate:modelValue":l[19]||(l[19]=e=>q.value.takeName=e),clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),_(o,{label:"接单人电话:"},{default:k((()=>[_(u,{modelValue:q.value.takePhone,"onUpdate:modelValue":l[20]||(l[20]=e=>q.value.takePhone=e),clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),_(o,{label:"接单时间:"},{default:k((()=>[_(u,{modelValue:q.value.takeTime,"onUpdate:modelValue":l[21]||(l[21]=e=>q.value.takeTime=e),clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),_(o,{label:"用户编号:"},{default:k((()=>[_(u,{modelValue:q.value.userNum,"onUpdate:modelValue":l[22]||(l[22]=e=>q.value.userNum=e),clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),_(o,{label:"订单状态:"},{default:k((()=>[_(d,{modelValue:q.value.status,"onUpdate:modelValue":l[23]||(l[23]=e=>q.value.status=e),placeholder:"请选择",style:{width:"100%"},clearable:""},{default:k((()=>[(h(!0),V(y,null,w(H.value,((e,l)=>(h(),P(t,{key:l,label:e.label,value:e.value},null,8,["label","value"])))),128))])),_:1},8,["modelValue"])])),_:1}),_(o,{label:"完成时间:"},{default:k((()=>[_(u,{modelValue:q.value.doneTime,"onUpdate:modelValue":l[24]||(l[24]=e=>q.value.doneTime=e),clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),_(o,{label:"城市:"},{default:k((()=>[_(u,{modelValue:q.value.city,"onUpdate:modelValue":l[25]||(l[25]=e=>q.value.city=e),clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),_(o,{label:"地区:"},{default:k((()=>[_(u,{modelValue:q.value.area,"onUpdate:modelValue":l[26]||(l[26]=e=>q.value.area=e),clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1})])),_:1},8,["model"])])),_:1},8,["modelValue"])])}}});export{F as default};
