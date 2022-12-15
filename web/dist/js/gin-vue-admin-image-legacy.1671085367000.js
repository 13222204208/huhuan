/*! 
 Build based on gin-vue-admin 
 Time : 1671085367000 */
!function(){function e(e,t){if(!(e instanceof t))throw new TypeError("Cannot call a class as a function")}function t(e,t){for(var i=0;i<t.length;i++){var a=t[i];a.enumerable=a.enumerable||!1,a.configurable=!0,"value"in a&&(a.writable=!0),Object.defineProperty(e,a.key,a)}}var i=document.createElement("style");i.innerHTML=".image-uploader[data-v-15c9badb]{border:1px dashed #d9d9d9;width:180px;border-radius:6px;cursor:pointer;position:relative;overflow:hidden}.image-uploader[data-v-15c9badb]{border-color:#409eff}.image-uploader-icon[data-v-15c9badb]{font-size:28px;color:#8c939d;width:178px;height:178px;line-height:178px;text-align:center}.image[data-v-15c9badb]{width:178px;height:178px;display:block}\n",document.head.appendChild(i),System.register(["../gva/gin-vue-admin-index-legacy.1671085367000.js","./gin-vue-admin-vendor-legacy.1671085367000.js"],(function(i){"use strict";var a,n,r,o,l,d,u,c,s,f,h,p;return{setters:[function(e){a=e.s,n=e._,r=e.u},function(e){o=e.r,l=e.i,d=e.o,u=e.j,c=e.k,s=e.t,f=e.x,h=e.A,p=e.d}],execute:function(){i("g",(function(e){return a({url:"/fileUploadAndDownload/getFileList",method:"post",data:e})})),i("d",(function(e){return a({url:"/fileUploadAndDownload/deleteFile",method:"post",data:e})})),i("e",(function(e){return a({url:"/fileUploadAndDownload/editFileName",method:"post",data:e})}));var g=function(){function i(t,a){var n=arguments.length>2&&void 0!==arguments[2]?arguments[2]:1920;e(this,i),this.file=t,this.fileSize=a,this.maxWH=n}var a,n,r;return a=i,(n=[{key:"compress",value:function(){var e=this,t=this.file.type,i=this.file.size/1024;return new Promise((function(a){var n=new FileReader;n.readAsDataURL(e.file),n.onload=function(){var r=document.createElement("canvas"),o=document.createElement("img");o.src=n.result,o.onload=function(){var n=r.getContext("2d"),l=e.dWH(o.width,o.height,e.maxWH);r.width=l.width,r.height=l.height,n.clearRect(0,0,r.width,r.height),n.drawImage(o,0,0,r.width,r.height);var d=r.toDataURL(t,.9),u=e.fileSizeKB(d);u>e.fileSize&&console.log("图片尺寸太大!"+i+" >> "+u);var c=e.dataURLtoBlob(d,t),s=new File([c],e.file.name);a(s)}}}))}},{key:"dWH",value:function(e,t,i){var a={width:e,height:t};return Math.max(e,t)>i?e>t?(a.width=i,a.height=Math.round(t*(i/e)),a):(a.height=i,a.width=Math.round(e*(i/t)),a):a}},{key:"fileSizeKB",value:function(e){return Math.round(3*e.split(",")[1].length/4/1024)}},{key:"dataURLtoBlob",value:function(e,t){for(var i=atob(e.split(",")[1]),a=e.split(",")[0].split(":")[1].split(";")[0],n=new ArrayBuffer(i.length),r=new Uint8Array(n),o=0;o<i.length;o++)r[o]=i.charCodeAt(o);return t&&(a=t),new Blob([n],{type:a,lastModifiedDate:new Date})}}])&&t(a.prototype,n),r&&t(a,r),Object.defineProperty(a,"prototype",{writable:!1}),i}(),m=h("压缩上传"),v={name:"UploadImage",methods:{}};i("U",n(Object.assign(v,{props:{imageUrl:{type:String,default:""},fileSize:{type:Number,default:2048},maxWH:{type:Number,default:1920}},emits:["on-success"],setup:function(e,t){var i=t.emit,a=e,n=o("/api"),h=r(),v=function(e){var t="image/jpeg"===e.type,i="image/png"===e.type;if(!t&&!i)return p.error("上传头像图片只能是 jpg或png 格式!"),!1;var n=e.size/1024<a.fileSize;return n||new g(e,a.fileSize,a.maxWH).compress()},b=function(e){var t=e.data;t.file&&i("on-success",t.file.url)};return function(e,t){var i=l("el-button"),a=l("el-upload");return d(),u("div",null,[c(a,{action:"".concat(n.value,"/fileUploadAndDownload/upload"),headers:{"x-token":f(h).token},"show-file-list":!1,"on-success":b,"before-upload":v,multiple:!1},{default:s((function(){return[c(i,{size:"small",type:"primary"},{default:s((function(){return[m]})),_:1})]})),_:1},8,["action","headers"])])}}}),[["__scopeId","data-v-15c9badb"]]))}}}))}();
