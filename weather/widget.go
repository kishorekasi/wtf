package weather

import (
	"fmt"
	"time"

	owm "github.com/briandowns/openweathermap"
	"github.com/rivo/tview"
)

type Widget struct {
	RefreshedAt time.Time
	View        *tview.TextView
}

func NewWidget() *Widget {
	widget := Widget{
		RefreshedAt: time.Now(),
	}

	widget.addView()

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	data := Fetch()

	widget.View.SetTitle(fmt.Sprintf(" %s Weather - %s ", icon(data), data.Name))
	widget.RefreshedAt = time.Now()

	fmt.Fprintf(widget.View, "%s", widget.contentFrom(data))
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) addView() {
	view := tview.NewTextView()

	view.SetBorder(true)
	view.SetDynamicColors(true)
	view.SetTitle(" Weather ")

	widget.View = view
}

func centerText(str string, width int) string {
	return fmt.Sprintf("%[1]*s", -width, fmt.Sprintf("%[1]*s", (width+len(str))/2, str))
}

func (widget *Widget) contentFrom(data *owm.CurrentWeatherData) string {
	str := fmt.Sprintf("\n")

	for _, weather := range data.Weather {
		str = str + fmt.Sprintf(" %s\n\n", weather.Description)
	}

	str = str + fmt.Sprintf("%10s: %4.1f° C\n\n", "Current", data.Main.Temp)
	str = str + fmt.Sprintf("%10s: %4.1f° C\n", "High", data.Main.TempMax)
	str = str + fmt.Sprintf("%10s: %4.1f° C\n", "Low", data.Main.TempMin)
	str = str + "\n\n\n\n"
	str = str + centerText(fmt.Sprintf("Refreshed at %s", widget.refreshedAt()), 38)

	return str
}

// icon returns an emoji for the current weather
// src: https://github.com/chubin/wttr.in/blob/master/share/translations/en.txt
func icon(data *owm.CurrentWeatherData) string {
	var icon string

	switch data.Weather[0].Description {
	case "clear":
		icon = "☀️"
	case "cloudy":
		icon = "⛅️"
	case "heavy rain":
		icon = "💦"
	case "heavy snow":
		icon = "⛄️"
	case "light intensity shower rain":
		icon = "☔️"
	case "light rain":
		icon = "🌦"
	case "light snow":
		icon = "🌨"
	case "moderate rain":
		icon = "🌧"
	case "moderate snow":
		icon = "🌨"
	case "overcast":
		icon = "🌥"
	case "overcast clouds":
		icon = "🌥"
	case "partly cloudy":
		icon = "🌤"
	case "snow":
		icon = "❄️"
	case "sunny":
		icon = "☀️"
	default:
		icon = "💥"
	}

	return icon
}

func (widget *Widget) refreshedAt() string {
	return widget.RefreshedAt.Format("15:04:05")
}