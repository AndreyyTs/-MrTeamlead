// +build amd64

#include "textflag.h"

// func SumSIMD(arr []Foo) uint64
TEXT ·SumSIMD(SB), $0-32
    MOVQ arr_base+0(FP),    SI  
    MOVQ arr_len+8(FP),     CX  
    XORQ BX, BX

    CMPQ CX, $1                 
    JLE  done                                 
    
    XORQ AX, AX      

    VPXORD Z0, Z0, Z0        // аккумулятор 1
    VPXORD Z1, Z1, Z1        // аккумулятор 2
    VPXORD Z2, Z2, Z2        // аккумулятор 3
    VPXORD Z3, Z3, Z3        // аккумулятор 4

avx_loop:   
    VMOVDQU64 (SI), Z4
    VMOVDQU64 64(SI), Z5
    VMOVDQU64 128(SI), Z6
    VMOVDQU64 192(SI), Z7
    
    VPADDQ Z4, Z0, Z0
    VPADDQ Z5, Z1, Z1
    VPADDQ Z6, Z2, Z2
    VPADDQ Z7, Z3, Z3


    ADDQ $16, AX                
    ADDQ $256, SI 
    CMPQ AX, CX                
    JL   avx_loop


    VPADDQ Z0, Z3, Z3
    VPADDQ Z1, Z3, Z3
    VPADDQ Z2, Z3, Z3             


    VEXTRACTI64X4 $0, Z3, Y1      
    VEXTRACTI128 $0, Y1, X2      
    VMOVQ X2, AX      

    VEXTRACTI64X4 $1, Z3, Y1    
    VEXTRACTI128 $0, Y1, X2      
    VMOVQ X2, BX      
    ADDQ  BX, AX

    VEXTRACTI64X4 $0, Z3, Y1     
    VEXTRACTI128 $1, Y1, X2       
    VMOVQ X2, BX      
    ADDQ  BX, AX

    VEXTRACTI64X4 $1, Z3, Y1    
    VEXTRACTI128 $1, Y1, X2      
    VMOVQ X2, BX      
    ADDQ  BX, AX


done:
    MOVQ AX,   uint64_+24(FP)
    RET
