package main

import (
	"fmt"
	"log"

	"github.com/gotk3/gotk3/gtk"
)

func createHeader() *gtk.HeaderBar {
	header, err := gtk.HeaderBarNew()
	if err != nil {
		log.Fatal("Could not create header bar:", err)
	}
	header.SetShowCloseButton(true)
	header.SetTitle("GO GTK")
	header.SetSubtitle("The Sandbox")
	return header
}

func createAdder() *gtk.Box {

	counter := 0

	box, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 2)

	label, err := gtk.LabelNew(fmt.Sprintf("%d", counter))
	if err != nil {
		log.Fatal("Unable to create label:", err)
	}

	button, err := gtk.ButtonNewWithLabel("Add 1")
	if err != nil {
		log.Fatal("Could not create button:", err)
	}
	button.Connect("clicked", func() {
		counter++
		labelText := fmt.Sprintf("%d", counter)
		label.SetText(labelText)
	})

	box.Add(label)
	box.Add(button)

	return box
}

func main() {
	gtk.Init(nil)

	// Create a toplevel window, set its title and
	// connect it to the destroy signal to exit the GTK
	// main loop when its is destroyed.
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}

	win.SetTitle("Sandbox")
	win.SetTitlebar(createHeader())

	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	win.Add(createAdder())

	win.SetDefaultSize(300, 100)
	win.ShowAll()

	gtk.Main()

}
