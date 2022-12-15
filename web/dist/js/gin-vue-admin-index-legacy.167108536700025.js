/*! 
 Build based on gin-vue-admin 
 Time : 1671085367000 */
!function(){function e(e,t,n,a,l,u,r){try{var i=e[u](r),o=i.value}catch(c){return void n(c)}i.done?t(o):Promise.resolve(o).then(a,l)}function t(t){return function(){var n=this,a=arguments;return new Promise((function(l,u){var r=t.apply(n,a);function i(t){e(r,l,u,i,o,"next",t)}function o(t){e(r,l,u,i,o,"throw",t)}i(void 0)}))}}var n=document.createElement("style");n.innerHTML=".previewCodeTool[data-v-9a9c27fe]{display:flex;align-items:center;padding:5px 0}.button-box[data-v-9a9c27fe]{padding:10px 20px}.button-box .el-button[data-v-9a9c27fe]{margin-right:20px;float:right}.auto-btn-list[data-v-9a9c27fe]{margin-top:16px}\n",document.head.appendChild(n),System.register(["./gin-vue-admin-fieldDialog-legacy.1671085367000.js","./gin-vue-admin-previewCodeDialg-legacy.1671085367000.js","./gin-vue-admin-stringFun-legacy.1671085367000.js","./gin-vue-admin-autoCode-legacy.1671085367000.js","./gin-vue-admin-dictionary-legacy.1671085367000.js","./gin-vue-admin-vendor-legacy.1671085367000.js","../gva/gin-vue-admin-index-legacy.1671085367000.js","./gin-vue-admin-sysDictionary-legacy.1671085367000.js","./gin-vue-admin-warningBar-legacy.1671085367000.js","./gin-vue-admin-dictionary-legacy.16710853670002.js"],(function(e){"use strict";var n,a,l,u,r,i,o,c,d,f,s,p,m,v,b,g,y,x,N,h,k,w,_,C,V,T,R,U,S,z,j,F,L,J;return{setters:[function(e){n=e.default},function(e){a=e.default},function(e){l=e.a,u=e.b,r=e.t,i=e.c},function(e){o=e.p,c=e.c,d=e.g,f=e.a,s=e.b,p=e.d},function(e){m=e.g},function(e){v=e.U,b=e.p,g=e.r,y=e.i,x=e.o,N=e.j,h=e.q,k=e.k,w=e.t,_=e.J,C=e.L,V=e.x,T=e.a2,R=e.H,U=e.v,S=e.A,z=e.F,j=e.G,F=e.d,L=e.a3},function(e){J=e._},function(){},function(){},function(){}],execute:function(){var q=function(e){return z("data-v-9a9c27fe"),e=e(),j(),e},A={class:"gva-search-box"},D={style:{fontSize:"16px",paddingLeft:"20px"}},O=S(" 点这里从现有数据库创建代码 "),E=S("使用此表创建"),M={class:"gva-search-box"},P=q((function(){return h("div",null," 自动创建API ",-1)})),$=q((function(){return h("div",null," 自动移动文件 ",-1)})),B={class:"gva-table-box"},X={class:"gva-btn-list"},I=S("新增Field"),H=S("编辑"),G=S("上移"),K=S("下移"),Q=q((function(){return h("p",null,"确定删除吗？",-1)})),W={style:{"text-align":"right","margin-top":"8px"}},Y=S("取消"),Z=S("确定"),ee=S("删除"),te={class:"gva-btn-list justify-content-flex-end auto-btn-list"},ne=S("预览代码"),ae=S("生成代码"),le={class:"dialog-footer"},ue=S("取 消"),re=S("确 定"),ie={class:"previewCodeTool"},oe=q((function(){return h("p",null,"操作栏：",-1)})),ce=S("全选"),de=S("复制"),fe={class:"dialog-footer",style:{"padding-top":"14px","padding-right":"14px"}},se=S("确 定"),pe={name:"AutoCode"};e("default",J(Object.assign(pe,{setup:function(e){var S={fieldName:"",fieldDesc:"",fieldType:"",dataType:"",fieldJson:"",columnName:"",dataTypeLong:"",comment:"",fieldSearchType:"",dictType:""},z=v(),j=b([]),J=g({}),q=g({dbName:"",tableName:""}),pe=g([]),me=g([]),ve=g(""),be=g({}),ge=g({structName:"",tableName:"",packageName:"",abbreviation:"",description:"",autoCreateApiToSql:!1,autoMoveFile:!1,fields:[]}),ye=g({structName:[{required:!0,message:"请输入结构体名称",trigger:"blur"}],abbreviation:[{required:!0,message:"请输入结构体简称",trigger:"blur"}],description:[{required:!0,message:"请输入结构体描述",trigger:"blur"}],packageName:[{required:!0,message:"文件名称：sysXxxxXxxx",trigger:"blur"}]}),xe=g({}),Ne=g({}),he=g(!1),ke=g(!1),we=g(null),_e=function(){we.value.selectText()},Ce=function(){we.value.copy()},Ve=function(e){he.value=!0,e?(ve.value="edit",Ne.value=JSON.parse(JSON.stringify(e)),xe.value=e):(ve.value="add",xe.value=JSON.parse(JSON.stringify(S)))},Te=L(),Re=function(){Te.refs.fieldDialogNode.fieldDialogFrom.validate((function(e){if(!e)return!1;xe.value.fieldName=u(xe.value.fieldName),"add"===ve.value&&ge.value.fields.push(xe.value),he.value=!1}))},Ue=function(){"edit"===ve.value&&(xe.value=Ne.value),he.value=!1},Se=g(null),ze=function(){var e=t(regeneratorRuntime.mark((function e(n){return regeneratorRuntime.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:if(!(ge.value.fields.length<=0)){e.next=3;break}return F({type:"error",message:"请填写至少一个field"}),e.abrupt("return",!1);case 3:if(!ge.value.fields.some((function(e){return e.fieldName===ge.value.structName}))){e.next=6;break}return F({type:"error",message:"存在与结构体同名的字段"}),e.abrupt("return",!1);case 6:Se.value.validate(function(){var e=t(regeneratorRuntime.mark((function e(t){var a,l,i,d,f,s,p,m;return regeneratorRuntime.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:if(!t){e.next=32;break}for(a in ge.value)"string"==typeof ge.value[a]&&(ge.value[a]=ge.value[a].trim());if(ge.value.structName=u(ge.value.structName),ge.value.tableName&&(ge.value.tableName=ge.value.tableName.replace(" ","")),ge.value.structName!==ge.value.abbreviation){e.next=7;break}return F({type:"error",message:"structName和struct简称不能相同"}),e.abrupt("return",!1);case 7:if(ge.value.humpPackageName=r(ge.value.packageName),!n){e.next=16;break}return e.next=11,o(ge.value);case 11:l=e.sent,J.value=l.data.autoCode,ke.value=!0,e.next=30;break;case 16:return e.next=18,c(ge.value);case 18:if(d=e.sent,"false"!==(null===(i=d.headers)||void 0===i?void 0:i.success)){e.next=23;break}return e.abrupt("return");case 23:if(!ge.value.autoMoveFile){e.next=26;break}return F({type:"success",message:"自动化代码创建成功，自动移动成功"}),e.abrupt("return");case 26:F({type:"success",message:"自动化代码创建成功，正在下载"});case 27:f=new Blob([d]),s="ginvueadmin.zip","download"in document.createElement("a")?(p=window.URL.createObjectURL(f),(m=document.createElement("a")).style.display="none",m.href=p,m.setAttribute("download",s),document.body.appendChild(m),m.click(),document.body.removeChild(m),window.URL.revokeObjectURL(p)):window.navigator.msSaveBlob(f,s);case 30:e.next=33;break;case 32:return e.abrupt("return",!1);case 33:case"end":return e.stop()}}),e)})));return function(t){return e.apply(this,arguments)}}());case 7:case"end":return e.stop()}}),e)})));return function(t){return e.apply(this,arguments)}}(),je=function(){var e=t(regeneratorRuntime.mark((function e(){var t;return regeneratorRuntime.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,s();case 2:0===(t=e.sent).code&&(pe.value=t.data.dbs);case 4:case"end":return e.stop()}}),e)})));return function(){return e.apply(this,arguments)}}(),Fe=function(){var e=t(regeneratorRuntime.mark((function e(){var t;return regeneratorRuntime.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,d({dbName:q.value.dbName});case 2:0===(t=e.sent).code&&(me.value=t.data.tables),q.value.tableName="";case 5:case"end":return e.stop()}}),e)})));return function(){return e.apply(this,arguments)}}(),Le=function(){var e=t(regeneratorRuntime.mark((function e(){var t,n,a;return regeneratorRuntime.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return t=["id","created_at","updated_at","deleted_at"],e.next=3,f(q.value);case 3:0===(n=e.sent).code&&(a=i(q.value.tableName),ge.value.structName=u(a),ge.value.tableName=q.value.tableName,ge.value.packageName=a,ge.value.abbreviation=a,ge.value.description=a+"表",ge.value.autoCreateApiToSql=!0,ge.value.fields=[],n.data.columns&&n.data.columns.forEach((function(e){if(!t.some((function(t){return t===e.columnName}))){var n=i(e.columnName);ge.value.fields.push({fieldName:u(n),fieldDesc:e.columnComment||n+"字段",fieldType:be.value[e.dataType],dataType:e.dataType,fieldJson:n,dataTypeLong:e.dataTypeLong&&e.dataTypeLong.split(",")[0],columnName:e.columnName,comment:e.columnComment,fieldSearchType:"",dictType:""})}})));case 5:case"end":return e.stop()}}),e)})));return function(){return e.apply(this,arguments)}}(),Je=function(){var e=t(regeneratorRuntime.mark((function e(){return regeneratorRuntime.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:["string","int","bool","float64","time.Time"].forEach(function(){var e=t(regeneratorRuntime.mark((function e(t){var n;return regeneratorRuntime.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,m(t);case 2:(n=e.sent)&&n.forEach((function(e){be.value[e.label]=t}));case 4:case"end":return e.stop()}}),e)})));return function(t){return e.apply(this,arguments)}}());case 2:case"end":return e.stop()}}),e)})));return function(){return e.apply(this,arguments)}}(),qe=function(){var e=t(regeneratorRuntime.mark((function e(t){var n;return regeneratorRuntime.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,p({id:Number(t)});case 2:0===(n=e.sent).code&&(ge.value=JSON.parse(n.data.meta));case 4:case"end":return e.stop()}}),e)})));return function(t){return e.apply(this,arguments)}}();return function(){je(),Je();var e=z.params.id;e&&qe(e)}(),function(e,t){var u=y("pointer"),r=y("el-icon"),i=y("el-option"),o=y("el-select"),c=y("el-form-item"),d=y("el-button"),f=y("el-form"),s=y("el-collapse-item"),p=y("el-collapse"),m=y("el-input"),v=y("el-tooltip"),b=y("el-checkbox"),g=y("el-table-column"),S=y("el-popover"),z=y("el-table"),F=y("el-dialog");return x(),N("div",null,[h("div",A,[k(p,{modelValue:V(j),"onUpdate:modelValue":t[2]||(t[2]=function(e){return T(j)?j.value=e:function(e){throw new TypeError('"'+e+'" is read-only')}("activeNames")}),style:{"margin-bottom":"12px"}},{default:w((function(){return[k(s,{name:"1"},{title:w((function(){return[h("div",D,[O,k(r,{class:"header-icon"},{default:w((function(){return[k(u)]})),_:1})])]})),default:w((function(){return[k(f,{ref:"getTableForm",style:{"margin-top":"24px"},inline:!0,model:q.value,"label-width":"120px"},{default:w((function(){return[k(c,{label:"数据库名",prop:"structName"},{default:w((function(){return[k(o,{modelValue:q.value.dbName,"onUpdate:modelValue":t[0]||(t[0]=function(e){return q.value.dbName=e}),filterable:"",placeholder:"请选择数据库",onChange:Fe},{default:w((function(){return[(x(!0),N(_,null,C(pe.value,(function(e){return x(),R(i,{key:e.database,label:e.database,value:e.database},null,8,["label","value"])})),128))]})),_:1},8,["modelValue"])]})),_:1}),k(c,{label:"表名",prop:"structName"},{default:w((function(){return[k(o,{modelValue:q.value.tableName,"onUpdate:modelValue":t[1]||(t[1]=function(e){return q.value.tableName=e}),disabled:!q.value.dbName,filterable:"",placeholder:"请选择表"},{default:w((function(){return[(x(!0),N(_,null,C(me.value,(function(e){return x(),R(i,{key:e.tableName,label:e.tableName,value:e.tableName},null,8,["label","value"])})),128))]})),_:1},8,["modelValue","disabled"])]})),_:1}),k(c,null,{default:w((function(){return[k(d,{size:"small",type:"primary",onClick:Le},{default:w((function(){return[E]})),_:1})]})),_:1})]})),_:1},8,["model"])]})),_:1})]})),_:1},8,["modelValue"])]),h("div",M,[k(f,{ref_key:"autoCodeForm",ref:Se,rules:ye.value,model:ge.value,"label-width":"120px",inline:!0},{default:w((function(){return[k(c,{label:"Struct名称",prop:"structName"},{default:w((function(){return[k(m,{modelValue:ge.value.structName,"onUpdate:modelValue":t[3]||(t[3]=function(e){return ge.value.structName=e}),placeholder:"首字母自动转换大写"},null,8,["modelValue"])]})),_:1}),k(c,{label:"TableName",prop:"tableName"},{default:w((function(){return[k(m,{modelValue:ge.value.tableName,"onUpdate:modelValue":t[4]||(t[4]=function(e){return ge.value.tableName=e}),placeholder:"指定表名（非必填）"},null,8,["modelValue"])]})),_:1}),k(c,{label:"Struct简称",prop:"abbreviation"},{default:w((function(){return[k(m,{modelValue:ge.value.abbreviation,"onUpdate:modelValue":t[5]||(t[5]=function(e){return ge.value.abbreviation=e}),placeholder:"简称会作为入参对象名和路由group"},null,8,["modelValue"])]})),_:1}),k(c,{label:"Struct中文名称",prop:"description"},{default:w((function(){return[k(m,{modelValue:ge.value.description,"onUpdate:modelValue":t[6]||(t[6]=function(e){return ge.value.description=e}),placeholder:"中文描述作为自动api描述"},null,8,["modelValue"])]})),_:1}),k(c,{label:"文件名称",prop:"packageName"},{default:w((function(){return[k(m,{modelValue:ge.value.packageName,"onUpdate:modelValue":t[7]||(t[7]=function(e){return ge.value.packageName=e}),placeholder:"生成文件的默认名称(建议为驼峰格式,首字母小写,如sysXxxXxxx)",onBlur:t[8]||(t[8]=function(e){return function(e,t){e[t]=l(e[t])}(ge.value,"packageName")})},null,8,["modelValue"])]})),_:1}),k(c,null,{label:w((function(){return[k(v,{content:"注：把自动生成的API注册进数据库",placement:"bottom",effect:"light"},{default:w((function(){return[P]})),_:1})]})),default:w((function(){return[k(b,{modelValue:ge.value.autoCreateApiToSql,"onUpdate:modelValue":t[9]||(t[9]=function(e){return ge.value.autoCreateApiToSql=e})},null,8,["modelValue"])]})),_:1}),k(c,null,{label:w((function(){return[k(v,{content:"注：自动迁移生成的文件到ymal配置的对应位置",placement:"bottom",effect:"light"},{default:w((function(){return[$]})),_:1})]})),default:w((function(){return[k(b,{modelValue:ge.value.autoMoveFile,"onUpdate:modelValue":t[10]||(t[10]=function(e){return ge.value.autoMoveFile=e})},null,8,["modelValue"])]})),_:1})]})),_:1},8,["rules","model"])]),h("div",B,[h("div",X,[k(d,{size:"small",type:"primary",onClick:t[11]||(t[11]=function(e){return Ve()})},{default:w((function(){return[I]})),_:1})]),k(z,{data:ge.value.fields},{default:w((function(){return[k(g,{align:"left",type:"index",label:"序列",width:"100"}),k(g,{align:"left",prop:"fieldName",label:"Field名"}),k(g,{align:"left",prop:"fieldDesc",label:"中文名"}),k(g,{align:"left",prop:"fieldJson",label:"FieldJson"}),k(g,{align:"left",prop:"fieldType",label:"Field数据类型",width:"130"}),k(g,{align:"left",prop:"dataTypeLong",label:"数据库字段长度",width:"130"}),k(g,{align:"left",prop:"columnName",label:"数据库字段",width:"130"}),k(g,{align:"left",prop:"comment",label:"数据库字段描述",width:"130"}),k(g,{align:"left",prop:"fieldSearchType",label:"搜索条件",width:"130"}),k(g,{align:"left",prop:"dictType",label:"字典",width:"130"}),k(g,{align:"left",label:"操作",width:"300"},{default:w((function(e){return[k(d,{size:"small",type:"text",icon:"edit",onClick:function(t){return Ve(e.row)}},{default:w((function(){return[H]})),_:2},1032,["onClick"]),k(d,{size:"small",type:"text",disabled:0===e.$index,onClick:function(t){return function(e){if(0!==e){var t=ge.value.fields[e-1];ge.value.fields.splice(e-1,1),ge.value.fields.splice(e,0,t)}}(e.$index)}},{default:w((function(){return[G]})),_:2},1032,["disabled","onClick"]),k(d,{size:"small",type:"text",disabled:e.$index+1===ge.value.fields.length,onClick:function(t){return function(e){if(e!==ge.value.fields.length-1){var t=ge.value.fields[e+1];ge.value.fields.splice(e+1,1),ge.value.fields.splice(e,0,t)}}(e.$index)}},{default:w((function(){return[K]})),_:2},1032,["disabled","onClick"]),k(S,{visible:e.row.visible,"onUpdate:visible":function(t){return e.row.visible=t},placement:"top"},{reference:w((function(){return[k(d,{size:"small",type:"text",icon:"delete",onClick:function(t){return e.row.visible=!0}},{default:w((function(){return[ee]})),_:2},1032,["onClick"])]})),default:w((function(){return[Q,h("div",W,[k(d,{size:"small",type:"text",onClick:function(t){return e.row.visible=!1}},{default:w((function(){return[Y]})),_:2},1032,["onClick"]),k(d,{type:"primary",size:"small",onClick:function(t){return function(e){ge.value.fields.splice(e,1)}(e.$index)}},{default:w((function(){return[Z]})),_:2},1032,["onClick"])])]})),_:2},1032,["visible","onUpdate:visible"])]})),_:1})]})),_:1},8,["data"]),h("div",te,[k(d,{size:"small",type:"primary",onClick:t[12]||(t[12]=function(e){return ze(!0)})},{default:w((function(){return[ne]})),_:1}),k(d,{size:"small",type:"primary",onClick:t[13]||(t[13]=function(e){return ze(!1)})},{default:w((function(){return[ae]})),_:1})])]),k(F,{modelValue:he.value,"onUpdate:modelValue":t[14]||(t[14]=function(e){return he.value=e}),title:"组件内容"},{footer:w((function(){return[h("div",le,[k(d,{size:"small",onClick:Ue},{default:w((function(){return[ue]})),_:1}),k(d,{size:"small",type:"primary",onClick:Re},{default:w((function(){return[re]})),_:1})])]})),default:w((function(){return[he.value?(x(),R(n,{key:0,ref:"fieldDialogNode","dialog-middle":xe.value},null,8,["dialog-middle"])):U("",!0)]})),_:1},8,["modelValue"]),k(F,{modelValue:ke.value,"onUpdate:modelValue":t[16]||(t[16]=function(e){return ke.value=e})},{title:w((function(){return[h("div",ie,[oe,k(d,{size:"small",type:"primary",onClick:_e},{default:w((function(){return[ce]})),_:1}),k(d,{size:"small",type:"primary",onClick:Ce},{default:w((function(){return[de]})),_:1})])]})),footer:w((function(){return[h("div",fe,[k(d,{size:"small",type:"primary",onClick:t[15]||(t[15]=function(e){return ke.value=!1})},{default:w((function(){return[se]})),_:1})])]})),default:w((function(){return[ke.value?(x(),R(a,{key:0,ref_key:"previewNode",ref:we,"preview-code":J.value},null,8,["preview-code"])):U("",!0)]})),_:1},8,["modelValue"])])}}}),[["__scopeId","data-v-9a9c27fe"]]))}}}))}();
