/*! 
 Build based on gin-vue-admin 
 Time : 1671085367000 */
!function(){function t(t,e,n,r,u,a,i){try{var o=t[a](i),l=o.value}catch(c){return void n(c)}o.done?e(l):Promise.resolve(l).then(r,u)}function e(t){return function(t){if(Array.isArray(t))return n(t)}(t)||function(t){if("undefined"!=typeof Symbol&&null!=t[Symbol.iterator]||null!=t["@@iterator"])return Array.from(t)}(t)||function(t,e){if(!t)return;if("string"==typeof t)return n(t,e);var r=Object.prototype.toString.call(t).slice(8,-1);"Object"===r&&t.constructor&&(r=t.constructor.name);if("Map"===r||"Set"===r)return Array.from(t);if("Arguments"===r||/^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(r))return n(t,e)}(t)||function(){throw new TypeError("Invalid attempt to spread non-iterable instance.\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method.")}()}function n(t,e){(null==e||e>t.length)&&(e=t.length);for(var n=0,r=new Array(e);n<e;n++)r[n]=t[n];return r}System.register(["./gin-vue-admin-authority-legacy.16710853670002.js","./gin-vue-admin-warningBar-legacy.1671085367000.js","./gin-vue-admin-vendor-legacy.1671085367000.js","../gva/gin-vue-admin-index-legacy.1671085367000.js"],(function(n){"use strict";var r,u,a,i,o,l,c,f,s,d,y,h,v,m,p;return{setters:[function(t){r=t.s},function(t){u=t.w},function(t){a=t.r,i=t.i,o=t.o,l=t.j,c=t.q,f=t.k,s=t.t,d=t.J,y=t.L,h=t.A,v=t.d,m=t.H,p=t.C},function(){}],execute:function(){var g={class:"clearflex",style:{margin:"18px"}},w=h("确 定"),I=h("全选"),b=h("本角色"),A=h("本角色及子角色"),x={name:"Datas"};n("default",Object.assign(x,{props:{row:{default:function(){return{}},type:Object},authority:{default:function(){return[]},type:Array}},emits:["changeRow"],setup:function(n,x){var j=x.expose,k=x.emit,C=n,R=a([]),S=a(!1),_=function t(e){e&&e.forEach((function(e){var n={};n.authorityId=e.authorityId,n.authorityName=e.authorityName,R.value.push(n),e.children&&e.children.length&&t(e.children)}))},O=a([]);_(C.authority),C.row.dataAuthorityId&&C.row.dataAuthorityId.forEach((function(t){var e=R.value&&R.value.filter((function(e){return e.authorityId===t.authorityId}))&&R.value.filter((function(e){return e.authorityId===t.authorityId}))[0];O.value.push(e)}));var z=function(){O.value=e(R.value),k("changeRow","dataAuthorityId",O.value),S.value=!0},E=function(){O.value=R.value.filter((function(t){return t.authorityId===C.row.authorityId})),k("changeRow","dataAuthorityId",O.value),S.value=!0},N=function(){var t=[];V(C.row,t),O.value=R.value.filter((function(e){return t.indexOf(e.authorityId)>-1})),k("changeRow","dataAuthorityId",O.value),S.value=!0},V=function t(e,n){n.push(e.authorityId),e.children&&e.children.forEach((function(e){t(e,n)}))},P=function(){var e,n=(e=regeneratorRuntime.mark((function t(){return regeneratorRuntime.wrap((function(t){for(;;)switch(t.prev=t.next){case 0:return t.next=2,r(C.row);case 2:0===t.sent.code&&v({type:"success",message:"资源设置成功"});case 4:case"end":return t.stop()}}),t)})),function(){var n=this,r=arguments;return new Promise((function(u,a){var i=e.apply(n,r);function o(e){t(i,u,a,o,l,"next",e)}function l(e){t(i,u,a,o,l,"throw",e)}o(void 0)}))});return function(){return n.apply(this,arguments)}}(),U=function(){k("changeRow","dataAuthorityId",O.value),S.value=!0};return j({enterAndNext:function(){P()},needConfirm:S}),function(t,e){var n=i("el-button"),r=i("el-checkbox"),a=i("el-checkbox-group");return o(),l("div",null,[c("div",g,[f(n,{class:"fl-right",size:"small",type:"primary",onClick:P},{default:s((function(){return[w]})),_:1}),f(n,{class:"fl-left",size:"small",type:"primary",onClick:z},{default:s((function(){return[I]})),_:1}),f(n,{class:"fl-left",size:"small",type:"primary",onClick:E},{default:s((function(){return[b]})),_:1}),f(n,{class:"fl-left",size:"small",type:"primary",onClick:N},{default:s((function(){return[A]})),_:1})]),f(a,{modelValue:O.value,"onUpdate:modelValue":e[0]||(e[0]=function(t){return O.value=t}),onChange:U},{default:s((function(){return[(o(!0),l(d,null,y(R.value,(function(t,e){return o(),m(r,{key:e,label:t},{default:s((function(){return[h(p(t.authorityName),1)]})),_:2},1032,["label"])})),128))]})),_:1},8,["modelValue"]),f(u,{title:"此功能仅用于创建角色和角色的many2many关系表，具体使用还须自己结合表实现业务，详情参考示例代码（客户示例）"})])}}}))}}}))}();
