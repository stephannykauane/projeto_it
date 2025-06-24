import { h } from 'vue'
import { QBadge } from 'quasar'

export default {
  name: 'quasar',

  setup () {
    return () => h(QBadge, {
      class: 'quasar',
      label: 'quasar'
    })
  }
}
