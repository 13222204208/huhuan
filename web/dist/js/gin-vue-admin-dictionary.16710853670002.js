/*! 
 Build based on gin-vue-admin 
 Time : 1671085367000 */
var e=Object.defineProperty,r=Object.getOwnPropertySymbols,t=Object.prototype.hasOwnProperty,a=Object.prototype.propertyIsEnumerable,i=(r,t,a)=>t in r?e(r,t,{enumerable:!0,configurable:!0,writable:!0,value:a}):r[t]=a,n=(e,n)=>{for(var o in n||(n={}))t.call(n,o)&&i(e,o,n[o]);if(r)for(var o of r(n))a.call(n,o)&&i(e,o,n[o]);return e};import{f as o}from"./gin-vue-admin-sysDictionary.1671085367000.js";import{f as s,r as l}from"./gin-vue-admin-vendor.1671085367000.js";const y=s("dictionary",(()=>{const e=l({}),r=r=>{e.value=n(n({},e.value),r)};return{dictionaryMap:e,setDictionaryMap:r,getDictionary:t=>{return a=this,i=null,n=function*(){if(e.value[t]&&e.value[t].length)return e.value[t];{const a=yield o({type:t});if(0===a.code){const i={},n=[];return a.data.resysDictionary.sysDictionaryDetails&&a.data.resysDictionary.sysDictionaryDetails.forEach((e=>{n.push({label:e.label,value:e.value})})),i[a.data.resysDictionary.type]=n,r(i),e.value[t]}}},new Promise(((e,r)=>{var t=e=>{try{s(n.next(e))}catch(t){r(t)}},o=e=>{try{s(n.throw(e))}catch(t){r(t)}},s=r=>r.done?e(r.value):Promise.resolve(r.value).then(t,o);s((n=n.apply(a,i)).next())}));var a,i,n}}}));export{y as u};
