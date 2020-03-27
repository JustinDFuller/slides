TEXT ·CompareAndSwapInt32(SB),NOSPLIT,$0
	JMP	runtime∕internal∕atomic·Cas(SB) // HL

TEXT runtime∕internal∕atomic·Cas(SB),NOSPLIT,$0-17 // HL
	MOVQ	ptr+0(FP), BX
	MOVL	old+8(FP), AX
	MOVL	new+12(FP), CX
	LOCK // HL
	CMPXCHGL	CX, 0(BX)
	SETEQ	ret+16(FP)
	RET
