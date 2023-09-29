from flask import Flask, render_template, request, abort, redirect
import flask.json as flask_json  # 尝试利用 json , 但始终不明白方法, 这条 import 就留着吧
import requests
import sqlite3
import json
import os  # 检测网络状况, 获取目录
import time  # 获取当前时间，用于生成日志, 以及实现签到功能
import webbrowser  # 浏览器, app启动后直接跳转至网页

def now_time():  # 
    return time.strftime("%Y-%m-%d %H-%M-%S", time.localtime())

app = Flask(__name__)

flag_signed_in = False  # 初始状态: 未登录
user_signed_in = [0, 0, 0]  # 记录登录状态下 user 的 id, username, password 
database_path = "%s\\src\\main.db"%os.path.dirname(__file__)  # main.db 数据库文件的绝对路径
log_filename = "%s\\log\\log_%s.txt"%(os.path.dirname(__file__), now_time())  # 本次运行可能产生的日志文件的绝对路径

"""
实现日志记录功能:
各网页视图函数的编写利】采用 try-except 结构, 
当代码运行出现错误时将错误信息Error写出至 .txt 文件, 并且记录错误产生的 part , 以供程序的修复
如果采用这种结构, 那么:
    如果一次运行过程中没有产生错误, 那么就不会有日志文件的产生;
    如果产生了错误, 那么会在日志文件中写入错误信息
"""
def write_log(part, Error):
    global log_filename
    
    f = open(log_filename, 'x')
    f.write('Error:    Part-{}    {}'.format(part, Error))
    f.close()

"""
实现签到每日一次功能:
在 main.db users 表中插入新的列: checkin ,设默认值为 'None', 
(如果不设置默认值, 则对于id=1,2...等先前没有 checkin 的记录, checkin 处值为 None , 无法直接通过 update 语句更新 checkin 的状态.)
签到过程:  user 签到, 在数据库中取出记录, 比对 checkin 的值与签到当天的日期. 如果一致, 说明当天已经签到过；如果不一致, 则将数据库中 checkin 值改为当天日期
"""
def database_add_col_Checkin():
    global database_path
    
    db = sqlite3.connect(database_path)
    cursor = db.cursor()
    temp_record = cursor.execute("select * from users where id=1").fetchall()
    """
    随便取出一条记录
    对于所给定的数据库, 在初始状态下, users 中只有 id, username, password 三个字段.
    那么, 令 temp_record = 该记录
    则: len( temp_record[0] ) = 3
    
    Why:
    签到: 在数据库中加入Checkin字段, user 签到时将自己的Checkin值改为“今天”日期
          如果Checkin值与今天日期不同, 则今天没签到。
    """
    if len(temp_record[0]) == 3:
        cursor.execute("alter table users add column checkin text default 'None'")
    db.commit()
    cursor.close()
    db.close()

"""
测试网络连通性
如果网络连接正常, 那么 os.system(u"ping %s") s: 网络正常状态下可顺利访问的 url , 例如 www.baidu.com
"""
@app.route('/ping', methods=["GET"])  # https://blog.csdn.net/weixin_40449300/article/details/79193872
def test_ping():
    try:
        result = os.system(u'ping www.baidu.com')
        if result == 0:
            data = {
            "code": 0,
            "msg": "pong!"
            }
        
            data = json.dumps(data)
            headers = {'Content-Type': 'application/json'}
            requests.post(url='http://127.0.0.1:8080/ping', json=data, headers=headers)
            # 尝试发送json, 但似乎无效, 后来没有再尝试用 json
        
            return "Internet Connected!"
        else:
            abort(408)
    except Exception as Error:
        write_log('ping', Error)
        
"""
signup 注册部分
"""
@app.route('/signup', methods=['POST', 'GET'])
def signup():
    global database_path, flag_signed_in, user_signed_in
    try:
        condition="None"
        if request.method == "POST":
            username = request.form.get("username")
            password = request.form.get("password")
            # 在数据库中查重
            
            db = sqlite3.connect(database_path)
            cursor = db.cursor()
            username_db = cursor.execute("select username from users where username=%s"%username).fetchall()
            if username_db == []:  # 数据库中没有用户信息, 即未注册
                id_latest = cursor.execute("select id from users").fetchall()[-1][0]  # 取最大的id
                # 因为数据库中已经有记录, 不再处理id=1的情况.
                user_id = id_latest + 1
                cursor.execute("insert into users values(%d, '%s', '%s', 'None')"%(user_id, username, password))  # 初始未签到
                db.commit()
                condition = "注册成功!\nID :%d\n用户名:%s\n密码 :%s"%(user_id, username, password)
                flag_signed_in = False  # 注册成功后自动登出账号
                user_signed_in = [0, 0, 0]
            else:
                condition = "用户名已存在!"
            cursor.close()
            db.close()
        return render_template("signup.html", condition=condition)
    except Exception as Error:
        write_log('signup', Error)

"""
signin 登录部分
"""
@app.route('/signin', methods=['POST', 'GET'])
def signin():
    global flag_signed_in, database_path, user_signed_in
    try:
        if flag_signed_in == True:
            condition = "%s 已登录"%user_signed_in[1]  # 取 username
        else:
            condition = "未登录"
        if request.method == "POST":
            db = sqlite3.connect(database_path)
            cursor = db.cursor()
            username = request.form.get("username")
            password = request.form.get("password")
            password_db = cursor.execute("select id,password from users where username='%s'"%username).fetchall()
            if password_db == []:
                condition = "用户名不存在!"
                flag_signed_in = False
                user_signed_in = [0, 0, 0]
            elif password_db[0][1] == password:  # password_db = [(id, password)]
                user_signed_in[0] = password_db[0][0]
                user_signed_in[1] = username
                user_signed_in[2] = password
                flag_signed_in = True
                condition = "欢迎, %s !"%username
            else:  # 密码错误
                flag_signed_in = False  # 密码错误, 登录失败, 自动登出当前帐号
                user_signed_in = [0, 0, 0]
                condition = "密码错误!"
            cursor.close()
            db.close()
        return render_template("signin.html", condition=condition)
    except Exception as Error:
        write_log('signin', Error)
            
"""
签到部分
"""
@app.route('/checkin', methods=['POST', 'GET'])
def checkin():  # 先登录, 再签到
    global flag_signed_in, user_signed_in, database_path
    try:
        if flag_signed_in == True:
            condition = "%s 已登录"%user_signed_in[1]
        else:
            condition = "未登录。请先登录, 后签到"
        if request.method == 'POST':
            if flag_signed_in == False:
                condition = "未登录。请先登录, 后签到"
            else:
                db = sqlite3.connect(database_path)
                cursor = db.cursor()
                temp_checked_in = cursor.execute("select checkin from users where id=%d"%user_signed_in[0]).fetchall()[0][0]
                checkin_time = now_time()[:10]  # 只用取年-月-日: 签到一天一次
                if temp_checked_in == checkin_time:
                    condition = "今日已经签到"
                else:
                    cursor.execute("update users set checkin='%s' where id=%d"%(checkin_time, user_signed_in[0]))
                    db.commit()
                    condition = "签到成功! %s"%checkin_time
                cursor.close()
                db.close()
        return render_template("checkin.html", condition=condition)
    except Exception as Error:
        write_log('checkin', Error)
 
"""           
@app.route('/readme')
def show_info():
    file = open('%s\\README.md'%(os.path.dirname(__file__)), 'r')
    text_list = []
    line = file.readline()
    while line:
        text_list.append(line)
        line = file.readline()
    file.close()        
    return '\n'.join(text_list)
"""

@app.route('/', methods=['POST', 'GET'])
def homepage():
    try:
        if request.method == "POST":
            index = request.form.get("index")
        
            if index == "1":
                return redirect('/ping')
        
            elif index == "2":
                return redirect('/signup')
        
            elif index == "3":
                return redirect('/signin')
        
            elif index == "4":
                return redirect('/checkin')
        return render_template("homepage.html")
    except Exception as Error:
        write_log('homepage', Error)

if __name__ == '__main__':
    database_add_col_Checkin()
    webbrowser.open('http://127.0.0.1:8080')
    app.run(port=8080)
