/** Notice * This file contains works from many authors under various (but compatible) licenses. Please see core.txt for more information. **/
(function(){(window.wpCoreControlsBundle=window.wpCoreControlsBundle||[]).push([[22],{450:function(ia){(function(){ia.exports={XW:function(){function ba(e,f){this.scrollLeft=e;this.scrollTop=f}function e(e){if(null===e||"object"!==typeof e||void 0===e.behavior||"auto"===e.behavior||"instant"===e.behavior)return!0;if("object"===typeof e&&"smooth"===e.behavior)return!1;throw new TypeError("behavior member of ScrollOptions "+e.behavior+" is not a valid value for enumeration ScrollBehavior.");}function ea(e,h){if("Y"===
h)return e.clientHeight+f<e.scrollHeight;if("X"===h)return e.clientWidth+f<e.scrollWidth}function x(e,f){e=aa.getComputedStyle(e,null)["overflow"+f];return"auto"===e||"scroll"===e}function ha(e){var f=ea(e,"Y")&&x(e,"Y");e=ea(e,"X")&&x(e,"X");return f||e}function ca(e){var f=(h()-e.startTime)/468;var n=.5*(1-Math.cos(Math.PI*(1<f?1:f)));f=e.xy+(e.x-e.xy)*n;n=e.zy+(e.y-e.zy)*n;e.method.call(e.qE,f,n);f===e.x&&n===e.y||aa.requestAnimationFrame(ca.bind(aa,e))}function da(e,f,x){var r=h();if(e===w.body){var y=
aa;var z=aa.scrollX||aa.pageXOffset;e=aa.scrollY||aa.pageYOffset;var da=n.scroll}else y=e,z=e.scrollLeft,e=e.scrollTop,da=ba;ca({qE:y,method:da,startTime:r,xy:z,zy:e,x:f,y:x})}var aa=window,w=document;if(!("scrollBehavior"in w.documentElement.style&&!0!==aa.ona)){var y=aa.HTMLElement||aa.Element,n={scroll:aa.scroll||aa.scrollTo,scrollBy:aa.scrollBy,qS:y.prototype.scroll||ba,scrollIntoView:y.prototype.scrollIntoView},h=aa.performance&&aa.performance.now?aa.performance.now.bind(aa.performance):Date.now,
f=/MSIE |Trident\/|Edge\//.test(aa.navigator.userAgent)?1:0;aa.scroll=aa.scrollTo=function(f,h){void 0!==f&&(!0===e(f)?n.scroll.call(aa,void 0!==f.left?f.left:"object"!==typeof f?f:aa.scrollX||aa.pageXOffset,void 0!==f.top?f.top:void 0!==h?h:aa.scrollY||aa.pageYOffset):da.call(aa,w.body,void 0!==f.left?~~f.left:aa.scrollX||aa.pageXOffset,void 0!==f.top?~~f.top:aa.scrollY||aa.pageYOffset))};aa.scrollBy=function(f,h){void 0!==f&&(e(f)?n.scrollBy.call(aa,void 0!==f.left?f.left:"object"!==typeof f?f:
0,void 0!==f.top?f.top:void 0!==h?h:0):da.call(aa,w.body,~~f.left+(aa.scrollX||aa.pageXOffset),~~f.top+(aa.scrollY||aa.pageYOffset)))};y.prototype.scroll=y.prototype.scrollTo=function(f,h){if(void 0!==f)if(!0===e(f)){if("number"===typeof f&&void 0===h)throw new SyntaxError("Value could not be converted");n.qS.call(this,void 0!==f.left?~~f.left:"object"!==typeof f?~~f:this.scrollLeft,void 0!==f.top?~~f.top:void 0!==h?~~h:this.scrollTop)}else h=f.left,f=f.top,da.call(this,this,"undefined"===typeof h?
this.scrollLeft:~~h,"undefined"===typeof f?this.scrollTop:~~f)};y.prototype.scrollBy=function(f,h){void 0!==f&&(!0===e(f)?n.qS.call(this,void 0!==f.left?~~f.left+this.scrollLeft:~~f+this.scrollLeft,void 0!==f.top?~~f.top+this.scrollTop:~~h+this.scrollTop):this.scroll({left:~~f.left+this.scrollLeft,top:~~f.top+this.scrollTop,behavior:f.behavior}))};y.prototype.scrollIntoView=function(f){if(!0===e(f))n.scrollIntoView.call(this,void 0===f?!0:f);else{for(f=this;f!==w.body&&!1===ha(f);)f=f.parentNode||
f.host;var h=f.getBoundingClientRect(),r=this.getBoundingClientRect();f!==w.body?(da.call(this,f,f.scrollLeft+r.left-h.left,f.scrollTop+r.top-h.top),"fixed"!==aa.getComputedStyle(f).position&&aa.scrollBy({left:h.left,top:h.top,behavior:"smooth"})):aa.scrollBy({left:r.left,top:r.top,behavior:"smooth"})}}}}}})()}}]);}).call(this || window)
