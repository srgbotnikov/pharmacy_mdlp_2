package sign

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

type CriptoSign struct {
	Thumbprint string
	SignShPath string
}

func NewCriptoSign(thumbprint, signShPath string) *CriptoSign {
	return &CriptoSign{
		Thumbprint: thumbprint,
		SignShPath: signShPath,
	}
}

func (p *CriptoSign) SignFile(data, requestID string) (string, string, error) {
	// log.Println("Sign file")

	filePath, err := p.createTmpFile(data, requestID)
	if err != nil {
		return "", "", err
	}
	cmd := exec.Command("/bin/sh", p.SignShPath+"/sign.sh", p.Thumbprint, filePath, filePath+"_out.sig", filePath+"_signed.txt")

	_, err = cmd.CombinedOutput()
	if err != nil {
		return "", "", errors.New("Error call sign.sh " + err.Error())
	}

	cmd = exec.Command("/bin/sh", p.SignShPath+"/base64.sh", filePath, filePath+"_base64.txt")

	_, err = cmd.CombinedOutput()

	if err != nil {
		return "", "", err
	}

	// Читать файл подписи
	b, err := ioutil.ReadFile(filePath + "_signed.txt")
	if err != nil {
		return "", "", err
	}
	sign := string(b)

	b, err = ioutil.ReadFile(filePath + "_base64.txt")
	if err != nil {
		return "", "", err
	}
	datasigned := string(b)

	//Удалить временные файлы
	p.deleteTmpFile(filePath)
	p.deleteTmpFile(filePath + "_out.sig")
	p.deleteTmpFile(filePath + "_signed.txt")
	p.deleteTmpFile(filePath + "_base64.txt")

	return datasigned, sign, nil
}

func (p *CriptoSign) GetThumbprint() (string, error) {
	if p.Thumbprint != "" {
		return p.Thumbprint, nil
	}
	return "", errors.New("the fingerprint does not exist")
}

//CurlUploadWebDAVFile загрузка документов большого объема через curl. В link надо менять url на адрес stunnel и передавать его
func (p *CriptoSign) CurlUploadWebDAVFile(token, data, requestID, link string) error {
	st := "'" + "Authorization: token " + token + "'"

	filePath, err := p.createTmpFile(data, requestID)
	if err != nil {
		return err
	}

	shStr := "/opt/cprocsp/bin/amd64/curl -T " + filePath + " " + link + " -H " + st

	file, err := os.Create(p.SignShPath + "/curlWebDav.sh")
	if err != nil {
		return err
	}
	fmt.Fprint(file, shStr)
	file.Close()

	cmd := exec.Command("/bin/sh", p.SignShPath+"/curlWebDav.sh")
	_, err = cmd.CombinedOutput()
	if err != nil {
		p.deleteTmpFile(filePath)
		return err
	}
	p.deleteTmpFile(filePath)
	return nil
}

//CurlDownloadFile скачать документ с помощью curl
func (p *CriptoSign) CurlDownloadFile(token, link string) (string, error) {
	st := "'" + "Authorization: token " + token + "'"

	shStr := "/opt/cprocsp/bin/amd64/curl -sb GET " + link + " -H " + st

	file, err := os.Create(p.SignShPath + "/curl1.sh")
	if err != nil {
		return "", err
	}
	fmt.Fprint(file, shStr)
	file.Close()

	cmd := exec.Command("/bin/sh", p.SignShPath+"/curl1.sh")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(out), nil
}

//CreateTmpFile Создание временного файла для подписи
func (p *CriptoSign) createTmpFile(data, requestID string) (string, error) {
	file, err := os.Create(p.SignShPath + "/" + requestID)
	if err != nil {
		return "", err
	}
	defer file.Close()

	file.WriteString(data)

	return p.SignShPath + "/" + requestID, nil
}

//DeleteTmpFile Удаление файлов
func (p *CriptoSign) deleteTmpFile(fileName string) (err error) {
	err = os.Remove(fileName)
	if err != nil {
		return err
	}

	return nil
}
