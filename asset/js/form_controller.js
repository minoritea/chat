import {Controller} from './stimulus.js'

export class FormController extends Controller {
	static targets = ['reset', 'textInput']

	connect() {
		document.documentElement.clientWidth >= 684 &&
			this.textInputTarget.focus()
	}

	reset() {
		this.resetTarget.click()
	}
}
