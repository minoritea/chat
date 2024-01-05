import {Controller} from 'https://cdn.skypack.dev/@hotwired/stimulus@v3.2.2'
import * as Turbo from 'https://cdn.skypack.dev/@hotwired/turbo@v7.3.0'

export class MessageController extends Controller {
	static targets = ['message', 'more', 'terminal']

	connect() {
		this.#scrollToBottom()
		this.intervalId = setInterval(() => this.#loadNewMessages(), 5 * 1000)
	}

	disconnect() {
		clearInterval(this.intervalId)
	}

	moreTargetConnected(element) {
		element.remove()
		this.#loadNewMessages()
	}

	terminalTargetConnected() {
		this.isTerminal = true
	}

	messageTargetConnected() {
		if (this.scrollReachedToBottom) {
			this.#scrollToBottom()
		}
	}

	afterSubmit() {
		this.#scrollToBottom()
	}

	async scroll() {
		this.scrollReachedToBottom = this.#isLastMessageLastVisibleRow()

		if (!this.isTerminal && this.element.scrollTop == 0) {
			this.isScrollLoading = true
			await this.#loadPastMessages()
			const offsetTop = this.messageTargets.length > 0 ?
				this.messageTargets[0].offsetTop :
				0
			this.element.scrollTop = offsetTop
			this.isScrollLoading = false
		}
	}

	#isLastMessageLastVisibleRow() {
		const lastMessage = this.messageTargets.slice(-1)[0]
		if (!lastMessage) {
			return false
		}
		return this.element.scrollTop + this.element.clientHeight >= lastMessage.offsetTop
	}

	#scrollToBottom() {
		const bottom = this.messageTargets.slice(-1)[0]
		if (bottom) {
			bottom.scrollIntoView({block: 'end'})
		}
		this.scrollReachedToBottom = true
	}

	async #loadPastMessages() {
		if (this.messageTargets.length > 0) {
			const id = this.messageTargets[0].id
			const res = await fetch('/messages?before_id=' + id)
			const html = await res.text()
			Turbo.renderStreamMessage(html)
		} else {
			const res = await fetch('/messages')
			const html = await res.text()
			Turbo.renderStreamMessage(html)
		}
	}

	async #loadNewMessages() {
		if (this.messageTargets.length > 0) {
			const id = this.messageTargets.slice(-1)[0].id
			const res = await fetch('/messages?after_id=' + id)
			const html = await res.text()
			Turbo.renderStreamMessage(html)
		} else {
			const res = await fetch('/messages')
			const html = await res.text()
			Turbo.renderStreamMessage(html)
		}
	}
}
