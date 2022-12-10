package Processor

import (
	"net/url"
	"reflect"
	"testing"
)

var a string = `

<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN">
<html>
	<head>
				
		<title>站内搜索 - 齐鲁工业大学校友会</title>
		<meta http-equiv="keywords" content="站内搜索,齐鲁工业大学校友会,齐鲁工业大学,校友会" />
		<meta http-equiv="description" content="" />
		
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
		
	  	<script src="../source/jquery-1.6.2.min.js" type="text/javascript"></script>
		<link href="../source/common.css" rel="stylesheet" type="text/css" />
		<link href="../source/item.css" rel="stylesheet" type="text/css" />
		<script src="../source/common.js" type="text/javascript"></script>
		<!--[if lte IE 6]>
		<script src="../source/DD_belatedPNG_0.0.8a.js" type="text/javascript"></script>
	    <script type="text/javascript">
	        DD_belatedPNG.fix('div, td, img, a');
	    </script>
		<![endif]-->
  	</head>
  	
  	<body>
  		<div class="page">
			
<script type="text/javascript">
	$(document).ready(function(){
		$("input[name='search']").blur(function(){
			if($.trim($(this).val())==""){
				$(this).addClass("alpha");
			}
		}).focus(function(){
			$(this).removeClass("alpha");
		});
		$("#searchBtn").click(function(){
			window.location.href="../search?key="+$("#searchText").val();
		});
	});
</script>
	<!--header-->
	<div class="header">
		<div class="header_logo">
			<div class="banner" style="width:1024px;">
				<div class="logo" style="float:left;"></div>
				<div class="search" style="background:url(../source/images/search_back.png) no-repeat center;float:right;margin-top:30px;height:30px;width:180px;">
					<input type="text" id="searchText" placeholder="输入关键词" style="height:24px;
						line-height:24px;font-size:14px;margin:2px 5px;width:145px;float:left;margin-right:0px;
						background:#dddddd;border: none;outline: none;">
					<a href="#" id="searchBtn" style="height:24px;width:24px;margin:3px;margin-left:0px;float:right;display:inline-block;"></a>
				</div>
			</div>
		</div>
		<div class="header_nav">
			<div class="header_menu">
				<ul id="nav">
					<li><a class="fstpg" href="../index">首页</a></li>
					
						<li >
						
							
							
									<a class="lv1"  href="../item?id=2046" target="_blank">工作动态
									</a>
							
						
						<!-- 子菜单 -->
						
						
							<ul class="children">
								
									<li title="新闻动态" >
										
											
											
												<a href="../item?id=2052" target="_blank">&nbsp;新闻动态</a>
											
										
									</li>
								
									<li title="活动公告" >
										
											
											
												<a href="../item?id=2054" target="_blank">&nbsp;活动公告</a>
											
										
									</li>
								
							</ul>
						
						</li>
					
						<li >
						
							
							
									<a class="lv1"  href="../item?id=2048" target="_blank">校友组织
									</a>
							
						
						<!-- 子菜单 -->
						
						
							<ul class="children">
								
									<li title="学院校友会" >
										
											
											
												<a href="../item?id=2066" target="_blank">&nbsp;学院校友会</a>
											
										
									</li>
								
									<li title="国内校友会" >
										
											
											
												<a href="../item?id=2064" target="_blank">&nbsp;国内校友会</a>
											
										
									</li>
								
									<li title="海外校友会" >
										
											
											
												<a href="../item?id=2065" target="_blank">&nbsp;海外校友会</a>
											
										
									</li>
								
									<li title="行业校友会" >
										
											
											
												<a href="../item?id=2067" target="_blank">&nbsp;行业校友会</a>
											
										
									</li>
								
							</ul>
						
						</li>
					
						<li >
						
							
							
									<a class="lv1"  href="../item?id=2050" target="_blank">校友服务
									</a>
							
						
						<!-- 子菜单 -->
						
						
							<ul class="children">
								
									<li title="返校服务" >
										
											
											
												<a href="../item?id=2121" target="_blank">&nbsp;返校服务</a>
											
										
									</li>
								
									<li title="校友通讯" >
										
											
											
												<a href="../item?id=2063" target="_blank">&nbsp;校友通讯</a>
											
										
									</li>
								
									<li title="校友信使" >
										
											
											
												<a href="../item?id=2161" target="_blank">&nbsp;校友信使</a>
											
										
									</li>
								
							</ul>
						
						</li>
					
						<li >
						
							
									<a class="lv1"  href="../../../sdili_jjh/jjh/item/?id=2038"
										target="_blank">社会捐赠
									</a>
							
							
						
						<!-- 子菜单 -->
						
						
						</li>
					
						<li >
						
							
							
									<a class="lv1"  href="../item?id=2076" target="_blank">校友风采
									</a>
							
						
						<!-- 子菜单 -->
						
						
							<ul class="children">
								
									<li title="政界明星" >
										
											
											
												<a href="../item?id=2101" target="_blank">&nbsp;政界明星</a>
											
										
									</li>
								
									<li title="企业精英" >
										
											
											
												<a href="../item?id=2102" target="_blank">&nbsp;企业精英</a>
											
										
									</li>
								
									<li title="专家学者" >
										
											
											
												<a href="../item?id=2103" target="_blank">&nbsp;专家学者</a>
											
										
									</li>
								
									<li title="领域达人" >
										
											
											
												<a href="../item?id=2104" target="_blank">&nbsp;领域达人</a>
											
										
									</li>
								
									<li title="科创达人" >
										
											
											
												<a href="../item?id=2105" target="_blank">&nbsp;科创达人</a>
											
										
									</li>
								
									<li title="上市公司" >
										
											
											
												<a href="../item?id=2106" target="_blank">&nbsp;上市公司</a>
											
										
									</li>
								
							</ul>
						
						</li>
					
						<li >
						
							
									<a class="lv1"  href="http://hzfzc.qlu.edu.cn/"
										target="_blank">合作发展
									</a>
							
							
						
						<!-- 子菜单 -->
						
						
						</li>
					
						<li >
						
							
							
									<a class="lv1"  href="../item?id=2079" target="_blank">校史馆
									</a>
							
						
						<!-- 子菜单 -->
						
						
							<ul class="children">
								
									<li title="校友活动集" >
										
											
											
												<a href="../item?id=2110" target="_blank">&nbsp;校友活动集</a>
											
										
									</li>
								
									<li title="历史老照片" >
										
											
											
												<a href="../item?id=2111" target="_blank">&nbsp;历史老照片</a>
											
										
									</li>
								
									<li title="老校区照片" >
										
											
											
												<a href="../item?id=2112" target="_blank">&nbsp;老校区照片</a>
											
										
									</li>
								
							</ul>
						
						</li>
					
						<li >
						
							
							
									<a class="lv1" style="border:0;" href="../item?id=2078" target="_blank">联系我们
									</a>
							
						
						<!-- 子菜单 -->
						
						
						</li>
					
				</ul>
			</div>
		</div>
	</div>
	<!--header end-->
			<div class="allPage page1">
				<div class="maper">
					<div class="map">
						您所在的位置：
						<a href="../index">齐鲁工业大学校友会</a>
						>
						<a href="javascript:;">站内搜索</a>
					</div>
					<div class="local">
						<h1>站内搜索</h1>
					</div>
				</div>
				<div class="container">
					<div class="containerMain">
						<div class="containerMenu">
							<div class="containerNav">
													
					
					
						<div class="naver">
							<div class="title">
								<div class="titleCnt">
									栏目导航
								</div>
							</div>
							<div class="main">
								<div class="mainTop"></div>
								<div class="mainCnt">
									<ul>
									
										
										<li class="option">
											<a href="../item?id=2046" >工作动态</a>
										</li>
										
									
										
										<li class="option">
											<a href="../item?id=2048" >校友组织</a>
										</li>
										
									
										
										<li class="option">
											<a href="../item?id=2050" >校友服务</a>
										</li>
										
									
										
										<li class="option">
											<a href="../../../sdili_jjh/jjh/item/?id=2038" target="_blank">社会捐赠</a>
										</li>
										
									
										
										<li class="option">
											<a href="../item?id=2076" >校友风采</a>
										</li>
										
									
										
										<li class="option">
											<a href="http://hzfzc.qlu.edu.cn/" target="_self">合作发展</a>
										</li>
										
									
										
										<li class="option">
											<a href="../item?id=2079" >校史馆</a>
										</li>
										
									
										
										<li class="option">
											<a href="../item?id=2078" >联系我们</a>
										</li>
										
									
									</ul>
								</div>
							</div>
							<div class="bottom"></div>
						</div>
					
					
					
					
							</div>
						</div>
						<div class="containerCnt">
							<div class="contenterTop"></div>
							<div class="contenter">
							<div class="titler">
								<div class="titleCnt">
									站内搜索→[<font color="red"></font>]
								</div>
							</div>
							
								
									<div class="empty">暂无满足条件的相关信息！</div>
								
								
							
							
						</div>
						<div class="contenterBottom"></div>
					</div>
				</div>
			</div>
  		</div>
		


	<!--footer-->
	<div style="margin:20px 0 10px 0;border-bottom:1px solid #ccc;"></div>
	<div class="footer" style="height:160px;background:#104179;">
		<div style="width:1082px;margin:0 auto;height:100%;">
			<div style="width:360px;height:120px;margin: 20px auto;line-height: 30px;text-align: center;color:#fff;float:left;border-right:1px solid #fff;">
				<p>校园网 V1.0</p>
				<p>版权所有© 齐鲁工业大学校友会 </p>
				<p>技术支持：蓝矩科技</p>
			</div>
			<div style="width:360px;height:120px;margin: 20px auto;line-height: 30px;text-align: center;color:#fff;float:left;border-right:1px solid #fff;">
				<p style="font-size:16px;">联系我们</p>
				<p>0531-89631717</p>
				<p>xyh@qlu.edu.cn</p>
				<p>济南市长清区大学路3501号齐鲁工业大学行政楼116室</p>
			</div>
			<div style="width:360px;height:120px;margin: 20px auto;line-height: 30px;text-align: center;color:#fff;float:left;">
				<p style="font-size:16px;">友情链接</p>
				
				<p><a style="color:#fff;" href="
								
									http://www.qlu.edu.cn/
								
								
							" target="_blank">齐鲁工业大学</a></p>
				
				<p><a style="color:#fff;" href="
								
									http://xypt.qlu.edu.cn/alumniLogin
								
								
							" target="_blank">校友社区</a></p>
				
				<p><a style="color:#fff;" href="
								
									http://jjh.qlu.edu.cn/site/sdili_jjh/jjh/index/
								
								
							" target="_blank">齐鲁工业大学基金会</a></p>
				
			</div>
		</div>
	</div>
	<!--footer end-->
  		</div>
  	</body>
</html>

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
		{name: "test", args: args{body: "  color:#C55F5F;\n}\n.alert-info{\n  background-color:#F0F3F5;\n}\n\n\t</style>\n\t<title>齐鲁工业大学录取咨询平台</title>\n</head>\n<body>\n\n<div id='header'>\n\t<!--\n\t<div class='row navbar  navbar-inverse'>\n\t\t<div id='header' class='navbar-inner'>\n\t\t\t<div class='container-fluid'>\n\t\t\t\t<ul class='nav pull-right'>\n\t\t\t\t\t<li>"},
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
				Murl: "http://xyh.qlu.edu.cn/site/sdili_xyh/xyh/search?key=",
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
