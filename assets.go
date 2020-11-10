// Code generated by staticFile; DO NOT EDIT.
// source files: ["front/avatar.webp" "front/index.html" "front/app.js" "front/r2d2.webp" "front/style.css"]
// dev mode: false

package main

func FrontAvatar() []byte { return []byte("RIFF\f\x01\x00\x00WEBPVP8 \x00\x01\x00\x00\x10\x15\x00\x9d\x01*\xe0\x00\xe0\x00>\x91F\x9cI%\xa4#!+}\xb8\x00\xb0\x12\ten\xe1o\x0e\x00\xfe\x00\xd9\xf4\x0eӟ\xf6\r[\x1e\xc1S\xf8\r\xbbH\x8f\xe4\xba\x13ͫ\xcd\x04\u007f%Оl\xe5<$\b\x1b\xafY\x16\xc8\xe4\xe0Њ\xac:C\b\\\x805\xa8\xdc/*l\x00>B\aCsj\xbf\xd6\r\x91\xc9\xc1\t\x98vլ&\aCs2;c\xed\xb3\xa7Y#7t\x99\x02\\\xaf0\xb4̨\u0099\x8f\xea?\x92\xe8O6\xaf3\xdf{\xd6A\xae\xee)4\x1e\xd5\rBny\xf5\xe6\xd5\xe0\\m\xa1\x1c\x19^\x14\xc2r\x00\x81\xe8\xe1A#$G\xf2]\t0\x00\xfb\xb5_\xde_/\xa3\xec\n\x1f\x12\x8a¾\x8a\x10\xedW]\x99\x8bc`>\xe2Ŋ\xa1\x8b.\xccű\xb0\x1fqb\xc5\x00\x01\xc8ʁ\x9a\xdf\xfdc\nk\x15P\x9a<G\x19\xf3\xa9\xc0wC\xf8\aD\xed\xd7\x04QX\x06\xad\x1e?U\xf0\x00\x00\x00\x00") }

func FrontIndex() []byte { return []byte("<!doctype html><html lang=fr><head><meta charset=utf-8><link rel=stylesheet href=/style.css><script src=/index.js></script><title>Arveto Auth Dev</title></head><body><h1>Arveto Auth Dev</h1><p>Ce site fournis des JWT de faux utilisateurs. <strong>À\nutiliser uniquement pour le dévelopement&#8239;!</strong></p><input type=url id=u value=http://localhost:8000 placeholder=\"URL de l'application\"><br><input id=r placeholder=\"URL de redirection dans l'application\"><br><input id=app placeholder=\"ID de l'application\"><br><input id=teams placeholder=\"Équipes (ca, dev)\"><br><hr><input type=checkbox id=gen><label for=gen>Bot</label><br><button type=button id=go>GO!</button><hr><ul id=list><li><button type=button class=inline id=tomas>GENERATE</button><a id=tomas>Tomas (Visiteur)</a><li><button type=button class=inline id=zech>GENERATE</button><a id=zech>Zech (Standard)</a><li><button type=button class=inline id=tzu>GENERATE</button><a id=tzu>Tzu Shikaï (Administrateur)</a><li><button type=button class=inline id=r2d2>GENERATE</button><a id=r2d2>R2D2 (Standard, Bot)</a></ul><output name=result id=out hidden>JWT</output><footer><a href=https://github.com/Arveto/arvetoAuthDev>https://github.com/Arveto/arvetoAuthDev</a><br>Copyright (c) 2020, Arveto Ink. All rights reserved.<br>Use of this source code is governed by a BSD<br>license that can be found in the LICENSE file.</footer></body></html>") }

func FrontApp() []byte { return []byte("const stdsubjets={'tomas':{id:'tomas',pseudo:'Tomas Benedito Bonito',email:'tomas@fai.mil',level:'visitor',bot:false,},'zech':{id:'zech',pseudo:'Zech',email:'zech@fai.mil',level:'standard',bot:false,},'tzu':{id:'tzu',pseudo:'Tzu Shikaï',email:'tzu@fai.mil',level:'administrator',bot:false,},'r2d2':{id:'bot-r2d2',pseudo:'R2D2',email:'r2d2@bot',level:'standard',bot:true,},};document.addEventListener('DOMContentLoaded',()=>{document.querySelectorAll('#list button').forEach(b=>{const s=stdsubjets[b.id];b.addEventListener('click',()=>{s.teams=getTeams();s.app=getApp();generate(s);});});const search=new URLSearchParams(document.location.search);const app=search.get('app');if(app)document.getElementById('app').value=app;const r=search.get('r');if(r)document.getElementById('r').value=r;let updateAppID=[];let updateAppURL=[];let updateRedirect=[];let updateTeams=[];document.querySelectorAll('#list a').forEach(a=>{a.addEventListener('click',event=>{event.preventDefault();window.open(a.href).focus();});let s=stdsubjets[a.id];function update(){a.href='/redirect?'+new URLSearchParams(s);}\nfunction updater(ar,f){ar.push(v=>{f(v);update();})}\nupdater(updateAppID,id=>s.app=id);updater(updateAppURL,u=>s.u=u);updater(updateRedirect,r=>s.r=r);updater(updateTeams,t=>s.teams=t);});function udpatingTeams(){const teams=getTeams();updateTeams.forEach(u=>u(teams));}\nudpatingTeams();document.getElementById('teams').addEventListener('input',udpatingTeams);function updator(array,id){const input=document.getElementById(id);function up(){const v=input.value;array.forEach(u=>u(v));}\ninput.addEventListener('input',up);up();}\nupdator(updateAppID,'app');updator(updateAppURL,'u');updator(updateRedirect,'r');},{once:true,});function getTeams(){return document.getElementById('teams').value.split(',').map(t=>t.trim()).filter(t=>t.length).sort();}\nfunction getApp(){return document.getElementById('app').value;}\nfunction redirectURL(query){return '/redirect'+new URLSearchParams(query);}\nasync function generate(query){const out=document.getElementById('out');out.value='loading ...';out.hidden=false;const jwt=await fetch('/generate?'+new URLSearchParams(query)).then(rep=>rep.text()).then(text=>text.trim());out.value=jwt;out.onclick=()=>{const i=document.createElement('input');document.body.appendChild(i);i.type='text';i.value=jwt;i.select();document.execCommand('copy');i.remove();out.classList.add('copied');setTimeout(()=>out.classList.remove('copied'),500);};}") }

func FrontR2d2() []byte { return []byte("RIFF\x9c\x01\x00\x00WEBPVP8 \x90\x01\x00\x00p\x18\x00\x9d\x01*\xe0\x00\xe0\x00>\x91@\x99H%\xa4\"\xa1/\x1ah\x00\xb0\x12\tcn\xe0\xc0\x03\xf8\a\xe0\a\xe17\xe6\xcb\xf9P\xc4\xeb\xf8\xb7\xe3O\xf0\r\x9e\x8e-\xf8\xd3\xfc\x03܊E\xbe\xf4\xe6~\x04#Aqo\x1c\xa9\xb1\xf7,\xe1\x1b\x83\xf8.6>\xa2῟YnF\x80\xe2l\x92\xa94\xc8#.\vD\xe3Dj\x98\xf4Q\x01\x16_\x84\xed\x87\xce*\xe6%C\\ \xe0T\xb2\v\xb5m\xbf\x86\xb2\xc1&x>jpqH6\xab\x80\x17\xde\x19A\x9a9\x89\xab\xbf\xaf\xb2\xb1!\xaa${g\xc0\xc1\xe7\xa8\xe4\f\xc8\xd4\xe8\x1ff\xd1t\x1a\xe9\x9c+/^\x149W1*8\xcf\xf0[+\x87\xcdM\u007f\xfd-\xcb\x16\x86E\xe5\xf7\"Ϩ\xb2\xfc'm\x16\x00\x00\xec\xed\x01}q\xc9G\x0er3M\x88&%\xc6\xc0\x00\xedҺl\xdc\xfcNFi\xadk\x10\xbec\xee1\x1c\xa5D\xa3\x8a\x80\xd7\x1c\x94p\xe7#4\xd5gc\xab\x8c\xb5\xa8\x9e\xd3٧\xb3N\x9cbs\xb3\xf2j\xed\xedU \x80\x157a*<!\xdbt\xf7M\xbc\x01ζ\x87\xaa\x96\xf3\x99k\\rQò\xccH.\xd8Y\xe1*<!\xdbt\xf7M\xbc\x00%\xdd7\xa7&|\xac1\xa8\xfb\x10\v|\xbdߴٹ\xf8\xb2\xccH1U(\xf7\xc7\\Z\x86\x16\u070e\xc5!\xbc&\xb4X\x06\n^\x86\x16\u070e\xfd\x03\xae\x96\x896\xdc\xdf1\xea\xf9\xf3\x97\xa3/9z\xb1O,p\xb40\xb6\xe4v)\r\xdf\xf9>V\xc08\xd6\xca0\x00\x00") }

func FrontStyle() []byte { return []byte(":root{--bg:hsla(0 0% 20%);--bg-primary:hsla(0 0% 10%);--bg-secondary:hsla(0 0% 15%);--cline:hsla(0 0% 30%);--cline-focus:dodgerblue;--cline-invalid:#FB5200;--w:70ex;--w-large:100ex;--l1:0.5ex;--l2:0.1ex;--sp:2ex}*[hidden]{display:none!important}body{max-width:var(--w);margin:auto;padding:var(--sp);font-size:large;font-family:sans-serif;background:var(--bg);color:#fff}footer{margin-top:var(--sp);border-top:solid var(--cline)var(--l1);padding-top:var(--l1);font-style:italic}hr{border:none;border-bottom:solid var(--cline)var(--l1)}a{color:#1e90ff}input:not([type=checkbox]){display:block;width:100%;margin:var(--l2)0;border:var(--l1)var(--cline);border-style:solid none;padding:1ex;font-size:inherit;font-family:inherit;color:inherit;background-color:var(--bg-secondary)}button{margin:var(--sp)0;border:var(--l1)var(--cline)solid;padding:var(--l1)var(--sp);font-size:inherit;font-family:inherit;color:inherit;background-color:var(--bg-secondary)}button:hover,input:focus,input:hover{border-color:var(--cline-focus);color:var(--cline-focus);border-radius:0}button.inline{margin:0 var(--l1);border:solid var(--l2)var(--cline);border-radius:var(--l1);padding:var(--l2)var(--l1);font-size:smaller}#out{display:block;padding:var(--sp);word-break:break-all;font-family:monospace;background-color:var(--bg-secondary)}#out.copied{border-color:green;color:green}") }

