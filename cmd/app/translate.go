package main

func (app *application) getTranslation(text string) string {
	translated, err := app.translator.Translate(text)
	if err != nil {
		app.logger.Error(err.Error())
		translated = "`Some error occured during translation`"
	}
	return translated
}
