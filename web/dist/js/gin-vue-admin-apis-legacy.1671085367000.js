/*! 
 Build based on gin-vue-admin 
 Time : 1671085367000 */
!function(){function e(e,t,n,r,a,u,i){try{var o=e[u](i),c=o.value}catch(s){return void n(s)}o.done?t(c):Promise.resolve(c).then(r,a)}function t(t){return function(){var n=this,r=arguments;return new Promise((function(a,u){var i=t.apply(n,r);function o(t){e(i,a,u,o,c,"next",t)}function c(t){e(i,a,u,o,c,"throw",t)}o(void 0)}))}}System.register(["./gin-vue-admin-api-legacy.16710853670002.js","../gva/gin-vue-admin-index-legacy.1671085367000.js","./gin-vue-admin-vendor-legacy.1671085367000.js"],(function(e){"use strict";var n,r,a,u,i,o,c,s,l,p,d;return{setters:[function(e){n=e.e},function(e){r=e.s},function(e){a=e.r,u=e.i,i=e.o,o=e.j,c=e.q,s=e.k,l=e.t,p=e.A,d=e.d}],execute:function(){var f={class:"clearflex"},h=p("确 定"),v={name:"Apis"};e("default",Object.assign(v,{props:{row:{default:function(){return{}},type:Object}},setup:function(e,p){var v=p.expose,y=e,m=a({children:"children",label:"description"}),g=a([]),b=a([]),k=a("");(function(){var e=t(regeneratorRuntime.mark((function e(){var t,a,u;return regeneratorRuntime.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,n();case 2:return t=e.sent,a=t.data.apis,g.value=j(a),e.next=7,i={authorityId:y.row.authorityId},r({url:"/casbin/getPolicyPathByAuthorityId",method:"post",data:i});case 7:u=e.sent,k.value=y.row.authorityId,b.value=[],u.data.paths&&u.data.paths.forEach((function(e){b.value.push("p:"+e.path+"m:"+e.method)}));case 11:case"end":return e.stop()}var i}),e)})));return function(){return e.apply(this,arguments)}})()();var x=a(!1),w=function(){x.value=!0},j=function(e){var t={};e&&e.forEach((function(e){e.onlyId="p:"+e.path+"m:"+e.method,Object.prototype.hasOwnProperty.call(t,e.apiGroup)?t[e.apiGroup].push(e):Object.assign(t,function(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}({},e.apiGroup,[e]))}));var n=[];for(var r in t){var a={ID:r,description:r+"组",children:t[r]};n.push(a)}return n},I=a(null),O=function(){var e=t(regeneratorRuntime.mark((function e(){var t,n;return regeneratorRuntime.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return t=I.value.getCheckedNodes(!0),n=[],t&&t.forEach((function(e){var t={path:e.path,method:e.method};n.push(t)})),e.next=5,a={authorityId:k.value,casbinInfos:n},r({url:"/casbin/updateCasbin",method:"post",data:a});case 5:0===e.sent.code&&d({type:"success",message:"api设置成功"});case 7:case"end":return e.stop()}var a}),e)})));return function(){return e.apply(this,arguments)}}();return v({needConfirm:x,enterAndNext:function(){O()}}),function(e,t){var n=u("el-button"),r=u("el-tree");return i(),o("div",null,[c("div",f,[s(n,{class:"fl-right",size:"small",type:"primary",onClick:O},{default:l((function(){return[h]})),_:1})]),s(r,{ref_key:"apiTree",ref:I,data:g.value,"default-checked-keys":b.value,props:m.value,"default-expand-all":"","highlight-current":"","node-key":"onlyId","show-checkbox":"",onCheck:w},null,8,["data","default-checked-keys","props"])])}}}))}}}))}();
