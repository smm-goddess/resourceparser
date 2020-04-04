package items

const (
	MASK_UI_MODE_TYPE       = 0
	UI_MODE_TYPE_ANY        = 0x00
	UI_MODE_TYPE_NORMAL     = 0x01
	UI_MODE_TYPE_DESK       = 0x02
	UI_MODE_TYPE_CAR        = 0x03
	UI_MODE_TYPE_TELEVISION = 0x04
	UI_MODE_TYPE_APPLIANCE  = 0x05
	UI_MODE_TYPE_WATCH      = 0x06
	MASK_UI_MODE_NIGHT      = 0
	SHIFT_UI_MODE_NIGHT     = 0
	UI_MODE_NIGHT_ANY       = 0x00
	UI_MODE_NIGHT_NO        = 0x01
	UI_MODE_NIGHT_YES       = 0x02

	//screenLayout
	MASK_SCREENSIZE   = 0
	SCREENSIZE_ANY    = 0x00
	SCREENSIZE_SMALL  = 0x01
	SCREENSIZE_NORMAL = 0x02
	SCREENSIZE_LARGE  = 0x03
	SCREENSIZE_XLARGE = 0x04
	MASK_SCREENLONG   = 0
	SHIFT_SCREENLONG  = 0
	SCREENLONG_ANY    = 0x00
	SCREENLONG_NO     = 0x01
	SCREENLONG_YES    = 0x02
	MASK_LAYOUTDIR    = 0
	SHIFT_LAYOUTDIR   = 0
	LAYOUTDIR_ANY     = 0x00
	LAYOUTDIR_LTR     = 0x01
	LAYOUTDIR_RTL     = 0x02
)

type ResTableConfig struct {
	Size uint32
	/*
	   union {
	       struct {
	           // Mobile country code (from SIM).  0 means "any".
	           uint16_t mcc;
	           // Mobile network code (from SIM).  0 means "any".
	           uint16_t mnc;
	       };
	       uint32_t imsi;
	   };*/
	MccMncOrImsi uint32 // mcc + mnc 0 means any or imsi
	/** locale
	/*union {
	    struct {
	        char language[2];
	        char country[2];
	    };
	    uint32_t locale;
	};*/
	LanguageCountryOrLocale uint32
	/**
	   * union {
	      struct {
	          uint8_t orientation;
	          uint8_t touchscreen;
	          uint16_t density;
	      };
	      uint32_t screenType;
	  };
	*/
	ScreenTypeDataOrScreenType uint32
	/**
	   * union {
	      struct {
	          uint8_t keyboard;
	          uint8_t navigation;
	          uint8_t inputFlags;
	          uint8_t inputPad0;
	      };
	      uint32_t input;
	  };
	*/
	InputTypeDataOrInputType uint32
	/** screen size
	   * union {
	      struct {
	          uint16_t screenWidth;
	          uint16_t screenHeight;
	      };
	      uint32_t screenSize;
	  };
	*/
	ScreenSizeWHOrScreenSize uint32
	/** system version
	   * union {
	      struct {
	          uint16_t sdkVersion;
	          // For now minorVersion must always be 0!!!  Its meaning
	          // is currently undefined.
	          uint16_t minorVersion;
	      };
	      uint32_t version;
	  };
	*/
	VersionConfigOrVersion uint32
	// Screen config
	/**
	   * union {
	      struct {
	          uint8_t screenLayout;
	          uint8_t uiMode;
	          uint16_t smallestScreenWidthDp;
	      };
	      uint32_t screenConfig;
	  };
	*/
	ScreenConfigDataOrScreenConfig uint32
	/** screen size
	   * union {
	      struct {
	          uint16_t screenWidthDp;
	          uint16_t screenHeightDp;
	      };
	      uint32_t screenSizeDp;
	  };
	*/
	ScreenSizeDpWHOrScreenSize uint32
	LocaleScript               [4]byte
	LocaleVariant              [8]byte
}
