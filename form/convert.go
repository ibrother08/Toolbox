package form

import (
	"strings"

	"github.com/ChimeraCoder/gojson"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

type ConvertBox struct {
	PnlTop     *vcl.TPanel
	MeoLeft    *vcl.TMemo
	MeoRight   *vcl.TMemo
	BtnConvert *vcl.TButton
}

func (f *TBMainForm) CreateConvertBox() {
	f.ConvertBox = &ConvertBox{}
	owner := f
	parent := f.Sheet1
	box := f.ConvertBox

	box.PnlTop = vcl.NewPanel(owner)
	box.PnlTop.SetParent(parent)
	box.PnlTop.SetAlign(types.AlTop)
	box.PnlTop.SetHeight(25)

	box.MeoLeft = vcl.NewMemo(owner)
	box.MeoLeft.SetParent(parent)
	box.MeoLeft.SetAlign(types.AlLeft)
	box.MeoLeft.SetWidth(390)
	box.MeoLeft.SetScrollBars(types.SsAutoVertical)
	fontLeft := box.MeoLeft.Font()
	fontLeft.SetSize(10)

	box.MeoRight = vcl.NewMemo(owner)
	box.MeoRight.SetParent(parent)
	box.MeoRight.SetAlign(types.AlRight)
	box.MeoRight.SetWidth(390)
	box.MeoRight.SetScrollBars(types.SsAutoVertical)
	box.MeoRight.SetReadOnly(true)
	fontRight := box.MeoRight.Font()
	fontRight.SetSize(10)

	box.BtnConvert = vcl.NewButton(owner)
	box.BtnConvert.SetParent(box.PnlTop)
	box.BtnConvert.SetAlign(types.AlLeft)
	box.BtnConvert.SetCaption("转换")
	box.BtnConvert.SetOnClick(box.onBtnConvertClick)
}

func (b *ConvertBox) onBtnConvertClick(sender vcl.IObject) {
	text := b.MeoLeft.Text()
	if text == "" {
		return
	}

	i := strings.NewReader(text)
	byt, err := gojson.Generate(i, gojson.ParseJson, "StructName", "PackageName", []string{"json"}, false, true)
	if err != nil {
		vcl.ShowMessage("错误：" + err.Error())
		return
	}

	b.MeoRight.SetText(string(byt))
}
