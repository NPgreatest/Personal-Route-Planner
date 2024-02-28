<template>
    <h1>使用应用</h1>
</template>

<script>
// @ is an alias to /src
import { mapGetters } from 'vuex'
import { NDivider } from "naive-ui"

export default {
  data() {
    return {
      test: ''
    }
  },
  computed: {
    ...mapGetters(['app'])
  },
  name: 'AppUse',

  components: {
  },
  async mounted() {
    const _id = this.$route.params._id;
    await this.$store.dispatch('getApp/getApp',{_id:_id}).then((res) => {

    }).catch(() => {
      console.log('[HomeView]: 获取hello失败！')
    })
    console.log(this.app)
    //window.open('https://www.baidu.com/', '_self');
  }
}
</script>
