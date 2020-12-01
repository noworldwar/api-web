package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/noworldwar/api-web/internal/model"
	"github.com/spf13/viper"
)

func GenerateDoc() {
	InitLangConfig()

	// title,version
	InitAPIdoc("zh")
	InitAPIdoc("cn")
	InitAPIdoc("en")

	routerPage := ReadRouterFile()

	// 取得 router 中的 Tag
	tags := GetAPIRoutersFromRouterFile(routerPage)

	// 在 /internal/api 中讀取所有檔案
	files, err := ioutil.ReadDir("./internal/api")
	if err != nil {
		log.Fatal(err)
	}

	// 把所有的 function 接成一個字串
	allFunc := GetAllFilesContentFromAPI(files)

	// 沒有語系切換的通用 Tag 結構
	GetTagModel("zh", tags, allFunc)
	GetTagModel("cn", tags, allFunc)
	GetTagModel("en", tags, allFunc)

	// 插入轉換語系的內容
	InsertLangContent("zh", model.APIdocZh.Tag)
	InsertLangContent("cn", model.APIdocCn.Tag)
	InsertLangContent("en", model.APIdocEn.Tag)
}

func InitAPIdoc(lang string) {
	byteValue, err := ioutil.ReadFile("./doc/lang/menu_" + lang + ".json")
	if err != nil {
		fmt.Printf("Read %s.json File Error:%s\r\n", lang, err)
	}

	switch lang {
	case "zh":
		err = json.Unmarshal(byteValue, &model.APIdocZh)
	case "cn":
		err = json.Unmarshal(byteValue, &model.APIdocCn)
	case "en":
		err = json.Unmarshal(byteValue, &model.APIdocEn)
	}

	if err != nil {
		fmt.Printf("Parsing APIdoc%s Error:%s\r\n", lang, err)
	}
}

// read router.go
func ReadRouterFile() []byte {
	routerPage, err := ioutil.ReadFile("./internal/app/router.go")
	if err != nil {
		log.Fatal(err)
	}
	return routerPage
}

// get api routers from router.go
func GetAPIRoutersFromRouterFile(routerPage []byte) []string {
	regexPattern := "\\ @Api(.*?)\\ @end"
	getTags := regexp.MustCompile(regexPattern)
	tagsCont := strings.Replace(string(routerPage), "\n", "", -1)
	tagsCont = strings.Replace(tagsCont, "\r", "", -1)
	tags := getTags.FindAllString(tagsCont, -1)

	if len(tags) > 0 {
		return tags
	}

	return nil
}

// connect all files content
func GetAllFilesContentFromAPI(files []os.FileInfo) string {
	allFunc := ""
	for _, f := range files {
		absPath, _ := filepath.Abs("./internal/api/" + f.Name())
		ff, err := ioutil.ReadFile(absPath)
		if err != nil {
			fmt.Println("read fail", err)
		}
		allFunc = allFunc + string(ff) + "\n"
	}

	return allFunc
}

func GetAllStringMatch(rawContent, regexPattern string) [][]string {
	getContent := regexp.MustCompile(regexPattern)
	contents := getContent.FindAllStringSubmatch(rawContent, -1)
	if len(contents) > 0 {
		return contents
	}

	return nil
}

func GetTagModel(lang string, tags []string, allFunc string) {
	docTag := []model.DocTag{}
	// compile single api
	for _, tagData := range tags {
		regexPattern := "\\ @Api:(.*?)\\@"
		getTagsContent := regexp.MustCompile(regexPattern)
		tagsContent := getTagsContent.FindStringSubmatch(tagData)
		tagArr := strings.Split(tagsContent[1], " ")
		methods := GetAllStringMatch(tagData, "r\\.(POST|GET|PUT|DELETE)")
		paths := GetAllStringMatch(tagData, "\"(.*?)\"")
		funNames := GetAllStringMatch(tagData, "api\\.(.*?)\\)")
		docFunc := GetDocFunc(funNames, allFunc, methods, paths, tagArr[0])
		docTag = append(docTag, model.DocTag{
			Name:  tagArr[0],
			Funcs: docFunc,
		})
	}

	switch lang {
	case "zh":
		model.APIdocZh.Tag = docTag
	case "cn":
		model.APIdocCn.Tag = docTag
	case "en":
		model.APIdocEn.Tag = docTag
	}
}

func GetDocFunc(funNames [][]string, allFunc string, methods, paths [][]string, tag string) []model.DocFunc {
	docFunc := []model.DocFunc{}
	// compile single function
	for index, funName := range funNames {
		funcContent := GetFunctionContent(allFunc, funName[1])
		// 取得Request
		rqContent := GetAllStringMatch(funcContent, "@Request (.*?) @")
		// 取得area
		areaOption := GetAreaOption(rqContent)
		request := GetRequestDocBody(areaOption, rqContent, tag, funName[1])
		// 取得Response
		regexPattern := "@Response (.*?) @"
		getRsContent := regexp.MustCompile(regexPattern)
		rsContent := getRsContent.FindAllStringSubmatch(funcContent, -1)
		response := GetResponseDocBody(rsContent)
		docFunc = append(docFunc, model.DocFunc{
			FuncName:    funName[1],
			Method:      methods[index][1],
			Path:        paths[index][1],
			Description: "",
			Requests:    request,
			Responses:   response,
		})
	}

	return docFunc
}

func GetFunctionContent(allFunc, funName string) string {
	regexPattern := "@Func " + funName + "(.*?)func " + funName
	getFuncContent := regexp.MustCompile(regexPattern)
	funcCont := strings.Replace(allFunc, "\n", "", -1)
	funcCont = strings.Replace(funcCont, "\r", "", -1)
	funcContent := getFuncContent.FindString(funcCont)

	if len(funcContent) > 0 {
		return funcContent
	}

	return ""
}

func GetAreaOption(rqContent [][]string) []string {
	areaOption := []string{}
	for _, rq := range rqContent {
		rqArr := strings.Split(rq[1], "{")
		area := strings.Trim(rqArr[0], " ")
		areaOption = append(areaOption, area)
	}

	areaOption = SliceUniqMap(areaOption)
	if len(areaOption) > 0 {
		return areaOption
	}

	return nil
}

func GetRequestDocBody(areaOption []string, rqContent [][]string, tag, funcName string) []model.DocRequest {
	request := []model.DocRequest{}
	for _, area := range areaOption {
		requestDocBody := []model.DocBody{}
		for _, rq := range rqContent {
			rqArr := strings.Split(rq[1], "{")
			rqArea := strings.Trim(rqArr[0], " ")
			if area == rqArea {
				trimCont := strings.TrimLeft(rqArr[1], "{")
				trimCont = strings.TrimRight(trimCont, "}/")
				descArr := strings.Split(trimCont, ",")
				required, _ := strconv.ParseBool(descArr[2])
				requestDocBody = append(requestDocBody, model.DocBody{
					Name:        descArr[0],
					Type:        descArr[1],
					Required:    required,
					Description: "",
				})
			}
		}
		request = append(request, model.DocRequest{
			Area: area,
			Body: requestDocBody,
		})
	}
	return request
}

func GetResponseDocBody(rsContent [][]string) []model.DocResponse {
	response := []model.DocResponse{}
	for _, rs := range rsContent {
		rsArr := strings.Split(rs[1], " ")
		response = append(response, model.DocResponse{
			StateCode:   rsArr[0],
			Description: "",
			Message:     rsArr[1],
		})
	}
	return response
}

func SliceUniqMap(s []string) []string {
	seen := make(map[string]struct{}, len(s))
	j := 0
	for _, v := range s {
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		s[j] = v
		j++
	}
	return s[:j]
}

func InsertLangContent(lang string, apiDoc []model.DocTag) {
	for tagIdx, tag := range apiDoc {
		for funcIdx := range tag.Funcs {
			apiDoc[tagIdx].Funcs[funcIdx].Description = GetLangDocString(lang, "FuncName."+tag.Funcs[funcIdx].FuncName)
			apiDoc[tagIdx].Funcs[funcIdx].Explanation = GetLangDocString(lang, "Explanation."+tag.Funcs[funcIdx].FuncName)
			for reqIdx, req := range tag.Funcs[funcIdx].Requests {
				area := req.Area
				for bodyIdx, body := range req.Body {
					apiDoc[tagIdx].Funcs[funcIdx].Requests[reqIdx].Body[bodyIdx].Description = GetLangDocString(lang, tag.Funcs[funcIdx].FuncName+".Requests."+area+"."+body.Name)
				}
			}
			for resIdx, res := range tag.Funcs[funcIdx].Responses {
				apiDoc[tagIdx].Funcs[funcIdx].Responses[resIdx].Description = GetLangDocString(lang, tag.Funcs[funcIdx].FuncName+".Responses."+res.StateCode)
			}
		}
		apiDoc[tagIdx].Name = GetLangDocString(lang, "Tag."+tag.Name)
	}
}

func GetLangDocString(lang, path string) string {
	switch lang {
	case "zh":
		return model.Zh.GetString(path)
	case "cn":
		return model.Cn.GetString(path)
	case "en":
		return model.En.GetString(path)
	default:
		return ""
	}
}

func InitLangConfig() {
	model.Zh = viper.New()
	model.Zh.SetConfigName("content_zh")
	model.Zh.SetConfigType("json")
	model.Zh.AddConfigPath("./doc/lang")

	zhErr := model.Zh.ReadInConfig() // Find and read the config file
	if zhErr != nil {                // Handle errors reading the config file
		log.Fatalln("Fatal error content_zh.json:", zhErr)
	}

	model.Cn = viper.New()
	model.Cn.SetConfigName("content_cn")
	model.Cn.SetConfigType("json")
	model.Cn.AddConfigPath("./doc/lang")

	cnErr := model.Cn.ReadInConfig() // Find and read the config file
	if cnErr != nil {                // Handle errors reading the config file
		log.Fatalln("Fatal error content_zh.json:", zhErr)
	}

	model.En = viper.New()
	model.En.SetConfigName("content_en")
	model.En.SetConfigType("json")
	model.En.AddConfigPath("./doc/lang")

	enErr := model.En.ReadInConfig() // Find and read the config file
	if enErr != nil {                // Handle errors reading the config file
		log.Fatalln("Fatal error content_zh.json:", zhErr)
	}
}
