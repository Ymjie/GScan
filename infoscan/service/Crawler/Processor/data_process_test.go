package Processor

import (
	"golang.org/x/text/encoding/simplifiedchinese"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

var a string = `

<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=100%">

<link type="text/css" href="/_css/_system/system.css" rel="stylesheet"/>
<link type="text/css" href="/_upload/site/1/style/62/62.css" rel="stylesheet"/>
<link type="text/css" href="/_upload/site/00/14/20/style/64/64.css" rel="stylesheet"/>
<link type="text/css" href="/_js/_portletPlugs/simpleNews/css/simplenews.css" rel="stylesheet" />
<link type="text/css" href="/_js/_portletPlugs/sudyNavi/css/sudyNav.css" rel="stylesheet" />
<link type="text/css" href="/_js/_portletPlugs/datepicker/css/datepicker.css" rel="stylesheet" />

<script language="javascript" src="/_js/jquery.min.js" sudy-wp-context="" sudy-wp-siteId="20"></script>
<script language="javascript" src="/_js/jquery.sudy.wp.visitcount.js"></script>
<script type="text/javascript" src="/_js/_portletPlugs/sudyNavi/jquery.sudyNav.js"></script>
<script type="text/javascript" src="/_js/_portletPlugs/datepicker/js/jquery.datepicker.js"></script>
<script type="text/javascript" src="/_js/_portletPlugs/datepicker/js/datepicker_lang_HK.js"></script>
<link rel="icon" href="/_upload/tpl/07/29/1833/template1833/favicon.ico" mce_href="/_upload/tpl/07/29/1833/template1833/favicon.ico" type="image/x-icon">
<title>院所快讯</title>
<meta name="keywords" content="新闻网">
<meta name="description" content="齐鲁工业大学" >
<meta name="renderer" content="webkit" />
<meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
<meta name="viewport" content="width=device-width,user-scalable=0,initial-scale=1.0, minimum-scale=1.0, maximum-scale=1.0"/>
<script type="text/javascript" src="/_upload/tpl/07/29/1833/template1833/extends/extends.js"></script>
<script type="text/javascript" src="/_upload/tpl/07/29/1833/template1833/js/iscroll.js"></script>
<link rel="stylesheet" href="/_upload/tpl/07/29/1833/template1833/style.css" type="text/css" />
<link rel="stylesheet" href="/_upload/tpl/07/29/1833/template1833/mobile.css" type="text/css"/>
<link rel="stylesheet" href="/_upload/tpl/07/29/1833/template1833/media.css" type="text/css"/>
<!--[if lt IE 9]>
	<script src="/_upload/tpl/07/29/1833/template1833/extends/libs/html5.js"></script>
<![endif]-->
</head>
<body class="list">
<!--Start||headTop-->
<div class="wrapper headtop" id="headtop">
	<div class="inner">
		<div class="mod clearfix">
			<!--头部信息-->
			<div class="top-left">
				<div class="site-rale" frag="窗口04">
					
				</div>					
			</div>
			<div class="top-right" >
				<div class="site-lang" frag="窗口05">
					
				</div>				
			</div>			
		</div>
    </div>
</div>
<!--End||headTop-->
<!--Start||head-->
<div class="wrapper header" id="header">
	<div class="inner">
		<div class="mod clearfix">
			<a class="navi-aside-toggle"></a>
			<div class="head-left" frag="面板01">
				<!--logo开始-->
				<div class="sitelogo" frag="窗口01" portletmode="simpleSiteAttri">
					<a href="/main.htm" title="返回新闻网首页"><div class="logobox"><img src="/_upload/tpl/07/29/1833/template1833/images/logo.png"></div><span class="sitetitle">新闻网</span></a> 
				</div>
				<!--//logo结束-->
			</div>
			<div class="head-right">
				<div class="site-lang clearfix" frag="窗口02" portletmode="simpleNews" configs="{'c8':'1','c42':'320','c25':'157','c30':'0','c29':'1','c23':'3','c34':'300','c20':'0','c31':'0','c16':'1','c3':'1','c2':'缩略图','c27':'480','c43':'0','c17':'0','c5':'_blank','c24':'40','c32':'','c26':'1','c37':'1','c28':'640','c40':'1','c15':'0','c14':'1','c44':'0','c33':'500','c10':'50','c18':'yyyy-MM-dd','c36':'0','c1':'1','c6':'15','c19':'yyyy-MM-dd','c21':'0','c4':'1','c35':'-1:-1','c39':'300','c38':'100','c7':'1','c12':'0','c9':'0','c11':'1','c13':'200','c41':'240'}" contents="{'c2':'0', 'c1':'/投稿专区'}"> 
					
						<ul class="news_list clearfix">
							
							<li class="news i1 clearfix">
								<div class="news_imgs"><a href='http://webplus.qlu.edu.cn/' target='_blank' title='投稿专区'><img src='/_upload/article/images/e3/c4/818ab1284e209022d78a6b2fb6f5/73441c3a-2631-4d47-ac09-9f4b6135bbf7_s.jpg' width='157' height='40' /></a></div>
							</li>
							
						</ul>			
					
				</div>
				<div class="searchbox" frag="窗口03" portletmode="search">
					<div class="wp-search clearfix">
							<form action="/_web/search/doSearch.do?locale=zh_CN&request_locale=zh_CN&_p=YXM9MjAmdD0xODMzJmQ9NjIzNyZwPTImZj04MDgmbT1TTiZ8Ym5uQ29sdW1uVmlydHVhbE5hbWU9ODA4Jg__" method="post" target="_blank">
								<div class="search-input">
									<input name="keyword" class="search-title" type="text" placeholder="SEARCH..."/>
								</div>
								<div class="search-btn">
									<input name="submit" class="search-submit" type="submit" value=""/>
								</div>
							</form>
						</div>
				</div>			
			</div>

		</div>
	</div>
</div>
<!--End||head-->
<!--Start||nav-->
<div class="wrapper nav wp-navi" id="nav">
	<div class="inner clearfix">
		<div class="wp-panel">
			<div class="wp-window" frag="窗口1" portletmode="simpleSudyNavi" configs="{'c1':'1','c7':'2','c4':'_self','c3':'15','c8':'2','c9':'0','c2':'1','c5':'2'}" contents="{'c2':'0', 'c1':'/首页,/工大要闻,/综合新闻,/部门动态,/院所快讯,/媒体工大,/人物在线'}">
				
					<div class="navi-slide-head clearfix">
						<h3 class="navi-slide-title">导航</h3>
						<a class="navi-slide-arrow"></a> 
					</div>
					
					<ul class="wp-menu" data-nav-aside='{"title":"导航","index":0}'>
						
						<li class="menu-item i1"><a class="menu-link" href="/main.htm" target="_self">首页</a>
							
						</li>
						
						<li class="menu-item i2"><a class="menu-link" href="/811/list.htm" target="_self">工大要闻</a>
							
						</li>
						
						<li class="menu-item i3"><a class="menu-link" href="/zhxw/list.htm" target="_self">综合新闻</a>
							
						</li>
						
						<li class="menu-item i4"><a class="menu-link" href="/813/list.htm" target="_self">部门动态</a>
							
						</li>
						
						<li class="menu-item i5"><a class="menu-link" href="/814/list.htm" target="_self">院所快讯</a>
							
						</li>
						
						<li class="menu-item i6"><a class="menu-link" href="/826/list.htm" target="_self">媒体工大</a>
							
						</li>
						
						<li class="menu-item i7"><a class="menu-link" href="/rwzx/list.htm" target="_self">人物在线</a>
							
						</li>
						
					</ul>
					
				
			</div>
		</div>
	</div>
</div>
<!--aside导航-->
<div class="wp-navi-aside" id="wp-navi-aside">
  <div class="aside-inner">
    <div class="navi-aside-wrap"></div>
  </div>
  <div class="navi-aside-mask"></div>
</div>
<!--End||nav-->
<div class="shadow"></div>
<!--Start||focus-->
<div class="wp-wrapper" id="container-1">
	<div class="wp-inner" frag="面板84">
		<div class="l-banner" frag="窗口84" portletmode="simpleColumnAttri" configs="{'c3':'400','c1':'0','c2':'300'}">
			
				<img border="0" style="margin:0 auto;" src="" data-imgsrc="/_upload/tpl/07/29/1833/template1833/images/list_banner.jpg">
			
		</div>
	</div>
</div>
<!--End||focus-->
<!--Start||content-->
<div class="wrapper" id="l-container">
	<div class="inner">
		<div class="mod clearfix">
			<div class="col_menu">
				<div class="col_menu_head">
					<h3 class="col_name" frag="窗口3" portletmode="simpleColumnAnchor">
						<span class="col_name_text">
						<span class='Column_Anchor'>院所快讯</span>
						</span>
					</h3>
					<a class="column-switch"></a>
				</div>
				<div class="col_menu_con" frag="面板4">
					<div class="col_list" frag="窗口4" portletmode="simpleColumnList">
						
							
							<ul class="wp_listcolumn list-paddingleft-2">
								
							</ul>
										
						
					</div>
				</div>
			</div>
			<div class="col_news">
				<div class="col_news_box">
					<div class="col_news_head">
						<ul class="col_metas clearfix" frag="窗口5" portletmode="simpleColumnAttri">
						   
							<li class="col_path"><span class="path_name">当前位置：</span><a href="/main.htm" target="_self">首页</a><span class='possplit'>&nbsp;&nbsp;</span><a href="/814/list.htm" target="_self">院所快讯</a></li>
							<li class="col_title"><h2>院所快讯</h2></li>
						   
						</ul>
					</div>
					<div class="col_news_con" >
						<div class="col_news_list listcon">
							<div frag="窗口6" portletmode="simpleList" configs="{'c8':'1','c23':'320','c31':'1','c28':'0','c27':'1','c21':'1','c20':'0','c16':'1','c3':'14','c2':'序号,标题,发布时间','c25':'480','c17':'0','c5':'_blank','c22':'240','c32':'1','c24':'1','c26':'640','c15':'0','c14':'1','c29':'1','c33':'1','c34':'0','c10':'50','c18':'yyyy-MM-dd','c1':'1','c6':'-1','c19':'yyyy-MM-dd','c4':'1','c7':'1','c12':'0','c9':'0','c11':'1','c13':'200','c30':'0'}">
							  					
									<ul class="news_list list2">
										
										<li class="news n1 clearfix">
											<span class="news_title"><a href='/2022/1214/c814a216777/page.htm' target='_blank' title='山东省分析测试中心组织开展国家乡村振兴岚皋县科技特派团中药材小组线上交流活动'>山东省分析测试中心组织开展国家乡村振兴岚皋县科技特派团中药材小组线上交流活动</a></span>
											<span class="news_meta">2022-12-16</span>
										</li>
										
										<li class="news n2 clearfix">
											<span class="news_title"><a href='/2022/1214/c814a216775/page.htm' target='_blank' title='马克思主义学院举行党的二十大精神宣讲报告会'>马克思主义学院举行党的二十大精神宣讲报告会</a></span>
											<span class="news_meta">2022-12-15</span>
										</li>
										
										<li class="news n3 clearfix">
											<span class="news_title"><a href='/2022/1213/c814a216751/page.htm' target='_blank' title='校（院）第二届“红星”党员先锋宣讲团成员走进体音学院宣讲党的二十大精神'>校（院）第二届“红星”党员先锋宣讲团成员走进体音学院宣讲党的二十大精神</a></span>
											<span class="news_meta">2022-12-13</span>
										</li>
										
										<li class="news n4 clearfix">
											<span class="news_title"><a href='/2022/1212/c814a216710/page.htm' target='_blank' title='分析测试中心科研人员参加山东省中药材行业协会第二次会员代表大会'>分析测试中心科研人员参加山东省中药材行业协会第二次会员代表大会</a></span>
											<span class="news_meta">2022-12-12</span>
										</li>
										
										<li class="news n5 clearfix">
											<span class="news_title"><a href='/2022/1212/c814a216709/page.htm' target='_blank' title='分析测试中心参与的国家重点研发计划项目通过综合绩效评价并获得优秀等级'>分析测试中心参与的国家重点研发计划项目通过综合绩效评价并获得优秀等级</a></span>
											<span class="news_meta">2022-12-12</span>
										</li>
										
										<li class="news n6 clearfix">
											<span class="news_title"><a href='/2022/1210/c814a216326/page.htm' target='_blank' title='校（院）成功举办第一届超能学霸笔记活动'>校（院）成功举办第一届超能学霸笔记活动</a></span>
											<span class="news_meta">2022-12-12</span>
										</li>
										
										<li class="news n7 clearfix">
											<span class="news_title"><a href='/2022/1211/c814a216348/page.htm' target='_blank' title='经济与管理学部“人际交往的艺术”线上心理健康教育主题讲座顺利举办'>经济与管理学部“人际交往的艺术”线上心理健康教育主题讲座顺利举办</a></span>
											<span class="news_meta">2022-12-12</span>
										</li>
										
										<li class="news n8 clearfix">
											<span class="news_title"><a href='/2022/1206/c814a216260/page.htm' target='_blank' title='中试基地获批省重点产业知识产权运营服务平台项目'>中试基地获批省重点产业知识产权运营服务平台项目</a></span>
											<span class="news_meta">2022-12-11</span>
										</li>
										
										<li class="news n9 clearfix">
											<span class="news_title"><a href='/2022/1209/c814a216322/page.htm' target='_blank' title='分析测试中心开展2022年度实验室运行内部审核工作'>分析测试中心开展2022年度实验室运行内部审核工作</a></span>
											<span class="news_meta">2022-12-10</span>
										</li>
										
										<li class="news n10 clearfix">
											<span class="news_title"><a href='/2022/1209/c814a216316/page.htm' target='_blank' title='体育与音乐学院召开学习贯彻党的二十大精神动员部署会'>体育与音乐学院召开学习贯彻党的二十大精神动员部署会</a></span>
											<span class="news_meta">2022-12-10</span>
										</li>
										
										<li class="news n11 clearfix">
											<span class="news_title"><a href='/2022/1208/c814a216294/page.htm' target='_blank' title='化药学部山东省自然科学基金重大基础研究项目顺利启动'>化药学部山东省自然科学基金重大基础研究项目顺利启动</a></span>
											<span class="news_meta">2022-12-10</span>
										</li>
										
										<li class="news n12 clearfix">
											<span class="news_title"><a href='/2022/1207/c814a216276/page.htm' target='_blank' title='新一代技术标准化研究院牵头的国际标准 ISO/IEC TR 20169《信息技术 智慧城市标准化概述》获批立项'>新一代技术标准化研究院牵头的国际标准 ISO/IEC TR 20169《信息技术 智慧城市标准化概述》获批立项</a></span>
											<span class="news_meta">2022-12-08</span>
										</li>
										
										<li class="news n13 clearfix">
											<span class="news_title"><a href='/2022/1208/c814a216296/page.htm' target='_blank' title='材料科学与工程学部为考研学子送上暖心礼包'>材料科学与工程学部为考研学子送上暖心礼包</a></span>
											<span class="news_meta">2022-12-08</span>
										</li>
										
										<li class="news n14 clearfix">
											<span class="news_title"><a href='/2022/1205/c814a216251/page.htm' target='_blank' title='知名公益人宋娟做客工大（科院）青衿学堂'>知名公益人宋娟做客工大（科院）青衿学堂</a></span>
											<span class="news_meta">2022-12-08</span>
										</li>
										
									</ul>
							  
 <div id="wp_paging_w6"> 
<ul class="wp_paging clearfix"> 
     <li class="pages_count"> 
         <span class="per_page">每页&nbsp;<em class="per_count">14</em>&nbsp;记录&nbsp;</span> 
         <span class="all_count">总共&nbsp;<em class="all_count">12222</em>&nbsp;记录&nbsp;</span> 
     </li> 
     <li class="page_nav"> 
         <a class="first" href="javascript:void(0);" target="_self"><span>第一页</span></a> 
         <a class="prev" href="javascript:void(0);" target="_self"><span>&lt;&lt;上一页</span></a> 
         <a class="next" href="/814/list2.htm" target="_self"><span>下一页&gt;&gt;</span></a> 
         <a class="last" href="/814/list873.htm" target="_self"><span>尾页</span></a> 
     </li> 
     <li class="page_jump"> 
         <span class="pages">页码&nbsp;<em class="curr_page">1</em>/<em class="all_pages">873</em></span> 
         <span><input class="pageNum" type="text" /><input type="hidden" class="currPageURL" value=""></span></span> 
         <span><a class="pagingJump" href="javascript:void(0);" target="_self">跳转到&nbsp;</a></span> 
     </li> 
</ul> 
</div> 
<script type="text/javascript"> 
     $().ready(function() { 
         $("#wp_paging_w6 .pagingJump").click(function() { 
             var pageNum = $("#wp_paging_w6 .pageNum").val(); 
             if (pageNum === "") { alert('请输入页码！'); return; } 
             if (isNaN(pageNum) || pageNum <= 0 || pageNum > 873) { alert('请输入正确页码！'); return; } 
             var reg = new RegExp("/list", "g"); 
             var url = "/814/list.htm"; 
             window.location.href = url.replace(reg, "/list" + pageNum); 
         }); 
     }); 
</script> 

							</div>
						</div>
					</div>
				</div>
			</div>
			<div class="clear"></div>
		</div>
	</div>
</div>
<!--End||content-->
<!--Start||footer-->
<div class="wrapper footer" id="footer">
	<div class="inner">
		<div class="mod clearfix">
			<div class="foot-left" frag="窗口90" portletmode="simpleSiteAttri">
				<p class="copyright"><span>版权所有©:齐鲁工业大学  山东省济南市长清区大学路3501号 邮编250353　鲁ICP备05046217号</span></p>
			</div>
			<div class="foot-right">
				<div class="shares" frag="窗口91">
					
				</div>
				<div class="shares shares1" frag="窗口92" portletmode="simpleNews" configs="{'c8':'1','c42':'320','c25':'320','c30':'0','c29':'1','c23':'1','c34':'300','c20':'0','c31':'0','c16':'1','c3':'2','c2':'序号,标题,发布时间','c27':'480','c43':'0','c17':'0','c5':'_blank','c24':'240','c32':'','c26':'1','c37':'1','c28':'640','c40':'1','c15':'0','c14':'1','c44':'0','c33':'500','c10':'50','c18':'yyyy-MM-dd','c36':'0','c1':'1','c6':'15','c19':'yyyy-MM-dd','c21':'0','c4':'1','c35':'-1:-1','c39':'300','c38':'100','c7':'1','c12':'0','c9':'0','c11':'1','c13':'200','c41':'240'}" contents="{'c2':'0', 'c1':'/微博微信'}">
					
						<ul>
							
							<li class="i1">
								<a href="https://weibo.com/3863948712/profile?is_hot=1" target="_blank"></a>
								<div class="con"></div>
							</li>
						   
							<li class="i2">
								<a href="http://javavoid(0)" target="_blank"></a>
								<div class="con"><a href='http://javavoid(0)' target='_blank' title='微信'><img src='/_upload/article/images/d8/a0/0cbc1a6e422a91e60b7d60c1bc49/c037be6c-822f-400e-bc7a-64890f86fd30_s.png' width='320' /></a></div>
							</li>
						   
						</ul>
					
				</div>
			</div>
		</div>
	</div>
</div>
<!--End||footer-->
</body>
<script type="text/javascript" src="/_upload/tpl/07/29/1833/template1833/js/comcus.js"></script>
<script type="text/javascript" src="/_upload/tpl/07/29/1833/template1833/js/list.js"></script>
<script type="text/javascript" src="/_upload/tpl/07/29/1833/template1833/js/app.js"></script>
<script type="text/javascript">
$(function(){
	// 初始化SDAPP
	new SDAPP({
		"menu":{
			type:"slide,aside"
		}
	});
});
</script>
</html>

 <img src="/_visitcount?siteId=20&type=2&columnId=814" style="display:none" width="0" height="0"></image>
`

func TestPageFind(t *testing.T) {

	type args struct {
		source string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "1", args: args{a}, want: []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PageFind([]byte(tt.args.source)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PageFind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_absUrl(t *testing.T) {
	str := "https://www.tsinghua.edu.cn/jyjx/bksjy.htm"
	s := "../../../../../2.html"
	s1 := "./../2.html"
	s2 := "/2.html"
	s3 := "./2.html"
	s4 := "https://xxx.com/2.html"
	println(absUrl(s, str))  // https://xxx.com/2.html
	println(absUrl(s1, str)) // https://xxx.com/articles/2.html
	println(absUrl(s2, str)) // https://xxx.com/2.html
	println(absUrl(s3, str)) // https://xxx.com/articles/2876/2.html
	println(absUrl(s4, str)) // https://xxx.com/2.html
}

func TestGettitle(t *testing.T) {
	s, _ := simplifiedchinese.GBK.NewEncoder().String("安大大")
	get, _ := http.Get("https://www.sogou.com/")
	all, _ := io.ReadAll(get.Body)
	type args struct {
		body string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{body: string(all)}},
		{name: "test", args: args{body: "<title>" + s + "</title>"},
			want: "齐鲁工业大学录取咨询平台"},
		{name: "2", args: args{body: "<title>Annual Reviews&#x4f7f;&#x7528;&#x6307;&#x5357;&#x8bed;&#x97f3;&#x8bfe;&#x4ef6;.mp4</title><link rel=\"alternate\" type=\"application/json+oembed\" href=\"https://workdrive.zohopublic.com.cn/services/oembed?type=json&url=https%3A%2F%2Fworkdrive.zohopublic.com.cn%2Fexternal%2FGDy6ryIAG5-2APbh\" title=\"Zoho WorkDrive\"/>\n<base href=\"/\" /> "}, want: "吾爱破解 - LCG - LSG|安卓破解|病毒分析|www.52pojie.cn "},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Gettitle([]byte(tt.args.body)); got != tt.want {
				t.Errorf("Gettitle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindurl(t *testing.T) {
	type args struct {
		body []byte
		Murl string
	}
	tests := []struct {
		name string
		args args
		want [][]*url.URL
	}{
		{
			name: "findurl", args: args{
				body: []byte(a),
				Murl: "https://news.qlu.edu.cn/814/list.htm",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Findurl(tt.args.body, tt.args.Murl); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Findurl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_absUrl1(t *testing.T) {
	type args struct {
		currUrl string
		baseUrl string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				currUrl: "../item?id=1",
				baseUrl: "http://xyh.qlu.edu.cn/site/sdili_xyh/xyh/search/?key=",
			},
			want: "http://xyh.qlu.edu.cn/site/sdili_xyh/xyh/item?id=1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := absUrl(tt.args.currUrl, tt.args.baseUrl); got != tt.want {
				t.Errorf("absUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isGBK(t *testing.T) {
	//s, _ := simplifiedchinese.GBK.NewEncoder().String("大萨达")
	s, _ := simplifiedchinese.HZGB2312.NewEncoder().String("大萨达萨达")
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "dsa",
			args: args{data: []byte(s)},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isGBK(tt.args.data); got != tt.want {
				t.Errorf("isGBK() = %v, want %v", got, tt.want)
			}
		})
	}
}
