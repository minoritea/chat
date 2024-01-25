import {Controller} from './stimulus.js'
import * as Turbo from './turbo.js'

export class MessageController extends Controller {
	static targets = ['message', 'more', 'terminal']
	appearedMessages = []

	connect() {
		this.#scrollToBottom()
		this.intervalId = setInterval(() => this.#loadNewMessages(), 5 * 1000)
	}

	disconnect() {
		clearInterval(this.intervalId)
	}

	messageAppeared() {
		this.#visibleMessagesChanged()
	}

	messageDisappeared() {
		this.#visibleMessagesChanged()
	}

	messageTargetConnected() {
		if (this.scrollReachedToBottom) {
			this.#scrollToBottom()
		}
	}

	moreTargetConnected() {
		if (this.scrollReachedToBottom) {
			this.#scrollToBottom()
		}
	}

	afterSubmit() {
		if (this.scrollReachedToBottom) {
			this.#scrollToBottom()
		}
	}

	#visibleMessagesChanged() {
		this.appearedMessages = this.messageTargets.filter(m => m.appeared)
		this.scrollReachedToBottom = this.#isLastMessageLastVisibleRow()
		if (
			!this.hasTerminalTarget
			&& this.#isFirstMessageFirstVisibleRow()
			&& !this.loadingPastMessages
		) {
			this.#loadPastMessages()
		}

		if (this.hasMoreTarget && this.moreTarget.appeared) {
			this.moreTarget.remove()
			this.#loadNewMessages()
		}
	}

	#isLastMessageLastVisibleRow() {
		const lastMessage = this.messageTargets.slice(-1)[0],
			lastAppearedMessage = this.appearedMessages.slice(-1)[0]
		if (!lastMessage || !lastAppearedMessage) {
			return false
		}
		return lastMessage.id === lastAppearedMessage.id
	}

	#isFirstMessageFirstVisibleRow() {
		const firstMessage = this.messageTargets[0],
			firstAppearedMessage = this.appearedMessages[0]
		if (!firstMessage || !firstAppearedMessage) {
			return false
		}
		return firstMessage.id === firstAppearedMessage.id
	}

	#scrollToBottom() {
		const bottom = this.hasMoreTarget ? this.moreTarget : this.messageTargets.slice(-1)[0]
		if (bottom) {
			bottom.scrollIntoView({block: 'end'})
		}
		this.scrollReachedToBottom = true
	}

	async #loadPastMessages() {
		this.loadingPastMessages = true
		try {
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
		} catch (e) {
			throw e
		} finally {
			this.loadingPastMessages = false
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
