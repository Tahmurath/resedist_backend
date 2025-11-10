import{j as e,N as o,O as r}from"./index.js";import{T as c}from"./theme-provider.js";import{u as n}from"./useTranslation.js";import{c as s}from"./createLucideIcon.js";/**
 * @license lucide-react v0.513.0 - ISC
 *
 * This source code is licensed under the ISC license.
 * See the LICENSE file in the root directory of this source tree.
 */const m=[["path",{d:"M15 21v-8a1 1 0 0 0-1-1h-4a1 1 0 0 0-1 1v8",key:"5wwlr5"}],["path",{d:"M3 10a2 2 0 0 1 .709-1.528l7-5.999a2 2 0 0 1 2.582 0l7 5.999A2 2 0 0 1 21 10v9a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z",key:"1d0kgt"}]],d=s("house",m);/**
 * @license lucide-react v0.513.0 - ISC
 *
 * This source code is licensed under the ISC license.
 * See the LICENSE file in the root directory of this source tree.
 */const h=[["path",{d:"M12.22 2h-.44a2 2 0 0 0-2 2v.18a2 2 0 0 1-1 1.73l-.43.25a2 2 0 0 1-2 0l-.15-.08a2 2 0 0 0-2.73.73l-.22.38a2 2 0 0 0 .73 2.73l.15.1a2 2 0 0 1 1 1.72v.51a2 2 0 0 1-1 1.74l-.15.09a2 2 0 0 0-.73 2.73l.22.38a2 2 0 0 0 2.73.73l.15-.08a2 2 0 0 1 2 0l.43.25a2 2 0 0 1 1 1.73V20a2 2 0 0 0 2 2h.44a2 2 0 0 0 2-2v-.18a2 2 0 0 1 1-1.73l.43-.25a2 2 0 0 1 2 0l.15.08a2 2 0 0 0 2.73-.73l.22-.39a2 2 0 0 0-.73-2.73l-.15-.08a2 2 0 0 1-1-1.74v-.5a2 2 0 0 1 1-1.74l.15-.09a2 2 0 0 0 .73-2.73l-.22-.38a2 2 0 0 0-2.73-.73l-.15.08a2 2 0 0 1-2 0l-.43-.25a2 2 0 0 1-1-1.73V4a2 2 0 0 0-2-2z",key:"1qme2f"}],["circle",{cx:"12",cy:"12",r:"3",key:"1v7zrd"}]],x=s("settings",h);/**
 * @license lucide-react v0.513.0 - ISC
 *
 * This source code is licensed under the ISC license.
 * See the LICENSE file in the root directory of this source tree.
 */const p=[["path",{d:"M19 21v-2a4 4 0 0 0-4-4H9a4 4 0 0 0-4 4v2",key:"975kel"}],["circle",{cx:"12",cy:"7",r:"4",key:"17ys0d"}]],f=s("user",p),u=()=>{const{t}=n(),l=[{path:"/tg/miniapp/in/home",label:t("site.home"),icon:e.jsx(d,{className:"h-6 w-6"})},{path:"/tg/miniapp/in/profile",label:t("site.profile"),icon:e.jsx(f,{className:"h-6 w-6"})},{path:"/tg/miniapp/in/settings",label:t("site.settings"),icon:e.jsx(x,{className:"h-6 w-6"})}];return e.jsx("nav",{className:"fixed bottom-0 left-0 right-0 bg-background border-t border-border ",children:e.jsx("div",{className:"flex justify-around items-center h-16",children:l.map(a=>e.jsxs(o,{to:a.path,className:({isActive:i})=>`flex flex-col items-center justify-center flex-1 text-muted-foreground ${i?"text-primary":""}`,children:[a.icon,e.jsx("span",{className:"text-xs mt-1",children:a.label})]},a.path))})})},N=()=>e.jsx(c,{defaultTheme:"dark",storageKey:"vite-ui-theme",children:e.jsx("div",{className:"bg-muted flex min-h-svh flex-col items-center justify-center gap-6 p-6 md:p-10",children:e.jsxs("div",{className:"flex w-full max-w-sm flex-col gap-6",children:[e.jsx(u,{}),e.jsx(r,{})]})})});export{N as default};
