; Segment type: Pure code
; Segment permissions: Read/Execute
__text segment para public 'CODE' use64
assume cs:__text
;org 100003F90h
assume es:nothing, ss:nothing, ds:nothing, fs:nothing, gs:nothing


; Attributes: bp-based frame

; int __cdecl main(int argc, const char **argv, const char **envp)
public _main
_main proc near

var_C= dword ptr -0Ch
var_8= dword ptr -8
var_4= dword ptr -4

push    rbp
mov     rbp, rsp
xor     eax, eax
mov     [rbp+var_4], 1
mov     [rbp+var_8], 2
mov     ecx, [rbp+var_4]
add     ecx, [rbp+var_8]
mov     [rbp+var_C], ecx
pop     rbp
retn
_main endp
