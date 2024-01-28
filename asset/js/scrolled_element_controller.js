import {IntersectionController} from './stimulus-use.js'

export class ScrolledElementController extends IntersectionController {
	static outlets = ['message']

	appear() {
		this.element.appeared = true
		this.messageOutlet.visibleMessagesChanged()
	}

	disappear() {
		this.element.appeared = false
		this.messageOutlet.visibleMessagesChanged()
	}
}
