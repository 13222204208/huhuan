/*! 
 Build based on gin-vue-admin 
 Time : 1671085367000 */
var e=(e,a,t)=>new Promise(((l,n)=>{var i=e=>{try{o(t.next(e))}catch(a){n(a)}},s=e=>{try{o(t.throw(e))}catch(a){n(a)}},o=e=>e.done?l(e.value):Promise.resolve(e.value).then(i,s);o((t=t.apply(e,a)).next())}));import{U as a,g as t,d as l}from"./gin-vue-admin-image.1671085367000.js";import{_ as n,u as i}from"../gva/gin-vue-admin-index.1671085367000.js";import{C as s}from"./gin-vue-admin-index.167108536700011.js";import{f as o}from"./gin-vue-admin-format.1671085367000.js";import{r,i as u,M as d,N as p,o as c,j as g,q as m,k as v,t as f,x as h,C as y,A as w,e as _,d as x}from"./gin-vue-admin-vendor.1671085367000.js";/* empty css                                                                             */import"./gin-vue-admin-date.1671085367000.js";import"./gin-vue-admin-dictionary.1671085367000.js";import"./gin-vue-admin-dictionary.16710853670002.js";import"./gin-vue-admin-sysDictionary.1671085367000.js";const b=(e,a)=>{var t=new Image;t.setAttribute("crossOrigin","anonymous"),t.onload=function(){var e=document.createElement("canvas");e.width=t.width,e.height=t.height,e.getContext("2d").drawImage(t,0,0,t.width,t.height);var l=e.toDataURL("image/png"),n=document.createElement("a"),i=new MouseEvent("click");n.download=a||"photo",n.href=l,n.dispatchEvent(i)},t.src=e};const j={class:"gva-table-box"},z={class:"gva-btn-list"},C=w("普通上传"),U=w("下载"),k=w("删除"),A={class:"gva-pagination"},O={name:"Upload"};var E=n(Object.assign(O,{setup(n){const O=r("/api"),E=i(),S=r(""),B=r(1),D=r(0),I=r(10),M=r([]),P=e=>{I.value=e,q()},T=e=>{B.value=e,q()},q=()=>e(this,null,(function*(){const e=yield t({page:B.value,pageSize:I.value});0===e.code&&(M.value=e.data.list,D.value=e.data.total,B.value=e.data.page,I.value=e.data.pageSize)}));q();const K=a=>e(this,null,(function*(){_.confirm("此操作将永久文件, 是否继续?","提示",{confirmButtonText:"确定",cancelButtonText:"取消",type:"warning"}).then((()=>e(this,null,(function*(){0===(yield l(a)).code&&(x({type:"success",message:"删除成功!"}),1===M.value.length&&B.value>1&&B.value--,q())})))).catch((()=>{x({type:"info",message:"已取消删除"})}))})),L=r(!1),N=e=>{L.value=!0;const a="image/jpeg"===e.type,t="image/png"===e.type,l=e.size/1024/1024<.5;return a||t||(x.error("上传图片只能是 jpg或png 格式!"),L.value=!1),l||(x.error("未压缩未见上传图片大小不能超过 500KB，请使用压缩上传"),L.value=!1),(t||a)&&l},R=e=>{L.value=!1,0===e.code?(x({type:"success",message:"上传成功"}),0===e.code&&q()):x({type:"warning",message:e.msg})},$=()=>{x({type:"error",message:"上传失败"}),L.value=!1};return(e,t)=>{const l=u("el-button"),n=u("el-upload"),i=u("el-table-column"),r=u("el-tag"),_=u("el-table"),x=u("el-pagination"),F=d("loading");return p((c(),g("div",null,[m("div",j,[m("div",z,[v(n,{action:`${O.value}/fileUploadAndDownload/upload`,"before-upload":N,headers:{"x-token":h(E).token},"on-error":$,"on-success":R,"show-file-list":!1,class:"upload-btn"},{default:f((()=>[v(l,{size:"small",type:"primary"},{default:f((()=>[C])),_:1})])),_:1},8,["action","headers"]),v(a,{imageUrl:S.value,"onUpdate:imageUrl":t[0]||(t[0]=e=>S.value=e),"file-size":512,"max-w-h":1080,class:"upload-btn",onOnSuccess:q},null,8,["imageUrl"])]),v(_,{data:M.value},{default:f((()=>[v(i,{align:"left",label:"预览",width:"100"},{default:f((e=>[v(s,{"pic-type":"file","pic-src":e.row.url},null,8,["pic-src"])])),_:1}),v(i,{align:"left",label:"日期",prop:"UpdatedAt",width:"180"},{default:f((e=>[m("div",null,y(h(o)(e.row.UpdatedAt)),1)])),_:1}),v(i,{align:"left",label:"文件名",prop:"name",width:"180"}),v(i,{align:"left",label:"链接",prop:"url","min-width":"300"}),v(i,{align:"left",label:"标签",prop:"tag",width:"100"},{default:f((e=>[v(r,{type:"jpg"===e.row.tag?"primary":"success","disable-transitions":""},{default:f((()=>[w(y(e.row.tag),1)])),_:2},1032,["type"])])),_:1}),v(i,{align:"left",label:"操作",width:"160"},{default:f((e=>[v(l,{size:"small",icon:"download",type:"text",onClick:a=>{var t;(t=e.row).url.indexOf("http://")>-1||t.url.indexOf("https://")>-1?b(t.url,t.name):b(O.value+t.url,t.name)}},{default:f((()=>[U])),_:2},1032,["onClick"]),v(l,{size:"small",icon:"delete",type:"text",onClick:a=>K(e.row)},{default:f((()=>[k])),_:2},1032,["onClick"])])),_:1})])),_:1},8,["data"]),m("div",A,[v(x,{"current-page":B.value,"page-size":I.value,"page-sizes":[10,30,50,100],style:{float:"right",padding:"20px"},total:D.value,layout:"total, sizes, prev, pager, next, jumper",onCurrentChange:T,onSizeChange:P},null,8,["current-page","page-size","total"])])])])),[[F,L.value,void 0,{fullscreen:!0,lock:!0}]])}}}),[["__scopeId","data-v-142e059e"]]);export{E as default};