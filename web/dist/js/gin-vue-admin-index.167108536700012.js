/*! 
 Build based on gin-vue-admin 
 Time : 1671085367000 */
import{_ as e,u as a}from"../gva/gin-vue-admin-index.1671085367000.js";import{r as s,i as t,o as l,j as i,k as n,t as c,q as o,x as d,H as r,v as u,F as v,G as m}from"./gin-vue-admin-vendor.1671085367000.js";const g=e=>(v("data-v-174131b1"),e=e(),m(),e),p={class:"setting_body"},k={class:"setting_card"},b={class:"setting_content"},y={class:"theme-box"},h={class:"item-top"},f=g((()=>o("img",{src:"https://gw.alipayobjects.com/zos/antfincdn/NQ%24zoisaD2/jpRkZQMyYRryryPNtyIC.svg"},null,-1))),_=g((()=>o("p",null," 简约白 ",-1))),j={class:"item-top"},w=g((()=>o("img",{src:"https://gw.alipayobjects.com/zos/antfincdn/XwFOFbLkSM/LCkqqYNmvBEbokSDscrm.svg"},null,-1))),C=g((()=>o("p",null," 商务黑 ",-1))),S={name:"Setting"};var x=e(Object.assign(S,{setup(e){const v=s(!1),m=s("rtl"),g=a(),S=()=>{v.value=!1},x=()=>{v.value=!0},M=e=>{null!==e?g.changeSideMode(e):g.changeSideMode("dark")};return(e,a)=>{const s=t("el-button"),q=t("check"),z=t("el-icon"),F=t("el-drawer");return l(),i("div",null,[n(s,{type:"primary",class:"drawer-container",icon:"setting",onClick:x}),n(F,{modelValue:v.value,"onUpdate:modelValue":a[2]||(a[2]=e=>v.value=e),title:"系统配置",direction:m.value,"before-close":S},{default:c((()=>[o("div",p,[o("div",k,[o("div",b,[o("div",y,[o("div",{class:"item",onClick:a[0]||(a[0]=e=>M("light"))},[o("div",h,["light"===d(g).mode?(l(),r(z,{key:0,class:"check"},{default:c((()=>[n(q)])),_:1})):u("",!0),f]),_]),o("div",{class:"item",onClick:a[1]||(a[1]=e=>M("dark"))},[o("div",j,["dark"===d(g).mode?(l(),r(z,{key:0,class:"check"},{default:c((()=>[n(q)])),_:1})):u("",!0),w]),C])])])])])])),_:1},8,["modelValue","direction"])])}}}),[["__scopeId","data-v-174131b1"]]);export{x as default};