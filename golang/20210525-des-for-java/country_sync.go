package main

import (
	"crypto/des"
	"encoding/base64"
	"log"
	"net/http"
	// "io"
	"io/ioutil"
	"bytes"
	"errors"
)

func main() {
	projectJsonStr := []byte(`{"projectDeclare":{"id":"58219b44-251c-4817-8b41-48485ca5635a","xmbm":"1aac7c83-f6c7-48f4-acba-78f670bbf0c1","name":"曹蓓蓓测试","type":"B","investmentType":"","constructionType":"设计施工","projectDebriefing":"意向谋划","projectDebriefingDetial":"13122223333","bjdsfzc":3,"jfknqk":"","kxwjqqk":"","yyknqk":"","yytcqk":"","assess":"3,","assessContent":"","assessExplanation":"","szgyhzcxxsm":"","zgyhzcxxsm":"","standard":0,"gbhdz":"上海发改委","realm":"A0001","industry":"B0002","twoIndustry":"高速","jsnrhmb":"13122223333","xmjhztz":1.3122223333E10,"cjqshte":1.3122223333E10,"xmdwzje":1.3122223333E10,"xmwczje":1.3122223333E10,"xmzqzje":1.0,"xmgqzje":1.0,"lwyyqk":"无","zddbjkqk":2,"yhckmfxd":2,"qttjdw":2,"qttjdwqc":"","bsdwqc":"上帝街区","mlwt":"资金困难","mlwtms":"13122223333","xybgzkl":"13122223333","sjkgsj":"2020/11/29 08:00:00","sjwgsj":"2020/12/17 08:00:00","pshjq":2,"pshjqms":"","gfltcgqd":2,"gfltcgqdms":"","dsfhz":2,"dsfhzms":"","xcld":2,"xcldms":"","discountPolicy":"","countryDiscountPolicy":"","chinaDiscountPolicy":"","projectSignificance":"13122223333","isShareInfo":2,"chinaGqzb":"1","projectGqzb":"","otherGqzb":"","chinaZqzb":"1","projectZqzb":"","otherZqzb":"","gardenNum":"","gardenAmount":"","createJobNum":"","taxAmount":"","turnover":0.0,"profit":0.0,"createId":"a28aa669-5058-4836-9c2a-d9a3b432e3d0","createName":"上帝街区","createDate":"2020/12/03 10:02:58","updateId":"a28aa669-5058-4836-9c2a-d9a3b432e3d0","updateName":"上帝街区","updateDate":"2020/12/03 15:33:33","isSubmit":1,"stage":2,"isNewProject":1},"projcectLnvestments":[{"id":"206026ba-d9e1-441e-b08a-061de6b606de","projectId":"58219b44-251c-4817-8b41-48485ca5635a","tzdwmcName":"","btzdwmcName":"","sort":1}],"projectConstructions":[{"id":"6ec63eed-bd6c-44d7-8e47-d10da66394fc","projectId":"58219b44-251c-4817-8b41-48485ca5635a","yzdwName":"123","jsdwName":"123","sort":1}],"projectGqgcs":[{"id":"49f309bd-7bf0-44ae-807c-eb62024321fc","projectId":"58219b44-251c-4817-8b41-48485ca5635a","type":3,"gdmc":"","sort":3},{"id":"5c572690-2689-4e1b-89c2-4a312e0c3a9c","projectId":"58219b44-251c-4817-8b41-48485ca5635a","type":2,"gdmc":"","sort":2},{"id":"e6e21c9f-0f92-45d6-9141-8c1e14cb8caa","projectId":"58219b44-251c-4817-8b41-48485ca5635a","type":1,"gdmc":"1","sort":1}],"projectXmfxes":[{"id":"a863992a-f187-44a3-b8ff-3752d44c2fe0","projectId":"58219b44-251c-4817-8b41-48485ca5635a","fxType":"未分析或无法评估","fxcd":"","description":"","sort":1}],"projectXmlxrs":[{"id":"136b9b50-279e-41db-9b75-08c1b74d3af2","projectId":"58219b44-251c-4817-8b41-48485ca5635a","name":"上海发改委","phone":"13122223333","type":3,"sort":1},{"id":"d467a0c4-78c9-4ec2-8397-725371c38e2e","projectId":"58219b44-251c-4817-8b41-48485ca5635a","name":"13122223333","phone":"13122223333","type":4,"sort":2}],"projectZqgcs":[{"id":"000a5cc1-3bd9-44a2-9bf5-733a22bfbc42","projectId":"58219b44-251c-4817-8b41-48485ca5635a","type":2,"gdmc":"","sort":2},{"id":"28afe9b0-2a45-4ff4-932a-f543ff31246d","projectId":"58219b44-251c-4817-8b41-48485ca5635a","type":3,"gdmc":"","sort":3},{"id":"318bc5b4-34ad-496d-89a1-4ebb6744bcd5","projectId":"58219b44-251c-4817-8b41-48485ca5635a","type":1,"gdmc":"1","sort":1}],"projectCountryProvinces":[{"id":"6066cc99-f7e6-4c1c-b19a-524f03b62697","projectId":"58219b44-251c-4817-8b41-48485ca5635a","dz":"Oceania","dzName":"大洋洲","gbhdz":"AU","gbhdzName":"澳大利亚","province":"AU-ACT","provinceName":"澳大利亚首都直辖区","address":"13122223333","sort":1}]}`)
	// projectJsonStr := []byte("1234567890")
	log.Printf("data: %v", projectJsonStr)
	key := []byte("tyxthjwt") // tyxthjwtzxtKEY0001_A
	// 加密
	result, err := MyEncrypt(projectJsonStr, key)
	if err != nil {
		log.Printf("encrypt error: %v", err)
		return
	}
	log.Printf("result: %v", result)
	result2, err := MyDecrypt(result, key)
	if err != nil {
		log.Printf("decrypt error: %v", err)
		return
	}
	log.Printf("result2: %v", result2)
	// if result2 != projectJsonStr {
	// 	log.Printf("error encrypt and decrypt not match")
	// 	return
	// }
	// base64编码
	str := base64.StdEncoding.EncodeToString(result)
	log.Printf("base64: %v", str)
	// TODO 调用接口
	resp, err := http.Post("http://lyw547479149.oicp.net/ydyl/transport/saveProjectInfoJson",
		"application/json", bytes.NewReader([]byte(str)))
	if err != nil {
		log.Printf("http.Post error: %v", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("read body error: %v", err)
		return
	}
	log.Printf("body: %v", string(body))
}

func MyEncrypt(data, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	bs := block.BlockSize()
	data = PKCS5Padding(data, bs)
	if len(data)%bs != 0 {
		return nil, errors.New("Need a multiple of the blocksize")
	}
	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		block.Encrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	return out, nil
}
func MyDecrypt(data []byte, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	bs := block.BlockSize()
	if len(data)%bs != 0 {
		return nil, errors.New("crypto/cipher: input not full blocks")
	}
	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		block.Decrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	out = PKCS5UnPadding(out)
	return out, nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}