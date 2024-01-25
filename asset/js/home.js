import './turbo.js'
import {Application} from './stimulus.js'
import {FlashController} from './flash_controller.js'
import {FormController} from './form_controller.js'
import {MessageController} from './message_controller.js'
import {ScrolledElementController} from './scrolled_element_controller.js'

const application = (globalThis.application ||= Application.start())
application.register('flash', FlashController)
application.register('form', FormController)
application.register('message', MessageController)
application.register('scrolled-element', ScrolledElementController)
