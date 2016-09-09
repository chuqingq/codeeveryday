package main

/*
#cgo CFLAGS: -I /home/panshangbin/semantic/include
#cgo LDFLAGS: -L /home/panshangbin/semantic/libs -lmsc -ldl -lpthread

#include "msp_cmn.h"
#include "msp_errors.h"
#include "stdlib.h"
*/
import "C"
import (
	"bytes"
	"crypto/x509"
	"crypto/rsa"
	"crypto/rand"
	"crypto/aes"
	"crypto/cipher"
	"crypto/tls"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
	"runtime"
	"github.com/satori/go.uuid"
	"unsafe"
)

type Push struct {
	Version   uint8
	Type      uint8
	AppID     uint16
	RequestID uint32
	DeviceID  [8]byte // 8字节
	Length    uint16
	Checksum  uint16
	Option    [4]byte
}

const PUSH_LEN int = 24

const PUB_KEY_STR string = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAm8GxKGWpMimDFilRyRhWDmqtPKZfh5FmpoZVaGL6j8MhWuc0MRhYEL83KOXHavftdN+h6oeA+GVbHcL8y7jOcDR5zSaLz6kA5wjOn42zKWTSSzqZxpJj1uGN+qhibjWUK/xywI0jVIF6W+zStXMKMNMVZh81VVDhLmtgCw2OaKli46aBZ8v0Aw9/JA3iPfflUMYxQHVg9kdvQqKHrfsZYBw03ej/5gpv+EpRbudzhrdbTtjSTFQrPCW607kwqbe2npQbkUHW74SG1C6xAhA/qK7G7w0CPUzNzz/pdZn0a22cgVVKEBbkqOJt68lZJRH3nTEW40j3dkH0wInnl2ja2QIDAQAB"

const PRI_KEY_STR string = "MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCbwbEoZakyKYMWKVHJGFYOaq08pl+HkWamhlVoYvqPwyFa5zQxGFgQvzco5cdq9+1036Hqh4D4ZVsdwvzLuM5wNHnNJovPqQDnCM6fjbMpZNJLOpnGkmPW4Y36qGJuNZQr/HLAjSNUgXpb7NK1cwow0xVmHzVVUOEua2ALDY5oqWLjpoFny/QDD38kDeI99+VQxjFAdWD2R29Cooet+xlgHDTd6P/mCm/4SlFu53OGt1tO2NJMVCs8JbrTuTCpt7aelBuRQdbvhIbULrECED+orsbvDQI9TM3PP+l1mfRrbZyBVUoQFuSo4m3ryVklEfedMRbjSPd2QfTAieeXaNrZAgMBAAECggEAGcM34k6uZbWoEQpUlMaJtWi/rsB2HJ5YNEMT7WgxuYW1BqwnXdeA+YQnQ4R+L5tCk4pJ5djz5CIfqBSQa8Hto3GKk/xEM9zoYU57nrh5YedjQT44ITgle21jZopjfYcvMvdWo7K0nU2tR3csgwa8MMc5SuLul2YBWQQ5pppfa8AVmuJ9erBiXb+Jaw8phmFpYjYH8eo+QEKhg7ToCB8VImLIfxv0qsDnGesFpbvIzYNdvz3vyWgzf5u8X3gETlgkGY2UoJJKiF5opbrE+bOa/rDFcgzDkUTRoSMTdmTsmNyk4tp1huNHr2bfdcSUVFLiZvQFt3/zFNTKMQHmsQZV4QKBgQDuWkPZspm6R7nqsEBhDJri6fGeQocFkJQRuDh0oPiOjFyxZhXWc1w+tJCIbUySRrKKBePIs2XKQ52HEk7bSGYnVYRw6h3ctMLrtUMb/H0QTn17iz+xw0fSFo4/AU69uAHLVdq3ePzOxp/drGPRhJJA/mauhv42iD0QgxkLidFQ3wKBgQCnSecWHAZGPsjYyJ5aVNASwMZagM7NQ0tpe4a0FtrujCVHV/q/fSEEnEigXlX0fYLyGOHKIefW3hmVi6fV5CRZBcMxZ07Kjm4TLrWsV+9Fl4F99NkVk7RuQuEq5FHxyzJvgUS392/kst9m84pSU8/Zh2aK2fxJacVvJ85jEaEzRwKBgQDC6N2DMCG1yuGloOuEcSJXXKdQm2Z+jnQG6XaBKQEY0H8cMja5XyyXumBWr8pl85ocdCSJAurCM/ilc7s4ZkPi9nOPQmOZD9g1l8yBHj/HDehfFsfHPcGFcxxvOUqCqe4NsO9iCXXyQUqJo2cc9iQDMgYVwh4vycjlr87TOKgKUQKBgGR80RIH3YD+n9kQkYaDYcWSBNRCgXbtUHRZXi35eKNIjfAQGjBCcr35PusOH6XQawMQDTlFKqV4HnglPrkN5QOQoZKgksS7z8U4Dqsq2zC7dG570JbUddKx293O7qZGv9IZHXVAbfc7t1R5QIJ5k+YAHomTradPoOhHSgNaiLFrAoGARxanlXi0lnrJr4mv3gknwgS3WDN/uLNlJazB/kbZhMPx5ygwXO8/R4iZCtCqdu+bXgTE7Lw8F8+SNBBkdT2PXVcG0pGLFoQ+bNA0uvnfrPLh96XukvSgnnD5MIpz3w1HIMFkIMX8Apmmyl0va6uKD4RcPBnsXxCVXfB/owQMcuI="

func (msg *Push) Encode() []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, msg)
	by := buf.Bytes()
	// log.Printf("encode :%v\n", by)
	return by
}

func (msg *Push) Decode(b []byte) (int, int, string) {
	if len(b) < PUSH_LEN {
		return 0, 0, ""
	}

	buf := bytes.NewReader(b)
	binary.Read(buf, binary.BigEndian, msg)
	if len(b) >= (PUSH_LEN + int(msg.Length)) {
		return PUSH_LEN + int(msg.Length), int(b[20] & 0x01), string(b[PUSH_LEN:(PUSH_LEN + int(msg.Length))])
	}

	return 0, 0, ""
}

func main() {
	server := flag.String("server", "117.78.39.181:9877", "msgpusher tcp port for terminals")
	deviceid := flag.String("deviceid", "0555555555555555", "deviceid, 16 numbers")
	flag.Parse()

	// set gomaxprocs
    runtime.GOMAXPROCS(runtime.NumCPU())

	// semantic
	loginParam := C.CString("appid = 56d4eb1c, work_dir = .")
	defer C.free(unsafe.Pointer(loginParam))
	C.MSPLogin(nil, nil, loginParam)

	// terminal -> push: online
	conn, err := net.Dial("tcp", *server)
	if err != nil {
		log.Fatalf(" Dial error: %s\n", err.Error())
	}
	defer conn.Close()

	deviceidbytes, err := hex.DecodeString(*deviceid)
	if err != nil {
		log.Printf("deviceid invalid: %v\n", err)
		return
	}

	var msg *Push = &Push{
		Version: 0x01,
		// DeviceID: deviceidbytes, /*[8]byte{'\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00'}*/
		// [8]byte{'\xb9', '\x09', '\xe2', '\xc5', '\xb6', '\xb4', '\x4d', '\x10'}, // 魏剑锋手机
		AppID:     0,
		RequestID: 0,
		Type:      1,
		Length:    0,
		Checksum:  0,
		Option:    [4]byte{'\x00', '\x00', '\x00', '\x00'},
	}
	n := copy(msg.DeviceID[:], deviceidbytes)
	if n != len(msg.DeviceID) {
		log.Printf("deviceid len invalid: %d\n", n)
		return
	}

	n, err = conn.Write(msg.Encode())
	if err != nil {
		log.Printf("write error: %s\n", err.Error())
	}
	// log.Printf("1111terminal -> push: %d\n", n)

	buf := make([]byte, 1024)
	var rest int = 0
	var used int = 0

	for {
		n, err := conn.Read(buf[rest:1024])
		if err != nil {
			log.Printf("read error: %s", err.Error())
			return
		}
		// log.Printf("recv %d\n", n)

		// 从buf中尽可能多的处理包
		used = 0
		for {
			length, mt, message := msg.Decode(buf[used:(rest + n)])
			if length <= 0 {
				break
			}
			// log.Printf("push -> terminal: \n")
			if (len(message) > 0) {
				go handlePushMessage(mt, message)
			}
			used += length

			// 根据情况处理包
			if msg.Type == 0x10 { // 推送请求
				msg.Type = 0x11 // 推送响应
				msg.Length = 0
				_, err = conn.Write(msg.Encode())
				if err != nil {
					log.Printf("write error: %s\n", err.Error())
				}
				// log.Printf("terminal -> push: %#v\n", msg)
			}
		}

		// buf中内容前移
		for i := 0; i < rest+n-used; i++ {
			buf[i] = buf[used+i]
		}
		rest = rest + n - used
	}
}

type MPlus struct {
	Msg		MPlusMsg 	`json:"msg"`
	Action	string 		`json:"action"`
}

type MPlusMsg struct {
	Message 		string 	`json:"message"`
	MessageId 		string 	`json:"messageId"`
	MessageReplyId 	string 	`json:"messageReplyId"`
	MessageType		int 	`json:"messageType"`
	MsgLocalId		string 	`json:"msgLocalId"`
	Res				string 	`json:"res"`
	Timestamp		int64 	`json:"timestamp"`
}

type SendMPlus struct {
	Sender			string 	`json:"sender"`
	DeviceId		string	`json:"deviceId"`
	SenderType		int		`json:"senderType"`
	Receiver		string	`json:"receiver"`
	ReceiverType	int		`json:"receiverType"`
	MessageInfo		string 	`json:"messageInfo"`
}

type GetPubDest struct {
	Addr	string	`json:"addr"`
	Type 	int		`json:"type"`
}

type GetPub struct {
	Contact 	string			`json:"contact"`
	Type 		int				`json:"type"`
	DeviceId 	string			`json:"deviceId"`
	Dest 		[]GetPubDest	`json:"dest"`
}

type SearchResultDetail struct {
	Text	string 	`json:"text"`
	Type 	string 	`json:"type"`
}

type SearchResult struct {
	Rc			int 				`json:"rc"`
	Operation	string				`json:"operation"`
	Service		string				`json:"service"`
	Answer		SearchResultDetail	`json:"answer"`
	Text		string				`json:"text"`
}

func handlePushMessage(mt int, message string) {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	var content string
	if (mt == 1) {
		resp, _ := client.Get("http://lbpush.powerapp.io/push/message?device_id=555555555555555&message_id=" + message)
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		content = string(body)
	} else {
		content = message
	}

	var tMap map[string]string
	json.Unmarshal([]byte(content), &tMap)

	_, ok := tMap["msg"]
	if (ok) {
		return
	}

	var mplus MPlus
	json.Unmarshal([]byte(tMap["message"]), &mplus)

	// receipt := {}
	receipt := MPlusMsg{
		Message: mplus.Msg.MessageId,
		MessageId: uuid.NewV4().String(),
		MessageReplyId: "",
		MessageType: 666,
		MsgLocalId: "",
		Res: "55555",
		Timestamp: time.Now().UnixNano() / 1000000,
	}
	receiptJson, _ := json.Marshal(receipt)
	postBuf, _ := json.Marshal(SendMPlus{
		Sender: "55555",
		DeviceId: "555555555555555",
		SenderType: 0,
		Receiver: mplus.Msg.Res,
		ReceiverType: 0,
		MessageInfo: string(receiptJson),
	})
	receiptReq, _ := http.NewRequest("POST", "https://lbserver.powerapp.io/webmessage/v2/user/message/_send", bytes.NewReader(postBuf))
	receiptReq.Header.Add("Content-Type", "application/json")
	receiptReq.Header.Add("Device-Token", "1,developer,1464c17ea1a1112")
	client.Do(receiptReq)
	// receiptRespBody, _ := client.Do(receiptReq)
	// defer receiptResp.Body.Close()
	// receiptRespBody, _ := ioutil.ReadAll(receiptResp.Body)
	// log.Printf("%v\n", string(receiptRespBody))

	// unpacket package
	packet, _ := base64.StdEncoding.DecodeString(mplus.Msg.Message)
	aesKeyLen := binary.BigEndian.Uint32(packet[0:4])

	encryptedAESKey := packet[4:(4 + aesKeyLen)]
	encryptedAESData := packet[(4 + aesKeyLen):]

	priData, _ := base64.StdEncoding.DecodeString(PRI_KEY_STR)
	priInterface, _ := x509.ParsePKCS8PrivateKey(priData)
	prikey := priInterface.(*rsa.PrivateKey)
	decryptedAESKey, _ := rsa.DecryptPKCS1v15(rand.Reader, prikey, encryptedAESKey)

	aesKey := decryptedAESKey[0:32]
	aesIV := decryptedAESKey[32:48]
	block, _ := aes.NewCipher(aesKey)
	blockmode := cipher.NewCBCDecrypter(block, aesIV)
	dst := make([]byte, len(encryptedAESData))
	blockmode.CryptBlocks(dst, encryptedAESData)
	dst = PKCS5UnPadding(dst)
	log.Printf("receive message: %v\n", string(dst))

	searchParam := C.CString("nlp_version=2.0")
	searchContent := C.CString(string(dst))
	defer C.free(unsafe.Pointer(searchParam))
	defer C.free(unsafe.Pointer(searchContent))
	var searchLen uint;
	var searchErr int;
	searchRet := C.GoString(C.MSPSearch(searchParam, searchContent, (*C.uint)(unsafe.Pointer(&searchLen)), (*C.int)(unsafe.Pointer(&searchErr))))

	retMsg := ""
	// get answer
	var searchResult SearchResult
	json.Unmarshal([]byte(searchRet), &searchResult)
	if (searchResult.Rc == 0) {
		retMsg = searchResult.Answer.Text
	} else {
		retMsg = "对不起，现在我还不能理解你说的，请见谅..."
	}
	// return
	encrypterBlockmode := cipher.NewCBCEncrypter(block, aesIV)
	retMsgBytes := []byte(retMsg)
	blockSize := block.BlockSize()
	retMsgBytes = PKCS5Padding(retMsgBytes, blockSize)
	retCryptedBytes := make([]byte, len(retMsgBytes))
	encrypterBlockmode.CryptBlocks(retCryptedBytes, retMsgBytes)

	dest := make([]GetPubDest, 1)
	dest[0] = GetPubDest{
		Addr: mplus.Msg.Res,
		Type: 0,
	}
	postGetPubBuf, _ := json.Marshal(GetPub{
		Contact: "55555",
		Type: 0,
		DeviceId: "555555555555555",
		Dest: dest,
	})
	getPubReq, _ := http.NewRequest("POST", "https://lbserver.powerapp.io/webmessage/key/_get", bytes.NewReader(postGetPubBuf))
	getPubReq.Header.Add("Content-Type", "application/json")
	getPubReq.Header.Add("Device-Token", "1,developer,1464c17ea1a1112")
	getPubResp, _ := client.Do(getPubReq)
	defer getPubResp.Body.Close()
	getPubRespBody, _ := ioutil.ReadAll(getPubResp.Body)
	var tPeerMap []map[string]string
	json.Unmarshal(getPubRespBody, &tPeerMap)
	peerPublicKey := tPeerMap[0]["publicKey"]

	pubData, _ := base64.StdEncoding.DecodeString(peerPublicKey)
    pubInterface, _ := x509.ParsePKIXPublicKey(pubData)
    pubkey := pubInterface.(*rsa.PublicKey)
    peerEncryptedAESKey, _ := rsa.EncryptPKCS1v15(rand.Reader, pubkey, decryptedAESKey)

	retBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(retBytes, (uint32)(len(peerEncryptedAESKey)))
	retBytes = append(retBytes, peerEncryptedAESKey...)
	retBytes = append(retBytes, retCryptedBytes...)

	ret := MPlusMsg{
		Message: base64.StdEncoding.EncodeToString((retBytes)),
		MessageId: uuid.NewV4().String(),
		MessageReplyId: "",
		MessageType: 1,
		MsgLocalId: "",
		Res: "55555",
		Timestamp: time.Now().UnixNano() / 1000000,
	}
	retJson, _ := json.Marshal(ret)
	postRetBuf, _ := json.Marshal(SendMPlus{
		Sender: "55555",
		DeviceId: "555555555555555",
		SenderType: 0,
		Receiver: mplus.Msg.Res,
		ReceiverType: 0,
		MessageInfo: string(retJson),
	})
	retReq, _ := http.NewRequest("POST", "https://lbserver.powerapp.io/webmessage/v2/user/message/_send", bytes.NewReader(postRetBuf))
	retReq.Header.Add("Content-Type", "application/json")
	retReq.Header.Add("Device-Token", "1,developer,1464c17ea1a1112")
	client.Do(retReq)
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
