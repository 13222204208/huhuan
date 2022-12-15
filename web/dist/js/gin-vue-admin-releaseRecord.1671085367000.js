/*! 
 Build based on gin-vue-admin 
 Time : 1671085367000 */
var e=Object.defineProperty,l=Object.getOwnPropertySymbols,a=Object.prototype.hasOwnProperty,t=Object.prototype.propertyIsEnumerable,o=(l,a,t)=>a in l?e(l,a,{enumerable:!0,configurable:!0,writable:!0,value:t}):l[a]=t,u=(e,u)=>{for(var i in u||(u={}))a.call(u,i)&&o(e,i,u[i]);if(l)for(var i of l(u))t.call(u,i)&&o(e,i,u[i]);return e},i=(e,l,a)=>new Promise(((t,o)=>{var u=e=>{try{n(a.next(e))}catch(l){o(l)}},i=e=>{try{n(a.throw(e))}catch(l){o(l)}},n=e=>e.done?t(e.value):Promise.resolve(e.value).then(u,i);n((a=a.apply(e,l)).next())}));import{g as n,d,c as s,u as r}from"./gin-vue-admin-releaseRecord.16710853670002.js";import{a as p}from"./gin-vue-admin-minappUser.1671085367000.js";import{a as v}from"./gin-vue-admin-coupon.16710853670002.js";import{f as c}from"./gin-vue-admin-format.1671085367000.js";import{r as m,U as f,i as g,o as h,j as b,q as y,k as w,t as _,A as V,C as k,x as z,J as C,L as j,_ as x,d as I,e as U,H as S}from"./gin-vue-admin-vendor.1671085367000.js";import"../gva/gin-vue-admin-index.1671085367000.js";import"./gin-vue-admin-date.1671085367000.js";import"./gin-vue-admin-dictionary.1671085367000.js";import"./gin-vue-admin-dictionary.16710853670002.js";import"./gin-vue-admin-sysDictionary.1671085367000.js";const D={class:"gva-table-box"},O={class:"gva-btn-list"},P=V("新增"),T={style:{color:"green"}},R={style:{color:"red"}},B=V("删除"),q={class:"gva-pagination"},A=V(" 此处受总数限制，如果剩余张数不足以发放给选定会员数量，则无法发放 "),E=V("选择会员"),H={style:{color:"blue","margin-left":"5px"}},J={class:"dialog-footer"},L=V("取 消"),M=V("确 定"),F={class:"gva-search-box"},G=V("查询"),K=V("重置"),N={class:"gva-btn-list"},Q=y("p",null,"确定要删除吗？",-1),W={style:{"text-align":"right","margin-top":"8px"}},X=V("取消"),Y=V("确定选择"),Z=["src"],$={class:"gva-pagination"},ee={name:"ReleaseRecord"},le=Object.assign(ee,{setup(e){const l=m({uid:0,couponId:null,total:1,status:0}),a=m(1),t=m(0),o=m(10),ee=m([]),le=m({}),ae=m([]),te=m(1),oe=m(0),ue=m(10),ie=m([]),ne=m({}),de=m(!1),se=f(),re=m([]);(()=>{i(this,null,(function*(){se.params.id&&(le.value.uid=se.params.id)}))})();const pe=()=>{de.value=!0},ve=()=>{te.value=1,ue.value=10,fe()},ce=e=>{ue.value=e,fe()},me=e=>{te.value=e,fe()},fe=()=>i(this,null,(function*(){const e=yield p(u({page:a.value,pageSize:o.value},ne.value));0===e.code&&(ie.value=e.data.list,oe.value=e.data.total,te.value=e.data.page,ue.value=e.data.pageSize)}));fe();const ge=()=>{ne.value={}},he=e=>{o.value=e,ye()},be=e=>{a.value=e,ye()};(()=>{i(this,null,(function*(){const e=yield v();0===e.code&&(ae.value=e.data.allCoupons)}))})();const ye=()=>i(this,null,(function*(){const e=yield n(u({page:a.value,pageSize:o.value},le.value));0===e.code&&(ee.value=e.data.list,t.value=e.data.total,a.value=e.data.page,o.value=e.data.pageSize)}));ye();(()=>{i(this,null,(function*(){}))})();const we=m([]),_e=e=>{we.value=e},Ve=m([]),ke=e=>{Ve.value=e,console.log("测试",Ve.value);const l=[];0!==Ve.value.length?(Ve.value&&Ve.value.map((e=>{l.push(e.ID)})),console.log("选中",l),re.value=l,console.log("选中的会员",re)):I({type:"warning",message:"请选择要删除的数据"})},ze=m(!1),Ce=m(""),je=e=>i(this,null,(function*(){0===(yield d({ID:e.ID})).code&&(I({type:"success",message:"删除成功"}),1===ee.value.length&&a.value>1&&a.value--,ye())})),xe=m(!1),Ie=()=>{Ce.value="create",xe.value=!0},Ue=()=>{de.value=!1},Se=()=>{de.value=!1,console.log("测试",re)},De=()=>{xe.value=!1,l.value={uid:0,couponId:null,total:1,status:0}},Oe=()=>i(this,null,(function*(){if(l.value.couponId<1)return void I.error("请选择优惠券");if(re.value.length<1)return void I.error("请选择会员");let e;switch(l.value.uid=re.value,Ce.value){case"create":e=yield s(l.value);break;case"update":e=yield r(l.value);break;default:e=yield s(l.value)}0===e.code&&(I({type:"success",message:"创建/更改成功"}),De(),ye())}));return(e,u)=>{const i=g("el-button"),n=g("el-table-column"),d=g("el-table"),s=g("el-pagination"),r=g("el-option"),p=g("el-select"),v=g("el-form-item"),m=g("el-input"),f=g("el-form"),I=g("el-dialog"),le=g("el-popover");return h(),b("div",null,[y("div",D,[y("div",O,[w(i,{size:"small",type:"primary",icon:"plus",onClick:Ie},{default:_((()=>[P])),_:1})]),w(d,{ref:"multipleTable",style:{width:"100%"},"tooltip-effect":"dark",data:ee.value,"row-key":"ID",onSelectionChange:_e},{default:_((()=>[w(n,{align:"left",label:"有效期",width:"230"},{default:_((e=>[V(k(z(c)(e.row.start,"day"))+" 到 "+k(z(c)(e.row.over,"day")),1)])),_:1}),w(n,{align:"left",label:"用户",prop:"uid",width:"120"},{default:_((e=>[V(k(e.row.UserInfo.nickname),1)])),_:1}),w(n,{align:"left",label:"优惠券",prop:"title",width:"120"}),w(n,{align:"left",label:"使用条件",width:"180"},{default:_((e=>[V(" 满"+k(e.row.condition)+"可用 ",1)])),_:1}),w(n,{align:"left",label:"数量",prop:"total",width:"120"}),w(n,{align:"left",label:"状态",prop:"status",width:"120"},{default:_((e=>[y("span",T,k(0==e.row.status?"未使用":""),1),y("span",R,k(1==e.row.status?"已使用":""),1)])),_:1}),w(n,{align:"left",label:"按钮组"},{default:_((e=>[w(i,{type:"text",icon:"delete",size:"small",onClick:l=>{return a=e.row,void U.confirm("确定要删除吗?","提示",{confirmButtonText:"确定",cancelButtonText:"取消",type:"warning"}).then((()=>{je(a)}));var a}},{default:_((()=>[B])),_:2},1032,["onClick"])])),_:1})])),_:1},8,["data"]),y("div",q,[w(s,{layout:"total, sizes, prev, pager, next, jumper","current-page":a.value,"page-size":o.value,"page-sizes":[10,30,50,100],total:t.value,onCurrentChange:be,onSizeChange:he},null,8,["current-page","page-size","total"])])]),w(I,{modelValue:xe.value,"onUpdate:modelValue":u[2]||(u[2]=e=>xe.value=e),"before-close":De,title:"弹窗操作"},{footer:_((()=>[y("div",J,[w(i,{size:"small",onClick:De},{default:_((()=>[L])),_:1}),w(i,{size:"small",type:"primary",onClick:Oe},{default:_((()=>[M])),_:1})])])),default:_((()=>[w(f,{model:l.value,"label-position":"right","label-width":"110px"},{default:_((()=>[w(v,{label:"选择优惠券:"},{default:_((()=>[w(p,{modelValue:l.value.couponId,"onUpdate:modelValue":u[0]||(u[0]=e=>l.value.couponId=e),placeholder:"请选择优惠券",style:{width:"58rem"}},{default:_((()=>[(h(!0),b(C,null,j(ae.value,(e=>(h(),S(r,{key:e.ID,value:e.ID,label:e.title},null,8,["value","label"])))),128))])),_:1},8,["modelValue"])])),_:1}),w(v,{label:"每人发送张数:"},{default:_((()=>[w(m,{modelValue:l.value.total,"onUpdate:modelValue":u[1]||(u[1]=e=>l.value.total=e),modelModifiers:{number:!0},clearable:"",placeholder:"请输入"},null,8,["modelValue"]),A])),_:1}),w(v,{label:"关联会员:"},{default:_((()=>[w(i,{type:"primary",icon:z(x),onClick:pe},{default:_((()=>[E])),_:1},8,["icon"]),y("span",H," 已选中 "+k(re.value.length)+" 人",1)])),_:1})])),_:1},8,["model"])])),_:1},8,["modelValue"]),w(I,{modelValue:de.value,"onUpdate:modelValue":u[9]||(u[9]=e=>de.value=e),"before-close":Ue,title:"选择用户"},{default:_((()=>[y("div",F,[w(f,{inline:!0,model:ne.value,class:"demo-form-inline"},{default:_((()=>[w(v,{label:"昵称"},{default:_((()=>[w(m,{modelValue:ne.value.nickname,"onUpdate:modelValue":u[3]||(u[3]=e=>ne.value.nickname=e),placeholder:"搜索条件"},null,8,["modelValue"])])),_:1}),w(v,{label:"手机号"},{default:_((()=>[w(m,{modelValue:ne.value.phone,"onUpdate:modelValue":u[4]||(u[4]=e=>ne.value.phone=e),placeholder:"搜索条件"},null,8,["modelValue"])])),_:1}),w(v,{label:"编号"},{default:_((()=>[w(m,{modelValue:ne.value.uid,"onUpdate:modelValue":u[5]||(u[5]=e=>ne.value.uid=e),placeholder:"搜索条件"},null,8,["modelValue"])])),_:1}),w(v,{label:"地区"},{default:_((()=>[w(m,{modelValue:ne.value.area,"onUpdate:modelValue":u[6]||(u[6]=e=>ne.value.area=e),placeholder:"搜索条件"},null,8,["modelValue"])])),_:1}),w(v,null,{default:_((()=>[w(i,{size:"small",type:"primary",icon:"search",onClick:ve},{default:_((()=>[G])),_:1}),w(i,{size:"small",icon:"refresh",onClick:ge},{default:_((()=>[K])),_:1})])),_:1})])),_:1},8,["model"])]),y("div",N,[w(le,{visible:ze.value,"onUpdate:visible":u[8]||(u[8]=e=>ze.value=e),placement:"top",width:"160"},{reference:_((()=>[w(i,{size:"small",style:{"margin-left":"10px"},disabled:!Ve.value.length,onClick:Se},{default:_((()=>[Y])),_:1},8,["disabled"])])),default:_((()=>[Q,y("div",W,[w(i,{size:"small",type:"text",onClick:u[7]||(u[7]=e=>ze.value=!1)},{default:_((()=>[X])),_:1})])])),_:1},8,["visible"])]),w(d,{ref:"multipleTable",style:{width:"100%"},"tooltip-effect":"dark",data:ie.value,"row-key":"ID",onSelectionChange:ke},{default:_((()=>[w(n,{type:"selection",width:"55"}),w(n,{align:"left",label:"头像"},{default:_((e=>[y("img",{src:e.row.avatar,width:"60"},null,8,Z)])),_:1}),w(n,{align:"left",label:"昵称",prop:"nickname",width:"120"}),w(n,{align:"left",label:"手机号",prop:"phone",width:"120"}),w(n,{align:"left",label:"地区",prop:"area_id",width:"120"}),w(n,{align:"left",label:"证件号",prop:"number",width:"120"})])),_:1},8,["data"]),y("div",$,[w(s,{layout:"total, sizes, prev, pager, next, jumper","current-page":te.value,"page-size":ue.value,"page-sizes":[10,30,50,100],total:oe.value,onCurrentChange:me,onSizeChange:ce},null,8,["current-page","page-size","total"])])])),_:1},8,["modelValue"])])}}});export{le as default};