package leetcode

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/user"
	"path"
	"time"

	"github.com/openset/leetcode/internal/base"
	"github.com/openset/leetcode/internal/client"
)

var err error

var checkErr = base.CheckErr

const authInfo = base.AuthInfo

func graphQLRequest(filename, jsonStr string, v interface{}) {
	data := remember(filename, func() []byte {
		return client.PostJson(graphqlUrl, jsonStr)
	})
	err = json.Unmarshal(data, &v)
	checkErr(err)
	return
}

func remember(filename string, f func() []byte) []byte {
	filename = getCachePath(filename)
	if fi, err := os.Stat(filename); err == nil {
		u := time.Now().AddDate(0, 0, -3)
		if fi.ModTime().After(u) {
			return fileGetContents(filename)
		}
	}
	data := f()
	var buf bytes.Buffer
	err = json.Indent(&buf, data, "", "\t")
	checkErr(err)
	filePutContents(filename, buf.Bytes())
	return data
}

func getCsrfToken(cookies []*http.Cookie) string {
	for _, cookie := range cookies {
		if cookie.Name == "csrftoken" {
			return cookie.Value
		}
	}
	return ""
}

func getCachePath(f string) string {
	dir, err := os.UserCacheDir()
	checkErr(err)
	u, err := user.Current()
	if err == nil && u.HomeDir != "" {
		dir = u.HomeDir
	}
	return path.Join(dir, ".leetcode", f)
}

func Clean() {
	dir := getCachePath("")
	err := os.RemoveAll(dir)
	checkErr(err)
}

func filePutContents(filename string, data []byte) {
	base.FilePutContents(filename, data)
}

func fileGetContents(filename string) []byte {
	data, err := ioutil.ReadFile(filename)
	checkErr(err)
	return data
}

func jsonEncode(v interface{}) []byte {
	data, err := json.Marshal(v)
	checkErr(err)
	var buf bytes.Buffer
	err = json.Indent(&buf, data, "", "\t")
	checkErr(err)
	return buf.Bytes()
}

func jsonDecode(data []byte, v interface{}) {
	err = json.Unmarshal(data, v)
	checkErr(err)
}

func saveCookies(cookies []*http.Cookie) {
	filePutContents(getCachePath(cookiesFile), jsonEncode(cookies))
}

func getCookies() (cookies []*http.Cookie) {
	data := fileGetContents(getCachePath(cookiesFile))
	jsonDecode(data, &cookies)
	return
}

func saveCredential(data url.Values) {
	filePutContents(getCachePath(credentialsFile), jsonEncode(data))
}

func getCredential() (data url.Values) {
	body := fileGetContents(getCachePath(credentialsFile))
	jsonDecode(body, &data)
	return
}