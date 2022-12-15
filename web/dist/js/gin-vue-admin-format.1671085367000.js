/*! 
 Build based on gin-vue-admin 
 Time : 1671085367000 */
import{f as e}from"./gin-vue-admin-date.1671085367000.js";import{g as r}from"./gin-vue-admin-dictionary.1671085367000.js";const t=e=>null!==e?e?"是":"否":"",a=(r,t)=>{if(null!==r&&""!==r){var a=new Date(r);return e(a,"day"==t?"yyyy-MM-dd":"yyyy-MM-dd hh:mm:ss")}return""},n=(e,r)=>{const t=r&&r.filter((r=>r.value===e));return t&&t[0]&&t[0].label},s=e=>{return t=this,a=null,n=function*(){return yield r(e)},new Promise(((e,r)=>{var s=e=>{try{l(n.next(e))}catch(t){r(t)}},i=e=>{try{l(n.throw(e))}catch(t){r(t)}},l=r=>r.done?e(r.value):Promise.resolve(r.value).then(s,i);l((n=n.apply(t,a)).next())}));var t,a,n};export{n as a,t as b,a as f,s as g};
