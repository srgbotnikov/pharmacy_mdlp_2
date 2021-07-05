package sign

type CryptoproCSP interface {
	//SignFile подписать файл
	SignFile(data, requestID string) (string, string, error)
	//GetThumbprint получить отпечаток текущей подписи
	GetThumbprint() (string, error)
	//CurlUploadWebDAVFile загрузить документ большого объема с помощью curl
	CurlUploadWebDAVFile(token, data, requestID, link string) error
	//CurlDownloadFile скачать документ с помощью curl
	CurlDownloadFile(token, link string) (string, error)
}

type Sign struct {
	CryptoproCSP
}

type Config struct {
	Thumbprint string
	SignShPath string
}

func NewSign(cfg Config) *Sign {
	return &Sign{
		CryptoproCSP: NewCriptoSign(cfg.Thumbprint, cfg.SignShPath),
	}
}
