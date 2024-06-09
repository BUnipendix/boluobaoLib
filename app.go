package boluobaoLib

import (
	"fmt"
	"github.com/AlexiaVeronica/boluobaoLib/boluobaomodel"
	"github.com/AlexiaVeronica/input"
	"sync"
)

type APP struct {
	threadNum int
	client    *Client
	bookInfo  *boluobaomodel.BookInfoData
}

func (client *Client) APP() *APP {
	return &APP{client: client, threadNum: 32}
}

func (app *APP) MergeText(f func(chapter boluobaomodel.ChapterList)) {
	if app.bookInfo == nil {
		fmt.Println("Please set book info first!")
		return
	}
	app.eachChapter(f)
}

func (app *APP) SetThreadNum(threadNum int) *APP {
	app.threadNum = threadNum
	return app
}

func (app *APP) SetBookInfo(bookInfo *boluobaomodel.BookInfoData) *APP {
	app.bookInfo = bookInfo
	return app
}

func (app *APP) eachChapter(f func(boluobaomodel.ChapterList)) {
	divisionList, err := app.client.API().GetCatalogue(app.bookInfo.NovelId)
	if err != nil {
		fmt.Println("get division list error:", err)
		return
	}
	for _, division := range divisionList.Data.VolumeList {
		for _, chapter := range division.ChapterList {
			f(chapter)
		}
	}
}

func (app *APP) Download(f1 continueFunction, f2 contentFunction) {
	if app.bookInfo == nil {
		fmt.Println("Please set book info first!")
		return
	}

	var wg sync.WaitGroup
	ch := make(chan struct{}, app.threadNum)

	app.eachChapter(func(chapter boluobaomodel.ChapterList) {
		wg.Add(1)
		ch <- struct{}{}

		go func(chapter boluobaomodel.ChapterList) {
			defer func() {
				wg.Done()
				<-ch
			}()
			if f1(app.bookInfo, chapter) {
				content, err := app.client.API().GetChapterContent(chapter.ChapID)
				if err != nil {
					fmt.Println("get chapter content error:", err)
					return
				}
				f2(app.bookInfo, &content.Data)
			}
		}(chapter)
	})

	wg.Wait()
}

func (app *APP) Search(keyword string, f1 continueFunction, f2 contentFunction) {
	searchInfo, err := app.client.API().GetSearch(keyword, 0)
	if err != nil {
		fmt.Println("search failed! " + err.Error())
		return
	}
	searchInfo.EachBook(func(index int, book boluobaomodel.BookInfoData) {
		fmt.Println("Index:", index, "\tBookName:", book.NovelName)
	})
	app.bookInfo = searchInfo.GetBook(input.IntInput("Please input the index of the book you want to download"))
	app.Download(f1, f2)
}

func (app *APP) Bookshelf(f1 continueFunction, f2 contentFunction) {
	shelf, err := app.client.API().GetBookShelfInfo()
	if err != nil {
		fmt.Println("get bookshelf error:", err)
		return
	}
	shelf.EachShelf(func(index int, shelf boluobaomodel.ShelfData) {
		fmt.Println("Index:", index, "\tShelfName:", shelf.Name, "\tShelfNum:", len(shelf.Expand.Novels))
	})
	bookshelf := shelf.GetShelf(input.IntInput("Please input the index of the bookshelf you want to download"))
	bookshelf.EachBookshelf(func(index int, book boluobaomodel.BookInfoData) {
		fmt.Println("Index:", index, "\tBookName:", book.NovelName)
	})
	app.bookInfo = bookshelf.GetBookshelf(input.IntInput("Please input the index of the book you want to download"))
	app.Download(f1, f2)
}
