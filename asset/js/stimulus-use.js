import{Controller as J} from"@hotwired/stimulus";var Ga=function(g){const a=g.getBoundingClientRect(),f=window.innerHeight||document.documentElement.clientHeight,h=window.innerWidth||document.documentElement.clientWidth,k=a.top<=f&&a.top+a.height>0,j=a.left<=h&&a.left+a.width>0;return k&&j},$=function(g){return g.replace(/(?:[_-])([a-z0-9])/g,(a,f)=>f.toUpperCase())},Ha=function(g,a){var f={};for(var h in g)if(Object.prototype.hasOwnProperty.call(g,h)&&a.indexOf(h)<0)f[h]=g[h];if(g!=null&&typeof Object.getOwnPropertySymbols==="function"){for(var k=0,h=Object.getOwnPropertySymbols(g);k<h.length;k++)if(a.indexOf(h[k])<0&&Object.prototype.propertyIsEnumerable.call(g,h[k]))f[h[k]]=g[h[k]]}return f},e=function(g){const a=document.head.querySelector(`meta[name="${g}"]`);return a&&a.getAttribute("content")},aa=function(g){try{return JSON.parse(g)}catch(a){return g}},i=function(g,a=xa){let f;return function(){const h=arguments,k=this;if(!f)f=!0,g.apply(k,h),setTimeout(()=>f=!1,a)}},T=function(g,a,f){const h=`transition${g[0].toUpperCase()}${g.substr(1)}`,k=Wa[g],j=a[g]||f[h]||f[k]||" ";return oa(j)?[]:j.split(" ")};async function na(g){return new Promise((a)=>{const f=Number(getComputedStyle(g).transitionDuration.split(",")[0].replace("s",""))*1000;setTimeout(()=>{a(f)},f)})}async function ta(){return new Promise((g)=>{requestAnimationFrame(()=>{requestAnimationFrame(g)})})}var oa=function(g){return g.length===0||!g.trim()},Df=function(){throw"[stimulus-use] Notice: The import for `useHotkeys()` has been moved from `stimulus-use` to `stimulus-use/hotkeys`. \nPlease change the import accordingly and add `hotkey-js` as a dependency to your project. \n\nFor more information see: https://stimulus-use.github.io/stimulus-use/#/use-hotkeys?id=importing-the-behavior"},N=(g,a)=>{const f=g[a];if(typeof f=="function")return f;else return(...h)=>{}},Z=(g,a,f)=>{let h=g;if(f===!0)h=`${a.identifier}:${g}`;else if(typeof f==="string")h=`${f}:${g}`;return h},d=(g,a,f)=>{const{bubbles:h,cancelable:k,composed:j}=a||{bubbles:!0,cancelable:!0,composed:!0};if(a)Object.assign(f,{originalEvent:a});return new CustomEvent(g,{bubbles:h,cancelable:k,composed:j,detail:f})},R={debug:!1,logger:console,dispatchEvent:!0,eventPrefix:!0};class U{constructor(g,a={}){var f,h,k;this.log=(q,p)=>{if(!this.debug)return;this.logger.groupCollapsed(`%c${this.controller.identifier} %c#${q}`,"color: #3B82F6","color: unset"),this.logger.log(Object.assign({controllerId:this.controllerId},p)),this.logger.groupEnd()},this.warn=(q)=>{this.logger.warn(`%c${this.controller.identifier} %c${q}`,"color: #3B82F6; font-weight: bold","color: unset")},this.dispatch=(q,p={})=>{if(this.dispatchEvent){const{event:y}=p,A=Ha(p,["event"]),B=this.extendedEvent(q,y||null,A);this.targetElement.dispatchEvent(B),this.log("dispatchEvent",Object.assign({eventName:B.type},A))}},this.call=(q,p={})=>{const y=this.controller[q];if(typeof y=="function")return y.call(this.controller,p)},this.extendedEvent=(q,p,y)=>{const{bubbles:A,cancelable:B,composed:D}=p||{bubbles:!0,cancelable:!0,composed:!0};if(p)Object.assign(y,{originalEvent:p});return new CustomEvent(this.composeEventName(q),{bubbles:A,cancelable:B,composed:D,detail:y})},this.composeEventName=(q)=>{let p=q;if(this.eventPrefix===!0)p=`${this.controller.identifier}:${q}`;else if(typeof this.eventPrefix==="string")p=`${this.eventPrefix}:${q}`;return p},this.debug=(h=(f=a===null||a===void 0?void 0:a.debug)!==null&&f!==void 0?f:g.application.stimulusUseDebug)!==null&&h!==void 0?h:R.debug,this.logger=(k=a===null||a===void 0?void 0:a.logger)!==null&&k!==void 0?k:R.logger,this.controller=g,this.controllerId=g.element.id||g.element.dataset.id,this.targetElement=(a===null||a===void 0?void 0:a.element)||g.element;const{dispatchEvent:j,eventPrefix:m}=Object.assign({},R,a);Object.assign(this,{dispatchEvent:j,eventPrefix:m}),this.controllerInitialize=g.initialize.bind(g),this.controllerConnect=g.connect.bind(g),this.controllerDisconnect=g.disconnect.bind(g)}}var b={eventPrefix:!0,bubbles:!0,cancelable:!0};class l extends U{constructor(g,a={}){var f,h,k,j;super(g,a);this.dispatch=(m,q={})=>{const{controller:p,targetElement:y,eventPrefix:A,bubbles:B,cancelable:D,log:E,warn:I}=this;Object.assign(q,{controller:p});const O=Z(m,this.controller,A),P=new CustomEvent(O,{detail:q,bubbles:B,cancelable:D});return y.dispatchEvent(P),I("`useDispatch()` is deprecated. Please use the built-in `this.dispatch()` function from Stimulus. You can find more information on how to upgrade at: https://stimulus-use.github.io/stimulus-use/#/use-dispatch"),E("dispatch",{eventName:O,detail:q,bubbles:B,cancelable:D}),P},this.targetElement=(f=a.element)!==null&&f!==void 0?f:g.element,this.eventPrefix=(h=a.eventPrefix)!==null&&h!==void 0?h:b.eventPrefix,this.bubbles=(k=a.bubbles)!==null&&k!==void 0?k:b.bubbles,this.cancelable=(j=a.cancelable)!==null&&j!==void 0?j:b.cancelable,this.enhanceController()}enhanceController(){Object.assign(this.controller,{dispatch:this.dispatch})}}var Ia=(g,a={})=>new l(g,a),Ja={overwriteDispatch:!0},Ka=(g,a={})=>{const{overwriteDispatch:f}=Object.assign({},Ja,a);if(Object.defineProperty(g,"isPreview",{get(){return document.documentElement.hasAttribute("data-turbolinks-preview")||document.documentElement.hasAttribute("data-turbo-preview")}}),Object.defineProperty(g,"isConnected",{get(){return!!Array.from(this.context.module.connectedContexts).find((h)=>h===this.context)}}),Object.defineProperty(g,"csrfToken",{get(){return this.metaValue("csrf-token")}}),f)Ia(g,a);Object.assign(g,{metaValue(h){const k=document.head.querySelector(`meta[name="${h}"]`);return k&&k.getAttribute("content")}})};class La extends J{constructor(g){super(g);this.isPreview=!1,this.isConnected=!1,this.csrfToken="",Ka(this,this.options)}}var Ma={events:["click","touchend"],onlyVisible:!0,dispatchEvent:!0,eventPrefix:!0},Na=(g,a={})=>{const f=g,{onlyVisible:h,dispatchEvent:k,events:j,eventPrefix:m}=Object.assign({},Ma,a),q=(B)=>{const D=(a===null||a===void 0?void 0:a.element)||f.element;if(D.contains(B.target)||!Ga(D)&&h)return;if(f.clickOutside)f.clickOutside(B);if(k){const E=Z("click:outside",f,m),I=d(E,B,{controller:f});D.dispatchEvent(I)}},p=()=>{j===null||j===void 0||j.forEach((B)=>{window.addEventListener(B,q,!0)})},y=()=>{j===null||j===void 0||j.forEach((B)=>{window.removeEventListener(B,q,!0)})},A=f.disconnect.bind(f);return Object.assign(f,{disconnect(){y(),A()}}),p(),[p,y]};class v extends J{}class Oa extends v{constructor(g){super(g);requestAnimationFrame(()=>{const[a,f]=Na(this,this.options);Object.assign(this,{observe:a,unobserve:f})})}}class x extends J{}x.debounces=[];var Pa=200,c=(g,a=Pa)=>{let f=null;return function(){const h=Array.from(arguments),k=this,j=h.map((q)=>q.params),m=()=>{return h.forEach((q,p)=>q.params=j[p]),g.apply(k,h)};if(f)clearTimeout(f);f=setTimeout(m,a)}},pf=(g,a)=>{const f=g;f.constructor.debounces.forEach((k)=>{if(typeof k==="string")f[k]=c(f[k],a===null||a===void 0?void 0:a.wait);if(typeof k==="object"){const{name:j,wait:m}=k;if(!j)return;f[j]=c(f[j],m||(a===null||a===void 0?void 0:a.wait))}})};class W extends U{constructor(g,a={}){super(g,a);this.observe=()=>{this.targetElement.addEventListener("mouseenter",this.onEnter),this.targetElement.addEventListener("mouseleave",this.onLeave)},this.unobserve=()=>{this.targetElement.removeEventListener("mouseenter",this.onEnter),this.targetElement.removeEventListener("mouseleave",this.onLeave)},this.onEnter=(f)=>{this.call("mouseEnter",f),this.log("mouseEnter",{hover:!0}),this.dispatch("mouseEnter",{hover:!1})},this.onLeave=(f)=>{this.call("mouseLeave",f),this.log("mouseLeave",{hover:!1}),this.dispatch("mouseLeave",{hover:!1})},this.controller=g,this.enhanceController(),this.observe()}enhanceController(){const g=this.controller.disconnect.bind(this.controller),a=()=>{this.unobserve(),g()};Object.assign(this.controller,{disconnect:a})}}var Qa=(g,a={})=>{const h=new W(g,a);return[h.observe,h.unobserve]};class s extends J{}class Sa extends s{constructor(g){super(g);requestAnimationFrame(()=>{const[a,f]=Qa(this,this.options);Object.assign(this,{observe:a,unobserve:f})})}}var Ta=["mousemove","mousedown","resize","keydown","touchstart","wheel"],Ua=60000,Va={ms:Ua,initialState:!1,events:Ta,dispatchEvent:!0,eventPrefix:!0},Xa=(g,a={})=>{const f=g,{ms:h,initialState:k,events:j,dispatchEvent:m,eventPrefix:q}=Object.assign({},Va,a);let p=k,y=setTimeout(()=>{p=!0,A()},h);const A=(G)=>{const L=Z("away",f,q);if(f.isIdle=!0,N(f,"away").call(f,G),m){const Q=d(L,G||null,{controller:f});f.element.dispatchEvent(Q)}},B=(G)=>{const L=Z("back",f,q);if(f.isIdle=!1,N(f,"back").call(f,G),m){const Q=d(L,G||null,{controller:f});f.element.dispatchEvent(Q)}},D=(G)=>{if(p)B(G);p=!1,clearTimeout(y),y=setTimeout(()=>{p=!0,A(G)},h)},E=(G)=>{if(!document.hidden)D(G)};if(p)A();else B();const I=f.disconnect.bind(f),O=()=>{j.forEach((G)=>{window.addEventListener(G,D)}),document.addEventListener("visibilitychange",E)},P=()=>{clearTimeout(y),j.forEach((G)=>{window.removeEventListener(G,D)}),document.removeEventListener("visibilitychange",E)};return Object.assign(f,{disconnect(){P(),I()}}),O(),[O,P]};class r extends J{constructor(){super(...arguments);this.isIdle=!1}}class Ya extends r{constructor(g){super(g);requestAnimationFrame(()=>{const[a,f]=Xa(this,this.options);Object.assign(this,{observe:a,unobserve:f})})}}var Za={dispatchEvent:!0,eventPrefix:!0,visibleAttribute:"isVisible"},_a=(g,a={})=>{const f=g,{dispatchEvent:h,eventPrefix:k,visibleAttribute:j}=Object.assign({},Za,a),m=(a===null||a===void 0?void 0:a.element)||f.element;if(!f.intersectionElements)f.intersectionElements=[];f.intersectionElements.push(m);const p=new IntersectionObserver((H)=>{const[S]=H;if(S.isIntersecting)y(S);else if(m.hasAttribute(j))A(S)},a),y=(H)=>{if(m.setAttribute(j,"true"),N(f,"appear").call(f,H,p),h){const S=Z("appear",f,k),u=d(S,null,{controller:f,entry:H,observer:p});m.dispatchEvent(u)}},A=(H)=>{if(m.removeAttribute(j),N(f,"disappear").call(f,H,p),h){const S=Z("disappear",f,k),u=d(S,null,{controller:f,entry:H,observer:p});m.dispatchEvent(u)}},B=f.disconnect.bind(f),D=()=>{I(),B()},E=()=>{p.observe(m)},I=()=>{p.unobserve(m)},O=()=>f.intersectionElements.filter((H)=>H.hasAttribute(j)).length===0,P=()=>f.intersectionElements.filter((H)=>H.hasAttribute(j)).length===1,G=()=>f.intersectionElements.some((H)=>H.hasAttribute(j)),L=()=>f.intersectionElements.every((H)=>H.hasAttribute(j));return Object.assign(f,{isVisible:L,noneVisible:O,oneVisible:P,atLeastOneVisible:G,allVisible:L,disconnect:D}),E(),[E,I]};class n extends J{}class $a extends n{constructor(g){super(g);requestAnimationFrame(()=>{const[a,f]=_a(this,this.options);Object.assign(this,{observe:a,unobserve:f})})}}var da=(g,a)=>{const f=(y)=>{const[A]=y;if(A.isIntersecting&&!g.isLoaded)h()},h=(y)=>{const A=g.data.get("src");if(!A)return;const B=g.element;g.isLoading=!0,N(g,"loading").call(g,A),B.onload=()=>{k(A)},B.src=A},k=(y)=>{g.isLoading=!1,g.isLoaded=!0,N(g,"loaded").call(g,y)},j=g.disconnect.bind(g),m=new IntersectionObserver(f,a),q=()=>{m.observe(g.element)},p=()=>{m.unobserve(g.element)};return Object.assign(g,{isVisible:!1,disconnect(){p(),j()}}),q(),[q,p]};class t extends J{constructor(){super(...arguments);this.isLoading=!1,this.isLoaded=!1}}class ua extends t{constructor(g){super(g);this.options={rootMargin:"10%"},requestAnimationFrame(()=>{const[a,f]=da(this,this.options);Object.assign(this,{observe:a,unobserve:f})})}}var C={mediaQueries:{},dispatchEvent:!0,eventPrefix:!0,debug:!1};class o extends U{constructor(g,a={}){var f,h,k,j;super(g,a);if(this.matches=[],this.callback=(m)=>{const q=Object.keys(this.mediaQueries).find((A)=>this.mediaQueries[A]===m.media);if(!q)return;const{media:p,matches:y}=m;this.changed({name:q,media:p,matches:y,event:m})},this.changed=(m)=>{const{name:q}=m;if(m.event)this.call($(`${q}_changed`),m),this.dispatch(`${q}:changed`,m),this.log(`media query "${q}" changed`,m);if(m.matches)this.call($(`is_${q}`),m),this.dispatch(`is:${q}`,m);else this.call($(`not_${q}`),m),this.dispatch(`not:${q}`,m)},this.observe=()=>{Object.keys(this.mediaQueries).forEach((m)=>{const q=this.mediaQueries[m],p=window.matchMedia(q);p.addListener(this.callback),this.matches.push(p),this.changed({name:m,media:q,matches:p.matches})})},this.unobserve=()=>{this.matches.forEach((m)=>m.removeListener(this.callback))},this.controller=g,this.mediaQueries=(f=a.mediaQueries)!==null&&f!==void 0?f:C.mediaQueries,this.dispatchEvent=(h=a.dispatchEvent)!==null&&h!==void 0?h:C.dispatchEvent,this.eventPrefix=(k=a.eventPrefix)!==null&&k!==void 0?k:C.eventPrefix,this.debug=(j=a.debug)!==null&&j!==void 0?j:C.debug,!window.matchMedia){console.error("window.matchMedia() is not available");return}this.enhanceController(),this.observe()}enhanceController(){const g=this.controller.disconnect.bind(this.controller),a=()=>{this.unobserve(),g()};Object.assign(this.controller,{disconnect:a})}}var qf=(g,a={})=>{const f=new o(g,a);return[f.observe,f.unobserve]},wa=(g,a,f)=>{return Object.defineProperty(g,a,{value:f}),f},yf=(g)=>{var a;(a=g.constructor.memos)===null||a===void 0||a.forEach((f)=>{wa(g,f,g[f])})},za=(g,a,f)=>{const h=f?`${$(a)}Meta`:$(a);Object.defineProperty(g,h,{get(){return aa(e(a))}})},Af=(g,a={suffix:!0})=>{const f=g.constructor.metaNames,h=a.suffix;f===null||f===void 0||f.forEach((k)=>{za(g,k,h)}),Object.defineProperty(g,"metas",{get(){const k={};return f===null||f===void 0||f.forEach((j)=>{const m=aa(e(j));if(m!==void 0&&m!==null)k[$(j)]=m}),k}})};class fa extends U{constructor(g,a={}){super(g,a);this.observe=()=>{try{this.observer.observe(this.targetElement,this.options)}catch(f){this.controller.application.handleError(f,"At a minimum, one of childList, attributes, and/or characterData must be true",{})}},this.unobserve=()=>{this.observer.disconnect()},this.mutation=(f)=>{this.call("mutate",f),this.log("mutate",{entries:f}),this.dispatch("mutate",{entries:f})},this.targetElement=(a===null||a===void 0?void 0:a.element)||g.element,this.controller=g,this.options=a,this.observer=new MutationObserver(this.mutation),this.enhanceController(),this.observe()}enhanceController(){const g=this.controller.disconnect.bind(this.controller),a=()=>{this.unobserve(),g()};Object.assign(this.controller,{disconnect:a})}}var Ca=(g,a={})=>{const f=new fa(g,a);return[f.observe,f.unobserve]};class ga extends J{}class Ra extends ga{constructor(g){super(g);requestAnimationFrame(()=>{const[a,f]=Ca(this,this.options);Object.assign(this,{observe:a,unobserve:f})})}}var ba={dispatchEvent:!0,eventPrefix:!0},ca=(g,a={})=>{const f=g,{dispatchEvent:h,eventPrefix:k}=Object.assign({},ba,a),j=(a===null||a===void 0?void 0:a.element)||f.element,m=(B)=>{const[D]=B;if(N(f,"resize").call(f,D.contentRect),h){const E=Z("resize",f,k),I=d(E,null,{controller:f,entry:D});j.dispatchEvent(I)}},q=f.disconnect.bind(f),p=new ResizeObserver(m),y=()=>{p.observe(j)},A=()=>{p.unobserve(j)};return Object.assign(f,{disconnect(){A(),q()}}),y(),[y,A]};class ha extends J{}class ia extends ha{constructor(g){super(g);requestAnimationFrame(()=>{const[a,f]=ca(this,this.options);Object.assign(this,{observe:a,unobserve:f})})}}class ja extends U{constructor(g,a={}){super(g,a);this.observe=()=>{this.observer.observe(this.targetElement,{subtree:!0,characterData:!0,childList:!0,attributes:!0,attributeOldValue:!0,attributeFilter:[this.targetSelector,this.scopedTargetSelector]})},this.unobserve=()=>{this.observer.disconnect()},this.mutation=(f)=>{for(let h of f)switch(h.type){case"attributes":let k=h.target.getAttribute(h.attributeName),j=h.oldValue;if(h.attributeName===this.targetSelector||h.attributeName===this.scopedTargetSelector){let y=this.targetsUsedByThisController(j),A=this.targetsUsedByThisController(k),B=y.filter((E)=>!A.includes(E)),D=A.filter((E)=>!y.includes(E));B.forEach((E)=>this.targetRemoved(this.stripIdentifierPrefix(E),h.target,"attributeChange")),D.forEach((E)=>this.targetAdded(this.stripIdentifierPrefix(E),h.target,"attributeChange"))}break;case"characterData":let m=this.findTargetInAncestry(h.target);if(m==null)return;else this.targetsUsedByThisControllerFromNode(m).forEach((A)=>{this.targetChanged(this.stripIdentifierPrefix(A),m,"domMutation")});break;case"childList":let{addedNodes:q,removedNodes:p}=h;q.forEach((y)=>this.processNodeDOMMutation(y,this.targetAdded)),p.forEach((y)=>this.processNodeDOMMutation(y,this.targetRemoved));break}},this.controller=g,this.options=a,this.targetElement=g.element,this.identifier=g.scope.identifier,this.identifierPrefix=`${this.identifier}.`,this.targetSelector=g.scope.schema.targetAttribute,this.scopedTargetSelector=`data-${this.identifier}-target`,this.targets=a.targets||g.constructor.targets,this.prefixedTargets=this.targets.map((f)=>`${this.identifierPrefix}${f}`),this.observer=new MutationObserver(this.mutation),this.enhanceController(),this.observe()}processNodeDOMMutation(g,a){let f=g,h=a,k=[];if(f.nodeName=="#text"||this.targetsUsedByThisControllerFromNode(f).length==0)h=this.targetChanged,f=this.findTargetInAncestry(g);else k=this.targetsUsedByThisControllerFromNode(f);if(f==null)return;else if(k.length==0)k=this.targetsUsedByThisControllerFromNode(f);k.forEach((j)=>{h.call(this,this.stripIdentifierPrefix(j),f,"domMutation")})}findTargetInAncestry(g){let a=g,f=[];if(a.nodeName!="#text")f=this.targetsUsedByThisControllerFromNode(a);while(a.parentNode!==null&&a.parentNode!=this.targetElement&&f.length==0)if(a=a.parentNode,a.nodeName!=="#text"){if(this.targetsUsedByThisControllerFromNode(a).length>0)return a}if(a.nodeName=="#text")return null;if(a.parentNode==null)return null;if(a.parentNode==this.targetElement){if(this.targetsUsedByThisControllerFromNode(a).length>0)return a;return null}return null}targetAdded(g,a,f){let h=`${g}TargetAdded`;this.controller[h]&&N(this.controller,h).call(this.controller,a),this.log("targetAdded",{target:g,node:a,trigger:f})}targetRemoved(g,a,f){let h=`${g}TargetRemoved`;this.controller[h]&&N(this.controller,h).call(this.controller,a),this.log("targetRemoved",{target:g,node:a,trigger:f})}targetChanged(g,a,f){let h=`${g}TargetChanged`;this.controller[h]&&N(this.controller,h).call(this.controller,a),this.log("targetChanged",{target:g,node:a,trigger:f})}targetsUsedByThisControllerFromNode(g){if(g.nodeName=="#text"||g.nodeName=="#comment")return[];let a=g;return this.targetsUsedByThisController(a.getAttribute(this.scopedTargetSelector)||a.getAttribute(this.targetSelector))}targetsUsedByThisController(g){g=g||"";let a=this.stripIdentifierPrefix(g).split(" ");return this.targets.filter((f)=>a.indexOf(f)!==-1)}stripIdentifierPrefix(g){return g.replace(new RegExp(this.identifierPrefix,"g"),"")}enhanceController(){const g=this.controller.disconnect.bind(this.controller),a=()=>{this.unobserve(),g()};Object.assign(this.controller,{disconnect:a})}}var la=(g,a={})=>{const h=new ja(g,a);return[h.observe,h.unobserve]};class ka extends J{}class va extends ka{constructor(g){super(g);requestAnimationFrame(()=>{const[a,f]=la(this,this.options);Object.assign(this,{observe:a,unobserve:f})})}}class ma extends J{}ma.throttles=[];var xa=200,Bf=(g,a={})=>{var f;const h=g;(f=h.constructor.throttles)===null||f===void 0||f.forEach((j)=>{if(typeof j==="string")h[j]=i(h[j],a===null||a===void 0?void 0:a.wait);if(typeof j==="object"){const{name:m,wait:q}=j;if(!m)return;h[m]=i(h[m],q||(a===null||a===void 0?void 0:a.wait))}})},Wa={enterFromClass:"enter",enterActiveClass:"enterStart",enterToClass:"enterEnd",leaveFromClass:"leave",leaveActiveClass:"leaveStart",leaveToClass:"leaveEnd"},sa={transitioned:!1,hiddenClass:"hidden",preserveOriginalClass:!0,removeToClasses:!0},ra=(g,a={})=>{var f,h,k;const j=g,m=j.element.dataset.transitionTarget;let q;if(m)q=j[`${m}Target`];const p=(a===null||a===void 0?void 0:a.element)||q||j.element;if(!(p instanceof HTMLElement||p instanceof SVGElement))return;const y=p.dataset,A=parseInt(y.leaveAfter||"")||a.leaveAfter||0,{transitioned:B,hiddenClass:D,preserveOriginalClass:E,removeToClasses:I}=Object.assign({},sa,a),O=(f=j.enter)===null||f===void 0?void 0:f.bind(j),P=(h=j.leave)===null||h===void 0?void 0:h.bind(j),G=(k=j.toggleTransition)===null||k===void 0?void 0:k.bind(j);async function L(F){if(j.transitioned)return;j.transitioned=!0,O&&O(F);const K=T("enterFrom",a,y),V=T("enterActive",a,y),X=T("enterTo",a,y),Y=T("leaveTo",a,y);if(D)p.classList.remove(D);if(!I)_(p,Y);if(await S(p,K,V,X,D,E,I),A>0)setTimeout(()=>{Q(F)},A)}async function Q(F){if(!j.transitioned)return;j.transitioned=!1,P&&P(F);const K=T("leaveFrom",a,y),V=T("leaveActive",a,y),X=T("leaveTo",a,y),Y=T("enterTo",a,y);if(!I)_(p,Y);if(await S(p,K,V,X,D,E,I),D)p.classList.add(D)}function H(F){if(G&&G(F),j.transitioned)Q();else L()}async function S(F,K,V,X,Y,Ea,Fa){const w=[];if(Ea)K.forEach((M)=>F.classList.contains(M)&&M!==Y&&w.push(M)),V.forEach((M)=>F.classList.contains(M)&&M!==Y&&w.push(M)),X.forEach((M)=>F.classList.contains(M)&&M!==Y&&w.push(M));if(z(F,K),_(F,w),z(F,V),await ta(),_(F,K),z(F,X),await na(F),_(F,V),Fa)_(F,X);z(F,w)}function u(){if(j.transitioned=B,B){if(D)p.classList.remove(D);L()}else{if(D)p.classList.add(D);Q()}}function z(F,K){if(K.length>0)F.classList.add(...K)}function _(F,K){if(K.length>0)F.classList.remove(...K)}return u(),Object.assign(j,{enter:L,leave:Q,toggleTransition:H}),[L,Q,H]};class pa extends J{constructor(){super(...arguments);this.transitioned=!1}}class ea extends pa{constructor(g){super(g);requestAnimationFrame(()=>{ra(this,this.options)})}}class qa extends U{constructor(g,a={}){super(g,a);this.observe=()=>{this.controller.isVisible=!document.hidden,document.addEventListener("visibilitychange",this.handleVisibilityChange),this.handleVisibilityChange()},this.unobserve=()=>{document.removeEventListener("visibilitychange",this.handleVisibilityChange)},this.becomesInvisible=(f)=>{this.controller.isVisible=!1,this.call("invisible",f),this.log("invisible",{isVisible:!1}),this.dispatch("invisible",{event:f,isVisible:!1})},this.becomesVisible=(f)=>{this.controller.isVisible=!0,this.call("visible",f),this.log("visible",{isVisible:!0}),this.dispatch("visible",{event:f,isVisible:!0})},this.handleVisibilityChange=(f)=>{if(document.hidden)this.becomesInvisible(f);else this.becomesVisible(f)},this.controller=g,this.enhanceController(),this.observe()}enhanceController(){const g=this.controllerDisconnect,a=()=>{this.unobserve(),g()};Object.assign(this.controller,{disconnect:a})}}var af=(g,a={})=>{const h=new qa(g,a);return[h.observe,h.unobserve]};class ya extends J{constructor(){super(...arguments);this.isVisible=!1}}class ff extends ya{constructor(g){super(g);requestAnimationFrame(()=>{const[a,f]=af(this,this.options);Object.assign(this,{observe:a,unobserve:f})})}}class Aa extends U{constructor(g,a={}){super(g,a);this.observe=()=>{if(document.hasFocus())this.becomesFocused();else this.becomesUnfocused();this.interval=setInterval(()=>{this.handleWindowFocusChange()},this.intervalDuration)},this.unobserve=()=>{clearInterval(this.interval)},this.becomesUnfocused=(f)=>{this.controller.hasFocus=!1,this.call("unfocus",f),this.log("unfocus",{hasFocus:!1}),this.dispatch("unfocus",{event:f,hasFocus:!1})},this.becomesFocused=(f)=>{this.controller.hasFocus=!0,this.call("focus",f),this.log("focus",{hasFocus:!0}),this.dispatch("focus",{event:f,hasFocus:!0})},this.handleWindowFocusChange=(f)=>{if(document.hasFocus()&&!this.controller.hasFocus)this.becomesFocused(f);else if(!document.hasFocus()&&this.controller.hasFocus)this.becomesUnfocused(f)},this.controller=g,this.intervalDuration=a.interval||200,this.enhanceController(),this.observe()}enhanceController(){const g=this.controllerDisconnect,a=()=>{this.unobserve(),g()};Object.assign(this.controller,{disconnect:a})}}var gf=(g,a={})=>{const h=new Aa(g,a);return[h.observe,h.unobserve]};class Ba extends J{constructor(){super(...arguments);this.hasFocus=!1}}class hf extends Ba{constructor(g){super(g);requestAnimationFrame(()=>{const[a,f]=gf(this,this.options);Object.assign(this,{observe:a,unobserve:f})})}}var jf=(g)=>{const a=g,f=(m)=>{const{innerWidth:q,innerHeight:p}=window,y={height:p||Infinity,width:q||Infinity,event:m};N(a,"windowResize").call(a,y)},h=a.disconnect.bind(a),k=()=>{window.addEventListener("resize",f),f()},j=()=>{window.removeEventListener("resize",f)};return Object.assign(a,{disconnect(){j(),h()}}),k(),[k,j]};class Da extends J{}class kf extends Da{constructor(g){super(g);requestAnimationFrame(()=>{const[a,f]=jf(this);Object.assign(this,{observe:a,unobserve:f})})}}export{jf as useWindowResize,gf as useWindowFocus,af as useVisibility,ra as useTransition,Bf as useThrottle,la as useTargetMutation,ca as useResize,Ca as useMutation,Af as useMeta,yf as useMemo,qf as useMatchMedia,da as useLazyLoad,_a as useIntersection,Xa as useIdle,Qa as useHover,Df as useHotkeys,Ia as useDispatch,pf as useDebounce,Na as useClickOutside,Ka as useApplication,c as debounce,kf as WindowResizeController,hf as WindowFocusController,ff as VisibilityController,Aa as UseWindowFocus,qa as UseVisibility,ja as UseTargetMutation,fa as UseMutation,W as UseHover,ea as TransitionController,va as TargetMutationController,ia as ResizeController,Ra as MutationController,ua as LazyLoadController,$a as IntersectionController,Ya as IdleController,Sa as HoverController,Oa as ClickOutsideController,La as ApplicationController};

//# debugId=E15C6420B4535B3764756e2164756e21
