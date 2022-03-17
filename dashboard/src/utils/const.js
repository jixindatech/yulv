const ROLE_OPTIONS = [
  { value: 'admin', lable: 'admin' }
]

const LOGIN_OPTIONS = [
  { value: 'standard', label: 'standard' },
  { value: 'ldap', label: 'ldap' }
]

const LDAP_TYPE_OPTIONS = [
  { value: 'tcp', label: 'tcp' },
  { value: 'udp', label: 'udp' }
]

const USER_STATUS_OPTIONS = [
  { value: '未锁定', label: 1 },
  { value: '已锁定', label: 2 }
]

const RULE_TYPE_OPTIONS = [
  { label: '允许类型', value: 1 },
  { label: '拒绝类型', value: 2 },
  { label: '日志', value: 3 }
]

const RULE_TYPE_MAP = {
  1: '允许类型',
  2: '拒绝类型',
  3: '日志'
}

const IP_TYPE_OPTIONS = [
  { label: '允许类型', value: 1 },
  { label: '拒绝类型', value: 2 }
]

const IP_TYPE_MAP = {
  1: '允许类型',
  2: '拒绝类型'
}

const RULE_MATCH_OPTIONS = [
  { label: '字符串查找', value: 'str_find' },
  { label: '正则匹配', value: 're' }
]

const SQL_TYPE_OPTIONS = [
  { label: 'INSERT', value: 'insert' },
  { label: 'UPDATE', value: 'update' },
  { label: 'DELETE', value: 'delete' },
  { label: 'REPLACE', value: 'replace' },
  { label: 'SET', value: 'set' },
  { label: 'BEGIN', value: 'begin' },
  { label: 'COMMIT', value: 'commit' },
  { label: 'ROLLBACK', value: 'rollback' },
  { label: 'ADMIN', value: 'admin' },
  { label: 'SELECT', value: 'select' },
  { label: 'USE', value: 'use' },
  { label: 'START', value: 'start' },
  { label: 'TRANSACTION', value: 'transaction' },
  { label: 'SHOW', value: 'show' },
  { label: 'TRUNCATE', value: 'truncate' }
]

const SQL_CHARSET = [
  { label: 'big5', value: 'big5' },
  { label: 'dec8', value: 'dec8' },
  { label: 'cp850', value: 'cp850' },
  { label: 'hp8', value: 'hp8' },
  { label: 'koi8r', value: 'koi8r' },
  { label: 'latin1', value: 'latin1' },
  { label: 'latin2', value: 'latin2' },
  { label: 'swe7', value: 'swe7' },
  { label: 'ascii', value: 'ascii' },
  { label: 'ujis', value: 'ujis' },
  { label: 'sjis', value: 'sjis' },
  { label: 'hebrew', value: 'hebrew' },
  { label: 'tis620', value: 'tis620' },
  { label: 'euckr', value: 'euckr' },
  { label: 'koi8u', value: 'koi8u' },
  { label: 'gb2312', value: 'gb2312' },
  { label: 'greek', value: 'greek' },
  { label: 'cp1250', value: 'cp1250' },
  { label: 'gbk', value: 'gbk' },
  { label: 'latin5', value: 'latin5' },
  { label: 'armscii8', value: 'armscii8' },
  { label: 'utf8', value: 'utf8' },
  { label: 'ucs2', value: 'ucs2' },
  { label: 'cp866', value: 'cp866' },
  { label: 'keybcs2', value: 'keybcs2' },
  { label: 'macce', value: 'macce' },
  { label: 'macroman', value: 'macroman' },
  { label: 'cp852', value: 'cp852' },
  { label: 'latin7', value: 'latin7' },
  { label: 'utf8mb4', value: 'utf8mb4' },
  { label: 'cp1251', value: 'cp1251' },
  { label: 'utf16', value: 'utf16' },
  { label: 'utf16le', value: 'utf16le' },
  { label: 'cp1256', value: 'cp1256' },
  { label: 'cp1257', value: 'cp1257' },
  { label: 'utf32', value: 'utf32' },
  { label: 'binary', value: 'binary' },
  { label: 'geostd8', value: 'geostd8' },
  { label: 'cp932', value: 'cp932' },
  { label: 'eucjpms', value: 'eucjpms' }
]

const SQL_COLLATION_BIG5 = [
  { label: 'big5_chinese_ci', value: 'big5_chinese_ci' },
  { label: 'big5_bin', value: 'big5_bin' }]

const SQL_COLLATION_DEC8 = [
  { label: 'dec8_swedish_ci', value: 'dec8_swedish_ci' },
  { label: 'dec8_bin', value: 'dec8_bin' }
]
const SQL_COLLATION_CP850 = [
  { label: 'cp850_general_ci', value: 'cp850_general_ci' },
  { label: 'cp850_bin', value: 'cp850_bin' }
]
const SQL_COLLATION_HP8 = [
  { label: 'hp8_english_ci', value: 'hp8_english_ci' },
  { label: 'hp8_bin', value: 'hp8_bin' }
]

const SQL_COLLATION_KOI8R = [
  { label: 'koi8r_general_ci', value: 'koi8r_general_ci' },
  { label: 'koi8r_bin', value: 'koi8r_bin' }
]
const SQL_COLLATION_LATIN1 = [
  { label: 'latin1_german1_ci', value: 'latin1_german1_ci' },
  { label: 'latin1_swedish_ci', value: 'latin1_swedish_ci' },
  { label: 'latin1_danish_ci', value: 'latin1_danish_ci' },
  { label: 'latin1_german2_ci', value: 'latin1_german2_ci' },
  { label: 'latin1_bin', value: 'latin1_bin' },
  { label: 'latin1_general_ci', value: 'latin1_general_ci' },
  { label: 'latin1_general_cs', value: 'latin1_general_cs' },
  { label: 'latin1_spanish_ci', value: 'latin1_spanish_ci' }
]
const SQL_COLLATION_LATIN2 = [
  { label: 'latin2_czech_cs', value: 'latin2_czech_cs' },
  { label: 'latin2_general_ci', value: 'latin2_general_ci' },
  { label: 'latin2_hungarian_ci', value: 'latin2_hungarian_ci' },
  { label: 'latin2_bin', value: 'latin2_bin' }
]
const SQL_COLLATION_SWE7 = [
  { label: 'swe7_swedish_ci', value: 'swe7_swedish_ci' },
  { label: 'swe7_bin', value: 'swe7_bin' }
]
const SQL_COLLATION_ASCII = [
  { label: 'ascii_general_ci', value: 'ascii_general_ci' },
  { label: 'ascii_bin', value: 'ascii_bin' }
]
const SQL_COLLATION_UJIS = [
  { label: 'ujis_japanese_ci', value: 'ujis_japanese_ci' },
  { label: 'ujis_bin', value: 'ujis_bin' }
]
const SQL_COLLATION_SJIS = [
  { label: 'sjis_japanese_ci', value: 'sjis_japanese_ci' },
  { label: 'sjis_bin', value: 'sjis_bin' }
]
const SQL_COLLATION_HEBREW = [
  { label: 'hebrew_general_ci', value: 'hebrew_general_ci' },
  { label: 'hebrew_bin', value: 'hebrew_bin' }
]
const SQL_COLLATION_TIS620 = [
  { label: 'tis620_thai_ci', value: 'tis620_thai_ci' },
  { label: 'tis620_bin', value: 'tis620_bin' }
]
const SQL_COLLATION_EUCKR = [
  { label: 'euckr_korean_ci', value: 'euckr_korean_ci' },
  { label: 'euckr_bin', value: 'euckr_bin' }
]
const SQL_COLLATION_KOI8U = [
  { label: 'koi8u_general_ci', value: 'koi8u_general_ci' },
  { label: 'koi8u_bin', value: 'koi8u_bin' }
]
const SQL_COLLATION_GB2312 = [
  { label: 'gb2312_chinese_ci', value: 'gb2312_chinese_ci' },
  { label: 'gb2312_bin', value: 'gb2312_bin' }
]
const SQL_COLLATION_GREEK = [
  { label: 'greek_general_ci', value: 'greek_general_ci' },
  { label: 'greek_bin', value: 'greek_bin' }
]
const SQL_COLLATION_CP1250 = [
  { label: 'cp1250_general_ci', value: 'cp1250_general_ci' },
  { label: 'cp1250_czech_cs', value: 'cp1250_czech_cs' },
  { label: 'cp1250_bin', value: 'cp1250_bin' },
  { label: 'cp1250_polish_ci', value: 'cp1250_polish_ci' }
]
const SQL_COLLATION_GBK = [
  { label: 'gbk_chinese_ci', value: 'gbk_chinese_ci' },
  { label: 'gbk_bin', value: 'gbk_bin' }
]
const SQL_COLLATION_LATIN5 = [
  { label: 'latin5_turkish_ci', value: 'latin5_turkish_ci' },
  { label: 'latin5_bin', value: 'latin5_bin' }
]
const SQL_COLLATION_ARMSCII8 = [
  { label: 'armscii8_general_ci', value: 'armscii8_general_ci' },
  { label: 'armscii8_bin', value: 'armscii8_bin' }
]
const SQL_COLLATION_UTF8 = [
  { label: 'utf8_general_ci', value: 'utf8_general_ci' },
  { label: 'utf8_bin', value: 'utf8_bin' },
  { label: 'utf8_unicode_ci', value: 'utf8_unicode_ci' },
  { label: 'utf8_latvian_ci', value: 'utf8_latvian_ci' },
  { label: 'utf8_romanian_ci', value: 'utf8_romanian_ci' },
  { label: 'utf8_slovenian_ci', value: 'utf8_slovenian_ci' },
  { label: 'utf8_estonian_ci', value: 'utf8_estonian_ci' },
  { label: 'utf8_spanish_ci', value: 'utf8_spanish_ci' },
  { label: 'utf8_swedish_ci', value: 'utf8_swedish_ci' },
  { label: 'utf8_turkish_ci', value: 'utf8_turkish_ci' },
  { label: 'utf8_czech_ci', value: 'utf8_czech_ci' },
  { label: 'utf8_danish_ci', value: 'utf8_danish_ci' },
  { label: 'utf8_lithuanian_ci', value: 'utf8_lithuanian_ci' },
  { label: 'utf8_slovak_ci', value: 'utf8_slovak_ci' },
  { label: 'utf8_spanish2_ci', value: 'utf8_spanish2_ci' },
  { label: 'utf8_roman_ci', value: 'utf8_roman_ci' },
  { label: 'utf8_persian_ci', value: 'utf8_persian_ci' },
  { label: 'utf8_esperanto_ci', value: 'utf8_esperanto_ci' },
  { label: 'utf8_hungarian_ci', value: 'utf8_hungarian_ci' },
  { label: 'utf8_sinhala_ci', value: 'utf8_sinhala_ci' },
  { label: 'utf8_german2_ci', value: 'utf8_german2_ci' },
  { label: 'utf8_croatian_ci', value: 'utf8_croatian_ci' },
  { label: 'utf8_unicode_520_ci', value: 'utf8_unicode_520_ci' },
  { label: 'utf8_vietnamese_ci', value: 'utf8_vietnamese_ci' },
  { label: 'utf8_general_mysql500_ci', value: 'utf8_general_mysql500_ci' }
]
const SQL_COLLATION_UCS2 = [
  { label: 'ucs2_general_ci', value: 'ucs2_general_ci' },
  { label: 'ucs2_bin', value: 'ucs2_bin' },
  { label: 'ucs2_icelandic_ci', value: 'ucs2_icelandic_ci' },
  { label: 'ucs2_general_ci', value: 'ucs2_general_ci' },
  { label: 'ucs2_bin', value: 'ucs2_bin' },
  { label: 'ucs2_icelandic_ci', value: 'ucs2_icelandic_ci' },
  { label: 'ucs2_latvian_ci', value: 'ucs2_latvian_ci' },
  { label: 'ucs2_romanian_ci', value: 'ucs2_romanian_ci' },
  { label: 'ucs2_slovenian_ci', value: 'ucs2_slovenian_ci' },
  { label: 'ucs2_polish_ci', value: 'ucs2_polish_ci' },
  { label: 'ucs2_estonian_ci', value: 'ucs2_estonian_ci' },
  { label: 'ucs2_spanish_ci', value: 'ucs2_spanish_ci' },
  { label: 'ucs2_swedish_ci', value: 'ucs2_swedish_ci' },
  { label: 'ucs2_turkish_ci', value: 'ucs2_turkish_ci' },
  { label: 'ucs2_czech_ci', value: 'ucs2_czech_ci' },
  { label: 'ucs2_danish_ci', value: 'ucs2_danish_ci' },
  { label: 'ucs2_lithuanian_ci', value: 'ucs2_lithuanian_ci' },
  { label: 'ucs2_slovak_ci', value: 'ucs2_slovak_ci' },
  { label: 'ucs2_spanish2_ci', value: 'ucs2_spanish2_ci' },
  { label: 'ucs2_roman_ci', value: 'ucs2_roman_ci' },
  { label: 'ucs2_persian_ci', value: 'ucs2_persian_ci' },
  { label: 'ucs2_esperanto_ci', value: 'ucs2_esperanto_ci' },
  { label: 'ucs2_hungarian_ci', value: 'ucs2_hungarian_ci' },
  { label: 'ucs2_sinhala_ci', value: 'ucs2_sinhala_ci' },
  { label: 'ucs2_german2_ci', value: 'ucs2_german2_ci' },
  { label: 'ucs2_croatian_ci', value: 'ucs2_croatian_ci' },
  { label: 'ucs2_unicode_520_ci', value: 'ucs2_unicode_520_ci' },
  { label: 'ucs2_vietnamese_ci', value: 'ucs2_vietnamese_ci' },
  { label: 'ucs2_general_mysql500_ci', value: 'ucs2_general_mysql500_ci' }
]
const SQL_COLLATION_CP886 = [
  { label: 'cp866_general_ci', value: 'cp866_general_ci' },
  { label: 'cp866_bin', value: 'cp866_bin' }
]
const SQL_COLLATION_KEYBCS2 = [
  { label: 'keybcs2_general_ci', value: 'keybcs2_general_ci' },
  { label: 'keybcs2_bin', value: 'keybcs2_bin' }
]
const SQL_COLLATION_MACCE = [
  { label: 'macce_general_ci', value: 'macce_general_ci' },
  { label: 'macce_bin', value: 'macce_bin' }
]

const SQL_COLLATION_MACROMAN = [
  { label: 'macroman_general_ci', value: 'macroman_general_ci' },
  { label: 'macroman_bin', value: 'macroman_bin' }
]
const SQL_COLLATION_CP852 = [
  { label: 'cp852_general_ci', value: 'cp852_general_ci' },
  { label: 'cp852_bin', value: 'cp852_bin' }
]
const SQL_COLLATION_LATIN7 = [
  { label: 'latin7_estonian_cs', value: 'latin7_estonian_cs' },
  { label: 'latin7_general_ci', value: 'latin7_general_ci' },
  { label: 'latin7_general_cs', value: 'latin7_general_cs' },
  { label: 'latin7_bin', value: 'latin7_bin' }
]
const SQL_COLLATION_UTF8MB4 = [
  { label: 'utf8mb4_general_ci', value: 'utf8mb4_general_ci' },
  { label: 'utf8mb4_bin', value: 'utf8mb4_bin' },
  { label: 'utf8mb4_unicode_ci', value: 'utf8mb4_unicode_ci' },
  { label: 'utf8mb4_icelandic_ci', value: 'utf8mb4_icelandic_ci' },
  { label: 'utf8mb4_latvian_ci', value: 'utf8mb4_latvian_ci' },
  { label: 'utf8mb4_romanian_ci', value: 'utf8mb4_romanian_ci' },
  { label: 'utf8mb4_slovenian_ci', value: 'utf8mb4_slovenian_ci' },
  { label: 'utf8mb4_polish_ci', value: 'utf8mb4_polish_ci' },
  { label: 'utf8mb4_estonian_ci', value: 'utf8mb4_estonian_ci' },
  { label: 'utf8mb4_spanish_ci', value: 'utf8mb4_spanish_ci' },
  { label: 'utf8mb4_swedish_ci', value: 'utf8mb4_swedish_ci' },
  { label: 'utf8mb4_turkish_ci', value: 'utf8mb4_turkish_ci' },
  { label: 'utf8mb4_czech_ci', value: 'utf8mb4_czech_ci' },
  { label: 'utf8mb4_danish_ci', value: 'utf8mb4_danish_ci' },
  { label: 'utf8mb4_lithuanian_ci', value: 'utf8mb4_lithuanian_ci' },
  { label: 'utf8mb4_slovak_ci', value: 'utf8mb4_slovak_ci' },
  { label: 'utf8mb4_spanish2_ci', value: 'utf8mb4_spanish2_ci' },
  { label: 'utf8mb4_roman_ci', value: 'utf8mb4_roman_ci' },
  { label: 'utf8mb4_persian_ci', value: 'utf8mb4_persian_ci' },
  { label: 'utf8mb4_esperanto_ci', value: 'utf8mb4_esperanto_ci' },
  { label: 'utf8mb4_hungarian_ci', value: 'utf8mb4_hungarian_ci' },
  { label: 'utf8mb4_sinhala_ci', value: 'utf8mb4_sinhala_ci' },
  { label: 'utf8mb4_german2_ci', value: 'utf8mb4_german2_ci' },
  { label: 'utf8mb4_croatian_ci', value: 'utf8mb4_croatian_ci' },
  { label: 'utf8mb4_unicode_520_ci', value: 'utf8mb4_unicode_520_ci' },
  { label: 'utf8mb4_vietnamese_ci', value: 'utf8mb4_vietnamese_ci' }
]
const SQL_COLLATION_CP1251 = [
  { label: 'cp1251_bulgarian_ci', value: 'cp1251_bulgarian_ci' },
  { label: 'cp1251_ukrainian_ci', value: 'cp1251_ukrainian_ci' },
  { label: 'cp1251_general_ci', value: 'cp1251_general_ci' },
  { label: 'cp1251_general_cs', value: 'cp1251_general_cs' }
]
const SQL_COLLATION_UTF16 = [
  { label: 'utf16_general_ci', value: 'utf16_general_ci' },
  { label: 'utf16_bin', value: 'utf16_bin' },
  { label: 'utf16_unicode_ci', value: 'utf16_unicode_ci' },
  { label: 'utf16_icelandic_ci', value: 'utf16_icelandic_ci' },
  { label: 'utf16_latvian_ci', value: 'utf16_latvian_ci' },
  { label: 'utf16_romanian_ci', value: 'utf16_romanian_ci' },
  { label: 'utf16_slovenian_ci', value: 'utf16_slovenian_ci' },
  { label: 'utf16_polish_ci', value: 'utf16_polish_ci' },
  { label: 'utf16_estonian_ci', value: 'utf16_estonian_ci' },
  { label: 'utf16_spanish_ci', value: 'utf16_spanish_ci' },
  { label: 'utf16_swedish_ci', value: 'utf16_swedish_ci' },
  { label: 'utf16_turkish_ci', value: 'utf16_turkish_ci' },
  { label: 'utf16_czech_ci', value: 'utf16_czech_ci' },
  { label: 'utf16_danish_ci', value: 'utf16_danish_ci' },
  { label: 'utf16_lithuanian_ci', value: 'utf16_lithuanian_ci' },
  { label: 'utf16_slovak_ci', value: 'utf16_slovak_ci' },
  { label: 'utf16_spanish2_ci', value: 'utf16_spanish2_ci' },
  { label: 'utf16_roman_ci', value: 'utf16_roman_ci' },
  { label: 'utf16_persian_ci', value: 'utf16_persian_ci' },
  { label: 'utf16_esperanto_ci', value: 'utf16_esperanto_ci' },
  { label: 'utf16_hungarian_ci', value: 'utf16_hungarian_ci' },
  { label: 'utf16_sinhala_ci', value: 'utf16_sinhala_ci' },
  { label: 'utf16_german2_ci', value: 'utf16_german2_ci' },
  { label: 'utf16_croatian_ci', value: 'utf16_croatian_ci' },
  { label: 'utf16_vietnamese_ci', value: 'utf16_vietnamese_ci' },
  { label: 'utf16_unicode_520_ci', value: 'utf16_unicode_520_ci' }
]
const SQL_COLLATION_UTF16LE = [
  { label: 'utf16le_general_ci', value: 'utf16le_general_ci' },
  { label: 'utf16le_bin', value: 'utf16le_bin' }
]
const SQL_COLLATION_CP1256 = [
  { label: 'cp1256_general_ci', value: 'cp1256_general_ci' },
  { label: 'cp1256_bin', value: 'cp1256_bin' }
]
const SQL_COLLATION_CP1257 = [
  { label: 'cp1257_lithuanian_ci', value: 'cp1257_lithuanian_ci' },
  { label: 'cp1257_bin', value: 'cp1257_bin' },
  { label: 'cp1257_general_ci', value: 'cp1257_general_ci' }
]
const SQL_COLLATION_UTF32 = [
  { label: 'utf32_general_ci', value: 'utf32_general_ci' },
  { label: 'utf32_bin', value: 'utf32_bin' },
  { label: 'utf32_unicode_ci', value: 'utf32_unicode_ci' },
  { label: 'utf32_icelandic_ci', value: 'utf32_icelandic_ci' },
  { label: 'utf32_latvian_ci', value: 'utf32_latvian_ci' },
  { label: 'utf32_romanian_ci', value: 'utf32_romanian_ci' },
  { label: 'utf32_slovenian_ci', value: 'utf32_slovenian_ci' },
  { label: 'utf32_polish_ci', value: 'utf32_polish_ci' },
  { label: 'utf32_estonian_ci', value: 'utf32_estonian_ci' },
  { label: 'utf32_spanish_ci', value: 'utf32_spanish_ci' },
  { label: 'utf32_swedish_ci', value: 'utf32_swedish_ci' },
  { label: 'utf32_turkish_ci', value: 'utf32_turkish_ci' },
  { label: 'utf32_czech_ci', value: 'utf32_czech_ci' },
  { label: 'utf32_danish_ci', value: 'utf32_danish_ci' },
  { label: 'utf32_lithuanian_ci', value: 'utf32_lithuanian_ci' },
  { label: 'utf32_slovak_ci', value: 'utf32_slovak_ci' },
  { label: 'utf32_spanish2_ci', value: 'utf32_spanish2_ci' },
  { label: 'utf32_roman_ci', value: 'utf32_roman_ci' },
  { label: 'utf32_persian_ci', value: 'utf32_persian_ci' },
  { label: 'utf32_esperanto_ci', value: 'utf32_esperanto_ci' },
  { label: 'utf32_hungarian_ci', value: 'utf32_hungarian_ci' },
  { label: 'utf32_sinhala_ci', value: 'utf32_sinhala_ci' },
  { label: 'utf32_german2_ci', value: 'utf32_german2_ci' },
  { label: 'utf32_croatian_ci', value: 'utf32_croatian_ci' },
  { label: 'utf32_unicode_520_ci', value: 'utf32_unicode_520_ci' },
  { label: 'utf32_vietnamese_ci', value: 'utf32_vietnamese_ci' }
]

const SQL_COLLATION_BINARAY = [
  { label: 'binary', value: 'binary' }
]
const SQL_COLLATION_GEOSTD8 = [
  { label: 'geostd8_general_ci', value: 'geostd8_general_ci' },
  { label: 'geostd8_bin', value: 'geostd8_bin' }
]
const SQL_COLLATION_CP932 = [
  { label: 'cp932_japanese_ci', value: 'cp932_japanese_ci' },
  { label: 'cp932_bin', value: 'cp932_bin' }
]
const SQL_COLLATION_EUCJPMS = [
  { label: 'eucjpms_japanese_ci', value: 'eucjpms_japanese_ci' },
  { label: 'eucjpms_bin', value: 'eucjpms_bin' }
]

export {
  ROLE_OPTIONS,
  LOGIN_OPTIONS,
  LDAP_TYPE_OPTIONS,
  USER_STATUS_OPTIONS,
  RULE_TYPE_OPTIONS,
  RULE_TYPE_MAP,
  IP_TYPE_OPTIONS,
  IP_TYPE_MAP,
  RULE_MATCH_OPTIONS,
  SQL_TYPE_OPTIONS,

  SQL_CHARSET,
  SQL_COLLATION_BIG5,
  SQL_COLLATION_DEC8,
  SQL_COLLATION_CP850,
  SQL_COLLATION_HP8,
  SQL_COLLATION_KOI8R,
  SQL_COLLATION_LATIN1,
  SQL_COLLATION_LATIN2,
  SQL_COLLATION_SWE7,
  SQL_COLLATION_ASCII,
  SQL_COLLATION_UJIS,
  SQL_COLLATION_SJIS,
  SQL_COLLATION_HEBREW,
  SQL_COLLATION_TIS620,
  SQL_COLLATION_EUCKR,
  SQL_COLLATION_KOI8U,
  SQL_COLLATION_GB2312,
  SQL_COLLATION_GREEK,
  SQL_COLLATION_CP1250,
  SQL_COLLATION_GBK,
  SQL_COLLATION_LATIN5,
  SQL_COLLATION_ARMSCII8,
  SQL_COLLATION_UTF8,
  SQL_COLLATION_UCS2,
  SQL_COLLATION_CP886,
  SQL_COLLATION_KEYBCS2,
  SQL_COLLATION_MACCE,
  SQL_COLLATION_MACROMAN,
  SQL_COLLATION_CP852,
  SQL_COLLATION_LATIN7,
  SQL_COLLATION_UTF8MB4,
  SQL_COLLATION_CP1251,
  SQL_COLLATION_UTF16,
  SQL_COLLATION_UTF16LE,
  SQL_COLLATION_CP1256,
  SQL_COLLATION_CP1257,
  SQL_COLLATION_UTF32,
  SQL_COLLATION_BINARAY,
  SQL_COLLATION_GEOSTD8,
  SQL_COLLATION_CP932,
  SQL_COLLATION_EUCJPMS
}
