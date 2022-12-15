/*! 
 Build based on gin-vue-admin 
 Time : 1671085367000 */
var e=Object.defineProperty,l=Object.getOwnPropertySymbols,a=Object.prototype.hasOwnProperty,o=Object.prototype.propertyIsEnumerable,u=(l,a,o)=>a in l?e(l,a,{enumerable:!0,configurable:!0,writable:!0,value:o}):l[a]=o,t=(e,l,a)=>new Promise(((o,u)=>{var t=e=>{try{n(a.next(e))}catch(l){u(l)}},d=e=>{try{n(a.throw(e))}catch(l){u(l)}},n=e=>e.done?o(e.value):Promise.resolve(e.value).then(t,d);n((a=a.apply(e,l)).next())}));import{g as d,d as n,c as i,u as r}from"./gin-vue-admin-helpWork.16710853670002.js";import{f as m,a as p,g as c}from"./gin-vue-admin-format.1671085367000.js";import{_ as s}from"./gin-vue-admin-index.167108536700024.js";import{U as v,r as f,i as b,o as h,j as V,q as g,k,t as _,J as y,L as w,A as U,C as j,x as P,e as N,d as x,H as T}from"./gin-vue-admin-vendor.1671085367000.js";import"../gva/gin-vue-admin-index.1671085367000.js";import"./gin-vue-admin-date.1671085367000.js";import"./gin-vue-admin-dictionary.1671085367000.js";import"./gin-vue-admin-dictionary.16710853670002.js";import"./gin-vue-admin-sysDictionary.1671085367000.js";import"./gin-vue-admin-area.16710853670002.js";const z={class:"gva-search-box"},C=U("查询"),I=U("重置"),D={class:"gva-table-box"},A=U("删除"),O={class:"gva-pagination"},S={class:"dialog-footer"},Y=U("取 消"),M=U("确 定"),B={name:"HelpWork"},H=Object.assign(B,{setup(e){const B=v(),H=e=>t(this,null,(function*(){console.log("子组件传的值",e),G.value.area=e})),W=f([]),q=f({orderNum:"",contents:"",linkman:"",contactPhone:"",address:"",time:"",remark:"",couponId:0,couponAmount:0,price:0,takeName:"",takePhone:"",takeTime:"",uid:"",status:void 0,doneTime:"",area:""}),E=f(1),J=f(0),L=f(10),F=f([]),G=f({}),K=()=>{G.value={}},Q=()=>{E.value=1,L.value=10,Z()},R=e=>{L.value=e,Z()},X=e=>{E.value=e,Z()},Z=()=>t(this,null,(function*(){const e=yield d(((e,t)=>{for(var d in t||(t={}))a.call(t,d)&&u(e,d,t[d]);if(l)for(var d of l(t))o.call(t,d)&&u(e,d,t[d]);return e})({page:E.value,pageSize:L.value},G.value));0===e.code&&(F.value=e.data.list,J.value=e.data.total,E.value=e.data.page,L.value=e.data.pageSize)}));Z();(()=>{t(this,null,(function*(){B.params.id&&(G.value.uid=B.params.id,Z())}))})();(()=>{t(this,null,(function*(){W.value=yield c("orderState")}))})();const $=f([]),ee=e=>{$.value=e};f(!1);const le=f(""),ae=e=>t(this,null,(function*(){0===(yield n({ID:e.ID})).code&&(x({type:"success",message:"删除成功"}),1===F.value.length&&E.value>1&&E.value--,Z())})),oe=f(!1),ue=()=>{oe.value=!1,q.value={orderNum:"",contents:"",linkman:"",contactPhone:"",address:"",time:"",remark:"",couponId:0,couponAmount:0,price:0,takeName:"",takePhone:"",takeTime:"",uid:"",status:void 0,doneTime:"",area:""}},te=()=>t(this,null,(function*(){let e;switch(le.value){case"create":e=yield i(q.value);break;case"update":e=yield r(q.value);break;default:e=yield i(q.value)}0===e.code&&(x({type:"success",message:"创建/更改成功"}),ue(),Z())}));return(e,l)=>{const a=b("el-date-picker"),o=b("el-form-item"),u=b("el-input"),t=b("el-option"),d=b("el-select"),n=b("el-button"),i=b("el-form"),r=b("el-table-column"),c=b("el-table"),v=b("el-pagination"),f=b("el-input-number"),x=b("el-dialog");return h(),V("div",null,[g("div",z,[k(i,{inline:!0,model:G.value,class:"demo-form-inline"},{default:_((()=>[k(o,{label:"下单日期"},{default:_((()=>[k(a,{modelValue:G.value.creationDate,"onUpdate:modelValue":l[0]||(l[0]=e=>G.value.creationDate=e),type:"date","value-format":"YYYY-MM-DD",placeholder:"选择日期"},null,8,["modelValue"])])),_:1}),k(o,{label:"订单号"},{default:_((()=>[k(u,{modelValue:G.value.orderNum,"onUpdate:modelValue":l[1]||(l[1]=e=>G.value.orderNum=e),placeholder:"搜索条件"},null,8,["modelValue"])])),_:1}),k(o,{label:"联系人"},{default:_((()=>[k(u,{modelValue:G.value.linkman,"onUpdate:modelValue":l[2]||(l[2]=e=>G.value.linkman=e),placeholder:"搜索条件"},null,8,["modelValue"])])),_:1}),k(o,{label:"联系电话"},{default:_((()=>[k(u,{modelValue:G.value.contactPhone,"onUpdate:modelValue":l[3]||(l[3]=e=>G.value.contactPhone=e),placeholder:"搜索条件"},null,8,["modelValue"])])),_:1}),k(o,{label:"接单人"},{default:_((()=>[k(u,{modelValue:G.value.takeName,"onUpdate:modelValue":l[4]||(l[4]=e=>G.value.takeName=e),placeholder:"搜索条件"},null,8,["modelValue"])])),_:1}),k(o,{label:"用户编号"},{default:_((()=>[k(u,{modelValue:G.value.uid,"onUpdate:modelValue":l[5]||(l[5]=e=>G.value.uid=e),placeholder:"搜索条件"},null,8,["modelValue"])])),_:1}),k(o,{label:"订单状态"},{default:_((()=>[k(d,{modelValue:G.value.status,"onUpdate:modelValue":l[6]||(l[6]=e=>G.value.status=e),class:"m-2",placeholder:"选择状态"},{default:_((()=>[(h(!0),V(y,null,w(W.value,(e=>(h(),T(t,{key:e.value,label:e.label,value:e.value},null,8,["label","value"])))),128))])),_:1},8,["modelValue"])])),_:1}),k(o,{label:"地区"},{default:_((()=>[k(s,{onSelectArea:H})])),_:1}),k(o,null,{default:_((()=>[k(n,{size:"small",type:"primary",icon:"search",onClick:Q},{default:_((()=>[C])),_:1}),k(n,{size:"small",icon:"refresh",onClick:K},{default:_((()=>[I])),_:1})])),_:1})])),_:1},8,["model"])]),g("div",D,[k(c,{ref:"multipleTable",style:{width:"100%"},"tooltip-effect":"dark",data:F.value,"row-key":"ID",onSelectionChange:ee},{default:_((()=>[k(r,{align:"left",label:"日期",width:"180"},{default:_((e=>[U(j(P(m)(e.row.CreatedAt)),1)])),_:1}),k(r,{align:"left",label:"订单号",prop:"orderNum",width:"220"}),k(r,{align:"left",label:"事务详情",prop:"contents",width:"220"}),k(r,{align:"left",label:"联系人",prop:"linkman",width:"120"}),k(r,{align:"left",label:"联系电话",prop:"contactPhone",width:"220"}),k(r,{align:"left",label:"维修地址",prop:"address",width:"220"}),k(r,{align:"left",label:"帮忙时长",prop:"time",width:"120"}),k(r,{align:"left",label:"留言",prop:"remark",width:"120"}),k(r,{align:"left",label:"优惠券",prop:"couponInfo",width:"120"},{default:_((e=>[U(j(""==e.row.couponInfo.title?"无":e.row.couponInfo.title),1)])),_:1}),k(r,{align:"left",label:"优惠金额",prop:"couponAmount",width:"120"}),k(r,{align:"left",label:"赏金",prop:"price",width:"120"}),k(r,{align:"left",label:"接单人",prop:"takeName",width:"120"}),k(r,{align:"left",label:"接单人电话",prop:"takePhone",width:"120"}),k(r,{align:"left",label:"接单时间",prop:"takeTime",width:"120"}),k(r,{align:"left",label:"用户编号",prop:"uid",width:"120"}),k(r,{align:"left",label:"订单状态",prop:"status",width:"120"},{default:_((e=>[U(j(P(p)(e.row.status,W.value)),1)])),_:1}),k(r,{align:"left",label:"完成时间",prop:"doneTime",width:"120"}),k(r,{align:"left",label:"地区",prop:"area",width:"120"}),k(r,{align:"left",label:"按钮组"},{default:_((e=>[k(n,{type:"text",icon:"delete",size:"small",onClick:l=>{return a=e.row,void N.confirm("确定要删除吗?","提示",{confirmButtonText:"确定",cancelButtonText:"取消",type:"warning"}).then((()=>{ae(a)}));var a}},{default:_((()=>[A])),_:2},1032,["onClick"])])),_:1})])),_:1},8,["data"]),g("div",O,[k(v,{layout:"total, sizes, prev, pager, next, jumper","current-page":E.value,"page-size":L.value,"page-sizes":[10,30,50,100],total:J.value,onCurrentChange:X,onSizeChange:R},null,8,["current-page","page-size","total"])])]),k(x,{modelValue:oe.value,"onUpdate:modelValue":l[24]||(l[24]=e=>oe.value=e),"before-close":ue,title:"弹窗操作"},{footer:_((()=>[g("div",S,[k(n,{size:"small",onClick:ue},{default:_((()=>[Y])),_:1}),k(n,{size:"small",type:"primary",onClick:te},{default:_((()=>[M])),_:1})])])),default:_((()=>[k(i,{model:q.value,"label-position":"right","label-width":"80px"},{default:_((()=>[k(o,{label:"订单号:"},{default:_((()=>[k(u,{modelValue:q.value.orderNum,"onUpdate:modelValue":l[7]||(l[7]=e=>q.value.orderNum=e),clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),k(o,{label:"事务详情:"},{default:_((()=>[k(u,{modelValue:q.value.contents,"onUpdate:modelValue":l[8]||(l[8]=e=>q.value.contents=e),clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),k(o,{label:"联系人:"},{default:_((()=>[k(u,{modelValue:q.value.linkman,"onUpdate:modelValue":l[9]||(l[9]=e=>q.value.linkman=e),clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),k(o,{label:"联系电话:"},{default:_((()=>[k(u,{modelValue:q.value.contactPhone,"onUpdate:modelValue":l[10]||(l[10]=e=>q.value.contactPhone=e),clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),k(o,{label:"维修地址:"},{default:_((()=>[k(u,{modelValue:q.value.address,"onUpdate:modelValue":l[11]||(l[11]=e=>q.value.address=e),clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),k(o,{label:"帮忙时长:"},{default:_((()=>[k(u,{modelValue:q.value.time,"onUpdate:modelValue":l[12]||(l[12]=e=>q.value.time=e),clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),k(o,{label:"留言:"},{default:_((()=>[k(u,{modelValue:q.value.remark,"onUpdate:modelValue":l[13]||(l[13]=e=>q.value.remark=e),clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),k(o,{label:"优惠券ID:"},{default:_((()=>[k(u,{modelValue:q.value.couponId,"onUpdate:modelValue":l[14]||(l[14]=e=>q.value.couponId=e),modelModifiers:{number:!0},clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),k(o,{label:"优惠金额:"},{default:_((()=>[k(f,{modelValue:q.value.couponAmount,"onUpdate:modelValue":l[15]||(l[15]=e=>q.value.couponAmount=e),style:{width:"100%"},precision:2,clearable:""},null,8,["modelValue"])])),_:1}),k(o,{label:"赏金:"},{default:_((()=>[k(f,{modelValue:q.value.price,"onUpdate:modelValue":l[16]||(l[16]=e=>q.value.price=e),style:{width:"100%"},precision:2,clearable:""},null,8,["modelValue"])])),_:1}),k(o,{label:"接单人:"},{default:_((()=>[k(u,{modelValue:q.value.takeName,"onUpdate:modelValue":l[17]||(l[17]=e=>q.value.takeName=e),clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),k(o,{label:"接单人电话:"},{default:_((()=>[k(u,{modelValue:q.value.takePhone,"onUpdate:modelValue":l[18]||(l[18]=e=>q.value.takePhone=e),clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),k(o,{label:"接单时间:"},{default:_((()=>[k(u,{modelValue:q.value.takeTime,"onUpdate:modelValue":l[19]||(l[19]=e=>q.value.takeTime=e),clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),k(o,{label:"用户编号:"},{default:_((()=>[k(u,{modelValue:q.value.uid,"onUpdate:modelValue":l[20]||(l[20]=e=>q.value.uid=e),clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),k(o,{label:"订单状态:"},{default:_((()=>[k(d,{modelValue:q.value.status,"onUpdate:modelValue":l[21]||(l[21]=e=>q.value.status=e),placeholder:"请选择",style:{width:"100%"},clearable:""},{default:_((()=>[(h(!0),V(y,null,w(W.value,((e,l)=>(h(),T(t,{key:l,label:e.label,value:e.value},null,8,["label","value"])))),128))])),_:1},8,["modelValue"])])),_:1}),k(o,{label:"完成时间:"},{default:_((()=>[k(u,{modelValue:q.value.doneTime,"onUpdate:modelValue":l[22]||(l[22]=e=>q.value.doneTime=e),clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),k(o,{label:"地区:"},{default:_((()=>[k(u,{modelValue:q.value.area,"onUpdate:modelValue":l[23]||(l[23]=e=>q.value.area=e),clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1})])),_:1},8,["model"])])),_:1},8,["modelValue"])])}}});export{H as default};
