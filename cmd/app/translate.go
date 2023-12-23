package main

func (app *application) getTranslation(text string, args ...string) string {
	translated, err := app.translator.Translate(text, args...)
	if err != nil {
		app.logger.Error(err.Error())
		translated = "`Some error occured during translation`"
	}
	return translated
}
