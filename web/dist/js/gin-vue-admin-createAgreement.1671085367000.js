/*! 
 Build based on gin-vue-admin 
 Time : 1671085367000 */
var e=(e,t,a)=>new Promise(((o,n)=>{var l=e=>{try{r(a.next(e))}catch(t){n(t)}},i=e=>{try{r(a.throw(e))}catch(t){n(t)}},r=e=>e.done?o(e.value):Promise.resolve(e.value).then(l,i);r((a=a.apply(e,t)).next())}));import{f as t,c as a,u as o}from"./gin-vue-admin-agreement.16710853670002.js";import{p as n,w as l,Q as i,$ as r,a0 as s,a1 as u,o as d,j as c,q as p,x as m,C as g,v as f,J as b,u as v,r as y,i as _,k as h,t as k,A as j,d as w}from"./gin-vue-admin-vendor.1671085367000.js";import{u as C}from"../gva/gin-vue-admin-index.1671085367000.js";var O=Object.defineProperty,x=Object.defineProperties,V=Object.getOwnPropertyDescriptors,D=Object.getOwnPropertySymbols,P=Object.prototype.hasOwnProperty,E=Object.prototype.propertyIsEnumerable,T=(e,t,a)=>t in e?O(e,t,{enumerable:!0,configurable:!0,writable:!0,value:a}):e[t]=a;const S=e=>e+"_"+(Date.now()+Math.floor(1e6*Math.random())),z=()=>{const e="undefined"!=typeof window?window:global;return e&&"tinymce"in e?e.tinymce:null};function A(e){return e?e.getContent():""}function L(e,t){t&&t.setContent(e)}function q(e,t=!0){e&&e.mode.set(t?"readonly":"design")}function I(e,t,a,o,n){let l,i,{images_upload_url:r,images_upload_credentials:s,custom_images_upload_header:u,custom_images_upload_param:d,custom_images_upload_callback:c}=e||{};l=new XMLHttpRequest,l.withCredentials=!!s,l.open("POST",r||""),u&&Object.keys(u).forEach((e=>{l.setRequestHeader(e,u[e])})),l.upload.onprogress=function(e){n(e.loaded/e.total*100)},l.onload=function(){if(403===l.status)return void o("HTTP Error (custom): status "+l.status,{remove:!0});if(l.status<200||l.status>=300)return void o("HTTP Error (custom): status "+l.status);let e=JSON.parse(l.responseText);if(!e)return void o("Invalid JSON (custom): "+l.responseText);let t="function"==typeof c?c(e):e.location;t?a(t):o("Failed Path (custom): location image path is error/empty")},l.onerror=function(){o("Image upload failed due to a XHR Transport error. Code: "+l.status)},i=new FormData,i.append("file",t.blob(),t.filename()),d&&Object.keys(d).forEach((e=>{i.append(e,d[e])})),l.send(i)}const N=(()=>{let e={status:{},loadedCallbacks:{}};const t=t=>{e.loadedCallbacks[t]&&(e.loadedCallbacks[t].forEach((e=>{"function"==typeof e.handler&&e.handler.call(e.scope)})),e.loadedCallbacks[t]=void 0)};return{load:(a,o,n)=>{o&&(e.loadedCallbacks[a]||(e.loadedCallbacks[a]=[]),e.loadedCallbacks[a].push({handler:o,scope:n||void 0})),"LOADED"!==e.status[a]?"LOADING"!==e.status[a]&&(e.status[a]="LOADING",((o,n)=>{const l=document.createElement("script");l.id=S("tiny-script"),l.type="application/javascript",l.src=o,l.referrerPolicy="origin";const i=()=>{l.removeEventListener("load",i),e.status[a]="LOADED",t(a)};l.addEventListener("load",i),(document.head||document.body).appendChild(l)})(a)):t(a)}}})(),H=["id"],M={key:0},R=Object.assign({name:"Vue3Tinymce"},{props:{modelValue:String,setting:{type:Object,default:()=>({})},setup:Function,disabled:Boolean,scriptSrc:String,debug:Boolean},emits:["update:modelValue","init","change"],setup:function(e,{expose:t,emit:a}){const o=e;let v=!0;const y=n({editor:null,id:S("tinymce"),err:""}),_=()=>{var e;return String(null!=(e=o.modelValue)?e:"")},h=(e,t,a)=>{o.debug&&console.log(`来自：${e.type} | \n ${t} \n ${a||"--"}`)},k=()=>{if(null===z())return void(y.err="tinymce is null");o.debug&&console.warn("vue3-tinymce 进入debug模式");let e=(t=((e,t)=>{for(var a in t||(t={}))P.call(t,a)&&T(e,a,t[a]);if(D)for(var a of D(t))E.call(t,a)&&T(e,a,t[a]);return e})({},o.setting),n={selector:"#"+y.id,setup:e=>{o.setup&&o.setup(e),e.on("init",(()=>{return t=e,y.editor=t,L(_(),t),o.disabled&&"readonly"!==t.mode.get()&&q(t),t.on("change input undo redo",(e=>{((e,t)=>{t||(t=y.editor);const o=A(t);h(e,o),a("update:modelValue",o),a("change",o)})(e,t)})),void a("init",t);var t}))}},x(t,V(n)));var t,n;o.setting.custom_images_upload&&(e.images_upload_handler=(...e)=>{I.apply(null,[o.setting||{},...e])}),z().init(e),v=!1};return l((()=>o.modelValue),((e,t)=>{if(y.editor&&y.editor.initialized&&t!==e&&e!==A(y.editor)){if(h({type:"propsChanged to setContent"},e,t),null===e)return function(e,t){if(t){if(t.resetContent)return t.resetContent("");t.setContent(""),t.setDirty(!1),t.undoManager.clear()}}(0,y.editor);L(_(),y.editor)}})),l((()=>o.disabled),(e=>{y.editor&&y.editor.initialized&&q(y.editor,e)})),t({id:y.id,editor:y.editor}),i((()=>{var e;if(null!==z())return void k();const t=null!=(e=o.scriptSrc)?e:"https://cdn.bootcdn.net/ajax/libs/tinymce/5.8.2/tinymce.min.js";N.load(t,k)})),r((()=>{v||k()})),s((()=>{y.editor&&y.editor.remove()})),u((()=>{y.editor&&y.editor.remove()})),(e,t)=>(d(),c(b,null,[p("div",{id:m(y).id,class:"tiny-textarea"},null,8,H),m(y).err?(d(),c("p",M,g(m(y).err),1)):f("",!0)],64))}});R.install=function(e){e.component("Vue3Tinymce",R)};const $={class:"gva-search-box"},F=j("立即保存"),J=j("取消"),U={name:"Agreement"},X=Object.assign(U,{setup(n){const l=v(),i=y({title:"",contents:""}),r=C(),s=y("");s.value="create";const u=a=>e(this,null,(function*(){const e=yield t({ID:a});console.log("详情",e),0===e.code&&(i.value=e.data.reagreement,b.value.content=e.data.reagreement.contents)})),g=l.currentRoute.value.query.id;console.log("id",g),g&&(s.value="update",u(g)),console.log("表单信息",i.value);const f=()=>e(this,null,(function*(){let e;switch(i.value.contents=b.value.content,console.log(i.value),s.value){case"create":e=yield a(i.value);break;case"update":e=yield o(i.value);break;default:e=yield a(i.value)}0===e.code&&(w({type:"success",message:"创建/更改成功"}),l.push({name:"agreement"}))})),b=y({content:"",setting:{height:560,toolbar:"undo redo | fullscreen | formatselect alignleft aligncenter alignright alignjustify | link unlink | numlist bullist | image media table | fontsizeselect forecolor backcolor | bold italic underline strikethrough | indent outdent | superscript subscript | removeformat |",toolbar_drawer:"sliding",quickbars_selection_toolbar:"removeformat | bold italic underline strikethrough | fontsizeselect forecolor backcolor",plugins:"link image media table lists fullscreen quickbars",fontsize_formats:"12px 14px 16px 18px",default_link_target:"_blank",link_title:!1,nonbreaking_force_tab:!0,custom_images_upload:!0,images_upload_url:"/api/fileUploadAndDownload/upload",custom_images_upload_callback:e=>"/api/"+e.data.file.url,custom_images_upload_header:{"X-Token":r.token},language:"zh_CN",language_url:"https://unpkg.com/@jsdawn/vue3-tinymce@1.1.6/dist/tinymce/langs/zh_CN.js"}});return(e,t)=>{const a=_("el-input"),o=_("el-form-item"),n=_("el-button"),l=_("el-form");return d(),c("div",null,[p("div",$,[h(l,{model:i.value,"label-width":"80px"},{default:k((()=>[h(o,{label:"标题"},{default:k((()=>[h(a,{modelValue:i.value.title,"onUpdate:modelValue":t[0]||(t[0]=e=>i.value.title=e),clearable:"",placeholder:"请输入"},null,8,["modelValue"])])),_:1}),h(o,{label:"内容"},{default:k((()=>[h(m(R),{modelValue:b.value.content,"onUpdate:modelValue":t[1]||(t[1]=e=>b.value.content=e),setting:b.value.setting},null,8,["modelValue","setting"])])),_:1}),h(o,null,{default:k((()=>[h(n,{type:"primary",onClick:f},{default:k((()=>[F])),_:1}),h(n,null,{default:k((()=>[J])),_:1})])),_:1})])),_:1},8,["model"])])])}}});export{X as default};