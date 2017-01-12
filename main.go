package main

import (
    "crypto/md5"
    "encoding/hex"

    "github.com/andlabs/ui"
)

func GetMD5Hash(text string) string {
    hasher := md5.New()
    hasher.Write([]byte(text))
    return hex.EncodeToString(hasher.Sum(nil))
}

func main() {
    err := ui.Main(func() {
        text := ui.NewEntry()
        hashed := ui.NewEntry()
        button := ui.NewButton("Hash")
        upperLabel := ui.NewLabel("Enter text:")
        lowerLabel := ui.NewLabel("Hash:")
        box := ui.NewVerticalBox()
        box.Append(upperLabel, false)
        box.Append(text, false)
        box.Append(lowerLabel, false)
        box.Append(hashed, false)
        box.Append(button, false)
        window := ui.NewWindow("MD5", 300, 100, false)
        window.SetChild(box)
        button.OnClicked(func(*ui.Button) {
            hashed.SetText(GetMD5Hash(text.Text()))
        })
        window.OnClosing(func(*ui.Window) bool {
            ui.Quit()
            return true
        })
        window.Show()
    })
    if err != nil {
        panic(err)
    }
}
