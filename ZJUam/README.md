# zjuam

## 前言
这里没有完成的项目, 只有一个 README.md, 记录着我对这次开发尝试的种种想法。

## 感受
自始至终, 都没有找到一个合适的解决问题的途径.

做完原型图, 信心满满, 准备学习 GO 语言
边学GO语言, 边试着用开发者工具在统一身份认证的网站上找出有用的信息
GO语言学会了一点, 网站的 security.js, login.js, html 也找到了
exponent, modulus 在 https://zjuam.zju.edu.cn/cas/v2/getPubKey, 可以直接GET!
然后: 
    数据怎么发送? 我要把数据发送到哪? 数据发送到网站后是怎么处理的?
    RSA的加密是哪一步实现的? 我要把 security.js 中的加密函数翻译成 GO 语言吗?
    要实现自动登录, 我要先手动成功登录一遍吗? Cookie 怎么获取?
思路断了。
我事实上没有任何的 Web 开发基础, 对 Web 元素的基本概念也没有形成系统性的认知
然而, 如果我要对某方面的知识形成一定的认识与理解, 我必须要有这方面的系统性的基础知识.
但是, 时间有限, 而在这过程中我有点好高骛远了, 没有尝试构建 Web开发 系统性的知识基础。
因此, 经过 3 天 的死磕与查资料, 我仍然没有形成一定的 Web 开发的概念。(这也让我没法实现 Naive Server 的 'token 处理', 'api 请求和响应采用 json' 两个基本功能)

怎么说呢...虽然 ZJUam 这个任务我最后没能完成, Web 开发的知识也没能掌握
但是我也收获了对自己思维方式的认知: 在给定基础性、系统性的知识的前提下, 偏向于自我独立探索、开创型的思维
以及一场自我提醒过无数遍的教训: 重视基础.
希望以后有机会能弄明白这个过程.

## 回忆
1. GO 的形象好可爱
2. {
    go mod init test
    go run .
    Hello World!
   }
3. F12 login:       <div class="login-middle">
                        <input id="username" name="username" class="form-control user-input" ...
                        <input id="password" name="password" class="form-control pwd-input"  ...
                        <input type="hidden" name="execution" value=...
        secutity.js RSAUtils.xxx
        login.js    function checkForm() {...}
4.  /try/try_http/test_get_url/test_get.go     运行, 输出modulus, exponent
5.  /try/try_ZJUam/Project_zjuam/Encrypt.go    翻译 security.js 部分
6.  import "github.com/robertkrimen/toot"      用 otto 运行 JS 代码
7.  没时间了, 写不完了

## 文件
/try       将所有的尝试都放在这里了, 包括 GO 学习最初的代码
/try/tey_hello 应该是: Hello World, 以及从同目录下其它文件导入函数
/try/try_http  试图学习 http 请求一类的 Web 开发相关的内容
/try/try_import  尝试从同目录下其它文件导入函数
/try/try_fyne    既然是客户端, 总得有个 UI 吧 ; 计划学习 fyne 的使用做 UI 界面
/try/try_otto    导入 otto 库, 尝试直接执行 security.js 中的加密函数
/try/try_ZJUam   尝试完成题目的要求 (Project_zjuam)

## 过程
2023.09.25 安装 Go 实现 Hello World

2023.09.26 研究 zjuam.zju.edu.cn/cas/login 
2023.09.26 分析html, js (0基础瞎分析)  
2023.09.26 security.js SHA加密相关
           Line 190 hexToDigit
           Line 200 biFromHex
           Line627 getKeyPair获取exponent, modulus   
           Line 640 加密函数主体
2023.09.26 login.js       Line 121 checkForm() 关键函数    密码先倒序, 后处理
2023.09.26 开发者工具分析: zjuam.zju.edu.cn/cas/v2/getPubKey

2023.09.27 Get 获取 exponent, modulus
2023.09.27 尝试翻译 Security.js 至 GO (Encrypt.go)
2023.09.27 突然发现有跑 JavaScript 的 GO 引擎

2023.09.28 迫于时间压力放弃项目, 转向Naive Server

## 参考(部分, 约占有价值部分的 1/3)
登录流程分析  https://blog.csdn.net/qq_42598133/article/details/125658235
GO 使用Cookie 做用户登录管理 https://blog.csdn.net/paterl/article/details/130400139

GO 使用Cookie 实现登录(gin + gorm) https://blog.csdn.net/weixin_51299478/article/details/122848795?spm=1001.2101.3001.6650.1&utm_medium=distribute.pc_relevant.none-task-blog-2%7Edefault%7EBlogCommendFromBaidu%7ERate-1-122848795-blog-130400139.235%5Ev38%5Epc_relevant_sort_base2&depth_1-utm_source=distribute.pc_relevant.none-task-blog-2%7Edefault%7EBlogCommendFromBaidu%7ERate-1-122848795-blog-130400139.235%5Ev38%5Epc_relevant_sort_base2&utm_relevant_index=2

GO 设置、读取和删除 Cookie http://www.guoxiaolong.cn/blog/?id=9812
GO 获取网页内容 https://bbs.huaweicloud.com/blogs/403624
Go 为什么Response.Body需要被关闭&&内存泄漏数量 https://blog.csdn.net/qq_43778308/article/details/115448697
GO 遍历字符串  https://www.php.cn/faq/500058.html
GO 创建字符串    https://juejin.cn/post/6989118832379953159    https://juejin.cn/post/6844903796917682189?from=search-suggest
GO 语言运算符 https://www.runoob.com/go/go-operators.html
GO Type  http://kangkona.github.io/oo-in-golang/
GO 中的类和对象    https://www.geeksforgeeks.org/class-and-object-in-golang/
GO container/list List http://c.biancheng.net/view/35.html
GO containter/list List 反向遍历 https://haicoder.net/golang/golang-list-walk.html
GO 中调用 JS 代码: https://blog.csdn.net/qq_42527676/article/details/86481835
GO 中调用 JS 代码   https://juejin.cn/post/6844904002975432717
GO 中调用  JS  代码    https://blog.51cto.com/u_16099241/6467254
JS substr  https://www.runoob.com/jsref/jsref-substr.html
JS ZERO_ARRAY