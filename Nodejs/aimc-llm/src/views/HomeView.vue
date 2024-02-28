<template>
  <div class="home">
    <KnowledgeBase />
  </div>
</template>

<script>
// @ is an alias to /src
import KnowledgeBase from '@/components/KnowledgeBase.vue'
import { NButton } from "naive-ui"
import { mapGetters } from 'vuex'
import { ref } from 'vue'
import { upload } from '@/api/test_services'
export default {
  data() {
    return {
      test: 'test',
      file: ref([])
    }
  },
  name: 'HomeView',
  computed: {
    ...mapGetters(['test_hello'])
  },
  components: {
    KnowledgeBase
  },
  mounted() {
    //this.$store.dispatch('test_services/getHello').then((res) => {

    //}).catch(() => {
    //  console.log('[HomeView]: 获取hello失败！')
    // })


  },
  methods: {
    submitUpload() {
      // const uploadRef = ref(null);
      // uploadRef.value?.submit();
      console.log(this.file[0])
      const formData = new FormData();
      formData.append("file", this.file[0].file);
      upload(formData).then(response => {
        console.log(response.data)
      })
    },
    handleChange(data) {
      this.file = data.fileList
    }
  }
}
</script>
