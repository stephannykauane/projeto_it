import { boot } from 'quasar/wrappers'
import VuePlugin from 'quasar-ui-quasar'

export default boot(({ app }) => {
  app.use(VuePlugin)
})
