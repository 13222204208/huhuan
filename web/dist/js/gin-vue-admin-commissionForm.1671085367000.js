/*! 
 Build based on gin-vue-admin 
 Time : 1671085367000 */
var e=(e,l,a)=>new Promise(((i,o)=>{var u=e=>{try{n(a.next(e))}catch(l){o(l)}},t=e=>{try{n(a.throw(e))}catch(l){o(l)}},n=e=>e.done?i(e.value):Promise.resolve(e.value).then(u,t);n((a=a.apply(e,l)).next())}));import{f as l,c as a,u as i}from"./gin-vue-admin-commission.16710853670002.js";import{g as o}from"./gin-vue-admin-format.1671085367000.js";import{U as u,u as t,r as n,i as s,o as d,j as m,q as r,k as c,t as v,J as p,L as f,A as y,d as b,H as g}from"./gin-vue-admin-vendor.1671085367000.js";import"../gva/gin-vue-admin-index.1671085367000.js";import"./gin-vue-admin-date.1671085367000.js";import"./gin-vue-admin-dictionary.1671085367000.js";import"./gin-vue-admin-dictionary.16710853670002.js";import"./gin-vue-admin-sysDictionary.1671085367000.js";const h={class:"gva-form-box"},V=y("保存"),j=y("返回"),_={name:"Commission"},k=Object.assign(_,{setup(y){const _=u(),k=t(),x=n(""),U=n([]),w=n({title:"",money:0,genre:void 0,commission:0});(()=>{e(this,null,(function*(){if(_.query.id){const e=yield l({ID:_.query.id});0===e.code&&(w.value=e.data.recommission,x.value="update")}else x.value="create";U.value=yield o("genre")}))})();const q=()=>e(this,null,(function*(){let e;switch(x.value){case"create":e=yield a(w.value);break;case"update":e=yield i(w.value);break;default:e=yield a(w.value)}0===e.code&&b({type:"success",message:"创建/更改成功"})})),C=()=>{k.go(-1)};return(e,l)=>{const a=s("el-input"),i=s("el-form-item"),o=s("el-option"),u=s("el-select"),t=s("el-button"),n=s("el-form");return d(),m("div",null,[r("div",h,[c(n,{model:w.value,"label-position":"right","label-width":"80px"},{default:v((()=>[c(i,{label:"名称:"},{default:v((()=>[c(a,{modelValue:w.value.title,"onUpdate:modelValue":l[0]||(l[0]=e=>w.value.title=e),clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),c(i,{label:"小于:"},{default:v((()=>[c(a,{modelValue:w.value.money,"onUpdate:modelValue":l[1]||(l[1]=e=>w.value.money=e),modelModifiers:{number:!0},clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),c(i,{label:"类型:"},{default:v((()=>[c(u,{modelValue:w.value.genre,"onUpdate:modelValue":l[2]||(l[2]=e=>w.value.genre=e),placeholder:"请选择",clearable:""},{default:v((()=>[(d(!0),m(p,null,f(U.value,((e,l)=>(d(),g(o,{key:l,label:e.label,value:e.value},null,8,["label","value"])))),128))])),_:1},8,["modelValue"])])),_:1}),c(i,{label:"佣金:"},{default:v((()=>[c(a,{modelValue:w.value.commission,"onUpdate:modelValue":l[3]||(l[3]=e=>w.value.commission=e),modelModifiers:{number:!0},clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),c(i,null,{default:v((()=>[c(t,{size:"mini",type:"primary",onClick:q},{default:v((()=>[V])),_:1}),c(t,{size:"mini",type:"primary",onClick:C},{default:v((()=>[j])),_:1})])),_:1})])),_:1},8,["model"])])])}}});export{k as default};