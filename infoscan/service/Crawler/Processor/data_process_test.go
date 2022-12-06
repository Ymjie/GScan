package Processor

import (
	"net/url"
	"reflect"
	"testing"
)

var a string = `<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width,user-scalable=0,initial-scale=1.0, minimum-scale=1.0, maximum-scale=1.0"/>
<meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
<title>2019 Qilu University of Technology (Shandong Academy of Sciences)Admissions Guide for International Students in China</title>
<meta name="keywords" content="">
<meta name="description" content="" >
<meta name="description" content=" I.Target of enrollment1.Bachelor degree or above, a non-Chinese citizen with a postgraduate qualification and a healthy body.2.Immigrants from mainland China (inland), Hong Kong, Macau and Taiwan, as applicants for international students, must have valid foreign passport or nationality documents for 4 years or more, and the last four years (up to the school year) There is a record of actually living abroad for more than 2 years (before March 1st) (the actual residence abroad for 9 months in a y" />

<link type="text/css" href="/_css/_system/system.css" rel="stylesheet"/>
<link type="text/css" href="/_upload/site/1/style/62/62.css" rel="stylesheet"/>
<link type="text/css" href="/_upload/site/00/15/21/style/244/244.css" rel="stylesheet"/>
<link type="text/css" href="/_js/_portletPlugs/simpleNews/css/simplenews.css" rel="stylesheet" />
<link type="text/css" href="/_js/_portletPlugs/sudyNavi/css/sudyNav.css" rel="stylesheet" />
<link type="text/css" href="/_js/_portletPlugs/datepicker/css/datepicker.css" rel="stylesheet" />

<script language="javascript" src="/_js/jquery.min.js" sudy-wp-context="" sudy-wp-siteId="21"></script>
<script language="javascript" src="/_js/jquery.sudy.wp.visitcount.js"></script>
<script type="text/javascript" src="/_js/_portletPlugs/sudyNavi/jquery.sudyNav.js"></script>
<script type="text/javascript" src="/_js/_portletPlugs/datepicker/js/jquery.datepicker.js"></script>
<script type="text/javascript" src="/_js/_portletPlugs/datepicker/js/datepicker_lang_HK.js"></script>
<script type="text/javascript" src="/_upload/tpl/07/9f/1951/template1951/extends/extends.js"></script>
<link href="/_upload/tpl/07/9f/1951/template1951/css/base.css" rel="stylesheet">
<link href="/_upload/tpl/07/9f/1951/template1951/css/media.css" rel="stylesheet">
<!--[if lt IE 9]>
	<script src="/_upload/tpl/07/9f/1951/template1951/extends/libs/html5.js"></script>
<![endif]-->
</head>
<body class="wp-info-page">
<!--头部开始--><header class="wp-wrapper wp-header">
  <div class="wp-inner logo clearfix"> 
    <!--logo开始-->
    <div class="wp-panel logo-panel panel-1" frag="面板1"> <a class="navi-aside-toggle"></a>
      <div class="wp-window logo-window window-1" frag="窗口1" portletmode="simpleSiteAttri">
        
		<div style=""><a class="site-url" href="/main.htm"><img border='0' src='/_upload/site/00/15/21/logo.png' /></a></div>	
      <div class="site-name">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;国际合作处（国际研究生院）</div>
	  
	  </div>
    </div>
    <!--//logo结束--> 

    <!--搜索开始-->
    <div class="wp-panel search-panel panel-3" frag="面板3">
		<div class="site-lang clearfix"  >
		            <ul class="clearfix">
		              <li class="links i1"><a href="http://www.qlu.edu.cn/" target="_blank">学校首页</a> </li>
		              <li class="links i2"><a href="http://os.qlu.edu.cn" target="_blank">一网通办</a> </li>
		              <li class="links i3"><a href="http://www.qlu.edu.cn/5141/list.htm" target="_blank">信息系统</a> </li>
		               </ul>
		        </div>
      <div class="wp-window search-window window-3" frag="窗口3" portletmode="search">
                 <!--搜索组件-->
          <div class="wp-search clearfix">
            <form method="post" action="/_web/search/doSearch.do?locale=zh_CN&request_locale=zh_CN&_p=YXM9MjEmdD0xOTUxJmQ9NjY0NiZwPTMmZj04NDQmYT0wJm09U04mfGJubkNvbHVtblZpcnR1YWxOYW1lPTg0NCY_" target="_blank">
              <div class="search-input">
                <input class="search-title" name="keyword" type="text" placeholder="请输入关键字">
              </div>
              <div class="search-btn">
                <input class="search-submit" name="submit" type="submit" value="">
              </div>
            </form>
          </div>
          <!--//复制以上代码到自定义搜索--> 
      </div>
    </div>
    <!--//搜索结束--> 

    <!--校训-->
	<div class="header_xx"><img src="/_upload/tpl/07/9f/1951/template1951/images/banner_xx.jpg"></div>
	<!--//校训-->
  </div>
</header>
<!--//头部结束--> 

<!--导航开始-->
<nav class="wp-wrapper wp-navi">
  <div class="wp-inner">
    <div class="wp-panel main-nav-panel panel-5" frag="面板5">
      <div class="wp-window main-nav-window window-5" frag="窗口5" portletmode="simpleSudyNavi" contents="{'c2':'0', 'c1':'/机构介绍,/合作与交流,/专家工作,/因公出国（境）,/留学工大,/规章制度,/下载专区'}">
        
          <div class="navi-slide-head">
            <h3 class="navi-slide-title">导航</h3>
            <a class="navi-slide-arrow"></a> </div>
          
          <ul class="wp-menu clearfix" data-nav-aside='{"title":"导航","index":0}'>
            
            <li class="menu-item i1"><a class="menu-link" href="/6610/list.htm" target="_self">机构介绍</a> 
               
            </li>
            
            <li class="menu-item i2"><a class="menu-link" href="/6611/list.htm" target="_self">合作与交流</a> 
               
            </li>
            
            <li class="menu-item i3"><a class="menu-link" href="/6612/list.htm" target="_self">专家工作</a> 
               
              <i class="menu-switch-arrow"></i>
              <ul class="sub-menu clearfix">
                
                <li class="sub-item i3-1"><a class="sub-link" href="/zjpr/list.htm" target="_self">专家聘任</a> 
                   
                </li>
                
                <li class="sub-item i3-2"><a class="sub-link" href="/yzxm/list.htm" target="_self">引智项目</a> 
                   
                </li>
                
                <li class="sub-item i3-3"><a class="sub-link" href="/zcgd/list.htm" target="_self">政策规定</a> 
                   
                </li>
                
              </ul>
               
            </li>
            
            <li class="menu-item i4"><a class="menu-link" href="/6613/list.htm" target="_self">因公出国（境）</a> 
               
            </li>
            
            <li class="menu-item i5"><a class="menu-link" href="/6614/list.htm" target="_self">留学工大</a> 
               
            </li>
            
            <li class="menu-item i6"><a class="menu-link" href="/6615/list.htm" target="_self">规章制度</a> 
               
            </li>
            
            <li class="menu-item i7"><a class="menu-link" href="/6616/list.htm" target="_self">下载专区</a> 
               
            </li>
            
          </ul>
           
        
      </div>
    </div>
  </div>
</nav>

<!--aside导航-->
<div class="wp-navi-aside" id="wp-navi-aside">
  <div class="aside-inner">
    <div class="navi-aside-wrap"></div>
  </div>
  <div class="navi-aside-mask"></div>
</div>
<!--//导航结束--> 
<!--栏目图片开始-->
<div class="wp-wrapper wp-banner" >
  <div class="wp-inner list_bg">
      <div class="banner"><img src="/_upload/tpl/07/9f/1951/template1951/images/banner.jpg"></div>
  </div>
</div>
<!--//栏目图片结束-->
<!--主体开始-->
<div class="wp-wrapper wp-container">
  <div class="wp-wrapper wp-banner list_bg" >
    <div class="wp-inner" >
      <div class="banner"></div>
    </div>
  </div>
  <div class="wp-inner clearfix">
    <div class="info-box" frag="面板6">
      <div class="article" frag="窗口6" portletmode="simpleArticleAttri">
        
          <h1 class="arti-title">2019 Qilu University of Technology (Shandong Academy of Sciences)Admissions Guide for International Students in China</h1>
          <p class="arti-metas"><span class="arti-update">发布时间：2019-05-30</span><span class="arti-views">作者：</span><a class="read-options" id="read-options" href="javascript:void(0)">设置</a></p>
          <div class="read-setting" id="read-setting">
            <p class="setting-item setting-font"> <a class="larger-font" id="larger-font" href="javascript:void(0)">A+</a> <a class="smaller-font" id="smaller-font" href="javascript:void(0)">A-</a> </p>
            <p class="setting-item setting-model"> <a class="read-model" id="read-model" href="javascript:void(0)">夜晚模式</a> </p>
          </div>
          <div class="entry">
            <article class="read"><div class='wp_articlecontent'><p><style type="text/css"> 	 	</style><p class="cjk" style="text-align:center;margin-bottom:0cm;"><span style="font-family:times new roman, serif;"><span lang="en-US"><span style="color:#00000a;font-family:times new roman, serif;font-size:20px;"><span lang="en-US"><strong><font color="#000000" face="宋体"></font><br /></strong></span></span></span></span></p><p><span style="color:#00000a;font-family:times new roman, serif;font-size:20px;"><span lang="en-US">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; I.Target of enrollment</span></span></p><p><br /></p><p class="cjk" style="text-align:justify;line-height:0.88cm;text-indent:0.99cm;page-break-after:auto;orphans:0;widows:0;page-break-inside:auto;margin-bottom:0cm;"><span style="color:#00000a;font-family:times new roman, serif;font-size:18px;"><span lang="en-US"><span style="font-weight:normal;">1.Bachelor degree or above, a non-Chinese citizen with a postgraduate qualification and a healthy body.</span></span></span></p><p class="cjk" style="text-align:justify;line-height:0.88cm;text-indent:0.99cm;page-break-after:auto;orphans:0;widows:0;page-break-inside:auto;margin-bottom:0cm;"><span style="color:#00000a;font-family:times new roman, serif;font-size:18px;"><span lang="en-US"><span style="font-weight:normal;">2.Immigrants from mainland China (inland), Hong Kong, Macau and Taiwan, as applicants for international students, must have valid foreign passport or nationality documents for 4 years or more, and the last four years (up to the school year) There is a record of actually living abroad for more than 2 years (before March 1st) (the actual residence abroad for 9 months in a year can be calculated on a one-year basis, subject to the entry and exit signatures).</span></span></span></p><p class="cjk" style="text-align:justify;line-height:0.88cm;text-indent:1.06cm;page-break-after:auto;orphans:0;widows:0;page-break-inside:auto;margin-bottom:0cm;"><span style="color:#00000a;font-family:times new roman, serif;font-size:20px;"><span lang="en-US">II.Application time</span></span></p><p class="cjk" style="text-align:justify;line-height:0.88cm;text-indent:0.99cm;page-break-after:auto;orphans:0;widows:0;page-break-inside:auto;margin-bottom:0cm;"><span style="color:#00000a;font-family:times new roman, serif;font-size:18px;"><span lang="en-US"><span style="font-weight:normal;">March 1st - May 31st</span></span></span></p><p class="cjk" style="text-align:justify;line-height:0.88cm;text-indent:1.06cm;page-break-after:auto;orphans:0;widows:0;page-break-inside:auto;margin-bottom:0cm;"><span style="color:#00000a;font-family:times new roman, serif;font-size:20px;"><span lang="en-US">III.Admissions profession</span></span></p><p class="cjk" style="text-align:justify;line-height:0.88cm;text-indent:0.99cm;page-break-after:auto;orphans:0;widows:0;page-break-inside:auto;margin-bottom:0cm;"><span style="color:#00000a;font-family:times new roman, serif;font-size:18px;"><span lang="en-US"><span style="font-weight:normal;">See attached</span></span></span></p><p class="cjk" style="text-align:justify;line-height:0.88cm;text-indent:1.06cm;page-break-after:auto;orphans:0;widows:0;page-break-inside:auto;margin-bottom:0cm;"><span style="color:#00000a;font-family:times new roman, serif;font-size:20px;"><span lang="en-US">IV.Qualification required</span></span></p><p class="cjk" style="text-align:justify;line-height:0.88cm;text-indent:1.06cm;page-break-after:auto;orphans:0;widows:0;page-break-inside:auto;margin-bottom:0cm;"><span style="color:#00000a;font-family:times new roman, serif;font-size:20px;"><span lang="en-US">1.Academic requirements</span></span><span style="color:#00000a;font-family:黑体;font-size:20px;"><span lang="zh-CN">：</span></span></p><p class="cjk" style="text-align:justify;line-height:0.88cm;text-indent:0.99cm;page-break-after:auto;orphans:0;widows:0;page-break-inside:auto;margin-bottom:0cm;"><span style="color:#00000a;font-family:times new roman, serif;font-size:18px;"><span lang="en-US"><span style="font-weight:normal;">Entry to a master&#39;s degree requires a bachelor&#39;s degree or equivalent (see “Successful completion of the International Standard Classification of Education (ISCED 2011) Level 6 or Level 7 course).</span></span></span></p><p class="cjk" style="text-align:justify;line-height:0.88cm;text-indent:1.06cm;page-break-after:auto;orphans:0;widows:0;page-break-inside:auto;margin-bottom:0cm;"><span style="color:#00000a;font-family:times new roman, serif;font-size:20px;"><span lang="en-US">2.Language requirements</span></span><span style="color:#00000a;font-family:黑体;font-size:20px;"><span lang="zh-CN">：</span></span></p><p class="cjk" style="text-align:justify;line-height:0.88cm;text-indent:0.99cm;page-break-after:auto;orphans:0;widows:0;page-break-inside:auto;margin-bottom:0cm;"><span style="color:#00000a;font-family:times new roman, serif;font-size:18px;"><span lang="en-US"><span style="font-weight:normal;">The requirements for Chinese language and professional Chinese language proficiency should be at least four levels of the International Chinese Proficiency Standard. For subjects and majors whose English is the language of instruction, the native language is English, and no language certificate is required. If the mother tongue is not English, the English level should be at least IELTS 5.5, 70% new TOEFL or other equivalent English proficiency test materials. Applicants who have obtained the highest degree and degree through English teaching must submit an English teaching certificate.</span></span></span></p><p class="cjk" style="text-align:justify;line-height:0.88cm;text-indent:1.06cm;page-break-after:auto;orphans:0;widows:0;page-break-inside:auto;margin-bottom:0cm;"><span style="color:#00000a;font-family:times new roman, serif;font-size:20px;"><span lang="en-US">3.Age requirement</span></span></p><p class="cjk" style="text-align:justify;line-height:0.88cm;text-indent:0.99cm;page-break-after:auto;orphans:0;widows:0;page-break-inside:auto;margin-bottom:0cm;"><span style="color:#00000a;font-family:times new roman, serif;font-size:18px;"><span lang="en-US"><span style="font-weight:normal;">Postgraduate entry requirements are 16 years old or older and 35 years old or younger</span></span></span></p><p class="cjk" style="text-align:justify;line-height:0.88cm;text-indent:0.99cm;page-break-after:auto;orphans:0;widows:0;page-break-inside:auto;margin-bottom:0cm;"><span style="color:#00000a;font-family:times new roman, serif;font-size:18px;"><span lang="en-US"><span style="font-weight:normal;">Those under the age of 18 must be sure that their parents are resident in China; if their parents are not resident in China, they must submit a notary certificate issued by the Jinan Notary Office.</span></span></span></p><p class="cjk" style="text-align:justify;line-height:0.88cm;text-indent:1.06cm;page-break-after:auto;orphans:0;widows:0;page-break-inside:auto;margin-bottom:0cm;"><span style="color:#00000a;font-family:times new roman, serif;font-size:20px;"><span lang="en-US">4.Economic requirements</span></span></p><p class="cjk" style="text-align:justify;line-height:0.88cm;text-indent:0.99cm;page-break-after:auto;orphans:0;widows:0;page-break-inside:auto;margin-bottom:0cm;"><span style="color:#00000a;font-family:times new roman, serif;font-size:18px;"><span lang="en-US"><span style="font-weight:normal;">Economic certificate of more than 20,000 RMB (scholarship students may pay less or exempt according to specific circumstances)</span></span></span></p><p class="cjk" style="text-align:justify;line-height:0.88cm;text-indent:1.06cm;page-break-after:auto;orphans:0;widows:0;page-break-inside:auto;margin-bottom:0cm;"><span style="color:#00000a;font-family:times new roman, serif;font-size:20px;"><span lang="en-US">V.Application materials</span></span></p><p class="cjk" style="text-align:justify;line-height:0.88cm;text-indent:0.99cm;page-break-after:auto;orphans:0;widows:0;page-break-inside:auto;margin-bottom:0cm;"><span style="color:#00000a;"><span style="font-family:仿宋_gb2312;font-size:18px;"><span lang="zh-CN"><span style="font-weight:normal;">（</span></span></span><span style="font-family:times new roman, serif;font-size:18px;"><span lang="en-US"><span style="font-weight:normal;">1</span></span></span></span><span style="color:#00000a;"><span style="font-family:仿宋_gb2312;font-size:18px;"><span lang="zh-CN"><span style="font-weight:normal;">）</span></span></span><span style="font-family:times new roman, serif;font-size:18px;"><span lang="en-US"><span style="font-weight:normal;">Application form for foreigners to study at Qilu University of technology (Shandong Academy of Sciences)</span></span></span></span></p><p class="cjk" style="text-align:justify;line-height:0.88cm;text-indent:0.99cm;page-break-after:auto;orphans:0;widows:0;page-break-inside:auto;margin-bottom:0cm;"><span style="color:#00000a;"><span style="font-family:times new roman, serif;font-size:18px;"><span lang="en-US"><span style="font-weight:normal;"><img src="/_ueditor/themes/default/images/icon_doc.gif" /><a href="/_upload/article/files/7c/7e/ae49e522431e85de65411631080e/188c6762-13d8-4b46-a199-3fe8942a173e.doc" sudyfile-attr="{'title':'附件1：齐鲁工业大学（山东省科学院）来华留学生申请表.doc'}">附件1：Application form for foreigners to study at Qilu University of technology(Shandong Academy of Sciences).doc</a></span></span></span></span></p><p class="cjk" style="text-align:justify;line-height:0.88cm;text-indent:0.99cm;page-break-after:auto;orphans:0;widows:0;page-break-inside:auto;margin-bottom:0cm;"><span style="color:#00000a;"><span style="font-family:仿宋_gb2312;font-size:18px;"><span lang="zh-CN"><span style="font-weight:normal;">（</span></span></span><span style="font-family:times new roman, serif;font-size:18px;"><span lang="en-US"><span style="font-weight:normal;">2</span></span></span></span><span style="color:#00000a;"><span style="font-family:仿宋_gb2312;font-size:18px;"><span lang="zh-CN"><span style="font-weight:normal;">）</span></span></span><span style="font-family:times new roman, serif;font-size:18px;"><span lang="en-US"><span style="font-weight:normal;">Registration Form for International Students</span></span></span></span></p><p class="cjk" style="text-align:justify;line-height:0.88cm;text-indent:0.99cm;page-break-after:auto;orphans:0;widows:0;page-break-inside:auto;margin-bottom:0cm;"><span style="color:#00000a;"><span style="font-family:times new roman, serif;font-size:18px;"><span lang="en-US"><span style="font-weight:normal;"><img src="/_ueditor/themes/default/images/icon_doc.gif" /><a href="/_upload/article/files/7c/7e/ae49e522431e85de65411631080e/33589c35-9589-4a61-b2e7-50194d37f3a2.doc" sudyfile-attr="{'title':'附件2：来华留学生登记表.doc'}">附件2：Registration Form for International Students.doc</a></span></span></span></span></p><p class="cjk" style="text-align:justify;line-height:0.88cm;text-indent:0.99cm;page-break-after:auto;orphans:0;widows:0;page-break-inside:auto;margin-bottom:0cm;"><span style="color:#00000a;"><span style="font-family:仿宋_gb2312;font-size:18px;"><span lang="zh-CN"><span style="font-weight:normal;">（</span></span></span><span style="font-family:times new roman, serif;font-size:18px;"><span lang="en-US"><span style="font-weight:normal;">3</span></span></span></span><span style="color:#00000a;"><span style="font-family:仿宋_gb2312;font-size:18px;"><span lang="zh-CN"><span style="font-weight:normal;">）</span></span></span><span style="font-family:times new roman, serif;font-size:18px;"><span lang="en-US"><span style="font-weight:normal;">Accommodation application form </span></span></span></span></p><p class="cjk" style="text-align:justify;line-height:0.88cm;text-indent:0.99cm;page-break-after:auto;orphans:0;widows:0;page-break-inside:auto;margin-bottom:0cm;"><span style="color:#00000a;"><span style="font-family:times new roman, serif;font-size:18px;"><span lang="en-US"><span style="font-weight:normal;"><img src="/_ueditor/themes/default/images/icon_doc.gif" /><a href="/_upload/article/files/7c/7e/ae49e522431e85de65411631080e/bce576d0-c5b7-46b7-8f04-e723ee8e2b35.doc" sudyfile-attr="{'title':'附件3：留学生宿舍申请.doc'}">附件3：Accommodation application form.doc</a></span></span></span></span></p><p class="cjk" style="text-align:justify;line-height:0.88cm;text-indent:0.99cm;page-break-after:auto;orphans:0;widows:0;page-break-inside:auto;margin-bottom:0cm;"><span style="color:#00000a;"><span style="font-family:仿宋_gb2312;font-size:18px;"><span lang="zh-CN"><span style="font-weight:normal;">（</span></span></span><span style="font-family:times new roman, serif;font-size:18px;"><span lang="en-US"><span style="font-weight:normal;">4</span></span></span></span><span style="color:#00000a;"><span style="font-family:仿宋_gb2312;font-size:18px;"><span lang="zh-CN"><span style="font-weight:normal;">）</span></span></span><span style="font-family:times new roman, serif;font-size:18px;"><span lang="en-US"><span style="font-weight:normal;">Copy of passport (valid period)</span></span></span></span></p><p class="cjk" style="text-align:justify;line-height:0.88cm;text-indent:0.99cm;page-break-after:auto;orphans:0;widows:0;page-break-inside:auto;margin-bottom:0cm;"><span style="color:#00000a;"><span style="font-family:仿宋_gb2312;font-size:18px;"><span lang="zh-CN"><span style="font-weight:normal;">（</span></span></span><span style="font-family:times new roman, serif;font-size:18px;"><span lang="en-US"><span style="font-weight:normal;">5</span></span></span></span><span style="color:#00000a;"><span style="font-family:仿宋_gb2312;font-size:18px;"><span lang="zh-CN"><span style="font-weight:normal;">）</span></span></span><span style="font-family:times new roman, serif;font-size:18px;"><span lang="en-US"><span style="font-weight:normal;">A copy of the undergraduate diploma or degree certificate (original notarized documents issued by the Embassy or Consulate of the People&#39;s Republic of China) or pre-graduation certificate (for students in school)</span></span></span></span></p><p class="cjk" style="text-align:justify;line-height:0.88cm;text-indent:0.99cm;page-break-after:auto;orphans:0;widows:0;page-break-inside:auto;margin-bottom:0cm;"><span style="color:#00000a;"><span style="font-family:仿宋_gb2312;font-size:18px;"><span lang="zh-CN"><span style="font-weight:normal;">（</span></span></span><span style="font-family:times new roman, serif;font-size:18px;"><span lang="en-US"><span style="font-weight:normal;">6</span></span></span></span><span style="color:#00000a;"><span style="font-family:仿宋_gb2312;font-size:18px;"><span lang="zh-CN"><span style="font-weight:normal;">）</span></span></span><span style="font-family:times new roman, serif;font-size:18px;"><span lang="en-US"><span style="font-weight:normal;">Copy of all course transcripts of the undergraduate course (original notarized documents issued by the Embassy or Consulate of the People&#39;s Republic of China)</span></span></span></span></p><p class="cjk" style="text-align:justify;line-height:0.88cm;text-indent:0.99cm;page-break-after:auto;orphans:0;widows:0;page-break-inside:auto;margin-bottom:0cm;"><span style="color:#00000a;"><span style="font-family:仿宋_gb2312;font-size:18px;"><span lang="zh-CN"><span style="font-weight:normal;">（</span></span></span><span style="font-family:times new roman, serif;font-size:18px;"><span lang="en-US"><span style="font-weight:normal;">7</span></span></span></span><span style="color:#00000a;"><span style="font-family:仿宋_gb2312;font-size:18px;"><span lang="zh-CN"><span style="font-weight:normal;">）</span></span></span><span style="font-family:times new roman, serif;font-size:18px;"><span lang="en-US"><span style="font-weight:normal;">Two original letters of recommendation from university </span></span></span></span></p><p class="cjk" style="text-align:justify;line-height:0.88cm;text-indent:0.99cm;page-break-after:auto;orphans:0;widows:0;page-break-inside:auto;margin-bottom:0cm;"><span style="color:#00000a;"><span style="font-family:仿宋_gb2312;font-size:18px;"><span lang="zh-CN"><span style="font-weight:normal;">（</span></span></span><span style="font-family:times new roman, serif;font-size:18px;"><span lang="en-US"><span style="font-weight:normal;">8</span></span></span></span><span style="color:#00000a;"><span style="font-family:仿宋_gb2312;font-size:18px;"><span lang="zh-CN"><span style="font-weight:normal;">）</span></span></span><span style="font-family:times new roman, serif;font-size:18px;"><span lang="en-US"><span style="font-weight:normal;">Study Plan in China (written in Chinese or English, no less than 400 words)</span></span></span></span></p><p class="cjk" style="text-align:justify;line-height:0.88cm;text-indent:0.99cm;page-break-after:auto;orphans:0;widows:0;page-break-inside:auto;margin-bottom:0cm;"><span style="color:#00000a;"><span style="font-family:仿宋_gb2312;font-size:18px;"><span lang="zh-CN"><span style="font-weight:normal;">（</span></span></span><span style="font-family:times new roman, serif;font-size:18px;"><span lang="en-US"><span style="font-weight:normal;">9</span></span></span></span><span style="color:#00000a;"><span style="font-family:仿宋_gb2312;font-size:18px;"><span lang="zh-CN"><span style="font-weight:normal;">）</span></span></span><span style="font-family:times new roman, serif;font-size:18px;"><span lang="en-US"><span style="font-weight:normal;">English/Chinese proficiency certificate</span></span></span></span></p><p class="cjk" style="text-align:justify;line-height:0.88cm;text-indent:0.99cm;page-break-after:auto;orphans:0;widows:0;page-break-inside:auto;margin-bottom:0cm;"><span style="color:#00000a;"><span style="font-family:仿宋_gb2312;font-size:18px;"><span lang="zh-CN"><span style="font-weight:normal;">（</span></span></span><span style="font-family:times new roman, serif;font-size:18px;"><span lang="en-US"><span style="font-weight:normal;">10</span></span></span></span><span style="color:#00000a;"><span style="font-family:仿宋_gb2312;font-size:18px;"><span lang="zh-CN"><span style="font-weight:normal;">） </span></span></span><span style="font-family:times new roman, serif;font-size:18px;"><span lang="en-US"><span style="font-weight:normal;">Financial certificate</span></span></span></span></p><p class="cjk" style="text-align:justify;line-height:0.88cm;text-indent:0.99cm;page-break-after:auto;orphans:0;widows:0;page-break-inside:auto;margin-bottom:0cm;"><span style="color:#00000a;"><span style="font-family:仿宋_gb2312;font-size:18px;"><span lang="zh-CN"><span style="font-weight:normal;">（</span></span></span><span style="font-family:times new roman, serif;font-size:18px;"><span lang="en-US"><span style="font-weight:normal;">11</span></span></span></span><span style="color:#00000a;"><span style="font-family:仿宋_gb2312;font-size:18px;"><span lang="zh-CN"><span style="font-weight:normal;">）</span></span></span><span style="font-family:times new roman, serif;font-size:18px;"><span lang="en-US"><span style="font-weight:normal;">Non-criminal Record</span></span></span></span></p><p class="cjk" style="text-align:justify;line-height:0.88cm;text-indent:0.99cm;page-break-after:auto;orphans:0;widows:0;page-break-inside:auto;margin-bottom:0cm;"><span style="color:#00000a;"><span style="font-family:仿宋_gb2312;font-size:18px;"><span lang="zh-CN"><span style="font-weight:normal;">（</span></span></span><span style="font-family:times new roman, serif;font-size:18px;"><span lang="en-US"><span style="font-weight:normal;">12</span></span></span></span><span style="color:#00000a;"><span style="font-family:仿宋_gb2312;font-size:18px;"><span lang="zh-CN"><span style="font-weight:normal;">）</span></span></span><span style="font-family:times new roman, serif;font-size:18px;"><span lang="en-US"><span style="font-weight:normal;">Foreign medical examination records (within 6 months)</span></span></span></span></p><p class="cjk" style="text-align:justify;line-height:0.88cm;text-indent:0.99cm;page-break-after:auto;orphans:0;widows:0;page-break-inside:auto;margin-bottom:0cm;"><span style="color:#00000a;"><span style="font-family:仿宋_gb2312;font-size:18px;"><span lang="zh-CN"><span style="font-weight:normal;">（</span></span></span><span style="font-family:times new roman, serif;font-size:18px;"><span lang="en-US"><span style="font-weight:normal;">13</span></span></span></span><span style="color:#00000a;"><span style="font-family:仿宋_gb2312;font-size:18px;"><span lang="zh-CN"><span style="font-weight:normal;">）</span></span></span><span style="font-family:times new roman, serif;font-size:18px;"><span lang="en-US"><span style="font-weight:normal;">Passport photos (electronic version and 4 paper photos signed on the back)</span></span></span></span></p><p class="cjk" style="text-align:justify;line-height:0.88cm;text-indent:1.06cm;page-break-after:auto;orphans:0;widows:0;page-break-inside:auto;margin-bottom:0cm;"><span style="color:#00000a;font-family:times new roman, serif;font-size:20px;"><span lang="en-US">VI.Application procedures</span></span></p><p class="cjk" style="text-align:justify;line-height:0.88cm;text-indent:0.99cm;page-break-after:auto;orphans:0;widows:0;page-break-inside:auto;margin-bottom:0cm;"><span style="color:#00000a;font-family:times new roman, serif;font-size:18px;"><span lang="en-US"><span style="font-weight:normal;">1. The applicant intends to download the application materials through the website (http://international.qlu.edu.cn/) and fill in as required;</span></span></span></p><p class="cjk" style="text-align:justify;line-height:0.88cm;text-indent:0.99cm;page-break-after:auto;orphans:0;widows:0;page-break-inside:auto;margin-bottom:0cm;"><span style="color:#00000a;font-family:times new roman, serif;font-size:18px;"><span lang="en-US"><span style="font-weight:normal;">2. Send the application materials and relevant certification materials to the International Cooperation Office (International Education Institute): guojichu@hotmail.com;</span></span></span></p><p class="cjk" style="text-align:justify;line-height:0.88cm;text-indent:0.99cm;page-break-after:auto;orphans:0;widows:0;page-break-inside:auto;margin-bottom:0cm;"><span style="color:#00000a;font-family:times new roman, serif;font-size:18px;"><span lang="en-US"><span style="font-weight:normal;">3. The International Cooperation Office (International Education Institute) organizes expert review materials and determines the admission list;</span></span></span></p><p class="cjk" style="text-align:justify;line-height:0.88cm;text-indent:0.99cm;page-break-after:auto;orphans:0;widows:0;page-break-inside:auto;margin-bottom:0cm;"><span style="color:#00000a;font-family:times new roman, serif;font-size:18px;"><span lang="en-US"><span style="font-weight:normal;">4. Inform the applicant in writing of the admission.</span></span></span></p><p class="cjk" style="text-align:justify;line-height:0.88cm;text-indent:1.06cm;page-break-after:auto;orphans:0;widows:0;page-break-inside:auto;margin-bottom:0cm;"><span style="color:#00000a;font-family:times new roman, serif;font-size:20px;"><span lang="en-US">VII.Admission time</span></span></p><p class="cjk" style="text-align:justify;line-height:0.88cm;text-indent:0.99cm;page-break-after:auto;orphans:0;widows:0;page-break-inside:auto;margin-bottom:0cm;"><span style="color:#00000a;font-family:times new roman, serif;font-size:18px;"><span lang="en-US"><span style="font-weight:normal;">September each year</span></span></span></p><p class="cjk" style="text-align:justify;line-height:0.88cm;text-indent:1.06cm;page-break-after:auto;orphans:0;widows:0;page-break-inside:auto;margin-bottom:0cm;"><span style="color:#00000a;font-family:times new roman, serif;font-size:20px;"><span lang="en-US">VIII.Tuition standard</span></span></p><p class="cjk" style="text-align:justify;line-height:0.88cm;text-indent:0.99cm;page-break-after:auto;orphans:0;widows:0;page-break-inside:auto;margin-bottom:0cm;"><span style="color:#00000a;font-family:times new roman, serif;font-size:18px;"><span lang="en-US"><span style="font-weight:normal;">Liberal arts class 25,000 yuan &nbsp;&nbsp;&nbsp;</span></span></span></p><p class="cjk" style="text-align:justify;line-height:0.88cm;text-indent:0.99cm;page-break-after:auto;orphans:0;widows:0;page-break-inside:auto;margin-bottom:0cm;"><span style="color:#00000a;font-family:times new roman, serif;font-size:18px;"><span lang="en-US"><span style="font-weight:normal;">Science and engineering class 30000 yuan &nbsp;&nbsp;&nbsp;</span></span></span></p><p class="cjk" style="text-align:justify;line-height:0.88cm;text-indent:0.99cm;page-break-after:auto;orphans:0;widows:0;page-break-inside:auto;margin-bottom:0cm;"><span style="color:#00000a;font-family:times new roman, serif;font-size:18px;"><span lang="en-US"><span style="font-weight:normal;">Art class 40000 yuan</span></span></span></p><p class="cjk" style="text-align:justify;line-height:0.88cm;text-indent:1.06cm;page-break-after:auto;orphans:0;widows:0;page-break-inside:auto;margin-bottom:0cm;"><span style="color:#00000a;font-family:times new roman, serif;font-size:20px;"><span lang="en-US">IX.Scholarship</span></span></p><p class="cjk" style="text-align:justify;line-height:0.88cm;text-indent:0.99cm;page-break-after:auto;orphans:0;widows:0;page-break-inside:auto;margin-bottom:0cm;"><span style="color:#00000a;font-family:times new roman, serif;font-size:18px;"><span lang="en-US"><span style="font-weight:normal;">International Student Scholarship 40,000 yuan per year</span></span></span></p><p class="cjk" style="text-align:justify;line-height:0.88cm;text-indent:0.99cm;page-break-after:auto;orphans:0;widows:0;page-break-inside:auto;margin-bottom:0cm;"><span style="color:#00000a;font-family:times new roman, serif;font-size:18px;"><span lang="en-US"><span style="font-weight:normal;">Shandong Provincial Government Foreign Student Scholarship 30,000 yuan per year</span></span></span></p><p class="cjk" style="text-align:justify;line-height:0.88cm;text-indent:0.99cm;page-break-after:auto;orphans:0;widows:0;page-break-inside:auto;margin-bottom:0cm;"><span style="color:#00000a;font-family:times new roman, serif;font-size:18px;"><span lang="en-US"><span style="font-weight:normal;">Applicants are required to complete the “Application Form for Scholarships for International Students of Qilu University of Technology (Shandong Academy of Sciences)”. The application procedure is the same as above.</span></span></span></p><p class="cjk" style="text-align:justify;line-height:0.88cm;text-indent:1.06cm;page-break-after:auto;orphans:0;widows:0;page-break-inside:auto;margin-bottom:0cm;"><span style="color:#00000a;font-family:times new roman, serif;font-size:20px;"><span lang="en-US">X.Contact information</span></span></p><p class="cjk" style="text-align:justify;line-height:0.88cm;text-indent:0.99cm;page-break-after:auto;orphans:0;widows:0;page-break-inside:auto;margin-bottom:0cm;"><span style="color:#00000a;font-family:times new roman, serif;font-size:18px;"><span lang="en-US"><span style="font-weight:normal;">Qilu University of Technology International Cooperation Office(International Education Institute)</span></span></span></p><p class="cjk" style="text-align:justify;line-height:0.88cm;text-indent:0.99cm;page-break-after:auto;orphans:0;widows:0;page-break-inside:auto;margin-bottom:0cm;"><span style="color:#00000a;font-family:times new roman, serif;font-size:18px;"><span lang="en-US"><span style="font-weight:normal;">Address: No. 3501, University Road, Changqing District, Jinan City, Shandong Province, China Post Code: 250353</span></span></span></p><p class="cjk" style="text-align:justify;line-height:0.88cm;text-indent:0.99cm;page-break-after:auto;orphans:0;widows:0;page-break-inside:auto;margin-bottom:0cm;"><span style="color:#00000a;font-family:times new roman, serif;font-size:18px;"><span lang="en-US"><span style="font-weight:normal;">Tel: 0531-89631559/89631016 Fax: 0531-89631016</span></span></span></p><ol class=" list-paddingleft-2" type="A" start="5"><li><p class="cjk" style="text-align:justify;line-height:0.88cm;page-break-after:auto;orphans:0;widows:0;page-break-inside:auto;margin-bottom:0cm;"><span style="color:#00000a;font-family:times new roman, serif;font-size:18px;"><span lang="en-US"><span style="font-weight:normal;">mail: guojichu@hotmail.com </span></span></span></p></li><li><p class="cjk" style="text-align:justify;line-height:0.88cm;page-break-after:auto;orphans:0;widows:0;page-break-inside:auto;margin-bottom:0cm;"><a name="_GoBack"></a><span style="color:#00000a;font-family:times new roman, serif;font-size:18px;"><span lang="en-US"><span style="font-weight:normal;">Website: http://international.qlu.edu.cn</span></span></span></p></li></ol><p class="cjk" style="margin-bottom:0cm;"><br /></p><p class="cjk" style="text-align:left;margin-bottom:0cm;"><span style="font-family:times new roman, serif;font-size:18px;"><span lang="en-US">Schedule: List of Program</span></span></p><p class="cjk" style="text-align:left;margin-bottom:0cm;"><br /></p><p class="cjk" style="margin-bottom:0cm;"><br /></p><center><table class="wp_editor_art_paste_table" cellspacing="0" cellpadding="0"><colgroup><col width="40" /><col width="162" /><col width="183" /><col width="179" /><tbody><tr><td height="5" bgcolor="#ffffff"><p class="cjk" style="text-align:center;page-break-after:auto;orphans:2;widows:2;page-break-inside:auto;"><span style="color:#000000;font-family:times new roman, serif;font-size:16px;"><span lang="en-US"><span style="font-style:normal;"><span style="text-decoration:none;"><strong>No.</strong></span></span></span></span></p></td><td bgcolor="#ffffff"><p class="cjk" style="text-align:center;page-break-after:auto;orphans:2;widows:2;page-break-inside:auto;"><span style="color:#000000;font-family:times new roman, serif;font-size:16px;text-decoration:none;"><span lang="en-US"><span style="font-style:normal;"><strong>Program</strong></span></span></span></p></td><td bgcolor="#ffffff"><p class="cjk" style="text-align:center;page-break-after:auto;orphans:2;widows:2;page-break-inside:auto;"><span style="color:#000000;"><span style="text-decoration:none;"></span></span><span style="font-family:times new roman, serif;"><span lang="en-US"><span style="color:#000000;font-family:times new roman, serif;font-size:16px;"><span lang="en-US"><span style="font-style:normal;"><span style="text-decoration:none;"><strong>S</strong></span></span></span></span><span style="color:#000000;font-family:times new roman, serif;font-size:16px;"><span lang="en-US"><span style="font-style:normal;"><span style="text-decoration:none;"><strong>chool</strong></span></span></span></span></span></span></p></td><td bgcolor="#ffffff"><p class="cjk" style="text-align:center;page-break-after:auto;orphans:2;widows:2;page-break-inside:auto;"><span style="color:#000000;font-family:times new roman, serif;font-size:16px;"><span lang="en-US"><span style="font-style:normal;"><span style="text-decoration:none;"><strong>Language of Instruction</strong></span></span></span></span></p></td></tr><tr><td height="20" bgcolor="#ffffff"><p class="cjk" style="text-align:center;page-break-after:auto;orphans:2;widows:2;page-break-inside:auto;"><span style="color:#000000;font-family:times new roman, serif;font-size:16px;"><span lang="en-US"><span style="font-style:normal;"><span style="text-decoration:none;">1</span></span></span></span></p></td><td bgcolor="#ffffff"><p class="cjk" style="text-align:left;page-break-after:auto;orphans:2;widows:2;page-break-inside:auto;"><span style="font-family:times new roman, serif;"><span lang="en-US"><span style="color:#000000;font-family:times new roman, serif;font-size:16px;"><span lang="en-US"><span style="font-style:normal;"><span style="text-decoration:none;">A</span></span></span></span><span style="color:#000000;font-family:times new roman, serif;font-size:16px;"><span style="font-style:normal;"><span style="text-decoration:none;">pplied </span></span></span><span style="color:#000000;font-family:times new roman, serif;font-size:16px;"><span lang="en-US"><span style="font-style:normal;"><span style="text-decoration:none;">E</span></span></span></span><span style="color:#000000;font-family:times new roman, serif;font-size:16px;"><span style="font-style:normal;"><span style="text-decoration:none;">conomics</span></span></span></span></span></p></td><td bgcolor="#ffffff"><p class="cjk" style="text-align:left;page-break-after:auto;orphans:2;widows:2;page-break-inside:auto;"><span style="color:#000000;font-family:times new roman, serif;font-size:16px;"><span lang="en-US"><span style="font-style:normal;"><span style="text-decoration:none;">School of Fiance</span></span></span></span></p></td><td bgcolor="#ffffff"><p class="cjk" style="text-align:center;page-break-after:auto;orphans:2;widows:2;page-break-inside:auto;"><span style="font-family:times new roman, serif;font-size:16px;"><span lang="en-US">Chinese or English</span></span></p></td></tr><tr><td height="20" bgcolor="#ffffff"><p class="cjk" style="text-align:center;page-break-after:auto;orphans:2;widows:2;page-break-inside:auto;"><span style="color:#000000;font-family:times new roman, serif;font-size:16px;"><span lang="en-US"><span style="font-style:normal;"><span style="text-decoration:none;">2</span></span></span></span></p></td><td bgcolor="#ffffff"><p class="cjk" style="text-align:left;page-break-after:auto;orphans:2;widows:2;page-break-inside:auto;"><span style="font-family:times new roman, serif;"><span lang="en-US"><span style="color:#000000;font-family:times new roman, serif;font-size:16px;"><span lang="en-US"><span style="font-style:normal;"><span style="text-decoration:none;">M</span></span></span></span><span style="color:#000000;font-family:times new roman, serif;font-size:16px;"><span style="font-style:normal;"><span style="text-decoration:none;">athematics</span></span></span></span></span></p></td><td bgcolor="#ffffff"><p class="cjk" style="text-align:left;page-break-after:auto;orphans:2;widows:2;page-break-inside:auto;"><span style="color:#000000;font-family:times new roman, serif;font-size:16px;"><span lang="en-US"><span style="font-style:normal;"><span style="text-decoration:none;">School of Mathematics and Statistics</span></span></span></span></p></td><td bgcolor="#ffffff"><p class="cjk" style="text-align:center;page-break-after:auto;orphans:2;widows:2;page-break-inside:auto;"><span style="font-family:times new roman, serif;font-size:16px;">Chinese or English</span></p></td></tr><tr><td height="20" bgcolor="#ffffff"><p class="cjk" style="text-align:center;page-break-after:auto;orphans:2;widows:2;page-break-inside:auto;"><span style="color:#000000;font-family:times new roman, serif;font-size:16px;"><span lang="en-US"><span style="font-style:normal;"><span style="text-decoration:none;">3</span></span></span></span></p></td><td bgcolor="#ffffff"><p class="cjk" style="text-align:left;page-break-after:auto;orphans:2;widows:2;page-break-inside:auto;"><span style="color:#000000;font-family:times new roman, serif;font-size:16px;"><span style="font-style:normal;"><span style="text-decoration:none;">Chemistry</span></span></span></p></td><td bgcolor="#ffffff"><p class="cjk" style="text-align:left;page-break-after:auto;orphans:2;widows:2;page-break-inside:auto;"><span style="color:#000000;font-family:times new roman, serif;font-size:16px;"><span lang="en-US"><span style="font-style:normal;"><span style="text-decoration:none;">School of Chemical and Pharmaceutical Engineering</span></span></span></span></p></td><td bgcolor="#ffffff"><p class="cjk" style="text-align:center;page-break-after:auto;orphans:2;widows:2;page-break-inside:auto;"><span style="font-family:times new roman, serif;font-size:16px;">Chinese or English</span></p></td></tr><tr><td height="20" bgcolor="#ffffff"><p class="cjk" style="text-align:center;page-break-after:auto;orphans:2;widows:2;page-break-inside:auto;"><span style="color:#000000;font-family:times new roman, serif;font-size:16px;"><span lang="en-US"><span style="font-style:normal;"><span style="text-decoration:none;">4</span></span></span></span></p></td><td bgcolor="#ffffff"><p class="cjk" style="text-align:left;page-break-after:auto;orphans:2;widows:2;page-break-inside:auto;"><span style="font-family:times new roman, serif;"><span lang="en-US"><span style="color:#000000;font-family:times new roman, serif;font-size:16px;"><span lang="en-US"><span style="font-style:normal;"><span style="text-decoration:none;">B</span></span></span></span><span style="color:#000000;font-family:times new roman, serif;font-size:16px;"><span style="font-style:normal;"><span style="text-decoration:none;">iology</span></span></span></span></span></p></td><td bgcolor="#ffffff"><p class="cjk" style="text-align:left;page-break-after:auto;orphans:2;widows:2;page-break-inside:auto;"><span style="color:#000000;font-family:times new roman, serif;font-size:16px;"><span lang="en-US"><span style="font-style:normal;"><span style="text-decoration:none;">School of Bioengineering</span></span></span></span></p></td><td bgcolor="#ffffff"><p class="cjk" style="text-align:center;page-break-after:auto;orphans:2;widows:2;page-break-inside:auto;"><span style="font-family:times new roman, serif;font-size:16px;">Chinese or English</span></p></td></tr><tr><td height="20" bgcolor="#ffffff"><p class="cjk" style="text-align:center;page-break-after:auto;orphans:2;widows:2;page-break-inside:auto;"><span style="color:#000000;font-family:times new roman, serif;font-size:16px;"><span lang="en-US"><span style="font-style:normal;"><span style="text-decoration:none;">5</span></span></span></span></p></td><td bgcolor="#ffffff"><p class="cjk" style="text-align:left;page-break-after:auto;orphans:2;widows:2;page-break-inside:auto;"><span style="font-family:times new roman, serif;"><span lang="en-US"><span style="color:#000000;font-family:times new roman, serif;font-size:16px;"><span style="font-style:normal;"><span style="text-decoration:none;">Mechanical </span></span></span><span style="color:#000000;font-family:times new roman, serif;font-size:16px;"><span lang="en-US"><span style="font-style:normal;"><span style="text-decoration:none;">E</span></span></span></span><span style="color:#000000;font-family:times new roman, serif;font-size:16px;"><span style="font-style:normal;"><span style="text-decoration:none;">ngineering</span></span></span></span></span></p></td><td bgcolor="#ffffff"><p class="cjk" style="text-align:left;page-break-after:auto;orphans:2;widows:2;page-break-inside:auto;"><span style="color:#000000;font-family:times new roman, serif;font-size:16px;"><span lang="en-US"><span style="font-style:normal;"><span style="text-decoration:none;">School of Mechanical and Automotive Engineering</span></span></span></span></p></td><td bgcolor="#ffffff"><p class="cjk" style="text-align:center;page-break-after:auto;orphans:2;widows:2;page-break-inside:auto;"><span style="font-family:times new roman, serif;font-size:16px;">Chinese or English</span></p></td></tr><tr><td height="20" bgcolor="#ffffff"><p class="cjk" style="text-align:center;page-break-after:auto;orphans:2;widows:2;page-break-inside:auto;"><span style="color:#000000;font-family:times new roman, serif;font-size:16px;"><span lang="en-US"><span style="font-style:normal;"><span style="text-decoration:none;">6</span></span></span></span></p></td><td bgcolor="#ffffff"><p class="cjk" style="text-align:left;page-break-after:auto;orphans:2;widows:2;page-break-inside:auto;"><span style="color:#000000;font-family:times new roman, serif;font-size:16px;"><span lang="en-US"><span style="font-style:normal;"><span style="text-decoration:none;">Material Science and Engineering</span></span></span></span></p></td><td bgcolor="#ffffff"><p class="cjk" style="text-align:left;page-break-after:auto;orphans:2;widows:2;page-break-inside:auto;"><span style="color:#000000;font-family:times new roman, serif;font-size:16px;"><span lang="en-US"><span style="font-style:normal;"><span style="text-decoration:none;">School of Material science and Engineering</span></span></span></span></p></td><td bgcolor="#ffffff"><p class="cjk" style="text-align:center;page-break-after:auto;orphans:2;widows:2;page-break-inside:auto;"><span style="font-family:times new roman, serif;font-size:16px;">Chinese or English</span></p></td></tr><tr><td height="20" bgcolor="#ffffff"><p class="cjk" style="text-align:center;page-break-after:auto;orphans:2;widows:2;page-break-inside:auto;"><span style="color:#000000;font-family:times new roman, serif;font-size:16px;"><span lang="en-US"><span style="font-style:normal;"><span style="text-decoration:none;">7</span></span></span></span></p></td><td bgcolor="#ffffff"><p class="cjk" style="text-align:left;page-break-after:auto;orphans:2;widows:2;page-break-inside:auto;"><span style="color:#000000;font-family:times new roman, serif;font-size:16px;"><span style="font-style:normal;"><span style="text-decoration:none;">Control Science and Engineering</span></span></span></p></td><td bgcolor="#ffffff"><p class="cjk" style="text-align:left;page-break-after:auto;orphans:2;widows:2;page-break-inside:auto;"><span style="color:#000000;font-family:times new roman, serif;font-size:16px;"><span lang="en-US"><span style="font-style:normal;"><span style="text-decoration:none;">School of Electrical Engineering and Automation</span></span></span></span></p></td><td bgcolor="#ffffff"><p class="cjk" style="text-align:center;page-break-after:auto;orphans:2;widows:2;page-break-inside:auto;"><span style="font-family:times new roman, serif;font-size:16px;">Chinese or English</span></p></td></tr><tr><td height="20" bgcolor="#ffffff"><p class="cjk" style="text-align:center;page-break-after:auto;orphans:2;widows:2;page-break-inside:auto;"><span style="color:#000000;font-family:times new roman, serif;font-size:16px;"><span lang="en-US"><span style="font-style:normal;"><span style="text-decoration:none;">8</span></span></span></span></p></td><td bgcolor="#ffffff"><p class="cjk" style="text-align:left;page-break-after:auto;orphans:2;widows:2;page-break-inside:auto;"><span style="color:#000000;font-family:times new roman, serif;font-size:16px;"><span lang="en-US"><span style="font-style:normal;"><span style="text-decoration:none;">Computer Science and Technology</span></span></span></span></p></td><td bgcolor="#ffffff"><p class="cjk" style="text-align:left;page-break-after:auto;orphans:2;widows:2;page-break-inside:auto;"><span style="color:#000000;font-family:times new roman, serif;font-size:16px;"><span lang="en-US"><span style="font-style:normal;"><span style="text-decoration:none;">School of Computer Science and Technology</span></span></span></span></p></td><td bgcolor="#ffffff"><p class="cjk" style="text-align:center;page-break-after:auto;orphans:2;widows:2;page-break-inside:auto;"><span style="font-family:times new roman, serif;font-size:16px;">Chinese or English</span></p></td></tr><tr><td height="20" bgcolor="#ffffff"><p class="cjk" style="text-align:center;page-break-after:auto;orphans:2;widows:2;page-break-inside:auto;"><span style="color:#000000;font-family:times new roman, serif;font-size:16px;"><span lang="en-US"><span style="font-style:normal;"><span style="text-decoration:none;">9</span></span></span></span></p></td><td bgcolor="#ffffff"><p class="cjk" style="text-align:left;page-break-after:auto;orphans:2;widows:2;page-break-inside:auto;"><span style="font-family:times new roman, serif;"><span lang="en-US"><span style="color:#000000;font-family:times new roman, serif;font-size:16px;"><span lang="en-US"><span style="font-style:normal;"><span style="text-decoration:none;">C</span></span></span></span><span style="color:#000000;font-family:times new roman, serif;font-size:16px;"><span style="font-style:normal;"><span style="text-decoration:none;">hemical engineering and Technology</span></span></span></span></span></p></td><td bgcolor="#ffffff"><p class="cjk" style="text-align:left;page-break-after:auto;orphans:2;widows:2;page-break-inside:auto;"><span style="color:#000000;font-family:times new roman, serif;font-size:16px;"><span lang="en-US"><span style="font-style:normal;"><span style="text-decoration:none;">School of Chemical and Pharmaceutical Engineering</span></span></span></span></p></td><td bgcolor="#ffffff"><p class="cjk" style="text-align:center;page-break-after:auto;orphans:2;widows:2;page-break-inside:auto;"><span style="font-family:times new roman, serif;font-size:16px;">Chinese or English</span></p></td></tr><tr><td height="20" bgcolor="#ffffff"><p class="cjk" style="text-align:center;page-break-after:auto;orphans:2;widows:2;page-break-inside:auto;"><span style="color:#000000;font-family:times new roman, serif;font-size:16px;"><span lang="en-US"><span style="font-style:normal;"><span style="text-decoration:none;">10</span></span></span></span></p></td><td bgcolor="#ffffff"><p class="cjk" style="text-align:left;page-break-after:auto;orphans:2;widows:2;page-break-inside:auto;"><span style="color:#000000;font-family:times new roman, serif;font-size:16px;"><span lang="en-US"><span style="font-style:normal;"><span style="text-decoration:none;">Light industry Machinery Engineering</span></span></span></span></p></td><td bgcolor="#ffffff"><p class="cjk" style="text-align:left;page-break-after:auto;orphans:2;widows:2;page-break-inside:auto;"><span style="color:#000000;font-family:times new roman, serif;font-size:16px;"><span lang="en-US"><span style="font-style:normal;"><span style="text-decoration:none;">State Key Laboratory</span></span></span></span></p></td><td bgcolor="#ffffff"><p class="cjk" style="text-align:center;page-break-after:auto;orphans:2;widows:2;page-break-inside:auto;"><span style="font-family:times new roman, serif;font-size:16px;">Chinese or English</span></p></td></tr><tr><td height="20" bgcolor="#ffffff"><p class="cjk" style="text-align:center;page-break-after:auto;orphans:2;widows:2;page-break-inside:auto;"><span style="color:#000000;font-family:times new roman, serif;font-size:16px;"><span lang="en-US"><span style="font-style:normal;"><span style="text-decoration:none;">11</span></span></span></span></p></td><td bgcolor="#ffffff"><p class="cjk" style="text-align:left;page-break-after:auto;orphans:2;widows:2;page-break-inside:auto;"><span style="color:#000000;font-family:times new roman, serif;font-size:16px;"><span lang="en-US"><span style="font-style:normal;"><span style="text-decoration:none;">Environmental Science and Engineering</span></span></span></span></p></td><td bgcolor="#ffffff"><p class="cjk" style="text-align:left;page-break-after:auto;orphans:2;widows:2;page-break-inside:auto;"><span style="color:#000000;font-family:times new roman, serif;font-size:16px;"><span lang="en-US"><span style="font-style:normal;"><span style="text-decoration:none;">School of Environmental Science and Engineering</span></span></span></span></p></td><td bgcolor="#ffffff"><p class="cjk" style="text-align:center;page-break-after:auto;orphans:2;widows:2;page-break-inside:auto;"><span style="font-family:times new roman, serif;font-size:16px;">Chinese or English</span></p></td></tr><tr><td height="20" bgcolor="#ffffff"><p class="cjk" style="text-align:center;page-break-after:auto;orphans:2;widows:2;page-break-inside:auto;"><span style="color:#000000;font-family:times new roman, serif;font-size:16px;"><span lang="en-US"><span style="font-style:normal;"><span style="text-decoration:none;">12</span></span></span></span></p></td><td bgcolor="#ffffff"><p class="cjk" style="text-align:left;page-break-after:auto;orphans:2;widows:2;page-break-inside:auto;"><span style="color:#000000;font-family:times new roman, serif;font-size:16px;"><span lang="en-US"><span style="font-style:normal;"><span style="text-decoration:none;">Food Science and Engineering</span></span></span></span></p></td><td bgcolor="#ffffff"><p class="cjk" style="text-align:left;page-break-after:auto;orphans:2;widows:2;page-break-inside:auto;"><span style="color:#000000;font-family:times new roman, serif;font-size:16px;"><span lang="en-US"><span style="font-style:normal;"><span style="text-decoration:none;">School of Food Science and Engineering</span></span></span></span></p></td><td bgcolor="#ffffff"><p class="cjk" style="text-align:center;page-break-after:auto;orphans:2;widows:2;page-break-inside:auto;"><span style="font-family:times new roman, serif;font-size:16px;">Chinese or English</span></p></td></tr><tr><td height="20" bgcolor="#ffffff"><p class="cjk" style="text-align:center;page-break-after:auto;orphans:2;widows:2;page-break-inside:auto;"><span style="color:#000000;font-family:times new roman, serif;font-size:16px;"><span lang="en-US"><span style="font-style:normal;"><span style="text-decoration:none;">13</span></span></span></span></p></td><td bgcolor="#ffffff"><p class="cjk" style="text-align:left;page-break-after:auto;orphans:2;widows:2;page-break-inside:auto;"><span style="color:#000000;font-family:times new roman, serif;font-size:16px;"><span style="font-style:normal;"><span style="text-decoration:none;">Management Science and Engineering</span></span></span></p></td><td bgcolor="#ffffff"><p class="cjk" style="text-align:left;page-break-after:auto;orphans:2;widows:2;page-break-inside:auto;"><span style="color:#000000;font-family:times new roman, serif;font-size:16px;"><span lang="en-US"><span style="font-style:normal;"><span style="text-decoration:none;">School of Management</span></span></span></span></p></td><td bgcolor="#ffffff"><p class="cjk" style="text-align:center;page-break-after:auto;orphans:2;widows:2;page-break-inside:auto;"><span style="font-family:times new roman, serif;font-size:16px;">Chinese or English</span></p></td></tr><tr><td height="20" bgcolor="#ffffff"><p class="cjk" style="text-align:center;page-break-after:auto;orphans:2;widows:2;page-break-inside:auto;"><span style="color:#000000;font-family:times new roman, serif;font-size:16px;"><span lang="en-US"><span style="font-style:normal;"><span style="text-decoration:none;">14</span></span></span></span></p></td><td bgcolor="#ffffff"><p class="cjk" style="text-align:left;page-break-after:auto;orphans:2;widows:2;page-break-inside:auto;"><span style="font-family:times new roman, serif;"><span lang="en-US"><span style="color:#000000;font-family:times new roman, serif;font-size:16px;"><span style="font-style:normal;"><span style="text-decoration:none;">Business </span></span></span><span style="color:#000000;font-family:times new roman, serif;font-size:16px;"><span lang="en-US"><span style="font-style:normal;"><span style="text-decoration:none;">M</span></span></span></span><span style="color:#000000;font-family:times new roman, serif;font-size:16px;"><span style="font-style:normal;"><span style="text-decoration:none;">anagement</span></span></span></span></span></p></td><td bgcolor="#ffffff"><p class="cjk" style="text-align:left;page-break-after:auto;orphans:2;widows:2;page-break-inside:auto;"><span style="color:#000000;font-family:times new roman, serif;font-size:16px;"><span lang="en-US"><span style="font-style:normal;"><span style="text-decoration:none;">School of Management</span></span></span></span></p></td><td bgcolor="#ffffff"><p class="cjk" style="text-align:center;page-break-after:auto;orphans:2;widows:2;page-break-inside:auto;"><span style="font-family:times new roman, serif;font-size:16px;">Chinese or English</span></p></td></tr><tr><td height="19" bgcolor="#ffffff"><p class="cjk" style="text-align:center;page-break-after:auto;orphans:2;widows:2;page-break-inside:auto;"><span style="color:#000000;font-family:times new roman, serif;font-size:16px;"><span lang="en-US"><span style="font-style:normal;"><span style="text-decoration:none;">15</span></span></span></span></p></td><td bgcolor="#ffffff"><p class="cjk" style="text-align:left;page-break-after:auto;orphans:2;widows:2;page-break-inside:auto;"><span style="color:#000000;font-family:times new roman, serif;font-size:16px;"><span style="font-style:normal;"><span style="text-decoration:none;">Design</span></span></span></p></td><td bgcolor="#ffffff"><p class="cjk" style="text-align:left;page-break-after:auto;orphans:2;widows:2;page-break-inside:auto;"><span style="color:#000000;font-family:times new roman, serif;font-size:16px;"><span lang="en-US"><span style="font-style:normal;"><span style="text-decoration:none;">School of Art and Design</span></span></span></span></p></td><td style="-ms-word-break:break-all;" bgcolor="#ffffff"><p class="cjk" style="text-align:center;page-break-after:auto;orphans:2;widows:2;page-break-inside:auto;"><span style="font-family:times new roman, serif;font-size:16px;">Chinese or English</span></p></td></tr></tbody></colgroup></table></center><p class="cjk" style="margin-bottom:0cm;"><br /></p></p></div></article>
          </div>
         
        
      </div>
    </div>
  </div>
</div>
<!--//主体结束--> 


<!--底部开始-->
<div class="wrapper footer" id="footer">
  <div class="inner" frag="面板09">
    <div class="mod clearfix">
      <div class="foot_logo fl"><img src="/_upload/tpl/07/9f/1951/template1951/images/logonew.png"></div>
      <div class="foot-right" frag="窗口09" portletmode="simpleSiteAttri">
        <p class="copyright"><span>版权所有 © 齐鲁工业大学</span><span>鲁ICP备05046217号</span></p>
            <p class="copyright"><span>地址:山东省济南市长清区大学路3501号  邮编250353&nbsp;&nbsp;联系电话:0531-89631016</span></p>
        
      </div>
    </div>
  </div>
</div>

<!--//底部结束-->

<!--读图模式开始-->
<div id="wp-view-page" class="view-image">
  <div id="view-head"><a id="back-read-page" href="javascript:;"><span>返回</span></a><a target="_blank" id="view-original-image"><span>原图</span></a></div>
  <div id="view-body">
    <ul id="view-image-items">
    </ul>
  </div>
  <div id="view-foot">
    <div id="view-index"><span id="view-current"></span>/<span id="view-total"></span></div>
    <p id="view-title"></p>
  </div>
</div>
<!--//读图模式结束-->

</body>
<script type="text/javascript" src="/_upload/tpl/07/9f/1951/template1951/js/app.js"></script>
<script type="text/javascript">
$(function(){
	// 初始化SDAPP
	new SDAPP({
		"menu":{
			type:"aside"
		},
		"view":{
			target:".read img",
			minSize:40
		}
	});
});
</script>
<script>
$(function(){
      $('#wp_nav_w5>ul>li').hover(function(){
        $(this).addClass('active');
      },function(){
        $(this).removeClass('active');
      })
    })
</script>
</html>
 <img src="/_visitcount?siteId=21&type=3&articleId=129377" style="display:none" width="0" height="0"></image>

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
	type args struct {
		body string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test", args: args{body: "\t<!DOCTYPE html PUBLIC \"-//W3C//DTD XHTML 1.0 Transitional//EN\" \"http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd\">\n<html xmlns=\"http://www.w3.org/1999/xhtml\">\n<head>\n<meta http-equiv=\"Content-Type\" content=\"text/html; charset=gbk\" />\n\n<meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=no\">\n<title>吾爱破解 - LCG - LSG|安卓破解|病毒分析|www.52pojie.cn </title>\n            <style type=\"text/css\">\n\t\t\t    html{-webkit-filter:grayscale(100%);filter:grayscale(100%);}\n\t\t\t    body{background-blend-mode:luminosity;}\n\t\t\t    .hdc h2:after{content: \"\";font-size: 17px;}\n\t\t\t    @media (max-width: 650px){\n\t\t\t      .hdc h2:after{font-size: 0.5rem;margin-left: 0.3rem;display: block;}\n\t\t\t    }\n\t\t    </style>\n<meta name=\"keywords\" content=\"吾爱破解,破解论坛,破解软件,软件论坛,病毒分析,脱壳破解,安卓破解,加密解密,软件安全\" />\n<meta name=\"description\" content=\"吾爱破解论坛致力于软件安全与病毒分析的前沿，丰富的技术版块交相辉映，由无数热衷于软件加密解密及反病毒爱好者共同维护 \" />\n<meta name=\"generator\" content=\"Discuz!\" />\n<meta name=\"MSSmartTagsPreventParsing\" content=\"True\" />\n<meta http-equiv=\"MSThemeCompatible\" content=\"Yes\" />\n<base href=\"https://www.52pojie.cn/\" /><link rel=\"stylesheet\" type=\"text/css\" href=\"https://static.52pojie.cn/cache/style_1_common.css?hf4\" /><link rel=\"stylesheet\" type=\"text/css\" href=\"https://static.52pojie.cn/cache/style_1_forum_index.css?hf4\" /><link rel=\"stylesheet\" id=\"css_extstyle\" type=\"text/css\" href=\"./template/default/style/t1/style.css\" />\n\n<meta name=\"application-name\" content=\"吾爱破解 - LCG - LSG |安卓破解|病毒分析|www.52pojie.cn\" />\n<link rel=\"archives\" title=\"吾爱破解 - LCG - LSG |安卓破解|病毒分析|www.52pojie.cn\" href=\"https://www.52pojie.cn/archiver/\" />\n<link rel=\"stylesheet\" id=\"css_widthauto\" type=\"text/css\" href='https://static.52pojie.cn/cache/style_1_widthauto.css?hf4' />\n\n\n<link rel=\"stylesheet\" id=\"css_responsive\" type=\"text/css\" href='https://static.52pojie.cn/cache/style_1_responsive.css?hf4' />\n\n</head>\n\n<body id=\"nv_forum\" class=\"pg_index\" onkeydown=\"if(event.keyCode==27) return false;\">"},
			want: "吾爱破解 - LCG - LSG|安卓破解|病毒分析|www.52pojie.cn "},
		{name: "2", args: args{body: "<titlE>吾爱破解 - LCG - LSG|安卓破解|病毒分析|www.52pojie.cn </tItle>"}, want: "吾爱破解 - LCG - LSG|安卓破解|病毒分析|www.52pojie.cn "},
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
				Murl: "https://international.qlu.edu.cn/2019/0530/c6614a129377/page.htm",
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
