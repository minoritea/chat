import {IntersectionController} from './stimulus-use.js'

export class ScrolledElementController extends IntersectionController {
	appear() {
		this.element.appeared = true
		this.dispatch('appeared')
	}

	disappear() {
		this.element.appeared = false
		this.dispatch('disappeared')
	}
}
