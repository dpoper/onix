(window.webpackJsonp=window.webpackJsonp||[]).push([[7],{107:function(e,t,a){"use strict";a.r(t),a.d(t,"frontMatter",(function(){return o})),a.d(t,"metadata",(function(){return c})),a.d(t,"rightToc",(function(){return l})),a.d(t,"default",(function(){return p}));var n=a(1),r=a(6),i=(a(0),a(136)),o={id:"concept_roles",title:"Roles"},c={id:"concept_roles",title:"Roles",description:"A role is a group of [privileges](/onix/docs/concept_roles#privileges) that can be assigned to [users](/onix/docs/concept_users) to *create*, *list* and/or *delete* configuration data within ",source:"@site/docs/concept_roles.md",permalink:"/onix/docs/concept_roles",editUrl:"https://github.com/facebook/docusaurus/edit/master/website/docs/concept_roles.md",sidebar:"someSidebar",previous:{title:"Users",permalink:"/onix/docs/concept_users"}},l=[{value:"Partitions",id:"partitions",children:[{value:"Default partitions",id:"default-partitions",children:[]}]},{value:"Administrator roles",id:"administrator-roles",children:[]},{value:"Privileges",id:"privileges",children:[]}],b={rightToc:l};function p(e){var t=e.components,a=Object(r.a)(e,["components"]);return Object(i.b)("wrapper",Object(n.a)({},b,a,{components:t,mdxType:"MDXLayout"}),Object(i.b)("p",null,"A role is a group of ",Object(i.b)("a",Object(n.a)({parentName:"p"},{href:"/onix/docs/concept_roles#privileges"}),"privileges")," that can be assigned to ",Object(i.b)("a",Object(n.a)({parentName:"p"},{href:"/onix/docs/concept_users"}),"users")," to ",Object(i.b)("em",{parentName:"p"},"create"),", ",Object(i.b)("em",{parentName:"p"},"list")," and/or ",Object(i.b)("em",{parentName:"p"},"delete")," configuration data within\na ",Object(i.b)("a",Object(n.a)({parentName:"p"},{href:"/onix/docs/concept_roles#partitions"}),"partition"),".\nPartitions and privileges are explained below."),Object(i.b)("p",null,"When a user logs in, Onix obtains a list of roles that have been assigned to the user and pass them to its data service."),Object(i.b)("p",null,"The data service ",Object(i.b)("em",{parentName:"p"},"only")," creates, lists or deletes data to which the user has been granted privileges - through their role(s)."),Object(i.b)("p",null,"In order to apply privileges to ",Object(i.b)("em",{parentName:"p"},"only certain portions of the data stored in Onix"),", it is necessary to partition the information    as explained below."),Object(i.b)("h2",{id:"partitions"},"Partitions"),Object(i.b)("p",null,"Partitions are a logical way to divide the Onix data into areas so that privileges can be given to access such areas\nindependently. "),Object(i.b)("p",null,"For example, suppose that a specific group within an organisation wants to be able to create and manage configuration\ndata that is not supposed to be seen or managed by another group within the organisation. In this case, a partition is\ncreated for such group, and privileges are given to a role to ",Object(i.b)("em",{parentName:"p"},"Create, Read and / or Delete")," data in the partition."),Object(i.b)("p",null,"Once users are ",Object(i.b)("a",Object(n.a)({parentName:"p"},{href:"/onix/docs/concept_users#memberships"}),"members")," of the role, they inherit those privileges."),Object(i.b)("h3",{id:"default-partitions"},"Default partitions"),Object(i.b)("p",null,"In order to facilitate getting started and to provide the simplest access controls possible, Onix ships out of the box with\ntwo default partitions:"),Object(i.b)("table",null,Object(i.b)("thead",{parentName:"table"},Object(i.b)("tr",{parentName:"thead"},Object(i.b)("th",Object(n.a)({parentName:"tr"},{align:null}),"key"),Object(i.b)("th",Object(n.a)({parentName:"tr"},{align:null}),"partition"))),Object(i.b)("tbody",{parentName:"table"},Object(i.b)("tr",{parentName:"tbody"},Object(i.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"REF"),Object(i.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"The default ",Object(i.b)("em",{parentName:"td"},"reference data")," partition. By default, any reference data in Onix is put into the REF partition. Reference data is any data that is part of configuration models (i.e. models, item types, link types and link rules).")),Object(i.b)("tr",{parentName:"tbody"},Object(i.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"INS"),Object(i.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"The default ",Object(i.b)("em",{parentName:"td"},"instance data")," partition. By default, any instance data in Onix is put into the INS partition. Instance data is comprised by items and links. Items and links are explained in later sections.")))),Object(i.b)("p",null,"The following figure shows the two types of partitions (i.e. reference and instance) as well as the data entities to which\nthey apply:"),Object(i.b)("p",null,Object(i.b)("img",Object(n.a)({parentName:"p"},{src:"/onix/img/concept_partitions.png",alt:"partitions"}))),Object(i.b)("h2",{id:"administrator-roles"},"Administrator roles"),Object(i.b)("p",null,"In addition to the default partitions, administrators can create and update partitions."),Object(i.b)("p",null,"Administrator roles can create or update partitions.\nThey are differentiated from ordinary roles by their ",Object(i.b)("em",{parentName:"p"},"level")," which is either 1 or 2."),Object(i.b)("p",null,"The following role levels are possible:"),Object(i.b)("table",null,Object(i.b)("thead",{parentName:"table"},Object(i.b)("tr",{parentName:"thead"},Object(i.b)("th",Object(n.a)({parentName:"tr"},{align:"center"}),"role level"),Object(i.b)("th",Object(n.a)({parentName:"tr"},{align:null}),"type"),Object(i.b)("th",Object(n.a)({parentName:"tr"},{align:null}),"description"))),Object(i.b)("tbody",{parentName:"table"},Object(i.b)("tr",{parentName:"tbody"},Object(i.b)("td",Object(n.a)({parentName:"tr"},{align:"center"}),"0"),Object(i.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"ordinary user"),Object(i.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"cannot manage partitions")),Object(i.b)("tr",{parentName:"tbody"},Object(i.b)("td",Object(n.a)({parentName:"tr"},{align:"center"}),"1"),Object(i.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"local administrator"),Object(i.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"can only manage the partitions they create")),Object(i.b)("tr",{parentName:"tbody"},Object(i.b)("td",Object(n.a)({parentName:"tr"},{align:"center"}),"2"),Object(i.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"super administrator"),Object(i.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"can manage all partitions irrespective of who have created them")))),Object(i.b)("h2",{id:"privileges"},"Privileges"),Object(i.b)("p",null,"A privilege is a permission for a user to perform an operation on a piece of configuration data."),Object(i.b)("p",null,"There are three types of operations for which a User can be granted privileged:"),Object(i.b)("table",null,Object(i.b)("thead",{parentName:"table"},Object(i.b)("tr",{parentName:"thead"},Object(i.b)("th",Object(n.a)({parentName:"tr"},{align:"center"}),"privilege operation"),Object(i.b)("th",Object(n.a)({parentName:"tr"},{align:null}),"description"))),Object(i.b)("tbody",{parentName:"table"},Object(i.b)("tr",{parentName:"tbody"},Object(i.b)("td",Object(n.a)({parentName:"tr"},{align:"center"}),"CREATE"),Object(i.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"allows users to create or update data. Onix operations are ",Object(i.b)("a",Object(n.a)({parentName:"td"},{href:"https://en.wikipedia.org/wiki/Idempotence"}),"idempotent"),". No distinction is made between create and update operations.")),Object(i.b)("tr",{parentName:"tbody"},Object(i.b)("td",Object(n.a)({parentName:"tr"},{align:"center"}),"READ"),Object(i.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"allows users to list data.")),Object(i.b)("tr",{parentName:"tbody"},Object(i.b)("td",Object(n.a)({parentName:"tr"},{align:"center"}),"DELETE"),Object(i.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"allows users to delete data.")))),Object(i.b)("p",null,Object(i.b)("img",Object(n.a)({parentName:"p"},{src:"/onix/img/concept_users_to_parts.png",alt:"fro users to partitions"}))))}p.isMDXComponent=!0},136:function(e,t,a){"use strict";a.d(t,"a",(function(){return s})),a.d(t,"b",(function(){return O}));var n=a(0),r=a.n(n);function i(e,t,a){return t in e?Object.defineProperty(e,t,{value:a,enumerable:!0,configurable:!0,writable:!0}):e[t]=a,e}function o(e,t){var a=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),a.push.apply(a,n)}return a}function c(e){for(var t=1;t<arguments.length;t++){var a=null!=arguments[t]?arguments[t]:{};t%2?o(Object(a),!0).forEach((function(t){i(e,t,a[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(a)):o(Object(a)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(a,t))}))}return e}function l(e,t){if(null==e)return{};var a,n,r=function(e,t){if(null==e)return{};var a,n,r={},i=Object.keys(e);for(n=0;n<i.length;n++)a=i[n],t.indexOf(a)>=0||(r[a]=e[a]);return r}(e,t);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(n=0;n<i.length;n++)a=i[n],t.indexOf(a)>=0||Object.prototype.propertyIsEnumerable.call(e,a)&&(r[a]=e[a])}return r}var b=r.a.createContext({}),p=function(e){var t=r.a.useContext(b),a=t;return e&&(a="function"==typeof e?e(t):c({},t,{},e)),a},s=function(e){var t=p(e.components);return r.a.createElement(b.Provider,{value:t},e.children)},d={inlineCode:"code",wrapper:function(e){var t=e.children;return r.a.createElement(r.a.Fragment,{},t)}},u=Object(n.forwardRef)((function(e,t){var a=e.components,n=e.mdxType,i=e.originalType,o=e.parentName,b=l(e,["components","mdxType","originalType","parentName"]),s=p(a),u=n,O=s["".concat(o,".").concat(u)]||s[u]||d[u]||i;return a?r.a.createElement(O,c({ref:t},b,{components:a})):r.a.createElement(O,c({ref:t},b))}));function O(e,t){var a=arguments,n=t&&t.mdxType;if("string"==typeof e||n){var i=a.length,o=new Array(i);o[0]=u;var c={};for(var l in t)hasOwnProperty.call(t,l)&&(c[l]=t[l]);c.originalType=e,c.mdxType="string"==typeof e?e:n,o[1]=c;for(var b=2;b<i;b++)o[b]=a[b];return r.a.createElement.apply(null,o)}return r.a.createElement.apply(null,a)}u.displayName="MDXCreateElement"}}]);