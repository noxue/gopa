爬虫库，指定规则即可抓取指定页面，用法非常简单。

## 用法

`go get github.com/noxue/gopa`

## 案例，采集起点中文网完本小说列表

```go
package main

import (
	"fmt"
	"github.com/noxue/gopa"
)

func main() {
	err := func() (err error) {
		defer func() {
			if e := recover(); e != nil {
				err = e.(error)
			}
		}()

		/**
		规则说明
		{
			"all":true, 是否匹配全部结果，true表示匹配全部，用于获取列表。false表示之获取一个，用于采集文章内容等
			"rules":[
				// name指定采集的结果名称
				// selector指定要采集内容的正则表达式，注意第一个( )匹配的 则是最终的内容
				// do表示对结果进行处理，比如添加前缀，后缀，拼接等。一个用逗号分隔的字符串，逗号后面字符串中包含逗号前面指定的字符串用匹配到的内容替换
				// replace数组指定多个替换规则数组，selector是满足条件的正则，text则是用于替换的内容
				{"name":"name", "selector":"<td><a class=\"name\".*?>(.*?)<"},
				{"name":"author", "selector":"<a class=\"author\".*?>(.*?)<"},
				{"name":"date", "selector":"<td class=\"date\">(.*?)<"},
				{"name":"url", "selector":"<td><a class=\"name\".*?href=\"(.*?)\"", "do":"#,https:#","replace":[{"selector":"qi","text":"bai"},{"selector":"dian","text":"du"}]}
			]
		}
		 */
		rules := `
		{
			"all":true,
			"rules":[
				{"name":"name", "selector":"<td><a class=\"name\".*?>(.*?)<"},
				{"name":"author", "selector":"<a class=\"author\".*?>(.*?)<"},
				{"name":"date", "selector":"<td class=\"date\">(.*?)<"},
				{"name":"url", "selector":"<td><a class=\"name\".*?href=\"(.*?)\"", "do":"#,https:#","replace":[{"selector":"qi","text":"bai"},{"selector":"dian","text":"du"}]}
			]
		}`

		pa := gopa.NewPa()
		data := pa.
			Get("https://www.qidian.com/finish?action=hidden&orderId=&page=1&style=2&pageSize=50&siteid=1&pubflag=0&hiddenField=2").
			Rules(rules)

		// 还可以直接调用Download方法保存到指定文件，比如保存图片，下载文件等
		// data.Download("c:/noxue.com.jpg")

		// 获取所有结果
		//data.All()

		// 获取第一个结果
		//data.One()


		fmt.Println(data)

		// 加入自己的逻辑，比如下载图片并保存到七牛云，然后结果替换成七牛云的地址，方便存储
		data.Do(func(data *gopa.Data) *gopa.Data {
			// 对每个结果循环处理
			// Data中保存的是采集到的结果数组，以key=>value形式保存
			for _, v := range data.Data {
				// 在这里可以对地址做任何操作，如果是图片你可以保存到七牛，然后修改原来的数据，
				// 假设下面是经过操作后变成另一个地址，然后修改采集到的结果
				v["url"] = v["url"]+"#test11111"
			}
			return data
		})

		fmt.Println(data)
		return
	}()

	if err != nil {
		fmt.Println(err)
	}
}
```

* 运行结果

```
[
	[0] =>(
		[name] => 带着仓库到大明,
		[author] => 迪巴拉爵士,
		[date] => 2018-11-06,
		[url] => https://book.baidu.com/info/1004185492,
	),

	[1] =>(
		[name] => 奶爸的文艺人生,
		[author] => 寒门,
		[date] => 2018-06-29,
		[url] => https://book.baidu.com/info/1009915605,
	),

	[2] =>(
		[name] => 重生之财源滚滚,
		[author] => 老鹰吃小鸡,
		[date] => 2018-05-07,
		[url] => https://book.baidu.com/info/1003580078,
	),

	[3] =>(
		[name] => 全能游戏设计师,
		[author] => 青衫取醉,
		[date] => 2018-09-20,
		[url] => https://book.baidu.com/info/1010377389,
	),

	[4] =>(
		[name] => 逆流纯真年代,
		[author] => 人间武库,
		[date] => 2018-09-10,
		[url] => https://book.baidu.com/info/1009398284,
	),

	[5] =>(
		[url] => https://book.baidu.com/info/3439785,
		[name] => 修真四万年,
		[author] => 卧牛真人,
		[date] => 2018-08-31,
	),

	[6] =>(
		[author] => 文抄公,
		[date] => 2018-11-30,
		[url] => https://book.baidu.com/info/1009942824,
		[name] => 逍遥梦路,
	),

	[7] =>(
		[name] => 大魏宫廷,
		[author] => 贱宗首席弟子,
		[date] => 2018-08-06,
		[url] => https://book.baidu.com/info/3662715,
	),

	[8] =>(
		[name] => 老衲要还俗,
		[author] => 一梦黄粱,
		[date] => 2018-09-29,
		[url] => https://book.baidu.com/info/1005405437,
	),

	[9] =>(
		[name] => 民国之文豪崛起,
		[author] => 王梓钧,
		[date] => 2018-07-30,
		[url] => https://book.baidu.com/info/1005064061,
	),

	[10] =>(
		[date] => 2018-04-20,
		[url] => https://book.baidu.com/info/3242304,
		[name] => 异常生物见闻录,
		[author] => 远瞳,
	),

	[11] =>(
		[url] => https://book.baidu.com/info/1010463703,
		[name] => 系统的黑科技网吧,
		[author] => 逆水之叶,
		[date] => 2018-10-08,
	),

	[12] =>(
		[author] => 我最白,
		[date] => 2018-11-19,
		[url] => https://book.baidu.com/info/1010438082,
		[name] => 文娱万岁,
	),

	[13] =>(
		[date] => 2018-10-22,
		[url] => https://book.baidu.com/info/1003504656,
		[name] => 史上最强店主,
		[author] => 南极烈日,
	),

	[14] =>(
		[name] => 大唐贞观第一纨绔,
		[author] => 危险的世界,
		[date] => 昨日14:01,
		[url] => https://book.baidu.com/info/1004984707,
	),

	[15] =>(
		[name] => 凡人修仙传,
		[author] => 忘语,
		[date] => 2016-01-05,
		[url] => https://book.baidu.com/info/107580,
	),

	[16] =>(
		[name] => 绝命手游,
		[author] => 奥比椰,
		[date] => 2018-08-29,
		[url] => https://book.baidu.com/info/1010318085,
	),

	[17] =>(
		[name] => 永恒武道,
		[author] => 月中阴,
		[date] => 2018-11-01,
		[url] => https://book.baidu.com/info/1009329519,
	),

	[18] =>(
		[name] => 走进修仙,
		[author] => 吾道长不孤,
		[date] => 2018-12-01,
		[url] => https://book.baidu.com/info/3406500,
	),

	[19] =>(
		[name] => 位面电梯,
		[author] => 千翠百恋,
		[date] => 2018-11-02,
		[url] => https://book.baidu.com/info/1003353824,
	),

	[20] =>(
		[url] => https://book.baidu.com/info/1209977,
		[name] => 斗破苍穹,
		[author] => 天蚕土豆,
		[date] => 2018-09-19,
	),

	[21] =>(
		[name] => 英雄联盟：上帝之手,
		[author] => 三千勿忘尽,
		[date] => 2018-07-04,
		[url] => https://book.baidu.com/info/1003900126,
	),

	[22] =>(
		[author] => 无言不信,
		[date] => 2018-11-26,
		[url] => https://book.baidu.com/info/1004949546,
		[name] => 盛唐剑圣,
	),

	[23] =>(
		[name] => 神门,
		[author] => 薪意,
		[date] => 2018-07-16,
		[url] => https://book.baidu.com/info/3600493,
	),

	[24] =>(
		[date] => 2018-11-20,
		[url] => https://book.baidu.com/info/1005313052,
		[name] => 请回答火影,
		[author] => 蒙着面的Sama,
	),

	[25] =>(
		[author] => 懵比的小提莫,
		[date] => 昨日04:49,
		[url] => https://book.baidu.com/info/1009953965,
		[name] => 梦幻西游大主播,
	),

	[26] =>(
		[date] => 2018-11-05,
		[url] => https://book.baidu.com/info/1004939100,
		[name] => 杀神永生,
		[author] => 恐怖的阿肥,
	),

	[27] =>(
		[url] => https://book.baidu.com/info/1010846649,
		[name] => 助鬼为乐系统,
		[author] => 左断手,
		[date] => 2018-11-02,
	),

	[28] =>(
		[author] => 小伈,
		[date] => 2018-05-04,
		[url] => https://book.baidu.com/info/1010795768,
		[name] => 德猎,
	),

	[29] =>(
		[name] => 垂钓诸天,
		[author] => 道在不可鸣,
		[date] => 2018-11-19,
		[url] => https://book.baidu.com/info/1004893588,
	),

	[30] =>(
		[date] => 2018-05-24,
		[url] => https://book.baidu.com/info/1004175804,
		[name] => 逍遥小书生,
		[author] => 荣小荣,
	),

	[31] =>(
		[name] => 活在诸天,
		[author] => 你好再见见,
		[date] => 2018-10-10,
		[url] => https://book.baidu.com/info/1009607977,
	),

	[32] =>(
		[name] => 美国牧场的小生活,
		[author] => 醛石,
		[date] => 2018-11-16,
		[url] => https://book.baidu.com/info/1011337448,
	),

	[33] =>(
		[date] => 2018-10-10,
		[url] => https://book.baidu.com/info/3660723,
		[name] => 抗日之特战兵王,
		[author] => 寂寞剑客,
	),

	[34] =>(
		[name] => 斗罗大陆,
		[author] => 唐家三少,
		[date] => 2009-12-24,
		[url] => https://book.baidu.com/info/1115277,
	),

	[35] =>(
		[name] => 都市至强者降临,
		[author] => 极地风刃,
		[date] => 2018-09-26,
		[url] => https://book.baidu.com/info/1009958915,
	),

	[36] =>(
		[name] => 超级潇洒人生,
		[author] => 胖达福,
		[date] => 2018-05-10,
		[url] => https://book.baidu.com/info/1009845790,
	),

	[37] =>(
		[name] => 超级神掠夺,
		[author] => 奇燃,
		[date] => 2018-11-12,
		[url] => https://book.baidu.com/info/1010695687,
	),

	[38] =>(
		[author] => 墨渊九砚,
		[date] => 2018-05-31,
		[url] => https://book.baidu.com/info/1005115417,
		[name] => 火影之最强卡卡西,
	),

	[39] =>(
		[name] => 壹号卫,
		[author] => 葛洛夫街兄弟,
		[date] => 2018-11-15,
		[url] => https://book.baidu.com/info/1010863551,
	),

	[40] =>(
		[name] => 这个天国不太平,
		[author] => 三江口水,
		[date] => 2018-11-30,
		[url] => https://book.baidu.com/info/1003640751,
	),

	[41] =>(
		[name] => 末世之宠物为王,
		[author] => 六枭,
		[date] => 2018-10-23,
		[url] => https://book.baidu.com/info/1010314675,
	),

	[42] =>(
		[name] => 全职高手,
		[author] => 蝴蝶蓝,
		[date] => 2014-04-30,
		[url] => https://book.baidu.com/info/1887208,
	),

	[43] =>(
		[name] => 奇遇无限,
		[author] => 龙鳞道V,
		[date] => 2018-10-18,
		[url] => https://book.baidu.com/info/1010519813,
	),

	[44] =>(
		[date] => 2018-10-31,
		[url] => https://book.baidu.com/info/1010957891,
		[name] => 浴血兵锋,
		[author] => 丑牛1985,
	),

	[45] =>(
		[url] => https://book.baidu.com/info/1010766239,
		[name] => 大唐农圣,
		[author] => 爱吃鱼的胖子,
		[date] => 2018-09-30,
	),

	[46] =>(
		[author] => 昨夜大雨,
		[date] => 2018-08-20,
		[url] => https://book.baidu.com/info/1010256698,
		[name] => 位面之纨绔生涯,
	),

	[47] =>(
		[name] => 穷鬼的上下两千年,
		[author] => 非玩家角色,
		[date] => 2018-10-25,
		[url] => https://book.baidu.com/info/1010696881,
	),

	[48] =>(
		[date] => 2018-09-03,
		[url] => https://book.baidu.com/info/1004857359,
		[name] => 当个法师闹革命,
		[author] => 尹四,
	),

	[49] =>(
		[url] => https://book.baidu.com/info/1010690206,
		[name] => 我的舢舨能升级,
		[author] => 二哥的小号,
		[date] => 2018-11-02,
	),
]
[
	[0] =>(
		[author] => 迪巴拉爵士,
		[date] => 2018-11-06,
		[url] => https://book.baidu.com/info/1004185492#test11111,
		[name] => 带着仓库到大明,
	),

	[1] =>(
		[name] => 奶爸的文艺人生,
		[author] => 寒门,
		[date] => 2018-06-29,
		[url] => https://book.baidu.com/info/1009915605#test11111,
	),

	[2] =>(
		[date] => 2018-05-07,
		[url] => https://book.baidu.com/info/1003580078#test11111,
		[name] => 重生之财源滚滚,
		[author] => 老鹰吃小鸡,
	),

	[3] =>(
		[name] => 全能游戏设计师,
		[author] => 青衫取醉,
		[date] => 2018-09-20,
		[url] => https://book.baidu.com/info/1010377389#test11111,
	),

	[4] =>(
		[name] => 逆流纯真年代,
		[author] => 人间武库,
		[date] => 2018-09-10,
		[url] => https://book.baidu.com/info/1009398284#test11111,
	),

	[5] =>(
		[name] => 修真四万年,
		[author] => 卧牛真人,
		[date] => 2018-08-31,
		[url] => https://book.baidu.com/info/3439785#test11111,
	),

	[6] =>(
		[author] => 文抄公,
		[date] => 2018-11-30,
		[url] => https://book.baidu.com/info/1009942824#test11111,
		[name] => 逍遥梦路,
	),

	[7] =>(
		[date] => 2018-08-06,
		[url] => https://book.baidu.com/info/3662715#test11111,
		[name] => 大魏宫廷,
		[author] => 贱宗首席弟子,
	),

	[8] =>(
		[url] => https://book.baidu.com/info/1005405437#test11111,
		[name] => 老衲要还俗,
		[author] => 一梦黄粱,
		[date] => 2018-09-29,
	),

	[9] =>(
		[author] => 王梓钧,
		[date] => 2018-07-30,
		[url] => https://book.baidu.com/info/1005064061#test11111,
		[name] => 民国之文豪崛起,
	),

	[10] =>(
		[name] => 异常生物见闻录,
		[author] => 远瞳,
		[date] => 2018-04-20,
		[url] => https://book.baidu.com/info/3242304#test11111,
	),

	[11] =>(
		[date] => 2018-10-08,
		[url] => https://book.baidu.com/info/1010463703#test11111,
		[name] => 系统的黑科技网吧,
		[author] => 逆水之叶,
	),

	[12] =>(
		[name] => 文娱万岁,
		[author] => 我最白,
		[date] => 2018-11-19,
		[url] => https://book.baidu.com/info/1010438082#test11111,
	),

	[13] =>(
		[name] => 史上最强店主,
		[author] => 南极烈日,
		[date] => 2018-10-22,
		[url] => https://book.baidu.com/info/1003504656#test11111,
	),

	[14] =>(
		[name] => 大唐贞观第一纨绔,
		[author] => 危险的世界,
		[date] => 昨日14:01,
		[url] => https://book.baidu.com/info/1004984707#test11111,
	),

	[15] =>(
		[url] => https://book.baidu.com/info/107580#test11111,
		[name] => 凡人修仙传,
		[author] => 忘语,
		[date] => 2016-01-05,
	),

	[16] =>(
		[author] => 奥比椰,
		[date] => 2018-08-29,
		[url] => https://book.baidu.com/info/1010318085#test11111,
		[name] => 绝命手游,
	),

	[17] =>(
		[date] => 2018-11-01,
		[url] => https://book.baidu.com/info/1009329519#test11111,
		[name] => 永恒武道,
		[author] => 月中阴,
	),

	[18] =>(
		[name] => 走进修仙,
		[author] => 吾道长不孤,
		[date] => 2018-12-01,
		[url] => https://book.baidu.com/info/3406500#test11111,
	),

	[19] =>(
		[name] => 位面电梯,
		[author] => 千翠百恋,
		[date] => 2018-11-02,
		[url] => https://book.baidu.com/info/1003353824#test11111,
	),

	[20] =>(
		[name] => 斗破苍穹,
		[author] => 天蚕土豆,
		[date] => 2018-09-19,
		[url] => https://book.baidu.com/info/1209977#test11111,
	),

	[21] =>(
		[name] => 英雄联盟：上帝之手,
		[author] => 三千勿忘尽,
		[date] => 2018-07-04,
		[url] => https://book.baidu.com/info/1003900126#test11111,
	),

	[22] =>(
		[name] => 盛唐剑圣,
		[author] => 无言不信,
		[date] => 2018-11-26,
		[url] => https://book.baidu.com/info/1004949546#test11111,
	),

	[23] =>(
		[name] => 神门,
		[author] => 薪意,
		[date] => 2018-07-16,
		[url] => https://book.baidu.com/info/3600493#test11111,
	),

	[24] =>(
		[author] => 蒙着面的Sama,
		[date] => 2018-11-20,
		[url] => https://book.baidu.com/info/1005313052#test11111,
		[name] => 请回答火影,
	),

	[25] =>(
		[date] => 昨日04:49,
		[url] => https://book.baidu.com/info/1009953965#test11111,
		[name] => 梦幻西游大主播,
		[author] => 懵比的小提莫,
	),

	[26] =>(
		[name] => 杀神永生,
		[author] => 恐怖的阿肥,
		[date] => 2018-11-05,
		[url] => https://book.baidu.com/info/1004939100#test11111,
	),

	[27] =>(
		[name] => 助鬼为乐系统,
		[author] => 左断手,
		[date] => 2018-11-02,
		[url] => https://book.baidu.com/info/1010846649#test11111,
	),

	[28] =>(
		[name] => 德猎,
		[author] => 小伈,
		[date] => 2018-05-04,
		[url] => https://book.baidu.com/info/1010795768#test11111,
	),

	[29] =>(
		[url] => https://book.baidu.com/info/1004893588#test11111,
		[name] => 垂钓诸天,
		[author] => 道在不可鸣,
		[date] => 2018-11-19,
	),

	[30] =>(
		[author] => 荣小荣,
		[date] => 2018-05-24,
		[url] => https://book.baidu.com/info/1004175804#test11111,
		[name] => 逍遥小书生,
	),

	[31] =>(
		[name] => 活在诸天,
		[author] => 你好再见见,
		[date] => 2018-10-10,
		[url] => https://book.baidu.com/info/1009607977#test11111,
	),

	[32] =>(
		[date] => 2018-11-16,
		[url] => https://book.baidu.com/info/1011337448#test11111,
		[name] => 美国牧场的小生活,
		[author] => 醛石,
	),

	[33] =>(
		[url] => https://book.baidu.com/info/3660723#test11111,
		[name] => 抗日之特战兵王,
		[author] => 寂寞剑客,
		[date] => 2018-10-10,
	),

	[34] =>(
		[author] => 唐家三少,
		[date] => 2009-12-24,
		[url] => https://book.baidu.com/info/1115277#test11111,
		[name] => 斗罗大陆,
	),

	[35] =>(
		[name] => 都市至强者降临,
		[author] => 极地风刃,
		[date] => 2018-09-26,
		[url] => https://book.baidu.com/info/1009958915#test11111,
	),

	[36] =>(
		[url] => https://book.baidu.com/info/1009845790#test11111,
		[name] => 超级潇洒人生,
		[author] => 胖达福,
		[date] => 2018-05-10,
	),

	[37] =>(
		[name] => 超级神掠夺,
		[author] => 奇燃,
		[date] => 2018-11-12,
		[url] => https://book.baidu.com/info/1010695687#test11111,
	),

	[38] =>(
		[url] => https://book.baidu.com/info/1005115417#test11111,
		[name] => 火影之最强卡卡西,
		[author] => 墨渊九砚,
		[date] => 2018-05-31,
	),

	[39] =>(
		[name] => 壹号卫,
		[author] => 葛洛夫街兄弟,
		[date] => 2018-11-15,
		[url] => https://book.baidu.com/info/1010863551#test11111,
	),

	[40] =>(
		[name] => 这个天国不太平,
		[author] => 三江口水,
		[date] => 2018-11-30,
		[url] => https://book.baidu.com/info/1003640751#test11111,
	),

	[41] =>(
		[name] => 末世之宠物为王,
		[author] => 六枭,
		[date] => 2018-10-23,
		[url] => https://book.baidu.com/info/1010314675#test11111,
	),

	[42] =>(
		[name] => 全职高手,
		[author] => 蝴蝶蓝,
		[date] => 2014-04-30,
		[url] => https://book.baidu.com/info/1887208#test11111,
	),

	[43] =>(
		[author] => 龙鳞道V,
		[date] => 2018-10-18,
		[url] => https://book.baidu.com/info/1010519813#test11111,
		[name] => 奇遇无限,
	),

	[44] =>(
		[name] => 浴血兵锋,
		[author] => 丑牛1985,
		[date] => 2018-10-31,
		[url] => https://book.baidu.com/info/1010957891#test11111,
	),

	[45] =>(
		[date] => 2018-09-30,
		[url] => https://book.baidu.com/info/1010766239#test11111,
		[name] => 大唐农圣,
		[author] => 爱吃鱼的胖子,
	),

	[46] =>(
		[url] => https://book.baidu.com/info/1010256698#test11111,
		[name] => 位面之纨绔生涯,
		[author] => 昨夜大雨,
		[date] => 2018-08-20,
	),

	[47] =>(
		[name] => 穷鬼的上下两千年,
		[author] => 非玩家角色,
		[date] => 2018-10-25,
		[url] => https://book.baidu.com/info/1010696881#test11111,
	),

	[48] =>(
		[name] => 当个法师闹革命,
		[author] => 尹四,
		[date] => 2018-09-03,
		[url] => https://book.baidu.com/info/1004857359#test11111,
	),

	[49] =>(
		[name] => 我的舢舨能升级,
		[author] => 二哥的小号,
		[date] => 2018-11-02,
		[url] => https://book.baidu.com/info/1010690206#test11111,
	),
]
```