/*! 
 Build based on gin-vue-admin 
 Time : 1671085367000 */
var a=Object.defineProperty,e=Object.getOwnPropertySymbols,t=Object.prototype.hasOwnProperty,s=Object.prototype.propertyIsEnumerable,r=(e,t,s)=>t in e?a(e,t,{enumerable:!0,configurable:!0,writable:!0,value:s}):e[t]=s;import{g as n}from"./gin-vue-admin-minappUser.1671085367000.js";import l from"./gin-vue-admin-echartsLine.1671085367000.js";import o from"./gin-vue-admin-dashboardTable.1671085367000.js";import{u as c,r as i,i as d,M as g,o as u,j as v,q as m,N as p,H as b,t as f,k as h,A as y,C as x,J as _,L as j,F as w,G as S,O as k,I as D}from"./gin-vue-admin-vendor.1671085367000.js";import{_ as M}from"../gva/gin-vue-admin-index.1671085367000.js";import"./gin-vue-admin-github.1671085367000.js";import"./gin-vue-admin-date.1671085367000.js";const O=a=>(w("data-v-3baf0d73"),a=a(),S(),a),Y={class:"page"},q={class:"gva-card-box"},C={class:"gva-card gva-top-card"},P={class:"gva-top-card-left"},I=O((()=>m("div",{class:"gva-top-card-left-title"},"早安，管理员，请开始一天的工作吧",-1))),U={class:"gva-top-card-left-rows"},F={class:"flex-center"},H=O((()=>m("img",{src:"./assets/dashboard.70e55b71.png",class:"gva-top-card-right",alt:""},null,-1))),E={class:"gva-card-box"},L=O((()=>m("div",{class:"card-header"},[m("span",null,"快捷入口")],-1))),T={class:"quick-entrance-item"},A={class:"gva-card-box"},G={class:"gva-card"},J=O((()=>m("div",{class:"card-header"},[m("span",null,"数据统计")],-1))),N={class:"echart-box"},R={name:"Dashboard"};var z=M(Object.assign(R,{setup(a){const w=c();i("addUserInfo"),i("");const S=()=>{w.push({name:"minappUser"})};var M=new Date,O=(new Date).getTime()-6048e5,R=new Date(O);R.getMonth(),R.getDate(),R=R.getFullYear()+"-"+(R.getMonth()>9?R.getMonth()+1:"0"+(R.getMonth()+1))+"-"+(R.getDate()>9?R.getDate():"0"+R.getDate()),console.log("开始时间 ",R);const z=i({start:R,end:function(a,e){let t;const s={"Y+":e.getFullYear().toString(),"m+":(e.getMonth()+1).toString(),"d+":e.getDate().toString(),"H+":e.getHours().toString(),"M+":e.getMinutes().toString(),"S+":e.getSeconds().toString()};for(let r in s)t=new RegExp("("+r+")").exec(a),t&&(a=a.replace(t[1],1==t[1].length?s[r]:s[r].padStart(t[1].length,"0")));return a}("YYYY-mm-dd",M),way:""}),B=i({});(()=>{return a=this,l=null,o=function*(){const a=yield n(((a,n)=>{for(var l in n||(n={}))t.call(n,l)&&r(a,l,n[l]);if(e)for(var l of e(n))s.call(n,l)&&r(a,l,n[l]);return a})({},z.value));0===a.code&&(B.value=a.data.data)},new Promise(((e,t)=>{var s=a=>{try{n(o.next(a))}catch(e){t(e)}},r=a=>{try{n(o.throw(a))}catch(e){t(e)}},n=a=>a.done?e(a.value):Promise.resolve(a.value).then(s,r);n((o=o.apply(a,l)).next())}));var a,l,o})();const K=i([{label:"员工管理",icon:"monitor",name:"user",color:"#ff9c6e",bg:"rgba(255, 156, 110,.3)"},{label:"角色管理",icon:"setting",name:"authority",color:"#69c0ff",bg:"rgba(105, 192, 255,.3)"},{label:"小程序用户",icon:"avatar",name:"minappUser",color:"#b37feb",bg:"rgba(179, 127, 235,.3)"},{label:"小程序设置",icon:"setting",name:"minappSet",color:"#ffd666",bg:"rgba(255, 214, 102,.3)"},{label:"优惠券管理",icon:"setting",name:"coupon",color:"#ff85c0",bg:"rgba(255, 133, 192,.3)"},{label:"地区管理",icon:"setting",name:"area",color:"#5cdbd3",bg:"rgba(92, 219, 211,.3)"}]);return(a,e)=>{const t=d("sort"),s=d("el-icon"),r=d("el-col"),n=d("avatar"),c=d("el-row"),i=d("el-card"),M=g("auth");return u(),v("div",Y,[m("div",q,[m("div",C,[m("div",P,[I,m("div",U,[p((u(),b(c,null,{default:f((()=>[h(r,{span:8,xs:24,sm:8},{default:f((()=>[m("div",F,[h(s,{class:"dasboard-icon"},{default:f((()=>[h(t)])),_:1}),y(" 今日流量 ("+x(B.value.active)+") ",1)])])),_:1}),h(r,{span:8,xs:24,sm:8},{default:f((()=>[m("div",{class:"flex-center",style:{cursor:"pointer"},onClick:S},[h(s,{class:"dasboard-icon"},{default:f((()=>[h(n)])),_:1}),y(" 总用户数 ("+x(B.value.all)+") ",1)])])),_:1})])),_:1})),[[M,888]])])]),H])]),m("div",E,[h(i,{class:"gva-card quick-entrance"},{header:f((()=>[L])),default:f((()=>[h(c,{gutter:20},{default:f((()=>[(u(!0),v(_,null,j(K.value,((a,e)=>(u(),b(r,{key:e,span:4,xs:8,class:"quick-entrance-items",onClick:e=>{return t=a.name,void w.push({name:t});var t}},{default:f((()=>[m("div",T,[m("div",{class:"quick-entrance-item-icon",style:k({backgroundColor:a.bg})},[h(s,null,{default:f((()=>[(u(),b(D(a.icon),{style:k({color:a.color})},null,8,["style"]))])),_:2},1024)],4),m("p",null,x(a.label),1)])])),_:2},1032,["onClick"])))),128))])),_:1})])),_:1})]),m("div",A,[m("div",G,[J,m("div",N,[h(c,{gutter:20},{default:f((()=>[h(r,{xs:24,sm:18},{default:f((()=>[h(l)])),_:1}),h(r,{xs:24,sm:6},{default:f((()=>[h(o)])),_:1})])),_:1})])])])])}}}),[["__scopeId","data-v-3baf0d73"]]);export{z as default};
