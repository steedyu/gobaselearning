package gousualsample

import (
	"encoding/json"
	"fmt"
)

type IT struct {
	Company  string   `json:"company"`
	Subjects []string `json:"subjects"` //二次编码
	IsOk     bool     `json:",string"`  //输出不是bool是string
	Price    float64  `json:"price"`
}

func StructTojsonDemo() {
	//定义一个结构体变量，同时初始化
	s := IT{"itcast", []string{"Go", "C++", "Python", "Test"}, true, 666.666}

	//编码，根据内容生成json文本
	//{"Company":"itcast","Subjects":["Go","C++","Python","Test"],"IsOk":true,"Price":666.666}
	//buf, err := json.Marshal(s)
	buf, err := json.MarshalIndent(s, "", "	") //格式化编码
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	fmt.Println("buf = ", string(buf))
}

func MapTojsonDemo() {

	//创建一个map
	m := make(map[string]interface{}, 4)
	m["company"] = "itcast"
	m["subjects"] = []string{"Go", "C++", "Python", "Test"}
	m["isok"] = true
	m["price"] = 666.666

	//编码成json
	//result, err := json.Marshal(m)
	result, err := json.MarshalIndent(m, "", "	")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Println("result = ", string(result))
}

func JsonToStructDemo() {
	jsonBuf := `
	{
    "company": "itcast",
    "subjects": [
        "Go",
        "C++",
        "Python",
        "Test"
    ],
    "isok": "true",
    "price": 666.666
}`

	var tmp IT                                   //定义一个结构体变量
	err := json.Unmarshal([]byte(jsonBuf), &tmp) //第二个参数要地址传递
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	//fmt.Println("tmp = ", tmp)
	fmt.Printf("tmp = %+v\n", tmp)

	type IT2 struct {
		Subjects []string `json:"subjects"` //二次编码
	}

	var tmp2 IT2
	err = json.Unmarshal([]byte(jsonBuf), &tmp2) //第二个参数要地址传递
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Printf("tmp2 = %+v\n", tmp2)
}

func JsonToMapDemo() {

	jsonBuf := `
	{
    "company": "itcast",
    "subjects": [
        "Go",
        "C++",
        "Python",
        "Test"
    ],
    "isok": "true",
    "price": 666.666
}`

	//创建一个map
	m := make(map[string]interface{}, 4)

	err := json.Unmarshal([]byte(jsonBuf), &m) //第二个参数要地址传递
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Printf("m = %+v\n", m)

	//	var str string
	//	str = string(m["company"]) //err， 无法转换
	//	fmt.Println("str = ", str)

	var str string

	//类型断言, 值，它是value类型
	for key, value := range m {
		//fmt.Printf("%v ============> %v\n", key, value)
		switch data := value.(type) {
		case string:
			str = data
			fmt.Printf("map[%s]的值类型为string, value = %s\n", key, str)
		case bool:
			fmt.Printf("map[%s]的值类型为bool, value = %v\n", key, data)
		case float64:
			fmt.Printf("map[%s]的值类型为float64, value = %f\n", key, data)
		case []string:
			fmt.Printf("map[%s]的值类型为[]string, value = %v\n", key, data)
		case []interface{}:
			fmt.Printf("map[%s]的值类型为[]interface, value = %v\n", key, data)
		}

	}
}

type Message struct {
	Title   string        `json:"title"`
	Content []interface{} `json:"content"`
}

/*
其实这里的主要矛盾，就是 TextContent 本身的类型已经说明它是文本元素了，并不需要一个 Type 字段来显性的地说明它是 text。
换句话说，所有的 TextContent 实例的 Type 字段的值都应该固定是 text，那这个字段压根儿就没有存在的必要不是吗？
但 json 的数据类型系统并不买账，对于 json 来说，object 就是 object，没有“object 是什么类型”这么一说，如果要区分两个 object，
只能通过设置不同的字段或不同的字段值。所以它并不认为 type 字段是可有可无的。
所以思路来了，能不能让 TextContent 没有 Type 字段，但是在序列化成 json 时自动添加上 "type": "text" 呢？如果能实现成这样，那使用时就相当优雅了：
 */

/*
type TextContent struct {
   Type string `json:"type"`
   Text string `json:"text"`
}

type ImageContent struct {
   Type string `json:"type"`
   Src  string `json:"src"`
}

type LinkContent struct {
   Type string `json:"type"`
   Text string `json:"text"`
   Link string `json:"link"`
}

func CommonJsonMarshalSample() {
   msg := Message{
	   Title:   "A apple",
	   Content: []interface{}{
		   TextContent{
			   Type: "text",
			   Text: "There is a apple.",
		   },
		   TextContent{
			   Type: "text",
			   Text: "It is red.",
		   },
		   ImageContent{
			   Type: "image",
			   Src:  "http://example.com/apple.png",
		   },
		   LinkContent{
			   Type: "link",
			   Text: "more info",
			   Link: "http://example.com/more.info.html",
		   },
	   },
   }

   buffer, _ := json.Marshal(msg)
   fmt.Printf("%s\n", buffer)
   // 输出：
   // {"title":"A apple","content":[{"type":"text","text":"There is a apple."},{"type":"text","text":"It is red."},{"type":"image","src":"http://example.com/apple.png"},{"type":"link","text":"more info","link":"http://example.com/more.info.html"}]}
}
*/

type TextContent struct {
	Text string `json:"text"`
}

//type jsonTextContent TextContent
//
//func (c TextContent) MarshalJSON() ([]byte, error) {
//	return json.Marshal(struct {
//		jsonTextContent
//		Type string `json:"type"`
//	}{
//		jsonTextContent: jsonTextContent(c),
//		Type:            "text",
//	})
//}


/*
runtime: goroutine stack exceeds 1000000000-byte limit
fatal error: stack overflow

使用MarshalJSON可以自定义marshall时的动作，但注意无限循环相互调用
 */
//func (c TextContent) MarshalJSON() ([]byte, error) {
//	return json.Marshal(struct {
//		TextContent
//		Type string `json:"type"`
//	}{
//		TextContent: c,
//		Type:        "text",
//	})
//}

type jsonImageContent ImageContent

func (c ImageContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		jsonImageContent
		Type string `json:"type"`
	}{
		jsonImageContent: jsonImageContent(c),
		Type:             "image",
	})
}

type ImageContent struct {
	Src string `json:"src"`
}

//func (c ImageContent) MarshalJSON() ([]byte, error) {
//	return json.Marshal(struct {
//		ImageContent
//		Type string `json:"type"`
//	}{
//		ImageContent: c,
//		Type:         "image",
//	})
//}

type jsonLinkContent LinkContent

func (c LinkContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		jsonLinkContent
		Type string `json:"type"`
	}{
		jsonLinkContent: jsonLinkContent(c),
		Type:            "link",
	})
}

//func (c LinkContent) MarshalJSON() ([]byte, error) {
//	return json.Marshal(struct {
//		LinkContent
//		Type string `json:"type"`
//	}{
//		LinkContent: c,
//		Type:        "link",
//	})
//}

type LinkContent struct {
	Text string `json:"text"`
	Link string `json:"link"`
}

func MarshalAppendAdditionalField() {
	msg := Message{
		Title: "A apple",
		Content: []interface{}{
			TextContent{
				Text: "There is a apple.",
			},
			TextContent{
				Text: "It is red.",
			},
			ImageContent{
				Src: "http://example.com/apple.png",
			},
			LinkContent{
				Text: "more info",
				Link: "http://example.com/more.info.html",
			},
		},
	}

	buffer, _ := json.Marshal(msg)
	fmt.Printf("%s\n", buffer)
	// 输出：
	// {"title":"A apple","content":[{"text":"There is a apple.","type":"text"},{"text":"It is red.","type":"text"},{"src":"http://example.com/apple.png","type":"image"},{"text":"more info","link":"http://example.com/more.info.html","type":"link"}]}
}
