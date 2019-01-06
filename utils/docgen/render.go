package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func fileWriter(path string) *os.File {
	f, err := os.OpenFile(path, os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err.Error())
	}
	return f
}

func renderAll(docs documents) {
	for cat := range Config.Categories {
		for i := range Config.Categories[cat].Templates {
			renderInnerLoop(&Config.Categories[cat].Templates[i], docs)
		}
	}

}

func renderInnerLoop(t *templates, docs documents) {
	makePath(t.OutputPath)

	for d := range docs {
		if docs[d].CategoryID == t.ref.ID {
			renderDocument(t, &docs[d], docs)
		}
	}
	renderCategory(t, docs)
}

func renderDocument(t *templates, d *document, docs documents) {
	if t.docTemplate == nil {
		panic(fmt.Sprintf("No document template loaded for %s[%d]/*", t.ref.ID, t.index))
	}

	f := fileWriter(t.DocumentPath(d))
	b := new(bytes.Buffer)

	log("Rendering document", d.DocumentID)

	err := t.docTemplate.Execute(b, t.DocumentValues(d, docs, true))
	if err != nil {
		panic(err.Error())
	}

	if len(Config.renderedDocuments[t.ref.ID]) == 0 {
		Config.renderedDocuments[t.ref.ID] = make(map[string][]string)
	}
	Config.renderedDocuments[t.ref.ID][d.DocumentID] = append(
		Config.renderedDocuments[t.ref.ID][d.DocumentID],
		b.String(),
	)

	write(f, b)
}

func renderCategory(t *templates, docs documents) {
	if t.catTemplate == nil {
		panic(fmt.Sprintf("No category template loaded for %s[%d]", t.ref.ID, t.index))
	}

	f := fileWriter(t.CategoryFilePath())
	b := new(bytes.Buffer)

	log("Rendering category", t.ref.ID)

	err := t.catTemplate.Execute(b, t.CategoryValues(docs))
	if err != nil {
		panic(err.Error())
	}

	Config.renderedCategories[t.ref.ID] = append(
		Config.renderedCategories[t.ref.ID],
		b.String(),
	)

	write(f, b)
}

func write(f io.WriteCloser, b io.Reader) {
	_, err := io.Copy(f, b)
	if err != nil {
		panic(err.Error())
	}

	err = f.Close()
	if err != nil {
		panic(err.Error())
	}
}
