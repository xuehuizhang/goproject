package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"math"
	"path"
	"quickstart/models"
	"strconv"
	"strings"
	"time"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	/*o:=orm.NewOrm()

	user:=models.User{}

	//user.Id=1
	user.Name="李四"
	err:=o.Read(&user,"Name")

	if err!=nil{
		beego.Info("查询失败",err)
		return
	}

	beego.Info("查询成功",user)
	c.Data["data"] = "Hello world"
	c.TplName = "test.html"*/
    //创建一个orm
   /* o:=orm.NewOrm()
    //创建一个对象
    user:=models.User{}
    //根据name查询
    user.Name="李四"

    err:=o.Read(&user,"Name")
    if err!=nil{
    	beego.Info("查询失败",err)
		return
	}
    beego.Info("查询成功：",user)
	c.Data["data"]="hello world"
	c.TplName="test.html"*/
	c.TplName="register.html"
}

func (c *MainController)Post()  {
/*	//创建一个orm对象
	o:=orm.NewOrm()
    //创建一个要插入的对象
    user:=models.User{}
    //初始化
    user.Name="李四"
    user.Pwd="123"

    r,err:=o.Insert(&user)
	if err!=nil{
		beego.Info("插入失败",err)
		return
	}*/
	//创建一个orm对象
	/*o:=orm.NewOrm()

	user :=models.User{Name:"马武",Pwd:"345"}

	r,err:=o.Insert(&user)
	if err!=nil{
		beego.Info("插入失败",err)
		return
	}
	beego.Info("插入成功",user)*/

	/*o :=orm.NewOrm()

	user:=models.User{}

	user.Id=1

	err:=o.Read(&user)

	if err==nil{
		user.Name="赵六"
		user.Pwd="123"
		_,err=o.Update(&user)
		if err!=nil{
			beego.Info("更新失败",err)
		}else{
			beego.Info("更新成功")
		}
	}*/

	/*o:=orm.NewOrm()
	user:=models.User{}
	user.Id=1
	_,err:=o.Delete(&user)
	if err!=nil{
		beego.Info("删除失败",err)
		return
	}
	beego.Info("删除成功")
	c.Data["data"]=fmt.Sprintf("你好世界：%d",1)
	c.TplName="test.html"*/
	name:=c.GetString("name")
	pwd:=c.GetString("pwd")
	beego.Info("新添加的用户",name,pwd)
    if name==""||pwd==""{
    	beego.Info("数据不能为空")
    	c.Data["info"]="数据不能为空"
    	c.Redirect("/reg",302)
		return
	}
	o:=orm.NewOrm()
	user:=models.User{}
	user.Name=name
	user.Pwd=pwd
	_,err:=o.Insert(&user)
	if err!=nil{
		beego.Info("注册失败",err)
		c.Redirect("/reg",302)
		return
	}
	beego.Info("注册成功")
	c.Redirect("/",302)
	return
}

func (c *MainController)ShowLogin()  {
	userName:=c.Ctx.GetCookie("userName")
	pwd:=c.Ctx.GetCookie("pwd")
	beego.Info(userName,pwd)

	if userName!=""&&pwd!=""{
		c.Data["rem"]=true
		c.Data["userName"]=userName
		c.Data["pwd"]=pwd
	}
	c.TplName="login.html"
}

func (c *MainController)HandleLogin()  {
	name:=c.GetString("name")
	pwd:=c.GetString("pwd")
	rem:=c.GetString("rem")
	beego.Info(rem)
	if name==""||pwd==""{
		beego.Info("用户名密码不能为空")
		c.Redirect("/",302)
		return
	}

	o:=orm.NewOrm()
	user:=models.User{}
	user.Name=name
	err:=o.Read(&user,"name")
	beego.Info(user)
	if err!=nil{
		beego.Info("不存在当前用户")
		c.Redirect("/",302)
		return
	}
	if user.Pwd!=pwd{
		beego.Info("密码错误")
		c.Redirect("/",302)
		return
	}
	if rem=="on" {
		c.Ctx.SetCookie("userName", name, time.Second*100)
		c.Ctx.SetCookie("pwd", pwd, time.Second*10)
	}else{
		c.Ctx.SetCookie("userName", name, -1)
		c.Ctx.SetCookie("pwd", pwd, -1)
	}
	c.SetSession("userName",name)
	beego.Info("登录成功")
	c.Redirect("/index",302)
}

func (c *MainController)Index()  {
	name:=c.GetSession("userName")
	if name==nil{
		beego.Info("请登录")
		c.Redirect("/",302)
		return
	}
	o:=orm.NewOrm()
    id,_:=c.GetInt("artiType")
	var articles []models.Article
	qs:=o.QueryTable("article")

	pageIndex1:=c.GetString("pageIndex")
	pageIndex,err:=strconv.Atoi(pageIndex1)
	if err!=nil{
		pageIndex=1
	}

	pageSize:=2
	var  count int64
	if id!=0 {
		count,err= qs.RelatedSel("ArticleType").Filter("ArticleType__Id", id).Count()
	}else{
		count,err=qs.RelatedSel("ArticleType").Count()
	}
	beego.Info("总记录数：",count)
	pageCount:=math.Ceil(float64(count)/float64(pageSize))

	start:=pageSize*(pageIndex-1)

	if id!=0 {
		_, err = qs.Limit(pageSize, start).RelatedSel("ArticleType").Filter("ArticleType__Id", id).All(&articles)
	}else {
		_,err=qs.Limit(pageSize,start).RelatedSel("ArticleType").All(&articles)
	}
	if err!=nil{
		beego.Info("查询失败")
		c.Redirect("/index",302)
		return
	}
	artiTypes:=make([]models.ArticleType,0)
	_,err=o.QueryTable("article_type").All(&artiTypes)
	if err!=nil{
		beego.Info("查询文章类型失败")
		c.Redirect("/index",302)
		return
	}

	firstPage:=false
	if pageIndex==1{
		firstPage=true
	}
	endPage:=false
	if float64(pageIndex)==pageCount{
		endPage=true
	}
	c.Data["artiTypes"]=artiTypes
	c.Data["firstPage"]=firstPage
	c.Data["endPage"]=endPage
	c.Data["count"]=count
	c.Data["pageCount"]=pageCount
	c.Data["articles"]=articles
	c.Data["pageIndex"]=pageIndex
	c.Data["typeId"]=id

	c.LayoutSections=make(map[string]string)
	c.LayoutSections["head"]="indexhead.html"
	c.Layout="layout.html"
	c.TplName="index.html"
}

func (c *MainController)ShowAdd()  {
	artiTypes:=make([]models.ArticleType,0)
	o:=orm.NewOrm()
	_,err:=o.QueryTable("article_type").All(&artiTypes)
	if err!=nil{
		beego.Info("获取文章类型失败")
		c.Redirect("/addarti",302)
		return
	}
	c.Data["artiTypes"]=artiTypes
	c.TplName="add.html"
}

func (c *MainController)HandleAdd()  {
	aname:=c.GetString("aname")
	acontent:=c.GetString("acontent")
	atypeId,_:=c.GetInt("atype")
	beego.Info("atype=",atypeId)
	f,h,err:=c.GetFile("aimg")
	defer f.Close()

	//限制后缀
	fileExt:=path.Ext(h.Filename)
	if strings.ToLower(fileExt)!=".jpg"&&strings.ToLower(fileExt)!=".png"{
		beego.Info("图片格式错误")
		c.Redirect("/addarti",302)
		return
	}
	//限制大小
	if h.Size>50000{
		beego.Info("图片太大")
		c.Redirect("/addarti",302)
		return
	}
	//文件重命名
	fileName:=time.Now().Format("2006-01-02-15-04-05")+fileExt
	if err!=nil{
		beego.Info("图片上传失败",err)
		c.Redirect("/addarti",302)
		return
	}else{
		beego.Info(h.Filename)
		c.SaveToFile("aimg","./static/img/"+fileName)
	}
	if aname==""||acontent==""{
		beego.Info("数据不能为空")
		c.Redirect("/addarti",302)
		return
	}
	o:=orm.NewOrm()
	var artiType models.ArticleType
	artiType.Id=atypeId
	err=o.Read(&artiType)
	if err!=nil{
		beego.Info("查询类型失败")
		c.Redirect("/addarti",302)
		return
	}

	arti:=models.Article{}
	arti.Aname=aname
	arti.Acontent=acontent
	arti.Aimg="/static/img/"+fileName
	arti.Acount=0
	arti.ArticleType=&artiType
	_,err=o.Insert(&arti)
	if err!=nil{
		beego.Info("添加失败",err)
		c.Redirect("/addarti",302)
		return
	}
	c.Redirect("/index",302)
}

func (c *MainController)ShowContent()  {
	id:=c.GetString("id")
	beego.Info(id)
	o:=orm.NewOrm()
	arti:=models.Article{}
	i,err:=strconv.Atoi(id)
	beego.Info(i)
	arti.Id=i
	err=o.Read(&arti)
	if err!=nil{
		beego.Info("获取详情出错",err)
		c.Redirect("/index",302)
		return
	}
	arti.Acount+=1
	_,err=o.Update(&arti)

	if err!=nil{
		beego.Info("更新阅读数失败",err)
		c.Redirect("/index",302)
		return
	}

	m2m:=o.QueryM2M(&arti,"User")
	user:=models.User{}
	userName := c.GetSession("userName").(string)
	beego.Info(userName)
	user.Name=userName
	o.Read(&user,"Name")
	m2m.Add(&user)


	_,err=o.LoadRelated(&arti,"User")
/*	var users []models.User
	o.QueryTable("User").RelatedSel("Article").Filter("Artilce__Artilce__Id",i).Distinct().All(&users)
	*/
	beego.Info(arti.User)
	if err!=nil{
		beego.Info("获取阅读人失败")
		c.Redirect("/index",302)
		return
	}
	beego.Info(arti.User)
	c.Data["arti"]=arti
	c.TplName="content.html"
}

func (c *MainController)ShowUpdate()  {
	id,err:=c.GetInt("id")
	if err!=nil{
		beego.Info("获取Id出错")
		c.Redirect("/index",302)
		return
	}
	o:=orm.NewOrm()
	arti:=models.Article{}
	arti.Id=id
	err=o.Read(&arti)
	if err!=nil{
		beego.Info("获取信息失败")
		c.Redirect("/index",302)
		return
	}
	c.Data["arti"]=arti
	c.TplName="update.html"
}
func (c *MainController)HandleUpdate()  {
	id,err:=c.GetInt("id")
	beego.Info(id)
	if err!=nil{
		beego.Info("更新：获取id失败",err)
		c.Redirect("/index",302)
		return
	}
	aname:=c.GetString("aname")
	acontent:=c.GetString("acontent")

	if aname==""||acontent==""{
		beego.Info("数据不能为空",err)
		c.Redirect("/index",302)
		return
	}

	f,h,err:=c.GetFile("aimg")
	defer f.Close()
	if err!=nil{
		beego.Info("上传图片失败",err)
		c.Redirect("/index",302)
		return
	}else{
		//判断格式
		fileExt:=path.Ext(h.Filename)
		if fileExt!=".jpg"&&fileExt!=".png"{
			beego.Info("格式不正确")
			c.Redirect("/index",302)
			return
		}
		//判断大小
		if h.Size>50000{
			beego.Info("文件太大")
			c.Redirect("/index",302)
			return
		}
		//重命名
		fileName:=time.Now().Format("2006-01-02-15-04-05")+fileExt
		c.SaveToFile("aimg","./static/img/"+fileName)
		arti:=models.Article{Id:id}
		o:=orm.NewOrm()
		err:=o.Read(&arti)
		if err!=nil{
			beego.Info("不存在当前文章")
			c.Redirect("/index",302)
			return
		}
		arti.Aname=aname
		arti.Acontent=acontent
		arti.Aimg="/static/img/"+fileName
		beego.Info(arti)
		_,err=o.Update(&arti,"Aname","Acontent","Aimg")
		if err!=nil{
			beego.Info("修改出错",err)
			c.Redirect("/index",302)
			return
		}
	}
	c.Redirect("/index",302)
}

func (c *MainController)DeleteArti()  {
	id,err:=c.GetInt("id")
	if err!=nil{
		beego.Info("获取Id出错",err)
		c.Redirect("/index",302)
		return
	}
	o:=orm.NewOrm()
	arti:=models.Article{Id:id}
	err=o.Read(&arti)
	if err!=nil{
		beego.Info("不存在当前文章",err)
		c.Redirect("/index",302)
		return
	}
	_,err=o.Delete(&arti)
	if err!=nil{
		beego.Info("删除失败",err)
		c.Redirect("/index",302)
	}
	beego.Info("删除成功")
	c.Redirect("/index",302)
}

func (c *MainController)ShowAddArtiType()  {
	o:=orm.NewOrm()
	artiTypes:=make([] models.ArticleType,1)
	qt:=o.QueryTable("article_type")
	_,err:=qt.All(&artiTypes)
	if err!=nil{
		beego.Info("查询类型失败")
		c.Redirect("/addartitype",302)
		return
	}
	c.Data["artiTypes"]=artiTypes
	c.TplName="addType.html"
}

func (c *MainController)HandleAddArtiType()  {
	typeName:=c.GetString("typeName")
	if typeName==""{
		beego.Info("类型名称不能为空")
		c.Redirect("/addartitype",302)
		return
	}
	o:=orm.NewOrm()
	artiType:=models.ArticleType{}
	artiType.TypeName=typeName
	_,err:=o.Insert(&artiType)
	if err!=nil{
		beego.Info("添加失败")
		c.Redirect("/addartitype",302)
		return
	}
	c.Redirect("/addartitype",302)
}

func (c *MainController)DelArtiType()  {
	id,err:=c.GetInt("id")
	if err!=nil{
		beego.Info("没获得id")
		c.Redirect("/addartitype",302)
		return
	}
	o:=orm.NewOrm()
	artiType:=models.ArticleType{Id:id}
	_,err=o.Delete(&artiType)
	if err!=nil{
		beego.Info("删除文章类型失败")
		c.Redirect("/addartitype",302)
		return
	}
	beego.Info("删除文章类型成功")
	c.Redirect("/addartitype",302)
}

func (c *MainController)HandleArtiType()  {
	typeName:=c.GetString("artiType")
	beego.Info(typeName)

	var articles []models.Article
	o:=orm.NewOrm()
	_,err:=o.QueryTable("Article").RelatedSel("ArticleType").Filter("ArticleType__TypeName",typeName).All(&articles)
	if err!=nil{
		beego.Info("查询失败")
		c.Redirect("/index",302)
		return
	}
	beego.Info(articles)
	c.Data["articles"]=articles
	//c.TplName="index.html"
	c.Ctx.WriteString("查询成功")
}

func (c *MainController)LogOut()  {
	c.DelSession("userName")
	c.Redirect("/",302)
}

func (c *MainController)Query()  {
	o:=orm.NewOrm()  //创建orm对象
	r:=o.Raw("select * from user")
	var users []models.User
	res,err:=r.QueryRows(&users)
	if err!=nil{
		beego.Info("查询出错",err)
		return
	}
	beego.Info(users,res)
}
