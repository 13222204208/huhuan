/*! 
 Build based on gin-vue-admin 
 Time : 1671085367000 */
var e=(e,l,a)=>new Promise(((u,o)=>{var d=e=>{try{i(a.next(e))}catch(l){o(l)}},t=e=>{try{i(a.throw(e))}catch(l){o(l)}},i=e=>e.done?u(e.value):Promise.resolve(e.value).then(d,t);i((a=a.apply(e,l)).next())}));import{f as l,c as a,u}from"./gin-vue-admin-releaseRecord.16710853670002.js";import"./gin-vue-admin-date.1671085367000.js";import"./gin-vue-admin-dictionary.16710853670002.js";import{U as o,u as d,r as t,i,o as s,j as r,q as n,k as m,t as c,A as v,d as p}from"./gin-vue-admin-vendor.1671085367000.js";import"../gva/gin-vue-admin-index.1671085367000.js";import"./gin-vue-admin-sysDictionary.1671085367000.js";const f={class:"gva-form-box"},b=v("保存"),y=v("返回"),h={name:"ReleaseRecord"},g=Object.assign(h,{setup(v){const h=o(),g=d(),V=t(""),j=t({uid:0,couponId:0,total:0,status:0});(()=>{e(this,null,(function*(){if(h.query.id){const e=yield l({ID:h.query.id});0===e.code&&(j.value=e.data.rereleaseRecord,V.value="update")}else V.value="create"}))})();const _=()=>e(this,null,(function*(){let e;switch(V.value){case"create":e=yield a(j.value);break;case"update":e=yield u(j.value);break;default:e=yield a(j.value)}0===e.code&&p({type:"success",message:"创建/更改成功"})})),x=()=>{g.go(-1)};return(e,l)=>{const a=i("el-input"),u=i("el-form-item"),o=i("el-button"),d=i("el-form");return s(),r("div",null,[n("div",f,[m(d,{model:j.value,"label-position":"right","label-width":"80px"},{default:c((()=>[m(u,{label:"用户ID:"},{default:c((()=>[m(a,{modelValue:j.value.uid,"onUpdate:modelValue":l[0]||(l[0]=e=>j.value.uid=e),modelModifiers:{number:!0},clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),m(u,{label:"优惠券id:"},{default:c((()=>[m(a,{modelValue:j.value.couponId,"onUpdate:modelValue":l[1]||(l[1]=e=>j.value.couponId=e),modelModifiers:{number:!0},clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),m(u,{label:"数量:"},{default:c((()=>[m(a,{modelValue:j.value.total,"onUpdate:modelValue":l[2]||(l[2]=e=>j.value.total=e),modelModifiers:{number:!0},clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),m(u,{label:"状态:"},{default:c((()=>[m(a,{modelValue:j.value.status,"onUpdate:modelValue":l[3]||(l[3]=e=>j.value.status=e),modelModifiers:{number:!0},clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),m(u,null,{default:c((()=>[m(o,{size:"mini",type:"primary",onClick:_},{default:c((()=>[b])),_:1}),m(o,{size:"mini",type:"primary",onClick:x},{default:c((()=>[y])),_:1})])),_:1})])),_:1},8,["model"])])])}}});export{g as default};