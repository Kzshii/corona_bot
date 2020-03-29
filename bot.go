package main

import (
	"bytes"
	"fmt"
	"log"
	"sort"
	"strings"
	"time"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"

	tb "gopkg.in/tucnak/telebot.v2"
)

func initBot(b *tb.Bot) {
	location, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/global", func(m *tb.Message) {
		data := getData()

		fmt.Println(m.Sender.Username)
		t, _ := time.Parse(time.RFC3339, data.LastUpdated)
		t = t.In(location)

		formatted := fmt.Sprintf("Casos confirmados: %s\nVitimas fatais: %s\nCasos recuperados: %s\nUltima atualização: %s",
			formatCommas(data.TotalConfirmed), formatCommas(data.TotalDeaths), formatCommas(data.TotalRecovered), t.Format(time.RFC1123))

		b.Send(m.Chat, formatted)
	})

	b.Handle("/list", func(m *tb.Message) {
		var list []string
		var buffer bytes.Buffer
		data := getData()

		fmt.Println(m.Sender.Username)

		listSize := fmt.Sprintf("São ao todo %d países\nAqui vai a lista para você pesquisar:", len(data.Areas))
		b.Send(m.Chat, listSize)

		for _, c := range data.Areas {
			list = append(list, c.DisplayName+"\n")
		}

		sort.Strings(list)

		for _, c := range list {
			buffer.WriteString(c)
		}

		b.Send(m.Chat, buffer.String())
	})

	b.Handle(tb.OnText, func(m *tb.Message) {
		if m.Text[0:9] == "/teixeira" {
			b.Send(m.Chat, "caraio, sdds de mainha, sdds de painho, sdds acarajé, sdds fumar um hollywood, sdds vitorinha")
			return
		}

		if m.Text[0:7] == "/search" {
			b.Send(m.Chat, "Procurando....")
			fmt.Println(m.Sender.Username, m.Payload)

			data := getData()
			t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)

			for _, c := range data.Areas {
				normPayload, _, _ := transform.String(t, m.Payload)
				normName, _, _ := transform.String(t, c.DisplayName)
				if strings.ToLower(normName) == strings.ToLower(normPayload) {

					t, _ := time.Parse(time.RFC3339, c.LastUpdated)
					t = t.In(location)

					formatted := fmt.Sprintf("País: %s\nCasos confirmados: %s\nVitimas fatais: %s\nCasos recuperados: %s\nUltima atualização: %s",
						c.DisplayName, formatCommas(c.TotalConfirmed), formatCommas(c.TotalDeaths), formatCommas(c.TotalRecovered), t.Format(time.RFC1123))

					b.Send(m.Chat, formatted)
					return
				}
			}
			b.Send(m.Chat, "Nenhum resultado encontrado")
		}
	})

	b.Start()
}
