package avcodec

// #include <codecID.h>
import "C"

type CodecID C.enum_AVCodecID

var (
	CodecIDMPEG1VIDEO      = CodecID(C.AV_CODEC_ID_MPEG1VIDEO)
	CodecIDMPEG2VIDEO      = CodecID(C.AV_CODEC_ID_MPEG2VIDEO)
	CodecIDH261            = CodecID(C.AV_CODEC_ID_H261)
	CodecIDH263            = CodecID(C.AV_CODEC_ID_H263)
	CodecIDRV10            = CodecID(C.AV_CODEC_ID_RV10)
	CodecIDRV20            = CodecID(C.AV_CODEC_ID_RV20)
	CodecIDMJPEG           = CodecID(C.AV_CODEC_ID_MJPEG)
	CodecIDMJPEGB          = CodecID(C.AV_CODEC_ID_MJPEGB)
	CodecIDLJPEG           = CodecID(C.AV_CODEC_ID_LJPEG)
	CodecIDSP5X            = CodecID(C.AV_CODEC_ID_SP5X)
	CodecIDJPEGLS          = CodecID(C.AV_CODEC_ID_JPEGLS)
	CodecIDMPEG4           = CodecID(C.AV_CODEC_ID_MPEG4)
	CodecIDRAWVIDEO        = CodecID(C.AV_CODEC_ID_RAWVIDEO)
	CodecIDMSMPEG4V1       = CodecID(C.AV_CODEC_ID_MSMPEG4V1)
	CodecIDMSMPEG4V2       = CodecID(C.AV_CODEC_ID_MSMPEG4V2)
	CodecIDMSMPEG4V3       = CodecID(C.AV_CODEC_ID_MSMPEG4V3)
	CodecIDWMV1            = CodecID(C.AV_CODEC_ID_WMV1)
	CodecIDWMV2            = CodecID(C.AV_CODEC_ID_WMV2)
	CodecIDH263P           = CodecID(C.AV_CODEC_ID_H263P)
	CodecIDH263I           = CodecID(C.AV_CODEC_ID_H263I)
	CodecIDFLV1            = CodecID(C.AV_CODEC_ID_FLV1)
	CodecIDSVQ1            = CodecID(C.AV_CODEC_ID_SVQ1)
	CodecIDSVQ3            = CodecID(C.AV_CODEC_ID_SVQ3)
	CodecIDDVVIDEO         = CodecID(C.AV_CODEC_ID_DVVIDEO)
	CodecIDHUFFYUV         = CodecID(C.AV_CODEC_ID_HUFFYUV)
	CodecIDCYUV            = CodecID(C.AV_CODEC_ID_CYUV)
	CodecIDH264            = CodecID(C.AV_CODEC_ID_H264)
	CodecIDINDEO3          = CodecID(C.AV_CODEC_ID_INDEO3)
	CodecIDVP3             = CodecID(C.AV_CODEC_ID_VP3)
	CodecIDTHEORA          = CodecID(C.AV_CODEC_ID_THEORA)
	CodecIDASV1            = CodecID(C.AV_CODEC_ID_ASV1)
	CodecIDASV2            = CodecID(C.AV_CODEC_ID_ASV2)
	CodecIDFFV1            = CodecID(C.AV_CODEC_ID_FFV1)
	CodecID4XM             = CodecID(C.AV_CODEC_ID_4XM)
	CodecIDVCR1            = CodecID(C.AV_CODEC_ID_VCR1)
	CodecIDCLJR            = CodecID(C.AV_CODEC_ID_CLJR)
	CodecIDMDEC            = CodecID(C.AV_CODEC_ID_MDEC)
	CodecIDROQ             = CodecID(C.AV_CODEC_ID_ROQ)
	CodecIDINTERPLAY_VIDEO = CodecID(C.AV_CODEC_ID_INTERPLAY_VIDEO)
	CodecIDXAN_WC3         = CodecID(C.AV_CODEC_ID_XAN_WC3)
	CodecIDXAN_WC4         = CodecID(C.AV_CODEC_ID_XAN_WC4)
	CodecIDRPZA            = CodecID(C.AV_CODEC_ID_RPZA)
	CodecIDCINEPAK         = CodecID(C.AV_CODEC_ID_CINEPAK)
	CodecIDWS_VQA          = CodecID(C.AV_CODEC_ID_WS_VQA)
	CodecIDMSRLE           = CodecID(C.AV_CODEC_ID_MSRLE)
	CodecIDMSVIDEO1        = CodecID(C.AV_CODEC_ID_MSVIDEO1)
	CodecIDIDCIN           = CodecID(C.AV_CODEC_ID_IDCIN)
	CodecID8BPS            = CodecID(C.AV_CODEC_ID_8BPS)
	CodecIDSMC             = CodecID(C.AV_CODEC_ID_SMC)
	CodecIDFLIC            = CodecID(C.AV_CODEC_ID_FLIC)
	CodecIDTRUEMOTION1     = CodecID(C.AV_CODEC_ID_TRUEMOTION1)
	CodecIDVMDVIDEO        = CodecID(C.AV_CODEC_ID_VMDVIDEO)
	CodecIDMSZH            = CodecID(C.AV_CODEC_ID_MSZH)
	CodecIDZLIB            = CodecID(C.AV_CODEC_ID_ZLIB)
	CodecIDQTRLE           = CodecID(C.AV_CODEC_ID_QTRLE)
	CodecIDTSCC            = CodecID(C.AV_CODEC_ID_TSCC)
	CodecIDULTI            = CodecID(C.AV_CODEC_ID_ULTI)
	CodecIDQDRAW           = CodecID(C.AV_CODEC_ID_QDRAW)
	CodecIDVIXL            = CodecID(C.AV_CODEC_ID_VIXL)
	CodecIDQPEG            = CodecID(C.AV_CODEC_ID_QPEG)
	CodecIDPNG             = CodecID(C.AV_CODEC_ID_PNG)
	CodecIDPPM             = CodecID(C.AV_CODEC_ID_PPM)
	CodecIDPBM             = CodecID(C.AV_CODEC_ID_PBM)
	CodecIDPGM             = CodecID(C.AV_CODEC_ID_PGM)
	CodecIDPGMYUV          = CodecID(C.AV_CODEC_ID_PGMYUV)
	CodecIDPAM             = CodecID(C.AV_CODEC_ID_PAM)
	CodecIDFFVHUFF         = CodecID(C.AV_CODEC_ID_FFVHUFF)
	CodecIDRV30            = CodecID(C.AV_CODEC_ID_RV30)
	CodecIDRV40            = CodecID(C.AV_CODEC_ID_RV40)
	CodecIDVC1             = CodecID(C.AV_CODEC_ID_VC1)
	CodecIDWMV3            = CodecID(C.AV_CODEC_ID_WMV3)
	CodecIDLOCO            = CodecID(C.AV_CODEC_ID_LOCO)
	CodecIDWNV1            = CodecID(C.AV_CODEC_ID_WNV1)
	CodecIDAASC            = CodecID(C.AV_CODEC_ID_AASC)
	CodecIDINDEO2          = CodecID(C.AV_CODEC_ID_INDEO2)
	CodecIDFRAPS           = CodecID(C.AV_CODEC_ID_FRAPS)
	CodecIDTRUEMOTION2     = CodecID(C.AV_CODEC_ID_TRUEMOTION2)
	CodecIDBMP             = CodecID(C.AV_CODEC_ID_BMP)
	CodecIDCSCD            = CodecID(C.AV_CODEC_ID_CSCD)
	CodecIDMMVIDEO         = CodecID(C.AV_CODEC_ID_MMVIDEO)
	CodecIDZMBV            = CodecID(C.AV_CODEC_ID_ZMBV)
	CodecIDAVS             = CodecID(C.AV_CODEC_ID_AVS)
	CodecIDSMACKVIDEO      = CodecID(C.AV_CODEC_ID_SMACKVIDEO)
	CodecIDNUV             = CodecID(C.AV_CODEC_ID_NUV)
	CodecIDKMVC            = CodecID(C.AV_CODEC_ID_KMVC)
	CodecIDFLASHSV         = CodecID(C.AV_CODEC_ID_FLASHSV)
	CodecIDCAVS            = CodecID(C.AV_CODEC_ID_CAVS)
	CodecIDJPEG2000        = CodecID(C.AV_CODEC_ID_JPEG2000)
	CodecIDVMNC            = CodecID(C.AV_CODEC_ID_VMNC)
	CodecIDVP5             = CodecID(C.AV_CODEC_ID_VP5)
	CodecIDVP6             = CodecID(C.AV_CODEC_ID_VP6)
	CodecIDVP6F            = CodecID(C.AV_CODEC_ID_VP6F)
	CodecIDTARGA           = CodecID(C.AV_CODEC_ID_TARGA)
	CodecIDDSICINVIDEO     = CodecID(C.AV_CODEC_ID_DSICINVIDEO)
	CodecIDTIERTEXSEQVIDEO = CodecID(C.AV_CODEC_ID_TIERTEXSEQVIDEO)
	CodecIDTIFF            = CodecID(C.AV_CODEC_ID_TIFF)
	CodecIDGIF             = CodecID(C.AV_CODEC_ID_GIF)
	CodecIDDXA             = CodecID(C.AV_CODEC_ID_DXA)
	CodecIDDNXHD           = CodecID(C.AV_CODEC_ID_DNXHD)
	CodecIDTHP             = CodecID(C.AV_CODEC_ID_THP)
	CodecIDSGI             = CodecID(C.AV_CODEC_ID_SGI)
	CodecIDC93             = CodecID(C.AV_CODEC_ID_C93)
	CodecIDBETHSOFTVID     = CodecID(C.AV_CODEC_ID_BETHSOFTVID)
	CodecIDPTX             = CodecID(C.AV_CODEC_ID_PTX)
	CodecIDTXD             = CodecID(C.AV_CODEC_ID_TXD)
	CodecIDVP6A            = CodecID(C.AV_CODEC_ID_VP6A)
	CodecIDAMV             = CodecID(C.AV_CODEC_ID_AMV)
	CodecIDVB              = CodecID(C.AV_CODEC_ID_VB)
	CodecIDPCX             = CodecID(C.AV_CODEC_ID_PCX)
	CodecIDSUNRAST         = CodecID(C.AV_CODEC_ID_SUNRAST)
	CodecIDINDEO4          = CodecID(C.AV_CODEC_ID_INDEO4)
	CodecIDINDEO5          = CodecID(C.AV_CODEC_ID_INDEO5)
	CodecIDMIMIC           = CodecID(C.AV_CODEC_ID_MIMIC)
	CodecIDRL2             = CodecID(C.AV_CODEC_ID_RL2)
	CodecIDESCAPE124       = CodecID(C.AV_CODEC_ID_ESCAPE124)
	CodecIDDIRAC           = CodecID(C.AV_CODEC_ID_DIRAC)
	CodecIDBFI             = CodecID(C.AV_CODEC_ID_BFI)
	CodecIDCMV             = CodecID(C.AV_CODEC_ID_CMV)
	CodecIDMOTIONPIXELS    = CodecID(C.AV_CODEC_ID_MOTIONPIXELS)
	CodecIDTGV             = CodecID(C.AV_CODEC_ID_TGV)
	CodecIDTGQ             = CodecID(C.AV_CODEC_ID_TGQ)
	CodecIDTQI             = CodecID(C.AV_CODEC_ID_TQI)
	CodecIDAURA            = CodecID(C.AV_CODEC_ID_AURA)
	CodecIDAURA2           = CodecID(C.AV_CODEC_ID_AURA2)
	CodecIDV210X           = CodecID(C.AV_CODEC_ID_V210X)
	CodecIDTMV             = CodecID(C.AV_CODEC_ID_TMV)
	CodecIDV210            = CodecID(C.AV_CODEC_ID_V210)
	CodecIDDPX             = CodecID(C.AV_CODEC_ID_DPX)
	CodecIDMAD             = CodecID(C.AV_CODEC_ID_MAD)
	CodecIDFRWU            = CodecID(C.AV_CODEC_ID_FRWU)
	CodecIDFLASHSV2        = CodecID(C.AV_CODEC_ID_FLASHSV2)
	CodecIDCDGRAPHICS      = CodecID(C.AV_CODEC_ID_CDGRAPHICS)
	CodecIDR210            = CodecID(C.AV_CODEC_ID_R210)
	CodecIDANM             = CodecID(C.AV_CODEC_ID_ANM)
	CodecIDBINKVIDEO       = CodecID(C.AV_CODEC_ID_BINKVIDEO)
	CodecIDIFF_ILBM        = CodecID(C.AV_CODEC_ID_IFF_ILBM)

	CodecIDKGV1      = CodecID(C.AV_CODEC_ID_KGV1)
	CodecIDYOP       = CodecID(C.AV_CODEC_ID_YOP)
	CodecIDVP8       = CodecID(C.AV_CODEC_ID_VP8)
	CodecIDPICTOR    = CodecID(C.AV_CODEC_ID_PICTOR)
	CodecIDANSI      = CodecID(C.AV_CODEC_ID_ANSI)
	CodecIDA64MULTI  = CodecID(C.AV_CODEC_ID_A64_MULTI)
	CodecIDA64MULTI5 = CodecID(C.AV_CODEC_ID_A64_MULTI5)
	CodecIDR10K      = CodecID(C.AV_CODEC_ID_R10K)
	CodecIDMXPEG     = CodecID(C.AV_CODEC_ID_MXPEG)
	CodecIDLAGARITH  = CodecID(C.AV_CODEC_ID_LAGARITH)
	CodecIDPRORES    = CodecID(C.AV_CODEC_ID_PRORES)
	CodecIDJV        = CodecID(C.AV_CODEC_ID_JV)
	CodecIDDFA       = CodecID(C.AV_CODEC_ID_DFA)
	CodecIDWMV3IMAGE = CodecID(C.AV_CODEC_ID_WMV3IMAGE)
	CodecIDVC1IMAGE  = CodecID(C.AV_CODEC_ID_VC1IMAGE)
	CodecIDUTVIDEO   = CodecID(C.AV_CODEC_ID_UTVIDEO)
	CodecIDBMVVIDEO  = CodecID(C.AV_CODEC_ID_BMV_VIDEO)
	CodecIDVBLE      = CodecID(C.AV_CODEC_ID_VBLE)
	CodecIDDXTORY    = CodecID(C.AV_CODEC_ID_DXTORY)
	CodecIDV410      = CodecID(C.AV_CODEC_ID_V410)
	CodecIDXWD       = CodecID(C.AV_CODEC_ID_XWD)
	CodecIDCDXL      = CodecID(C.AV_CODEC_ID_CDXL)
	CodecIDXBM       = CodecID(C.AV_CODEC_ID_XBM)
	CodecIDZEROCODEC = CodecID(C.AV_CODEC_ID_ZEROCODEC)
	CodecIDMSS1      = CodecID(C.AV_CODEC_ID_MSS1)
	CodecIDMSA1      = CodecID(C.AV_CODEC_ID_MSA1)
	CodecIDTSCC2     = CodecID(C.AV_CODEC_ID_TSCC2)
	CodecIDMTS2      = CodecID(C.AV_CODEC_ID_MTS2)
	CodecIDCLLC      = CodecID(C.AV_CODEC_ID_CLLC)
	CodecIDMSS2      = CodecID(C.AV_CODEC_ID_MSS2)
	CodecIDVP9       = CodecID(C.AV_CODEC_ID_VP9)
	CodecIDAIC       = CodecID(C.AV_CODEC_ID_AIC)
	CodecIDESCAPE130 = CodecID(C.AV_CODEC_ID_ESCAPE130)
	CodecIDG2M       = CodecID(C.AV_CODEC_ID_G2M)
	CodecIDWEBP      = CodecID(C.AV_CODEC_ID_WEBP)
	CodecIDHNM4VIDEO = CodecID(C.AV_CODEC_ID_HNM4_VIDEO)
	CodecIDHEVC      = CodecID(C.AV_CODEC_ID_HEVC)

	CodecIDFIC          = CodecID(C.AV_CODEC_ID_FIC)
	CodecIDALIAS_PIX    = CodecID(C.AV_CODEC_ID_ALIAS_PIX)
	CodecIDBRENDER_PIX  = CodecID(C.AV_CODEC_ID_BRENDER_PIX)
	CodecIDPAF_VIDEO    = CodecID(C.AV_CODEC_ID_PAF_VIDEO)
	CodecIDEXR          = CodecID(C.AV_CODEC_ID_EXR)
	CodecIDVP7          = CodecID(C.AV_CODEC_ID_VP7)
	CodecIDSANM         = CodecID(C.AV_CODEC_ID_SANM)
	CodecIDSGIRLE       = CodecID(C.AV_CODEC_ID_SGIRLE)
	CodecIDMVC1         = CodecID(C.AV_CODEC_ID_MVC1)
	CodecIDMVC2         = CodecID(C.AV_CODEC_ID_MVC2)
	CodecIDHQX          = CodecID(C.AV_CODEC_ID_HQX)
	CodecIDTDSC         = CodecID(C.AV_CODEC_ID_TDSC)
	CodecIDHQ_HQA       = CodecID(C.AV_CODEC_ID_HQ_HQA)
	CodecIDHAP          = CodecID(C.AV_CODEC_ID_HAP)
	CodecIDDDS          = CodecID(C.AV_CODEC_ID_DDS)
	CodecIDDXV          = CodecID(C.AV_CODEC_ID_DXV)
	CodecIDSCREENPRESSO = CodecID(C.AV_CODEC_ID_SCREENPRESSO)
	CodecIDRSCC         = CodecID(C.AV_CODEC_ID_RSCC)

	CodecIDY41P          = CodecID(C.AV_CODEC_ID_Y41P)
	CodecIDAVRP          = CodecID(C.AV_CODEC_ID_AVRP)
	CodecID012V          = CodecID(C.AV_CODEC_ID_012V)
	CodecIDAVUI          = CodecID(C.AV_CODEC_ID_AVUI)
	CodecIDAYUV          = CodecID(C.AV_CODEC_ID_AYUV)
	CodecIDTARGA_Y216    = CodecID(C.AV_CODEC_ID_TARGA_Y216)
	CodecIDV308          = CodecID(C.AV_CODEC_ID_V308)
	CodecIDV408          = CodecID(C.AV_CODEC_ID_V408)
	CodecIDYUV4          = CodecID(C.AV_CODEC_ID_YUV4)
	CodecIDAVRN          = CodecID(C.AV_CODEC_ID_AVRN)
	CodecIDCPIA          = CodecID(C.AV_CODEC_ID_CPIA)
	CodecIDXFACE         = CodecID(C.AV_CODEC_ID_XFACE)
	CodecIDSNOW          = CodecID(C.AV_CODEC_ID_SNOW)
	CodecIDSMVJPEG       = CodecID(C.AV_CODEC_ID_SMVJPEG)
	CodecIDAPNG          = CodecID(C.AV_CODEC_ID_APNG)
	CodecIDDAALA         = CodecID(C.AV_CODEC_ID_DAALA)
	CodecIDCFHD          = CodecID(C.AV_CODEC_ID_CFHD)
	CodecIDTRUEMOTION2RT = CodecID(C.AV_CODEC_ID_TRUEMOTION2RT)
	CodecIDM101          = CodecID(C.AV_CODEC_ID_M101)
	CodecIDMAGICYUV      = CodecID(C.AV_CODEC_ID_MAGICYUV)
	CodecIDSHEERVIDEO    = CodecID(C.AV_CODEC_ID_SHEERVIDEO)
	CodecIDYLC           = CodecID(C.AV_CODEC_ID_YLC)
	CodecIDPSD           = CodecID(C.AV_CODEC_ID_PSD)
	CodecIDPIXLET        = CodecID(C.AV_CODEC_ID_PIXLET)
	CodecIDSPEEDHQ       = CodecID(C.AV_CODEC_ID_SPEEDHQ)
	CodecIDFMVC          = CodecID(C.AV_CODEC_ID_FMVC)
	CodecIDSCPR          = CodecID(C.AV_CODEC_ID_SCPR)
	CodecIDCLEARVIDEO    = CodecID(C.AV_CODEC_ID_CLEARVIDEO)
	CodecIDXPM           = CodecID(C.AV_CODEC_ID_XPM)
	CodecIDAV1           = CodecID(C.AV_CODEC_ID_AV1)
	CodecIDBITPACKED     = CodecID(C.AV_CODEC_ID_BITPACKED)
	CodecIDMSCC          = CodecID(C.AV_CODEC_ID_MSCC)
	CodecIDSRGC          = CodecID(C.AV_CODEC_ID_SRGC)
	CodecIDSVG           = CodecID(C.AV_CODEC_ID_SVG)
	CodecIDGDV           = CodecID(C.AV_CODEC_ID_GDV)
	CodecIDFITS          = CodecID(C.AV_CODEC_ID_FITS)

	CodecIDFIRST_AUDIO      = CodecID(C.AV_CODEC_ID_FIRST_AUDIO)
	CodecIDPCM_S16LE        = CodecID(C.AV_CODEC_ID_PCM_S16LE)
	CodecIDPCM_S16BE        = CodecID(C.AV_CODEC_ID_PCM_S16BE)
	CodecIDPCM_U16LE        = CodecID(C.AV_CODEC_ID_PCM_U16LE)
	CodecIDPCM_U16BE        = CodecID(C.AV_CODEC_ID_PCM_U16BE)
	CodecIDPCM_S8           = CodecID(C.AV_CODEC_ID_PCM_S8)
	CodecIDPCM_U8           = CodecID(C.AV_CODEC_ID_PCM_U8)
	CodecIDPCM_MULAW        = CodecID(C.AV_CODEC_ID_PCM_MULAW)
	CodecIDPCM_ALAW         = CodecID(C.AV_CODEC_ID_PCM_ALAW)
	CodecIDPCM_S32LE        = CodecID(C.AV_CODEC_ID_PCM_S32LE)
	CodecIDPCM_S32BE        = CodecID(C.AV_CODEC_ID_PCM_S32BE)
	CodecIDPCM_U32LE        = CodecID(C.AV_CODEC_ID_PCM_U32LE)
	CodecIDPCM_U32BE        = CodecID(C.AV_CODEC_ID_PCM_U32BE)
	CodecIDPCM_S24LE        = CodecID(C.AV_CODEC_ID_PCM_S24LE)
	CodecIDPCM_S24BE        = CodecID(C.AV_CODEC_ID_PCM_S24BE)
	CodecIDPCM_U24LE        = CodecID(C.AV_CODEC_ID_PCM_U24LE)
	CodecIDPCM_U24BE        = CodecID(C.AV_CODEC_ID_PCM_U24BE)
	CodecIDPCM_S24DAUD      = CodecID(C.AV_CODEC_ID_PCM_S24DAUD)
	CodecIDPCM_ZORK         = CodecID(C.AV_CODEC_ID_PCM_ZORK)
	CodecIDPCM_S16LE_PLANAR = CodecID(C.AV_CODEC_ID_PCM_S16LE_PLANAR)
	CodecIDPCM_DVD          = CodecID(C.AV_CODEC_ID_PCM_DVD)
	CodecIDPCM_F32BE        = CodecID(C.AV_CODEC_ID_PCM_F32BE)
	CodecIDPCM_F32LE        = CodecID(C.AV_CODEC_ID_PCM_F32LE)
	CodecIDPCM_F64BE        = CodecID(C.AV_CODEC_ID_PCM_F64BE)
	CodecIDPCM_F64LE        = CodecID(C.AV_CODEC_ID_PCM_F64LE)
	CodecIDPCM_BLURAY       = CodecID(C.AV_CODEC_ID_PCM_BLURAY)
	CodecIDPCM_LXF          = CodecID(C.AV_CODEC_ID_PCM_LXF)
	CodecIDS302M            = CodecID(C.AV_CODEC_ID_S302M)
	CodecIDPCM_S8_PLANAR    = CodecID(C.AV_CODEC_ID_PCM_S8_PLANAR)
	CodecIDPCM_S24LE_PLANAR = CodecID(C.AV_CODEC_ID_PCM_S24LE_PLANAR)
	CodecIDPCM_S32LE_PLANAR = CodecID(C.AV_CODEC_ID_PCM_S32LE_PLANAR)
	CodecIDPCM_S16BE_PLANAR = CodecID(C.AV_CODEC_ID_PCM_S16BE_PLANAR)

	CodecIDPCM_S64LE = CodecID(C.AV_CODEC_ID_PCM_S64LE)
	CodecIDPCM_S64BE = CodecID(C.AV_CODEC_ID_PCM_S64BE)
	CodecIDPCM_F16LE = CodecID(C.AV_CODEC_ID_PCM_F16LE)
	CodecIDPCM_F24LE = CodecID(C.AV_CODEC_ID_PCM_F24LE)

	CodecIDADPCM_IMA_QT      = CodecID(C.AV_CODEC_ID_ADPCM_IMA_QT)
	CodecIDADPCM_IMA_WAV     = CodecID(C.AV_CODEC_ID_ADPCM_IMA_WAV)
	CodecIDADPCM_IMA_DK3     = CodecID(C.AV_CODEC_ID_ADPCM_IMA_DK3)
	CodecIDADPCM_IMA_DK4     = CodecID(C.AV_CODEC_ID_ADPCM_IMA_DK4)
	CodecIDADPCM_IMA_WS      = CodecID(C.AV_CODEC_ID_ADPCM_IMA_WS)
	CodecIDADPCM_IMA_SMJPEG  = CodecID(C.AV_CODEC_ID_ADPCM_IMA_SMJPEG)
	CodecIDADPCM_MS          = CodecID(C.AV_CODEC_ID_ADPCM_MS)
	CodecIDADPCM_4XM         = CodecID(C.AV_CODEC_ID_ADPCM_4XM)
	CodecIDADPCM_XA          = CodecID(C.AV_CODEC_ID_ADPCM_XA)
	CodecIDADPCM_ADX         = CodecID(C.AV_CODEC_ID_ADPCM_ADX)
	CodecIDADPCM_EA          = CodecID(C.AV_CODEC_ID_ADPCM_EA)
	CodecIDADPCM_G726        = CodecID(C.AV_CODEC_ID_ADPCM_G726)
	CodecIDADPCM_CT          = CodecID(C.AV_CODEC_ID_ADPCM_CT)
	CodecIDADPCM_SWF         = CodecID(C.AV_CODEC_ID_ADPCM_SWF)
	CodecIDADPCM_YAMAHA      = CodecID(C.AV_CODEC_ID_ADPCM_YAMAHA)
	CodecIDADPCM_SBPRO_4     = CodecID(C.AV_CODEC_ID_ADPCM_SBPRO_4)
	CodecIDADPCM_SBPRO_3     = CodecID(C.AV_CODEC_ID_ADPCM_SBPRO_3)
	CodecIDADPCM_SBPRO_2     = CodecID(C.AV_CODEC_ID_ADPCM_SBPRO_2)
	CodecIDADPCM_THP         = CodecID(C.AV_CODEC_ID_ADPCM_THP)
	CodecIDADPCM_IMA_AMV     = CodecID(C.AV_CODEC_ID_ADPCM_IMA_AMV)
	CodecIDADPCM_EA_R1       = CodecID(C.AV_CODEC_ID_ADPCM_EA_R1)
	CodecIDADPCM_EA_R3       = CodecID(C.AV_CODEC_ID_ADPCM_EA_R3)
	CodecIDADPCM_EA_R2       = CodecID(C.AV_CODEC_ID_ADPCM_EA_R2)
	CodecIDADPCM_IMA_EA_SEAD = CodecID(C.AV_CODEC_ID_ADPCM_IMA_EA_SEAD)
	CodecIDADPCM_IMA_EA_EACS = CodecID(C.AV_CODEC_ID_ADPCM_IMA_EA_EACS)
	CodecIDADPCM_EA_XAS      = CodecID(C.AV_CODEC_ID_ADPCM_EA_XAS)
	CodecIDADPCM_EA_MAXIS_XA = CodecID(C.AV_CODEC_ID_ADPCM_EA_MAXIS_XA)
	CodecIDADPCM_IMA_ISS     = CodecID(C.AV_CODEC_ID_ADPCM_IMA_ISS)
	CodecIDADPCM_G722        = CodecID(C.AV_CODEC_ID_ADPCM_G722)
	CodecIDADPCM_IMA_APC     = CodecID(C.AV_CODEC_ID_ADPCM_IMA_APC)
	CodecIDADPCM_VIMA        = CodecID(C.AV_CODEC_ID_ADPCM_VIMA)

	CodecIDADPCM_AFC      = CodecID(C.AV_CODEC_ID_ADPCM_AFC)
	CodecIDADPCM_IMA_OKI  = CodecID(C.AV_CODEC_ID_ADPCM_IMA_OKI)
	CodecIDADPCM_DTK      = CodecID(C.AV_CODEC_ID_ADPCM_DTK)
	CodecIDADPCM_IMA_RAD  = CodecID(C.AV_CODEC_ID_ADPCM_IMA_RAD)
	CodecIDADPCM_G726LE   = CodecID(C.AV_CODEC_ID_ADPCM_G726LE)
	CodecIDADPCM_THP_LE   = CodecID(C.AV_CODEC_ID_ADPCM_THP_LE)
	CodecIDADPCM_PSX      = CodecID(C.AV_CODEC_ID_ADPCM_PSX)
	CodecIDADPCM_AICA     = CodecID(C.AV_CODEC_ID_ADPCM_AICA)
	CodecIDADPCM_IMA_DAT4 = CodecID(C.AV_CODEC_ID_ADPCM_IMA_DAT4)
	CodecIDADPCM_MTAF     = CodecID(C.AV_CODEC_ID_ADPCM_MTAF)

	CodecIDAMR_NB = CodecID(C.AV_CODEC_ID_AMR_NB)
	CodecIDAMR_WB = CodecID(C.AV_CODEC_ID_AMR_WB)

	CodecIDRA_144 = CodecID(C.AV_CODEC_ID_RA_144)
	CodecIDRA_288 = CodecID(C.AV_CODEC_ID_RA_288)

	CodecIDROQ_DPCM       = CodecID(C.AV_CODEC_ID_ROQ_DPCM)
	CodecIDINTERPLAY_DPCM = CodecID(C.AV_CODEC_ID_INTERPLAY_DPCM)
	CodecIDXAN_DPCM       = CodecID(C.AV_CODEC_ID_XAN_DPCM)
	CodecIDSOL_DPCM       = CodecID(C.AV_CODEC_ID_SOL_DPCM)

	CodecIDSDX2_DPCM    = CodecID(C.AV_CODEC_ID_SDX2_DPCM)
	CodecIDGREMLIN_DPCM = CodecID(C.AV_CODEC_ID_GREMLIN_DPCM)

	CodecIDMP2            = CodecID(C.AV_CODEC_ID_MP2)
	CodecIDMP3            = CodecID(C.AV_CODEC_ID_MP3)
	CodecIDAAC            = CodecID(C.AV_CODEC_ID_AAC)
	CodecIDAC3            = CodecID(C.AV_CODEC_ID_AC3)
	CodecIDDTS            = CodecID(C.AV_CODEC_ID_DTS)
	CodecIDVORBIS         = CodecID(C.AV_CODEC_ID_VORBIS)
	CodecIDDVAUDIO        = CodecID(C.AV_CODEC_ID_DVAUDIO)
	CodecIDWMAV1          = CodecID(C.AV_CODEC_ID_WMAV1)
	CodecIDWMAV2          = CodecID(C.AV_CODEC_ID_WMAV2)
	CodecIDMACE3          = CodecID(C.AV_CODEC_ID_MACE3)
	CodecIDMACE6          = CodecID(C.AV_CODEC_ID_MACE6)
	CodecIDVMDAUDIO       = CodecID(C.AV_CODEC_ID_VMDAUDIO)
	CodecIDFLAC           = CodecID(C.AV_CODEC_ID_FLAC)
	CodecIDMP3ADU         = CodecID(C.AV_CODEC_ID_MP3ADU)
	CodecIDMP3ON4         = CodecID(C.AV_CODEC_ID_MP3ON4)
	CodecIDSHORTEN        = CodecID(C.AV_CODEC_ID_SHORTEN)
	CodecIDALAC           = CodecID(C.AV_CODEC_ID_ALAC)
	CodecIDWESTWOOD_SND1  = CodecID(C.AV_CODEC_ID_WESTWOOD_SND1)
	CodecIDGSM            = CodecID(C.AV_CODEC_ID_GSM)
	CodecIDQDM2           = CodecID(C.AV_CODEC_ID_QDM2)
	CodecIDCOOK           = CodecID(C.AV_CODEC_ID_COOK)
	CodecIDTRUESPEECH     = CodecID(C.AV_CODEC_ID_TRUESPEECH)
	CodecIDTTA            = CodecID(C.AV_CODEC_ID_TTA)
	CodecIDSMACKAUDIO     = CodecID(C.AV_CODEC_ID_SMACKAUDIO)
	CodecIDQCELP          = CodecID(C.AV_CODEC_ID_QCELP)
	CodecIDWAVPACK        = CodecID(C.AV_CODEC_ID_WAVPACK)
	CodecIDDSICINAUDIO    = CodecID(C.AV_CODEC_ID_DSICINAUDIO)
	CodecIDIMC            = CodecID(C.AV_CODEC_ID_IMC)
	CodecIDMUSEPACK7      = CodecID(C.AV_CODEC_ID_MUSEPACK7)
	CodecIDMLP            = CodecID(C.AV_CODEC_ID_MLP)
	CodecIDGSM_MS         = CodecID(C.AV_CODEC_ID_GSM_MS)
	CodecIDATRAC3         = CodecID(C.AV_CODEC_ID_ATRAC3)
	CodecIDAPE            = CodecID(C.AV_CODEC_ID_APE)
	CodecIDNELLYMOSER     = CodecID(C.AV_CODEC_ID_NELLYMOSER)
	CodecIDMUSEPACK8      = CodecID(C.AV_CODEC_ID_MUSEPACK8)
	CodecIDSPEEX          = CodecID(C.AV_CODEC_ID_SPEEX)
	CodecIDWMAVOICE       = CodecID(C.AV_CODEC_ID_WMAVOICE)
	CodecIDWMAPRO         = CodecID(C.AV_CODEC_ID_WMAPRO)
	CodecIDWMALOSSLESS    = CodecID(C.AV_CODEC_ID_WMALOSSLESS)
	CodecIDATRAC3P        = CodecID(C.AV_CODEC_ID_ATRAC3P)
	CodecIDEAC3           = CodecID(C.AV_CODEC_ID_EAC3)
	CodecIDSIPR           = CodecID(C.AV_CODEC_ID_SIPR)
	CodecIDMP1            = CodecID(C.AV_CODEC_ID_MP1)
	CodecIDTWINVQ         = CodecID(C.AV_CODEC_ID_TWINVQ)
	CodecIDTRUEHD         = CodecID(C.AV_CODEC_ID_TRUEHD)
	CodecIDMP4ALS         = CodecID(C.AV_CODEC_ID_MP4ALS)
	CodecIDATRAC1         = CodecID(C.AV_CODEC_ID_ATRAC1)
	CodecIDBINKAUDIO_RDFT = CodecID(C.AV_CODEC_ID_BINKAUDIO_RDFT)
	CodecIDBINKAUDIO_DCT  = CodecID(C.AV_CODEC_ID_BINKAUDIO_DCT)
	CodecIDAAC_LATM       = CodecID(C.AV_CODEC_ID_AAC_LATM)
	CodecIDQDMC           = CodecID(C.AV_CODEC_ID_QDMC)
	CodecIDCELT           = CodecID(C.AV_CODEC_ID_CELT)
	CodecIDG723_1         = CodecID(C.AV_CODEC_ID_G723_1)
	CodecIDG729           = CodecID(C.AV_CODEC_ID_G729)
	CodecID8SVX_EXP       = CodecID(C.AV_CODEC_ID_8SVX_EXP)
	CodecID8SVX_FIB       = CodecID(C.AV_CODEC_ID_8SVX_FIB)
	CodecIDBMV_AUDIO      = CodecID(C.AV_CODEC_ID_BMV_AUDIO)
	CodecIDRALF           = CodecID(C.AV_CODEC_ID_RALF)
	CodecIDIAC            = CodecID(C.AV_CODEC_ID_IAC)
	CodecIDILBC           = CodecID(C.AV_CODEC_ID_ILBC)
	CodecIDOPUS           = CodecID(C.AV_CODEC_ID_OPUS)
	CodecIDCOMFORT_NOISE  = CodecID(C.AV_CODEC_ID_COMFORT_NOISE)
	CodecIDTAK            = CodecID(C.AV_CODEC_ID_TAK)
	CodecIDMETASOUND      = CodecID(C.AV_CODEC_ID_METASOUND)
	CodecIDPAF_AUDIO      = CodecID(C.AV_CODEC_ID_PAF_AUDIO)
	CodecIDON2AVC         = CodecID(C.AV_CODEC_ID_ON2AVC)
	CodecIDDSS_SP         = CodecID(C.AV_CODEC_ID_DSS_SP)
	CodecIDCODEC2         = CodecID(C.AV_CODEC_ID_CODEC2)

	CodecIDFFWAVESYNTH        = CodecID(C.AV_CODEC_ID_FFWAVESYNTH)
	CodecIDSONIC              = CodecID(C.AV_CODEC_ID_SONIC)
	CodecIDSONIC_LS           = CodecID(C.AV_CODEC_ID_SONIC_LS)
	CodecIDEVRC               = CodecID(C.AV_CODEC_ID_EVRC)
	CodecIDSMV                = CodecID(C.AV_CODEC_ID_SMV)
	CodecIDDSD_LSBF           = CodecID(C.AV_CODEC_ID_DSD_LSBF)
	CodecIDDSD_MSBF           = CodecID(C.AV_CODEC_ID_DSD_MSBF)
	CodecIDDSD_LSBF_PLANAR    = CodecID(C.AV_CODEC_ID_DSD_LSBF_PLANAR)
	CodecIDDSD_MSBF_PLANAR    = CodecID(C.AV_CODEC_ID_DSD_MSBF_PLANAR)
	CodecID4GV                = CodecID(C.AV_CODEC_ID_4GV)
	CodecIDINTERPLAY_ACM      = CodecID(C.AV_CODEC_ID_INTERPLAY_ACM)
	CodecIDXMA1               = CodecID(C.AV_CODEC_ID_XMA1)
	CodecIDXMA2               = CodecID(C.AV_CODEC_ID_XMA2)
	CodecIDDST                = CodecID(C.AV_CODEC_ID_DST)
	CodecIDATRAC3AL           = CodecID(C.AV_CODEC_ID_ATRAC3AL)
	CodecIDATRAC3PAL          = CodecID(C.AV_CODEC_ID_ATRAC3PAL)
	CodecIDDOLBY_E            = CodecID(C.AV_CODEC_ID_DOLBY_E)
	CodecIDAPTX               = CodecID(C.AV_CODEC_ID_APTX)
	CodecIDAPTX_HD            = CodecID(C.AV_CODEC_ID_APTX_HD)
	CodecIDSBC                = CodecID(C.AV_CODEC_ID_SBC)
	CodecIDFIRST_SUBTITLE     = CodecID(C.AV_CODEC_ID_FIRST_SUBTITLE)
	CodecIDDVD_SUBTITLE       = CodecID(C.AV_CODEC_ID_DVD_SUBTITLE)
	CodecIDDVB_SUBTITLE       = CodecID(C.AV_CODEC_ID_DVB_SUBTITLE)
	CodecIDTEXT               = CodecID(C.AV_CODEC_ID_TEXT)
	CodecIDXSUB               = CodecID(C.AV_CODEC_ID_XSUB)
	CodecIDSSA                = CodecID(C.AV_CODEC_ID_SSA)
	CodecIDMOV_TEXT           = CodecID(C.AV_CODEC_ID_MOV_TEXT)
	CodecIDHDMV_PGS_SUBTITLE  = CodecID(C.AV_CODEC_ID_HDMV_PGS_SUBTITLE)
	CodecIDDVB_TELETEXT       = CodecID(C.AV_CODEC_ID_DVB_TELETEXT)
	CodecIDSRT                = CodecID(C.AV_CODEC_ID_SRT)
	CodecIDMICRODVD           = CodecID(C.AV_CODEC_ID_MICRODVD)
	CodecIDEIA_608            = CodecID(C.AV_CODEC_ID_EIA_608)
	CodecIDJACOSUB            = CodecID(C.AV_CODEC_ID_JACOSUB)
	CodecIDSAMI               = CodecID(C.AV_CODEC_ID_SAMI)
	CodecIDREALTEXT           = CodecID(C.AV_CODEC_ID_REALTEXT)
	CodecIDSTL                = CodecID(C.AV_CODEC_ID_STL)
	CodecIDSUBVIEWER1         = CodecID(C.AV_CODEC_ID_SUBVIEWER1)
	CodecIDSUBVIEWER          = CodecID(C.AV_CODEC_ID_SUBVIEWER)
	CodecIDSUBRIP             = CodecID(C.AV_CODEC_ID_SUBRIP)
	CodecIDWEBVTT             = CodecID(C.AV_CODEC_ID_WEBVTT)
	CodecIDMPL2               = CodecID(C.AV_CODEC_ID_MPL2)
	CodecIDVPLAYER            = CodecID(C.AV_CODEC_ID_VPLAYER)
	CodecIDPJS                = CodecID(C.AV_CODEC_ID_PJS)
	CodecIDASS                = CodecID(C.AV_CODEC_ID_ASS)
	CodecIDHDMV_TEXT_SUBTITLE = CodecID(C.AV_CODEC_ID_HDMV_TEXT_SUBTITLE)

	CodecIDFIRST_UNKNOWN   = CodecID(C.AV_CODEC_ID_FIRST_UNKNOWN)
	CodecIDTTF             = CodecID(C.AV_CODEC_ID_TTF)
	CodecIDSCTE_35         = CodecID(C.AV_CODEC_ID_SCTE_35)
	CodecIDBINTEXT         = CodecID(C.AV_CODEC_ID_BINTEXT)
	CodecIDXBIN            = CodecID(C.AV_CODEC_ID_XBIN)
	CodecIDIDF             = CodecID(C.AV_CODEC_ID_IDF)
	CodecIDOTF             = CodecID(C.AV_CODEC_ID_OTF)
	CodecIDSMPTE_KLV       = CodecID(C.AV_CODEC_ID_SMPTE_KLV)
	CodecIDDVD_NAV         = CodecID(C.AV_CODEC_ID_DVD_NAV)
	CodecIDTIMED_ID3       = CodecID(C.AV_CODEC_ID_TIMED_ID3)
	CodecIDBIN_DATA        = CodecID(C.AV_CODEC_ID_BIN_DATA)
	CodecIDPROBE           = CodecID(C.AV_CODEC_ID_PROBE)
	CodecIDMPEG2TS         = CodecID(C.AV_CODEC_ID_MPEG2TS)
	CodecIDMPEG4SYSTEMS    = CodecID(C.AV_CODEC_ID_MPEG4SYSTEMS)
	CodecIDFFMETADATA      = CodecID(C.AV_CODEC_ID_FFMETADATA)
	CodecIDWRAPPED_AVFRAME = CodecID(C.AV_CODEC_ID_WRAPPED_AVFRAME)
)
