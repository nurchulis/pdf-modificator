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

package security ;import (_bc "bytes";_gf "crypto/aes";_g "crypto/cipher";_c "crypto/md5";_e "crypto/rand";_fb "crypto/rc4";_af "crypto/sha256";_f "crypto/sha512";_fe "encoding/binary";_gb "errors";_db "fmt";_bd "github.com/unidoc/unipdf/v3/common";_a "hash";
_b "io";_de "math";);var _ StdHandler =stdHandlerR4 {};const (EventDocOpen =AuthEvent ("\u0044o\u0063\u004f\u0070\u0065\u006e");EventEFOpen =AuthEvent ("\u0045\u0046\u004f\u0070\u0065\u006e"););var _ StdHandler =stdHandlerR6 {};func (_cfd *ecbDecrypter )CryptBlocks (dst ,src []byte ){if len (src )%_cfd ._gd !=0{_bd .Log .Error ("\u0045\u0052\u0052\u004f\u0052:\u0020\u0045\u0043\u0042\u0020\u0064\u0065\u0063\u0072\u0079\u0070\u0074\u003a \u0069\u006e\u0070\u0075\u0074\u0020\u006e\u006f\u0074\u0020\u0066\u0075\u006c\u006c\u0020\u0062\u006c\u006f\u0063\u006b\u0073");
return ;};if len (dst )< len (src ){_bd .Log .Error ("\u0045R\u0052\u004fR\u003a\u0020\u0045C\u0042\u0020\u0064\u0065\u0063\u0072\u0079p\u0074\u003a\u0020\u006f\u0075\u0074p\u0075\u0074\u0020\u0073\u006d\u0061\u006c\u006c\u0065\u0072\u0020t\u0068\u0061\u006e\u0020\u0069\u006e\u0070\u0075\u0074");
return ;};for len (src )> 0{_cfd ._cb .Decrypt (dst ,src [:_cfd ._gd ]);src =src [_cfd ._gd :];dst =dst [_cfd ._gd :];};};func (_gcf *ecbDecrypter )BlockSize ()int {return _gcf ._gd };func _bg (_fc _g .Block )_g .BlockMode {return (*ecbDecrypter )(_dba (_fc ))};
func (_fbf stdHandlerR4 )alg2 (_ecf *StdEncryptDict ,_ab []byte )[]byte {_bd .Log .Trace ("\u0061\u006c\u0067\u0032");_aeg :=_fbf .paddedPass (_ab );_baf :=_c .New ();_baf .Write (_aeg );_baf .Write (_ecf .O );var _dg [4]byte ;_fe .LittleEndian .PutUint32 (_dg [:],uint32 (_ecf .P ));
_baf .Write (_dg [:]);_bd .Log .Trace ("\u0067o\u0020\u0050\u003a\u0020\u0025\u0020x",_dg );_baf .Write ([]byte (_fbf .ID0 ));_bd .Log .Trace ("\u0074\u0068\u0069\u0073\u002e\u0052\u0020\u003d\u0020\u0025d\u0020\u0065\u006e\u0063\u0072\u0079\u0070t\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061\u0020\u0025\u0076",_ecf .R ,_ecf .EncryptMetadata );
if (_ecf .R >=4)&&!_ecf .EncryptMetadata {_baf .Write ([]byte {0xff,0xff,0xff,0xff});};_bad :=_baf .Sum (nil );if _ecf .R >=3{_baf =_c .New ();for _ce :=0;_ce < 50;_ce ++{_baf .Reset ();_baf .Write (_bad [0:_fbf .Length /8]);_bad =_baf .Sum (nil );};};
if _ecf .R >=3{return _bad [0:_fbf .Length /8];};return _bad [0:5];};func _gcfe (_cadd []byte )([]byte ,error ){_gcfg :=_af .New ();_gcfg .Write (_cadd );return _gcfg .Sum (nil ),nil ;};

// Authenticate implements StdHandler interface.
func (_feb stdHandlerR4 )Authenticate (d *StdEncryptDict ,pass []byte )([]byte ,Permissions ,error ){_bd .Log .Trace ("\u0044\u0065b\u0075\u0067\u0067\u0069n\u0067\u0020a\u0075\u0074\u0068\u0065\u006e\u0074\u0069\u0063a\u0074\u0069\u006f\u006e\u0020\u002d\u0020\u006f\u0077\u006e\u0065\u0072 \u0070\u0061\u0073\u0073");
_bf ,_bbb :=_feb .alg7 (d ,pass );if _bbb !=nil {return nil ,0,_bbb ;};if _bf !=nil {_bd .Log .Trace ("\u0074h\u0069\u0073\u002e\u0061u\u0074\u0068\u0065\u006e\u0074i\u0063a\u0074e\u0064\u0020\u003d\u0020\u0054\u0072\u0075e");return _bf ,PermOwner ,nil ;
};_bd .Log .Trace ("\u0044\u0065bu\u0067\u0067\u0069n\u0067\u0020\u0061\u0075the\u006eti\u0063\u0061\u0074\u0069\u006f\u006e\u0020- \u0075\u0073\u0065\u0072\u0020\u0070\u0061s\u0073");_bf ,_bbb =_feb .alg6 (d ,pass );if _bbb !=nil {return nil ,0,_bbb ;
};if _bf !=nil {_bd .Log .Trace ("\u0074h\u0069\u0073\u002e\u0061u\u0074\u0068\u0065\u006e\u0074i\u0063a\u0074e\u0064\u0020\u003d\u0020\u0054\u0072\u0075e");return _bf ,d .P ,nil ;};return nil ,0,nil ;};

// StdHandler is an interface for standard security handlers.
type StdHandler interface{

// GenerateParams uses owner and user passwords to set encryption parameters and generate an encryption key.
// It assumes that R, P and EncryptMetadata are already set.
GenerateParams (_ed *StdEncryptDict ,_cbe ,_bgg []byte )([]byte ,error );

// Authenticate uses encryption dictionary parameters and the password to calculate
// the document encryption key. It also returns permissions that should be granted to a user.
// In case of failed authentication, it returns empty key and zero permissions with no error.
Authenticate (_cbd *StdEncryptDict ,_aa []byte )([]byte ,Permissions ,error );};type ecb struct{_cb _g .Block ;_gd int ;};

// NewHandlerR6 creates a new standard security handler for R=5 and R=6.
func NewHandlerR6 ()StdHandler {return stdHandlerR6 {}};func (_dee stdHandlerR4 )alg6 (_adf *StdEncryptDict ,_ga []byte )([]byte ,error ){var (_bdbg []byte ;_cee error ;);_cac :=_dee .alg2 (_adf ,_ga );if _adf .R ==2{_bdbg ,_cee =_dee .alg4 (_cac ,_ga );
}else if _adf .R >=3{_bdbg ,_cee =_dee .alg5 (_cac ,_ga );}else {return nil ,_gb .New ("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020R");};if _cee !=nil {return nil ,_cee ;};_bd .Log .Trace ("\u0063\u0068\u0065\u0063k:\u0020\u0025\u0020\u0078\u0020\u003d\u003d\u0020\u0025\u0020\u0078\u0020\u003f",string (_bdbg ),string (_adf .U ));
_afg :=_bdbg ;_deb :=_adf .U ;if _adf .R >=3{if len (_afg )> 16{_afg =_afg [0:16];};if len (_deb )> 16{_deb =_deb [0:16];};};if !_bc .Equal (_afg ,_deb ){return nil ,nil ;};return _cac ,nil ;};func (_fa errInvalidField )Error ()string {return _db .Sprintf ("\u0025s\u003a\u0020e\u0078\u0070\u0065\u0063t\u0065\u0064\u0020%\u0073\u0020\u0066\u0069\u0065\u006c\u0064\u0020\u0074o \u0062\u0065\u0020%\u0064\u0020b\u0079\u0074\u0065\u0073\u002c\u0020g\u006f\u0074 \u0025\u0064",_fa .Func ,_fa .Field ,_fa .Exp ,_fa .Got );
};func (_cfe stdHandlerR6 )alg11 (_efb *StdEncryptDict ,_gba []byte )([]byte ,error ){if _dcf :=_ad ("\u0061\u006c\u00671\u0031","\u0055",48,_efb .U );_dcf !=nil {return nil ,_dcf ;};_agb :=make ([]byte ,len (_gba )+8);_cedc :=copy (_agb ,_gba );_cedc +=copy (_agb [_cedc :],_efb .U [32:40]);
_adeg ,_gbc :=_cfe .alg2b (_efb .R ,_agb ,_gba ,nil );if _gbc !=nil {return nil ,_gbc ;};_adeg =_adeg [:32];if !_bc .Equal (_adeg ,_efb .U [:32]){return nil ,nil ;};return _adeg ,nil ;};func _gbd (_gge []byte )(_g .Block ,error ){_fgg ,_dge :=_gf .NewCipher (_gge );
if _dge !=nil {_bd .Log .Error ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0063\u0072\u0065\u0061\u0074\u0065\u0020A\u0045\u0053\u0020\u0063\u0069p\u0068\u0065r\u003a\u0020\u0025\u0076",_dge );
return nil ,_dge ;};return _fgg ,nil ;};

// Allowed checks if a set of permissions can be granted.
func (_ba Permissions )Allowed (p2 Permissions )bool {return _ba &p2 ==p2 };type ecbDecrypter ecb ;func (_dfa stdHandlerR6 )alg9 (_gffc *StdEncryptDict ,_cdd []byte ,_cdc []byte )error {if _adff :=_ad ("\u0061\u006c\u0067\u0039","\u004b\u0065\u0079",32,_cdd );
_adff !=nil {return _adff ;};if _dca :=_ad ("\u0061\u006c\u0067\u0039","\u0055",48,_gffc .U );_dca !=nil {return _dca ;};var _ffe [16]byte ;if _ ,_aac :=_b .ReadFull (_e .Reader ,_ffe [:]);_aac !=nil {return _aac ;};_adfb :=_ffe [0:8];_afgb :=_ffe [8:16];
_ebb :=_gffc .U [:48];_bec :=make ([]byte ,len (_cdc )+len (_adfb )+len (_ebb ));_bbff :=copy (_bec ,_cdc );_bbff +=copy (_bec [_bbff :],_adfb );_bbff +=copy (_bec [_bbff :],_ebb );_dfdg ,_aec :=_dfa .alg2b (_gffc .R ,_bec ,_cdc ,_ebb );if _aec !=nil {return _aec ;
};O :=make ([]byte ,len (_dfdg )+len (_adfb )+len (_afgb ));_bbff =copy (O ,_dfdg [:32]);_bbff +=copy (O [_bbff :],_adfb );_bbff +=copy (O [_bbff :],_afgb );_gffc .O =O ;_bbff =len (_cdc );_bbff +=copy (_bec [_bbff :],_afgb );_dfdg ,_aec =_dfa .alg2b (_gffc .R ,_bec ,_cdc ,_ebb );
if _aec !=nil {return _aec ;};_ggd ,_aec :=_gbd (_dfdg [:32]);if _aec !=nil {return _aec ;};_dea :=make ([]byte ,_gf .BlockSize );_gdbe :=_g .NewCBCEncrypter (_ggd ,_dea );OE :=make ([]byte ,32);_gdbe .CryptBlocks (OE ,_cdd [:32]);_gffc .OE =OE ;return nil ;
};func (stdHandlerR4 )paddedPass (_ade []byte )[]byte {_ac :=make ([]byte ,32);_ebd :=copy (_ac ,_ade );for ;_ebd < 32;_ebd ++{_ac [_ebd ]=_gfd [_ebd -len (_ade )];};return _ac ;};func (_deeg stdHandlerR6 )alg2b (R int ,_egd ,_ggc ,_ead []byte )([]byte ,error ){if R ==5{return _gcfe (_egd );
};return _daa (_egd ,_ggc ,_ead );};func (_ec *ecbEncrypter )BlockSize ()int {return _ec ._gd };func _dba (_cf _g .Block )*ecb {return &ecb {_cb :_cf ,_gd :_cf .BlockSize ()}};func (_bede stdHandlerR6 )alg12 (_ddc *StdEncryptDict ,_ccf []byte )([]byte ,error ){if _dbed :=_ad ("\u0061\u006c\u00671\u0032","\u0055",48,_ddc .U );
_dbed !=nil {return nil ,_dbed ;};if _acd :=_ad ("\u0061\u006c\u00671\u0032","\u004f",48,_ddc .O );_acd !=nil {return nil ,_acd ;};_ggeb :=make ([]byte ,len (_ccf )+8+48);_afb :=copy (_ggeb ,_ccf );_afb +=copy (_ggeb [_afb :],_ddc .O [32:40]);_afb +=copy (_ggeb [_afb :],_ddc .U [0:48]);
_gac ,_ebge :=_bede .alg2b (_ddc .R ,_ggeb ,_ccf ,_ddc .U [0:48]);if _ebge !=nil {return nil ,_ebge ;};_gac =_gac [:32];if !_bc .Equal (_gac ,_ddc .O [:32]){return nil ,nil ;};return _gac ,nil ;};type ecbEncrypter ecb ;func (_bde stdHandlerR4 )alg3Key (R int ,_ebg []byte )[]byte {_bdb :=_c .New ();
_ecc :=_bde .paddedPass (_ebg );_bdb .Write (_ecc );if R >=3{for _ceg :=0;_ceg < 50;_ceg ++{_cec :=_bdb .Sum (nil );_bdb =_c .New ();_bdb .Write (_cec );};};_fg :=_bdb .Sum (nil );if R ==2{_fg =_fg [0:5];}else {_fg =_fg [0:_bde .Length /8];};return _fg ;
};type errInvalidField struct{Func string ;Field string ;Exp int ;Got int ;};func (_aef stdHandlerR6 )alg8 (_bcf *StdEncryptDict ,_efg []byte ,_abc []byte )error {if _caf :=_ad ("\u0061\u006c\u0067\u0038","\u004b\u0065\u0079",32,_efg );_caf !=nil {return _caf ;
};var _gec [16]byte ;if _ ,_gab :=_b .ReadFull (_e .Reader ,_gec [:]);_gab !=nil {return _gab ;};_dfdd :=_gec [0:8];_eef :=_gec [8:16];_gag :=make ([]byte ,len (_abc )+len (_dfdd ));_eefa :=copy (_gag ,_abc );copy (_gag [_eefa :],_dfdd );_eab ,_dbfa :=_aef .alg2b (_bcf .R ,_gag ,_abc ,nil );
if _dbfa !=nil {return _dbfa ;};U :=make ([]byte ,len (_eab )+len (_dfdd )+len (_eef ));_eefa =copy (U ,_eab [:32]);_eefa +=copy (U [_eefa :],_dfdd );copy (U [_eefa :],_eef );_bcf .U =U ;_eefa =len (_abc );copy (_gag [_eefa :],_eef );_eab ,_dbfa =_aef .alg2b (_bcf .R ,_gag ,_abc ,nil );
if _dbfa !=nil {return _dbfa ;};_fab ,_dbfa :=_gbd (_eab [:32]);if _dbfa !=nil {return _dbfa ;};_fded :=make ([]byte ,_gf .BlockSize );_bac :=_g .NewCBCEncrypter (_fab ,_fded );UE :=make ([]byte ,32);_bac .CryptBlocks (UE ,_efg [:32]);_bcf .UE =UE ;return nil ;
};

// NewHandlerR4 creates a new standard security handler for R<=4.
func NewHandlerR4 (id0 string ,length int )StdHandler {return stdHandlerR4 {ID0 :id0 ,Length :length }};func (_dbg stdHandlerR4 )alg4 (_cg []byte ,_ded []byte )([]byte ,error ){_cad ,_cccc :=_fb .NewCipher (_cg );if _cccc !=nil {return nil ,_gb .New ("\u0066a\u0069l\u0065\u0064\u0020\u0072\u0063\u0034\u0020\u0063\u0069\u0070\u0068");
};_aad :=[]byte (_gfd );_edc :=make ([]byte ,len (_aad ));_cad .XORKeyStream (_edc ,_aad );return _edc ,nil ;};

// AuthEvent is an event type that triggers authentication.
type AuthEvent string ;

// GenerateParams is the algorithm opposite to alg2a (R>=5).
// It generates U,O,UE,OE,Perms fields using AESv3 encryption.
// There is no algorithm number assigned to this function in the spec.
// It expects R, P and EncryptMetadata fields to be set.
func (_gbca stdHandlerR6 )GenerateParams (d *StdEncryptDict ,opass ,upass []byte )([]byte ,error ){_aeef :=make ([]byte ,32);if _ ,_aefd :=_b .ReadFull (_e .Reader ,_aeef );_aefd !=nil {return nil ,_aefd ;};d .U =nil ;d .O =nil ;d .UE =nil ;d .OE =nil ;
d .Perms =nil ;if len (upass )> 127{upass =upass [:127];};if len (opass )> 127{opass =opass [:127];};if _bfg :=_gbca .alg8 (d ,_aeef ,upass );_bfg !=nil {return nil ,_bfg ;};if _gccb :=_gbca .alg9 (d ,_aeef ,opass );_gccb !=nil {return nil ,_gccb ;};if d .R ==5{return _aeef ,nil ;
};if _acfg :=_gbca .alg10 (d ,_aeef );_acfg !=nil {return nil ,_acfg ;};return _aeef ,nil ;};

// StdEncryptDict is a set of additional fields used in standard encryption dictionary.
type StdEncryptDict struct{R int ;P Permissions ;EncryptMetadata bool ;O ,U []byte ;OE ,UE []byte ;Perms []byte ;};func _cca (_bdc []byte ,_bbf int ){_dfd :=_bbf ;for _dfd < len (_bdc ){copy (_bdc [_dfd :],_bdc [:_dfd ]);_dfd *=2;};};func _gc (_eb _g .Block )_g .BlockMode {return (*ecbEncrypter )(_dba (_eb ))};
func (_gbf stdHandlerR6 )alg10 (_cde *StdEncryptDict ,_cgc []byte )error {if _ffa :=_ad ("\u0061\u006c\u00671\u0030","\u004b\u0065\u0079",32,_cgc );_ffa !=nil {return _ffa ;};_ag :=uint64 (uint32 (_cde .P ))|(_de .MaxUint32 <<32);Perms :=make ([]byte ,16);
_fe .LittleEndian .PutUint64 (Perms [:8],_ag );if _cde .EncryptMetadata {Perms [8]='T';}else {Perms [8]='F';};copy (Perms [9:12],"\u0061\u0064\u0062");if _ ,_fda :=_b .ReadFull (_e .Reader ,Perms [12:16]);_fda !=nil {return _fda ;};_dcg ,_cgfc :=_gbd (_cgc [:32]);
if _cgfc !=nil {return _cgfc ;};_fbg :=_gc (_dcg );_fbg .CryptBlocks (Perms ,Perms );_cde .Perms =Perms [:16];return nil ;};func _daa (_fcf ,_beag ,_fgge []byte )([]byte ,error ){var (_deg ,_beage ,_eba _a .Hash ;);_deg =_af .New ();_fff :=make ([]byte ,64);
_bef :=_deg ;_bef .Write (_fcf );K :=_bef .Sum (_fff [:0]);_acf :=make ([]byte ,64*(127+64+48));_gda :=func (_ebgf int )([]byte ,error ){_badb :=len (_beag )+len (K )+len (_fgge );_dbe :=_acf [:_badb ];_cbee :=copy (_dbe ,_beag );_cbee +=copy (_dbe [_cbee :],K [:]);
_cbee +=copy (_dbe [_cbee :],_fgge );if _cbee !=_badb {_bd .Log .Error ("E\u0052\u0052\u004f\u0052\u003a\u0020u\u006e\u0065\u0078\u0070\u0065\u0063t\u0065\u0064\u0020\u0072\u006f\u0075\u006ed\u0020\u0069\u006e\u0070\u0075\u0074\u0020\u0073\u0069\u007ae\u002e");
return nil ,_gb .New ("\u0077\u0072\u006f\u006e\u0067\u0020\u0073\u0069\u007a\u0065");};K1 :=_acf [:_badb *64];_cca (K1 ,_badb );_cfbf ,_gdbc :=_gbd (K [0:16]);if _gdbc !=nil {return nil ,_gdbc ;};_adg :=_g .NewCBCEncrypter (_cfbf ,K [16:32]);_adg .CryptBlocks (K1 ,K1 );
E :=K1 ;_dc :=0;for _fgff :=0;_fgff < 16;_fgff ++{_dc +=int (E [_fgff ]%3);};var _aff _a .Hash ;switch _dc %3{case 0:_aff =_deg ;case 1:if _beage ==nil {_beage =_f .New384 ();};_aff =_beage ;case 2:if _eba ==nil {_eba =_f .New ();};_aff =_eba ;};_aff .Reset ();
_aff .Write (E );K =_aff .Sum (_fff [:0]);return E ,nil ;};for _caed :=0;;{E ,_bgb :=_gda (_caed );if _bgb !=nil {return nil ,_bgb ;};_fdef :=E [len (E )-1];_caed ++;if _caed >=64&&_fdef <=uint8 (_caed -32){break ;};};return K [:32],nil ;};func (_ae *ecbEncrypter )CryptBlocks (dst ,src []byte ){if len (src )%_ae ._gd !=0{_bd .Log .Error ("\u0045\u0052\u0052\u004f\u0052:\u0020\u0045\u0043\u0042\u0020\u0065\u006e\u0063\u0072\u0079\u0070\u0074\u003a \u0069\u006e\u0070\u0075\u0074\u0020\u006e\u006f\u0074\u0020\u0066\u0075\u006c\u006c\u0020\u0062\u006c\u006f\u0063\u006b\u0073");
return ;};if len (dst )< len (src ){_bd .Log .Error ("\u0045R\u0052\u004fR\u003a\u0020\u0045C\u0042\u0020\u0065\u006e\u0063\u0072\u0079p\u0074\u003a\u0020\u006f\u0075\u0074p\u0075\u0074\u0020\u0073\u006d\u0061\u006c\u006c\u0065\u0072\u0020t\u0068\u0061\u006e\u0020\u0069\u006e\u0070\u0075\u0074");
return ;};for len (src )> 0{_ae ._cb .Encrypt (dst ,src [:_ae ._gd ]);src =src [_ae ._gd :];dst =dst [_ae ._gd :];};};const (PermOwner =Permissions (_de .MaxUint32 );PermPrinting =Permissions (1<<2);PermModify =Permissions (1<<3);PermExtractGraphics =Permissions (1<<4);
PermAnnotate =Permissions (1<<5);PermFillForms =Permissions (1<<8);PermDisabilityExtract =Permissions (1<<9);PermRotateInsert =Permissions (1<<10);PermFullPrintQuality =Permissions (1<<11););type stdHandlerR6 struct{};func (_ece stdHandlerR6 )alg13 (_aedd *StdEncryptDict ,_ddfb []byte )error {if _bdd :=_ad ("\u0061\u006c\u00671\u0033","\u004b\u0065\u0079",32,_ddfb );
_bdd !=nil {return _bdd ;};if _bdbd :=_ad ("\u0061\u006c\u00671\u0033","\u0050\u0065\u0072m\u0073",16,_aedd .Perms );_bdbd !=nil {return _bdbd ;};_cgg :=make ([]byte ,16);copy (_cgg ,_aedd .Perms [:16]);_gbff ,_abe :=_gf .NewCipher (_ddfb [:32]);if _abe !=nil {return _abe ;
};_fdeg :=_bg (_gbff );_fdeg .CryptBlocks (_cgg ,_cgg );if !_bc .Equal (_cgg [9:12],[]byte ("\u0061\u0064\u0062")){return _gb .New ("\u0064\u0065\u0063o\u0064\u0065\u0064\u0020p\u0065\u0072\u006d\u0069\u0073\u0073\u0069o\u006e\u0073\u0020\u0061\u0072\u0065\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064");
};_cea :=Permissions (_fe .LittleEndian .Uint32 (_cgg [0:4]));if _cea !=_aedd .P {return _gb .New ("\u0070\u0065r\u006d\u0069\u0073\u0073\u0069\u006f\u006e\u0073\u0020\u0076\u0061\u006c\u0069\u0064\u0061\u0074\u0069\u006f\u006e\u0020\u0066\u0061il\u0065\u0064");
};var _cdcg bool ;if _cgg [8]=='T'{_cdcg =true ;}else if _cgg [8]=='F'{_cdcg =false ;}else {return _gb .New ("\u0064\u0065\u0063\u006f\u0064\u0065\u0064 \u006d\u0065\u0074a\u0064\u0061\u0074\u0061 \u0065\u006e\u0063\u0072\u0079\u0070\u0074\u0069\u006f\u006e\u0020\u0066\u006c\u0061\u0067\u0020\u0069\u0073\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064");
};if _cdcg !=_aedd .EncryptMetadata {return _gb .New ("\u006d\u0065t\u0061\u0064\u0061\u0074a\u0020\u0065n\u0063\u0072\u0079\u0070\u0074\u0069\u006f\u006e \u0076\u0061\u006c\u0069\u0064\u0061\u0074\u0069\u006f\u006e\u0020\u0066a\u0069\u006c\u0065\u0064");
};return nil ;};func (_cd stdHandlerR6 )alg2a (_ecd *StdEncryptDict ,_bea []byte )([]byte ,Permissions ,error ){if _fde :=_ad ("\u0061\u006c\u00672\u0061","\u004f",48,_ecd .O );_fde !=nil {return nil ,0,_fde ;};if _efc :=_ad ("\u0061\u006c\u00672\u0061","\u0055",48,_ecd .U );
_efc !=nil {return nil ,0,_efc ;};if len (_bea )> 127{_bea =_bea [:127];};_bcc ,_ced :=_cd .alg12 (_ecd ,_bea );if _ced !=nil {return nil ,0,_ced ;};var (_afd []byte ;_gcc []byte ;_eeb []byte ;);var _cbea Permissions ;if len (_bcc )!=0{_cbea =PermOwner ;
_gcef :=make ([]byte ,len (_bea )+8+48);_beb :=copy (_gcef ,_bea );_beb +=copy (_gcef [_beb :],_ecd .O [40:48]);copy (_gcef [_beb :],_ecd .U [0:48]);_afd =_gcef ;_gcc =_ecd .OE ;_eeb =_ecd .U [0:48];}else {_bcc ,_ced =_cd .alg11 (_ecd ,_bea );if _ced ==nil &&len (_bcc )==0{_bcc ,_ced =_cd .alg11 (_ecd ,[]byte (""));
};if _ced !=nil {return nil ,0,_ced ;}else if len (_bcc )==0{return nil ,0,nil ;};_cbea =_ecd .P ;_bda :=make ([]byte ,len (_bea )+8);_ege :=copy (_bda ,_bea );copy (_bda [_ege :],_ecd .U [40:48]);_afd =_bda ;_gcc =_ecd .UE ;_eeb =nil ;};if _cfb :=_ad ("\u0061\u006c\u00672\u0061","\u004b\u0065\u0079",32,_gcc );
_cfb !=nil {return nil ,0,_cfb ;};_gcc =_gcc [:32];_bed ,_ced :=_cd .alg2b (_ecd .R ,_afd ,_bea ,_eeb );if _ced !=nil {return nil ,0,_ced ;};_ddf ,_ced :=_gf .NewCipher (_bed [:32]);if _ced !=nil {return nil ,0,_ced ;};_edd :=make ([]byte ,_gf .BlockSize );
_ccg :=_g .NewCBCDecrypter (_ddf ,_edd );_edcc :=make ([]byte ,32);_ccg .CryptBlocks (_edcc ,_gcc );if _ecd .R ==5{return _edcc ,_cbea ,nil ;};_ced =_cd .alg13 (_ecd ,_edcc );if _ced !=nil {return nil ,0,_ced ;};return _edcc ,_cbea ,nil ;};type stdHandlerR4 struct{Length int ;
ID0 string ;};

// Authenticate implements StdHandler interface.
func (_bedg stdHandlerR6 )Authenticate (d *StdEncryptDict ,pass []byte )([]byte ,Permissions ,error ){return _bedg .alg2a (d ,pass );};

// GenerateParams generates and sets O and U parameters for the encryption dictionary.
// It expects R, P and EncryptMetadata fields to be set.
func (_eeg stdHandlerR4 )GenerateParams (d *StdEncryptDict ,opass ,upass []byte )([]byte ,error ){O ,_df :=_eeg .alg3 (d .R ,upass ,opass );if _df !=nil {_bd .Log .Debug ("\u0045R\u0052\u004fR\u003a\u0020\u0045r\u0072\u006f\u0072\u0020\u0067\u0065\u006ee\u0072\u0061\u0074\u0069\u006e\u0067 \u004f\u0020\u0066\u006f\u0072\u0020\u0065\u006e\u0063\u0072\u0079p\u0074\u0069\u006f\u006e\u0020\u0028\u0025\u0073\u0029",_df );
return nil ,_df ;};d .O =O ;_bd .Log .Trace ("\u0067\u0065\u006e\u0020\u004f\u003a\u0020\u0025\u0020\u0078",O );_adb :=_eeg .alg2 (d ,upass );U ,_df :=_eeg .alg5 (_adb ,upass );if _df !=nil {_bd .Log .Debug ("\u0045R\u0052\u004fR\u003a\u0020\u0045r\u0072\u006f\u0072\u0020\u0067\u0065\u006ee\u0072\u0061\u0074\u0069\u006e\u0067 \u004f\u0020\u0066\u006f\u0072\u0020\u0065\u006e\u0063\u0072\u0079p\u0074\u0069\u006f\u006e\u0020\u0028\u0025\u0073\u0029",_df );
return nil ,_df ;};d .U =U ;_bd .Log .Trace ("\u0067\u0065\u006e\u0020\u0055\u003a\u0020\u0025\u0020\u0078",U );return _adb ,nil ;};const _gfd ="\x28\277\116\136\x4e\x75\x8a\x41\x64\000\x4e\x56\377"+"\xfa\001\010\056\x2e\x00\xb6\xd0\x68\076\x80\x2f\014"+"\251\xfe\x64\x53\x69\172";


// Permissions is a bitmask of access permissions for a PDF file.
type Permissions uint32 ;func _ad (_gca ,_cc string ,_gff int ,_ccc []byte )error {if len (_ccc )< _gff {return errInvalidField {Func :_gca ,Field :_cc ,Exp :_gff ,Got :len (_ccc )};};return nil ;};func (_fd stdHandlerR4 )alg7 (_bcb *StdEncryptDict ,_fec []byte )([]byte ,error ){_gbg :=_fd .alg3Key (_bcb .R ,_fec );
_edce :=make ([]byte ,len (_bcb .O ));if _bcb .R ==2{_fgf ,_dd :=_fb .NewCipher (_gbg );if _dd !=nil {return nil ,_gb .New ("\u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0063\u0069\u0070\u0068\u0065\u0072");};_fgf .XORKeyStream (_edce ,_bcb .O );}else if _bcb .R >=3{_dda :=append ([]byte {},_bcb .O ...);
for _ff :=0;_ff < 20;_ff ++{_afa :=append ([]byte {},_gbg ...);for _gce :=0;_gce < len (_gbg );_gce ++{_afa [_gce ]^=byte (19-_ff );};_ggb ,_ee :=_fb .NewCipher (_afa );if _ee !=nil {return nil ,_gb .New ("\u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0063\u0069\u0070\u0068\u0065\u0072");
};_ggb .XORKeyStream (_edce ,_dda );_dda =append ([]byte {},_edce ...);};}else {return nil ,_gb .New ("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020R");};_egf ,_cae :=_fd .alg6 (_bcb ,_edce );if _cae !=nil {return nil ,nil ;};return _egf ,nil ;};func (_bcd stdHandlerR4 )alg5 (_cgf []byte ,_da []byte )([]byte ,error ){_aede :=_c .New ();
_aede .Write ([]byte (_gfd ));_aede .Write ([]byte (_bcd .ID0 ));_ebdb :=_aede .Sum (nil );_bd .Log .Trace ("\u0061\u006c\u0067\u0035");_bd .Log .Trace ("\u0065k\u0065\u0079\u003a\u0020\u0025\u0020x",_cgf );_bd .Log .Trace ("\u0049D\u003a\u0020\u0025\u0020\u0078",_bcd .ID0 );
if len (_ebdb )!=16{return nil ,_gb .New ("\u0068a\u0073\u0068\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006eo\u0074\u0020\u0031\u0036\u0020\u0062\u0079\u0074\u0065\u0073");};_bae ,_ef :=_fb .NewCipher (_cgf );if _ef !=nil {return nil ,_gb .New ("\u0066a\u0069l\u0065\u0064\u0020\u0072\u0063\u0034\u0020\u0063\u0069\u0070\u0068");
};_dbf :=make ([]byte ,16);_bae .XORKeyStream (_dbf ,_ebdb );_caa :=make ([]byte ,len (_cgf ));for _fca :=0;_fca < 19;_fca ++{for _fed :=0;_fed < len (_cgf );_fed ++{_caa [_fed ]=_cgf [_fed ]^byte (_fca +1);};_bae ,_ef =_fb .NewCipher (_caa );if _ef !=nil {return nil ,_gb .New ("\u0066a\u0069l\u0065\u0064\u0020\u0072\u0063\u0034\u0020\u0063\u0069\u0070\u0068");
};_bae .XORKeyStream (_dbf ,_dbf );_bd .Log .Trace ("\u0069\u0020\u003d\u0020\u0025\u0064\u002c\u0020\u0065\u006b\u0065\u0079:\u0020\u0025\u0020\u0078",_fca ,_caa );_bd .Log .Trace ("\u0069\u0020\u003d\u0020\u0025\u0064\u0020\u002d\u003e\u0020\u0025\u0020\u0078",_fca ,_dbf );
};_ecfb :=make ([]byte ,32);for _ecg :=0;_ecg < 16;_ecg ++{_ecfb [_ecg ]=_dbf [_ecg ];};_ ,_ef =_e .Read (_ecfb [16:32]);if _ef !=nil {return nil ,_gb .New ("\u0066a\u0069\u006c\u0065\u0064 \u0074\u006f\u0020\u0067\u0065n\u0020r\u0061n\u0064\u0020\u006e\u0075\u006d\u0062\u0065r");
};return _ecfb ,nil ;};func (_aee stdHandlerR4 )alg3 (R int ,_ea ,_aegd []byte )([]byte ,error ){var _be []byte ;if len (_aegd )> 0{_be =_aee .alg3Key (R ,_aegd );}else {_be =_aee .alg3Key (R ,_ea );};_bb ,_gdb :=_fb .NewCipher (_be );if _gdb !=nil {return nil ,_gb .New ("\u0066a\u0069l\u0065\u0064\u0020\u0072\u0063\u0034\u0020\u0063\u0069\u0070\u0068");
};_gcd :=_aee .paddedPass (_ea );_bbe :=make ([]byte ,len (_gcd ));_bb .XORKeyStream (_bbe ,_gcd );if R >=3{_dgf :=make ([]byte ,len (_be ));for _aed :=0;_aed < 19;_aed ++{for _eg :=0;_eg < len (_be );_eg ++{_dgf [_eg ]=_be [_eg ]^byte (_aed +1);};_ca ,_dgfc :=_fb .NewCipher (_dgf );
if _dgfc !=nil {return nil ,_gb .New ("\u0066a\u0069l\u0065\u0064\u0020\u0072\u0063\u0034\u0020\u0063\u0069\u0070\u0068");};_ca .XORKeyStream (_bbe ,_bbe );};};return _bbe ,nil ;};