package Crawler

import (
	"GScan/infoscan/dao"
	"GScan/pkg/logger"
	"encoding/json"
	"fmt"
	"github.com/axgle/mahonia"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
	"html"
	"io/ioutil"
	"log"
	"strings"
)

func OutPutRes(jobid uint, DAO dao.IDAO) string {
	bufmap := map[string]*strings.Builder{}
	result := DAO.GetResult(jobid)
	for _, r := range result {
		if _, ok := bufmap[r.Type]; !ok {
			bufmap[r.Type] = &strings.Builder{}
			bufmap[r.Type].WriteString(fmt.Sprintf("<table><caption>%s</caption>", r.Type))
			bufmap[r.Type].WriteString(fmt.Sprintf("<tr><th >%s</th><th>%s</th><th>%s</th></tr>", "URL", "父URL", "数据"))

		}
		url1 := DAO.GetOnePages(&dao.Page{
			Model: gorm.Model{ID: r.PageID},
			JobID: jobid,
		})
		get, err := DAO.WebTreeGet(r.JobID, r.PageID)
		url2 := &dao.Page{}
		if err == nil {
			url2 = DAO.GetOnePages(&dao.Page{
				Model: gorm.Model{ID: get[0]},
				JobID: jobid,
			})
		} else {
			url2.URL = "查询失败"
		}
		bufmap[r.Type].WriteString(fmt.Sprintf("<tr><td>%s</td><td>%s</td><td>%s</td></tr>", url1.URL, url2.URL, html.EscapeString(r.Data)))
		//bufmap[r.Type].WriteString(fmt.Sprintf("[%s](%s)发现外链%s\n", url1.URL, url2.URL, r.Data))
	}
	var all strings.Builder
	all.WriteString(`<html>
<head>
    <style>
        table {
            width: 90%;
            background: #ccc;
            margin: 10px auto;
            border-collapse: collapse;
            /*border-collapse:collapse合并内外边距
(去除表格单元格默认的2个像素内外边距*/
        }

        th,
        td {
            height: 25px;
            line-height: 25px;
            text-align: center;
            border: 1px solid #ccc;
        }

        th {
            background: #eee;
            font-weight: normal;
        }

        tr {
            background: #fff;
        }

        tr:hover {
            background: rgb(80, 220, 224);
        }

        td a {
            color: #06f;
            text-decoration: none;
        }

        td a:hover {
            color: #06f;
            text-decoration: underline;
        }
    </style>
</head>

<body>
`)
	for _, builder := range bufmap {
		all.WriteString(builder.String())
		all.WriteString("</table>")
		all.WriteString("<br>")
	}
	all.WriteString(`</body>
</html>`)
	encoder := mahonia.NewEncoder("UTF-8")
	encoder.ConvertString(all.String())
	return encoder.ConvertString(all.String())
}

func Out2Excel(jobid uint, DAO dao.IDAO, filename string) {
	result := DAO.GetResult(jobid)
	if len(result) == 0 {
		log.Fatalln("结果为空！")
	}
	logger.PF(logger.LINFO, "<Out2Excel>正在输出结果")

	// 数据写入内存
	sheetData := make(map[string][][]string)
	pagesData := DAO.GetAllPages(&dao.Page{
		JobID: jobid,
	})
	webTreeData, _ := DAO.WebTreeGetAll(jobid)
	webTrees := make(map[uint]*dao.WebTree)
	pages := make(map[uint]*dao.Page)

	for _, page := range pagesData {
		pages[page.ID] = page
	}
	for _, webtree := range webTreeData {
		webTrees[webtree.PageID] = webtree
	}
	for _, r := range result {
		if _, ok := sheetData[r.Type]; !ok {
			sheetData[r.Type] = [][]string{{
				"URL",
				"父URL",
				"数据",
			}}
		}
		// 获取 URL 和父URL
		url1, ok := pages[r.PageID]
		if !ok {
			logger.PF(logger.LERROR, "<Out2Excel>[JobID:%d]结果列表中有不存在的PageID:%d", jobid, r.PageID) //目前造成这个bug的原因未知
			continue
		}
		url2 := ""
		var get []uint
		if webTrees[r.PageID] == nil {
			url2 = "查询失败"
		} else {
			for _, v := range webTrees[r.PageID].FiD {
				get = append(get, v)
			}
			if pages[get[0]] != nil {
				url2 = pages[get[0]].URL
			} else {
				url2 = "查询失败"
			}
		}
		sheetData[r.Type] = append(sheetData[r.Type], []string{
			url1.URL,
			url2,
			r.Data,
		})
	}

	f := excelize.NewFile()
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// 将数据写入文件
	for sheetName, data := range sheetData {
		if f.GetSheetIndex(sheetName) == -1 {
			f.NewSheet(sheetName)
		}
		err := f.SetSheetRow(sheetName, "A1", &data[0])
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		err = f.SetSheetRow(sheetName, "A2", &data[1])
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		for i := 2; i < len(data); i++ {
			axis := fmt.Sprintf("A%d", i+1)
			err = f.SetSheetRow(sheetName, axis, &data[i])
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
		}
	}

	f.DeleteSheet("Sheet1")
	if err := f.SaveAs(filename); err != nil {
		fmt.Println(err.Error())
		return
	}
	logger.PF(logger.LINFO, "<Out2Excel>输出结果完成，%s", filename)
}

func Out2Json(jobid uint, DAO dao.IDAO, filename string) {
	result := DAO.GetResult(jobid)
	if len(result) == 0 {
		log.Fatalln("结果为空！")
	}
	logger.PF(logger.LINFO, "<Out2Json>正在输出结果")
	marshal, err := json.Marshal(&result)
	if err != nil {
		log.Fatalln(err)
	}
	ioutil.WriteFile(filename, marshal, 0644)
}
