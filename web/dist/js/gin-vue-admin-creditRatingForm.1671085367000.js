/*! 
 Build based on gin-vue-admin 
 Time : 1671085367000 */
var e=(e,a,l)=>new Promise(((i,t)=>{var d=e=>{try{n(l.next(e))}catch(a){t(a)}},u=e=>{try{n(l.throw(e))}catch(a){t(a)}},n=e=>e.done?i(e.value):Promise.resolve(e.value).then(d,u);n((l=l.apply(e,a)).next())}));import{f as a,c as l,u as i}from"./gin-vue-admin-creditRating.16710853670002.js";import"./gin-vue-admin-date.1671085367000.js";import"./gin-vue-admin-dictionary.16710853670002.js";import{U as t,u as d,r as u,i as n,o as r,j as o,q as s,k as m,t as c,A as v,d as p}from"./gin-vue-admin-vendor.1671085367000.js";import"../gva/gin-vue-admin-index.1671085367000.js";import"./gin-vue-admin-sysDictionary.1671085367000.js";const f={class:"gva-form-box"},y=v("保存"),b=v("返回"),g={name:"CreditRating"},h=Object.assign(g,{setup(v){const g=t(),h=d(),x=u(""),V=u({title:"",minGrade:0,maxGrade:0});(()=>{e(this,null,(function*(){if(g.query.id){const e=yield a({ID:g.query.id});0===e.code&&(V.value=e.data.recreditRating,x.value="update")}else x.value="create"}))})();const j=()=>e(this,null,(function*(){let e;switch(x.value){case"create":e=yield l(V.value);break;case"update":e=yield i(V.value);break;default:e=yield l(V.value)}0===e.code&&p({type:"success",message:"创建/更改成功"})})),_=()=>{h.go(-1)};return(e,a)=>{const l=n("el-input"),i=n("el-form-item"),t=n("el-button"),d=n("el-form");return r(),o("div",null,[s("div",f,[m(d,{model:V.value,"label-position":"right","label-width":"80px"},{default:c((()=>[m(i,{label:"标题:"},{default:c((()=>[m(l,{modelValue:V.value.title,"onUpdate:modelValue":a[0]||(a[0]=e=>V.value.title=e),clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),m(i,{label:"最小分数:"},{default:c((()=>[m(l,{modelValue:V.value.minGrade,"onUpdate:modelValue":a[1]||(a[1]=e=>V.value.minGrade=e),modelModifiers:{number:!0},clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),m(i,{label:"最大分数:"},{default:c((()=>[m(l,{modelValue:V.value.maxGrade,"onUpdate:modelValue":a[2]||(a[2]=e=>V.value.maxGrade=e),modelModifiers:{number:!0},clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),m(i,null,{default:c((()=>[m(t,{size:"mini",type:"primary",onClick:j},{default:c((()=>[y])),_:1}),m(t,{size:"mini",type:"primary",onClick:_},{default:c((()=>[b])),_:1})])),_:1})])),_:1},8,["model"])])])}}});export{h as default};
