package roolink

type RooLink struct {
	APIKey       string
	ProtectedURL string
	UserAgent    string
}

type RequestLimit struct {
	Requests int64 `json:"requests"`
}

type SensorPayload struct {
	URL        string      `json:"url"`
	UserAgent  string      `json:"userAgent"`
	Abck       string      `json:"_abck"`
	BmSz       string      `json:"bm_sz"`
	ScriptData *ScriptData `json:"scriptData"`
	SecCpt     bool        `json:"sec_cpt"`
	Stepper    bool        `json:"stepper"`
	Index      int         `json:"index"`
	Keyboard   bool        `json:"keyboard"`
	Flags      string      `json:"flags"`
}

type SbsdPayload struct {
	UserAgent string `json:"userAgent"`
	Vid       string `json:"vid"`
	Cookie    string `json:"bm_o"`
	Static    bool   `json:"static"`
}

type PixelPayload struct {
	UserAgent            string `json:"userAgent"`
	Bazadebezolkohpepadr int64  `json:"bazadebezolkohpepadr"`
	Hash                 string `json:"hash"`
}

type CptChallenge struct {
	Token      string `json:"token"`
	Timestamp  int64  `json:"timestamp"`
	Nonce      string `json:"nonce"`
	Difficulty int64  `json:"difficulty"`
	Cookie     string `json:"cookie"`
}

type SensorResponse struct {
	Sensor string `json:"sensor_data"`
}

type SbsdBody struct {
	Body string `json:"body"`
}

type ScriptData struct {
	Ver string `json:"ver"`
	Dvc string `json:"dvc"`
	Key int64  `json:"key"`
	Din []int  `json:"din"`
}
