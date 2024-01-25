import './turbo.js'
import {Application} from './stimulus.js'
import {FlashController} from './flash_controller.js'

const application = (globalThis.application ||= Application.start())
application.register('flash', FlashController)
