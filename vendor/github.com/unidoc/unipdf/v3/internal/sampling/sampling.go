//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

package sampling ;import (_a "github.com/unidoc/unipdf/v3/internal/bitwise";_b "github.com/unidoc/unipdf/v3/internal/imageutil";_e "io";);func (_fb *Reader )ReadSample ()(uint32 ,error ){if _fb ._af ==_fb ._bc .Height {return 0,_e .EOF ;};_ff ,_fe :=_fb ._bf .ReadBits (byte (_fb ._bc .BitsPerComponent ));
if _fe !=nil {return 0,_fe ;};_fb ._ag --;if _fb ._ag ==0{_fb ._ag =_fb ._bc .ColorComponents ;_fb ._aa ++;};if _fb ._aa ==_fb ._bc .Width {if _fb ._ga {_fb ._bf .ConsumeRemainingBits ();};_fb ._aa =0;_fb ._af ++;};return uint32 (_ff ),nil ;};func NewWriter (img _b .ImageBase )*Writer {return &Writer {_cge :_a .NewWriterMSB (img .Data ),_gd :img ,_ab :img .ColorComponents ,_gb :img .BytesPerLine *8!=img .ColorComponents *img .BitsPerComponent *img .Width };
};type Reader struct{_bc _b .ImageBase ;_bf *_a .Reader ;_aa ,_af ,_ag int ;_ga bool ;};func ResampleUint32 (data []uint32 ,bitsPerInputSample int ,bitsPerOutputSample int )[]uint32 {var _cbb []uint32 ;_bg :=bitsPerOutputSample ;var _gg uint32 ;var _egb uint32 ;
_eb :=0;_egbb :=0;_aef :=0;for _aef < len (data ){if _eb > 0{_fd :=_eb ;if _bg < _fd {_fd =_bg ;};_gg =(_gg <<uint (_fd ))|(_egb >>uint (bitsPerInputSample -_fd ));_eb -=_fd ;if _eb > 0{_egb =_egb <<uint (_fd );}else {_egb =0;};_bg -=_fd ;if _bg ==0{_cbb =append (_cbb ,_gg );
_bg =bitsPerOutputSample ;_gg =0;_egbb ++;};}else {_gf :=data [_aef ];_aef ++;_d :=bitsPerInputSample ;if _bg < _d {_d =_bg ;};_eb =bitsPerInputSample -_d ;_gg =(_gg <<uint (_d ))|(_gf >>uint (_eb ));if _d < bitsPerInputSample {_egb =_gf <<uint (_d );};
_bg -=_d ;if _bg ==0{_cbb =append (_cbb ,_gg );_bg =bitsPerOutputSample ;_gg =0;_egbb ++;};};};for _eb >=bitsPerOutputSample {_ce :=_eb ;if _bg < _ce {_ce =_bg ;};_gg =(_gg <<uint (_ce ))|(_egb >>uint (bitsPerInputSample -_ce ));_eb -=_ce ;if _eb > 0{_egb =_egb <<uint (_ce );
}else {_egb =0;};_bg -=_ce ;if _bg ==0{_cbb =append (_cbb ,_gg );_bg =bitsPerOutputSample ;_gg =0;_egbb ++;};};if _bg > 0&&_bg < bitsPerOutputSample {_gg <<=uint (_bg );_cbb =append (_cbb ,_gg );};return _cbb ;};func (_bd *Writer )WriteSamples (samples []uint32 )error {for _fdb :=0;
_fdb < len (samples );_fdb ++{if _aac :=_bd .WriteSample (samples [_fdb ]);_aac !=nil {return _aac ;};};return nil ;};type Writer struct{_gd _b .ImageBase ;_cge *_a .Writer ;_gc ,_ab int ;_gb bool ;};type SampleWriter interface{WriteSample (_ge uint32 )error ;
WriteSamples (_cg []uint32 )error ;};type SampleReader interface{ReadSample ()(uint32 ,error );ReadSamples (_f []uint32 )error ;};func ResampleBytes (data []byte ,bitsPerSample int )[]uint32 {var _bb []uint32 ;_fg :=bitsPerSample ;var _bbg uint32 ;var _afb byte ;
_eg :=0;_ac :=0;_cb :=0;for _cb < len (data ){if _eg > 0{_cd :=_eg ;if _fg < _cd {_cd =_fg ;};_bbg =(_bbg <<uint (_cd ))|uint32 (_afb >>uint (8-_cd ));_eg -=_cd ;if _eg > 0{_afb =_afb <<uint (_cd );}else {_afb =0;};_fg -=_cd ;if _fg ==0{_bb =append (_bb ,_bbg );
_fg =bitsPerSample ;_bbg =0;_ac ++;};}else {_ed :=data [_cb ];_cb ++;_cc :=8;if _fg < _cc {_cc =_fg ;};_eg =8-_cc ;_bbg =(_bbg <<uint (_cc ))|uint32 (_ed >>uint (_eg ));if _cc < 8{_afb =_ed <<uint (_cc );};_fg -=_cc ;if _fg ==0{_bb =append (_bb ,_bbg );
_fg =bitsPerSample ;_bbg =0;_ac ++;};};};for _eg >=bitsPerSample {_ae :=_eg ;if _fg < _ae {_ae =_fg ;};_bbg =(_bbg <<uint (_ae ))|uint32 (_afb >>uint (8-_ae ));_eg -=_ae ;if _eg > 0{_afb =_afb <<uint (_ae );}else {_afb =0;};_fg -=_ae ;if _fg ==0{_bb =append (_bb ,_bbg );
_fg =bitsPerSample ;_bbg =0;_ac ++;};};return _bb ;};func NewReader (img _b .ImageBase )*Reader {return &Reader {_bf :_a .NewReader (img .Data ),_bc :img ,_ag :img .ColorComponents ,_ga :img .BytesPerLine *8!=img .ColorComponents *img .BitsPerComponent *img .Width };
};func (_ffb *Reader )ReadSamples (samples []uint32 )(_c error ){for _ea :=0;_ea < len (samples );_ea ++{samples [_ea ],_c =_ffb .ReadSample ();if _c !=nil {return _c ;};};return nil ;};func (_ba *Writer )WriteSample (sample uint32 )error {if _ ,_dee :=_ba ._cge .WriteBits (uint64 (sample ),_ba ._gd .BitsPerComponent );
_dee !=nil {return _dee ;};_ba ._ab --;if _ba ._ab ==0{_ba ._ab =_ba ._gd .ColorComponents ;_ba ._gc ++;};if _ba ._gc ==_ba ._gd .Width {if _ba ._gb {_ba ._cge .FinishByte ();};_ba ._gc =0;};return nil ;};