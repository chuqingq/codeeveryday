package main

import (
	"crypto/cipher"
	"crypto/des"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

const rose_url = "https://daisy.powerapp.io/rose/v1/dodb/1/databases/ota/collections/packages/documents?sort={publish_time:-1}&limit=200"
const authorization = "Bearer 2465247d798fc55e5fe4166ceb17163f"

type roseResult struct {
	Total   uint64       `json:"total"`
	Results []roseRecord `json:"results"`
}

type roseRecord struct {
	Id          string `json:"_id"`
	Version     string `json:"version"`
	DownloadUrl string `json:"download_url"`
}

func main() {
	res, err := getOtaResult()
	if err != nil {
		log.Panic(err)
	}

	// log.Printf("res:%#v\n", res)

	for _, record := range res.Results {
		downloadUrl, err := getNewOtaDownloadUrl(record.DownloadUrl)

		if err != nil {
			// log.Printf("%v: replace %v error: %v\n", record.Id, record.DownloadUrl, err)
			continue
		}

		log.Printf("%v: replace %v to %v\n", record.Id, record.DownloadUrl, downloadUrl)
		// TODO update to ota rose
		err = updateOtaRoseRecord(record.Id, downloadUrl)
		if err != nil {
			log.Printf("update error: %v\n", err)
		}
	}
}

///////////////////////////////////

func getOtaResult() (*roseResult, error) {
	client := &http.Client{}

	// get ota records
	request, _ := http.NewRequest("GET", rose_url, nil)
	request.Header.Add("Authorization", authorization)

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, errors.New(response.Status)
	}

	var res roseResult
	err = json.Unmarshal(bytes, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

var oldRe *regexp.Regexp

func init() {
	oldRe = regexp.MustCompile("http://.*/r/([^/]*)/(.*)")
}

const newDownloadUrlPrefix = "http://lfs.powerapp.io"

func getNewOtaDownloadUrl(old string) (string, error) {
	encryptedfid := oldRe.ReplaceAllString(old, "$1")
	if encryptedfid == old {
		return "", errors.New("url should not change")
	}
	fidBytes, err := ReadFidDecrypt(encryptedfid)
	if err != nil {
		return "", err
	}
	fid := string(fidBytes)
	filename := oldRe.ReplaceAllString(old, "$2")

	newDownloadUrl := newDownloadUrlPrefix + encodeDownloadUrl("/f/"+fid+"/"+filename+"?e=0")
	// TODO fid
	return newDownloadUrl, nil
}

// 读权限文件ID解密
func ReadFidDecrypt(crypted string) (string, error) {
	decoded, err := hex.DecodeString(crypted)
	if err != nil {
		return "", err
	}

	decrypted, err := desDecrypt(decoded, readKey)
	if err != nil {
		return "", err
	}

	return string(decrypted), nil
}

func desDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key)
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData, err = pKCS5UnPadding(origData)
	if err != nil {
		return nil, err
	}
	// origData = ZeroUnPadding(origData)
	return origData, nil
}

var readKey []byte = []byte("8uhb(OL>")

func pKCS5UnPadding(origData []byte) ([]byte, error) {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	if length < unpadding {
		return nil, errors.New("unpadding error!")
	} else {
		return origData[:(length - unpadding)], nil
	}
}

func encodeDownloadUrl(url string) string {
	h := hmac.New(sha1.New, []byte("5GIaNcWym9MLGLdM"))
	h.Write([]byte(url))

	token := hex.EncodeToString(h.Sum(nil))

	return url + "&t=PNGmeAEZ0sTwgfDh:" + token
}

////////

func updateOtaRoseRecord(id, downloadUrl string) error {
	client := &http.Client{}

	// get ota records
	request, _ := http.NewRequest("PUT",
		"https://daisy.powerapp.io/rose/v1/dodb/1/databases/ota/collections/packages/documents/"+id,
		strings.NewReader("{content:{$set:{\"download_url\": \""+downloadUrl+"\"}}}"))
	request.Header.Add("Authorization", authorization)

	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// bytes, err := ioutil.ReadAll(response.Body)
	// if err != nil {
	// 	return err
	// }

	if response.StatusCode != http.StatusNoContent {
		return errors.New(response.Status)
	}

	return nil
}
