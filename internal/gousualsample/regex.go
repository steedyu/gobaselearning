package gousualsample

import (
	"regexp"
	"fmt"
)


//#1 普通字符组
func RegexCommonStringArray() {

	//普通字符组运算
	fmt.Println(regexp.MatchString("[0123456789]", "5"))//true
	fmt.Println(regexp.MatchString("[0123456789]", "a"))//false
	fmt.Println(regexp.MatchString("[0-9]", "5"))//true
	fmt.Println(regexp.MatchString("[a-z]", "a"))//true
	fmt.Println("---------------------------------")

	//排除字符组
	fmt.Println(regexp.MatchString("[^0-9]", "5"))//false
	fmt.Println(regexp.MatchString("[^a-z]", "a"))//false
	fmt.Println("---------------------------------")

	//字符组运算  golang这个例子不支持
	//fmt.Println(regexp.MatchString("[a-z-[aeiou]]","a"))
	//fmt.Println(regexp.MatchString("[a-z-[aeiou]]","b"))

}

//#2 量词
func RegexMeasureWord() {

	//为什么这里会全部匹配呢，因为这里并没有限定输入字符串的开始和结束的位置（在断言篇你就会知道原因）
	fmt.Println(regexp.MatchString("\\d{3}", "123456789"))//true
	fmt.Println(regexp.MatchString("\\d{3,}", "123456789"))//true
	fmt.Println(regexp.MatchString("\\d{3,5}", "123456789"))//true
	fmt.Println("---------------------------------")

	fmt.Println(regexp.MatchString("^\\d{3}$", "123456789"))//false
	fmt.Println(regexp.MatchString("^\\d{3,}$", "123456789"))//true
	fmt.Println(regexp.MatchString("^\\d{3,5}$", "123456789"))//false
	fmt.Println("---------------------------------")

	fmt.Println(regexp.MatchString("^\\d*$", "123456789"))//true
	fmt.Println(regexp.MatchString("^\\d?$", "123456789"))//false
	fmt.Println(regexp.MatchString("^\\d+$", "123456789"))//true
	fmt.Println("---------------------------------")


	fmt.Println(regexp.MatchString("\\*", "*"))//true
	fmt.Println(regexp.MatchString("\\?", "?"))//true
	fmt.Println(regexp.MatchString("\\+", "+"))//true
	fmt.Println(regexp.MatchString("\\{2,3}", "{2,3}"))//true
	fmt.Println(regexp.MatchString("\\{2,3}", "{4,3}"))//false

}


//C# sample
//一个用于输出匹配子字符串的辅助方法
//        public static void ShowAllMatch(Match match, bool isCapture)
//        {
//            int MatchCount = 0;//匹配次数
//            while (true)
//            {
//                if (match.Success)
//                {
//                    MatchCount++;
//                    Console.WriteLine("第{0}个匹配的字符串为:{1}", MatchCount.ToString(), match.Value);
//                    int GroupIndex = 0;
//                    /*
//                     * 对于任意一个Match对象，其Gruops集合中至少有一个元素，下标为0。
//                     * 若Match对象能成功匹配子字符串，则下标0的位置默认存储整个匹配的字符串。若不成功则为空。
//                     * 若通过下标寻找不存在的项，则只返回空字符串，不会报错。
//                     * 其余下标对应正则表达式中括号出现的位置。如第一个Gruops[1],对应匹配第一个括号的分组。
//                     */
//foreach (Group group in match.Groups)
//{
//Console.WriteLine("分组下标:{0},分组的值为:{1}", GroupIndex.ToString(), group.Value);
///*
// * 当在括号分组后加上量词时(如'(\d){m,n}'这种形式)，可通过Capture类获得每一次捕获的字符串。
// * 如，若匹配了字符串"2015"，则Capture会分别捕获，"2","0","1","5",group.Value的值为"5"
// * 此时，group.Value仅获取最后捕获的子字符串。
// */
//if (isCapture)
//{
//int CaptureIndex = 0;
//foreach (Capture capture in group.Captures)
//{
//Console.WriteLine("分组下标:{0},第{1}个捕获的子字符串为:{2}", GroupIndex.ToString(), CaptureIndex.ToString(), capture.Value);
//CaptureIndex++;
//}
//}
//GroupIndex++;
//}
//match = match.NextMatch();//获取从该匹配位置结束之后的下一个匹配对象。
//Console.WriteLine();
//}
//else
//{
//if (MatchCount == 0)
//Console.WriteLine("没有匹配项！");
//else
//Console.WriteLine("匹配结束！");
//break;
//}
//}
//}

//Console.WriteLine("---捕获分组示例1---");
//string InputB = "2015-06-15-20-30-2014-02-28-01-015-2015-08-09-00";
//string RegexStrB = @"(\d{4})-(\d{2})-\d{2}";//匹配一个日期，格式为YYYY-MM-DD。再通过分组获取匹配字符串中的年份和月份。
//Regex RegexB = new Regex(RegexStrB);
//Match MatchB = RegexB.Match(InputB);
//ShowAllMatch(MatchB, false);
//Console.WriteLine();
//
//
///*
// *  Groups[0] 是整个表达式匹配的内容
//    Groups[1] 是第一个捕获组捕获的内容
//    以此类推。。。
//
//    表达式中的捕获组只匹配成功一次的情况比较简单，就不作说明了
//    当表达式中的捕获组可以匹配成功多次时，Group只保留最后一次匹配结果，而匹配过程的中间值，就被保留成CaptureCollection了
// */
//Console.WriteLine("---捕获分组示例2---");
//string InputC = "2015-06-15-20-30-2014-02-28-01-015-2015-08-09-00";
//string RegexStrC = @"(\d){4}-(\d){2}-\d{2}";//匹配一个日期，格式为YYYY-MM-DD。再通过分组获取匹配字符串中的年份和月份。
//Regex RegexC = new Regex(RegexStrC);
//Match MatchC = RegexC.Match(InputC);
//ShowAllMatch(MatchC, true);
//Console.WriteLine();
//
//Console.WriteLine("---捕获分组示例3---");
//string InputD = "2015-06-15-20-30-2014-02-28-01-015-2015-08-09-00";
///*
// * 匹配一个日期，格式为YYYY-MM-DD。再通过分组获取匹配字符串中的年份和月份。
// * 若括号之间有嵌套，则Group的下标先算外层括号，再算内层。
// */
//string RegexStrD = @"((\d){4}-(\d){2})-(\d){2}";
//Regex RegexD = new Regex(RegexStrD);
//Match MatchD = RegexD.Match(InputD);
//ShowAllMatch(MatchD, true);
//Console.WriteLine();


//#3 括号
func RegexBracket() {
	//false error parsing regexp: invalid nested repetition operator: `{2}+`
	//fmt.Println(regexp.MatchString("^\\d{2}+$", "12"))//true
	fmt.Println(regexp.MatchString("^(\\d{2})+$", "12"))//true


	var InputB string = "2015-06-15-20-30-2014-02-28-01-015-2015-08-09-00"
	var regexStrB string = "(\\d{4})-(\\d{2})-\\d{2}"
	showallmatch(regexStrB,InputB)
	fmt.Println("-----------------InputB End--------------------------------------")


	var InputC string= "2015-06-15-20-30-2014-02-28-01-015-2015-08-09-00"
	var RegexStrC string = "(\\d){4}-(\\d){2}-\\d{2}";//匹配一个日期，格式为YYYY-MM-DD。再通过分组获取匹配字符串中的年份和月份。
	showallmatch(RegexStrC,InputC)
	fmt.Println("-----------------InputC End--------------------------------------")


	var InputD string= "2015-06-15-20-30-2014-02-28-01-015-2015-08-09-00"
	var RegexStrD string = "((\\d){4}-(\\d){2})-(\\d){2}"//匹配一个日期，格式为YYYY-MM-DD。再通过分组获取匹配字符串中的年份和月份。
	showallmatch(RegexStrD,InputD)
	fmt.Println("-----------------InputD End--------------------------------------")

}

func showallmatch(matchstr string, s string) {

	var match = regexp.MustCompile(matchstr)

	//FindStringSubmatch returns a slice of strings holding the text of the
	// leftmost match of the regular expression in s and the matches
	//match.FindStringSubmatch()

	//FindStringSubmatchIndex returns a slice holding the index pairs
	// identifying the leftmost match of the regular expression in s
	//每组开始下标及结束下标的后一位
	//fmt.Println(match.FindStringSubmatchIndex(s))

	strings := match.FindAllStringSubmatch(s, -1)
	fmt.Println(strings)
	for i, match := range strings {
		fmt.Printf("Match %v: ", i)
		for j, group := range match {
			fmt.Printf("Group %v: %v ", j, group)
		}
		fmt.Println()
	}

	fmt.Println(match.FindAllStringSubmatchIndex(s, -1))

}

//#4 反相引用
func RegexReverseReference() {


	//panic: regexp: CompilePOSIX(`([a-z])\1`): error parsing regexp: invalid escape sequence: `\1`
	//golang 正则表达式，在这块不支持

	//regex := regexp.MustCompilePOSIX("([a-z])\\1")
	//fmt.Println(regex.FindAllStringSubmatch("book", -1))
	//fmt.Println(regex.MatchString("book"))//true
	//fmt.Println(regex.MatchString("sleep"))//true
	//fmt.Println(regex.MatchString("where"))//false
}

var uidmatch = regexp.MustCompile("^\\d{16}$")

func IsAllow(uid string) (ok bool) {

	if uidmatch.MatchString(uid) {
		ok = true
	}

	return
}
//"(?=.*[0-9].*)(?=.*[A-Z].*)(?=.*[a-z].*).{12,20}"
var (
	bidformatregmatch = regexp.MustCompile("^[a-zA-z0-9]{32}$")
	onlynummatch = regexp.MustCompile("^[0-9]{32}$")
	onlylettermatch = regexp.MustCompile("^[a-zA-z]{32}$")
)

func IsBid(bid string) (result bool) {
	if bidformatregmatch.MatchString(bid) && !onlylettermatch.MatchString(bid) && !onlynummatch.MatchString(bid) {
		result = true
	}

	return
}

func Xx() {
	//\@=   (?= 顺序环视
	//查找后面是sql的my： /my\(sql\)\@=
	reg := regexp.MustCompile("(.*[0-9].*)@(.*[A-Z].*)@(.*[a-z].*)@.{8}")

	reg1 := regexp.MustCompile(`my(?=sql)`)
	fmt.Println(reg1.MatchString("mysql"))
	//reg, _ := regexp.Compile("(.*[0-9].*)(.*[a-zA-Z].*){8}")
	fmt.Println(reg.MatchString("111"))
	fmt.Println(reg.MatchString("11z1aa1"))
	fmt.Println(reg.MatchString("z11z1Aa1"))
}