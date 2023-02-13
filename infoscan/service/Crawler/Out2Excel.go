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
	f := excelize.NewFile()
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	for _, r := range result {
		a := []string{}
		if f.GetSheetIndex(r.Type) == -1 {
			f.NewSheet(r.Type)
			//raw := map[string]interface{}{}
			//err := json.Unmarshal([]byte(r.Data), &raw)
			//if err != nil {
			//	fmt.Println(err.Error())
			//	continue
			//}
			a = append(a, "URL")
			a = append(a, "父URL")
			//for k, _ := range raw {
			//	a = append(a, k)
			//}
			a = append(a, "数据")
			err := f.SetSheetRow(r.Type, "A1", &a)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
		} else {
			continue
		}
	}
	id := map[string]int{}
	for k, r := range result {
		lensssss := len(result)
		fmt.Printf("<Out2Excel>正在输出(%d/%d)\r", k, lensssss)
		//raw := map[string]interface{}{}
		if _, ok := id[r.Type]; !ok {
			id[r.Type] = 2
		} else {
			id[r.Type]++
		}
		axis := fmt.Sprintf("A%d", id[r.Type])
		a := []string{}
		//err := json.Unmarshal([]byte(r.Data), &raw)
		//if err != nil {
		//	fmt.Println(err.Error())
		//	continue
		//}
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
		a = append(a, url1.URL)
		a = append(a, url2.URL)
		a = append(a, r.Data)
		err = f.SetSheetRow(r.Type, axis, &a)
		if err != nil {
			fmt.Println(err.Error())
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
