package main

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/kyl2016/Play-With-Golang/pkg/random"
	"github.com/kyl2016/Play-With-Golang/utility"
	"golang.org/x/net/html"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func download(url, saveTo string) {
	// Request the HTML page.
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	r, err := ioutil.ReadAll(res.Body)
	os.WriteFile(saveTo, r, os.ModePerm)
}

func getFakeChineseUserInfos() {
	rootFolder := "/Users/kyl/Documents/ihandy/fake_user/fakeUsers"
	url := "https://www.myfakeinfo.com/nationalidno/get-china-citizenidandname.php"
	os.Mkdir(rootFolder, os.ModePerm)
	for i := 1000; i < 2000; i++ {
		download(url, rootFolder+"/"+strconv.Itoa(i)+".php")
		time.Sleep(time.Second * 3)
	}
}

func getFakeEnglishUserPhones() {
	rootFolder := "/Users/kyl/Documents/ihandy/fake_user/fakeUsers-CN-Phone"
	url := "https://fake-name-generator.com/Fake-Chinese-Name-Generator"
	os.Mkdir(rootFolder, os.ModePerm)
	for i := 11199; i < 20000; i++ {
		download(url, rootFolder+"/"+strconv.Itoa(i)+".php")
		time.Sleep(time.Millisecond * 30)
	}
}

func getFakeEnglishUserInfos() {
	rootFolder := "/Users/kyl/Documents/ihandy/fake_user/fakeUsers-US"
	url := "https://fake-name-generator.com/Fake-US-Name-Generator"
	os.Mkdir(rootFolder, os.ModePerm)
	for i := 13879; i < 20000; i++ {
		download(url, rootFolder+"/"+strconv.Itoa(i)+".php")
		time.Sleep(time.Millisecond * 30)
	}
}

func getFakeEmail() string {
	domains := map[int]string{
		0: "qq.com", 1: "gmail.com", 2: "yahoo.com.cn", 3: "sina.com", 4: "126.com", 5: "163.com", 6: "hotmail.com", 7: "sohu.com", 8: "vip.163.com", 9: "netease.com", 10: "188.com", 11: "aliyun.com", 12: "189.cn", 13: "foxmail.com", 14: "outlook.com",
	}

	var letters = []interface{}{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", ".", "_"}
	var lettersWeightInt []float32
	for i := 0; i < 28; i++ {
		lettersWeightInt = append(lettersWeightInt, 10)
	}

	var email string
	for i := 0; i < 10; i++ {
		email += random.Discrete(letters, lettersWeightInt).(string)
	}

	email += "@" + domains[random.GetRandomNumber(14)]
	return email
}

func main() {
	//getFakeChineseUserInfos
	//go getFakeEnglishUserInfos()
	//go getFakeEnglishUserPhones()
	//select {}
	//println(getFakeEmail())

	//parseUserInfoCN()
	//parseUserInfoEN()
	//parseUserInfoCN_Phone()

	combine2()
}

func combine() {
	// process male
	var cnUsers []ShowlistMockUser
	usersCN := "users-cn.json"
	buf, _ := os.ReadFile(usersCN)
	json.Unmarshal(buf, &cnUsers)

	var enUsers []ShowlistMockUser
	usersEN := "users-en.json"
	buf, _ = os.ReadFile(usersEN)
	json.Unmarshal(buf, &enUsers)

	var phones []ShowlistMockUser
	phonesFile := "users-cn-phoneAndUserName.json"
	buf, _ = os.ReadFile(phonesFile)
	json.Unmarshal(buf, &phones)

	var newUsers []ShowlistMockUser

	genders := []string{"male", "female"}
	okCount := 0
	for _, gender := range genders {
		indexOfEnUsers := 0
		for i, u := range cnUsers {
			if okCount == 5000 {
				okCount++
				break
			} else if okCount >= 10000 {
				break
			}

			if strings.ToLower(u.Gender) == gender {
				if indexOfEnUsers >= len(enUsers) {
					break
				}
				for strings.ToLower(enUsers[indexOfEnUsers].Gender) != gender {
					indexOfEnUsers++
				}
				cnUsers[i].RealNameEn = enUsers[indexOfEnUsers].RealNameEn
				if enUsers[indexOfEnUsers].UserNameEn == "" {
					fmt.Println(enUsers[indexOfEnUsers])
				}
				cnUsers[i].UserNameEn = enUsers[indexOfEnUsers].UserNameEn
				indexOfEnUsers++
				okCount++
			}
			if len(phones) > i {
				if phones[i].UserName == "" || phones[i].PhoneNumber == "" {
					fmt.Println(phones[i])
				}

				cnUsers[i].PhoneNumber = strings.ReplaceAll(phones[i].PhoneNumber, "-", "")
				cnUsers[i].UserName = phones[i].UserName
			}
			cnUsers[i].Email = getFakeEmail()
			cnUsers[i].Avatar = getAvatar()
			newUsers = append(newUsers, cnUsers[i])
		}
	}

	for _, u := range newUsers {
		if u.UserName == "" || u.PhoneNumber == "" {
			fmt.Println(u)
		}
	}

	buf, _ = json.Marshal(newUsers)
	os.WriteFile("users.json", buf, os.ModePerm)

	fmt.Println(okCount)
}

func combine2() {
	// process male
	var cnUsers []ShowlistMockUser
	usersCN := "users-cn.json"
	buf, _ := os.ReadFile(usersCN)
	json.Unmarshal(buf, &cnUsers)

	var enUsers []ShowlistMockUser
	usersEN := "users-en.json"
	buf, _ = os.ReadFile(usersEN)
	json.Unmarshal(buf, &enUsers)

	var phones []ShowlistMockUser
	phonesFile := "users-cn-phoneAndUserName.json"
	buf, _ = os.ReadFile(phonesFile)
	json.Unmarshal(buf, &phones)

	var newUsers []ShowlistMockUser

	for i, _ := range cnUsers {
		if i == 10000 {
			break
		}

		cnUsers[i].RealNameEn = enUsers[i].RealNameEn
		if enUsers[i].UserNameEn == "" {
			fmt.Println(enUsers[i])
		}
		cnUsers[i].UserNameEn = enUsers[i].UserNameEn

		cnUsers[i].PhoneNumber = strings.ReplaceAll(phones[i].PhoneNumber, "-", "")
		cnUsers[i].UserName = phones[i].UserName

		cnUsers[i].Email = getFakeEmail()
		cnUsers[i].Avatar = getAvatar()

		newUsers = append(newUsers, cnUsers[i])
	}

	for _, u := range newUsers {
		if u.UserName == "" || u.PhoneNumber == "" {
			fmt.Println(u)
		}
	}

	buf, _ = json.Marshal(newUsers)
	os.WriteFile("users.json", buf, os.ModePerm)
}

var users []ShowlistMockUser

func parseUserInfoCN_Phone() {
	enInfoFiles, _ := utility.GetAllFiles("/Users/kyl/Documents/ihandy/fake_user/fakeUsers-CN-Phone", ".php")
	for _, file := range enInfoFiles {
		reader, err := os.Open(file)
		utility.PanicIfNotNil(err)
		doc, err := goquery.NewDocumentFromReader(reader)
		utility.PanicIfNotNil(err)

		var u ShowlistMockUser
		u.RealName = doc.Find("h3[class=h3title]").First().Text()

		index := 0
		doc.Find("dd[class=col-sm-8]").Each(func(i int, s *goquery.Selection) {
			switch index {
			case 0:
				u.Gender = s.Text()
			case 6:
				u.PhoneNumber = s.Text()
				if u.PhoneNumber == "" {
					fmt.Println("phone is nil")
				}
			case 18:
				u.UserName = s.Text()
				if u.UserName == "" {
					fmt.Println("uname is nil")
				}
			}
			index++
			//fmt.Println(s.Text())
		})

		users = append(users, u)
	}

	buf, _ := json.Marshal(users)
	os.WriteFile("users-cn-phoneAndUserName.json", buf, os.ModePerm)
}

var enUsers = map[string]bool{}

func parseUserInfoEN() {
	enInfoFiles, _ := utility.GetAllFiles("/Users/kyl/Documents/ihandy/fake_user/fakeUsers-US", ".php")
	for _, file := range enInfoFiles {
		reader, err := os.Open(file)
		utility.PanicIfNotNil(err)
		doc, err := goquery.NewDocumentFromReader(reader)
		utility.PanicIfNotNil(err)

		var u ShowlistMockUser
		u.RealNameEn = doc.Find("h3[class=h3title]").First().Text()

		index := 0
		doc.Find("dd[class=col-sm-8]").Each(func(i int, s *goquery.Selection) {
			switch index {
			case 0:
				u.Gender = s.Text()
			case 18:
				u.UserNameEn = s.Text()
				//case 21:
				//	u.Email = s.Text()
			}
			index++
		})

		if enUsers[u.RealNameEn] {
			println(u.RealNameEn)
		} else {
			enUsers[u.RealNameEn] = true
			users = append(users, u)
		}
	}

	buf, _ := json.Marshal(users)
	os.WriteFile("users-en.json", buf, os.ModePerm)
}

func parseUserInfoCN() {
	rootFolder := "/Users/kyl/Documents/ihandy/fake_user/fakeUsers-CN"
	distinctNames = map[string]bool{}
	files, err := utility.GetAllFiles(rootFolder, ".php")
	utility.PanicIfNotNil(err)
	for _, f := range files {
		r, _ := os.Open(f)

		// Load the HTML document
		doc, err := goquery.NewDocumentFromReader(r)
		//doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		var fakeUsers [][]string
		headerLine := true
		doc.Find("table").Each(func(i int, s *goquery.Selection) {
			// For each item found, get the band and title
			var fields []string
			var f func(*html.Node)
			f = func(n *html.Node) {
				if n.Type == html.TextNode {
					fields = append(fields, n.Data)
				}
				if n.FirstChild != nil {
					for c := n.FirstChild; c != nil; c = c.NextSibling {
						f(c)
					}
				}
			}
			for _, n := range s.Find("tr").Nodes {
				if headerLine {
					headerLine = false
					continue
				}
				fields = []string{}
				f(n)
				fakeUsers = append(fakeUsers, fields)
			}
		})

		add(fakeUsers)
	}

	buf, _ := json.Marshal(users)
	os.WriteFile("users-cn.json", buf, os.ModePerm)
	fmt.Println(len(users))
}

type ShowlistMockUser struct {
	UserName    string
	UserNameEn  string
	RealName    string
	RealNameEn  string
	PhoneNumber string
	Email       string
	Image       string
	Avatar      string
	Birthday    string
	Address     string
	Gender      string
	ID          string
}

var distinctNames map[string]bool

func add(originInfos [][]string) {
	for _, info := range originInfos {
		if distinctNames[info[0]] {
			println("already exists ", info[0])
			continue
		}

		distinctNames[info[0]] = true

		users = append(users, ShowlistMockUser{
			RealName: info[0],
			Birthday: info[3],
			Address:  info[5],
			Gender:   info[2],
			ID:       info[1],
		})
	}
}

var avatars []string
var avatarsCount int

func getAvatar() string {
	if len(avatars) == 0 {
		avatars, _ = utility.GetLines("/Users/kyl/Downloads/avatars/avatars.txt")
		avatarsCount = len(avatars)
	}

	return "https://bytepower-server-public.s3.cn-northwest-1.amazonaws.com.cn/avatars/" + avatars[rand.Int31n(int32(avatarsCount))]
}
