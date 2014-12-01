package gen

import "reflect"

/*

#cgo linux LDFLAGS: -liup
#include <iup.h>
#include <stdlib.h>

extern int go_cb_handle_none_0(Ihandle* h);
extern int go_cb_handle_none_1(Ihandle* h);
extern int go_cb_handle_none_2(Ihandle* h);
extern int go_cb_handle_none_3(Ihandle* h);
extern int go_cb_handle_none_4(Ihandle* h);
extern int go_cb_handle_none_5(Ihandle* h);
extern int go_cb_handle_none_6(Ihandle* h);
extern int go_cb_handle_none_7(Ihandle* h);
extern int go_cb_handle_none_8(Ihandle* h);
extern int go_cb_handle_none_9(Ihandle* h);
extern int go_cb_handle_none_10(Ihandle* h);
extern int go_cb_handle_none_11(Ihandle* h);
extern int go_cb_handle_none_12(Ihandle* h);
extern int go_cb_handle_none_13(Ihandle* h);
extern int go_cb_handle_none_14(Ihandle* h);
extern int go_cb_handle_none_15(Ihandle* h);
extern int go_cb_handle_none_16(Ihandle* h);
extern int go_cb_handle_none_17(Ihandle* h);
extern int go_cb_handle_none_18(Ihandle* h);
extern int go_cb_handle_none_19(Ihandle* h);
extern int go_cb_handle_none_20(Ihandle* h);
extern int go_cb_handle_none_21(Ihandle* h);
extern int go_cb_handle_none_22(Ihandle* h);
extern int go_cb_handle_none_23(Ihandle* h);
extern int go_cb_handle_none_24(Ihandle* h);
extern int go_cb_handle_none_25(Ihandle* h);
extern int go_cb_handle_none_26(Ihandle* h);
extern int go_cb_handle_none_27(Ihandle* h);
extern int go_cb_handle_none_28(Ihandle* h);
extern int go_cb_handle_none_29(Ihandle* h);
extern int go_cb_handle_none_30(Ihandle* h);
extern int go_cb_handle_none_31(Ihandle* h);
extern int go_cb_handle_none_32(Ihandle* h);
extern int go_cb_handle_none_33(Ihandle* h);
extern int go_cb_handle_none_34(Ihandle* h);
extern int go_cb_handle_none_35(Ihandle* h);
extern int go_cb_handle_none_36(Ihandle* h);
extern int go_cb_handle_none_37(Ihandle* h);
extern int go_cb_handle_none_38(Ihandle* h);
extern int go_cb_handle_none_39(Ihandle* h);
extern int go_cb_handle_none_40(Ihandle* h);
extern int go_cb_handle_none_41(Ihandle* h);
extern int go_cb_handle_none_42(Ihandle* h);
extern int go_cb_handle_none_43(Ihandle* h);
extern int go_cb_handle_none_44(Ihandle* h);
extern int go_cb_handle_none_45(Ihandle* h);
extern int go_cb_handle_none_46(Ihandle* h);
extern int go_cb_handle_none_47(Ihandle* h);
extern int go_cb_handle_none_48(Ihandle* h);
extern int go_cb_handle_none_49(Ihandle* h);
extern int go_cb_handle_none_50(Ihandle* h);
extern int go_cb_handle_none_51(Ihandle* h);
extern int go_cb_handle_none_52(Ihandle* h);
extern int go_cb_handle_none_53(Ihandle* h);
extern int go_cb_handle_none_54(Ihandle* h);
extern int go_cb_handle_none_55(Ihandle* h);
extern int go_cb_handle_none_56(Ihandle* h);
extern int go_cb_handle_none_57(Ihandle* h);
extern int go_cb_handle_none_58(Ihandle* h);
extern int go_cb_handle_none_59(Ihandle* h);
extern int go_cb_handle_none_60(Ihandle* h);
extern int go_cb_handle_none_61(Ihandle* h);
extern int go_cb_handle_none_62(Ihandle* h);
extern int go_cb_handle_none_63(Ihandle* h);
extern int go_cb_handle_none_64(Ihandle* h);
extern int go_cb_handle_none_65(Ihandle* h);
extern int go_cb_handle_none_66(Ihandle* h);
extern int go_cb_handle_none_67(Ihandle* h);
extern int go_cb_handle_none_68(Ihandle* h);
extern int go_cb_handle_none_69(Ihandle* h);
extern int go_cb_handle_none_70(Ihandle* h);
extern int go_cb_handle_none_71(Ihandle* h);
extern int go_cb_handle_none_72(Ihandle* h);
extern int go_cb_handle_none_73(Ihandle* h);
extern int go_cb_handle_none_74(Ihandle* h);
extern int go_cb_handle_none_75(Ihandle* h);
extern int go_cb_handle_none_76(Ihandle* h);
extern int go_cb_handle_none_77(Ihandle* h);
extern int go_cb_handle_none_78(Ihandle* h);
extern int go_cb_handle_none_79(Ihandle* h);
extern int go_cb_handle_none_80(Ihandle* h);
extern int go_cb_handle_none_81(Ihandle* h);
extern int go_cb_handle_none_82(Ihandle* h);
extern int go_cb_handle_none_83(Ihandle* h);
extern int go_cb_handle_none_84(Ihandle* h);
extern int go_cb_handle_none_85(Ihandle* h);
extern int go_cb_handle_none_86(Ihandle* h);
extern int go_cb_handle_none_87(Ihandle* h);
extern int go_cb_handle_none_88(Ihandle* h);
extern int go_cb_handle_none_89(Ihandle* h);
extern int go_cb_handle_none_90(Ihandle* h);
extern int go_cb_handle_none_91(Ihandle* h);
extern int go_cb_handle_none_92(Ihandle* h);
extern int go_cb_handle_none_93(Ihandle* h);
extern int go_cb_handle_none_94(Ihandle* h);
extern int go_cb_handle_none_95(Ihandle* h);
extern int go_cb_handle_none_96(Ihandle* h);
extern int go_cb_handle_none_97(Ihandle* h);
extern int go_cb_handle_none_98(Ihandle* h);
extern int go_cb_handle_none_99(Ihandle* h);
extern int go_cb_handle_int_0(Ihandle* h, int arg0);
extern int go_cb_handle_int_1(Ihandle* h, int arg0);
extern int go_cb_handle_int_2(Ihandle* h, int arg0);
extern int go_cb_handle_int_3(Ihandle* h, int arg0);
extern int go_cb_handle_int_4(Ihandle* h, int arg0);
extern int go_cb_handle_int_5(Ihandle* h, int arg0);
extern int go_cb_handle_int_6(Ihandle* h, int arg0);
extern int go_cb_handle_int_7(Ihandle* h, int arg0);
extern int go_cb_handle_int_8(Ihandle* h, int arg0);
extern int go_cb_handle_int_9(Ihandle* h, int arg0);
extern int go_cb_handle_int_10(Ihandle* h, int arg0);
extern int go_cb_handle_int_11(Ihandle* h, int arg0);
extern int go_cb_handle_int_12(Ihandle* h, int arg0);
extern int go_cb_handle_int_13(Ihandle* h, int arg0);
extern int go_cb_handle_int_14(Ihandle* h, int arg0);
extern int go_cb_handle_int_15(Ihandle* h, int arg0);
extern int go_cb_handle_int_16(Ihandle* h, int arg0);
extern int go_cb_handle_int_17(Ihandle* h, int arg0);
extern int go_cb_handle_int_18(Ihandle* h, int arg0);
extern int go_cb_handle_int_19(Ihandle* h, int arg0);
extern int go_cb_handle_int_20(Ihandle* h, int arg0);
extern int go_cb_handle_int_21(Ihandle* h, int arg0);
extern int go_cb_handle_int_22(Ihandle* h, int arg0);
extern int go_cb_handle_int_23(Ihandle* h, int arg0);
extern int go_cb_handle_int_24(Ihandle* h, int arg0);
extern int go_cb_handle_int_25(Ihandle* h, int arg0);
extern int go_cb_handle_int_26(Ihandle* h, int arg0);
extern int go_cb_handle_int_27(Ihandle* h, int arg0);
extern int go_cb_handle_int_28(Ihandle* h, int arg0);
extern int go_cb_handle_int_29(Ihandle* h, int arg0);
extern int go_cb_handle_int_30(Ihandle* h, int arg0);
extern int go_cb_handle_int_31(Ihandle* h, int arg0);
extern int go_cb_handle_int_32(Ihandle* h, int arg0);
extern int go_cb_handle_int_33(Ihandle* h, int arg0);
extern int go_cb_handle_int_34(Ihandle* h, int arg0);
extern int go_cb_handle_int_35(Ihandle* h, int arg0);
extern int go_cb_handle_int_36(Ihandle* h, int arg0);
extern int go_cb_handle_int_37(Ihandle* h, int arg0);
extern int go_cb_handle_int_38(Ihandle* h, int arg0);
extern int go_cb_handle_int_39(Ihandle* h, int arg0);
extern int go_cb_handle_int_40(Ihandle* h, int arg0);
extern int go_cb_handle_int_41(Ihandle* h, int arg0);
extern int go_cb_handle_int_42(Ihandle* h, int arg0);
extern int go_cb_handle_int_43(Ihandle* h, int arg0);
extern int go_cb_handle_int_44(Ihandle* h, int arg0);
extern int go_cb_handle_int_45(Ihandle* h, int arg0);
extern int go_cb_handle_int_46(Ihandle* h, int arg0);
extern int go_cb_handle_int_47(Ihandle* h, int arg0);
extern int go_cb_handle_int_48(Ihandle* h, int arg0);
extern int go_cb_handle_int_49(Ihandle* h, int arg0);
extern int go_cb_handle_int_50(Ihandle* h, int arg0);
extern int go_cb_handle_int_51(Ihandle* h, int arg0);
extern int go_cb_handle_int_52(Ihandle* h, int arg0);
extern int go_cb_handle_int_53(Ihandle* h, int arg0);
extern int go_cb_handle_int_54(Ihandle* h, int arg0);
extern int go_cb_handle_int_55(Ihandle* h, int arg0);
extern int go_cb_handle_int_56(Ihandle* h, int arg0);
extern int go_cb_handle_int_57(Ihandle* h, int arg0);
extern int go_cb_handle_int_58(Ihandle* h, int arg0);
extern int go_cb_handle_int_59(Ihandle* h, int arg0);
extern int go_cb_handle_int_60(Ihandle* h, int arg0);
extern int go_cb_handle_int_61(Ihandle* h, int arg0);
extern int go_cb_handle_int_62(Ihandle* h, int arg0);
extern int go_cb_handle_int_63(Ihandle* h, int arg0);
extern int go_cb_handle_int_64(Ihandle* h, int arg0);
extern int go_cb_handle_int_65(Ihandle* h, int arg0);
extern int go_cb_handle_int_66(Ihandle* h, int arg0);
extern int go_cb_handle_int_67(Ihandle* h, int arg0);
extern int go_cb_handle_int_68(Ihandle* h, int arg0);
extern int go_cb_handle_int_69(Ihandle* h, int arg0);
extern int go_cb_handle_int_70(Ihandle* h, int arg0);
extern int go_cb_handle_int_71(Ihandle* h, int arg0);
extern int go_cb_handle_int_72(Ihandle* h, int arg0);
extern int go_cb_handle_int_73(Ihandle* h, int arg0);
extern int go_cb_handle_int_74(Ihandle* h, int arg0);
extern int go_cb_handle_int_75(Ihandle* h, int arg0);
extern int go_cb_handle_int_76(Ihandle* h, int arg0);
extern int go_cb_handle_int_77(Ihandle* h, int arg0);
extern int go_cb_handle_int_78(Ihandle* h, int arg0);
extern int go_cb_handle_int_79(Ihandle* h, int arg0);
extern int go_cb_handle_int_80(Ihandle* h, int arg0);
extern int go_cb_handle_int_81(Ihandle* h, int arg0);
extern int go_cb_handle_int_82(Ihandle* h, int arg0);
extern int go_cb_handle_int_83(Ihandle* h, int arg0);
extern int go_cb_handle_int_84(Ihandle* h, int arg0);
extern int go_cb_handle_int_85(Ihandle* h, int arg0);
extern int go_cb_handle_int_86(Ihandle* h, int arg0);
extern int go_cb_handle_int_87(Ihandle* h, int arg0);
extern int go_cb_handle_int_88(Ihandle* h, int arg0);
extern int go_cb_handle_int_89(Ihandle* h, int arg0);
extern int go_cb_handle_int_90(Ihandle* h, int arg0);
extern int go_cb_handle_int_91(Ihandle* h, int arg0);
extern int go_cb_handle_int_92(Ihandle* h, int arg0);
extern int go_cb_handle_int_93(Ihandle* h, int arg0);
extern int go_cb_handle_int_94(Ihandle* h, int arg0);
extern int go_cb_handle_int_95(Ihandle* h, int arg0);
extern int go_cb_handle_int_96(Ihandle* h, int arg0);
extern int go_cb_handle_int_97(Ihandle* h, int arg0);
extern int go_cb_handle_int_98(Ihandle* h, int arg0);
extern int go_cb_handle_int_99(Ihandle* h, int arg0);
extern int go_cb_handle_int_int_0(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_1(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_2(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_3(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_4(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_5(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_6(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_7(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_8(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_9(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_10(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_11(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_12(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_13(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_14(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_15(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_16(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_17(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_18(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_19(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_20(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_21(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_22(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_23(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_24(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_25(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_26(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_27(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_28(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_29(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_30(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_31(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_32(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_33(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_34(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_35(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_36(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_37(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_38(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_39(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_40(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_41(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_42(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_43(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_44(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_45(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_46(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_47(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_48(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_49(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_50(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_51(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_52(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_53(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_54(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_55(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_56(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_57(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_58(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_59(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_60(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_61(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_62(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_63(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_64(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_65(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_66(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_67(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_68(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_69(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_70(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_71(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_72(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_73(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_74(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_75(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_76(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_77(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_78(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_79(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_80(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_81(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_82(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_83(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_84(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_85(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_86(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_87(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_88(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_89(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_90(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_91(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_92(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_93(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_94(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_95(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_96(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_97(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_98(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_99(Ihandle* h, int arg0, int arg1);
extern int go_cb_handle_int_int_string_0(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_1(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_2(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_3(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_4(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_5(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_6(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_7(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_8(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_9(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_10(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_11(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_12(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_13(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_14(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_15(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_16(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_17(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_18(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_19(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_20(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_21(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_22(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_23(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_24(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_25(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_26(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_27(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_28(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_29(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_30(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_31(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_32(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_33(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_34(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_35(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_36(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_37(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_38(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_39(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_40(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_41(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_42(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_43(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_44(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_45(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_46(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_47(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_48(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_49(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_50(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_51(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_52(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_53(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_54(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_55(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_56(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_57(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_58(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_59(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_60(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_61(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_62(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_63(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_64(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_65(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_66(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_67(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_68(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_69(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_70(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_71(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_72(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_73(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_74(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_75(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_76(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_77(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_78(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_79(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_80(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_81(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_82(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_83(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_84(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_85(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_86(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_87(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_88(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_89(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_90(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_91(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_92(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_93(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_94(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_95(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_96(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_97(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_98(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_string_99(Ihandle* h, int arg0, int arg1, char* arg2);
extern int go_cb_handle_int_int_int_int_string_0(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_1(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_2(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_3(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_4(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_5(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_6(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_7(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_8(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_9(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_10(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_11(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_12(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_13(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_14(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_15(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_16(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_17(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_18(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_19(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_20(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_21(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_22(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_23(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_24(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_25(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_26(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_27(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_28(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_29(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_30(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_31(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_32(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_33(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_34(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_35(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_36(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_37(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_38(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_39(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_40(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_41(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_42(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_43(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_44(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_45(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_46(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_47(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_48(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_49(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_50(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_51(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_52(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_53(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_54(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_55(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_56(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_57(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_58(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_59(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_60(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_61(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_62(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_63(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_64(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_65(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_66(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_67(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_68(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_69(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_70(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_71(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_72(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_73(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_74(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_75(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_76(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_77(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_78(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_79(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_80(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_81(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_82(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_83(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_84(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_85(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_86(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_87(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_88(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_89(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_90(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_91(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_92(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_93(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_94(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_95(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_96(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_97(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_98(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_int_int_int_int_string_99(Ihandle* h, int arg0, int arg1, int arg2, int arg3, char* arg4);
extern int go_cb_handle_string_int_int_int_0(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_1(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_2(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_3(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_4(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_5(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_6(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_7(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_8(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_9(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_10(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_11(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_12(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_13(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_14(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_15(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_16(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_17(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_18(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_19(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_20(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_21(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_22(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_23(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_24(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_25(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_26(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_27(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_28(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_29(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_30(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_31(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_32(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_33(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_34(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_35(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_36(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_37(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_38(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_39(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_40(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_41(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_42(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_43(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_44(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_45(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_46(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_47(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_48(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_49(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_50(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_51(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_52(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_53(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_54(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_55(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_56(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_57(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_58(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_59(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_60(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_61(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_62(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_63(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_64(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_65(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_66(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_67(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_68(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_69(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_70(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_71(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_72(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_73(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_74(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_75(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_76(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_77(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_78(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_79(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_80(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_81(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_82(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_83(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_84(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_85(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_86(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_87(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_88(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_89(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_90(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_91(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_92(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_93(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_94(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_95(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_96(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_97(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_98(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);
extern int go_cb_handle_string_int_int_int_99(Ihandle* h, char* arg0, int arg1, int arg2, int arg3);

*/
import "C"

//export go_cb_handle_none_0
func go_cb_handle_none_0(h *C.Ihandle) C.int { return C.int(callbacks_none[0]()) }

//export go_cb_handle_none_1
func go_cb_handle_none_1(h *C.Ihandle) C.int { return C.int(callbacks_none[1]()) }

//export go_cb_handle_none_2
func go_cb_handle_none_2(h *C.Ihandle) C.int { return C.int(callbacks_none[2]()) }

//export go_cb_handle_none_3
func go_cb_handle_none_3(h *C.Ihandle) C.int { return C.int(callbacks_none[3]()) }

//export go_cb_handle_none_4
func go_cb_handle_none_4(h *C.Ihandle) C.int { return C.int(callbacks_none[4]()) }

//export go_cb_handle_none_5
func go_cb_handle_none_5(h *C.Ihandle) C.int { return C.int(callbacks_none[5]()) }

//export go_cb_handle_none_6
func go_cb_handle_none_6(h *C.Ihandle) C.int { return C.int(callbacks_none[6]()) }

//export go_cb_handle_none_7
func go_cb_handle_none_7(h *C.Ihandle) C.int { return C.int(callbacks_none[7]()) }

//export go_cb_handle_none_8
func go_cb_handle_none_8(h *C.Ihandle) C.int { return C.int(callbacks_none[8]()) }

//export go_cb_handle_none_9
func go_cb_handle_none_9(h *C.Ihandle) C.int { return C.int(callbacks_none[9]()) }

//export go_cb_handle_none_10
func go_cb_handle_none_10(h *C.Ihandle) C.int { return C.int(callbacks_none[10]()) }

//export go_cb_handle_none_11
func go_cb_handle_none_11(h *C.Ihandle) C.int { return C.int(callbacks_none[11]()) }

//export go_cb_handle_none_12
func go_cb_handle_none_12(h *C.Ihandle) C.int { return C.int(callbacks_none[12]()) }

//export go_cb_handle_none_13
func go_cb_handle_none_13(h *C.Ihandle) C.int { return C.int(callbacks_none[13]()) }

//export go_cb_handle_none_14
func go_cb_handle_none_14(h *C.Ihandle) C.int { return C.int(callbacks_none[14]()) }

//export go_cb_handle_none_15
func go_cb_handle_none_15(h *C.Ihandle) C.int { return C.int(callbacks_none[15]()) }

//export go_cb_handle_none_16
func go_cb_handle_none_16(h *C.Ihandle) C.int { return C.int(callbacks_none[16]()) }

//export go_cb_handle_none_17
func go_cb_handle_none_17(h *C.Ihandle) C.int { return C.int(callbacks_none[17]()) }

//export go_cb_handle_none_18
func go_cb_handle_none_18(h *C.Ihandle) C.int { return C.int(callbacks_none[18]()) }

//export go_cb_handle_none_19
func go_cb_handle_none_19(h *C.Ihandle) C.int { return C.int(callbacks_none[19]()) }

//export go_cb_handle_none_20
func go_cb_handle_none_20(h *C.Ihandle) C.int { return C.int(callbacks_none[20]()) }

//export go_cb_handle_none_21
func go_cb_handle_none_21(h *C.Ihandle) C.int { return C.int(callbacks_none[21]()) }

//export go_cb_handle_none_22
func go_cb_handle_none_22(h *C.Ihandle) C.int { return C.int(callbacks_none[22]()) }

//export go_cb_handle_none_23
func go_cb_handle_none_23(h *C.Ihandle) C.int { return C.int(callbacks_none[23]()) }

//export go_cb_handle_none_24
func go_cb_handle_none_24(h *C.Ihandle) C.int { return C.int(callbacks_none[24]()) }

//export go_cb_handle_none_25
func go_cb_handle_none_25(h *C.Ihandle) C.int { return C.int(callbacks_none[25]()) }

//export go_cb_handle_none_26
func go_cb_handle_none_26(h *C.Ihandle) C.int { return C.int(callbacks_none[26]()) }

//export go_cb_handle_none_27
func go_cb_handle_none_27(h *C.Ihandle) C.int { return C.int(callbacks_none[27]()) }

//export go_cb_handle_none_28
func go_cb_handle_none_28(h *C.Ihandle) C.int { return C.int(callbacks_none[28]()) }

//export go_cb_handle_none_29
func go_cb_handle_none_29(h *C.Ihandle) C.int { return C.int(callbacks_none[29]()) }

//export go_cb_handle_none_30
func go_cb_handle_none_30(h *C.Ihandle) C.int { return C.int(callbacks_none[30]()) }

//export go_cb_handle_none_31
func go_cb_handle_none_31(h *C.Ihandle) C.int { return C.int(callbacks_none[31]()) }

//export go_cb_handle_none_32
func go_cb_handle_none_32(h *C.Ihandle) C.int { return C.int(callbacks_none[32]()) }

//export go_cb_handle_none_33
func go_cb_handle_none_33(h *C.Ihandle) C.int { return C.int(callbacks_none[33]()) }

//export go_cb_handle_none_34
func go_cb_handle_none_34(h *C.Ihandle) C.int { return C.int(callbacks_none[34]()) }

//export go_cb_handle_none_35
func go_cb_handle_none_35(h *C.Ihandle) C.int { return C.int(callbacks_none[35]()) }

//export go_cb_handle_none_36
func go_cb_handle_none_36(h *C.Ihandle) C.int { return C.int(callbacks_none[36]()) }

//export go_cb_handle_none_37
func go_cb_handle_none_37(h *C.Ihandle) C.int { return C.int(callbacks_none[37]()) }

//export go_cb_handle_none_38
func go_cb_handle_none_38(h *C.Ihandle) C.int { return C.int(callbacks_none[38]()) }

//export go_cb_handle_none_39
func go_cb_handle_none_39(h *C.Ihandle) C.int { return C.int(callbacks_none[39]()) }

//export go_cb_handle_none_40
func go_cb_handle_none_40(h *C.Ihandle) C.int { return C.int(callbacks_none[40]()) }

//export go_cb_handle_none_41
func go_cb_handle_none_41(h *C.Ihandle) C.int { return C.int(callbacks_none[41]()) }

//export go_cb_handle_none_42
func go_cb_handle_none_42(h *C.Ihandle) C.int { return C.int(callbacks_none[42]()) }

//export go_cb_handle_none_43
func go_cb_handle_none_43(h *C.Ihandle) C.int { return C.int(callbacks_none[43]()) }

//export go_cb_handle_none_44
func go_cb_handle_none_44(h *C.Ihandle) C.int { return C.int(callbacks_none[44]()) }

//export go_cb_handle_none_45
func go_cb_handle_none_45(h *C.Ihandle) C.int { return C.int(callbacks_none[45]()) }

//export go_cb_handle_none_46
func go_cb_handle_none_46(h *C.Ihandle) C.int { return C.int(callbacks_none[46]()) }

//export go_cb_handle_none_47
func go_cb_handle_none_47(h *C.Ihandle) C.int { return C.int(callbacks_none[47]()) }

//export go_cb_handle_none_48
func go_cb_handle_none_48(h *C.Ihandle) C.int { return C.int(callbacks_none[48]()) }

//export go_cb_handle_none_49
func go_cb_handle_none_49(h *C.Ihandle) C.int { return C.int(callbacks_none[49]()) }

//export go_cb_handle_none_50
func go_cb_handle_none_50(h *C.Ihandle) C.int { return C.int(callbacks_none[50]()) }

//export go_cb_handle_none_51
func go_cb_handle_none_51(h *C.Ihandle) C.int { return C.int(callbacks_none[51]()) }

//export go_cb_handle_none_52
func go_cb_handle_none_52(h *C.Ihandle) C.int { return C.int(callbacks_none[52]()) }

//export go_cb_handle_none_53
func go_cb_handle_none_53(h *C.Ihandle) C.int { return C.int(callbacks_none[53]()) }

//export go_cb_handle_none_54
func go_cb_handle_none_54(h *C.Ihandle) C.int { return C.int(callbacks_none[54]()) }

//export go_cb_handle_none_55
func go_cb_handle_none_55(h *C.Ihandle) C.int { return C.int(callbacks_none[55]()) }

//export go_cb_handle_none_56
func go_cb_handle_none_56(h *C.Ihandle) C.int { return C.int(callbacks_none[56]()) }

//export go_cb_handle_none_57
func go_cb_handle_none_57(h *C.Ihandle) C.int { return C.int(callbacks_none[57]()) }

//export go_cb_handle_none_58
func go_cb_handle_none_58(h *C.Ihandle) C.int { return C.int(callbacks_none[58]()) }

//export go_cb_handle_none_59
func go_cb_handle_none_59(h *C.Ihandle) C.int { return C.int(callbacks_none[59]()) }

//export go_cb_handle_none_60
func go_cb_handle_none_60(h *C.Ihandle) C.int { return C.int(callbacks_none[60]()) }

//export go_cb_handle_none_61
func go_cb_handle_none_61(h *C.Ihandle) C.int { return C.int(callbacks_none[61]()) }

//export go_cb_handle_none_62
func go_cb_handle_none_62(h *C.Ihandle) C.int { return C.int(callbacks_none[62]()) }

//export go_cb_handle_none_63
func go_cb_handle_none_63(h *C.Ihandle) C.int { return C.int(callbacks_none[63]()) }

//export go_cb_handle_none_64
func go_cb_handle_none_64(h *C.Ihandle) C.int { return C.int(callbacks_none[64]()) }

//export go_cb_handle_none_65
func go_cb_handle_none_65(h *C.Ihandle) C.int { return C.int(callbacks_none[65]()) }

//export go_cb_handle_none_66
func go_cb_handle_none_66(h *C.Ihandle) C.int { return C.int(callbacks_none[66]()) }

//export go_cb_handle_none_67
func go_cb_handle_none_67(h *C.Ihandle) C.int { return C.int(callbacks_none[67]()) }

//export go_cb_handle_none_68
func go_cb_handle_none_68(h *C.Ihandle) C.int { return C.int(callbacks_none[68]()) }

//export go_cb_handle_none_69
func go_cb_handle_none_69(h *C.Ihandle) C.int { return C.int(callbacks_none[69]()) }

//export go_cb_handle_none_70
func go_cb_handle_none_70(h *C.Ihandle) C.int { return C.int(callbacks_none[70]()) }

//export go_cb_handle_none_71
func go_cb_handle_none_71(h *C.Ihandle) C.int { return C.int(callbacks_none[71]()) }

//export go_cb_handle_none_72
func go_cb_handle_none_72(h *C.Ihandle) C.int { return C.int(callbacks_none[72]()) }

//export go_cb_handle_none_73
func go_cb_handle_none_73(h *C.Ihandle) C.int { return C.int(callbacks_none[73]()) }

//export go_cb_handle_none_74
func go_cb_handle_none_74(h *C.Ihandle) C.int { return C.int(callbacks_none[74]()) }

//export go_cb_handle_none_75
func go_cb_handle_none_75(h *C.Ihandle) C.int { return C.int(callbacks_none[75]()) }

//export go_cb_handle_none_76
func go_cb_handle_none_76(h *C.Ihandle) C.int { return C.int(callbacks_none[76]()) }

//export go_cb_handle_none_77
func go_cb_handle_none_77(h *C.Ihandle) C.int { return C.int(callbacks_none[77]()) }

//export go_cb_handle_none_78
func go_cb_handle_none_78(h *C.Ihandle) C.int { return C.int(callbacks_none[78]()) }

//export go_cb_handle_none_79
func go_cb_handle_none_79(h *C.Ihandle) C.int { return C.int(callbacks_none[79]()) }

//export go_cb_handle_none_80
func go_cb_handle_none_80(h *C.Ihandle) C.int { return C.int(callbacks_none[80]()) }

//export go_cb_handle_none_81
func go_cb_handle_none_81(h *C.Ihandle) C.int { return C.int(callbacks_none[81]()) }

//export go_cb_handle_none_82
func go_cb_handle_none_82(h *C.Ihandle) C.int { return C.int(callbacks_none[82]()) }

//export go_cb_handle_none_83
func go_cb_handle_none_83(h *C.Ihandle) C.int { return C.int(callbacks_none[83]()) }

//export go_cb_handle_none_84
func go_cb_handle_none_84(h *C.Ihandle) C.int { return C.int(callbacks_none[84]()) }

//export go_cb_handle_none_85
func go_cb_handle_none_85(h *C.Ihandle) C.int { return C.int(callbacks_none[85]()) }

//export go_cb_handle_none_86
func go_cb_handle_none_86(h *C.Ihandle) C.int { return C.int(callbacks_none[86]()) }

//export go_cb_handle_none_87
func go_cb_handle_none_87(h *C.Ihandle) C.int { return C.int(callbacks_none[87]()) }

//export go_cb_handle_none_88
func go_cb_handle_none_88(h *C.Ihandle) C.int { return C.int(callbacks_none[88]()) }

//export go_cb_handle_none_89
func go_cb_handle_none_89(h *C.Ihandle) C.int { return C.int(callbacks_none[89]()) }

//export go_cb_handle_none_90
func go_cb_handle_none_90(h *C.Ihandle) C.int { return C.int(callbacks_none[90]()) }

//export go_cb_handle_none_91
func go_cb_handle_none_91(h *C.Ihandle) C.int { return C.int(callbacks_none[91]()) }

//export go_cb_handle_none_92
func go_cb_handle_none_92(h *C.Ihandle) C.int { return C.int(callbacks_none[92]()) }

//export go_cb_handle_none_93
func go_cb_handle_none_93(h *C.Ihandle) C.int { return C.int(callbacks_none[93]()) }

//export go_cb_handle_none_94
func go_cb_handle_none_94(h *C.Ihandle) C.int { return C.int(callbacks_none[94]()) }

//export go_cb_handle_none_95
func go_cb_handle_none_95(h *C.Ihandle) C.int { return C.int(callbacks_none[95]()) }

//export go_cb_handle_none_96
func go_cb_handle_none_96(h *C.Ihandle) C.int { return C.int(callbacks_none[96]()) }

//export go_cb_handle_none_97
func go_cb_handle_none_97(h *C.Ihandle) C.int { return C.int(callbacks_none[97]()) }

//export go_cb_handle_none_98
func go_cb_handle_none_98(h *C.Ihandle) C.int { return C.int(callbacks_none[98]()) }

//export go_cb_handle_none_99
func go_cb_handle_none_99(h *C.Ihandle) C.int { return C.int(callbacks_none[99]()) }

//export go_cb_handle_int_0
func go_cb_handle_int_0(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[0](int(arg0))) }

//export go_cb_handle_int_1
func go_cb_handle_int_1(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[1](int(arg0))) }

//export go_cb_handle_int_2
func go_cb_handle_int_2(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[2](int(arg0))) }

//export go_cb_handle_int_3
func go_cb_handle_int_3(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[3](int(arg0))) }

//export go_cb_handle_int_4
func go_cb_handle_int_4(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[4](int(arg0))) }

//export go_cb_handle_int_5
func go_cb_handle_int_5(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[5](int(arg0))) }

//export go_cb_handle_int_6
func go_cb_handle_int_6(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[6](int(arg0))) }

//export go_cb_handle_int_7
func go_cb_handle_int_7(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[7](int(arg0))) }

//export go_cb_handle_int_8
func go_cb_handle_int_8(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[8](int(arg0))) }

//export go_cb_handle_int_9
func go_cb_handle_int_9(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[9](int(arg0))) }

//export go_cb_handle_int_10
func go_cb_handle_int_10(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[10](int(arg0))) }

//export go_cb_handle_int_11
func go_cb_handle_int_11(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[11](int(arg0))) }

//export go_cb_handle_int_12
func go_cb_handle_int_12(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[12](int(arg0))) }

//export go_cb_handle_int_13
func go_cb_handle_int_13(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[13](int(arg0))) }

//export go_cb_handle_int_14
func go_cb_handle_int_14(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[14](int(arg0))) }

//export go_cb_handle_int_15
func go_cb_handle_int_15(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[15](int(arg0))) }

//export go_cb_handle_int_16
func go_cb_handle_int_16(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[16](int(arg0))) }

//export go_cb_handle_int_17
func go_cb_handle_int_17(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[17](int(arg0))) }

//export go_cb_handle_int_18
func go_cb_handle_int_18(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[18](int(arg0))) }

//export go_cb_handle_int_19
func go_cb_handle_int_19(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[19](int(arg0))) }

//export go_cb_handle_int_20
func go_cb_handle_int_20(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[20](int(arg0))) }

//export go_cb_handle_int_21
func go_cb_handle_int_21(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[21](int(arg0))) }

//export go_cb_handle_int_22
func go_cb_handle_int_22(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[22](int(arg0))) }

//export go_cb_handle_int_23
func go_cb_handle_int_23(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[23](int(arg0))) }

//export go_cb_handle_int_24
func go_cb_handle_int_24(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[24](int(arg0))) }

//export go_cb_handle_int_25
func go_cb_handle_int_25(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[25](int(arg0))) }

//export go_cb_handle_int_26
func go_cb_handle_int_26(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[26](int(arg0))) }

//export go_cb_handle_int_27
func go_cb_handle_int_27(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[27](int(arg0))) }

//export go_cb_handle_int_28
func go_cb_handle_int_28(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[28](int(arg0))) }

//export go_cb_handle_int_29
func go_cb_handle_int_29(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[29](int(arg0))) }

//export go_cb_handle_int_30
func go_cb_handle_int_30(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[30](int(arg0))) }

//export go_cb_handle_int_31
func go_cb_handle_int_31(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[31](int(arg0))) }

//export go_cb_handle_int_32
func go_cb_handle_int_32(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[32](int(arg0))) }

//export go_cb_handle_int_33
func go_cb_handle_int_33(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[33](int(arg0))) }

//export go_cb_handle_int_34
func go_cb_handle_int_34(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[34](int(arg0))) }

//export go_cb_handle_int_35
func go_cb_handle_int_35(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[35](int(arg0))) }

//export go_cb_handle_int_36
func go_cb_handle_int_36(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[36](int(arg0))) }

//export go_cb_handle_int_37
func go_cb_handle_int_37(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[37](int(arg0))) }

//export go_cb_handle_int_38
func go_cb_handle_int_38(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[38](int(arg0))) }

//export go_cb_handle_int_39
func go_cb_handle_int_39(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[39](int(arg0))) }

//export go_cb_handle_int_40
func go_cb_handle_int_40(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[40](int(arg0))) }

//export go_cb_handle_int_41
func go_cb_handle_int_41(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[41](int(arg0))) }

//export go_cb_handle_int_42
func go_cb_handle_int_42(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[42](int(arg0))) }

//export go_cb_handle_int_43
func go_cb_handle_int_43(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[43](int(arg0))) }

//export go_cb_handle_int_44
func go_cb_handle_int_44(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[44](int(arg0))) }

//export go_cb_handle_int_45
func go_cb_handle_int_45(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[45](int(arg0))) }

//export go_cb_handle_int_46
func go_cb_handle_int_46(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[46](int(arg0))) }

//export go_cb_handle_int_47
func go_cb_handle_int_47(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[47](int(arg0))) }

//export go_cb_handle_int_48
func go_cb_handle_int_48(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[48](int(arg0))) }

//export go_cb_handle_int_49
func go_cb_handle_int_49(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[49](int(arg0))) }

//export go_cb_handle_int_50
func go_cb_handle_int_50(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[50](int(arg0))) }

//export go_cb_handle_int_51
func go_cb_handle_int_51(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[51](int(arg0))) }

//export go_cb_handle_int_52
func go_cb_handle_int_52(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[52](int(arg0))) }

//export go_cb_handle_int_53
func go_cb_handle_int_53(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[53](int(arg0))) }

//export go_cb_handle_int_54
func go_cb_handle_int_54(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[54](int(arg0))) }

//export go_cb_handle_int_55
func go_cb_handle_int_55(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[55](int(arg0))) }

//export go_cb_handle_int_56
func go_cb_handle_int_56(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[56](int(arg0))) }

//export go_cb_handle_int_57
func go_cb_handle_int_57(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[57](int(arg0))) }

//export go_cb_handle_int_58
func go_cb_handle_int_58(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[58](int(arg0))) }

//export go_cb_handle_int_59
func go_cb_handle_int_59(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[59](int(arg0))) }

//export go_cb_handle_int_60
func go_cb_handle_int_60(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[60](int(arg0))) }

//export go_cb_handle_int_61
func go_cb_handle_int_61(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[61](int(arg0))) }

//export go_cb_handle_int_62
func go_cb_handle_int_62(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[62](int(arg0))) }

//export go_cb_handle_int_63
func go_cb_handle_int_63(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[63](int(arg0))) }

//export go_cb_handle_int_64
func go_cb_handle_int_64(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[64](int(arg0))) }

//export go_cb_handle_int_65
func go_cb_handle_int_65(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[65](int(arg0))) }

//export go_cb_handle_int_66
func go_cb_handle_int_66(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[66](int(arg0))) }

//export go_cb_handle_int_67
func go_cb_handle_int_67(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[67](int(arg0))) }

//export go_cb_handle_int_68
func go_cb_handle_int_68(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[68](int(arg0))) }

//export go_cb_handle_int_69
func go_cb_handle_int_69(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[69](int(arg0))) }

//export go_cb_handle_int_70
func go_cb_handle_int_70(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[70](int(arg0))) }

//export go_cb_handle_int_71
func go_cb_handle_int_71(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[71](int(arg0))) }

//export go_cb_handle_int_72
func go_cb_handle_int_72(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[72](int(arg0))) }

//export go_cb_handle_int_73
func go_cb_handle_int_73(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[73](int(arg0))) }

//export go_cb_handle_int_74
func go_cb_handle_int_74(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[74](int(arg0))) }

//export go_cb_handle_int_75
func go_cb_handle_int_75(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[75](int(arg0))) }

//export go_cb_handle_int_76
func go_cb_handle_int_76(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[76](int(arg0))) }

//export go_cb_handle_int_77
func go_cb_handle_int_77(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[77](int(arg0))) }

//export go_cb_handle_int_78
func go_cb_handle_int_78(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[78](int(arg0))) }

//export go_cb_handle_int_79
func go_cb_handle_int_79(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[79](int(arg0))) }

//export go_cb_handle_int_80
func go_cb_handle_int_80(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[80](int(arg0))) }

//export go_cb_handle_int_81
func go_cb_handle_int_81(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[81](int(arg0))) }

//export go_cb_handle_int_82
func go_cb_handle_int_82(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[82](int(arg0))) }

//export go_cb_handle_int_83
func go_cb_handle_int_83(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[83](int(arg0))) }

//export go_cb_handle_int_84
func go_cb_handle_int_84(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[84](int(arg0))) }

//export go_cb_handle_int_85
func go_cb_handle_int_85(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[85](int(arg0))) }

//export go_cb_handle_int_86
func go_cb_handle_int_86(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[86](int(arg0))) }

//export go_cb_handle_int_87
func go_cb_handle_int_87(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[87](int(arg0))) }

//export go_cb_handle_int_88
func go_cb_handle_int_88(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[88](int(arg0))) }

//export go_cb_handle_int_89
func go_cb_handle_int_89(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[89](int(arg0))) }

//export go_cb_handle_int_90
func go_cb_handle_int_90(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[90](int(arg0))) }

//export go_cb_handle_int_91
func go_cb_handle_int_91(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[91](int(arg0))) }

//export go_cb_handle_int_92
func go_cb_handle_int_92(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[92](int(arg0))) }

//export go_cb_handle_int_93
func go_cb_handle_int_93(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[93](int(arg0))) }

//export go_cb_handle_int_94
func go_cb_handle_int_94(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[94](int(arg0))) }

//export go_cb_handle_int_95
func go_cb_handle_int_95(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[95](int(arg0))) }

//export go_cb_handle_int_96
func go_cb_handle_int_96(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[96](int(arg0))) }

//export go_cb_handle_int_97
func go_cb_handle_int_97(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[97](int(arg0))) }

//export go_cb_handle_int_98
func go_cb_handle_int_98(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[98](int(arg0))) }

//export go_cb_handle_int_99
func go_cb_handle_int_99(h *C.Ihandle, arg0 C.int) C.int { return C.int(callbacks_int[99](int(arg0))) }

//export go_cb_handle_int_int_0
func go_cb_handle_int_int_0(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[0](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_1
func go_cb_handle_int_int_1(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[1](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_2
func go_cb_handle_int_int_2(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[2](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_3
func go_cb_handle_int_int_3(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[3](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_4
func go_cb_handle_int_int_4(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[4](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_5
func go_cb_handle_int_int_5(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[5](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_6
func go_cb_handle_int_int_6(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[6](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_7
func go_cb_handle_int_int_7(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[7](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_8
func go_cb_handle_int_int_8(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[8](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_9
func go_cb_handle_int_int_9(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[9](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_10
func go_cb_handle_int_int_10(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[10](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_11
func go_cb_handle_int_int_11(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[11](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_12
func go_cb_handle_int_int_12(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[12](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_13
func go_cb_handle_int_int_13(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[13](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_14
func go_cb_handle_int_int_14(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[14](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_15
func go_cb_handle_int_int_15(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[15](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_16
func go_cb_handle_int_int_16(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[16](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_17
func go_cb_handle_int_int_17(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[17](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_18
func go_cb_handle_int_int_18(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[18](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_19
func go_cb_handle_int_int_19(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[19](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_20
func go_cb_handle_int_int_20(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[20](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_21
func go_cb_handle_int_int_21(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[21](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_22
func go_cb_handle_int_int_22(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[22](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_23
func go_cb_handle_int_int_23(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[23](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_24
func go_cb_handle_int_int_24(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[24](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_25
func go_cb_handle_int_int_25(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[25](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_26
func go_cb_handle_int_int_26(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[26](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_27
func go_cb_handle_int_int_27(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[27](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_28
func go_cb_handle_int_int_28(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[28](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_29
func go_cb_handle_int_int_29(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[29](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_30
func go_cb_handle_int_int_30(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[30](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_31
func go_cb_handle_int_int_31(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[31](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_32
func go_cb_handle_int_int_32(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[32](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_33
func go_cb_handle_int_int_33(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[33](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_34
func go_cb_handle_int_int_34(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[34](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_35
func go_cb_handle_int_int_35(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[35](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_36
func go_cb_handle_int_int_36(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[36](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_37
func go_cb_handle_int_int_37(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[37](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_38
func go_cb_handle_int_int_38(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[38](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_39
func go_cb_handle_int_int_39(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[39](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_40
func go_cb_handle_int_int_40(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[40](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_41
func go_cb_handle_int_int_41(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[41](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_42
func go_cb_handle_int_int_42(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[42](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_43
func go_cb_handle_int_int_43(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[43](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_44
func go_cb_handle_int_int_44(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[44](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_45
func go_cb_handle_int_int_45(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[45](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_46
func go_cb_handle_int_int_46(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[46](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_47
func go_cb_handle_int_int_47(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[47](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_48
func go_cb_handle_int_int_48(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[48](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_49
func go_cb_handle_int_int_49(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[49](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_50
func go_cb_handle_int_int_50(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[50](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_51
func go_cb_handle_int_int_51(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[51](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_52
func go_cb_handle_int_int_52(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[52](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_53
func go_cb_handle_int_int_53(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[53](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_54
func go_cb_handle_int_int_54(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[54](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_55
func go_cb_handle_int_int_55(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[55](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_56
func go_cb_handle_int_int_56(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[56](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_57
func go_cb_handle_int_int_57(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[57](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_58
func go_cb_handle_int_int_58(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[58](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_59
func go_cb_handle_int_int_59(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[59](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_60
func go_cb_handle_int_int_60(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[60](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_61
func go_cb_handle_int_int_61(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[61](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_62
func go_cb_handle_int_int_62(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[62](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_63
func go_cb_handle_int_int_63(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[63](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_64
func go_cb_handle_int_int_64(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[64](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_65
func go_cb_handle_int_int_65(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[65](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_66
func go_cb_handle_int_int_66(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[66](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_67
func go_cb_handle_int_int_67(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[67](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_68
func go_cb_handle_int_int_68(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[68](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_69
func go_cb_handle_int_int_69(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[69](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_70
func go_cb_handle_int_int_70(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[70](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_71
func go_cb_handle_int_int_71(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[71](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_72
func go_cb_handle_int_int_72(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[72](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_73
func go_cb_handle_int_int_73(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[73](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_74
func go_cb_handle_int_int_74(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[74](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_75
func go_cb_handle_int_int_75(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[75](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_76
func go_cb_handle_int_int_76(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[76](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_77
func go_cb_handle_int_int_77(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[77](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_78
func go_cb_handle_int_int_78(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[78](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_79
func go_cb_handle_int_int_79(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[79](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_80
func go_cb_handle_int_int_80(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[80](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_81
func go_cb_handle_int_int_81(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[81](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_82
func go_cb_handle_int_int_82(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[82](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_83
func go_cb_handle_int_int_83(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[83](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_84
func go_cb_handle_int_int_84(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[84](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_85
func go_cb_handle_int_int_85(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[85](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_86
func go_cb_handle_int_int_86(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[86](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_87
func go_cb_handle_int_int_87(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[87](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_88
func go_cb_handle_int_int_88(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[88](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_89
func go_cb_handle_int_int_89(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[89](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_90
func go_cb_handle_int_int_90(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[90](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_91
func go_cb_handle_int_int_91(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[91](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_92
func go_cb_handle_int_int_92(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[92](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_93
func go_cb_handle_int_int_93(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[93](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_94
func go_cb_handle_int_int_94(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[94](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_95
func go_cb_handle_int_int_95(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[95](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_96
func go_cb_handle_int_int_96(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[96](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_97
func go_cb_handle_int_int_97(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[97](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_98
func go_cb_handle_int_int_98(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[98](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_99
func go_cb_handle_int_int_99(h *C.Ihandle, arg0 C.int, arg1 C.int) C.int { return C.int(callbacks_int_int[99](int(arg0), int(arg1))) }

//export go_cb_handle_int_int_string_0
func go_cb_handle_int_int_string_0(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[0](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_1
func go_cb_handle_int_int_string_1(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[1](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_2
func go_cb_handle_int_int_string_2(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[2](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_3
func go_cb_handle_int_int_string_3(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[3](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_4
func go_cb_handle_int_int_string_4(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[4](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_5
func go_cb_handle_int_int_string_5(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[5](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_6
func go_cb_handle_int_int_string_6(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[6](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_7
func go_cb_handle_int_int_string_7(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[7](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_8
func go_cb_handle_int_int_string_8(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[8](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_9
func go_cb_handle_int_int_string_9(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[9](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_10
func go_cb_handle_int_int_string_10(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[10](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_11
func go_cb_handle_int_int_string_11(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[11](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_12
func go_cb_handle_int_int_string_12(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[12](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_13
func go_cb_handle_int_int_string_13(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[13](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_14
func go_cb_handle_int_int_string_14(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[14](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_15
func go_cb_handle_int_int_string_15(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[15](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_16
func go_cb_handle_int_int_string_16(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[16](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_17
func go_cb_handle_int_int_string_17(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[17](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_18
func go_cb_handle_int_int_string_18(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[18](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_19
func go_cb_handle_int_int_string_19(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[19](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_20
func go_cb_handle_int_int_string_20(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[20](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_21
func go_cb_handle_int_int_string_21(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[21](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_22
func go_cb_handle_int_int_string_22(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[22](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_23
func go_cb_handle_int_int_string_23(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[23](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_24
func go_cb_handle_int_int_string_24(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[24](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_25
func go_cb_handle_int_int_string_25(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[25](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_26
func go_cb_handle_int_int_string_26(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[26](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_27
func go_cb_handle_int_int_string_27(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[27](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_28
func go_cb_handle_int_int_string_28(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[28](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_29
func go_cb_handle_int_int_string_29(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[29](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_30
func go_cb_handle_int_int_string_30(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[30](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_31
func go_cb_handle_int_int_string_31(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[31](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_32
func go_cb_handle_int_int_string_32(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[32](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_33
func go_cb_handle_int_int_string_33(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[33](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_34
func go_cb_handle_int_int_string_34(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[34](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_35
func go_cb_handle_int_int_string_35(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[35](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_36
func go_cb_handle_int_int_string_36(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[36](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_37
func go_cb_handle_int_int_string_37(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[37](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_38
func go_cb_handle_int_int_string_38(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[38](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_39
func go_cb_handle_int_int_string_39(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[39](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_40
func go_cb_handle_int_int_string_40(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[40](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_41
func go_cb_handle_int_int_string_41(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[41](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_42
func go_cb_handle_int_int_string_42(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[42](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_43
func go_cb_handle_int_int_string_43(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[43](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_44
func go_cb_handle_int_int_string_44(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[44](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_45
func go_cb_handle_int_int_string_45(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[45](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_46
func go_cb_handle_int_int_string_46(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[46](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_47
func go_cb_handle_int_int_string_47(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[47](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_48
func go_cb_handle_int_int_string_48(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[48](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_49
func go_cb_handle_int_int_string_49(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[49](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_50
func go_cb_handle_int_int_string_50(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[50](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_51
func go_cb_handle_int_int_string_51(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[51](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_52
func go_cb_handle_int_int_string_52(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[52](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_53
func go_cb_handle_int_int_string_53(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[53](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_54
func go_cb_handle_int_int_string_54(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[54](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_55
func go_cb_handle_int_int_string_55(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[55](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_56
func go_cb_handle_int_int_string_56(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[56](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_57
func go_cb_handle_int_int_string_57(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[57](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_58
func go_cb_handle_int_int_string_58(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[58](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_59
func go_cb_handle_int_int_string_59(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[59](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_60
func go_cb_handle_int_int_string_60(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[60](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_61
func go_cb_handle_int_int_string_61(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[61](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_62
func go_cb_handle_int_int_string_62(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[62](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_63
func go_cb_handle_int_int_string_63(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[63](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_64
func go_cb_handle_int_int_string_64(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[64](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_65
func go_cb_handle_int_int_string_65(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[65](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_66
func go_cb_handle_int_int_string_66(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[66](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_67
func go_cb_handle_int_int_string_67(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[67](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_68
func go_cb_handle_int_int_string_68(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[68](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_69
func go_cb_handle_int_int_string_69(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[69](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_70
func go_cb_handle_int_int_string_70(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[70](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_71
func go_cb_handle_int_int_string_71(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[71](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_72
func go_cb_handle_int_int_string_72(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[72](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_73
func go_cb_handle_int_int_string_73(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[73](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_74
func go_cb_handle_int_int_string_74(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[74](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_75
func go_cb_handle_int_int_string_75(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[75](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_76
func go_cb_handle_int_int_string_76(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[76](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_77
func go_cb_handle_int_int_string_77(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[77](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_78
func go_cb_handle_int_int_string_78(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[78](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_79
func go_cb_handle_int_int_string_79(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[79](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_80
func go_cb_handle_int_int_string_80(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[80](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_81
func go_cb_handle_int_int_string_81(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[81](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_82
func go_cb_handle_int_int_string_82(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[82](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_83
func go_cb_handle_int_int_string_83(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[83](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_84
func go_cb_handle_int_int_string_84(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[84](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_85
func go_cb_handle_int_int_string_85(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[85](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_86
func go_cb_handle_int_int_string_86(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[86](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_87
func go_cb_handle_int_int_string_87(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[87](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_88
func go_cb_handle_int_int_string_88(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[88](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_89
func go_cb_handle_int_int_string_89(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[89](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_90
func go_cb_handle_int_int_string_90(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[90](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_91
func go_cb_handle_int_int_string_91(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[91](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_92
func go_cb_handle_int_int_string_92(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[92](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_93
func go_cb_handle_int_int_string_93(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[93](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_94
func go_cb_handle_int_int_string_94(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[94](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_95
func go_cb_handle_int_int_string_95(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[95](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_96
func go_cb_handle_int_int_string_96(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[96](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_97
func go_cb_handle_int_int_string_97(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[97](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_98
func go_cb_handle_int_int_string_98(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[98](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_string_99
func go_cb_handle_int_int_string_99(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 *C.char) C.int { return C.int(callbacks_int_int_string[99](int(arg0), int(arg1), C.GoString(arg2))) }

//export go_cb_handle_int_int_int_int_string_0
func go_cb_handle_int_int_int_int_string_0(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[0](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_1
func go_cb_handle_int_int_int_int_string_1(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[1](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_2
func go_cb_handle_int_int_int_int_string_2(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[2](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_3
func go_cb_handle_int_int_int_int_string_3(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[3](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_4
func go_cb_handle_int_int_int_int_string_4(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[4](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_5
func go_cb_handle_int_int_int_int_string_5(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[5](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_6
func go_cb_handle_int_int_int_int_string_6(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[6](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_7
func go_cb_handle_int_int_int_int_string_7(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[7](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_8
func go_cb_handle_int_int_int_int_string_8(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[8](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_9
func go_cb_handle_int_int_int_int_string_9(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[9](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_10
func go_cb_handle_int_int_int_int_string_10(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[10](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_11
func go_cb_handle_int_int_int_int_string_11(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[11](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_12
func go_cb_handle_int_int_int_int_string_12(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[12](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_13
func go_cb_handle_int_int_int_int_string_13(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[13](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_14
func go_cb_handle_int_int_int_int_string_14(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[14](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_15
func go_cb_handle_int_int_int_int_string_15(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[15](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_16
func go_cb_handle_int_int_int_int_string_16(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[16](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_17
func go_cb_handle_int_int_int_int_string_17(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[17](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_18
func go_cb_handle_int_int_int_int_string_18(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[18](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_19
func go_cb_handle_int_int_int_int_string_19(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[19](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_20
func go_cb_handle_int_int_int_int_string_20(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[20](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_21
func go_cb_handle_int_int_int_int_string_21(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[21](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_22
func go_cb_handle_int_int_int_int_string_22(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[22](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_23
func go_cb_handle_int_int_int_int_string_23(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[23](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_24
func go_cb_handle_int_int_int_int_string_24(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[24](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_25
func go_cb_handle_int_int_int_int_string_25(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[25](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_26
func go_cb_handle_int_int_int_int_string_26(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[26](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_27
func go_cb_handle_int_int_int_int_string_27(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[27](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_28
func go_cb_handle_int_int_int_int_string_28(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[28](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_29
func go_cb_handle_int_int_int_int_string_29(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[29](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_30
func go_cb_handle_int_int_int_int_string_30(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[30](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_31
func go_cb_handle_int_int_int_int_string_31(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[31](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_32
func go_cb_handle_int_int_int_int_string_32(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[32](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_33
func go_cb_handle_int_int_int_int_string_33(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[33](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_34
func go_cb_handle_int_int_int_int_string_34(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[34](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_35
func go_cb_handle_int_int_int_int_string_35(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[35](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_36
func go_cb_handle_int_int_int_int_string_36(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[36](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_37
func go_cb_handle_int_int_int_int_string_37(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[37](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_38
func go_cb_handle_int_int_int_int_string_38(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[38](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_39
func go_cb_handle_int_int_int_int_string_39(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[39](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_40
func go_cb_handle_int_int_int_int_string_40(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[40](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_41
func go_cb_handle_int_int_int_int_string_41(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[41](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_42
func go_cb_handle_int_int_int_int_string_42(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[42](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_43
func go_cb_handle_int_int_int_int_string_43(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[43](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_44
func go_cb_handle_int_int_int_int_string_44(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[44](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_45
func go_cb_handle_int_int_int_int_string_45(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[45](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_46
func go_cb_handle_int_int_int_int_string_46(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[46](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_47
func go_cb_handle_int_int_int_int_string_47(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[47](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_48
func go_cb_handle_int_int_int_int_string_48(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[48](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_49
func go_cb_handle_int_int_int_int_string_49(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[49](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_50
func go_cb_handle_int_int_int_int_string_50(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[50](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_51
func go_cb_handle_int_int_int_int_string_51(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[51](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_52
func go_cb_handle_int_int_int_int_string_52(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[52](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_53
func go_cb_handle_int_int_int_int_string_53(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[53](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_54
func go_cb_handle_int_int_int_int_string_54(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[54](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_55
func go_cb_handle_int_int_int_int_string_55(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[55](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_56
func go_cb_handle_int_int_int_int_string_56(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[56](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_57
func go_cb_handle_int_int_int_int_string_57(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[57](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_58
func go_cb_handle_int_int_int_int_string_58(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[58](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_59
func go_cb_handle_int_int_int_int_string_59(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[59](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_60
func go_cb_handle_int_int_int_int_string_60(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[60](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_61
func go_cb_handle_int_int_int_int_string_61(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[61](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_62
func go_cb_handle_int_int_int_int_string_62(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[62](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_63
func go_cb_handle_int_int_int_int_string_63(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[63](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_64
func go_cb_handle_int_int_int_int_string_64(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[64](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_65
func go_cb_handle_int_int_int_int_string_65(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[65](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_66
func go_cb_handle_int_int_int_int_string_66(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[66](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_67
func go_cb_handle_int_int_int_int_string_67(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[67](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_68
func go_cb_handle_int_int_int_int_string_68(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[68](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_69
func go_cb_handle_int_int_int_int_string_69(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[69](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_70
func go_cb_handle_int_int_int_int_string_70(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[70](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_71
func go_cb_handle_int_int_int_int_string_71(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[71](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_72
func go_cb_handle_int_int_int_int_string_72(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[72](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_73
func go_cb_handle_int_int_int_int_string_73(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[73](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_74
func go_cb_handle_int_int_int_int_string_74(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[74](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_75
func go_cb_handle_int_int_int_int_string_75(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[75](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_76
func go_cb_handle_int_int_int_int_string_76(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[76](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_77
func go_cb_handle_int_int_int_int_string_77(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[77](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_78
func go_cb_handle_int_int_int_int_string_78(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[78](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_79
func go_cb_handle_int_int_int_int_string_79(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[79](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_80
func go_cb_handle_int_int_int_int_string_80(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[80](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_81
func go_cb_handle_int_int_int_int_string_81(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[81](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_82
func go_cb_handle_int_int_int_int_string_82(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[82](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_83
func go_cb_handle_int_int_int_int_string_83(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[83](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_84
func go_cb_handle_int_int_int_int_string_84(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[84](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_85
func go_cb_handle_int_int_int_int_string_85(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[85](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_86
func go_cb_handle_int_int_int_int_string_86(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[86](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_87
func go_cb_handle_int_int_int_int_string_87(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[87](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_88
func go_cb_handle_int_int_int_int_string_88(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[88](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_89
func go_cb_handle_int_int_int_int_string_89(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[89](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_90
func go_cb_handle_int_int_int_int_string_90(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[90](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_91
func go_cb_handle_int_int_int_int_string_91(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[91](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_92
func go_cb_handle_int_int_int_int_string_92(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[92](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_93
func go_cb_handle_int_int_int_int_string_93(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[93](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_94
func go_cb_handle_int_int_int_int_string_94(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[94](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_95
func go_cb_handle_int_int_int_int_string_95(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[95](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_96
func go_cb_handle_int_int_int_int_string_96(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[96](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_97
func go_cb_handle_int_int_int_int_string_97(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[97](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_98
func go_cb_handle_int_int_int_int_string_98(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[98](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_int_int_int_int_string_99
func go_cb_handle_int_int_int_int_string_99(h *C.Ihandle, arg0 C.int, arg1 C.int, arg2 C.int, arg3 C.int, arg4 *C.char) C.int { return C.int(callbacks_int_int_int_int_string[99](int(arg0), int(arg1), int(arg2), int(arg3), C.GoString(arg4))) }

//export go_cb_handle_string_int_int_int_0
func go_cb_handle_string_int_int_int_0(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[0](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_1
func go_cb_handle_string_int_int_int_1(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[1](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_2
func go_cb_handle_string_int_int_int_2(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[2](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_3
func go_cb_handle_string_int_int_int_3(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[3](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_4
func go_cb_handle_string_int_int_int_4(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[4](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_5
func go_cb_handle_string_int_int_int_5(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[5](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_6
func go_cb_handle_string_int_int_int_6(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[6](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_7
func go_cb_handle_string_int_int_int_7(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[7](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_8
func go_cb_handle_string_int_int_int_8(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[8](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_9
func go_cb_handle_string_int_int_int_9(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[9](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_10
func go_cb_handle_string_int_int_int_10(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[10](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_11
func go_cb_handle_string_int_int_int_11(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[11](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_12
func go_cb_handle_string_int_int_int_12(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[12](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_13
func go_cb_handle_string_int_int_int_13(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[13](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_14
func go_cb_handle_string_int_int_int_14(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[14](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_15
func go_cb_handle_string_int_int_int_15(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[15](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_16
func go_cb_handle_string_int_int_int_16(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[16](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_17
func go_cb_handle_string_int_int_int_17(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[17](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_18
func go_cb_handle_string_int_int_int_18(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[18](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_19
func go_cb_handle_string_int_int_int_19(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[19](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_20
func go_cb_handle_string_int_int_int_20(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[20](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_21
func go_cb_handle_string_int_int_int_21(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[21](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_22
func go_cb_handle_string_int_int_int_22(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[22](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_23
func go_cb_handle_string_int_int_int_23(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[23](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_24
func go_cb_handle_string_int_int_int_24(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[24](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_25
func go_cb_handle_string_int_int_int_25(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[25](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_26
func go_cb_handle_string_int_int_int_26(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[26](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_27
func go_cb_handle_string_int_int_int_27(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[27](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_28
func go_cb_handle_string_int_int_int_28(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[28](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_29
func go_cb_handle_string_int_int_int_29(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[29](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_30
func go_cb_handle_string_int_int_int_30(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[30](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_31
func go_cb_handle_string_int_int_int_31(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[31](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_32
func go_cb_handle_string_int_int_int_32(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[32](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_33
func go_cb_handle_string_int_int_int_33(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[33](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_34
func go_cb_handle_string_int_int_int_34(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[34](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_35
func go_cb_handle_string_int_int_int_35(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[35](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_36
func go_cb_handle_string_int_int_int_36(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[36](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_37
func go_cb_handle_string_int_int_int_37(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[37](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_38
func go_cb_handle_string_int_int_int_38(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[38](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_39
func go_cb_handle_string_int_int_int_39(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[39](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_40
func go_cb_handle_string_int_int_int_40(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[40](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_41
func go_cb_handle_string_int_int_int_41(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[41](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_42
func go_cb_handle_string_int_int_int_42(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[42](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_43
func go_cb_handle_string_int_int_int_43(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[43](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_44
func go_cb_handle_string_int_int_int_44(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[44](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_45
func go_cb_handle_string_int_int_int_45(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[45](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_46
func go_cb_handle_string_int_int_int_46(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[46](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_47
func go_cb_handle_string_int_int_int_47(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[47](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_48
func go_cb_handle_string_int_int_int_48(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[48](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_49
func go_cb_handle_string_int_int_int_49(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[49](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_50
func go_cb_handle_string_int_int_int_50(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[50](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_51
func go_cb_handle_string_int_int_int_51(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[51](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_52
func go_cb_handle_string_int_int_int_52(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[52](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_53
func go_cb_handle_string_int_int_int_53(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[53](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_54
func go_cb_handle_string_int_int_int_54(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[54](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_55
func go_cb_handle_string_int_int_int_55(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[55](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_56
func go_cb_handle_string_int_int_int_56(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[56](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_57
func go_cb_handle_string_int_int_int_57(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[57](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_58
func go_cb_handle_string_int_int_int_58(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[58](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_59
func go_cb_handle_string_int_int_int_59(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[59](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_60
func go_cb_handle_string_int_int_int_60(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[60](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_61
func go_cb_handle_string_int_int_int_61(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[61](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_62
func go_cb_handle_string_int_int_int_62(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[62](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_63
func go_cb_handle_string_int_int_int_63(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[63](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_64
func go_cb_handle_string_int_int_int_64(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[64](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_65
func go_cb_handle_string_int_int_int_65(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[65](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_66
func go_cb_handle_string_int_int_int_66(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[66](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_67
func go_cb_handle_string_int_int_int_67(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[67](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_68
func go_cb_handle_string_int_int_int_68(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[68](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_69
func go_cb_handle_string_int_int_int_69(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[69](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_70
func go_cb_handle_string_int_int_int_70(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[70](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_71
func go_cb_handle_string_int_int_int_71(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[71](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_72
func go_cb_handle_string_int_int_int_72(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[72](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_73
func go_cb_handle_string_int_int_int_73(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[73](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_74
func go_cb_handle_string_int_int_int_74(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[74](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_75
func go_cb_handle_string_int_int_int_75(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[75](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_76
func go_cb_handle_string_int_int_int_76(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[76](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_77
func go_cb_handle_string_int_int_int_77(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[77](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_78
func go_cb_handle_string_int_int_int_78(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[78](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_79
func go_cb_handle_string_int_int_int_79(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[79](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_80
func go_cb_handle_string_int_int_int_80(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[80](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_81
func go_cb_handle_string_int_int_int_81(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[81](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_82
func go_cb_handle_string_int_int_int_82(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[82](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_83
func go_cb_handle_string_int_int_int_83(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[83](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_84
func go_cb_handle_string_int_int_int_84(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[84](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_85
func go_cb_handle_string_int_int_int_85(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[85](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_86
func go_cb_handle_string_int_int_int_86(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[86](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_87
func go_cb_handle_string_int_int_int_87(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[87](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_88
func go_cb_handle_string_int_int_int_88(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[88](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_89
func go_cb_handle_string_int_int_int_89(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[89](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_90
func go_cb_handle_string_int_int_int_90(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[90](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_91
func go_cb_handle_string_int_int_int_91(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[91](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_92
func go_cb_handle_string_int_int_int_92(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[92](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_93
func go_cb_handle_string_int_int_int_93(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[93](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_94
func go_cb_handle_string_int_int_int_94(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[94](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_95
func go_cb_handle_string_int_int_int_95(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[95](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_96
func go_cb_handle_string_int_int_int_96(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[96](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_97
func go_cb_handle_string_int_int_int_97(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[97](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_98
func go_cb_handle_string_int_int_int_98(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[98](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

//export go_cb_handle_string_int_int_int_99
func go_cb_handle_string_int_int_int_99(h *C.Ihandle, arg0 *C.char, arg1 C.int, arg2 C.int, arg3 C.int) C.int { return C.int(callbacks_string_int_int_int[99](C.GoString(arg0), int(arg1), int(arg2), int(arg3))) }

var cCallbacks_none = []C.Icallback{
	(C.Icallback)(C.go_cb_handle_none_0),
	(C.Icallback)(C.go_cb_handle_none_1),
	(C.Icallback)(C.go_cb_handle_none_2),
	(C.Icallback)(C.go_cb_handle_none_3),
	(C.Icallback)(C.go_cb_handle_none_4),
	(C.Icallback)(C.go_cb_handle_none_5),
	(C.Icallback)(C.go_cb_handle_none_6),
	(C.Icallback)(C.go_cb_handle_none_7),
	(C.Icallback)(C.go_cb_handle_none_8),
	(C.Icallback)(C.go_cb_handle_none_9),
	(C.Icallback)(C.go_cb_handle_none_10),
	(C.Icallback)(C.go_cb_handle_none_11),
	(C.Icallback)(C.go_cb_handle_none_12),
	(C.Icallback)(C.go_cb_handle_none_13),
	(C.Icallback)(C.go_cb_handle_none_14),
	(C.Icallback)(C.go_cb_handle_none_15),
	(C.Icallback)(C.go_cb_handle_none_16),
	(C.Icallback)(C.go_cb_handle_none_17),
	(C.Icallback)(C.go_cb_handle_none_18),
	(C.Icallback)(C.go_cb_handle_none_19),
	(C.Icallback)(C.go_cb_handle_none_20),
	(C.Icallback)(C.go_cb_handle_none_21),
	(C.Icallback)(C.go_cb_handle_none_22),
	(C.Icallback)(C.go_cb_handle_none_23),
	(C.Icallback)(C.go_cb_handle_none_24),
	(C.Icallback)(C.go_cb_handle_none_25),
	(C.Icallback)(C.go_cb_handle_none_26),
	(C.Icallback)(C.go_cb_handle_none_27),
	(C.Icallback)(C.go_cb_handle_none_28),
	(C.Icallback)(C.go_cb_handle_none_29),
	(C.Icallback)(C.go_cb_handle_none_30),
	(C.Icallback)(C.go_cb_handle_none_31),
	(C.Icallback)(C.go_cb_handle_none_32),
	(C.Icallback)(C.go_cb_handle_none_33),
	(C.Icallback)(C.go_cb_handle_none_34),
	(C.Icallback)(C.go_cb_handle_none_35),
	(C.Icallback)(C.go_cb_handle_none_36),
	(C.Icallback)(C.go_cb_handle_none_37),
	(C.Icallback)(C.go_cb_handle_none_38),
	(C.Icallback)(C.go_cb_handle_none_39),
	(C.Icallback)(C.go_cb_handle_none_40),
	(C.Icallback)(C.go_cb_handle_none_41),
	(C.Icallback)(C.go_cb_handle_none_42),
	(C.Icallback)(C.go_cb_handle_none_43),
	(C.Icallback)(C.go_cb_handle_none_44),
	(C.Icallback)(C.go_cb_handle_none_45),
	(C.Icallback)(C.go_cb_handle_none_46),
	(C.Icallback)(C.go_cb_handle_none_47),
	(C.Icallback)(C.go_cb_handle_none_48),
	(C.Icallback)(C.go_cb_handle_none_49),
	(C.Icallback)(C.go_cb_handle_none_50),
	(C.Icallback)(C.go_cb_handle_none_51),
	(C.Icallback)(C.go_cb_handle_none_52),
	(C.Icallback)(C.go_cb_handle_none_53),
	(C.Icallback)(C.go_cb_handle_none_54),
	(C.Icallback)(C.go_cb_handle_none_55),
	(C.Icallback)(C.go_cb_handle_none_56),
	(C.Icallback)(C.go_cb_handle_none_57),
	(C.Icallback)(C.go_cb_handle_none_58),
	(C.Icallback)(C.go_cb_handle_none_59),
	(C.Icallback)(C.go_cb_handle_none_60),
	(C.Icallback)(C.go_cb_handle_none_61),
	(C.Icallback)(C.go_cb_handle_none_62),
	(C.Icallback)(C.go_cb_handle_none_63),
	(C.Icallback)(C.go_cb_handle_none_64),
	(C.Icallback)(C.go_cb_handle_none_65),
	(C.Icallback)(C.go_cb_handle_none_66),
	(C.Icallback)(C.go_cb_handle_none_67),
	(C.Icallback)(C.go_cb_handle_none_68),
	(C.Icallback)(C.go_cb_handle_none_69),
	(C.Icallback)(C.go_cb_handle_none_70),
	(C.Icallback)(C.go_cb_handle_none_71),
	(C.Icallback)(C.go_cb_handle_none_72),
	(C.Icallback)(C.go_cb_handle_none_73),
	(C.Icallback)(C.go_cb_handle_none_74),
	(C.Icallback)(C.go_cb_handle_none_75),
	(C.Icallback)(C.go_cb_handle_none_76),
	(C.Icallback)(C.go_cb_handle_none_77),
	(C.Icallback)(C.go_cb_handle_none_78),
	(C.Icallback)(C.go_cb_handle_none_79),
	(C.Icallback)(C.go_cb_handle_none_80),
	(C.Icallback)(C.go_cb_handle_none_81),
	(C.Icallback)(C.go_cb_handle_none_82),
	(C.Icallback)(C.go_cb_handle_none_83),
	(C.Icallback)(C.go_cb_handle_none_84),
	(C.Icallback)(C.go_cb_handle_none_85),
	(C.Icallback)(C.go_cb_handle_none_86),
	(C.Icallback)(C.go_cb_handle_none_87),
	(C.Icallback)(C.go_cb_handle_none_88),
	(C.Icallback)(C.go_cb_handle_none_89),
	(C.Icallback)(C.go_cb_handle_none_90),
	(C.Icallback)(C.go_cb_handle_none_91),
	(C.Icallback)(C.go_cb_handle_none_92),
	(C.Icallback)(C.go_cb_handle_none_93),
	(C.Icallback)(C.go_cb_handle_none_94),
	(C.Icallback)(C.go_cb_handle_none_95),
	(C.Icallback)(C.go_cb_handle_none_96),
	(C.Icallback)(C.go_cb_handle_none_97),
	(C.Icallback)(C.go_cb_handle_none_98),
	(C.Icallback)(C.go_cb_handle_none_99),
}

var cCallbacks_int = []C.Icallback{
	(C.Icallback)(C.go_cb_handle_int_0),
	(C.Icallback)(C.go_cb_handle_int_1),
	(C.Icallback)(C.go_cb_handle_int_2),
	(C.Icallback)(C.go_cb_handle_int_3),
	(C.Icallback)(C.go_cb_handle_int_4),
	(C.Icallback)(C.go_cb_handle_int_5),
	(C.Icallback)(C.go_cb_handle_int_6),
	(C.Icallback)(C.go_cb_handle_int_7),
	(C.Icallback)(C.go_cb_handle_int_8),
	(C.Icallback)(C.go_cb_handle_int_9),
	(C.Icallback)(C.go_cb_handle_int_10),
	(C.Icallback)(C.go_cb_handle_int_11),
	(C.Icallback)(C.go_cb_handle_int_12),
	(C.Icallback)(C.go_cb_handle_int_13),
	(C.Icallback)(C.go_cb_handle_int_14),
	(C.Icallback)(C.go_cb_handle_int_15),
	(C.Icallback)(C.go_cb_handle_int_16),
	(C.Icallback)(C.go_cb_handle_int_17),
	(C.Icallback)(C.go_cb_handle_int_18),
	(C.Icallback)(C.go_cb_handle_int_19),
	(C.Icallback)(C.go_cb_handle_int_20),
	(C.Icallback)(C.go_cb_handle_int_21),
	(C.Icallback)(C.go_cb_handle_int_22),
	(C.Icallback)(C.go_cb_handle_int_23),
	(C.Icallback)(C.go_cb_handle_int_24),
	(C.Icallback)(C.go_cb_handle_int_25),
	(C.Icallback)(C.go_cb_handle_int_26),
	(C.Icallback)(C.go_cb_handle_int_27),
	(C.Icallback)(C.go_cb_handle_int_28),
	(C.Icallback)(C.go_cb_handle_int_29),
	(C.Icallback)(C.go_cb_handle_int_30),
	(C.Icallback)(C.go_cb_handle_int_31),
	(C.Icallback)(C.go_cb_handle_int_32),
	(C.Icallback)(C.go_cb_handle_int_33),
	(C.Icallback)(C.go_cb_handle_int_34),
	(C.Icallback)(C.go_cb_handle_int_35),
	(C.Icallback)(C.go_cb_handle_int_36),
	(C.Icallback)(C.go_cb_handle_int_37),
	(C.Icallback)(C.go_cb_handle_int_38),
	(C.Icallback)(C.go_cb_handle_int_39),
	(C.Icallback)(C.go_cb_handle_int_40),
	(C.Icallback)(C.go_cb_handle_int_41),
	(C.Icallback)(C.go_cb_handle_int_42),
	(C.Icallback)(C.go_cb_handle_int_43),
	(C.Icallback)(C.go_cb_handle_int_44),
	(C.Icallback)(C.go_cb_handle_int_45),
	(C.Icallback)(C.go_cb_handle_int_46),
	(C.Icallback)(C.go_cb_handle_int_47),
	(C.Icallback)(C.go_cb_handle_int_48),
	(C.Icallback)(C.go_cb_handle_int_49),
	(C.Icallback)(C.go_cb_handle_int_50),
	(C.Icallback)(C.go_cb_handle_int_51),
	(C.Icallback)(C.go_cb_handle_int_52),
	(C.Icallback)(C.go_cb_handle_int_53),
	(C.Icallback)(C.go_cb_handle_int_54),
	(C.Icallback)(C.go_cb_handle_int_55),
	(C.Icallback)(C.go_cb_handle_int_56),
	(C.Icallback)(C.go_cb_handle_int_57),
	(C.Icallback)(C.go_cb_handle_int_58),
	(C.Icallback)(C.go_cb_handle_int_59),
	(C.Icallback)(C.go_cb_handle_int_60),
	(C.Icallback)(C.go_cb_handle_int_61),
	(C.Icallback)(C.go_cb_handle_int_62),
	(C.Icallback)(C.go_cb_handle_int_63),
	(C.Icallback)(C.go_cb_handle_int_64),
	(C.Icallback)(C.go_cb_handle_int_65),
	(C.Icallback)(C.go_cb_handle_int_66),
	(C.Icallback)(C.go_cb_handle_int_67),
	(C.Icallback)(C.go_cb_handle_int_68),
	(C.Icallback)(C.go_cb_handle_int_69),
	(C.Icallback)(C.go_cb_handle_int_70),
	(C.Icallback)(C.go_cb_handle_int_71),
	(C.Icallback)(C.go_cb_handle_int_72),
	(C.Icallback)(C.go_cb_handle_int_73),
	(C.Icallback)(C.go_cb_handle_int_74),
	(C.Icallback)(C.go_cb_handle_int_75),
	(C.Icallback)(C.go_cb_handle_int_76),
	(C.Icallback)(C.go_cb_handle_int_77),
	(C.Icallback)(C.go_cb_handle_int_78),
	(C.Icallback)(C.go_cb_handle_int_79),
	(C.Icallback)(C.go_cb_handle_int_80),
	(C.Icallback)(C.go_cb_handle_int_81),
	(C.Icallback)(C.go_cb_handle_int_82),
	(C.Icallback)(C.go_cb_handle_int_83),
	(C.Icallback)(C.go_cb_handle_int_84),
	(C.Icallback)(C.go_cb_handle_int_85),
	(C.Icallback)(C.go_cb_handle_int_86),
	(C.Icallback)(C.go_cb_handle_int_87),
	(C.Icallback)(C.go_cb_handle_int_88),
	(C.Icallback)(C.go_cb_handle_int_89),
	(C.Icallback)(C.go_cb_handle_int_90),
	(C.Icallback)(C.go_cb_handle_int_91),
	(C.Icallback)(C.go_cb_handle_int_92),
	(C.Icallback)(C.go_cb_handle_int_93),
	(C.Icallback)(C.go_cb_handle_int_94),
	(C.Icallback)(C.go_cb_handle_int_95),
	(C.Icallback)(C.go_cb_handle_int_96),
	(C.Icallback)(C.go_cb_handle_int_97),
	(C.Icallback)(C.go_cb_handle_int_98),
	(C.Icallback)(C.go_cb_handle_int_99),
}

var cCallbacks_int_int = []C.Icallback{
	(C.Icallback)(C.go_cb_handle_int_int_0),
	(C.Icallback)(C.go_cb_handle_int_int_1),
	(C.Icallback)(C.go_cb_handle_int_int_2),
	(C.Icallback)(C.go_cb_handle_int_int_3),
	(C.Icallback)(C.go_cb_handle_int_int_4),
	(C.Icallback)(C.go_cb_handle_int_int_5),
	(C.Icallback)(C.go_cb_handle_int_int_6),
	(C.Icallback)(C.go_cb_handle_int_int_7),
	(C.Icallback)(C.go_cb_handle_int_int_8),
	(C.Icallback)(C.go_cb_handle_int_int_9),
	(C.Icallback)(C.go_cb_handle_int_int_10),
	(C.Icallback)(C.go_cb_handle_int_int_11),
	(C.Icallback)(C.go_cb_handle_int_int_12),
	(C.Icallback)(C.go_cb_handle_int_int_13),
	(C.Icallback)(C.go_cb_handle_int_int_14),
	(C.Icallback)(C.go_cb_handle_int_int_15),
	(C.Icallback)(C.go_cb_handle_int_int_16),
	(C.Icallback)(C.go_cb_handle_int_int_17),
	(C.Icallback)(C.go_cb_handle_int_int_18),
	(C.Icallback)(C.go_cb_handle_int_int_19),
	(C.Icallback)(C.go_cb_handle_int_int_20),
	(C.Icallback)(C.go_cb_handle_int_int_21),
	(C.Icallback)(C.go_cb_handle_int_int_22),
	(C.Icallback)(C.go_cb_handle_int_int_23),
	(C.Icallback)(C.go_cb_handle_int_int_24),
	(C.Icallback)(C.go_cb_handle_int_int_25),
	(C.Icallback)(C.go_cb_handle_int_int_26),
	(C.Icallback)(C.go_cb_handle_int_int_27),
	(C.Icallback)(C.go_cb_handle_int_int_28),
	(C.Icallback)(C.go_cb_handle_int_int_29),
	(C.Icallback)(C.go_cb_handle_int_int_30),
	(C.Icallback)(C.go_cb_handle_int_int_31),
	(C.Icallback)(C.go_cb_handle_int_int_32),
	(C.Icallback)(C.go_cb_handle_int_int_33),
	(C.Icallback)(C.go_cb_handle_int_int_34),
	(C.Icallback)(C.go_cb_handle_int_int_35),
	(C.Icallback)(C.go_cb_handle_int_int_36),
	(C.Icallback)(C.go_cb_handle_int_int_37),
	(C.Icallback)(C.go_cb_handle_int_int_38),
	(C.Icallback)(C.go_cb_handle_int_int_39),
	(C.Icallback)(C.go_cb_handle_int_int_40),
	(C.Icallback)(C.go_cb_handle_int_int_41),
	(C.Icallback)(C.go_cb_handle_int_int_42),
	(C.Icallback)(C.go_cb_handle_int_int_43),
	(C.Icallback)(C.go_cb_handle_int_int_44),
	(C.Icallback)(C.go_cb_handle_int_int_45),
	(C.Icallback)(C.go_cb_handle_int_int_46),
	(C.Icallback)(C.go_cb_handle_int_int_47),
	(C.Icallback)(C.go_cb_handle_int_int_48),
	(C.Icallback)(C.go_cb_handle_int_int_49),
	(C.Icallback)(C.go_cb_handle_int_int_50),
	(C.Icallback)(C.go_cb_handle_int_int_51),
	(C.Icallback)(C.go_cb_handle_int_int_52),
	(C.Icallback)(C.go_cb_handle_int_int_53),
	(C.Icallback)(C.go_cb_handle_int_int_54),
	(C.Icallback)(C.go_cb_handle_int_int_55),
	(C.Icallback)(C.go_cb_handle_int_int_56),
	(C.Icallback)(C.go_cb_handle_int_int_57),
	(C.Icallback)(C.go_cb_handle_int_int_58),
	(C.Icallback)(C.go_cb_handle_int_int_59),
	(C.Icallback)(C.go_cb_handle_int_int_60),
	(C.Icallback)(C.go_cb_handle_int_int_61),
	(C.Icallback)(C.go_cb_handle_int_int_62),
	(C.Icallback)(C.go_cb_handle_int_int_63),
	(C.Icallback)(C.go_cb_handle_int_int_64),
	(C.Icallback)(C.go_cb_handle_int_int_65),
	(C.Icallback)(C.go_cb_handle_int_int_66),
	(C.Icallback)(C.go_cb_handle_int_int_67),
	(C.Icallback)(C.go_cb_handle_int_int_68),
	(C.Icallback)(C.go_cb_handle_int_int_69),
	(C.Icallback)(C.go_cb_handle_int_int_70),
	(C.Icallback)(C.go_cb_handle_int_int_71),
	(C.Icallback)(C.go_cb_handle_int_int_72),
	(C.Icallback)(C.go_cb_handle_int_int_73),
	(C.Icallback)(C.go_cb_handle_int_int_74),
	(C.Icallback)(C.go_cb_handle_int_int_75),
	(C.Icallback)(C.go_cb_handle_int_int_76),
	(C.Icallback)(C.go_cb_handle_int_int_77),
	(C.Icallback)(C.go_cb_handle_int_int_78),
	(C.Icallback)(C.go_cb_handle_int_int_79),
	(C.Icallback)(C.go_cb_handle_int_int_80),
	(C.Icallback)(C.go_cb_handle_int_int_81),
	(C.Icallback)(C.go_cb_handle_int_int_82),
	(C.Icallback)(C.go_cb_handle_int_int_83),
	(C.Icallback)(C.go_cb_handle_int_int_84),
	(C.Icallback)(C.go_cb_handle_int_int_85),
	(C.Icallback)(C.go_cb_handle_int_int_86),
	(C.Icallback)(C.go_cb_handle_int_int_87),
	(C.Icallback)(C.go_cb_handle_int_int_88),
	(C.Icallback)(C.go_cb_handle_int_int_89),
	(C.Icallback)(C.go_cb_handle_int_int_90),
	(C.Icallback)(C.go_cb_handle_int_int_91),
	(C.Icallback)(C.go_cb_handle_int_int_92),
	(C.Icallback)(C.go_cb_handle_int_int_93),
	(C.Icallback)(C.go_cb_handle_int_int_94),
	(C.Icallback)(C.go_cb_handle_int_int_95),
	(C.Icallback)(C.go_cb_handle_int_int_96),
	(C.Icallback)(C.go_cb_handle_int_int_97),
	(C.Icallback)(C.go_cb_handle_int_int_98),
	(C.Icallback)(C.go_cb_handle_int_int_99),
}

var cCallbacks_int_int_string = []C.Icallback{
	(C.Icallback)(C.go_cb_handle_int_int_string_0),
	(C.Icallback)(C.go_cb_handle_int_int_string_1),
	(C.Icallback)(C.go_cb_handle_int_int_string_2),
	(C.Icallback)(C.go_cb_handle_int_int_string_3),
	(C.Icallback)(C.go_cb_handle_int_int_string_4),
	(C.Icallback)(C.go_cb_handle_int_int_string_5),
	(C.Icallback)(C.go_cb_handle_int_int_string_6),
	(C.Icallback)(C.go_cb_handle_int_int_string_7),
	(C.Icallback)(C.go_cb_handle_int_int_string_8),
	(C.Icallback)(C.go_cb_handle_int_int_string_9),
	(C.Icallback)(C.go_cb_handle_int_int_string_10),
	(C.Icallback)(C.go_cb_handle_int_int_string_11),
	(C.Icallback)(C.go_cb_handle_int_int_string_12),
	(C.Icallback)(C.go_cb_handle_int_int_string_13),
	(C.Icallback)(C.go_cb_handle_int_int_string_14),
	(C.Icallback)(C.go_cb_handle_int_int_string_15),
	(C.Icallback)(C.go_cb_handle_int_int_string_16),
	(C.Icallback)(C.go_cb_handle_int_int_string_17),
	(C.Icallback)(C.go_cb_handle_int_int_string_18),
	(C.Icallback)(C.go_cb_handle_int_int_string_19),
	(C.Icallback)(C.go_cb_handle_int_int_string_20),
	(C.Icallback)(C.go_cb_handle_int_int_string_21),
	(C.Icallback)(C.go_cb_handle_int_int_string_22),
	(C.Icallback)(C.go_cb_handle_int_int_string_23),
	(C.Icallback)(C.go_cb_handle_int_int_string_24),
	(C.Icallback)(C.go_cb_handle_int_int_string_25),
	(C.Icallback)(C.go_cb_handle_int_int_string_26),
	(C.Icallback)(C.go_cb_handle_int_int_string_27),
	(C.Icallback)(C.go_cb_handle_int_int_string_28),
	(C.Icallback)(C.go_cb_handle_int_int_string_29),
	(C.Icallback)(C.go_cb_handle_int_int_string_30),
	(C.Icallback)(C.go_cb_handle_int_int_string_31),
	(C.Icallback)(C.go_cb_handle_int_int_string_32),
	(C.Icallback)(C.go_cb_handle_int_int_string_33),
	(C.Icallback)(C.go_cb_handle_int_int_string_34),
	(C.Icallback)(C.go_cb_handle_int_int_string_35),
	(C.Icallback)(C.go_cb_handle_int_int_string_36),
	(C.Icallback)(C.go_cb_handle_int_int_string_37),
	(C.Icallback)(C.go_cb_handle_int_int_string_38),
	(C.Icallback)(C.go_cb_handle_int_int_string_39),
	(C.Icallback)(C.go_cb_handle_int_int_string_40),
	(C.Icallback)(C.go_cb_handle_int_int_string_41),
	(C.Icallback)(C.go_cb_handle_int_int_string_42),
	(C.Icallback)(C.go_cb_handle_int_int_string_43),
	(C.Icallback)(C.go_cb_handle_int_int_string_44),
	(C.Icallback)(C.go_cb_handle_int_int_string_45),
	(C.Icallback)(C.go_cb_handle_int_int_string_46),
	(C.Icallback)(C.go_cb_handle_int_int_string_47),
	(C.Icallback)(C.go_cb_handle_int_int_string_48),
	(C.Icallback)(C.go_cb_handle_int_int_string_49),
	(C.Icallback)(C.go_cb_handle_int_int_string_50),
	(C.Icallback)(C.go_cb_handle_int_int_string_51),
	(C.Icallback)(C.go_cb_handle_int_int_string_52),
	(C.Icallback)(C.go_cb_handle_int_int_string_53),
	(C.Icallback)(C.go_cb_handle_int_int_string_54),
	(C.Icallback)(C.go_cb_handle_int_int_string_55),
	(C.Icallback)(C.go_cb_handle_int_int_string_56),
	(C.Icallback)(C.go_cb_handle_int_int_string_57),
	(C.Icallback)(C.go_cb_handle_int_int_string_58),
	(C.Icallback)(C.go_cb_handle_int_int_string_59),
	(C.Icallback)(C.go_cb_handle_int_int_string_60),
	(C.Icallback)(C.go_cb_handle_int_int_string_61),
	(C.Icallback)(C.go_cb_handle_int_int_string_62),
	(C.Icallback)(C.go_cb_handle_int_int_string_63),
	(C.Icallback)(C.go_cb_handle_int_int_string_64),
	(C.Icallback)(C.go_cb_handle_int_int_string_65),
	(C.Icallback)(C.go_cb_handle_int_int_string_66),
	(C.Icallback)(C.go_cb_handle_int_int_string_67),
	(C.Icallback)(C.go_cb_handle_int_int_string_68),
	(C.Icallback)(C.go_cb_handle_int_int_string_69),
	(C.Icallback)(C.go_cb_handle_int_int_string_70),
	(C.Icallback)(C.go_cb_handle_int_int_string_71),
	(C.Icallback)(C.go_cb_handle_int_int_string_72),
	(C.Icallback)(C.go_cb_handle_int_int_string_73),
	(C.Icallback)(C.go_cb_handle_int_int_string_74),
	(C.Icallback)(C.go_cb_handle_int_int_string_75),
	(C.Icallback)(C.go_cb_handle_int_int_string_76),
	(C.Icallback)(C.go_cb_handle_int_int_string_77),
	(C.Icallback)(C.go_cb_handle_int_int_string_78),
	(C.Icallback)(C.go_cb_handle_int_int_string_79),
	(C.Icallback)(C.go_cb_handle_int_int_string_80),
	(C.Icallback)(C.go_cb_handle_int_int_string_81),
	(C.Icallback)(C.go_cb_handle_int_int_string_82),
	(C.Icallback)(C.go_cb_handle_int_int_string_83),
	(C.Icallback)(C.go_cb_handle_int_int_string_84),
	(C.Icallback)(C.go_cb_handle_int_int_string_85),
	(C.Icallback)(C.go_cb_handle_int_int_string_86),
	(C.Icallback)(C.go_cb_handle_int_int_string_87),
	(C.Icallback)(C.go_cb_handle_int_int_string_88),
	(C.Icallback)(C.go_cb_handle_int_int_string_89),
	(C.Icallback)(C.go_cb_handle_int_int_string_90),
	(C.Icallback)(C.go_cb_handle_int_int_string_91),
	(C.Icallback)(C.go_cb_handle_int_int_string_92),
	(C.Icallback)(C.go_cb_handle_int_int_string_93),
	(C.Icallback)(C.go_cb_handle_int_int_string_94),
	(C.Icallback)(C.go_cb_handle_int_int_string_95),
	(C.Icallback)(C.go_cb_handle_int_int_string_96),
	(C.Icallback)(C.go_cb_handle_int_int_string_97),
	(C.Icallback)(C.go_cb_handle_int_int_string_98),
	(C.Icallback)(C.go_cb_handle_int_int_string_99),
}

var cCallbacks_int_int_int_int_string = []C.Icallback{
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_0),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_1),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_2),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_3),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_4),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_5),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_6),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_7),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_8),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_9),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_10),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_11),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_12),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_13),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_14),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_15),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_16),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_17),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_18),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_19),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_20),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_21),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_22),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_23),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_24),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_25),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_26),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_27),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_28),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_29),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_30),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_31),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_32),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_33),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_34),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_35),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_36),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_37),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_38),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_39),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_40),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_41),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_42),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_43),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_44),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_45),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_46),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_47),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_48),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_49),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_50),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_51),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_52),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_53),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_54),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_55),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_56),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_57),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_58),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_59),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_60),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_61),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_62),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_63),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_64),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_65),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_66),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_67),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_68),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_69),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_70),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_71),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_72),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_73),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_74),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_75),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_76),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_77),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_78),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_79),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_80),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_81),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_82),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_83),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_84),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_85),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_86),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_87),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_88),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_89),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_90),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_91),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_92),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_93),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_94),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_95),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_96),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_97),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_98),
	(C.Icallback)(C.go_cb_handle_int_int_int_int_string_99),
}

var cCallbacks_string_int_int_int = []C.Icallback{
	(C.Icallback)(C.go_cb_handle_string_int_int_int_0),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_1),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_2),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_3),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_4),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_5),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_6),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_7),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_8),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_9),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_10),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_11),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_12),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_13),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_14),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_15),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_16),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_17),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_18),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_19),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_20),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_21),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_22),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_23),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_24),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_25),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_26),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_27),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_28),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_29),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_30),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_31),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_32),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_33),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_34),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_35),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_36),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_37),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_38),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_39),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_40),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_41),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_42),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_43),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_44),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_45),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_46),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_47),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_48),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_49),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_50),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_51),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_52),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_53),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_54),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_55),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_56),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_57),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_58),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_59),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_60),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_61),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_62),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_63),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_64),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_65),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_66),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_67),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_68),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_69),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_70),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_71),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_72),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_73),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_74),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_75),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_76),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_77),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_78),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_79),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_80),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_81),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_82),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_83),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_84),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_85),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_86),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_87),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_88),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_89),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_90),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_91),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_92),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_93),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_94),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_95),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_96),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_97),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_98),
	(C.Icallback)(C.go_cb_handle_string_int_int_int_99),
}

var callbacks_none []func() int
var callbacks_int []func(arg0 int) int
var callbacks_int_int []func(arg0 int, arg1 int) int
var callbacks_int_int_string []func(arg0 int, arg1 int, arg2 string) int
var callbacks_int_int_int_int_string []func(arg0 int, arg1 int, arg2 int, arg3 int, arg4 string) int
var callbacks_string_int_int_int []func(arg0 string, arg1 int, arg2 int, arg3 int) int

func GetNextCallback(cb interface{}) C.Icallback {
	typ := reflect.TypeOf(cb)
	checkIntReturningFunction(typ)
	if typ.NumIn() == 0 {
		checkEnoughCallbacks(len(callbacks_none), len(cCallbacks_none))
		callbacks_none = append(callbacks_none, cb.(func() int))
		return cCallbacks_none[len(callbacks_none)-1]
	}
	if typ.NumIn() == 1 && typ.In(0).String() == "int" {
		checkEnoughCallbacks(len(callbacks_int), len(cCallbacks_int))
		callbacks_int = append(callbacks_int, cb.(func(int) int))
		return cCallbacks_int[len(callbacks_int)-1]
	}
	if typ.NumIn() == 2 && typ.In(0).String() == "int" && typ.In(1).String() == "int" {
		checkEnoughCallbacks(len(callbacks_int_int), len(cCallbacks_int_int))
		callbacks_int_int = append(callbacks_int_int, cb.(func(int, int) int))
		return cCallbacks_int_int[len(callbacks_int_int)-1]
	}
	if typ.NumIn() == 3 && typ.In(0).String() == "int" && typ.In(1).String() == "int" && typ.In(2).String() == "string" {
		checkEnoughCallbacks(len(callbacks_int_int_string), len(cCallbacks_int_int_string))
		callbacks_int_int_string = append(callbacks_int_int_string, cb.(func(int, int, string) int))
		return cCallbacks_int_int_string[len(callbacks_int_int_string)-1]
	}
	if typ.NumIn() == 5 && typ.In(0).String() == "int" && typ.In(1).String() == "int" && typ.In(2).String() == "int" && typ.In(3).String() == "int" && typ.In(4).String() == "string" {
		checkEnoughCallbacks(len(callbacks_int_int_int_int_string), len(cCallbacks_int_int_int_int_string))
		callbacks_int_int_int_int_string = append(callbacks_int_int_int_int_string, cb.(func(int, int, int, int, string) int))
		return cCallbacks_int_int_int_int_string[len(callbacks_int_int_int_int_string)-1]
	}
	if typ.NumIn() == 4 && typ.In(0).String() == "string" && typ.In(1).String() == "int" && typ.In(2).String() == "int" && typ.In(3).String() == "int" {
		checkEnoughCallbacks(len(callbacks_string_int_int_int), len(cCallbacks_string_int_int_int))
		callbacks_string_int_int_int = append(callbacks_string_int_int_int, cb.(func(string, int, int, int) int))
		return cCallbacks_string_int_int_int[len(callbacks_string_int_int_int)-1]
	}
	panic("unsupported callback function type")
}

func checkIntReturningFunction(typ reflect.Type) {
	if typ.Kind() != reflect.Func {
		panic("callback must be a function")
	}
	if typ.NumOut() != 1 || typ.Out(0).String() != "int" {
		panic("callback must return int")
	}
}

func checkEnoughCallbacks(inGo, inC int) {
	if inGo > inC-1 {
		panic("not enough callbacks, regenerate more")
	}
}
