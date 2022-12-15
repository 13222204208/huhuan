/*! 
 Build based on gin-vue-admin 
 Time : 1671085367000 */
var e=Object.defineProperty,l=Object.getOwnPropertySymbols,a=Object.prototype.hasOwnProperty,o=Object.prototype.propertyIsEnumerable,u=(l,a,o)=>a in l?e(l,a,{enumerable:!0,configurable:!0,writable:!0,value:o}):l[a]=o,t=(e,l,a)=>new Promise(((o,u)=>{var t=e=>{try{n(a.next(e))}catch(l){u(l)}},d=e=>{try{n(a.throw(e))}catch(l){u(l)}},n=e=>e.done?o(e.value):Promise.resolve(e.value).then(t,d);n((a=a.apply(e,l)).next())}));import{g as d,d as n,c as r,u as i}from"./gin-vue-admin-helpIncidentally.16710853670002.js";import{f as m,a as p,g as s}from"./gin-vue-admin-format.1671085367000.js";import{_ as c}from"./gin-vue-admin-index.167108536700024.js";import{r as v,i as b,o as f,j as h,q as V,k as g,t as y,A as _,C as w,x as k,J as U,L as C,e as N,d as P,H as A}from"./gin-vue-admin-vendor.1671085367000.js";import"../gva/gin-vue-admin-index.1671085367000.js";import"./gin-vue-admin-date.1671085367000.js";import"./gin-vue-admin-dictionary.1671085367000.js";import"./gin-vue-admin-dictionary.16710853670002.js";import"./gin-vue-admin-sysDictionary.1671085367000.js";import"./gin-vue-admin-area.16710853670002.js";const j={class:"gva-search-box"},x=_("查询"),T=_("重置"),z={class:"gva-table-box"},I=_("删除"),O={class:"gva-pagination"},S={class:"dialog-footer"},D=_("取 消"),B=_("确 定"),H={name:"HelpIncidentally"},M=Object.assign(H,{setup(e){const H=e=>t(this,null,(function*(){console.log("子组件传的值",e),G.value.area=e})),M=v([]),q=v({orderNum:"",goodsName:"",goodsCount:"",linkman:"",contactPhone:"",fetchAddress:"",receiveAddress:"",time:"",remark:"",couponId:0,couponAmount:0,price:0,takeName:"",takePhone:"",takeTime:"",userNum:"",goodsPrice:0,status:void 0,doneTime:"",city:"",area:"",type:0,startCity:"",arriveCity:"",flightNumber:""}),E=v(1),J=v(0),L=v(10),F=v([]),G=v({}),K=()=>{G.value={}},Q=()=>{E.value=1,L.value=10,X()},R=e=>{L.value=e,X()},W=e=>{E.value=e,X()},X=()=>t(this,null,(function*(){const e=yield d(((e,t)=>{for(var d in t||(t={}))a.call(t,d)&&u(e,d,t[d]);if(l)for(var d of l(t))o.call(t,d)&&u(e,d,t[d]);return e})({page:E.value,pageSize:L.value},G.value));0===e.code&&(F.value=e.data.list,J.value=e.data.total,E.value=e.data.page,L.value=e.data.pageSize)}));X();(()=>{t(this,null,(function*(){M.value=yield s("orderState")}))})();const Y=v([]),Z=e=>{Y.value=e};v(!1);const $=v(""),ee=e=>t(this,null,(function*(){0===(yield n({ID:e.ID})).code&&(P({type:"success",message:"删除成功"}),1===F.value.length&&E.value>1&&E.value--,X())})),le=v(!1),ae=()=>{le.value=!1,q.value={orderNum:"",goodsName:"",goodsCount:"",linkman:"",contactPhone:"",fetchAddress:"",receiveAddress:"",time:"",remark:"",couponId:0,couponAmount:0,price:0,takeName:"",takePhone:"",takeTime:"",userNum:"",goodsPrice:0,status:void 0,doneTime:"",city:"",area:"",type:0,startCity:"",arriveCity:"",flightNumber:""}},oe=()=>t(this,null,(function*(){let e;switch($.value){case"create":e=yield r(q.value);break;case"update":e=yield i(q.value);break;default:e=yield r(q.value)}0===e.code&&(P({type:"success",message:"创建/更改成功"}),ae(),X())}));return(e,l)=>{const a=b("el-input"),o=b("el-form-item"),u=b("el-button"),t=b("el-form"),d=b("el-table-column"),n=b("el-table"),r=b("el-pagination"),i=b("el-input-number"),s=b("el-option"),v=b("el-select"),P=b("el-dialog");return f(),h("div",null,[V("div",j,[g(t,{inline:!0,model:G.value,class:"demo-form-inline"},{default:y((()=>[g(o,{label:"订单号"},{default:y((()=>[g(a,{modelValue:G.value.orderNum,"onUpdate:modelValue":l[0]||(l[0]=e=>G.value.orderNum=e),placeholder:"搜索条件"},null,8,["modelValue"])])),_:1}),g(o,{label:"联系人"},{default:y((()=>[g(a,{modelValue:G.value.linkman,"onUpdate:modelValue":l[1]||(l[1]=e=>G.value.linkman=e),placeholder:"搜索条件"},null,8,["modelValue"])])),_:1}),g(o,{label:"联系电话"},{default:y((()=>[g(a,{modelValue:G.value.contactPhone,"onUpdate:modelValue":l[2]||(l[2]=e=>G.value.contactPhone=e),placeholder:"搜索条件"},null,8,["modelValue"])])),_:1}),g(o,{label:"接单人"},{default:y((()=>[g(a,{modelValue:G.value.takeName,"onUpdate:modelValue":l[3]||(l[3]=e=>G.value.takeName=e),placeholder:"搜索条件"},null,8,["modelValue"])])),_:1}),g(o,{label:"用户编号"},{default:y((()=>[g(a,{modelValue:G.value.userNum,"onUpdate:modelValue":l[4]||(l[4]=e=>G.value.userNum=e),placeholder:"搜索条件"},null,8,["modelValue"])])),_:1}),g(o,{label:"订单状态"},{default:y((()=>[g(a,{modelValue:G.value.status,"onUpdate:modelValue":l[5]||(l[5]=e=>G.value.status=e),placeholder:"搜索条件"},null,8,["modelValue"])])),_:1}),g(o,{label:"地区"},{default:y((()=>[g(c,{onSelectArea:H})])),_:1}),g(o,{label:"类型"},{default:y((()=>[g(a,{modelValue:G.value.type,"onUpdate:modelValue":l[6]||(l[6]=e=>G.value.type=e),placeholder:"搜索条件"},null,8,["modelValue"])])),_:1}),g(o,{label:"出发城市"},{default:y((()=>[g(a,{modelValue:G.value.startCity,"onUpdate:modelValue":l[7]||(l[7]=e=>G.value.startCity=e),placeholder:"搜索条件"},null,8,["modelValue"])])),_:1}),g(o,{label:"抵达城市"},{default:y((()=>[g(a,{modelValue:G.value.arriveCity,"onUpdate:modelValue":l[8]||(l[8]=e=>G.value.arriveCity=e),placeholder:"搜索条件"},null,8,["modelValue"])])),_:1}),g(o,null,{default:y((()=>[g(u,{size:"small",type:"primary",icon:"search",onClick:Q},{default:y((()=>[x])),_:1}),g(u,{size:"small",icon:"refresh",onClick:K},{default:y((()=>[T])),_:1})])),_:1})])),_:1},8,["model"])]),V("div",z,[g(n,{ref:"multipleTable",style:{width:"100%"},"tooltip-effect":"dark",data:F.value,"row-key":"ID",onSelectionChange:Z},{default:y((()=>[g(d,{align:"left",label:"日期",width:"180"},{default:y((e=>[_(w(k(m)(e.row.CreatedAt)),1)])),_:1}),g(d,{align:"left",label:"订单号",prop:"orderNum",width:"220"}),g(d,{align:"left",label:"物品名称",prop:"goodsName",width:"120"}),g(d,{align:"left",label:"物品数量",prop:"goodsCount",width:"120"}),g(d,{align:"left",label:"联系人",prop:"linkman",width:"120"}),g(d,{align:"left",label:"联系电话",prop:"contactPhone",width:"120"}),g(d,{align:"left",label:"取货地址",prop:"fetchAddress",width:"120"}),g(d,{align:"left",label:"收货地址",prop:"receiveAddress",width:"220"}),g(d,{align:"left",label:"出发时间",prop:"time",width:"120"}),g(d,{align:"left",label:"留言",prop:"remark",width:"120"}),g(d,{align:"left",label:"优惠金额",prop:"couponAmount",width:"120"}),g(d,{align:"left",label:"赏金",prop:"price",width:"120"}),g(d,{align:"left",label:"接单人",prop:"takeName",width:"120"}),g(d,{align:"left",label:"接单人电话",prop:"takePhone",width:"120"}),g(d,{align:"left",label:"接单时间",prop:"takeTime",width:"120"}),g(d,{align:"left",label:"用户编号",prop:"userNum",width:"120"}),g(d,{align:"left",label:"商品价值",prop:"goodsPrice",width:"120"}),g(d,{align:"left",label:"订单状态",prop:"status",width:"120"},{default:y((e=>[_(w(k(p)(e.row.status,M.value)),1)])),_:1}),g(d,{align:"left",label:"完成时间",prop:"doneTime",width:"120"}),g(d,{align:"left",label:"城市",prop:"city",width:"120"}),g(d,{align:"left",label:"地区",prop:"area",width:"120"}),g(d,{align:"left",label:"类型",prop:"type",width:"120"}),g(d,{align:"left",label:"出发城市",prop:"startCity",width:"120"}),g(d,{align:"left",label:"抵达城市",prop:"arriveCity",width:"120"}),g(d,{align:"left",label:"航班号",prop:"flightNumber",width:"120"}),g(d,{align:"left",label:"按钮组"},{default:y((e=>[g(u,{type:"text",icon:"delete",size:"small",onClick:l=>{return a=e.row,void N.confirm("确定要删除吗?","提示",{confirmButtonText:"确定",cancelButtonText:"取消",type:"warning"}).then((()=>{ee(a)}));var a}},{default:y((()=>[I])),_:2},1032,["onClick"])])),_:1})])),_:1},8,["data"]),V("div",O,[g(r,{layout:"total, sizes, prev, pager, next, jumper","current-page":E.value,"page-size":L.value,"page-sizes":[10,30,50,100],total:J.value,onCurrentChange:W,onSizeChange:R},null,8,["current-page","page-size","total"])])]),g(P,{modelValue:le.value,"onUpdate:modelValue":l[34]||(l[34]=e=>le.value=e),"before-close":ae,title:"弹窗操作"},{footer:y((()=>[V("div",S,[g(u,{size:"small",onClick:ae},{default:y((()=>[D])),_:1}),g(u,{size:"small",type:"primary",onClick:oe},{default:y((()=>[B])),_:1})])])),default:y((()=>[g(t,{model:q.value,"label-position":"right","label-width":"80px"},{default:y((()=>[g(o,{label:"订单号:"},{default:y((()=>[g(a,{modelValue:q.value.orderNum,"onUpdate:modelValue":l[9]||(l[9]=e=>q.value.orderNum=e),clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),g(o,{label:"物品名称:"},{default:y((()=>[g(a,{modelValue:q.value.goodsName,"onUpdate:modelValue":l[10]||(l[10]=e=>q.value.goodsName=e),clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),g(o,{label:"物品数量:"},{default:y((()=>[g(a,{modelValue:q.value.goodsCount,"onUpdate:modelValue":l[11]||(l[11]=e=>q.value.goodsCount=e),clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),g(o,{label:"联系人:"},{default:y((()=>[g(a,{modelValue:q.value.linkman,"onUpdate:modelValue":l[12]||(l[12]=e=>q.value.linkman=e),clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),g(o,{label:"联系电话:"},{default:y((()=>[g(a,{modelValue:q.value.contactPhone,"onUpdate:modelValue":l[13]||(l[13]=e=>q.value.contactPhone=e),clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),g(o,{label:"取货地址:"},{default:y((()=>[g(a,{modelValue:q.value.fetchAddress,"onUpdate:modelValue":l[14]||(l[14]=e=>q.value.fetchAddress=e),clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),g(o,{label:"收货地址:"},{default:y((()=>[g(a,{modelValue:q.value.receiveAddress,"onUpdate:modelValue":l[15]||(l[15]=e=>q.value.receiveAddress=e),clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),g(o,{label:"出发时间:"},{default:y((()=>[g(a,{modelValue:q.value.time,"onUpdate:modelValue":l[16]||(l[16]=e=>q.value.time=e),clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),g(o,{label:"留言:"},{default:y((()=>[g(a,{modelValue:q.value.remark,"onUpdate:modelValue":l[17]||(l[17]=e=>q.value.remark=e),clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),g(o,{label:"优惠券ID:"},{default:y((()=>[g(a,{modelValue:q.value.couponId,"onUpdate:modelValue":l[18]||(l[18]=e=>q.value.couponId=e),modelModifiers:{number:!0},clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),g(o,{label:"优惠金额:"},{default:y((()=>[g(i,{modelValue:q.value.couponAmount,"onUpdate:modelValue":l[19]||(l[19]=e=>q.value.couponAmount=e),style:{width:"100%"},precision:2,clearable:""},null,8,["modelValue"])])),_:1}),g(o,{label:"赏金:"},{default:y((()=>[g(i,{modelValue:q.value.price,"onUpdate:modelValue":l[20]||(l[20]=e=>q.value.price=e),style:{width:"100%"},precision:2,clearable:""},null,8,["modelValue"])])),_:1}),g(o,{label:"接单人:"},{default:y((()=>[g(a,{modelValue:q.value.takeName,"onUpdate:modelValue":l[21]||(l[21]=e=>q.value.takeName=e),clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),g(o,{label:"接单人电话:"},{default:y((()=>[g(a,{modelValue:q.value.takePhone,"onUpdate:modelValue":l[22]||(l[22]=e=>q.value.takePhone=e),clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),g(o,{label:"接单时间:"},{default:y((()=>[g(a,{modelValue:q.value.takeTime,"onUpdate:modelValue":l[23]||(l[23]=e=>q.value.takeTime=e),clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),g(o,{label:"用户编号:"},{default:y((()=>[g(a,{modelValue:q.value.userNum,"onUpdate:modelValue":l[24]||(l[24]=e=>q.value.userNum=e),clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),g(o,{label:"商品价值:"},{default:y((()=>[g(i,{modelValue:q.value.goodsPrice,"onUpdate:modelValue":l[25]||(l[25]=e=>q.value.goodsPrice=e),style:{width:"100%"},precision:2,clearable:""},null,8,["modelValue"])])),_:1}),g(o,{label:"订单状态:"},{default:y((()=>[g(v,{modelValue:q.value.status,"onUpdate:modelValue":l[26]||(l[26]=e=>q.value.status=e),placeholder:"请选择",style:{width:"100%"},clearable:""},{default:y((()=>[(f(!0),h(U,null,C(M.value,((e,l)=>(f(),A(s,{key:l,label:e.label,value:e.value},null,8,["label","value"])))),128))])),_:1},8,["modelValue"])])),_:1}),g(o,{label:"完成时间:"},{default:y((()=>[g(a,{modelValue:q.value.doneTime,"onUpdate:modelValue":l[27]||(l[27]=e=>q.value.doneTime=e),clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),g(o,{label:"城市:"},{default:y((()=>[g(a,{modelValue:q.value.city,"onUpdate:modelValue":l[28]||(l[28]=e=>q.value.city=e),clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),g(o,{label:"地区:"},{default:y((()=>[g(a,{modelValue:q.value.area,"onUpdate:modelValue":l[29]||(l[29]=e=>q.value.area=e),clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),g(o,{label:"类型:"},{default:y((()=>[g(a,{modelValue:q.value.type,"onUpdate:modelValue":l[30]||(l[30]=e=>q.value.type=e),modelModifiers:{number:!0},clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),g(o,{label:"出发城市:"},{default:y((()=>[g(a,{modelValue:q.value.startCity,"onUpdate:modelValue":l[31]||(l[31]=e=>q.value.startCity=e),clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),g(o,{label:"抵达城市:"},{default:y((()=>[g(a,{modelValue:q.value.arriveCity,"onUpdate:modelValue":l[32]||(l[32]=e=>q.value.arriveCity=e),clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),g(o,{label:"航班号:"},{default:y((()=>[g(a,{modelValue:q.value.flightNumber,"onUpdate:modelValue":l[33]||(l[33]=e=>q.value.flightNumber=e),clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1})])),_:1},8,["model"])])),_:1},8,["modelValue"])])}}});export{M as default};