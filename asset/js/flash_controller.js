import {Controller} from 'https://cdn.skypack.dev/@hotwired/stimulus@v3.2.2'

export class FlashController extends Controller {
	static values = {index: Number}

	connect() {
		const index = this.indexValue + 1
		setTimeout(() => this.element.remove(), index * 10 * 1000)
	}

	close() {
		this.element.remove()
	}
}
