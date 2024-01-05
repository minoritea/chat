import {Controller} from 'https://cdn.skypack.dev/@hotwired/stimulus@v3.2.2'

export class FormController extends Controller {
	static targets = ['reset', 'textInput']

	connect() {
		this.textInputTarget.focus()
	}

	reset() {
		this.resetTarget.click()
	}
}
