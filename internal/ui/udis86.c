// 2 january 2016
#include "libudis86/decode.c"
#include "libudis86/input.c"
#include "libudis86/itab.c"
#include "libudis86/syn-att.c"
#define opr_cast opr_cast_intel
#define gen_operand gen_operand_intel
#include "libudis86/syn-intel.c"
#include "libudis86/syn.c"
#include "libudis86/udis86.c"
