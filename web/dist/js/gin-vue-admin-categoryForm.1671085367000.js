/*! 
 Build based on gin-vue-admin 
 Time : 1671085367000 */
var e=(e,a,l)=>new Promise(((t,i)=>{var u=e=>{try{d(l.next(e))}catch(a){i(a)}},o=e=>{try{d(l.throw(e))}catch(a){i(a)}},d=e=>e.done?t(e.value):Promise.resolve(e.value).then(u,o);d((l=l.apply(e,a)).next())}));import{f as a,c as l,u as t}from"./gin-vue-admin-category.16710853670002.js";import"./gin-vue-admin-date.1671085367000.js";import"./gin-vue-admin-dictionary.16710853670002.js";import{U as i,u,r as o,i as d,o as s,j as n,q as r,k as m,t as c,A as v,d as p}from"./gin-vue-admin-vendor.1671085367000.js";import"../gva/gin-vue-admin-index.1671085367000.js";import"./gin-vue-admin-sysDictionary.1671085367000.js";const y={class:"gva-form-box"},f=v("保存"),g=v("返回"),b={name:"Category"},h=Object.assign(b,{setup(v){const b=i(),h=u(),j=o(""),x=o({title:"",pid:0});(()=>{e(this,null,(function*(){if(b.query.id){const e=yield a({ID:b.query.id});0===e.code&&(x.value=e.data.recategory,j.value="update")}else j.value="create"}))})();const V=()=>e(this,null,(function*(){let e;switch(j.value){case"create":e=yield l(x.value);break;case"update":e=yield t(x.value);break;default:e=yield l(x.value)}0===e.code&&p({type:"success",message:"创建/更改成功"})})),_=()=>{h.go(-1)};return(e,a)=>{const l=d("el-input"),t=d("el-form-item"),i=d("el-button"),u=d("el-form");return s(),n("div",null,[r("div",y,[m(u,{model:x.value,"label-position":"right","label-width":"80px"},{default:c((()=>[m(t,{label:"分类名称:"},{default:c((()=>[m(l,{modelValue:x.value.title,"onUpdate:modelValue":a[0]||(a[0]=e=>x.value.title=e),clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),m(t,{label:"父ID:"},{default:c((()=>[m(l,{modelValue:x.value.pid,"onUpdate:modelValue":a[1]||(a[1]=e=>x.value.pid=e),modelModifiers:{number:!0},clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),m(t,null,{default:c((()=>[m(i,{size:"mini",type:"primary",onClick:V},{default:c((()=>[f])),_:1}),m(i,{size:"mini",type:"primary",onClick:_},{default:c((()=>[g])),_:1})])),_:1})])),_:1},8,["model"])])])}}});export{h as default};
