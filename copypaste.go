package main

import "net/http"

func copyPaste(o Model, w http.ResponseWriter, r *http.Request) {
	p := o.(*Paste)

	//pasteBody, err := ioutil.ReadFile(pasteStore.filenameForID(p.ID))

	//if err != nil {
	//	panic(err)
	//}

	newPaste, err := pasteStore.New(false)

	if err != nil {
		 panic(err)
	}

	pasteUpdateCore(newPaste, w, r, true, p)
}
