package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/rumblefrog/go-a2s"
)

const appId = "org.bulkmoerls.isshounicfull"

var (
	statusBar    *gtk.Label
	refreshBtn   *gtk.Button
	playersLabel *gtk.Label
	mapLabel     *gtk.Label
	keywordLabel *gtk.Label
)

/**
 * The code below is where we query and update the UI.
 * Seems like everything is going smoothly. Lovely!
 */

func queryUpdate(serverAddr string) {
	client, err := a2s.NewClient(serverAddr)
	if err != nil {
		log.Fatalln("Cannot add a client.")
	}
	defer client.Close()

	info, err := client.QueryInfo()
	if err != nil {
		statusBar.SetText("Can't connect to server.")
		return
	}

	if info.Players == info.MaxPlayers {
		statusBar.SetText("Server is full. Better luck next time, tough guy!")
	} else {
		statusBar.SetText("Server isn't full. Lucky you!")
	}
	/**
	 * had to add another variable because
	 * GoTK3 doesn't support formatting text
	 * (example "%s")
	 */
	mapText := fmt.Sprintf("Map: %s", info.Map)
	playersText := fmt.Sprintf("Players: %d/%d", info.Players, info.MaxPlayers)
	keywordText := fmt.Sprintf("Keywords:\n %s", info.ExtendedServerInfo.Keywords)

	mapLabel.SetText(mapText)
	playersLabel.SetText(playersText)
	keywordLabel.SetText(keywordText)
}
func main() {
	fmt.Println("Initializing GTK...")
	// Create a new application.
	application, err := gtk.ApplicationNew(appId, glib.APPLICATION_FLAGS_NONE)
	errorCheck(err)

	// Connect function to application startup event, this is not required.
	application.Connect("startup", func() {
		log.Println("application startup")
	})

	// Connect function to application activate event
	application.Connect("activate", func() {
		log.Println("application activate")

		// Get the GtkBuilder UI definition in the glade file.
		builder, err := gtk.BuilderNewFromFile("ui/isshounicfull.ui")
		if err != nil {
			log.Fatal("The UI file, critical for loading the whole ui, isn't present while Building.", err)
		}

		// Map the handlers to callback functions, and connect the signals
		// to the Builder.
		signals := map[string]interface{}{
			"on_main_window_destroy": onMainWindowDestroy,
		}
		builder.ConnectSignals(signals)
		/**
		 * GTK OBJ BLOCK STARTS HERE
		 *
		 * This block is where we define the UI so that it can do things.
		 */

		statusObj, err := builder.GetObject("status-label")
		if err != nil {
			log.Fatal("Can't find status-label in UI file.")
		}
		statusBar = statusObj.(*gtk.Label)

		// Get the object with the id of "main_window".
		obj, err := builder.GetObject("main_window")
		errorCheck(err)

		btnObj, err := builder.GetObject("refresh-btn")
		if err != nil {
			log.Fatal("Could not find Object: refresh-btn.")
		}
		refreshBtn := btnObj.(*gtk.Button)

		mapLabelObj, err := builder.GetObject("map-label")
		if err != nil {
			log.Fatalln("failed to find map-label.")
		}
		mapLabel = mapLabelObj.(*gtk.Label)

		playersLabelObj, err := builder.GetObject("players-label")
		if err != nil {
			log.Fatalln("failed to find players-label.")
		}
		playersLabel = playersLabelObj.(*gtk.Label)

		keywordLabelObj, err := builder.GetObject("keyword-label")
		if err != nil {
			log.Fatalln("failed to find keyword-label.")
		}
		keywordLabel = keywordLabelObj.(*gtk.Label)
		/**
		 * GTK OBJ BLOCK STOPS HERE
		 */
		refreshBtn.Connect("clicked", func(btn *gtk.Button) {
			queryUpdate("45.62.160.71:27015")
		})
		// Verify that the object is a pointer to a gtk.ApplicationWindow.
		win, err := isWindow(obj)
		errorCheck(err)

		// Show the Window and all of its components.
		win.ShowAll()
		application.AddWindow(win)
		queryUpdate("45.62.160.71:27015")
	})
	// Connect function to application shutdown event, this is not required.
	application.Connect("shutdown", func() {
		log.Println("application shutdown")
	})
	// Launch the application
	os.Exit(application.Run(os.Args))
}

func isWindow(obj glib.IObject) (*gtk.Window, error) {
	// Make type assertion (as per gtk.go).
	if win, ok := obj.(*gtk.Window); ok {
		return win, nil
	}
	return nil, errors.New("not a *gtk.Window")
}

func errorCheck(e error) {
	if e != nil {
		// panic for any errors.
		log.Panic(e)
	}
}

// onMainWindowDestory is the callback that is linked to the
// on_main_window_destroy handler. It is not required to map this,
// and is here to simply demo how to hook-up custom callbacks.
func onMainWindowDestroy() {
	log.Println("onMainWindowDestroy")
}
